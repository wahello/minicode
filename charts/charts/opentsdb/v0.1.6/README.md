opentsdb
========
Store and serve massive amounts of time series data without losing granularity.

Current chart version is `0.1.3`

Source code can be found [here](https://github.com/Gradiant/charts)

## Installing the Chart

Add gradiant helm repo:

```
helm repo add gradiant https://gradiant.github.io/charts
```

To install the chart with the release name `opentsdb`.

```
$ helm install --name opentsdb gradiant/opentsdb
```

For more detailed custom values.yaml file see `ci/custom-values.yaml`

## Chart Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://gradiant.github.io/charts | hbase | ~0.1.0 |

## Chart Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| antiAffinity | string | `"soft"` | Select antiAffinity as either hard or soft, default is 'soft' 'hard' is for production setups |
| conf | object | `{"tsd.core.auto_create_metrics":true,"tsd.core.auto_create_tagks":true,"tsd.core.auto_create_tagvs":true,"tsd.storage.hbase.zk_quorum":null}` | configure /etc/opentsdb/opentsdb.conf contents |
| daemons | int | `1` | Initial number of tsd replicas |
| env.init | object | `{"COMPRESSION":"GZ"}` | values for init container when creating hbase tables |
| env.opentsdb | string | `nil` | env values to pass to opentsdb, for example JAVA_OPTS |
| hbase | object | `{"enabled":true}` | pass vars towards hbase chart, from dependencies |
| hbase.enabled | bool | `true` | set to 'false' to disable automatically deploying dependent charts |
| hbaseConfigMapName | string | `nil` | name of the configmap used to fetch hbase configuration default hbaseConfigMapName is {{ .Release.Name}}-hbase |
| hbaseImage | object | `{"repository":"gradiant/hbase-base","tag":"2.0.1"}` | container umage used to run hbase client shell to create initial opentsdb tables |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"gradiant/opentsdb"` |  |
| image.tag | string | `"2.4.0"` |  |
| init_hbase_script | string | `"create '$UID_TABLE',\n{NAME => 'id', COMPRESSION => '$COMPRESSION', BLOOMFILTER => '$BLOOMFILTER'},\n{NAME => 'name', COMPRESSION => '$COMPRESSION', BLOOMFILTER => '$BLOOMFILTER'}\ncreate '$TSDB_TABLE',\n{NAME => 't', VERSIONS => 1, COMPRESSION => '$COMPRESSION', BLOOMFILTER => '$BLOOMFILTER'}\ncreate '$TREE_TABLE',\n{NAME => 't', VERSIONS => 1, COMPRESSION => '$COMPRESSION', BLOOMFILTER => '$BLOOMFILTER'}\ncreate '$META_TABLE',\n{NAME => 'name', COMPRESSION => '$COMPRESSION', BLOOMFILTER => '$BLOOMFILTER'}\n"` | hbase init script to create hbase tables, where $VARS are env vars from env.init (above), if empty then default will be used |
| logback | object | `{"level_conn_mgr":"WARN","level_core":"INFO","level_graph":"INFO","level_graph_handler":"WARN","level_hbase":"WARN","level_meta":"INFO","level_query":"WARN","level_querylog":"WARN","level_root":"DEBUG","level_search":"INFO","level_stats":"INFO","level_stumbleupon":"WARN","level_tools":"INFO","level_tree":"INFO","level_treshold":"WARN","level_tsd":"INFO","level_uid":"INFO","level_utils":"INFO","pattern":"%d{ISO8601} %-5level [%thread] %logger{0}: %msg%n"}` | configure /etc/opentsdb/logback.xml contents for more detailed logging |
| nodePort.enabled | bool | `false` | set to 'true' to change to nodeport instead of ClusterIP |
| nodePort.externalPort | int | `31042` | sport to use to expose service |
| port | int | `4242` | expose port when using nodePort |
| resources | object | `{"limits":{"cpu":"1000m","memory":"2048Mi"},"requests":{"cpu":"10m","memory":"256Mi"}}` | container resource requests and limits |
