<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>AEK</title>
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
	
<!-- <body class="hold-transition skin-blue  sidebar-collapse sidebar-mini">-->
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
{{ template "sidebar" .}}

  <!-- Content Wrapper. Contains page content -->
  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
        Exploring the knowledge of mankind without barriers.
        <small>You, are the center of blockchain world.</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
             </ol>
    </section>

    <!-- Main content -->
    <section class="content">
  <!-- 
  <center> 
	        <div class="input-group " style="width:600px;">
                <input type="text" class="form-control" style="height:45px;"> 
                    <span class="input-group-btn">
                     &nbsp;   <button type="button" class="btn btn-info btn-flat" style="height:40px;">Go</button>
                    </span>
              </div>
   </center> -->
    <!-- /.lockscreen credentials -->

<br />
<div class="row">
	       <div class="col-md-4">
          <!-- Custom Tabs (Pulled to the right) -->
          <div class="nav-tabs-custom">
            <ul class="nav nav-tabs pull-right">
				 <li class="dropdown">
                <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                  <span class="caret"></span>
                </a>
                <ul class="dropdown-menu">
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#"><i class="glyphicon glyphicon-film"></i> Video</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#"><i class="glyphicon glyphicon-picture"></i> Picture</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#"><i class="fa fa-database"></i> Data</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#"><i class="fa  fa-code"></i> Code</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#"><i class="glyphicon glyphicon-resize-horizontal"></i> Dictionary</a></li>
                  <li role="presentation" class="divider"></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#"><i class="glyphicon glyphicon-download-alt"></i>Tools/Software</a></li>
                </ul>
              </li>
              <li ><a href="#tab_libpack" data-toggle="tab"><i class="glyphicon glyphicon-tasks"></i> Pack</a></li>    
                                                   
              <li ><a href="#tab_libbook" data-toggle="tab"><i class="fa  fa-book"></i> Book</a></li>
              <li ><a href="#tab_libarticle" data-toggle="tab"><i class="glyphicon glyphicon-font"></i> Article</a></li>
              <li class="active"><a href="#tab_libwiki" data-toggle="tab"><i class="fa  fa-wikipedia-w"></i> Wiki</a></li>
             
              <li class="pull-left header"><i class="fa fa-th"></i>aeLibrary</li>
            </ul>
            <div class="tab-content" style="height:220px;">
              <div class="tab-pane active" id="tab_libwiki">
                <b>Wikis</b>

                <p>Exactly like the original bootstrap tabs except you should use
                  the custom wrapper <code>.nav-tabs-custom</code> to achieve this style.</p>
                A wonderful serenity has taken possession of my entire soul,
                like these sweet mornings of spring which I enjoy with my whole heart.
                I am alone, and feel the charm of existence in this spot,
                which was created for the bliss of souls like mine. I am so happy,
                my dear friend, so absorbed in the exquisite sense of mere tranquil existence,
                that I neglect my talents. I should be incapable of drawing a single stroke
                at the present moment; and yet I feel that I never was a greater artist than now.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_libarticle">
				  <b>Articles and papers.</b>
                The European languages are members of the same family. Their separate existence is a myth.
                For science, music, sport, etc, Europe uses the same vocabulary. The languages only differ
                in their grammar, their pronunciation and their most common words. Everyone realizes why a
                new common language would be desirable: one could refuse to pay expensive translators. To
                achieve this, it would be necessary to have uniform grammar, pronunciation and more common
                words. If several languages coalesce, the grammar of the resulting language is more simple
                and regular than that of the individual languages.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_libbook">
				  <b>Books online</b>
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                It has survived not only five centuries, but also the leap into electronic typesetting,
                remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
                like Aldus PageMaker including versions of Lorem Ipsum.
              </div>
              <!-- /.tab-pane -->
               <div class="tab-pane" id="tab_libpack">
				  <b>Knowledge packs</b>
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                It has survived not only five centuries, but also the leap into electronic typesetting,
                remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
                like Aldus PageMaker including versions of Lorem Ipsum.
              </div>
              <!-- /.tab-pane -->
            </div>
            <!-- /.tab-content -->
          </div>
          <!-- nav-tabs-custom -->
        </div>
        <!-- /.col lib-->
	
	  <div class="col-md-4">
          <!-- Custom Tabs (Pulled to the right) -->
          <div class="nav-tabs-custom">
            <ul class="nav nav-tabs pull-right">
				
				<li class="dropdown">
                <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                  <span class="caret"></span>
                </a>
                <ul class="dropdown-menu">
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Thoughts</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Another action</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Something else here</a></li>
                  <li role="presentation" class="divider"></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Separated link</a></li>
                </ul>
              </li>
              <li ><a href="#tab_lthoughts" data-toggle="tab">Skills</a></li>
              <li><a href="#tab_lskills" data-toggle="tab">Tools</a></li>
              <li class="active"><a href="#tab_ltools" data-toggle="tab">Courses</a></li>
              
              <li class="pull-left header"><i class="fa  fa-book"></i>Learn</li>
            </ul>
            <div class="tab-content" style="height:220px;">
              <div class="tab-pane active" id="tab_lthoughts">
                <b>How to use:</b>

                <p>Exactly like the original bootstrap tabs except you should use
                  the custom wrapper <code>.nav-tabs-custom</code> to achieve this style.</p>
                A wonderful serenity has taken possession of my entire soul,
                like these sweet mornings of spring which I enjoy with my whole heart.
                I am alone, and feel the charm of existence in this spot,
                which was created for the bliss of souls like mine. I am so happy,
                my dear friend, so absorbed in the exquisite sense of mere tranquil existence,
                that I neglect my talents. I should be incapable of drawing a single stroke
                at the present moment; and yet I feel that I never was a greater artist than now.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_lskills">
                The European languages are members of the same family. Their separate existence is a myth.
                For science, music, sport, etc, Europe uses the same vocabulary. The languages only differ
                in their grammar, their pronunciation and their most common words. Everyone realizes why a
                new common language would be desirable: one could refuse to pay expensive translators. To
                achieve this, it would be necessary to have uniform grammar, pronunciation and more common
                words. If several languages coalesce, the grammar of the resulting language is more simple
                and regular than that of the individual languages.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_ltools">
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                It has survived not only five centuries, but also the leap into electronic typesetting,
                remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
                like Aldus PageMaker including versions of Lorem Ipsum.
              </div>
              <!-- /.tab-pane -->
            </div>
            <!-- /.tab-content -->
          </div>
          <!-- nav-tabs-custom -->
        </div>
        <!-- /.col -->
	    
     <div class="col-md-4">
          <!-- Custom Tabs (Pulled to the right) -->
          <div class="nav-tabs-custom">
            <ul class="nav nav-tabs pull-right">
				<li class="dropdown">
                <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                <span class="caret"></span>
                </a>
                <ul class="dropdown-menu">
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Action</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Another action</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Something else here</a></li>
                  <li role="presentation" class="divider"></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Separated link</a></li>
                </ul>
              </li>
              <li ><a href="#tab_crcollection" data-toggle="tab">Collection</a></li>
              <li><a href="#tab_crbook" data-toggle="tab">Book</a></li>
              <li class="active"><a href="#tab_crarticle" data-toggle="tab">Article</a></li>
              
              <li class="pull-left header"><a href="/view?pubkey={{.Account}}&viewtype=author"><i class="glyphicon glyphicon-pencil"></i>My</a></li>
            </ul>
            <div class="tab-content" style="height:220px;">
              <div class="tab-pane active" id="tab_crcollection">
                <b>How to use:</b>

                <p>Exactly like the original bootstrap tabs except you should use
                  the custom wrapper <code>.nav-tabs-custom</code> to achieve this style.</p>
                A wonderful serenity has taken possession of my entire soul,
                like these sweet mornings of spring which I enjoy with my whole heart.
                I am alone, and feel the charm of existence in this spot,
                which was created for the bliss of souls like mine. I am so happy,
                my dear friend, so absorbed in the exquisite sense of mere tranquil existence,
                that I neglect my talents. I should be incapable of drawing a single stroke
                at the present moment; and yet I feel that I never was a greater artist than now.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_crbook">
                The European languages are members of the same family. Their separate existence is a myth.
                For science, music, sport, etc, Europe uses the same vocabulary. The languages only differ
                in their grammar, their pronunciation and their most common words. Everyone realizes why a
                new common language would be desirable: one could refuse to pay expensive translators. To
                achieve this, it would be necessary to have uniform grammar, pronunciation and more common
                words. If several languages coalesce, the grammar of the resulting language is more simple
                and regular than that of the individual languages.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_crarticle">
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                It has survived not only five centuries, but also the leap into electronic typesetting,
                remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
                like Aldus PageMaker including versions of Lorem Ipsum.
              </div>
              <!-- /.tab-pane -->
            </div>
            <!-- /.tab-content -->
          </div>
          <!-- nav-tabs-custom -->
        </div>
        <!-- /.col -->
     
      </div>
