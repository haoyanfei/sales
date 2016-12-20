{{template "public/header.html"}}

	<div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
        Page Header
        <small>Optional description</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> Level</a></li>
        <li class="active">Here</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
<form class="form-horizontal" role="form" action="/admin/role/add" method="post">
  <div class="form-group form-group-lg">
    <label class="col-sm-2 control-label" for="formGroupInputLarge">角色名称</label>
    <div class="col-sm-10">
      <input class="form-control" name="Name" type="text"  placeholder="Large input">
      <input class="form-control" name="Id" type="text"  placeholder="Large input">
	</div>
  </div>
 <div class="form-group form-group-lg">
    <label class="col-sm-2 control-label" for="formGroupInputLarge"></label>
    <div class="col-sm-10">
      <input class="form-control" type="submit" id="formGroupInputLarge" placeholder="Large input">
    </div>
  </div>
</form>

    </section>
    <!-- /.content -->
  </div>
  <!-- /.content-wrapper -->

{{template "public/footer.html"}}