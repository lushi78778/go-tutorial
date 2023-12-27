#!/bin/bash

# 未经测试！！！！！！

zipCount=$(ls -1 *.zip 2>/dev/null | wc -l)

# 检查是否没有找到 ZIP 文件
if [ "$zipCount" -eq 0 ]; then
    echo "未找到ZIP文件。退出。"
    exit 1
fi

# 检查是否有多个 ZIP 文件
if [ "$zipCount" -gt 1 ]; then
    echo "发现多个ZIP文件。退出。"
    exit 1
fi

# 如果只有一个 ZIP 文件，则继续解压缩
distFolder="dist"
zipFile=$(ls *.zip)
zipFileFullPath="$(pwd)/$zipFile"

# 检查 dist 文件夹是否存在，如果存在则清空其内容，否则创建它
if [ ! -d "$distFolder" ]; then
    mkdir "$distFolder"
else
    rm -rf "$distFolder"/*
fi

# 解压缩文件到 dist 文件夹
unzip -q "$zipFileFullPath" -d "$distFolder"

# 删除原始的 ZIP 文件
rm -f "$zipFileFullPath"

echo "ZIP文件成功解压缩到 $distFolder 文件夹。"
