<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>æKnow - Settings</title>
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
  {{ template "sidebar" .}}

  <!-- Content Wrapper. Contains page content -->
	<div class="content-wrapper">
<div class="col-md-9">	
<section class="content">		
		    <div class="box box-warning">
            <div class="box-header with-border">
              <h3 class="box-title">Haeme Settings</h3>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <form role="form"  action="/savesitesetting" method="POST" >
                <!-- text input -->
                
                <div class="form-group">
                  <label>AENS</label>
                  <input type="text" name="aens" class="form-control" value="{{.AENS}}">
                </div>
                
                
                <div class="form-group">
                  <label>Title</label>
                  <input type="text" name="title" class="form-control" value="{{.Title}}">
                </div>
                
                <div class="form-group">
                  <label>Subtitle</label>
                  <input type="text" name="subtitle" class="form-control" value="{{.Subtitle}}">
                </div>
                
               <!-- textarea -->
                <div class="form-group">
                  <label>Site Description</label>
                  <textarea class="form-control" name="sitedescription" rows="3" >{{.Description}}</textarea>
                </div>
              
               <div class="form-group">
                  <label>Author</label>
                  <input type="text" name="author" class="form-control" value="{{.Author}}">
                </div>  
              
               <div class="form-group">
                  <label>Author Description</label>
                  <textarea class="form-control" name="authordescription" rows="3" >{{.AuthorDescription}}</textarea>
                </div>
                
              
                <!-- select -->
                <div class="form-group">
                  <label>Theme</label>
                  <select class="form-control" name="theme">
                    <option>aeknow</option>                   
                  </select>
                </div>
               
                 <div class="box-footer">
                <button type="submit" class="btn btn-primary">Save</button>
              </div>
              </form>
            </div>
            <!-- /.box-body -->
          </div>
          <!-- /.box --> 
 	</section>	
</div>

</div>
</div>
</body>
</html>
