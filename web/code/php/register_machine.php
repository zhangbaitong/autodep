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

if("register_machine"==$method)
{
    register_machine();
}


/*
 * 功能：打包镜像
 */
function register_machine(){
    $result=array();
    $machine_name = isset($_POST['machine_name'])?$_POST['machine_name']:null;
    $machine_ip = isset($_POST['machine_ip'])?$_POST['machine_ip']:null;
    $docker_port = isset($_POST['docker_port'])?$_POST['docker_port']:null;
    $is_use = isset($_POST['is_use'])?$_POST['is_use']:null;
    $remark = isset($_POST['remark'])?$_POST['remark']:null;
    

    if(empty($machine_name)||''==$machine_name)
    {
        $result['code']="-1";
        $result['reason']="machine_name can't be empty";
        print_r(json_encode($result));
        return;
    }
 
    if(empty($machine_ip)||''==$machine_ip)
    {
        $result['code']="-1";
        $result['reason']="machine_ip can't be empty";
        print_r(json_encode($result));
        return;
    }
 
    if(empty($docker_port)||''==$docker_port)
    {
        $result['code']="-1";
        $result['reason']="docker_port can't be empty";
        print_r(json_encode($result));
        return;
    }
    if(!is_numeric($docker_port))
    {
        $result['code']="-1";
        $result['reason']="docker_port must be number";
        print_r(json_encode($result));
        return;
    }
    
    
    if($is_use==null ||''==$is_use)
    {
        $result['code']="-1";
        $result['reason']="is_use can't be empty";
        print_r(json_encode($result));
        return;
    }
    
    $param=array();
    $param['machine_name']=$machine_name;
    $param['machine_ip']=$machine_ip;
    $param['docker_port']=$docker_port;
    $param['is_use']=$is_use;
    $param['remark']=$remark;
    

    $client= new GearmanClient();
    $client->addServer("127.0.0.1", 4730);
    $result=$client->do("register_machine", json_encode($param));
    print_r(json_encode($result));
}


?>