<?php
$worker= new GearmanWorker();
$worker->addServer("127.0.0.1", 4730);
$worker->addFunction("create_images", "create_images");
$worker->addFunction("query_images", "query_images");
$worker->addFunction("create_fig", "create_fig");
$worker->addFunction("query_machines", "query_machines");
$worker->addFunction("register_machine", "register_machine");
while ($worker->work());


//创建镜像   
function create_images($job)  
{  
    include_once './work/images_work.php';  
    return json_encode((new images_work())->create_images($job));
}  


//查询镜像
function query_images($job)
{
    include_once './work/images_work.php';
    return json_encode((new images_work())->query_images($job)); 
}

//生成fig文件
function create_fig($job)
{
    include_once './work/fig_work.php';
    return json_encode((new fig_work())->create_fig($job));
}


//生成fig文件
function query_machines($job)
{
    include_once './work/machines_work.php';
    return json_encode((new machines_work())->query_machines($job));
}


//生成fig文件
function register_machine($job)
{
    include_once './work/machines_work.php';
    return json_encode((new machines_work())->register_machine($job));
}
?>
