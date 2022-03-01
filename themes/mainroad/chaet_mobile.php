<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, height=device-height, user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0">
<meta name="format-detection" content="telephone=no">
<title>LayIM 移动版</title>

<link rel="stylesheet" href="/views/layim/dist/css/layui.mobile.css">

</head>
<body>
<script src="/views/layim/dist/layui.js"></script>
<script src="/themes/mainroad/js/base64.js"></script>

<script>
 var myAccount = '{{.Account}}';


                    if (!/^http(s*):\/\//.test(location.href)) {
                        //alert('请部署到localhost');
                    }

                    var lockReconnect = false;
                    var socket;
                    var loc = window.location, wsUrl; 
                    if (loc.protocol === "https:"){
						 wsUrl = "wss:";
						}else{
							 wsUrl = "ws:";
							}
                    wsUrl += "//" + loc.host + "/websocket?user={{.Account}}";


                    self.setInterval("heart()", 60000);

                    function heart() {
						var timestamp=new Date().getTime();
						var heartbeatStr="ping"+timestamp;
						var httpRequest = new XMLHttpRequest();
						httpRequest.open('POST', '/signjson', true); 
						httpRequest.setRequestHeader( "Content-Type" , "application/x-www-form-urlencoded");
						//	httpRequest.send('body='+Base64.encode(JSON.stringify(data)));
						httpRequest.send('body='+heartbeatStr);
						/**
						 * 获取数据后的处理程序
						 */
						httpRequest.onreadystatechange = function () {//请求后的回调接口，可将请求成功后要执行的程序写在其中
							if (httpRequest.readyState == 4 && httpRequest.status == 200) {//验证请求是否发送成功
								var signature = httpRequest.responseText;//获取到服务端返回的数据
								//console.log(signature);	
								var mtype='ping';
								socket.send('{"Signature":"'+signature+'","Body":"'+heartbeatStr+'","Account":"{{.Account}}","Mtype":"'+mtype+'"}');
								//console.log('{"Signature":"'+signature+'","Body":"'+heartbeatStr+'","Account":"{{.Account}}"}');
							   // socket.send(JSON.stringify(data+":sig:"+json));
							}
						};
                       // socket.send('ping');
                       // console.log('ping')
                    }
                    
                    
                    //console.log(wsUrl);
                    socket = new WebSocket(wsUrl);

                    socket.onerror = function(event) {
                        console.log('websocket服务出错了');
                        alert("Websocket Server Error");
                        //reconnect(wsUrl);
                    };
                    socket.onclose = function(event) {
                        console.log('websocket服务关闭了');
                        window.alert("Websocket Server is Closed");
                        //reconnect(wsUrl);
                    };
                    socket.onopen = function(event) {
                        //heartCheck.reset().start(); //传递信息
                        console.log("连接成功!" + new Date().toUTCString());
                        
                        socket.send('{"Signature":"none","Body":"online","Account":"{{.Account}}","Mtype":"online"}');
                        //socket.send("{{.Account}} Online")
                    };


                    //收到消息推送
                    function doWithMsg(msg) {
                        //getdjxh()//这个函数是业务上面申请列表的函数 可以忽略
                        window.external.CallFun('receiveMsg'); //这个也是
                    }
	
	
	
