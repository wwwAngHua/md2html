# MCP Server

使用 `t1yOS MCP Server` 可以将你在 `t1yOS` 平台上创建的应用功能全部暴露给 `Claude Code`、`Codex`、`Cursor`、`Windsurf`、`VS Code（GitHub Copilot）`、`Zed` 等主流 `AI` 编程工具。如需查看完整 `MCP Tools` 列表请阅读：[MCP Tools 列表文档](./tools/index.html)

## 准备工作

开始前，你需要前往[控制台（console.t1y.net）](https://console.t1y.net/)中的`个人信息`页面中获取你账号的 `Access Token` 并创建一个应用，随后前往应用`云操作系统页面`中的`设置 App` 中获取 `AppID`、`API Key`、`Secret Key` 等密钥信息。

## MCP Server 地址

```bash
https://ai.t1y.net/mcp
```

想要调用 `t1yOS MCP Server` 非常简单，只需要使用 `HTTP` 类型接入即可，地址填写 `t1yOS MCP Server` 的地址，请求头中附带你在准备工作中准备的 4 个密钥即可，任何 `AI` 工具都是通用的思想。

**请求头例如：**

```json
{
  "X-T1Y-Application-ID": "${T1Y_APP_ID}",
  "X-T1Y-API-Key": "${T1Y_API_KEY}",
  "X-T1Y-Secret-Key": "${T1Y_SECRET_KEY}",
  "X-T1Y-Access-Token": "${T1Y_ACCESS_TOKEN}"
}
```

建议使用环境变量存储，切勿硬编码或提交到代码仓库。

## Claude Code（命令行）

**配置文件位置**：项目根目录 `.mcp.json`，或全局 `~/.claude/.mcp.json`

**方式一：编辑配置文件（推荐）**

```json
{
  "mcpServers": {
    "t1yOS": {
      "type": "http",
      "url": "https://ai.t1y.net/mcp",
      "headers": {
        "X-T1Y-Application-ID": "${T1Y_APP_ID}",
        "X-T1Y-API-Key": "${T1Y_API_KEY}",
        "X-T1Y-Secret-Key": "${T1Y_SECRET_KEY}",
        "X-T1Y-Access-Token": "${T1Y_ACCESS_TOKEN}"
      }
    }
  }
}
```

**方式二：命令行添加**

```bash
claude mcp add-json t1yOS '{
  "type": "http",
  "url": "https://ai.t1y.net/mcp",
  "headers": {
    "X-T1Y-Application-ID": "${T1Y_APP_ID}",
    "X-T1Y-API-Key": "${T1Y_API_KEY}",
    "X-T1Y-Secret-Key": "${T1Y_SECRET_KEY}",
    "X-T1Y-Access-Token": "${T1Y_ACCESS_TOKEN}"
  }
}'
```

**设置环境变量**（在 Shell 配置文件 `~/.zshrc` 或 `~/.bashrc` 中添加）：

```bash
export T1Y_APP_ID="your_app_id"
export T1Y_API_KEY="your_api_key"
export T1Y_SECRET_KEY="your_secret_key"
export T1Y_ACCESS_TOKEN="your_access_token"
```

**验证连接**：

```bash
claude mcp list        # 查看已添加的 MCP Server
/mcp                   # 在 Claude Code 会话中查看工具状态
```

## Cursor

**配置文件位置**：

- 项目级（推荐）：项目根目录 `.cursor/mcp.json`
- 全局：`~/.cursor/mcp.json`

```json
{
  "mcpServers": {
    "t1yOS": {
      "type": "http",
      "url": "https://ai.t1y.net/mcp",
      "headers": {
        "X-T1Y-Application-ID": "${env:T1Y_APP_ID}",
        "X-T1Y-API-Key": "${env:T1Y_API_KEY}",
        "X-T1Y-Secret-Key": "${env:T1Y_SECRET_KEY}",
        "X-T1Y-Access-Token": "${env:T1Y_ACCESS_TOKEN}"
      }
    }
  }
}
```

**也可通过 UI 添加**：

1. 打开 Cursor Settings（`Cmd/Ctrl + ,`）
2. 进入 **Features → MCP**（或 **Tools & MCP**）
3. 点击 **+ Add New MCP Server**
4. 填写 Server URL 及 Headers

> Cursor 使用 `${env:VAR}` 语法引用环境变量（注意与其他工具的 `${VAR}` 语法区别）。

## Windsurf（Codeium）

**配置文件位置**：`~/.codeium/windsurf/mcp_config.json`

> **注意**：Windsurf 使用 `serverUrl`（而非 `url`）字段，并需显式声明 `type: "streamable-http"`。

```json
{
  "mcpServers": {
    "t1yOS": {
      "type": "streamable-http",
      "serverUrl": "https://ai.t1y.net/mcp",
      "headers": {
        "X-T1Y-Application-ID": "${T1Y_APP_ID}",
        "X-T1Y-API-Key": "${T1Y_API_KEY}",
        "X-T1Y-Secret-Key": "${T1Y_SECRET_KEY}",
        "X-T1Y-Access-Token": "${T1Y_ACCESS_TOKEN}"
      }
    }
  }
}
```

**也可通过 UI 配置**：Settings → AI → External Tools

## VS Code（GitHub Copilot）

**配置文件位置**：

- 用户级：通过命令面板 `MCP: Open User Configuration` 打开
- 项目级：`.vscode/mcp.json`

> **注意**：VS Code 使用 `"servers"` 作为根键（其他工具使用 `"mcpServers"`），且需在 settings 中启用 MCP。

```json
{
  "servers": {
    "t1yOS": {
      "type": "http",
      "url": "https://ai.t1y.net/mcp",
      "headers": {
        "X-T1Y-Application-ID": "${input:T1Y_APP_ID}",
        "X-T1Y-API-Key": "${input:T1Y_API_KEY}",
        "X-T1Y-Secret-Key": "${input:T1Y_SECRET_KEY}",
        "X-T1Y-Access-Token": "${input:T1Y_ACCESS_TOKEN}"
      }
    }
  }
}
```

**启用 MCP 支持**（在 `settings.json` 中添加）：

```json
{
  "chat.mcp.enabled": true
}
```

## Gemini CLI

**配置文件位置**：`~/.gemini/settings.json`

```json
{
  "mcpServers": {
    "t1yOS": {
      "type": "http",
      "url": "https://ai.t1y.net/mcp",
      "headers": {
        "X-T1Y-Application-ID": "${T1Y_APP_ID}",
        "X-T1Y-API-Key": "${T1Y_API_KEY}",
        "X-T1Y-Secret-Key": "${T1Y_SECRET_KEY}",
        "X-T1Y-Access-Token": "${T1Y_ACCESS_TOKEN}"
      }
    }
  }
}
```

## Continue.dev

**配置文件位置**：`~/.continue/config.json`

```json
{
  "mcpServers": [
    {
      "name": "t1yOS",
      "transport": {
        "type": "streamableHttp",
        "url": "https://ai.t1y.net/mcp",
        "requestOptions": {
          "headers": {
            "X-T1Y-Application-ID": "${T1Y_APP_ID}",
            "X-T1Y-API-Key": "${T1Y_API_KEY}",
            "X-T1Y-Secret-Key": "${T1Y_SECRET_KEY}",
            "X-T1Y-Access-Token": "${T1Y_ACCESS_TOKEN}"
          }
        }
      }
    }
  ]
}
```

## Zed

**配置文件位置**：`~/.config/zed/settings.json`

> **注意**：Zed 使用 `context_servers` 作为根键。当前版本对自定义请求头支持有限，推荐使用 `mcp-remote` 代理方式接入。

```json
{
  "context_servers": {
    "t1yOS": {
      "command": {
        "path": "npx",
        "args": [
          "-y",
          "mcp-remote@latest",
          "https://ai.t1y.net/mcp",
          "--header",
          "X-T1Y-Application-ID:${T1Y_APP_ID}",
          "--header",
          "X-T1Y-API-Key:${T1Y_API_KEY}",
          "--header",
          "X-T1Y-Secret-Key:${T1Y_SECRET_KEY}",
          "--header",
          "X-T1Y-Access-Token:${T1Y_ACCESS_TOKEN}"
        ],
        "env": {
          "T1Y_APP_ID": "your_app_id",
          "T1Y_API_KEY": "your_api_key",
          "T1Y_SECRET_KEY": "your_secret_key",
          "T1Y_ACCESS_TOKEN": "your_access_token"
        }
      }
    }
  }
}
```

## 常见问题

**Q：连接失败，提示 401/403？**  
A：请检查四个认证 Header 是否全部传入，确认凭证未过期。

**Q：工具列表为空？**  
A：重启客户端后再试。部分工具（如 Cursor）在添加 Server 后需重载窗口（`Cmd/Ctrl + Shift + P → Reload Window`）。

**Q：配置文件修改后无效？**  
A：确认 JSON 格式正确（缺少逗号或括号会导致静默失败）。可用 `cat config.json | python3 -m json.tool` 验证语法。

**Q：Windsurf 连接报错？**  
A：确认使用了 `serverUrl` 字段（不是 `url`），并且添加了 `"type": "streamable-http"`。

**Q：VS Code 找不到 MCP 配置选项？**  
A：请确认 VS Code 版本 >= 1.102，并在 `settings.json` 中启用 `"chat.mcp.enabled": true`。