<div class="row">
       
        
         <div class="col-md-4">
          <!-- Custom Tabs (Pulled to the right) -->
          <div class="nav-tabs-custom">
            <ul class="nav nav-tabs pull-right">
				<li class="dropdown">
                <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                 <span class="caret"></span>
                </a>
                <ul class="dropdown-menu">
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Action</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Another action</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Something else here</a></li>
                  <li role="presentation" class="divider"></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Separated link</a></li>
                </ul>
              </li>
              
              
              <li ><a href="#tab_refringe" data-toggle="tab">Fringe</a></li>
              <li><a href="#tab_retheory" data-toggle="tab">Theory</a></li>
              <li class="active"><a href="#tab_retech" data-toggle="tab">Tech</a></li>
              
              <li class="pull-left header"><i class="glyphicon glyphicon-education"></i> Research</li>
            </ul>
            <div class="tab-content" style="height:220px;">
              <div class="tab-pane active" id="tab_refringe">
                <b>How to use:</b>

                <p>Exactly like the original bootstrap tabs except you should use
                  the custom wrapper <code>.nav-tabs-custom</code> to achieve this style.</p>
                A wonderful serenity has taken possession of my entire soul,
                like these sweet mornings of spring which I enjoy with my whole heart.
                I am alone, and feel the charm of existence in this spot,
                which was created for the bliss of souls like mine. I am so happy,
                my dear friend, so absorbed in the exquisite sense of mere tranquil existence,
                that I neglect my talents. I should be incapable of drawing a single stroke
                at the present moment; and yet I feel that I never was a greater artist than now.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_retheory">
                The European languages are members of the same family. Their separate existence is a myth.
                For science, music, sport, etc, Europe uses the same vocabulary. The languages only differ
                in their grammar, their pronunciation and their most common words. Everyone realizes why a
                new common language would be desirable: one could refuse to pay expensive translators. To
                achieve this, it would be necessary to have uniform grammar, pronunciation and more common
                words. If several languages coalesce, the grammar of the resulting language is more simple
                and regular than that of the individual languages.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_retech">
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                It has survived not only five centuries, but also the leap into electronic typesetting,
                remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
                like Aldus PageMaker including versions of Lorem Ipsum.
              </div>
              <!-- /.tab-pane -->
            </div>
            <!-- /.tab-content -->
          </div>
          <!-- nav-tabs-custom -->
        </div>
        <!-- /.col -->
    
    <div class="col-md-4">
          <!-- Custom Tabs (Pulled to the right) -->
          <div class="nav-tabs-custom">
            <ul class="nav nav-tabs pull-right">
				<li class="dropdown">
                <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                <span class="caret"></span>
                </a>
                <ul class="dropdown-menu">
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Action</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Another action</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Something else here</a></li>
                  <li role="presentation" class="divider"></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Separated link</a></li>
                </ul>
              </li>
              <li ><a href="#tab_coforum" data-toggle="tab">Forum</a></li>
              <li><a href="#tab_cochat" data-toggle="tab">Chat</a></li>
              <li class="active"><a href="#tab_comessage" data-toggle="tab">Message</a></li>
              
              <li class="pull-left header"><i class="glyphicon glyphicon-comment"></i>Communication</li>
            </ul>
            <div class="tab-content" style="height:220px;">
              <div class="tab-pane active" id="tab_coforum">
                <b>How to use:</b>

                <p>Exactly like the original bootstrap tabs except you should use
                  the custom wrapper <code>.nav-tabs-custom</code> to achieve this style.</p>
                A wonderful serenity has taken possession of my entire soul,
                like these sweet mornings of spring which I enjoy with my whole heart.
                I am alone, and feel the charm of existence in this spot,
                which was created for the bliss of souls like mine. I am so happy,
                my dear friend, so absorbed in the exquisite sense of mere tranquil existence,
                that I neglect my talents. I should be incapable of drawing a single stroke
                at the present moment; and yet I feel that I never was a greater artist than now.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_cochat">
                The European languages are members of the same family. Their separate existence is a myth.
                For science, music, sport, etc, Europe uses the same vocabulary. The languages only differ
                in their grammar, their pronunciation and their most common words. Everyone realizes why a
                new common language would be desirable: one could refuse to pay expensive translators. To
                achieve this, it would be necessary to have uniform grammar, pronunciation and more common
                words. If several languages coalesce, the grammar of the resulting language is more simple
                and regular than that of the individual languages.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_comessage">
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                It has survived not only five centuries, but also the leap into electronic typesetting,
                remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
                like Aldus PageMaker including versions of Lorem Ipsum.
              </div>
              <!-- /.tab-pane -->
            </div>
            <!-- /.tab-content -->
          </div>
          <!-- nav-tabs-custom -->
        </div>
        <!-- /.col -->
    
    
    <div class="col-md-4">
          <!-- Custom Tabs (Pulled to the right) -->
          <div class="nav-tabs-custom">
            <ul class="nav nav-tabs pull-right">
				<li class="dropdown">
                <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                 <span class="caret"></span>
                </a>
                <ul class="dropdown-menu">
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Action</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Another action</a></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Something else here</a></li>
                  <li role="presentation" class="divider"></li>
                  <li role="presentation"><a role="menuitem" tabindex="-1" href="#">Separated link</a></li>
                </ul>
              </li>
              
              
              
              <li ><a href="#tab_helse" data-toggle="tab">Else</a></li>
              <li><a href="#tab_hlinks" data-toggle="tab">Links</a></li>
              <li class="active"><a href="#tab_hdocuments" data-toggle="tab">Documents</a></li>
              
              
              <li class="pull-left header"><i class="glyphicon glyphicon-question-sign"></i> Help</li>
            </ul>
            <div class="tab-content" style="height:220px;">
              <div class="tab-pane active" id="tab_hdocuments">
                <b>How to use:</b>

                <p>Exactly like the original bootstrap tabs except you should use
                  the custom wrapper <code>.nav-tabs-custom</code> to achieve this style.</p>
                A wonderful serenity has taken possession of my entire soul,
                like these sweet mornings of spring which I enjoy with my whole heart.
                I am alone, and feel the charm of existence in this spot,
                which was created for the bliss of souls like mine. I am so happy,
                my dear friend, so absorbed in the exquisite sense of mere tranquil existence,
                that I neglect my talents. I should be incapable of drawing a single stroke
                at the present moment; and yet I feel that I never was a greater artist than now.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_hlinks">
                The European languages are members of the same family. Their separate existence is a myth.
                For science, music, sport, etc, Europe uses the same vocabulary. The languages only differ
                in their grammar, their pronunciation and their most common words. Everyone realizes why a
                new common language would be desirable: one could refuse to pay expensive translators. To
                achieve this, it would be necessary to have uniform grammar, pronunciation and more common
                words. If several languages coalesce, the grammar of the resulting language is more simple
                and regular than that of the individual languages.
              </div>
              <!-- /.tab-pane -->
              <div class="tab-pane" id="tab_helse">
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                It has survived not only five centuries, but also the leap into electronic typesetting,
                remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
                like Aldus PageMaker including versions of Lorem Ipsum.
              </div>
              <!-- /.tab-pane -->
            </div>
            <!-- /.tab-content -->
          </div>
          <!-- nav-tabs-custom -->
        </div>
        <!-- /.col -->
    
      </div>
      <!-- /.row -->
 
