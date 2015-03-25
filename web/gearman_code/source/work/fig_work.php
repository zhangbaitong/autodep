<?php
class fig_work
{       
    /**
     * 功能：创建fig文件
     * 参数：$job：参数
     * 返回： 
     */
    function create_fig($job)
    {
        $param=json_decode($job->workload());
        $result=array();
        $result['code']=0;
        
        //校验参数
        if(empty($param) || ""==$param)
        {
            $result['code']=1;
            $result['reason']="data Can't be empty";
            return $result;
        }
     
        $conf= require './config.php';
        $fig_directory=$conf['path']['fig']."/".$param->project_name;
        
        $data=array();
        $data['Version']='1.0';
        $data['ServerIP']=$param->machine_ip;
        $data['Port']=(int)$param->docker_port;
        $data['Method']='create_fig';
        $data['Params']=$this->deal_data($param->server,$fig_directory);
        

        include_once $conf['tool']['http_kit'];
        
        $http=new http();
        $result['back']=$http->post_web_result("http://117.78.19.76:8080/v1/fig/create", "request=".json_encode($data));
        
        print_r($data);     
        echo "---------------"."\r\n"."\r\n";
        print_r(json_encode($data));

        return $result;
    }
    
    /**
     * 功能：加工fig数据
     * 参数：从前端传来的数据
     * 返回：符合fig文件要求的数据
     */
    function deal_data($param,$fig_directory){
        $result=array();
        $commands=array();
        $fig_data="";
        foreach($param as $p)
        {
            $fig_data=$fig_data.$this->deal_server($p->server_name);
            $fig_data=$fig_data.$this->deal_one_value("image",$p->image);
            $fig_data=$fig_data.$this->deal_more_value("ports",$p->ports);
            $fig_data=$fig_data.$this->deal_more_value("links",$p->links);
            $fig_data=$fig_data.$this->deal_more_value("volumes",$p->volumes);
            $fig_data=$fig_data.$this->deal_command($p->server_name,$p->command,$fig_directory);
            $commands=array_merge($commands,$this->deal_command_content($p->server_name,$p->command));
        }
        
        $result['fig_data']=$fig_data;
        $result['commands']=$commands;
        $result['fig_directory']=$fig_directory;
        
        return $result;
    }
    
    
    /*
     *  功能：处理command
    */
    function deal_command($server_name,$command,$fig_directory)
    {
         $result="";
         if(""!=trim($command))
         {
            $result="  command: ".$fig_directory."/startup/".$server_name."/start.sh"."\r\n";
         }
         
         return $result;
    }
     
    
    /**
     * 功能：处理 command内容
     */
    function deal_command_content($server_name,$command)
    {
        $result=array();
        if(""!=trim($command))
        {
            $new_command="#!/bin/bash"."\n".$command;
            $result[$server_name]=$new_command;
        }
        return $result;
    }
    
    
   /**
    * 功能：server的数据
    */ 
   function deal_server($param)
   {
       return $param.":"."\n";
   } 
   
   
   /**
    * 功能：处理只有单个值的数据，比如image,command
    */
   function deal_one_value($name,$value)
   {
       return "  ".$name.": ".$value."\n";
   }
   
   
   /**
    * 处理拥有多个值的数据，比如ports,links
    * @param unknown $name
    * @param unknown $values
    * @return string
    */
   function deal_more_value($name,$values)
   {
       $result="";
       foreach($values as $value)
       {
           if(""!=trim($value))
           {
               if("ports"==$name)
               {
                   $result=$result."    - \"".$value."\""."\n";
               }
               else
               { 
                   $result=$result."    - ".$value."\n";
               }
           }
       }
       
       if(""!=$result)
       {
           $result="  ".$name.":"."\n".$result;
       }
       return $result;
   }
   

 
}
   

?>
