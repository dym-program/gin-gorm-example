#!/bin/bash

# 创建项目目录结构，如果目录已存在则不创建
mkdir -p \
    api \
    assets \
    build \
    cmd/_your_app_ \
    configs \
    deployments \
    docs \
    githooks \
    init \
    internal/router \
    internal/application \
    internal/command \
    internal/query \
    internal/middleware \
    internal/model \
    internal/repository \
    internal/response \
    internal/errmsg \
    pkg \
    test \
    third_party \
    tools \
    vendor \
    website

# 创建必要的文件
touch README.md LICENSE.md Makefile CONTRIBUTING.md

echo "Directory structure created successfully!"