<div class="row">
	<center>
      
          <div class="box box-solid"  style="width:70%;">
            <div class="box-header with-border">
              <i class="fa fa-text-width"></i>

              <h3 class="box-title">Block Quotes</h3>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <blockquote>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer posuere erat a ante.</p>
                <small>Someone famous in <cite title="Source Title">Source Title</cite></small>
              </blockquote>
            </div>
            <!-- /.box-body -->
          </div>
          <!-- /.box -->
        
        <!-- ./col -->
       </center>
</div>
 
 <div class="row">
 <div class="col-md-12">
<div class="box box-solid">
            <div class="box-header with-border">
              <i class="fa fa-text-width"></i>

              <h3 class="box-title">About</h3>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <h1>h1. Bootstrap heading</h1>

              <h2>h2. Bootstrap heading</h2>

              <h3>h3. Bootstrap heading</h3>
              <h4>h4. Bootstrap heading</h4>
              <h5>h5. Bootstrap heading</h5>
              <h6>h6. Bootstrap heading</h6>
            </div>
            <!-- /.box-body -->
          </div>
          <!-- /.box --></div>

</div> 
 

  
  <!-- /.lockscreen-item -->

    </section>
    <!-- /.content -->
  </div>
  <!-- /.content-wrapper -->
 {{ template "footer" .}}
  <div class="control-sidebar-bg"></div>
</div>
<!-- ./wrapper -->

<!-- jQuery 3 -->
<script src="/views/static/bower_components/jquery/dist/jquery.min.js"></script>
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
