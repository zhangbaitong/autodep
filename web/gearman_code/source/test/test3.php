<?php
$conf= require '../config.php';
include $conf['tool']['file_kit'];

$file_kit=new file_kit();
$post_data=$file_kit->read_file("/data/DockerManager/docker/Dockerfile/template/nginx/Dockerfile.tar.gz");



    $curl= curl_init();// 启动一个CURL会话 
    curl_setopt($curl, CURLOPT_URL,"http://127.0.0.1:4243/build?t=test6");// 要访问的地址 
    curl_setopt($curl, CURLOPT_REFERER,""); 
    curl_setopt($curl, CURLOPT_POST, 1);// 发送一个常规的Post请求 
    curl_setopt($curl, CURLOPT_POSTFIELDS,$post_data);// Post提交的数据包 
    curl_setopt($curl, CURLOPT_TIMEOUT, 30);// 设置超时限制防止死循环 
    curl_setopt($curl, CURLOPT_HTTPHEADER, array('Content-Type: application/tar'));
    curl_setopt($curl, CURLOPT_HEADER, 0);// 显示返回的Header区域内容 
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);// 获取的信息以文件流的形式返回 
    $tmpInfo= curl_exec($curl);// 执行操作 
    if(curl_errno($curl)) { 
       echo'Errno'.curl_error($curl);//捕抓异常 
    } 
    curl_close($curl);// 关闭CURL会话 
    echo  $tmpInfo;// 返回数据    
?>
