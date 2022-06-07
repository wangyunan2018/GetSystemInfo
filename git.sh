#!/bin/bash

echo "Start git code..."
echo

git add .
git commit -m "add"
git push -u origin master

if [ $? -eq 0 ]; then
        echo
        echo "Ruturn $?, SUCCESS !"
else
        echo
        echo "Ruturn $?, FAIL !"
fi