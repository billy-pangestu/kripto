
<?php
session_start();

if(!isset($_SESSION['id'])){
	header('Location: login/');
}

$url = "http://127.0.0.1:4000/v1/api/user/getnotes/";
$json = file_get_contents('http://127.0.0.1:4000/v1/api/user/getnotes/'.$_SESSION['id']);
$obj = json_decode($json, TRUE);


// foreach ($obj as $key => $value) {
//   echo $value["id"] . ", " . $value["subject"] . "<br>";
// }

// if(obj->{'stat_msg'} == "Success"){
// 	echo "berhasil";
// }
// else{
// 	echo "gagal";
// }
?>
<!DOCTYPE html>
<html>
<head>
	<meta name="description" content="">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="icon" type="image/png" href="assets/notes.png"/>
	<link rel="stylesheet" href="css/bootstrap.min.css">
	<script src="https://code.jquery.com/jquery-3.5.1.min.js"></script>
    
    <link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.10.21/css/dataTables.bootstrap4.min.css"/>
    
    <script type="text/javascript" src="https://cdn.datatables.net/1.10.21/js/jquery.dataTables.min.js"></script>
    <script type="text/javascript" src="https://cdn.datatables.net/1.10.21/js/dataTables.bootstrap4.min.js"></script>


	<title>User Home</title>
</head>
<nav class="navbar navbar-expand-lg navbar-light bg-light">
  <a class="navbar-brand" href="#">
    <img src="assets/notes.png" width="30" height="30" class="d-inline-block align-top" alt="">
    My Notes
  </a>

  <div class="collapse navbar-collapse" id="navbarNav">
    <ul class="navbar-nav">
      <li class="nav-item active">
        <a class="nav-link" href="user-main.php">Home <span class="sr-only">(current)</span></a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="user-upload.php">Upload</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="services/logout.php">Logout</a>
      </li>
    </ul>
  </div>
</nav>
<body>
	<div class="container">
		<div class="row mt-5">
			<div class="col-lg-12">
				<div class="card border-0">
					<div class="card-body">
						<table class="table">
							<thead>
								<tr>
									<th scope="col" style="width: 40%">Subjects</th>
									<th scope="col" style="width: 40%"></th>
								</tr>
							</thead>
							<tbody>
									<?php
									if ($obj["data"] == ""){
										
									}
									else{
										for ($i = 0; $i < count($obj["data"]); $i++) {
											$idrow = $obj["data"][$i]["id"];
											$subject = $obj["data"][$i]["subject"];
											$notes = $obj["data"][$i]["notes"];
	
											echo"
											<tr>
												<th scope='row'>".$subject."</th>
												<td>
												<a href='user-detail.php?id=".$idrow."' class='btn btn-primary'><img src='assets/text-format.png' width='30' height='30' alt=''></a>
												<a href='services/delete.php?id=".$idrow."'' class='btn btn-danger'><img src='assets/delete.png' width='30' height='30' alt=''></a>
												</td>
											</tr>";
										}
									}
									?>
									
							</tbody>
						</table>

					</div>
				</div>

			</div>
		</div>
	</div>
	<script src="../js/bootstrap.min.js"></script>

<script src="../js/plugins.js"></script>

<script type="text/javascript">
	$(document).ready(function(){
		$('table').DataTable();
	});

</script>

</body>
</html>