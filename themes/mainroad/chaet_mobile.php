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
  
  //演示自动回复
  var autoReplay = [
    '您好，我现在有事不在，一会再和您联系。', 
    '你没发错吧？face[微笑] ',
    '洗澡中，请勿打扰，偷窥请购票，个体四十，团体八折，订票电话：一般人我不告诉他！face[哈哈] ',
    '你好，我是主人的美女秘书，有什么事就跟我说吧，等他回来我会转告他的。face[心] face[心] face[心] ',
    'face[威武] face[威武] face[威武] face[威武] ',
    '<（@￣︶￣@）>',
    '你要和我说话？你真的要和我说话？你确定自己想说吗？你一定非说不可吗？那你说吧，这是自动回复。',
    'face[黑线]  你慢慢说，别急……',
    '(*^__^*) face[嘻嘻] ，是贤心吗？'
  ];
  
  layim.config({
    
    
    //上传图片接口
    uploadImage: {
      url: '/upload/image' //（返回的数据格式见下文）
      ,type: '' //默认post
    }
    
    //上传文件接口
    ,uploadFile: {
      url: '/upload/file' //（返回的数据格式见下文）
      ,type: '' //默认post
    }
    
    //,brief: true
 //初始化接口
                            ,init: {
                                "mine": {
      "username": "Liu Yang"
      ,"id": "ak_bKVvB7iFJKuzH6EvpzLfWKFUpG3qFxUvj8eGwdkFEb7TCTwP8"
      ,"status": "online"
      ,"sign": "From the blockchain, to the blockchain. "
      ,"avatar": "http://127.0.0.1:8080/ipfs/QmR3AmaREUvuPauo5wA1esBDckHH7BbnL7d7n5SEcvUpKY"
    }
    ,"friend": [{
      "groupname": "AENS Boss"
      ,"id": 1
      ,"online": 2
      ,"list": [{
        "username": "Caigen"
        ,"id": "ak_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5"
        ,"avatar": "http://127.0.0.1:8080/ipfs/QmR3AmaREUvuPauo5wA1esBDckHH7BbnL7d7n5SEcvUpKY"
        ,"aens":"1.chain"
        ,"sign": "The biggest boss of AENS, AENS最大的老板"
      },{
        "username": "疯子(7.chain)"
        ,"id": "108101"
        ,"avatar": "http://127.0.0.1:8080/ipfs/QmR3AmaREUvuPauo5wA1esBDckHH7BbnL7d7n5SEcvUpKY"
        ,"sign": "微电商达人"
      },{
        "username": "刘少(liu.chain)"
        ,"id": "168168"
        ,"avatar": "http://127.0.0.1:8080/ipfs/QmR3AmaREUvuPauo5wA1esBDckHH7BbnL7d7n5SEcvUpKY"
        ,"sign": "让天下没有难写的代码"
        ,"status": "offline"
      },{
        "username": "B.A.M(cryptopay.chain)"
        ,"id": "168168"
        ,"avatar": "http://tp4.sinaimg.cn/2145291155/180/5601307179/1"
        ,"sign": "让天下没有难写的代码"
        ,"status": "offline"
      },{
        "username": "Blockcity(Blockcity.chain)"
        ,"id": "1681681"
        ,"avatar": "http://tp4.sinaimg.cn/2145291155/180/5601307179/1"
        ,"sign": "让天下没有难写的代码"
        ,"status": "offline"
      },{
        "username": "善天(pay.chain)"
        ,"id": "1681683"
        ,"avatar": "http://tp4.sinaimg.cn/2145291155/180/5601307179/1"
        ,"sign": "让天下没有难写的代码"
        ,"status": "offline"
      },{
        "username": "冠头(guantoulaoshi.chain)"
        ,"id": "1681683"
        ,"avatar": "http://tp4.sinaimg.cn/2145291155/180/5601307179/1"
        ,"sign": "让天下没有难写的代码"
        ,"status": "offline"
      },{
        "username": "AE86(ae86.chain)"
        ,"id": "1681685"
        ,"avatar": "http://tp4.sinaimg.cn/2145291155/180/5601307179/1"
        ,"sign": "让天下没有难写的代码"
        ,"status": "offline"
      },{
        "username": "Neverland"
        ,"id": "6666665"
        ,"avatar": "http://tp2.sinaimg.cn/1783286485/180/5677568891/1"
        ,"sign": "代码在囧途，也要写到底"
      }]
    },{
      "groupname": "AE大佬分组"
      ,"id": 2
      ,"online": 3
      ,"list": [{
        "username": "罗玉凤"
        ,"id": "121286"
        ,"avatar": "http://tp1.sinaimg.cn/1241679004/180/5743814375/0"
        ,"sign": "在自己实力不济的时候，不要去相信什么媒体和记者。他们不是善良的人，有时候候他们的采访对当事人而言就是陷阱"
      },{
        "username": "长泽梓Azusa"
        ,"id": "100001222"
        ,"sign": "我是日本女艺人长泽あずさ"
        ,"avatar": "http://tva1.sinaimg.cn/crop.0.0.180.180.180/86b15b6cjw1e8qgp5bmzyj2050050aa8.jpg"
      },{
        "username": "大鱼_MsYuyu"
        ,"id": "12123454"
        ,"avatar": "http://tp1.sinaimg.cn/5286730964/50/5745125631/0"
        ,"sign": "我瘋了！這也太準了吧  超級笑點低"
      },{
        "username": "谢楠"
        ,"id": "10034001"
        ,"avatar": "http://tp4.sinaimg.cn/1665074831/180/5617130952/0"
        ,"sign": ""
      },{
        "username": "柏雪近在它香"
        ,"id": "3435343"
        ,"avatar": "http://tp2.sinaimg.cn/2518326245/180/5636099025/0"
        ,"sign": ""
      }]
    },{
      "groupname": "AE大神"
      ,"id": 3
      ,"online": 1
      ,"list": [{
        "username": "林心如"
        ,"id": "76543"
        ,"avatar": "http://tp3.sinaimg.cn/1223762662/180/5741707953/0"
        ,"sign": "我爱贤心"
      },{
        "username": "佟丽娅"
        ,"id": "4803920"
        ,"avatar": "http://tp4.sinaimg.cn/1345566427/180/5730976522/0"
        ,"sign": "我也爱贤心吖吖啊"
      }]
    }]
    ,"group": [{
      "groupname": "AENS群(aens.liuyang.chain)"
      ,"id": "group_bKVvB7iFJKuzH6EvpzLfWKFUpG3qFxUvj8eGwdkFEb7TCTwP8_1"
      ,"avatar": "http://tp2.sinaimg.cn/2211874245/180/40050524279/0"
    },{
      "groupname": "AE学习交流"
      ,"id": "group_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5_2"
      ,"avatar": "http://tp2.sinaimg.cn/5488749285/50/5719808192/1"
    }]
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
