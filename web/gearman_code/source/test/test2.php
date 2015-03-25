<?php
//$conf= require 'D:\\infobird\\program\\DockerManager\\config.php';
//include 'D:\\infobird\\program\\DockerManager\\tools\\file_kit.php';

// $date=date('Y-m-d',time());
// $folder=$conf['path']['dockerfile']."\\".$date."\\nginx";

//$file_kit=new file_kit();
// $file_kit->mk_folder($folder);
// $ori_content=$file_kit->read_file($conf['path']['dockerfile_template']."\\x.txt");
// $new_content=$file_kit->insert_content($ori_content, "ooo,99", "\r\n"."baic"."\r\n");
// $file_kit->mk_file($folder."\\Dockerfile", $new_content);

//echo $file_kit->relative_path("/home/pipework","/data/DockerManager/docker/Dockerfile/product/2015-03-05/nginx");

//$pos=strrpos("/a/bc/d","/");

//echo substr("/a/bc/d", 0,$pos);
//echo substr("/a/bc/d", $pos+1);
   copy('/data/DockerManager/docker/Dockerfile/product/2015-03-06/nginx/Dockerfile','/data/DockerManager/Dockerfile'); 
?>
