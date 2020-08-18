package alertLogHandle

import (
	"crypto/sha512"
	"encoding/hex"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"shebinbin.com/alertSyslog/cmdSend"
	"shebinbin.com/alertSyslog/config"
	"shebinbin.com/alertSyslog/cpaasapi"
	"shebinbin.com/alertSyslog/mysqlconn"
	"shebinbin.com/alertSyslog/utils"
	"shebinbin.com/alertSyslog/waitdbqueue"
)

const (
	INSERT = iota
	DELETE
	UPDATE
)

type wdbNode struct {
	keyword  string
	operator byte
}

// 存放告警重复值数据
var memData map[string]int64

// 存放项目信息
var projectDataMap map[string]string

// 当数据库断开连接的时候，对数据库的操作会先暂存在内存中。
// 等数据库可以连接了，那么这些暂存的操作会全部刷进数据库中
// 布尔值意思：true的时候为insert操作，false的时候为delete操作
var wDbQueue = waitdbqueue.NewQueue()

func init() {
	projectDataMap = make(map[string]string)
	if config.Dbping {
		memData = mysqlconn.MemDataInit()
	} else {
		memData = make(map[string]int64)
	}
}

func PrintMemData() {
	logger.Info("=========内存缓存重复告警如下==========")
	for k, v := range memData {
		logger.Info(k, v)
	}
	logger.Info("==================================")
}

func AlertSyslog(alertRs gjson.Result, ip string, portNum string, nodeip string) {
	msg, instId, alertType := getMsg(alertRs, nodeip)
	if msg == "" || instId == "" {
		return
	}
	// 如果为 恢复告警
	if alertType {
		logger.Info("正在处理恢复告警...")
		// 接收到恢复告警,如果这条告警存在，那么在内存中清理该条数据，在数据库中清理该条数据

		// 如果dbping 为true，访问数据库将数据写入
		//访问数据库时，如果数据库不可用，那么将 dbping 置为false，并且将数据写入waitDbQueue[key]false。
		// 如果dbping 为false，那么将数据写入:waitDbQueue[key]false

		// 接收到恢复告警，如果这条告警不存在，查看数据库中是否有该条数据，如果有就删除，如果没有直接丢弃这条数据
		if memData[instId] != 0 {
			delete(memData, instId)
			if config.Dbping {
				ret := mysqlconn.DeleteData(instId)
				if ret != 0 {
					config.Dbping = false
					wDbQueue.Push(wdbNode{keyword: instId, operator: DELETE})
				}
			} else {
				wDbQueue.Push(wdbNode{keyword: instId, operator: DELETE})
			}
		} else {
			// 如果数据库中有，但内存中没有，就将这条告警在数据库中删除，并且发送恢复告警
			insertFlag := mysqlconn.FindData(instId)
			if insertFlag == 0 {
				// 如果数据库中有该值，则删除
				mysqlconn.DeleteData(instId)
			} else if insertFlag == 2 {
				// 如果数据库无法访问
				config.Dbping = false
				wDbQueue.Push(wdbNode{keyword: instId, operator: DELETE})
			} else {
				// 数据库中没有，内存中也没有，那就将这条够告警丢弃。
				logger.Info("系统已将" + instId + "恢复告警丢弃！")
				return
			}
		}
	} else {
		nowUtime := time.Now().Unix()
		logger.Info("正在处理异常告警...")
		// 如果为 异常告警  且 重复告警
		if memData[instId] != 0 {
			// 如果该告警时间间隔超过半个小时 1800s，则再发一次
			if nowUtime-memData[instId] >= config.AlertDuration {
				logger.Info("重复告警，已超过", config.AlertDuration, "s，再次发送该告警！")
				memData[instId] = nowUtime
				if !mysqlconn.UpdateData(instId, memData[instId]) {
					config.Dbping = false
					wDbQueue.Push(wdbNode{keyword: instId, operator: UPDATE})
				}
			} else {
				// 如果内存中有这条告警，且时间间隔不超过半个小时，那直接将重复的告警丢弃
				//logger.Info(instId)
				logger.Info("重复告警，系统已将该告警丢弃！")
				return
			}
		} else {
			// 如果内存中不存在，那就将数据写入内存，并且写入数据库
			memData[instId] = nowUtime
			ret := mysqlconn.InsertData(instId, memData[instId])
			if !ret {
				config.Dbping = false
				wDbQueue.Push(wdbNode{keyword: instId, operator: INSERT})
			}
		}
	}

	logger.Info("发往<", ip, ":", portNum, ">的syslog :", msg)

	cmdSend.Send(ip, portNum, msg, nodeip)
}

