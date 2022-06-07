#!/bin/bash

sudo rm textfile/* -f
filePath=textfile

# GPU驱动版本：
nvidia-smi > $filePath/gpuinfo.txt && sed -n '3p' $filePath/gpuinfo.txt > $filePath/gpuVersion.txt && sed -i 's/|//g' $filePath/gpuVersion.txt && sed -i 's/^ *//' $filePath/gpuVersion.txt && sed -i 's/[ ]*$//g' $filePath/gpuVersion.txt

# GPU风扇温度|功率|显存利用率：
sed -n '10p' $filePath/gpuinfo.txt > $filePath/gpuPower.txt && sed -i 's/|//g' $filePath/gpuPower.txt && sed -i 's/^ *//' $filePath/gpuPower.txt && sed -i 's/[ ]*$//g' $filePath/gpuPower.txt

# GPU进程：
nvidia-smi | grep -A 100 "Processes" | grep -A 100 "==" | grep -v "==" | grep -v "+" > $filePath/gpuProcesses.txt && sed -i 's/|//g' $filePath/gpuProcesses.txt && sed -i 's/^ *//' $filePath/gpuProcesses.txt && sed -i 's/[ ]*$//g' $filePath/gpuProcesses.txt
