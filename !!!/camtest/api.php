<?php
//throw new Exception("Test");
$fileName = $_FILES['file']['name'];
$fileType = $_FILES['file']['type'];
$fileError = $_FILES['file']['error'];
$fileContent = file_get_contents($_FILES['file']['tmp_name']);
file_put_contents('new_file', $fileContent);
?>
