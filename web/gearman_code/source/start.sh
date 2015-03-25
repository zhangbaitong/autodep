#!/bin/bash
echo $# 
echo $2
echo $1
php_cmd="/usr/bin/php "

##start work default
if [ "$1"  = "start" ]
then
	work="docker"
    log=$work.out 
    nohup $php_cmd index_work.php  $work &>$log&
    exit
fi 

##stop work
if [ "$1"  = "stop" ]
then
    ps -ef  |grep index_work.php  |awk '{print $2}'  |while read pid
        do
            kill -9 $pid
        done   
    exit
fi 

##restart work
if [ "$1"  = "restart" ]
then
    stop();
    start();
fi 