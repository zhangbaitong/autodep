# AutoDEP
It's a auto deployment tool with docker.

#How to get
git clone https://github.com/zhangbaitong/autodep.git

#Install

1.install package

go get github.com/codeskyblue/go-sh

go get github.com/samalba/dockerclient

go get github.com/robfig/config

2.install gcc

yum install gcc

3.create your ssh keygen

#Test

go test api/common -v

#feature

1. support http REST API for do some action.

	eg.shell,docker client and so on.

2. support package for docker image Automatic.

3. It can run dep docker container for specical host.

...

#Our plan

see [plan](./docs/plan.md)

#Collaborators

[show168](https://github.com/show168)

[tomzhaogy](https://github.com/tomzhaogy)
