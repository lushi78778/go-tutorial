@echo off
set ZIP_FILE=webHelpGO-TUTORIAL2-all.zip
set DEST_FOLDER=dist

rem 清空dist文件夹
if exist "%DEST_FOLDER%" (
    echo 清空 %DEST_FOLDER% 文件夹
    rmdir /s /q "%DEST_FOLDER%"
)

rem 创建dist文件夹
mkdir "%DEST_FOLDER%"

rem 解压ZIP文件到dist文件夹
powershell -Command "Expand-Archive -Path '.\%ZIP_FILE%' -DestinationPath '.\%DEST_FOLDER%'"
