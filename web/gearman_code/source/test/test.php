<?php
//     $myfile = fopen("D:\\temp\\x.txt", "w");
//     fclose($myfile);
//       file_put_contents("D:\\temp\\a\\x.txt", "ddddddd\r\n\r\n", FILE_APPEND);


// if(file_exists("D:\\temp\\x.txt"))
// {
//     echo "文件存在";
// }
// else 
// {
//     echo "文件不存在";    
// }


// $filename = "D:\\temp\\x.txt";
// $handle = fopen($filename, "r");//读取二进制文件时，需要将第二个参数设置成'rb'

// //通过filesize获得文件大小，将整个文件一下子读到一个字符串中
// $contents = fread($handle, filesize ($filename));
// fclose($handle);

// // echo $contents;

//  $pos=strpos($contents,"323");
//  $prev=substr($contents,0,$pos-1);
 
//  $next=substr($contents,$pos);
 
//  echo $prev."\r\n"."被插入的内容"."\r\n".$next;
//  echo date('Y-m-d',time());

if(is_dir("D:\\temp\\reald"))
{
    echo "目录存在";
}
else 
{
    echo "目录不存在";
}
?>