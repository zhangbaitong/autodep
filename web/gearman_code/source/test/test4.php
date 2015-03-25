<?php
     require_once 'Archive/Tar.php';  
      
    $targetDir = 'xdebug';  
    $tarFile = 'xdebug.tar.gz';  
    $tar = new Archive_Tar($tarFile);  
    $tar->extract($targetDir);  
      
    $dp = opendir($targetDir);  
    while ($entry = readdir($dp)){  
        if(is_dir($entry))  
        {  
            echo '[DIR] '.$entry. '<br/>';  
        }elseif (is_file($entry))  
        {  
            echo '[FILE] '.$entry. '<br/>';  
        }  
    }  
    closedir($dp); 
?>
