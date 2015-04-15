新机器需先做下列操作:
1. 更新/etc/sysconfig/docker文件，重启docker
```
other_args='-H unix:///var/run/docker.sock -H 0.0.0.0:4243 --insecure-registry 
```

2. 防火墙开放4243端口

3. 拷贝ssh授权文件