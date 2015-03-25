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

if("images_package"==$method)
{
    images_package();
}


/*
 * 功能：打包镜像
 */
function images_package(){
    $result=array();
    $template = isset($_POST['template'])?$_POST['template']:null;
    $image_name = isset($_POST['image_name'])?$_POST['image_name']:null;
    $code_path = isset($_POST['code_path'])?$_POST['code_path']:null;
    $creator = isset($_POST['creator'])?$_POST['creator']:null;
    $remark = isset($_POST['remark'])?$_POST['remark']:null;

    if(empty($template)||''==$template)
    {
        $result['code']='-1';
        $result['reason']="template Can't be empty";
        print_r(json_encode($result));
        return;
    } 
    if(empty($image_name)||''==$image_name)
    {
        $result['code']='-1';
        $result['reason']="image_name Can't be empty";
        print_r(json_encode($result));
        return;
    }
    if(empty($code_path)||''==$code_path)
    {
        $result['code']='-1';
        $result['reason']="code_path Can't be empty";
        print_r(json_encode($result));
        return;
    }
    if(empty($creator)||''==$creator)
    {
        $creator='admin';
    }
    if(empty($remark))
    {
        $remark='';
    }
    
    
    
    $param=array();
    $param['template']=$template;
    $param['image_name']=$image_name;
    $param['code_path']=$code_path;
    $param['creator']=$creator;
    $param['remark']=$remark;

    $client= new GearmanClient();
    $client->addServer("127.0.0.1", 4730);
    $result=$client->do("create_images", json_encode($param));
    print_r($result);
}


?>