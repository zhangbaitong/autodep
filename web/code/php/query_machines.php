<?php

$result = array();
$method = isset($_POST["method"])?$_POST["method"]:null;

if(empty($method))
{  
    $result['code']="-1";
    $result['reason']="method can't be empty";
    print_r(json_encode($result));
    return;
}

if("query_machines"==$method)
{
    query_machines();
}


/*
 * 功能：打包镜像
 */
function query_machines(){
    $result=array();
    $machine_name = isset($_POST['machine_name'])?$_POST['machine_name']:null;
    $machine_ip = isset($_POST['machine_ip'])?$_POST['machine_ip']:null;
    $docker_port = isset($_POST['docker_port'])?$_POST['docker_port']:null;
    $is_use = isset($_POST['is_use'])?$_POST['is_use']:null;


    
    $param=array();
    $param['machine_name']=$machine_name;
    $param['machine_ip']=$machine_ip;
    $param['docker_port']=$docker_port;
    $param['is_use']=$is_use;

    $client= new GearmanClient();
    $client->addServer("127.0.0.1", 4730);
    $result=$client->do("query_machines", json_encode($param));
    print_r(json_encode($result));
}


?>