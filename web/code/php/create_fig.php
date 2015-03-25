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

if("create_fig"==$method)
{
    create_fig();
}


/*
 * 功能：打包镜像
 */
function create_fig(){
    $result=array();
    $data = isset($_POST['data'])?$_POST['data']:null;
    

    if(empty($data)||''==$data)
    {
        $result['code']="-1";
        $result['reason']="data can't be empty";
        print_r(json_encode($result));
        return;
    }
    

    $client= new GearmanClient();
    $client->addServer("127.0.0.1", 4730);
    $result=$client->do("create_fig", json_encode($data));
    print_r(json_encode($result));
}


?>