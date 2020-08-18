// ==UserScript==
// @name           预览
// @namespace      http://tampermonkey.net/
// @version        0.1
// @description    detect code and view the cover
// @author         xshrim
// @include        *
// @grant          GM_setValue
// @grant          GM_getValue
// @grant          GM_setClipboard
// @grant          unsafeWindow
// @grant          window.close
// @grant          window.focus
// @grant          GM_log
// @grant          GM_addStyle
// @grant          GM_xmlhttpRequest
// @grant          GM_getResourceText
// @require        http://code.jquery.com/jquery-2.1.1.min.js

// ==/UserScript==

// tampermonkey的脚本和脚本所运行的目标页面是不在一个scope中的， tampermonkey的脚本运行在一个隔离的沙箱中, 因此tampermonkey中定义的变量和函数在目标页面中是不可见的（即在浏览器console中不存在）。
// 为了将tampermonkey中定义的变量和函数真正注入到目标页面中， 需要使用unsafeWindow。 默认的window对象处于脚本自己的scope中， 而unsafeWindow对象则处于目标页面的scope中。
// 使用unsafeWindow需要在头部grant unsafeWindow。 不同scope中的对象是不能直接互操作的。
//var $ = unsafeWindow.jQuery;   // 注入目标页， 仅在目标页中有效
var $ = window.jQuery;     // 仅在当前tampermonkey脚本中有效
// unsafeWindow.on = true;  //表示变量on被注入目标页， 可以在目标页中直接使用
// cookie
/*
function getCookie(name) {
  var nameEQ = name + "=";
  var ca = document.cookie.split(';');
  for(var i=0;i < ca.length;i++) {
    var c = ca[i];
    while (c.charAt(0)==' ') c = c.substring(1,c.length);
    if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
  }
  return null;
}

function setCookie(c_name, value, expiredays) {
    var exdate = new Date();
    exdate.setDate(exdate.getDate()+expiredays);
    document.cookie = c_name + "=" + escape(value) + ((expiredays==null) ?
        "" :
        ";expires="+exdate.toUTCString() + ";path=/");
}

function delete_cookie( name ) {
      document.cookie = name + '=; path=/; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}
*/

function getSelectedText(){
    if(document.Selection){
        //ie浏览器
        return document.selection.createRange().text;
    }else{
        //标准浏览器
        return window.getSelection().toString();
    }
}

function getWordAtPoint(x, y) {
  var range = document.caretRangeFromPoint(x, y);

  if (range.startContainer.nodeType === Node.TEXT_NODE) {
    range.expand('word');
    return range.toString().trim();
  }

  return null;
}

function mousePosition(ev){
    ev = ev || window.event;
    if(ev.pageX || ev.pageY){
        return {x:ev.pageX, y:ev.pageY};
    }
    return {
        x:ev.clientX + document.body.scrollLeft - document.body.clientLeft,
        y:ev.clientY + document.body.scrollTop - document.body.clientTop
    };
}

function getVideoCode(title){
    /*
    var t = title.match(/[A-Za-z]+\-\d+/g);
    if(!t){
        t = title.match(/heyzo[\-\_]?\d{4}/g);
    }
    if(!t){
        t = title.match(/\d{6}[\-\_]\d{3}/g);
    }
    if(!t){
        t = title.match(/[A-Za-z]+\d+/g);
    }
    return t;
    */
    return title.match(/([A-Za-z0-9]+[\-\_]\d+)|(heyzo[\-\_]?\d{4})|(\d{6}[\-\_]\d{3})|([A-Za-z]+\d+)|(\d{5}[\-\_]\d{4})|(\d{5}[\-\_]\d{3})/g);
}

function getVideoInfo(id, content){
    var contenttitle = $(content).find("div")[0];
    var contentimg = $(content).find("img")[0];
    console.log(contenttitle);
    console.log(contentimg);
    $("#imgpoptitletext").html($("#imgpoptitletext").html() + " " + id);
    GM_xmlhttpRequest({
        method: "GET",
        url: "https://avmoo.asia/cn/search/" + id,
        onload: xhr => {
            var xhr_data = $(xhr.responseText);

            if(!(xhr_data.find("div.alert").length)){
                var title = xhr_data.find("div.photo-info span").html();
                if (title !== undefined) {
                    $(contenttitle).html("<h4>" + title + "</h4>");
                    //$("#imgpopcontenttitle").html("<h4>" + title + "</h4>");
                }
                var img_url = xhr_data.find("div.photo-frame img").attr("src");
                if (img_url !== undefined) {
                    $(contentimg).attr("src", img_url.replace("ps.j","pl.j"));
                    //$("#imgpopcontentimg").attr("src", img_url.replace("ps.j","pl.j"));
                }
            }else{
                getUncensored(id, content);
            }
        }
    })
}

function getUncensored(id, content){
    var contenttitle = $(content).find("div")[0];
    var contentimg = $(content).find("img")[0];
    $("#imgpoptitletext").html($("#imgpoptitletext").html() + " " + id);
    GM_xmlhttpRequest({
        method: "GET",
        url: "https://avsox.asia/cn/search/" + id,
        onload: xhr => {
            var xhr_data = $(xhr.responseText);

            if(!(xhr_data.find("div.alert").length)){
                var title = xhr_data.find("div.photo-info span").html();
                $(contenttitle).html("<h4>" + title + "</h4>");
                //$("#imgpopcontenttitle").html("<h4>" + title + "</h4>");
                var details_url = xhr_data.find("a.movie-box").attr("href");
                if (details_url !== undefined) {
                    GM_xmlhttpRequest({
                        method: "GET",
                        url: details_url,
                        onload: temp => {
                            var img = $(temp.responseText).find("a.bigImage").attr("href");
                            $(contentimg).attr("src", img);
                             //$("#imgpopcontentimg").attr("src", img);
                        }
                    });
                }
            }
        }
    })
}

