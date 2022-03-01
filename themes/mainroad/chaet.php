<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Chaet</title>
    <style type="text/css">
        .talk_con {
            width: 100%;
            height: 100%;
            border: 1px solid #666;
            margin: 50px auto 0;
            background: #f9f9f9;
        }
        
        .talk_show {
            width: 100%;
            height: 420px;
            border: 1px solid #666;
            background: #fff;
            margin: 10px auto 0;
            overflow: auto;
        }
        
        .talk_input {
            width: 100%;
        }
        
        .talk_word {
            width: 90%;
            height: 26px;
            float: left;
            text-indent: 10px;
            margin: 2% 5%;
        }
        
        .talk_sub {
            width: 100%;
            height: 30px;
            float: left;
        }
        
        .atalk {
            margin: 10px;
        }
        
        .atalk span {
            display: inline-block;
            background: #0181cc;
            border-radius: 10px;
            color: #fff;
            padding: 5px 10px;
        }
        
        .btalk {
            margin: 10px;
            text-align: right;
        }
        
        .btalk span {
            display: inline-block;
            background: #ef8201;
            border-radius: 10px;
            color: #fff;
            padding: 5px 10px;
        }
    </style>
    <link rel="stylesheet" href="/views/layim/dist/css/layui.css">
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


</head>

<body class="hold-transition skin-blue sidebar-mini">
    <!-- Site wrapper -->
    <div class="wrapper">

        <header class="main-header">
            <!-- Logo -->
            <a href="/" class="logo" style="background:#f7296e">
                <!-- mini logo for sidebar mini 50x50 pixels -->
                <span class="logo-mini"><b>a</b>K</span>
                <!-- logo for regular state and mobile devices -->
                <span class="logo-lg"><b>ae</b>Know</span>
            </a>
            <!-- Header Navbar: style can be found in header.less -->
            <nav class="navbar navbar-static-top" style="background:#f7296e">
                <!-- Sidebar toggle button-->
                <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
                    <span class="sr-only">Toggle navigation</span>
                </a>

                <div class="navbar-custom-menu">
                    <ul class="nav navbar-nav">
                        <!-- User Account: style can be found in dropdown.less -->
                        <li class="dropdown user user-menu">
                            <a href="#">
                                <img src="/views/static/dist/img/ae.png" class="user-image" alt="User Image">
                                <span class="hidden-xs">{{.Account}}</span>
                            </a>

                        </li>

                    </ul>
                </div>
            </nav>
        </header>


        <!-- Left side column. contains the sidebar -->
        {{ template "sidebar" .}}

        <!-- Content Wrapper. Contains page content -->
        <div class="content-wrapper">
            <!-- Content Header (Page header) -->
            <section class="content-header">
                <h1>
                    Chat

                </h1>
                <ol class="breadcrumb">
                    <li>
                        <a href="#">
                            <i class="fa fa-dashboard"></i>Home</a>
                    </li>
                    <li>
                        <a href="#">Knode</a>
                    </li>
                    <li class="active">chat</li>
                </ol>

            </section>

            <!-- Main content -->
            <section class="content">

          
        </div>
        <!-- /.box -->

        </section>
        <!-- /.content -->
    </div>
    <!-- /.content-wrapper -->

    {{ template "footer" .}}


    <!-- Add the sidebar's background. This div must be placed
       immediately after the control sidebar -->
    <div class="control-sidebar-bg"></div>
    </div>
    <!-- ./wrapper -->

    <!-- jQuery 3 -->


</body>

</html>