// 拼装 syslog 返回syslog 字符串
func getMsg(alertRs gjson.Result, nodeip string) (string, string, bool) {

	var alertType bool
	// 阶段阈值和value值长度
	valueStrLen := 5

	labelsObj := alertRs.Get("labels")

	alertMetaStr := utils.If(alertRs.Get("annotations.AlertMeta").Exists(),
		strings.ReplaceAll(alertRs.Get("annotations.AlertMeta").String(), `\"`, ""), "xx").(string)

	logger.Debug(alertMetaStr)

	metalabelsNs := utils.If(gjson.Valid(alertMetaStr),
		gjson.Get(alertMetaStr, `metric.queries.0.labels.#[name=="namespace"].value`),
		gjson.Result{}).(gjson.Result).String()
	metaScaleDownNs := utils.If(gjson.Valid(alertMetaStr),
		gjson.Get(alertMetaStr, `scale_down.0.namespace`),
		gjson.Result{}).(gjson.Result).String()

	metaScaleUpNs := utils.If(gjson.Valid(alertMetaStr),
		gjson.Get(alertMetaStr, `scale_up.0.namespace`),
		gjson.Result{}).(gjson.Result).String()

	metaThreshold := utils.If(gjson.Valid(alertMetaStr),
		gjson.Get(alertMetaStr, `labels.alert_threshold`),
		gjson.Result{}).(gjson.Result).String()

	var metaNamespace string

	if metalabelsNs != "" {
		metaNamespace = metalabelsNs
	} else if metaScaleDownNs != "" {
		metaNamespace = metaScaleDownNs
	} else if metaScaleUpNs != "" {
		metaNamespace = metaScaleUpNs
	}
	logger.Debug("AlertMeta中命名空间为 :", metaNamespace)

	if alertRs.Get("status").Exists() && alertRs.Get("status").String() == "resolved" {
		// 恢复告警为 true
		alertType = true
	} else {
		// 异常告警 firing && pending 为false
		alertType = false
	}

	if !labelsObj.Exists() || !alertRs.Get("status").Exists() {
		logger.Error("labelsObj or alertStatus is null!")
		return "", "", alertType
	}

	alertName := stringFromJson(labelsObj, "alert_name")

	sObKind := stringFromJson(labelsObj, "alert_involved_object_kind")

	sObCluster := stringFromJson(labelsObj, "alert_cluster")

	projectNameStr := stringFromJson(labelsObj, "alert_project")

	sNameSpace := stringFromJson(labelsObj, "alert_involved_object_namespace")

	sNameSpace2 := stringFromJson(labelsObj, "namespace")

	if sNameSpace == "" || sNameSpace == "NA" {
		if sNameSpace2 != "" && sNameSpace2 != "NA" {
			sNameSpace = sNameSpace2
		} else if metaNamespace != "" {
			sNameSpace = metaNamespace
		} else {
			sNameSpace = "NA"
		}
	}

	if projectNameStr == "" || projectNameStr == "NA" {
		projectNameStr = utils.If(strings.Contains(sNameSpace, "-"), strings.Split(sNameSpace, "-")[0], config.AlertTag).(string)
	}

	sAppName := stringFromJson(labelsObj, "application")

	sInvolObName := stringFromJson(labelsObj, "alert_involved_object_name")

	if sAppName == "" || sAppName == "NA" {
		sAppName = sInvolObName
	}

	sIndicator := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(stringFromJson(labelsObj, "alert_indicator"),
		"workload", ""), ".", ""), "utilization", "")

	// 用于去重的唯一ID
	primaryID := sObCluster + stringFromJson(labelsObj, "alertname")

	var sAlcompare string
	if !labelsObj.Get("alert_compare").Exists() {
		sAlcompare = "~"
	} else {
		sAlcompare = labelsObj.Get("alert_compare").String()
	}

	// 获取json中的阈值，有可能为空。
	althreshold := labelsObj.Get("alert_threshold")
	var sAlthreshold string
	if !althreshold.Exists() {
		sAlthreshold = "未知"
	} else {
		if len(althreshold.String()) >= valueStrLen {
			sAlthreshold = althreshold.String()[:valueStrLen]
		} else {
			sAlthreshold = althreshold.String()
		}
	}

	if metaThreshold != "" {
		if len(metaThreshold) >= valueStrLen {
			sAlthreshold = metaThreshold[:valueStrLen]
		} else {
			sAlthreshold = metaThreshold
		}
	}

	alertunit := labelsObj.Get("alert_unit")
	var sAlertunit string
	if alertunit.Exists() {
		sAlertunit = alertunit.String()
	}

	// 当前值
	annotationsObj := alertRs.Get("annotations")
	var acValue string
	if annotationsObj.Exists() && annotationsObj.Get("AlertCurrentValue").Exists() {
		tmpStr := annotationsObj.Get("AlertCurrentValue").String()
		if len(tmpStr) >= valueStrLen {
			acValue = tmpStr[:valueStrLen]
		} else {
			acValue = tmpStr
		}
	} else {
		acValue = "未知"
	}

	// 1、事件集成系统名称
	setEventIntegrationName := config.EINAME

	// 获取json中的项目名，有可能为空。

	// 2、业务系统名称
	setBusinessEventName := config.BENAME

	// 3、管理机构
	setManagementOrg := "1001"
	// 4、所属机构
	setSubOrg := "1001"
	// 5、节点IP
	setNodeIP := config.MonitorIP
	//setNodeIP := config.NodeIP
	// 6、节点名称，设置为NA
	setNodeName := config.NodeName

	// 7、事件名称
	setEventName := "(" + config.IsRecover + ")" +
		sIndicator + sAlcompare + sAlthreshold + sAlertunit
	// + "__" + labelsObj.Get("alert_name").String()

	// 8、实例ID
	msgInfo := "告警:" + alertName + ",类型:" + sObKind +
		",集群:" + sObCluster + ",项目:" + projectNameStr + ",NS:" + sNameSpace +
		",应用:" + sAppName + ",对象:" + sInvolObName + ",指标:" +
		sIndicator + sAlcompare + sAlthreshold + sAlertunit
	msgId := sAlcompare + sAlthreshold + sAlertunit + primaryID + msgInfo
	msgHash := getHash(msgId)
	setInstanceID := sObCluster + msgHash[len(msgId)%7:len(msgId)%7+5] +
		msgHash[len(msgHash)/2:(len(msgHash)/2)+5] + msgHash[(len(msgHash)-5):]
	//primaryID = setInstanceID

	// 9、实例值
	setInstanceValue := "PROBLEM"

	// 10、组件类型
	setModuleType := config.ModuleType
	// 11、组件
	setModuleName := config.AppTag

	// 12、组件子类
	setModuleSubClass := sIndicator

	// 13、事件级别
	var setEventLevel string
	if alertType {
		// 恢复告警
		setEventLevel = "100"
	} else {
		if !labelsObj.Get("severity").Exists() {
			setEventLevel = "4"
		} else {
			severity := labelsObj.Get("severity").String()
			switch severity {
			case "Critical":
				setEventLevel = "1"
			case "High":
				setEventLevel = "2"
			case "Medium":
				setEventLevel = "3"
			case "Low":
				setEventLevel = "4"
			default:
				setEventLevel = "4"
			}
		}
	}

	// 业务系统名
	var bussinessEventInfo string
	proNameStrTmp := strings.ToUpper(projectNameStr)
	if "SYSTEM" == proNameStrTmp || "NA" == proNameStrTmp ||
		"" == proNameStrTmp || config.AlertTag == proNameStrTmp ||
		"KUBE" == proNameStrTmp || "ALAUDA" == proNameStrTmp {
		bussinessEventInfo = config.BENAME
	} else {
		if config.YWAlertOn {
			porjectEname := projectNameStr
			var projectZhName string
			var flag bool
			if projectDataMap[porjectEname] == "" {
				//projectZhName, flag = mysqlconn.FindProjectName(porjectEname)
				projectZhName, flag = cpaasapi.GetCpaasProjectName(config.OrgName, porjectEname)
				if flag {
					projectDataMap[porjectEname] = projectZhName
				}
			}

			if projectDataMap[porjectEname] == "" { // && !flag
				bussinessEventInfo = config.BENAME
			} else {
				bussinessEventInfo = projectDataMap[porjectEname] + "（" + strings.ToUpper(projectNameStr) + "）"
			}
		} else {
			// setBusinessEventName = "CEB-CPAAS"
			// 如果 业务告警开关设置为false，则将该告警直接丢弃
			logger.Warn("环境变量YW_ALERT被设置为", config.YWAlertOn, "，应用项目组的告警被丢弃...")
			return "", "", false
		}
	}

	// 14、事件描述
	var setEventComment string
	setEventComment = msgInfo + ",平台:" + config.AppTag + ",节点:" + nodeip + ",当前值:" + acValue + ",业务:" + bussinessEventInfo + ",时间:" + time.Now().Format("2006/1/2 15:04:05")
	/*if annotationsObj.Exists() && annotationsObj.Get("message").Exists() {
		setEventComment = strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(annotationsObj.Get("message").String(),
						"\"", ""),
					"“", ""),
				"”", ""),
			" ", "") + "指标当前值:" + acValue
	} else {
		setEventComment = config.OrgName + "-msg"
	}*/

	// 15、发生时间
	setOccurrenceTime := "0"

	msg := setEventIntegrationName + seperator +
		setBusinessEventName + seperator +
		setManagementOrg + seperator +
		setSubOrg + seperator +
		setNodeIP + seperator +
		setNodeName + seperator +
		setEventName + seperator +
		setInstanceID + seperator +
		setInstanceValue + seperator +
		setModuleType + seperator +
		setModuleName + seperator +
		setModuleSubClass + seperator +
		setEventLevel + seperator +
		setEventComment + seperator +
		setOccurrenceTime

	// 用于去重的唯一ID
	//primaryID := alertName + sObKind + sObCluster + projectNameStr + sNameSpace +
	//	sAppName + sInvolObName + sIndicator + sAlcompare + sAlthreshold + sAlertunit

	return msg, primaryID, alertType
	// return msg, getSha256(primaryID), alertType
}

