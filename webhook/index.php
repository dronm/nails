<?php
/*
 * company_id 	number 	ID компании в рамках которой происходит событие
 * resource 	string 	Имя сущности. Одно из: company, service_category, service, staff, client, record
 * resource_id 	number 	ID сущности в YCLIENTS
 * status 	string 	Событие. create, update или delete
 * data
 * 
 * Порядок работы: загружаются транзакции, в триггире before insert/update ycl_transactions_process()
 * определяется specialist_id через specialists.ycl_staff_id и текущее значение data->'record'->>'staff_id'
 * Таблица visits не используется!!!
 * 
 * curl -X POST https://pronogtiwh.katren.org -d "key1=value1,key2=value2"
 * curl --data "@/home/andrey/www/htdocs/nails/webhook/post.txt" -X POST http://localhost/nails/webhook/index.php
 */

require_once('conf.php');
require_once(FRAME_WORK_PATH.'db/db_pgsql.php');
 
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

function getDbConn() {
	$conf = json_decode(file_get_contents(CONF), TRUE);
	$conn = parsePgConnStr($conf['db']['primary']);
	
	$link = new DB_Sql();
	$link->appname = 'nails';
	$link->technicalemail = '';
	$link->detailedError = FALSE;

	/*conneсtion*/
	$link->server	= $conn['host'];
	$link->user	= $conn['user'];
	$link->password	= $conn['pwd'];
	$link->database	= $conn['dbname'];
	$link->connect($conn['host'], $conn['user'], $conn['pwd']);	
	return $link;
}

$content = file_get_contents("php://input");

$res = json_decode($content, TRUE);
if(isset($res['resource']) && isset($res['status']) && isset($res['resource_id'])){
	if($res['resource'] == "staff" && ($res['status'] == "create" || $res['status'] == "update")) {
		//staff insert/update
		$data = json_encode($res['data']);
		$conn = getDbConn();
		$conn->query(sprintf(
			"INSERT INTO ycl_staff (id, data) VALUES (%d, '%s'::jsonb)
			ON CONFLICT (id) DO UPDATE
			SET data = '%s'::jsonb"
		,$res['data']['id']
		,$data
		,$data
		));
		
	}else if($res['resource'] == "finances_operation" && ($res['status'] == "create" || $res['status'] == "update")) {	
		$data = json_encode($res['data']);
		$conn = getDbConn();
		$conn->query(sprintf(
			"INSERT INTO ycl_transactions (id, data) VALUES (%d, '%s'::jsonb)
			ON CONFLICT (id) DO UPDATE
			SET data = '%s'::jsonb"
		,$res['data']['id']
		,$data
		,$data
		));
	
	}else{
		//file_put_contents("log/log.txt", $content.PHP_EOL.PHP_EOL, FILE_APPEND);
	}
}

?>
