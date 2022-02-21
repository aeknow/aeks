{{ define "footer" }}
  <footer class="footer">
	<div class="container footer__container flex">		
		<div class="footer__copyright">
			&copy; 2022 AEKs.chain. Powered by <a href="https://www.aeternity.com" target="_blank">Aeternity</a>, <a href="https://ipfs.io" target="_blank">IPFS</a>, <a href="https://www.aeknow.org" target="_blank">AEKnow</a>
			<span class="footer__copyright-credits">The theme is forked from <a href="https://github.com/Vimux/Mainroad/" target="_blank">Mainroad</a> theme.</span>
		</div>
	</div>
</footer>

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
                        base: '/views/layim/dist/js/'
                    }).extend({
                        //socket: 'socket',
                        // contextmenu:'contextMenu',
                    });

                    layui.use(['layim', 'contextmenu'], function(layim) {
                        var menu = layui.contextmenu;

                        //基础配置
                        layim.config({

                            //初始化接口
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

                            //扩展工具栏
                            ,
                            tool: [{
                                alias: 'code',
                                title: '代码',
                                icon: '&#xe64e;'
                            }]

                            //,brief: true //是否简约模式（若开启则不显示主面板）

                            ,title: 'Chaet' //自定义主面板最小化时的标题
                            ,
                            right: '10px' //主面板相对浏览器右侧距离
                                //,minRight: '90px' //聊天面板最小化时相对浏览器右侧距离
                                ,
                            initSkin: '2.jpg' //1-5 设置初始背景
                                //,skin: ['aaa.jpg'] //新增皮肤
                                //,isfriend: false //是否开启好友
                                //,isgroup: false //是否开启群组
                                ,min: true //是否始终最小化主面板，默认false
                                ,
                            notice: true //是否开启桌面消息提醒，默认false
                                //,voice: false //声音提醒，默认开启，声音文件为：default.mp3

                            ,
                            msgbox: '/views/layim/dist/css/modules/layim/html/msgbox.html?aid={{.Account}}' //消息盒子页面地址，若不开启，剔除该项即可
                                ,
                            find: '/views/layim/dist/css/modules/layim/html/find.html' //发现页面地址，若不开启，剔除该项即可
                                ,
                            chatLog: '/views/layim/dist/css/modules/layim/html/chatlog.html' //聊天记录页面地址，若不开启，剔除该项即可

                        });

                        //监听在线状态的切换事件
                        layim.on('online', function(data) {
                            //console.log(data);
                        });

                        //监听签名修改
                        layim.on('sign', function(value) {
                            //console.log(value);
                        });
                        
                        

                        //监听自定义工具栏点击，以添加代码为例
                        layim.on('tool(code)', function(insert) {
                            layer.prompt({
                                title: '插入代码',
                                formType: 2,
                                shade: 0
                            }, function(text, index) {
                                layer.close(index);
                                insert('[pre class=layui-code]' + text + '[/pre]'); //将内容插入到编辑器
                            });
                        });

                        //监听layim建立就绪
                        layim.on('ready', function(res) {
                            menu.init([{
                                target: '.layim-list-friend',
                                menu: [{
                                    text: "新增分组",
                                    callback: function(target) {
                                        layer.msg(target.find('span').text());
                                    }
                                }]
                            }, {
                                target: '.layim-list-friend >li>h5>span',
                                menu: [{
                                    text: "重命名",
                                    callback: function(target) {
                                        layer.msg(target.find('span').text());
                                    }
                                }, {
                                    text: "删除分组",
                                    callback: function(target) {
                                        layer.msg(target.find('span').text());
                                    }
                                }]
                            }]);


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

                        //监听查看群员
                        layim.on('members', function(data) {
                            console.log(data);
                        });

                        //监听聊天窗口的切换
                        layim.on('chatChange', function(res) {
                            var type = res.data.type;
                            //console.log(res.data.id)
                            if (type === 'friend') {
                                //模拟标注好友状态
                                //layim.setChatStatus('<span style="color:#FF5722;">在线</span>');
                            } else if (type === 'group') {
                                //模拟系统消息
                                /*
                                layim.getMessage({
                                  system: true
                                  ,id: res.data.id
                                  ,type: "group"
                                  ,content: '模拟群员'+(Math.random()*100|0) + '加入群聊'
                                });*/
                            }
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


                    });
               
               
               
                </script>

         
{{ end }}
