<?php 
class AES
{
    var $key = "AFd6N3v1ebLw711zxpZjxZ7iq4fYpNYa";
    var $iv = "MesA7nqIVa23b167";

    function encryptToken($data)
    {
        // Mcrypt library has been DEPRECATED since PHP 7.1, use openssl:
        // return openssl_encrypt($data, 'aes-256-cbc', $this->key, OPENSSL_RAW_DATA, $this->iv);
        $padding = 16 - (strlen($data) % 16);
		$data .= str_repeat(chr($padding), $padding);
		return openssl_encrypt($data, 'aes-256-cbc', $this->key, OPENSSL_DONT_ZERO_PAD_KEY, $this->iv);
        //return mcrypt_encrypt(MCRYPT_RIJNDAEL_256, $this->key, $data, MCRYPT_MODE_CBC, $this->iv);
    }
}

session_start();

if(!isset($_SESSION['id'])){
	header('Location: login/');
}

$uid = $_SESSION['id'];

if (isset($_POST['submit'])) {
	$aes = new AES();

	$subject = $_POST['subject'];
	$password = $_POST['password'];
	$notes = $_POST['notes'];

	$epin = base64_encode($aes->encryptToken($password));

	if ($subject == "" || $password == "" || $notes == ""){
		echo '<script>alert("Please Fill all form!")</script>';
	}
	else{
		$url = 'http://127.0.0.1:4000/v1/api/user/insertnote';
		$data = array(
			'user_id' => $uid,
			'subject' => $subject,
			'notes' => $notes,
			'password' => $epin
		);
	
		$options = array(
			'http' => array(
				'header'  => "Content-Type: application/json\r\n",
				'method'  => 'POST',
				'content' => json_encode($data)
			)
		);
		
		$context  = stream_context_create($options);
		$result = file_get_contents($url, false, $context);
	
		$obj = json_decode($result, TRUE);
	
		if ($obj["stat_code"] == 200){
			echo "<script>alert('Success to create new Note');</script>";
			header("Location: user-main.php");
		}
		else{
			echo "<meta http-equiv='refresh' content='0'>";
			echo "<script>alert('Credential Error');</script>";
		}
	}
}

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


	<title>User Upload</title>
</head>
<nav class="navbar navbar-expand-lg navbar-light bg-light">
  <a class="navbar-brand" href="#">
    <img src="assets/notes.png" width="30" height="30" class="d-inline-block align-top" alt="">
    My Notes
  </a>

  <div class="collapse navbar-collapse" id="navbarNav">
    <ul class="navbar-nav">
      <li class="nav-item">
        <a class="nav-link" href="user-main.php">Home</a>
      </li>
      <li class="nav-item active">
        <a class="nav-link" href="user-upload.php">Upload <span class="sr-only">(current)</span></a>
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
						<center><b>UPLOAD NEW NOTES</b></center>
						<hr>
						<form action="" method="POST" enctype="multipart/form-data">
							<div class="input-group validate-input">
								<input type="Text" name="subject" id = "subject" class="form-control" placeholder="Subjects">
							</div>
                            <br>
							<div class="input-group validate-input">
								<input type="password" name="password" id = "password" class="form-control" placeholder="Password">
							</div>
							
							<br>
							<div class="input-group">
								<textarea class="form-control" id="exampleFormControlTextarea1" type="text" id="notes" name="notes"  placeholder="Notes"></textarea>
							</div>
							<hr>
							<input type="submit" id="submitFormData" name="submit" class="btn btn-primary btn-block" value="upload">
						</form>
					</div>
				</div>

			</div>
		</div>
	</div>
	<!-- <script src="../js/bootstrap.min.js"></script>

<script src="../js/plugins.js"></script> -->


</body>
</html>