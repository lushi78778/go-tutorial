@echo off
setlocal enabledelayedexpansion

set "zipCount=0"

rem 统计当前目录下的 ZIP 文件数量
for %%i in (*.zip) do (
    set /a "zipCount+=1"
    set "zipFile=%%i"
)

rem 检查是否没有找到 ZIP 文件
if !zipCount! equ 0 (
    echo 未找到ZIP文件。退出。
    goto :eof
)

rem 检查是否有多个 ZIP 文件
if !zipCount! gtr 1 (
    echo 发现多个ZIP文件。退出。
    goto :eof
)

rem 如果只有一个 ZIP 文件，则继续解压缩
set "distFolder=dist"
set "zipFileFullPath=%cd%\%zipFile%"

rem 检查 dist 文件夹是否存在，如果存在则清空其内容，否则创建它
if not exist %distFolder% (
    mkdir %distFolder%
) else (
    del /q %distFolder%\*
)

rem 使用 PowerShell 的 Expand-Archive 命令将文件解压缩到 dist 文件夹
powershell Expand-Archive -Path "%zipFileFullPath%" -DestinationPath "%distFolder%"

rem 删除原始的 ZIP 文件
del /q "%zipFileFullPath%"

echo ZIP文件成功解压缩到 %distFolder% 文件夹。

:end