func stringFromJson(labelsObj gjson.Result, str string) string {
	indicator := labelsObj.Get(str)
	if !indicator.Exists() {
		return "NA"
	}
	return indicator.String()
}

func DBdataUpdate() {
	logger.Info("正在将内存中数据持久化至MySQL中...")
	wDbQueue.PrintAll()
	alertUtime := time.Now().Unix()
	qlen := wDbQueue.Len()
	for i := 0; i < qlen; i++ {
		dbop := wDbQueue.PrePop().(wdbNode)
		switch dbop.operator {
		case INSERT:
			ret := mysqlconn.InsertData(dbop.keyword, alertUtime)
			if !ret {
				// 数据库访问失败，将数据库状态标志位置位false,中断内存数据导入数据库，并退出循环
				config.Dbping = false
				break
			} else {
				// 如果数据库中有这条记录或者成功插入，则将waitDbQueue中的该条内容清空
				wDbQueue.PrePopClear()
			}
		case DELETE:
			ret := mysqlconn.DeleteData(dbop.keyword)
			if ret != 0 {
				// 数据库访问失败，将数据库状态标志位置位false，并退出循环
				config.Dbping = false
				break
			} else {
				// 如果数据库中该条数据被成功删除，则将waitDbQueue中的该条内容清空
				wDbQueue.PrePopClear()
			}
		case UPDATE:
			if mysqlconn.UpdateData(dbop.keyword, alertUtime) {
				// 数据库访问失败，将数据库状态标志位置位false,并退出循环
				config.Dbping = false
				break
			} else {
				// 如果数据库中该条数据被成功删除，则将waitDbQueue中的该条内容清空
				wDbQueue.PrePopClear()
			}

		}
	}
	if !config.Dbping {
		logger.Error("数据库仍然处于断开状态，内存数据无法存入MySQL中...")
	} else {
		logger.Info("持久化操作完成...")
	}
}

func getHash(msg string) string {
	hash := sha512.New()
	hash.Write([]byte(msg))
	bytes := hash.Sum([]byte("sha*-mud"))
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}
