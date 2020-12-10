<?php 

// konfigurasi koneksi
$host       =  "127.0.0.1";
$dbuser     =  "postgres";
$dbpass     =  "isabilly2200";
$port       =  "5432";
 $dbname    =  "kriptografi";

 $dbconn = pg_connect("dbname=kriptografi");

// $conn_string = "host=localhost port=5432 dbname=kriptografi user=postgres password=isabilly2200";
// $dbconn4 = pg_connect("host=localhost port=5432 dbname=kriptografi user=postgres password=isabilly2200");

// script koneksi php postgree
// $link = new PDO("pgsql:dbname=kriptografi;host=127.0.0.1", $dbuser, $dbpass); 
 
if($dbconn)
{
    echo "Koneksi Berhasil";
}else
{
    echo "Gagal melakukan Koneksi";
}

?>