# AutoDEP
一个基于Docker的快速镜像创建，镜像管理，仓库查询，容器组合，自动部署，系统监控的运维工具。

![AutoDEP screen shot](screeshot1s.gif?raw=true "AutoDEP scree shot")

#获取代码
git clone https://github.com/zhangbaitong/autodep.git

#安装依赖和环境设置

1.安装依赖包和工具。
```
$ ./dependence.sh
```
2.创建数据库接口
```
$ ./init_db.sh
```
3.创建你的ssh keygen以便连接你需要连接的机器。

详情查看[ssh keygen create](./docs/sshkeygen.md)

4.启动程序
```
go run src/api/main.go
```

#测试代码

go test api/common -v(暂不提供)

#主要功能特点

1. 支持开发环境的代码镜像创建。

2. 提供默认镜像模板（目前只支持nginx，gearman，zeromq）

3. 支持已创建镜像的查询

4. 私有仓库镜像查询（支持镜像标签）

5. 服务器镜像查询

6. 服务器容器查询

7. 基于项目的容器组合功能

8. 基于项目的容器组合管理功能

9. 服务器注册，查询功能

#设计
 
查看[设计详情](./docs/design.md)

#我们的计划

查看[计划列表](./docs/plan.md)

#合作开发者

[show168](https://github.com/show168)

[tomzhaogy](https://github.com/tomzhaogy)

#其他

欢迎提出宝贵意见并参与项目改进

[MD格式小记](./docs/markdown.md)

[代码贡献常见操作](./docs/contributor.md)
