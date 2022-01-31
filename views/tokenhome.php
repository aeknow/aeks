
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>æKnow -Tokens</title>
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
 <link rel="dns-prefetch" href="https://www.aeknow.org">
  <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
  <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
  <!--[if lt IE 9]>
  <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
  <![endif]-->

  <!-- Google Font -->
</head>
<body class="hold-transition skin-blue sidebar-mini">
<!-- Site wrapper -->
<div class="wrapper">

<header class="main-header">
    <!-- Logo -->
    <a href="/" class="logo"  style="background:#f7296e">
      <!-- mini logo for sidebar mini 50x50 pixels -->
      <span class="logo-mini"><b>a</b>K</span>
      <!-- logo for regular state and mobile devices -->
      <span class="logo-lg"><b>ae</b>Know</span>
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
  

  <!-- Left side column. contains the sidebar -->
  {{ template "sidebar" .}}

  <!-- Content Wrapper. Contains page content -->
  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
       Tokens
        <small>Aeternity blockchain based AEX-9 token</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
        <li><a href="#">Know</a></li>
        <li class="active">Token</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">

      <!-- Default box -->
      <div class="box">
        <div class="box-header with-border">
			<div class="col-md-3">

          <!-- Profile Image -->
          <div class="box box-primary">
			  <div class="box-header with-border">
              <h3 class="box-title">Summary </h3>
            </div>
            <div class="box-body box-profile">
			<center>
              <img class="" style="with:120px;" src="/views/qr_ak.png" alt="AK address">
			</center>
              <h3 class="profile-username text-center">QRcode</h3>

              <p class="text-muted text-center">of the address</p>

              <ul class="list-group list-group-unbordered">
                <li class="list-group-item">
                  <i class="fa fa-btc margin-r-5"></i><b>AE Balance:</b> <a class="pull-right"/>{{.Balance}} AE</a>
                </li>
                
                <li class="list-group-item">
                  <i class="fa fa-calculator margin-r-5"></i><b>Nonce:</b> <a class="pull-right"/>{{.Nonce}} </a>
                 
                </li> 
                
               <center> <a href="https://www.aeknow.org/address/wallet/{{.Account}}" target="_blank">Transaction History</a></center>
                 
                 
              </ul>

            </div>
            
            <!-- /.box-body -->
          </div>
          <!-- /.box -->

        </div>
        <!-- /.col -->
       
       
       <div class="col-md-9">
		<!-- Horizontal Form -->
          <div class="box box-info">
            <div class="box-header with-border">
              <h3 class="box-title">Check contract</h3>
            </div>
            
            <div class="box-footer">
                  <form name="checktoken" action="https://www.aeknow.org/token/checktoken" method="post" target="_blank" >
                    <div class="input-group">
                      <input type="text" name="tokenname" placeholder="Check Token name or ct_ contracts infomation ..." class="form-control">
                      <span class="input-group-btn">
                            <a href='#' οnclick='doLogin()'><button class="btn btn-warning btn-flat">Check</button></a>
                          </span>
                    </div>
                  </form>
                </div>
            <!-- /.box-header -->
        <script>
function doLogin(){
	document.checktoken.submit();
}
</script>
        
        <div class="nav-tabs-custom">
            <ul class="nav nav-tabs">
              <li class="active"><a href="/token" >My Tokens</a></li>
               <li  >
				    <a href=/deploytoken><button type="button" class="btn btn-block btn-info btn-sm">Create Token</button></a> 
				  </li>
              
             <!--  <li><a href="#" data-toggle="tab">Expiring</a></li> -->
            </ul>
            <div class="tab-content1">
           <div class="tab-pane table-responsive no-padding " id="tokenlist"></div>
			</div>
			</div>
			
			
<!-- jQuery 3 -->
<script src="/views/static/bower_components/jquery/dist/jquery.min.js"></script>
<script>
$.get("https://www.aeknow.org/api/tokenlist/{{.Account}}/{{.Account}}",function(response){
$("#tokenlist").html(response);
});
</script>


            <!-- <a href="/aens/all" class="pull-right">All AENS-></a>-->
              <!-- /.box-body -->
              <div class="box-footer">
               
              </div>
              <!-- /.box-footer -->
            </form>
          </div>
          <!-- /.box -->
        
       </div>
        <!-- /.box-body -->
        <div class="box-footer">

        </div>
        <!-- /.box-footer-->
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


<!-- Bootstrap 3.3.7 -->
<script src="/views/static/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
<!-- SlimScroll -->
<script src="/views/static/bower_components/jquery-slimscroll/jquery.slimscroll.min.js"></script>
<!-- FastClick -->
<script src="/views/static/bower_components/fastclick/lib/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="/views/static/dist/js/adminlte.min.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="/views/static/dist/js/demo.js"></script>
<script>
  $(document).ready(function () {
    $('.sidebar-menu').tree()
  })
</script>
</body>
</html>
