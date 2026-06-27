# t1yOS Skills 使用指南

t1yOS Skills 是一个专为 AI 编程助手（如 Claude Code、Cursor、Windsurf 等）设计的技能包，为 AI 提供 t1yOS 平台的完整知识，帮助开发者更高效地使用 AI 进行平台资源管理和云函数开发。

## 什么是 Skills？

Skills 是一种为 AI 助手注入专业领域知识的机制。通过加载 t1yOS Skills，AI 能够：

- 准确理解 t1yOS 平台的特性和约束
- 编写符合平台规范的云函数代码
- 正确使用数据库、加密、JWT 等内置模块
- 规避常见的平台开发陷阱

**没有 Skills 时**，AI 可能基于通用的 Node.js 或 Serverless 经验给出不符合 t1yOS 平台规则的建议（如使用 `module.exports`、`import` 语法、npm 包等）。

**使用 Skills 后**，AI 将严格遵循 t1yOS 的平台规范，生成的代码可直接部署运行。

## 获取

| 区域 | 仓库地址                                                           |
| ---- | ------------------------------------------------------------------ |
| 国际 | [github.com/t1yOS/t1y-skills](https://github.com/t1yOS/t1y-skills) |
| 国内 | [gitee.com/t1yOS/t1y-skills](https://gitee.com/t1yOS/t1y-skills)   |

## 安装方式

1. 从 [GitHub](https://github.com/t1yOS/t1y-skills)（国内用户从 [Gitee](https://gitee.com/t1yOS/t1y-skills)）下载 ZIP 包并解压
2. 将解压后的 `t1y-skills` 目录放置到对应位置：

| AI 工具               | 安装路径                                          |
| --------------------- | ------------------------------------------------- |
| **Claude Code**       | `~/.claude/skills/t1yOS/`                         |
| **Cursor**            | 项目根目录 `.claude/skills/` 或 `.cursor/skills/` |
| **Windsurf**          | 项目根目录 `.claude/skills/`                      |
| **VS Code (Copilot)** | 项目根目录 `.claude/skills/`                      |

放置后 AI 将自动识别并加载技能，无需额外配置。

## 触发方式

安装后，当你向 AI 提出与 t1yOS 平台相关的问题时，AI 将自动激活此技能。以下是一些典型的触发场景：

| 场景       | 示例提问                            |
| ---------- | ----------------------------------- |
| 编写云函数 | "帮我写一个用户登录的云函数"        |
| 数据库操作 | "如何在 t1yOS 中分页查询用户列表"   |
| 安全加密   | "给这个云函数加上签名验证"          |
| SDK 集成   | "帮我在 React 项目中集成 t1yOS SDK" |
| MCP 管理   | "帮我通过 MCP 创建一个云函数"       |
| 问题排查   | "为什么我的云函数部署后无法访问"    |

## 技能覆盖范围

t1yOS Skills 包含以下领域知识：

### 1. 云函数开发

- `.jsc` 文件编写规范
- `onRequest(ctx)` 入口函数
- 内置模块使用：`context`、`http`、`mongo`、`crypto`、`os`、`jwt`
- 本地/远程 require 模块加载
- Go 模板渲染语法
- 文件上传、Cookie 处理、重定向

### 2. 云数据库

- NoSQL 风格 CRUD 操作
- SQL-to-BSON 转换查询
- 聚合管道、分页、排序
- 查询操作符（`$gt`、`$in`、`$regex` 等）
- ObjectID 类型处理

### 3. 安全机制

- HMAC-SHA256 请求签名（`ctx.sign()`）
- AES-256-GCM 端到端加密（安全模式）
- Argon2id / bcrypt 密码哈希
- JWT 令牌签发与验证

### 4. MCP Server

- 配置方法（支持 Claude Code、Cursor、Windsurf、VS Code 等）
- 完整工具目录：数据库管理、环境变量、文件管理、云函数管理、应用设置、定时任务
- 权限模型说明

### 5. 客户端 SDK

- JavaScript/TypeScript（Web、Node.js、React Native、小程序）
- Android（Kotlin）
- Swift（iOS/macOS）
- Flutter（Dart）
- C#（.NET / Unity）
- Go
- 特殊类型标记（ObjectID、Date、Boolean 等）

## 最佳实践

### 与 MCP Server 配合使用

t1yOS 同时提供 **Skills** 和 **MCP Server** 两种 AI 辅助方式：

| 特性     | Skills                         | MCP Server                       |
| -------- | ------------------------------ | -------------------------------- |
| 用途     | 知识注入，指导 AI 编写正确代码 | 工具调用，让 AI 直接操作平台资源 |
| 工作方式 | AI 参考技能知识生成代码        | AI 通过 MCP 协议调用平台 API     |
| 典型场景 | 编写云函数、集成 SDK           | 管理数据库、部署函数、查看日志   |

**推荐做法**：同时安装 Skills 和配置 MCP Server。Skills 确保 AI 生成的代码正确，MCP Server 让 AI 能直接帮你管理平台资源。

### 开发工作流建议

1. **编写云函数时**：让 AI 参考 Skills 知识生成 `.jsc` 代码 → 通过 MCP `function_create` 或 WebIDE 部署
2. **调试问题时**：通过 MCP `function_get_log` 查看日志 → 让 AI 分析日志并修复代码
3. **管理数据时**：通过 MCP 数据库工具直接操作 → 如需批量处理，让 AI 编写云函数脚本

## 技能文件结构

```
t1y-os-skills/
├── SKILL.md                    # 主技能文件（核心知识和规范）
└── references/                 # 参考文档（按需加载）
    ├── cloud-function.md       # 云函数 API 完整参考
    ├── database.md             # 数据库 API 完整参考
    ├── crypto.md               # 加密模块 API 完整参考
    ├── mcp.md                  # MCP Server 工具参考
    ├── sdks.md                 # 多平台 SDK 集成指南
    ├── security.md             # 安全机制详解
    └── constants.md            # HTTP 常量和状态码
```

## 常见问题

### Q: Skills 和 MCP Server 有什么区别？

Skills 是"教 AI 知识"，让 AI 理解 t1yOS 的特性和规则；MCP Server 是"给 AI 工具"，让 AI 能直接操作平台。两者互补，建议同时使用。

### Q: 技能是否会消耗更多 Token？

技能加载会消耗一定的上下文 Token（约 3,000-5,000 词的基础知识 + 按需加载的参考文档），但能大幅减少 AI 因不了解平台而反复试错产生的 Token 消耗。基准测试显示，使用技能后代码通过率从 93% 提升到 100%，综合效率更优。

### Q: 如何更新技能？

当 t1yOS 平台发布新功能或 API 变更时，Skills 项目会同步更新。重新下载最新版本的 Skills 文件替换即可。

### Q: 支持哪些 AI 工具？

目前支持所有兼容 Skills 机制的 AI 编程工具，包括 Claude Code、Cursor、Windsurf、VS Code with GitHub Copilot 等。

## 获取方式

- **GitHub**：[github.com/t1yOS/t1y-skills](https://github.com/t1yOS/t1y-skills)
- **Gitee（国内镜像）**：[gitee.com/t1yOS/t1y-skills](https://gitee.com/t1yOS/t1y-skills)
- **反馈渠道**：如有问题或建议，请通过平台客服或 QQ 群反馈
