<html>

<head>
	<title>{{ .Title}}</title>
</head>

<body style="padding:10px;">
	<input type="text" id="messageTxt" />
	<button type="button" id="sendBtn">Send</button>
	<div id="messages" style="width: 375px;margin:10 0 0 0px;border-top: 1px solid black;">
	</div>

	<script type="text/javascript">
		var HOST = {{.Host }}
	</script>
	<script src="/views/js/chat.js" type="text/javascript"></script>
</body>

</html>
