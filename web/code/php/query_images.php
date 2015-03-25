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

if("query_images"==$method)
{
    query_images();
}


/*
 * 功能：打包镜像
 */
function query_images(){
    $result=array();
    $creator = isset($_POST['creator'])?$_POST['creator']:null;
    $image_name = isset($_POST['image_name'])?$_POST['image_name']:null;
    $start_time = isset($_POST['start_time'])?$_POST['start_time']:null;
    $end_time = isset($_POST['end_time'])?$_POST['end_time']:null;

    if(empty($creator)||''==$creator)
    {
        $creator='admin';
    }
    
    $param=array();
    $param['creator']=$creator;
    $param['image_name']=$image_name;
    $param['start_time']=$start_time;
    $param['end_time']=$end_time;

    $client= new GearmanClient();
    $client->addServer("127.0.0.1", 4730);
    $result=$client->do("query_images", json_encode($param));
    print_r(json_encode($result));
}


?>