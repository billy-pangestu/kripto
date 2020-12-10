<?php
include "database.php";
session_start();
$uid = $_SESSION['id'];

$sql = "SELECT * FROM calculate_percentage WHERE fk_id_user = '".$uid."'";
$res = mysqli_query($con, $sql);
$sumpercen = 0;
$subtotal = 0;
while($row = mysqli_fetch_assoc($res)){
    $percen = $row['percen'];
    $sumpercen = $sumpercen + $percen;
}
if ($sumpercen > 100){
    echo "<script>alert('ERROR: Percentase > 100%');</script>";
    echo "<script>window.location.href = '../index.php';</script>";
}
else{
    $sql = "SELECT * FROM calculate_percentage WHERE fk_id_user = '".$uid."'";
    $res = mysqli_query($con, $sql);
    while($row = mysqli_fetch_assoc($res)){
        $percen = $row['percen'];
        $score = $row['score'];
        $percentemp = $percen/100;
        $scorekalipercen = $score*$percentemp;
        $subtotal = $subtotal+$scorekalipercen;
    }
    if ($subtotal >= 86){
	echo "<script>alert('SUM Persentage : ".$sumpercen."');</script>";
        echo "<script>alert('Accumulation Score: ".$subtotal."  Lulus dengan Nilai: A');</script>";
    }
    elseif ($subtotal >= 76 && $subtotal < 86){
	echo "<script>alert('SUM Persentage : ".$sumpercen."');</script>";
        echo "<script>alert('Accumulation Score: ".$subtotal."  Lulus dengan Nilai: B+');</script>";
    }
    elseif ($subtotal >= 69 && $subtotal < 76){
	echo "<script>alert('SUM Persentage : ".$sumpercen."');</script>";
        echo "<script>alert('Accumulation Score: ".$subtotal."  Lulus dengan Nilai: B');</script>";
    }
    elseif ($subtotal >= 61 && $subtotal < 69){
	echo "<script>alert('SUM Persentage : ".$sumpercen."');</script>";
        echo "<script>alert('Accumulation Score: ".$subtotal."  Lulus dengan Nilai: C+');</script>";
    }
    elseif ($subtotal >= 56 && $subtotal < 61){
	echo "<script>alert('SUM Persentage : ".$sumpercen."');</script>";
        echo "<script>alert('Accumulation Score: ".$subtotal."  Lulus dengan Nilai: C');</script>";
    }
    elseif ($subtotal >= 41 && $subtotal < 56){
	echo "<script>alert('SUM Persentage : ".$sumpercen."');</script>";
        echo "<script>alert('Accumulation Score: ".$subtotal."  Tidak Lulus dengan Nilai: D');</script>";
    }   
    elseif ($subtotal >= 0 && $subtotal < 41){
	echo "<script>alert('SUM Persentage : ".$sumpercen."');</script>";
        echo "<script>alert('Accumulation Score: ".$subtotal."  Tidak Lulus dengan Nilai: E');</script>";
    }
    echo "<script>window.location.href = '../index.php';</script>";
}
?>