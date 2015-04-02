create table images(
    image_id  integer PRIMARY KEY autoincrement,
    image_name varchar(50) not null,
    creator varchar(50) not null,
    create_time integer not null,
    remark varchar(2000)
);


create table machines(
     machine_id  integer PRIMARY KEY autoincrement,
     machine_name varchar(50) not null,
     machine_ip  varchar(20) not null,
     docker_port  integer not null,
     is_use       integer default 0,
     remark varchar(2000)
);

create table fig_project(
     fig_project_id  integer PRIMARY KEY autoincrement,
     project_name varchar(50) not null,
     machine_ip  varchar(20) not null,
     fig_directory  varchar(200) not null,
     fig_content  text not null,
     create_time   integer not null
);