layui.config({
  version: true
}).use('mobile', function(){
  var mobile = layui.mobile
  ,layim = mobile.layim
  ,layer = mobile.layer;  

  layim.config({
    
   init: {
                                url: '/views/layim/sample/json/getList.json',
                                data: {}
                            }

                            //查看群员接口
                            ,
                            members: {
                                url: '/views/layim/sample/json/getMembers.json',
                                data: {}
                            }

                            //上传图片接口
                            ,
                            uploadImage: {
                                url: '/uploadimage' //（返回的数据格式见下文）
                                    ,
                                type: '' //默认post
                            }

                            //上传文件接口
                            ,
                            uploadFile: {
                                url: '/uploadfile' //（返回的数据格式见下文）
                                    ,
                                type: '' //默认post
                            }

                            ,
                            isAudio: false //开启聊天工具栏音频
                                ,
                            isVideo: false //开启聊天工具栏视频
    //扩展聊天面板工具栏
    ,tool: [{
      alias: 'code'
      ,title: '代码'
      ,iconUnicode: '&#xe64e;'
    }]
    
    //扩展更多列表
    ,moreList: [{
      alias: 'find'
      ,title: '发现'
      ,iconUnicode: '&#xe628;' //图标字体的unicode，可不填
      ,iconClass: '' //图标字体的class类名
    },{
      alias: 'share'
      ,title: '分享与邀请'
      ,iconUnicode: '&#xe641;' //图标字体的unicode，可不填
      ,iconClass: '' //图标字体的class类名
    }]
    
    //,tabIndex: 1 //用户设定初始打开的Tab项下标
    //,isNewFriend: false //是否开启“新的朋友”
    ,isgroup: true //是否开启“群聊”
    //,chatTitleColor: '#c00' //顶部Bar颜色
    ,title: 'Chaet' //应用名，默认：我的IM
  });
  
  //创建一个会话
  /*
  layim.chat({
    id: 111111
    ,name: '许闲心'
    ,type: 'kefu' //friend、group等字符，如果是group，则创建的是群聊
    ,avatar: 'http://tp1.sinaimg.cn/1571889140/180/40030060651/1'
  });
  */

  
  //监听点击“新的朋友”
  layim.on('newFriend', function(){
    layim.panel({
      title: '新的朋友' //标题
      ,tpl: '<div style="padding: 10px;">自定义模版，</div>' //模版
      ,data: { //数据
        test: '么么哒'
      }
    });
  });
  
  //查看聊天信息
  layim.on('detail', function(data){
    //console.log(data); //获取当前会话对象
    layim.panel({
      title: data.name + ' 聊天信息' //标题
      ,tpl: '<div style="padding: 10px;">自定义模版，<a href="http://www.layui.com/doc/modules/layim_mobile.html#ondetail" target="_blank">参考文档</a></div>' //模版
      ,data: { //数据
        test: '么么哒'
      }
    });
  });
  
  //监听点击更多列表
  layim.on('moreList', function(obj){
    switch(obj.alias){
      case 'find':
        layer.msg('自定义发现动作');
        
        //模拟标记“发现新动态”为已读
        layim.showNew('More', false);
        layim.showNew('find', false);
      break;
      case 'share':
        layim.panel({
          title: '邀请好友' //标题
          ,tpl: '<div style="padding: 10px;">自定义模版，</div>' //模版
          ,data: { //数据
            test: '么么哒'
          }
        });
      break;
    }
  });
  
  //监听返回
  layim.on('back', function(){
    //如果你只是弹出一个会话界面（不显示主面板），那么可通过监听返回，跳转到上一页面，如：history.back();
  });
  
  //监听自定义工具栏点击，以添加代码为例
  layim.on('tool(code)', function(insert, send){
    insert('[pre class=layui-code]123[/pre]'); //将内容插入到编辑器
    send();
  });
  
   //监听发送消息
                        layim.on('sendMessage', function(data) {
                            var To = data.to;
                            
                            data.to.timestamp=new Date().getTime();                       
                            //Get the signature of the message
							var httpRequest = new XMLHttpRequest();
							httpRequest.open('POST', '/signjson', true);						
							
							httpRequest.setRequestHeader( "Content-Type" , "application/x-www-form-urlencoded");
							httpRequest.send('body='+Base64.encode(JSON.stringify(data)));
						
							/**
							 * 获取数据后的处理程序
							 */
							httpRequest.onreadystatechange = function () {
								if (httpRequest.readyState == 4 && httpRequest.status == 200) {
									var signature = httpRequest.responseText;
									//console.log(signature);	
									if ('groupname' in data.to) {
											mtype='group';
										}else{
											mtype='private';
										}
									
									socket.send('{"Signature":"'+signature+'","Body":"'+Base64.encode(JSON.stringify(data))+'","Account":"{{.Account}}","Mtype":"'+mtype+'"}');
									console.log('{"Signature":"'+signature+'","Body":"'+Base64.encode(JSON.stringify(data))+'","Account":"{{.Account}}","Mtype":"'+mtype+'"}');
									console.log(data);
								  
								}
							};
							
							//socket.send(JSON.stringify(data));
                            //if(To.type === 'friend'){
                            //   layim.setChatStatus('<span style="color:#FF5722;">对方正在输入。。。</span>');
                            // }


                        });

 socket.onmessage = function(res) {
                           // console.log(res.data)
                            var msg=JSON.parse(res.data)
                            console.log(msg)
                            //heartCheck.reset().start();
                            
                            if (res.data.indexOf("ping")==-1){
								console.log("MSG:"+msg)
								layim.getMessage(JSON.parse(res.data)); //res.data即你发送消息传递的数据（阅读：监听发送的消息）
								}else{
									console.log("ping"+res.data)
									}

                            if(res.data != 'pong' && msg.username !='localakak'){		
                           // layim.getMessage(JSON.parse(res.data)); //res.data即你发送消息传递的数据（阅读：监听发送的消息）
                            }

                        };

  
  //监听查看更多记录
  layim.on('chatlog', function(data, ul){
    console.log(data);
    layim.panel({
      title: '与 '+ data.name +' 的聊天记录' //标题
      ,tpl: '<div style="padding: 10px;">这里是模版，</div>' //模版
      ,data: { //数据
        test: 'Hello'
      }
    });
  });
  
  //模拟"更多"有新动态
  layim.showNew('More', true);
  layim.showNew('find', true);
});
</script>
</body>
</html>
