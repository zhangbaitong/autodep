<?php
class file_kit
{
    /*
     * 功能：创建文件夹
     * 参数：$folder：文件夹目录
     * 返回：无
     */
    function mk_folder($folder)
    {
        if(!is_readable($folder))
        {
            $this->mk_folder(dirname($folder));
            if(!is_file($folder)) mkdir($folder,0777);
        }
    }
    
    
    /**
     * 功能：创建文件
     * 参数：$file_name：文件名称（包括目录）
     *       $content：文件内容
     * 返回：
     *      true：成功
     *      false：失败
     */
    function mk_file($file_name,$content)
    {
        if(!empty($content))
        {
            $folder=dirname($file_name);

            if(!is_dir($folder))
            {
                $this->mk_folder($folder);
            }
            
            $myfile = fopen($file_name, "w") or die("Unable to open file!");
            fwrite($myfile, $content);
            fclose($myfile);
            return true;
        }
        return false;
    }
    
    
    /*
     * 功能：读取文件内容，返回字符串
     * 参数：$file_path：文件路径
     * 返回：文件内容
     */
    function read_file($file_path)
    {
        if(file_exists($file_path))
        {
            //读取二进制文件时，需要将第二个参数设置成'rb'
            $handle = fopen($file_path, "r");
            //通过filesize获得文件大小，将整个文件一下子读到一个字符串中
            $contents = fread($handle, filesize ($file_path));
            fclose($handle);
            return $contents;
        }
        return "";
    }
    
    
    /*
     * 功能：在字符串中指定的位置插入内容
     * 参数： $ori_content:原有的字符串
     *       $search_keys:要查找的值（多个）,用逗号隔开,从第一个开始查找，找到就停止
     *       $insert_value:插入的值
     * 返回：
     *       新内容  
     *       "":表示插入失败 
     *             
     */
    function insert_content($ori_content,$search_keys,$insert_value)
    {
        if(empty($ori_content) || empty($search_keys) || empty($insert_value))
        {
            return "";
        }
        foreach (explode(",",$search_keys) as $key)
        {
            if(empty($key))
            {
                continue;
            }
            
            if(false==strpos($ori_content,$key))
            {
                continue;
            }
            
            $pos=strpos($ori_content,$key);
            return substr($ori_content,0,$pos-1).$insert_value.substr($ori_content,$pos);
        }
        return "";
    }

    /*
     * 功能：计算出B相对于A的相对路径
     * 参数：
     * 返回：
     */
    function relative_path($a, $b)
    {
        $patha = explode('/', $a);
        $pathb = explode('/', $b);
         
        $counta = count($patha) - 1;
        $countb = count($pathb) - 1;
         
        $path = "../";
        if ($countb > $counta) {
            while ($countb > $counta) {
                $path .= "../";
                $countb --;
            }
        }
         
        // 寻找第一个公共结点
        for ($i = $countb - 1; $i >= 0;) {
            if ($patha[$i] != $pathb[$i]) {
                $path .= "../";
                $i --;
            } else { // 判断是否为真正的第一个公共结点，防止出现子目录重名情况
                for ($j = $i - 1, $flag = 1; $j >= 0; $j --) {
                    if ($patha[$j] == $pathb[$j]) {
                        continue;
                    } else {
                        $flag = 0;
                        break;
                    }
                }
                 
                if ($flag)
                    break;
                else
                    $i ++;
            }
        }
         
        for ($i += 1; $i <= $counta; $i ++) {
            $path .= $patha[$i] . "/";
        }
         
        return $path;
    }
}
?>