--镜像表
create table images(
    image_id  integer PRIMARY KEY autoincrement,
    image_name varchar(50) not null,
    creator varchar(50) not null,
    create_time integer not null,
    remark varchar(2000)
);

--机器表
create table machines(
     machine_id  integer PRIMARY KEY autoincrement,
     machine_name varchar(50) not null,
     machine_ip  varchar(20) not null,
     docker_port  integer not null,
     is_use       integer default 0,
     remark varchar(2000)
);

--fig项目表
create table fig_project(
     fig_project_id  integer PRIMARY KEY autoincrement,
     project_name varchar(50) not null,
     machine_ip  varchar(20) not null,
     fig_directory  varchar(200) not null,
     fig_param  text not null,
     fig_content  text not null,
     create_time   integer not null
);

--模版表
create table template(
     template_id  integer PRIMARY KEY autoincrement,
     template_name varchar(50) not null,
     template_type varchar(50) not null,
     template_content  varchar(5000) not null,
     create_time   integer not null,
     remark varchar(2000)
);


--模版初始化sql
insert into template(template_name,template_type,template_content,create_time) values('fig_nginx','fig','{"image":"centos6/nginx:2015-03-13","ports":["80:80"],"links":["mysql:mysql_host"],"volumes":["/data/project_name/nginx_code:/data/project_name/nginx_code","/data/project_name/software/nginx:/etc/nginx","/data/project_name/software/nginx-php/php.ini:/etc/php.ini","/data/project_name/software/nginx-php/php.d:/etc/php.d","/data/project_name/software/nginx-php/php-fpm.conf:/etc/php-fpm.conf","/data/project_name/software/nginx-php/php-fpm.d:/etc/php-fpm.d","/data/project_name/log/nginx:/var/log/nginx","/data/project_name/log/nginx-php-fpm:/var/log/php-fpm"],"command":"service nginx start\nservice php-fpm start"}',datetime());
insert into template(template_name,template_type,template_content,create_time) values('fig_mysql','fig','{"image":"centos6/mysql:2015-01-09","ports":["3306:3306"],"links":[""],"volumes":["/data/project_name/mysql:/home/databases/mysql/data","/data/project_name/software/mysql/my.cnf:/etc/my.cnf","/data/project_name/log/mysql:/var/log/mysql"],"command":"service mysqld restart\nmysql -e \"grant all privileges on *.* to ''root''@''%'' identified by ''111111'';\"\nmysql -e \"grant all privileges on *.* to ''root''@''localhost''identified by ''111111'';\""}',datetime());
insert into template(template_name,template_type,template_content,create_time) values('fig_redis','fig','{"image":"centos6/redis:2015-01-09","ports":["6379:6379"],"links":[""],"volumes":["/data/project_name/software/redis:/etc/redis","/data/project_name/log/redis:/var/redis/log"],"command":""}',datetime());
insert into template(template_name,template_type,template_content,create_time) values('fig_gearman','fig','{"image":"centos6/gearman:2015-03-13","ports":["4730:4730"],"links":[""],"volumes":["/data/project_name/gearman_code:/data/project_name/gearman_code","/data/project_name/software/gearman-php/php.ini:/etc/php.ini","/data/project_name/software/gearman-php/php.d:/etc/php.d","/data/project_name/software/gearman-php/php-fpm.conf:/etc/php-fpm.conf","/data/project_name/software/gearman-php/php-fpm.d:/etc/php-fpm.d","/data/project_name/log/gearman-php-fpm:/var/log/php-fpm"],"command":""}',datetime());
insert into template(template_name,template_type,template_content,create_time) values('fig_zeromq','fig','{"image":"centos6/zeromq:2015-03-13","ports":[""],"links":[""],"volumes":["/data/project_name/zeromq_code:/data/project_name/zeromq_code","/data/project_name/software/zeromq-php/php.ini:/etc/php.ini","/data/project_name/software/zeromq-php/php.d:/etc/php.d","/data/project_name/software/zeromq-php/php-fpm.conf:/etc/php-fpm.conf","/data/project_name/software/zeromq-php/php-fpm.d:/etc/php-fpm.d","/data/project_name/log/zeromq-php-fpm:/var/log/php-fpm"],"command":""}',datetime());
insert into template(template_name,template_type,template_content,create_time) values('fig_nodejs','fig','{"image":"centos6/nodejs:2015-02-02","ports":["81:80"],"links":[""],"volumes":[""],"command":""}',datetime());