function createPopSubContent(pop, code) {
    var popcontent = document.createElement("div");
    var popcontenttitle = document.createElement("div");
    var popcontenttext = document.createElement("textarea");
    var popcontentimg = document.createElement("img");

    // popcontent.id="imgpopcontent";
    // popcontenttitle.id="imgpopcontenttitle";
    // popcontentimg.id="imgpopcontentimg";
    // popcontenttext.id = "imgpopcontenttext";

    popcontenttitle.style.cssText = "max-width:800px;";
    popcontenttext.style.cssText = "position: absolute;top: 0;left: 0;opacity: 0;z-index: -10;";

    popcontent.appendChild(popcontenttitle);
    popcontent.appendChild(popcontentimg);
    popcontent.appendChild(popcontenttext);
    pop.appendChild(popcontent);

    popcontentimg.onmousedown = function(event){
        if (event.ctrlKey) {
             window.open("https://www.javbus.com/" + code); // poptitle.innerText
        } else {
            popcontenttext.value = popcontenttitle.innerText;
            popcontenttext.select();
            document.execCommand("copy");
            popcontenttitle.style.cssText = "max-width:800px;background:green;";
        }
    }

    popcontenttitle.onmousedown = function(event) {
         if (!event.ctrlKey) {
            return
        }
         window.open("https://www.zhongzilou.com/list/" + code + "/1"); // poptitle.innerText
    }

    return popcontent;
}

function createPop(left, top, removeable) {
    $("#imgpop").remove();

    var pop = document.createElement("div");
    var poptitle = document.createElement("div");
    var poptitletext = document.createElement("span");
    var poptitlebtn = document.createElement("label");


    poptitle.id="imgpoptitle";
    poptitletext.id="imgpoptitletext";
    poptitlebtn.id="imgpoptitlebtn";

    pop.id="imgpop";

    poptitle.style.cssText = "height:30px;width:100%;text-align:center;vertical-align:middle;font-size:14px;font-weight:bold;background:gray;cursor:move;";
    poptitlebtn.innerText = "X";
    poptitlebtn.style.cssText = "float:right";


    //popcontenttitle.style.cssText = "height:30px;width:100%;background:red;";
    pop.style.cssText = "position:absolute;left:" + left + "px;top:" + top + "px;background:#f0f0f0;z-index:101;border:solid 2px #afccfe;";

    poptitle.appendChild(poptitletext);
    poptitle.appendChild(poptitlebtn);
    pop.appendChild(poptitle);

    poptitlebtn.onclick = function(event){
             $("#imgpop").remove();
    }

    if(removeable == true){
			var ismousedown = false;
			var popleft,poptop;
			var downX,downY;
			popleft = parseInt(pop.style.left);
			poptop = parseInt(pop.style.top);
			poptitle.onmousedown = function(e){
                ismousedown = true;
                downX = e.clientX;
                downY = e.clientY;
			}
			poptitle.onmousemove = function(e){
				if(ismousedown){
                    pop.style.top = e.clientY - downY + poptop + "px";
                    pop.style.left = e.clientX - downX + popleft + "px";
				}
			}
			/*松开鼠标时要重新计算当前窗口的位置*/
			poptitle.onmouseup = function(){
				popleft = parseInt(pop.style.left);
				poptop = parseInt(pop.style.top);
				ismousedown = false;
			}
		}

    return pop;
}

function show(text, pos) {
    if (text == undefined || text == "") {
        return;
    }
    var code = getVideoCode(text);
    if (code !== null && code !== undefined) {
        var pop = createPop(pos.x, pos.y, true);
        document.body.appendChild(pop);
        $.each(code ,function(index,value){
            var popcontent = createPopSubContent(pop, value);
            console.log("=====  ", value, "  =====");
            getVideoInfo(value, popcontent);
        });
    }
}

(function() {
    'use strict';

    document.body.onclick = function(event){
        //if (event.target.id !== "imgpop" && event.target.id !== "imgpoptitle" && event.target.id !== "imgpopcontent" && event.target.parentNode.id !== "imgpop" && event.target.parentNode.id !== "imgpoptitle" && event.target.parentNode.id !== "imgpopcontent") {
         if ($("#imgpop") !== undefined) {
             if (event.target.id != "imgpop" && event.target.parentNode.id != "imgpop" && event.target.parentNode.parentNode.id != "imgpop" && event.target.parentNode.parentNode.parentNode.id != "imgpop" && event.target.parentNode.parentNode.parentNode.parentNode.id != "imgpop") {
            //if ($("#imgpop") != event.target && !$.contains($("#imgpop"), event.target)) {
                $("#imgpop").remove();
            }
        }
    }

    // 三种获取特定文本的方式
    /* 通过选中文本获取
    document.body.onmouseup = function(event){
         if (!event.altKey) {
            return
        }

         var pos = mousePosition(event)
         var text = getSelectedText()
         show(text, pos);
	};
    */

    /* 通过鼠标所在位置单词获取
    document.body.onmousemove = function(event){
        if (!event.altKey) {
            return
        }

        var pos = mousePosition(event)
        var text = getWordAtPoint(event.x, event.y)
         show(text, pos);
    }
    */

    // 通过鼠标所在控件文本获取
    document.body.onmouseover = function(event){
        if (!event.altKey) {
            return
        }

        //console.log('当前鼠标在', el, '元素上');//在控制台中打印该变量

        var pos = mousePosition(event)

       if (event.target.children.length <= 1) {
           var text = event.target.innerText;
           show(text, pos);
       }
    }

})();