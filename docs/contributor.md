#github代码贡献常见操作

Creator : zhangbaitong

Github  : https://github.com/zhangbaitong

##如何参与项目代码贡献

* 打开 https://github.com/zhangbaitong/autodep，然后点击右上角 Fork 按钮

* 在自己的github仓库找到autodep项目，然后复制地址

* 在本地执行 git clone https://github.com/yourname/autodep.git

* 进行相关代码修改

* git add . 添加所有修改（git add yourfiles）

* git commit -m "your message" 提交所有修改到本地

* git push 提交所有修改到自己的仓库

* 在github网站上自己的仓库上点击 Pull requests 按钮并添加提交注释


##定期更新自己的项目与原始项目同步

* git remote add upstream https://github.com/zhangbaitong/autodep

* git fetch upstream

* git checkout master

* git rebase upstream/master

* git push -f origin master


##提交被拒绝后如何进行代码回滚

* git log <filename> 查看版本记录并获得commit id

* git reset --hard <commit id> 进行回滚

* git push -f 强制提交

##设置用户信息

* git config user.name "yourname"
* git config user.email "your email"

##提交时添加签名

* git commit --amend -s --no-edit && git push -f

##提交包含删除内容在内的所有内容

* git commit -am "issue9527"
	
