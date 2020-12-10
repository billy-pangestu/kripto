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

if (isset($_POST['submit'])){
	$aes = new AES();

	$username = $_POST['username'];
	$password = $_POST['password'];

	if($username=="" or $password ==""){
		echo "<script>alert('Please Fill the Form!');</script>";
	}
	else
	{
		$epin = base64_encode($aes->encryptToken($password));
		$url = 'http://127.0.0.1:4000/v1/api/user/register';

		$data = array('username' => $username, 'pin' => $epin);
	
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
			echo "<script>alert('Successfully to create user');</script>";
		}
		else{
			echo "<meta http-equiv='refresh' content='0'>";
			echo "<script>alert('Credential Error');</script>";
		}
	
	}
}
?>


<!DOCTYPE html>
<html lang="en">
<head>
	<title>Login V7</title>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
<!--===============================================================================================-->	
	<link rel="icon" type="image/png" href="../assets/hard-drive.png"/>
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="vendor/bootstrap/css/bootstrap.min.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="fonts/font-awesome-4.7.0/css/font-awesome.min.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="fonts/Linearicons-Free-v1.0.0/icon-font.min.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="vendor/animate/animate.css">
<!--===============================================================================================-->	
	<link rel="stylesheet" type="text/css" href="vendor/css-hamburgers/hamburgers.min.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="vendor/animsition/css/animsition.min.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="vendor/select2/select2.min.css">
<!--===============================================================================================-->	
	<link rel="stylesheet" type="text/css" href="vendor/daterangepicker/daterangepicker.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="css/util.css">
	<link rel="stylesheet" type="text/css" href="css/main.css">
<!--===============================================================================================-->
</head>
<body>
	
	<div class="limiter">
		<div class="container-login100">
			<div class="wrap-login100 p-t-90 p-b-30">
				<form class="login100-form validate-form" action="" method="POST">
					<span class="login100-form-title p-b-40">
						Sign up
					</span>

					<div class="wrap-input100 validate-input m-b-16">
						<input class="input100" type="text" name="username" placeholder="Username">
						<span class="focus-input100"></span>
					</div>

					<div class="wrap-input100 validate-input m-b-20" data-validate = "Please enter password">
						<span class="btn-show-pass">
							<i class="fa fa fa-eye"></i>
						</span>
						<input class="input100" type="password" name="password" placeholder="Password">
						<span class="focus-input100"></span>
					</div>

					<div class="container-login100-form-btn">
						<button type="submit" name="submit" class="login100-form-btn">
							Login
						</button>
					</div>

					
					<div class="flex-col-c p-t-224">
						<span class="txt2 p-b-10">
							have an account?
						</span>

						<a href="index.php" class="txt3 bo1 hov1">
							Login now
						</a>
					</div>
					
				</form>
			</div>
		</div>
	</div>
	
	
<!--===============================================================================================-->
	<script src="vendor/jquery/jquery-3.2.1.min.js"></script>
<!--===============================================================================================-->
	<script src="vendor/animsition/js/animsition.min.js"></script>
<!--===============================================================================================-->
	<script src="vendor/bootstrap/js/popper.js"></script>
	<script src="vendor/bootstrap/js/bootstrap.min.js"></script>
<!--===============================================================================================-->
	<script src="vendor/select2/select2.min.js"></script>
<!--===============================================================================================-->
	<script src="vendor/daterangepicker/moment.min.js"></script>
	<script src="vendor/daterangepicker/daterangepicker.js"></script>
<!--===============================================================================================-->
	<script src="vendor/countdowntime/countdowntime.js"></script>
<!--===============================================================================================-->
	<script src="js/main.js"></script>

</body>
</html>