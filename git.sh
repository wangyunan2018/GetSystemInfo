#!/bin/bash

echo "Start git code..."
echo

git add .
git commit -m "add"
#git push -u origin master
git push origin main

if [ $? -eq 0 ]; then
        echo
        echo "Return $?, SUCCESS !"
else
        echo
        echo "Return $?, FAIL !"
fi
