<?php
$status = 0;
$sta = 0;
include "../../../inc/dbinfo.inc";

$conn = mysqli_connect(DB_SERVER, DB_USERNAME, DB_PASSWORD);
if (mysqli_connect_errno()) echo "Failed to connect to MySQL: " . mysqli_connect_error();

$database = mysqli_select_db($conn, DB_DATABASE);
$t=time();
$t=date("Y-m-d h:m:s",$t);

if (isset($_POST['submit'])){
	$username = $_POST['username'];
	$password = $_POST['password'];

	$username = mysqli_real_escape_string($conn, $username);
	$password = mysqli_real_escape_string($conn, $password);

	if($username=="" or $password ==""){
		$status = 1;
		$error_msg = "Username or Password Invalid";
	}
	else
	{
		$sql = "SELECT * FROM USERS  WHERE username = '".$username."'";
		$res = mysqli_query($conn, $sql);
		if (mysqli_num_rows($res) > 0)
		{
			$data = mysqli_fetch_assoc($res);
			$status = 1;
			$error_msg = "This username already taken";
		}

		else{
			$sql = "INSERT INTO USERS (`username`, `password`, `created_at`, `updated_at`) VALUES ('".$username."','".$password."', '".$t."', '".$t."')";
			
			if (mysqli_query($conn, $sql)) {
				$sta= 1;
				$errors_msg = "New id is already create";
			} 
			else {
				//echo "Error: " . $sql . "<br>" . mysqli_error($conn);
			}
		}
	}
}
?>
<!DOCTYPE html>
<html>
<head>
	<title>
		Login to Projects
	</title>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width,initial-scale=1,shrink-to-fit=no">
	<link rel="icon" href="Asset/logo.png" type="image/x-icon">
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">

</head>
<body style="background-color: rgba(243, 241, 239, 1)">
	<div class="container">
		<div class="row mt-3 mb-3">
			<div class="col-lg-4">

			</div>
			<div class="col-lg-4">
				<div class="card">
					<div class="card-header">
						<div class="row">
							<div class="col-lg-4">
								
							</div>
							<div class="col-lg-8 text-right">
								<h5>Sign-up Page</h5>
							</div>
						</div>
					</div>
					<div class="card-body text-center mt-3 mr-3 ml-3 mb-4 ">
						<br>
						<div class="head mb-4">
							<h2>My Drive S3</h2>
							<br>
							<form action="" method="POST">
								<div class="form-group">
									<input type="text" name="username" id = "username" class="form-control" placeholder="Username">
								</div>
								<div class="form-group">
									<input type="password" name="password" id ="password" class="form-control" placeholder="Password">
								</div>

								<hr>
								<button type="submit" name="submit" class="btn btn-danger btn-block">
									Submit
								</button>
								<br>
								<p>already have account? <a href="index.php">Log in</a></p>
								<!--<a href="list.php" class="btn btn-primary btn-block">Login</a>-->
							</form>
						</div>
						<?php if ($status == 1) { ?>
							<div class="alert alert-danger">
								<?php echo $error_msg; ?>!	
							</div>
						<?php } ?>
						<?php if ($sta == 1) { ?>
							<div class="alert alert-success">
								<?php echo $errors_msg; ?>!	
							</div>
						<?php } ?>

					</div>
				</div>
			</div>
			<div class="col-lg-4">

			</div>
		</div>
	</div>
	</html>