<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>Editor</title>
    <link rel="stylesheet" href="/views/editor.md/examples/css/style.css" />
       <link rel="stylesheet" href="/views/editor.md/css/editormd.css" />
  <!-- Tell the browser to be responsive to screen width -->
  <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
  <!-- Bootstrap 3.3.7 -->
  <link rel="stylesheet" href="/views/static/bower_components/bootstrap/dist/css/bootstrap.min.css">
  <!-- Font Awesome -->
  <link rel="stylesheet" href="/views/static/bower_components/font-awesome/css/font-awesome.min.css">
  <!-- Ionicons -->
  <link rel="stylesheet" href="/views/static/bower_components/Ionicons/css/ionicons.min.css">
  <!-- Theme style -->
  <link rel="stylesheet" href="/views/static/dist/css/AdminLTE.css">
  <!-- AdminLTE Skins. Choose a skin from the css/skins
       folder instead of downloading all of them to reduce the load. -->
 <link rel="stylesheet" href="/views/static/dist/css/skins/skin.css">
  <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
  <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
  <!--[if lt IE 9]>
  <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
  <![endif]-->

	<meta http-equiv="pragma" content="no-cache">
	<!-- HTTP 1.0 -->
	<meta http-equiv="cache-control" content="no-cache">
	<!-- Prevent caching at the proxy server -->
	<meta http-equiv="expires" content="0">

  <!-- Google Font -->
</head>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
<header class="main-header">
    <!-- Logo -->
    <a href="/" class="logo"  style="background:#f7296e">
      <!-- mini logo for sidebar mini 50x50 pixels -->
      <span class="logo-mini"><b>AE</b>KS</span>
      <!-- logo for regular state and mobile devices -->
      <span class="logo-lg"><b>AEKS</b>.chain</span>
    </a>
    <!-- Header Navbar: style can be found in header.less -->
    <nav class="navbar navbar-static-top"  style="background:#f7296e">
      <!-- Sidebar toggle button-->
      <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
        <span class="sr-only">Toggle navigation</span>
      </a>

      <div class="navbar-custom-menu" >
        <ul class="nav navbar-nav">
       <!-- User Account: style can be found in dropdown.less -->
          <li class="dropdown user user-menu">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
              <img src="/views/static/dist/img/ae.png" class="user-image" alt="User Image">
              <span class="hidden-xs">{{.Account}}</span>
            </a>
           <ul class="dropdown-menu">
              <!-- User image -->
              <li class="user-header"  style="height:60px;background-color:wihte;">
                <a href="/logout">Logout</a>
              </li>
              </ul>
          </li>

        </ul>
      </div>
    </nav>
  </header>
 <!-- Left side column. contains the logo and sidebar -->
	  <!-- Left side column. contains the sidebar -->
<div style="text-align:left">
  {{ template "sidebar" .}}
</div>
  <!-- Content Wrapper. Contains page content -->
	<div class="content-wrapper">
	
<section class="content">		
		    <div class="box box-warning">
            <div class="box-header with-border">
              <h3 class="box-title">Writing</h3>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
<form class="form-horizontal" action="/saveblog" method="POST" >
<div style="text-align:left;margin-left:10px;">

<div class="input-group input-group-lg"  style="margin-top:5px;;">
                <div class="input-group-btn">
                  <button type="button" class="btn btn-danger">Title</button>
                </div>
                <!-- /btn-group -->
                <input type="text" class="form-control"  name="title" placeholder="title">
              </div>
 <div class="input-group input-group-lg"  style="margin-top:5px;;" >
                <div class="input-group-btn">
                  <button type="button" class="btn btn-danger">Keywords</button>
                </div>
                <!-- /btn-group -->
                <input type="text" class="form-control"  name="keywords" placeholder="Keywords of the article" >
              </div>
    <div class="input-group input-group-lg"  style="margin-top:5px;;" >
                <div class="input-group-btn">
                  <button type="button" class="btn btn-danger">Categories</button>
                </div>
                <!-- /btn-group -->
                <input type="text" class="form-control"  name="tags" placeholder="Categories"  >
              </div>           
              
<div class="form-group " style="margin-left:5px;margin-top:5px;;" >                 
                  <div class="input-group-btn">
                  <button type="button" class="btn btn-danger">Abstract</button>
                </div>
                  <textarea class="form-control" rows="4" name="description"  placeholder="Abstract ..."></textarea>
                </div>


</div>

 <input type="hidden" name="editpath" value="{{.EditPath}}">
		 <div id="test-editormd">
                <textarea style="display:none;" name="content"></textarea>
            </div>

        <button type="submit" class="btn btn-success pull-left" style="background-color:green;color:white">Post</button>
  <script src="/views/editor.md/examples/js/jquery.min.js"></script>
        <script src="/views/editor.md/editormd.min.js"></script>      
        <script type="text/javascript">
			var testEditor;

            $(function() {
                testEditor = editormd("test-editormd", {
                    width   : "100%",
                    height  : 640,
                    syncScrolling : "single",
                    tex:true,
                    imageUpload : true,
                    imageUploadURL : "/uploadblogimage",
                    imageFormats      : ["jpg", "jpeg", "gif", "png", "bmp", "webp","mp4","avi"],
                    codeFold : true,
                    taskList : true,                   
                    placeholder : "Enjoy Sharing! Write now...",
                    htmlDecode : true,
                    path    : "/views/editor.md/lib/"
                });
                
                /*
                // or
                testEditor = editormd({
                    id      : "test-editormd",
                    width   : "90%",
                    height  : 640,
                    path    : "/views/editor.md/lib/"
                });
                */
            });
        </script>

</form>
            
            </div>
            <!-- /.box-body -->
          </div>
          <!-- /.box --> 
 	</section>	


</div>
</div>
</body>
</html>
