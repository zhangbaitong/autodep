<?php
class machines_work
{       
  /**
     * 功能：查询镜像
     * 参数：
     * 返回：
     */
    function query_machines($job)
    {
        $param=json_decode($job->workload());
        $result=array();
        $result['code']=0;
        $where="";

        if(!empty($param->machine_name) && ""!=$param->machine_name)
        {
            $where=$where." and machine_name like '%{$param->machine_name}%' ";
        }
        if(!empty($param->machine_ip) && ""!=$param->machine_ip)
        {
            $where=$where." and  machine_ip like '%{$param->machine_ip}%' ";
        }
        if($param->docker_port!=null && ""!=$param->docker_port)
        {
            $where=$where." and docker_port={$param->docker_port} ";
        }
        if($param->is_use!=null && ""!=$param->is_use)
        {
            $where=$where." and is_use={$param->is_use} ";
        }
        
        
        $conf= require './config.php';
        include_once $conf['tool']['sqllite_kit'];

        try
        {
             $sql= "select * from machines where 1=1".$where;       
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
     * 功能：注册机器
     * @param unknown $job
     */
    function register_machine($job)
    { 
        $param=json_decode($job->workload());
        $result=array();
        $result['code']="0";
    
        //校验参数
         if(empty($param->machine_name)||''==$param->machine_name)
        {
            $result['code']="-1";
            $result['reason']="machine_name can't be empty";
            print_r(json_encode($result));
            return;
        }
 
        if(empty($param->machine_ip)||''==$param->machine_ip)
        {
            $result['code']="-1";
            $result['reason']="machine_ip can't be empty";
            print_r(json_encode($result));
            return;
        }
     
        if(empty($param->docker_port)||''==$param->docker_port)
        {
            $result['code']="-1";
            $result['reason']="docker_port can't be empty";
            print_r(json_encode($result));
            return;
        }
        if(!is_numeric($param->docker_port))
        {
            $result['code']="-1";
            $result['reason']="docker_port must be number";
            print_r(json_encode($result));
            return;
        }
        
        
        if($param->is_use==null ||''==$param->is_use)
        {
            $result['code']="-1";
            $result['reason']="is_use can't be empty";
            print_r(json_encode($result));
            return;
        }
    

    
        $conf= require './config.php';
        include_once $conf['tool']['sqllite_kit'];
    
        $DB=new sqllite($conf['path']['db']);
        $create_time=time();
        $sql= "insert into machines(machine_name,machine_ip,docker_port,is_use,remark) values('{$param->machine_name}','{$param->machine_ip}',{$param->docker_port},{$param->is_use},'{$param->remark}')";

        echo "sql:".$sql;
        
        $DB->query($sql);
    
        return $result;
        
    }
 
}
   

?>
