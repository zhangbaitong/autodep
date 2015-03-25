<?php
$client= new GearmanClient();
$client->addServer("127.0.0.1", 4730);
$param=array();
$param['$creator']='admin';
$result=$client->do("query_images", json_encode($param));
print_r(json_encode($result));
?>
