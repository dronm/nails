<?php

function parsePgConnStr($s) {
	$conn = [];
	$c = explode("//", $s);
	if(count($c) < 2){
		die('invalid parameter');
	}
	$parts = explode("@", $c[1]);
	if(count($parts) < 2){
		die('invalid parameter');
	}
	$user_pwd = explode(":", $parts[0]);
	if(count($user_pwd) < 2){
		die('invalid parameter');
	}
	$conn['user'] = $user_pwd[0];
	$conn['pwd'] = $user_pwd[1];

	$host_port_name = explode("/", $parts[1]);
	if(count($host_port_name) < 2){
		die('invalid parameter');
	}
	$conn['dbname'] = $host_port_name[1];

	$host_port = explode(":", $host_port_name[0]);
	if(count($host_port_name) < 2){
		die('invalid parameter');
	}
	$conn['host'] = $host_port[0];
	$conn['port'] = $host_port[1];
	
	return $conn;
}

$conn = parsePgConnStr('postgresql://nails:159753@localhost:5432/nails');
echo sprintf('User=%s, pwd=%s, dbname=%s, host=%s, port=%s', $conn['user'], $conn['pwd'], $conn['dbname'], $conn['host'], $conn['port']);
?>
