##生成key
```
ssh-keygen -t rsa
```

```
/root/.ssh/id_rsa(私有密钥)
/root/.ssh/id_rsa.pub(公有密钥)
```

##复制密匙

复制公有密钥到要访问的机器上的用户目录

scp id_rsa.pub root@root:/root/.ssh/authorized_keys

##修改权限

修改目标机器的.ssh目录权限

```
chmod 755 ~/.ssh
```

##加快ssh访问速度
```
ssh -M -S ~/.ssh/remote-host user@remotehost 

ssh -S ~/.ssh/remote-host user@remotehost
```