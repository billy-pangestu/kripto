<?php
session_start();

if(!isset($_SESSION['id'])){
	header('Location: login/');
}

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

$idnote = $_GET['id'];

$json = file_get_contents('http://127.0.0.1:4000/v1/api/user/getnote/'.$idnote);
$obj = json_decode($json, TRUE);


if (isset($_POST['submit'])) {
    $aes = new AES();
    
    $idrow = $obj["data"]["id"];
    $password = $_POST['password'];

	$epin = base64_encode($aes->encryptToken($password));

	if ($password == ""){
		echo '<script>alert("Please Fill password!")</script>';
	}
	else{
        $json = file_get_contents('http://127.0.0.1:4000/v1/api/user/getdecriptnote/'.$idnote.'/'.$epin);
        $obj = json_decode($json, TRUE);

        if ($obj["stat_code"] == 200){
            $notes = $obj["data"]["notes"];
        }
        else{
            echo "<meta http-equiv='refresh' content='0'>";
		    echo "<script>alert('Password not Match!');</script>";
        }
		// $url = 'http://127.0.0.1:4000/v1/api/user/insertnote';
		// $data = array(
		// 	'user_id' => $uid,
		// 	'subject' => $subject,
		// 	'notes' => $notes,
		// 	'password' => $epin
		// );
	
		// $options = array(
		// 	'http' => array(
		// 		'header'  => "Content-Type: application/json\r\n",
		// 		'method'  => 'POST',
		// 		'content' => json_encode($data)
		// 	)
		// );
		
		// $context  = stream_context_create($options);
		// $result = file_get_contents($url, false, $context);
	
		// $obj = json_decode($result, TRUE);
	
		// if ($obj["stat_code"] == 200){
		// 	echo $obj["data"];
		// }
		// else{
		// 	echo "<meta http-equiv='refresh' content='0'>";
		// 	echo "<script>alert('Credential Error');</script>";
		// }
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
                    <center><b>DETAILS NOTE</b></center>
						<hr>
                        <?php
									if ($obj["data"] == ""){
										
									}
									else{
										$idrow = $obj["data"]["id"];
										$subject = $obj["data"]["subject"];
										$notes = $obj["data"]["notes"];
						?>
						<form action="" method="POST" enctype="multipart/form-data">
                            Subject:
							<div class="input-group validate-input">
								<input type="Text" name="subject" id = "subject" class="form-control" placeholder="Subjects" disabled value="<?php echo $subject; ?>">
							</div>
                            <br>
							<div class="input-group validate-input">
                                <div class="col-lg-6"><input type="password" name="password" id = "password" class="form-control" placeholder="Password"></div>
                                <div class="col-lg-6"><input type="submit" id="submitFormData" name="submit" class="btn btn-warning btn-block" value="Decript Note"></div>
							</div>
							<hr>
                            Note:
							<div class="input-group">
								<textarea class="form-control" id="exampleFormControlTextarea1" type="text" id="notes" name="notes" disabled placeholder="Notes"><?php echo $notes;?></textarea>
							</div>
							
						</form>
                        <?php
                                    }
                        ?>

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