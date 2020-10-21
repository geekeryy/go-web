#!/bin/bash
GitProjectName=$1
GitBranchName=$2
GitCommitId=$3
PkgName=$4

#脚本参数检查
if [ -z "$GitProjectName" -o -z "$GitBranchName" -o -z "$GitCommitId" ]
then
    echo "Missing options"
    exit 1
fi

#选择Dockerfile
Dockerfile='./Dockerfile'
DockerContext=../

#构建Docker镜像
echo "start build dcoker image witch dockerfile : $Dockerfile ..."
docker build -f $Dockerfile -t $PkgName $DockerContext
if [ $? -eq 0 ]
then
    echo "build docker image success !"

    exit 0
else
    echo "build image failed !"
    exit 1
fi
