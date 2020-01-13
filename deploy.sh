#!/bin/bash

AppName = "w3fy"

#git冲突解决，默认覆盖本地代码
git fetch --all
git reset --hard origin/master

#构建项目
make -f MakeFile clean
make -f MakeFile

#启动项目
./$AppName

