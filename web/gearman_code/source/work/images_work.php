<?php
class images_work
{
    /**
     * 功能：查询镜像
     * 参数：
     * 返回：
     */
    function query_images($job)
    {
        $param=json_decode($job->workload());
        $result=array();
        $result['code']=0;
        
        //校验参数
        if(empty($param->creator) || ""==$param->creator)
        {
            $result['code']=1;
            $result['reason']="creator Can't be empty";
            return $result;
        }
        if(empty($param->image_name) || ""==$param->image_name)
        {
            $param->image_name='%';
        }
        if(empty($param->start_time) || ""==$param->start_time)
        {
            $param->start_time='0000-00-00 00:00:00';
        }
        if(empty($param->end_time) || ""==$param->end_time)
        {
            $param->end_time='9999-00-00 00:00:00';
        }
        
        $start_time=strtotime($param->start_time);
        $end_time=strtotime($param->end_time);
        
        
        $conf= require './config.php';
        include_once $conf['tool']['sqllite_kit'];

        try
        {
             $sql= "select * from images where image_name like '%{$param->image_name}%' and create_time>={$start_time} and create_time<={$end_time} and creator='{$param->creator}' order by image_id desc";       
             $result['data']=(new sqllite($conf['path']['db']))->getlist($sql);
        }
        catch (Exception $ex)
        {
            $result['code']=1;
            $result['reason']=$ex->getMessage();
        }

        return $result;
    }
    
    
    /**
     * 功能：创建镜像
     * 参数：$job：参数
     * 返回： 
     */
    function create_images($job)
    {
        $param=json_decode($job->workload());
        $result=array();
    
        //校验参数
        if(empty($param->template) || ""==$param->template)
        {
            $result['code']=1;
            $result['reason']="template Can't be empty";
            return $result;
        }
        if(empty($param->image_name) || ""==$param->image_name)
        {
            $result['code']=1;
            $result['reason']="image_name Can't be empty";
            return $result;
        }
        if(empty($param->code_path) || ""==$param->code_path)
        {
            $result['code']=1;
            $result['reason']="code_path Can't be empty";
            return $result;
        }
        if(empty($param->creator) || ""==$param->creator)
        {
            $param->creator='admin';
        }
        if(empty($param->remark))
        {
            $param->remark='';
        }
    
        //生存Dockerfile文件
        $dockerfile_directory=$this->create_dockerfile($param->template,$param->code_path);
        if(empty($dockerfile_directory) || ""==$dockerfile_directory)
        {
            $result['code']=1;
            $result['reason']="dockerfile_directory is empty,create_dockerfile error";
            return $result;
        }
    
        //生成镜像
        $result=$this->build_images($param->image_name,$dockerfile_directory);
    
        //保存镜像信息到数据库
        $this->save_images_to_db($param);
    
        return $result;
    }
    
    
    /*
     * 功能：根据模版生成Dockerfile文件
    * 参数：无
    * 返回：Dockerfile文件所在目录
    */
    function create_dockerfile($template,$code_path){
        $conf= require './config.php';
        include_once $conf['tool']['file_kit'];
    
        //生成Dockerfile文件中的目标指令
        $date=date('Y-m-d',time());
        $folder=$conf['path']['dockerfile']."/".$date."/".$template;
        $pos=strrpos($code_path,"/");
        $code_path_prev=substr($code_path,0,$pos);
        $code_path_next=substr($code_path,$pos);
        if(empty($code_path_prev) || ''==$code_path_prev)
        {
            $code_path_prev="/";
        }
        $relative_path=".".$code_path_next;
        $add_content="\n"."ADD  ".$relative_path."  ".$conf['path']['code']."/".$template."\n";
    
    
    
        //读取模版，生成目标Dockerfile文件
        $file_kit=new file_kit();
        $file_kit->mk_folder($folder);
        $ori_content=$file_kit->read_file($conf['path']['dockerfile_template']."/".$template."/"."Dockerfile");
        $new_content=$file_kit->insert_content($ori_content, "EXPOSE,CMD",$add_content);
        $file_kit->mk_file($folder."/Dockerfile", $new_content);
        $file_kit->mk_file($code_path_prev."/Dockerfile", $new_content);
        $file_kit=null;
        return  $code_path_prev;
    }
    
    
    
    /**
     * 功能：调用docker命令创建镜像
     * 参数：image_name：新创建镜像的名称
     *       dockerfile_directory：dockerfile文件所在的目录
     * 返回：result:
     *           code:
     *                0：创建成功
     *                1： 创建失败
     *           reason:执行过程
     */
    function build_images($image_name,$dockerfile_directory)
    {
    
        //调用Docker的build命令
        $cmd='/usr/bin/docker build -t '.$image_name.'  '.$dockerfile_directory;
        exec($cmd,$out,$status);
    
        $result=array();
        $result['code']=$status;
        if(1==$result['code'])
        {
            $result['reason']=$out;;
        }
        return json_encode($result);
    }
    
    
    /**
     * 功能：保存镜像信息到数据库
     * 参数：
     * 返回：
     */
    function save_images_to_db($param)
    {
        $conf= require './config.php';
        include_once $conf['tool']['sqllite_kit'];
    
        $DB=new sqllite($conf['path']['db']);
        $create_time=time();
        $sql= "insert into images(image_name,creator,create_time,remark) values('{$param->image_name}','{$param->creator}',{$create_time},'{$param->remark}')";
        $DB->query($sql);
    }
    
    
    /**
     * 功能：调用docker远程api创建镜像
     * 问题：当Dockerfile中有指令"ADD src des"时，无法找到src文件或目录，未解决
     */
    function build_images_http()
    {
        //     $conf= require './config.php';
        //     include $conf['tool']['file_kit'];
        //     $file_kit=new file_kit();
        //     $post_data=$file_kit->read_file("/data/DockerManager/Dockerfile.tar.gz");
        //     $curl= curl_init();
        //     curl_setopt($curl, CURLOPT_URL,"http://127.0.0.1:4243/build?t=test3");
        //     curl_setopt($curl, CURLOPT_REFERER,"");
        //     curl_setopt($curl, CURLOPT_POST, 1);
        //     curl_setopt($curl, CURLOPT_POSTFIELDS,$post_data);
        //     curl_setopt($curl, CURLOPT_TIMEOUT, 30);
        //     curl_setopt($curl, CURLOPT_HTTPHEADER, array('Content-Type: application/tar'));
        //     curl_setopt($curl, CURLOPT_HEADER, 0);
        //     curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
        //     $tmpInfo= curl_exec($curl);
        //     if(curl_errno($curl)) {
        //         echo'Errno'.curl_error($curl);
        //     }
        //     curl_close($curl);
        //     echo  $tmpInfo;
    }
}
   

?>
