#项目设计简介

主要语言：golang，HTML，JS

系统分为两部分：

1.前端

技术：HTML，JS，AJAX

2.后端

技术：Docker，fig，ssh，docker-Registry，shell，http，sqlite3

#架构设计图


![architecture](../architecture.png?raw=true "architecture")

#使用场景（待续-开发到测试到生产环境的流程管理）

1.开发环境使用

在开发环境中进行测试，修改，迭代。完成后通过autodep进行打包。

将打包的应用提交到私有仓库。

2.测试环境使用

通过autodep通过设置容器组合来完成项目依赖配比。

然后在autodep中选择指定机器运行项目。

3.生产环境使用

通过autodep同步测试环境项目到生产环境即可。

![workflow](../workflow.png?raw=true "workflow")



