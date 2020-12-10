<?php

session_start();

if(!isset($_SESSION['id'])){
	header('Location: ../login/');
}

$uid = $_SESSION['id'];

$idrow = $_GET['id'];

$url = 'http://127.0.0.1:4000/v1/api/user/deletenote/'.$idrow;

$options = array(
	'http' => array(
		'header'  => "Content-Type: application/json\r\n",
		'method'  => 'DELETE'
	)
);

$context  = stream_context_create($options);
$result = file_get_contents($url, false, $context);

$obj = json_decode($result, TRUE);

if ($obj["stat_code"] == 200){
	header('Location: deletenotif.php');
}
else{
	echo "<meta http-equiv='refresh' content='0'>";
	echo "<script>alert('Credential Error');</script>";
	header('Location: ../user-main.php');
}

?>