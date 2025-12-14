# gin-pathway

## 简介
gin-pathway 是一个基于 Gin 框架的目录结构示例项目，旨在为开发者提供一个行业最佳实践的目录结构模板。通过使用 gin-pathway，开发者可以快速上手 Gin 框架，无需在项目初期花费大量时间设计和调整目录结构。

## 主要功能
- 提供一个清晰、合理的目录结构示例
- 遵循行业最佳实践，便于代码管理和维护
- 内置已集成日志、配置、数据库及示例 demo，简化开发流程
- 适用于各种规模的 Gin 项目，帮助开发者快速搭建项目基础

## 项目特点
- 清晰的分层架构：将业务逻辑、数据访问、HTTP 处理等模块分离，便于维护和扩展
- 最佳实践遵循：参考了多个成功的 Gin 项目，确保目录结构符合社区标准
- 易于上手：提供了详细的文档和示例代码，帮助开发者快速入门
- 灵活性：可以根据实际需求调整目录结构，满足不同项目的个性化要求

## 目录结构示例
```html
gin-pathway/
├── cmd/                     # 应用启动入口
│   └── main.go              # 主程序入口
├── configs/                 # 配置文件存放目录
│   └── config.yaml          # 示例配置文件
├── docs/                    # 文档文件，如 Swagger api 等
├── internal/                # 内部包，存放核心业务逻辑
│   ├── api/                 # 版本路由
│   ├── app/                 # 包括应用程序启动、初始化等逻辑
│   ├── controller/          # HTTP 请求处理函数
│   ├── middleware/          # 中间件
│   ├── model/               # 数据模型定义
│   ├── repository/          # 数据访问层
│   ├── service/             # 业务逻辑层
│   └── utils/               # 工具函数
├── pkg/                     # 第三方依赖或公共工具包
├── scripts/                 # 脚本文件，如项目部署脚本等
├── tests/                   # 单元测试文件
├── .env                     # 环境变量文件
├── go.mod                   # Go 模块管理文件
├── go.sum                   # Go 模块依赖校验文件
└── README.md                # 项目说明文档
```
针对项目目录的详细介绍，可参考如下链接：
[Go 项目实战：搭建高效的 Gin Web 目录结构](https://vespeng.com/posts/go_practical_gin_directory_structure/)

## 使用方法
1. 克隆项目到本地
    ```shell
   git clone git@github.com:vespeng/gin-pathway.git
   cd gin-pathway
   ```
2. 安装依赖
    ```shell
    go mod tidy
    ```
3. 配置环境
    - 根据 `.env` 文件中的配置，修改 configs/config.yaml 文件以适应您的开发环境。
4. 开始开发你的 Gin 应用程序
5. 启动应用
    ```shell
    go run cmd/main.go
    ```
---
