@echo off

:: 创建项目目录结构，如果目录已存在则不创建
mkdir api
mkdir assets
mkdir build
mkdir cmd\_your_app_
mkdir configs
mkdir deployments
mkdir docs
mkdir githooks
mkdir init
mkdir internal\router
mkdir internal\application
mkdir internal\command
mkdir internal\query
mkdir internal\middleware
mkdir internal\model
mkdir internal\repository
mkdir internal\response
mkdir internal\errmsg
mkdir pkg
mkdir test
mkdir third_party
mkdir tools
mkdir vendor
mkdir website

:: 创建必要的文件
echo. > README.md
echo. > LICENSE.md
echo. > Makefile
echo. > CONTRIBUTING.md

echo "Directory structure created successfully!"
