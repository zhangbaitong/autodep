<?php
$DB=new sqllite('/data/DockerManager/db/test2.db');
$DB->query("create table test(id integer primary key,title varchar(50))");
$DB->query("insert into test(title) values('小屋')");

$DB->query("insert into test(title) values('蓝雨')");

$DB->query("insert into test(title) values('Ajan')");

$DB->query("insert into test(title) values('sjolzy')");
print_r($DB->getlist('select * from test order by id desc')); 


class sqllite
{
    function __construct($file)
    {
        try
        {
    
            $this->connection=new PDO('sqlite:'.$file);
    
        }
        catch(PDOException $e)
        {
            try
            {
                 $this->connection=new PDO('sqlite2:'.$file);
            }
            catch(PDOException $e)
            {
                exit('error!');
            }
        }
    }
    
    
    function __destruct()
    {
        $this->connection=null;
    }

    
    function query($sql)
    { 
        //直接运行SQL，可用于更新、删除数据
        return $this->connection->query($sql);
    }
    
    
    function getlist($sql)
    {
        //取得记录列表
        $recordlist=array();
        foreach($this->query($sql) as $rstmp)
        {
            $recordlist[]=$rstmp;
        }
        return $recordlist;
    }
    
   
    function Execute($sql)
    { 
        return $this->query($sql)->fetch();
    }
    
    
    function RecordArray($sql)
    {
        return $this->query($sql)->fetchAll();
    }
    
    
    function RecordCount($sql)
    {
        return count($this->RecordArray($sql));
    }
    
    function RecordLastID()
    {
        return $this->connection->lastInsertId();
    }   
}
?>
