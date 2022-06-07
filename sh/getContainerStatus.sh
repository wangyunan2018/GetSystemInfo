#!/bin/bash

path=../textfile
cd sh
rm $path/containerName.txt -f
rm $path/dockerStatus.txt -f
rm $path/containerExitName.txt -f
rm $path/dockerExitStatus.txt -f
rm $path/dockerRunRes.txt -f
rm $path/dockerRunStop.txt -f
rm $path/dockerRunDel.txt -f
rm $path/dockerRunLog.txt -f
rm $path/dockerExitRes.txt -f
rm $path/dockerExitStop.txt -f
rm $path/dockerExitDel.txt -f
rm $path/dockerExitLog.txt -f

# 获取在线容器名
docker ps > $path/dockerRunning.txt && sed -i '1d' $path/dockerRunning.txt && awk '{print $NF}' $path/dockerRunning.txt > $path/containerName.txt

# 获取在线容器ID
awk '{print $1}'  $path/dockerRunning.txt > $path/dockerRunId.txt

# 获取在线容器状态
for i in $(cat $path/containerName.txt); do
	echo "Running" >> $path/dockerStatus.txt; echo @$i
done

# 在线容器操作 重启、停止、删除、日志
for j in $(cat $path/containerName.txt); do
	echo "重启" >> $path/dockerRunRes.txt; echo @$j
	echo "停止" >> $path/dockerRunStop.txt; echo @$j
	echo "删除" >> $path/dockerRunDel.txt; echo @$j
	echo "日志" >> $path/dockerRunLog.txt; echo @$j
done

# 获取离线容器名
docker ps -a | grep -w "Exited" > $path/ExitName.txt && awk '{print $NF}' $path/ExitName.txt > $path/containerExitName.txt

# 获取离线容器ID
awk '{print $1}' $path/ExitName.txt > $path/dockerExitId.txt

# 获取离线容器状态
for k in $(cat $path/containerExitName.txt); do
	echo "Exited" >> $path/dockerExitStatus.txt; echo @$k
done

# 离线容器操作 重启、停止、删除、日志
for h in $(cat $path/containerExitName.txt); do
	echo "重启" >> $path/dockerExitRes.txt; echo @$h
	echo "停止" >> $path/dockerExitStop.txt; echo @$h
	echo "删除" >> $path/dockerExitDel.txt; echo @$h
	echo "日志" >> $path/dockerExitLog.txt; echo @$h
done