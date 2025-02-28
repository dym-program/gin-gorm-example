# 项目目录结构

本项目遵循 Go 官方推荐的目录结构，适用于大型应用的开发。包括后端服务和前端开发的文件结构。

## 目录说明

### `api/`
存放对外提供的 API 接口定义文件，可能包括 `swagger` 或 `openapi` 等文档。

### `assets/`
存放项目使用的静态资源，如图片、字体、样式文件等。

### `build/`
存放持续集成（CI）和构建相关文件，自动化部署和打包过程文件。

### `cmd/`
存放项目的可执行文件，每个子目录对应一个可执行文件。例如，主应用的入口文件应该位于 `cmd/_your_app_/` 中。

### `configs/`
存放项目的配置文件或配置模板，通常包含数据库配置、第三方 API 配置等。

### `deployments/`
存放 IaaS、PaaS 或容器编排的部署配置和模板。包括 Kubernetes 配置、Docker 配置等。

### `docs/`
项目的文档目录，存放设计文档、开发文档、用户文档等。

### `githooks/`
存放 Git 钩子脚本，用于自动化操作，例如在提交代码前运行代码检查。

### `init/`
存放系统和进程初始化配置文件，如数据库初始化脚本、启动脚本等。

### `internal/`
存放私有应用和库代码，不希望被外部导入的代码。一般包括：
- `/router` 路由配置文件。
- `/application` 存放命令和查询。
- `/command` 存放执行命令。
- `/query` 存放查询逻辑。
- `/middleware` 存放中间件。
- `/model` 存放模型定义。
- `/repository` 数据库访问层，封装数据操作。
- `/response` 统一响应格式。
- `/errmsg` 错误处理。

### `pkg/`
存放可以被外部应用使用的公共库代码。

### `test/`
存放测试代码，包含单元测试、集成测试等。需要测试应用的整体功能时，可以在此目录下添加测试。

### `third_party/`
存放第三方工具或库。

### `tools/`
存放项目支持工具，这些工具会导入 `/pkg` 和 `/internal` 中的代码。

### `vendor/`
存放项目依赖，通常由 `go mod vendor` 命令自动生成。

### `website/`
如果项目有网站数据或文档展示，可以将网站相关的文件放在此目录。

## 文件说明

- `README.md`: 项目的介绍、功能说明、安装和使用指引。
- `LICENSE.md`: 项目许可证，可以是私有的或开源的，如 MIT、GPL 等。
- `Makefile`: 用于管理项目，执行编译、测试、静态分析等任务。
- `CONTRIBUTING.md`: 贡献指南，说明如何参与项目的开发。
- ----------------------------



