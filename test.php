
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

    function decryptToken($data)
    {
        // Mcrypt library has been DEPRECATED since PHP 7.1, use openssl:
        return openssl_decrypt(base64_decode($data), 'aes-256-cbc', $this->key, OPENSSL_RAW_DATA, $this->iv);
        //$data = mcrypt_decrypt(MCRYPT_RIJNDAEL_128, $this->key, base64_decode($data), MCRYPT_MODE_CBC, $this->iv);
        //$padding = ord($data[strlen($data) - 1]);
        //return substr($data, 0, -$padding);
    }
}



if (php_sapi_name() === 'cli')
{
    $aes = new AES();
    echo ('PHP encrypt: '.base64_encode($aes->encryptToken('dmyz.org')))."\n";
    echo ('PHP decrypt: '.$aes->decryptToken('FSfhJ/gk3iEJOPVLyFVc2Q=='))."\n";
}

if (isset($_POST['submitchange'])){
	$aes = new AES();

	$phonenumber = $_POST['phone_number'];
	$newpin = $_POST['new_pin'];

	//$epin = $newpin;

	$epin = base64_encode($aes->encryptToken($newpin));
	$dpin = $aes->decryptToken($epin);

	echo "encrypt: ".$epin;
	echo "<br>";

	// $plaintext = "message to be encrypted";
	// //$ivlen = openssl_cipher_iv_length();
	// $cipher="AES-256-CBC";
	// $iv = "MesA7nqIVa23b167";
	// $key = "AFd6N3v1ebLw711zxpZjxZ7iq4fYpNYa";
	// $ciphertext_raw = openssl_encrypt($newpin, $cipher, $key, $options=OPENSSL_RAW_DATA, $iv);
	// $ciphertext = base64_encode( $iv.$ciphertext_raw );

	// echo $epin;
	// print_r($_POST);

	$url = 'http://127.0.0.1:4000/v1/api/user/changepinwotoken';
	$data = array('phone_number' => $phonenumber, 'new_pin' => $epin);

	// use key 'http' even if you send the request to https://...
	$options = array(
    	'http' => array(
       		'header'  => "Content-type: application/x-www-form-urlencoded\r\n",
        	'method'  => 'POST',
        	'content' => http_build_query($data)
    	)
	);
	$context  = stream_context_create($options);
	$result = file_get_contents($url, false, $context);
	echo $result;
	
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
								<h5>Kripto</h5>
							</div>
						</div>

					</div>
					<div class="card-body text-center mt-3 mr-3 ml-3 mb-4 ">
						<br>
						<div class="head mb-4">
							<br>
							<!-- http://127.0.0.1:4000/v1/api/user/changepinwotoken -->
                            <form action="" method="POST">
                                <div class="form-group">
									<input type="text" name="phone_number" id = "phonenumber" class="form-control" placeholder="phonenumber" value="+6281339223132">
								</div>
						
								<div class="form-group">
									<input type="password" name="new_pin" id ="newpin" class="form-control" placeholder="new password">
								</div>
								<hr>
								<button type="submit" name="submitchange" class="btn btn-primary btn-block">
									Set New Password
								</button>
								<br>
                            </form>
                            <hr>
                            
                            <form action="http://127.0.0.1:4000/v1/api/user/showepin" method="POST">
                                <div class="form-group">
									<input type="text" name="phone_number" id = "phonenumber" class="form-control" placeholder="phonenumber"value="+6281339223132">
								</div>
								<button type="submit" name="submit" class="btn btn-primary btn-block">
									show Encrypt Password
								</button>
								<br>
                            </form>
                            <hr>
                            <form action="http://127.0.0.1:4000/v1/api/user/showppin" method="POST">
                                <div class="form-group">
									<input type="text" name="phone_number" id = "phonenumber" class="form-control" placeholder="phonenumber"value="+6281339223132">
								</div>
								<button type="submit" name="submit" class="btn btn-primary btn-block">
									show Plain Password
								</button>
								<br>
							</form>
							<hr>
							<div class="col-lg">
								<a href="../../" class="btn btn-dark btn-block">Back</a>
							</div>
						</div>
						
					</div>
				</div>
			</div>
			<div class="col-lg-4">

			</div>
		</div>
	</div>
	</html>