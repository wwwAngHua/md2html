# 使用云函数构建 MCP Server

本文将介绍如何使用 t1yOS 云函数构建属于自己的 MCP Server，以下是一个最基本的示例：

```js
const ctx = require('context')
const mongo = require('mongo')

const mcp_name = 't1y_cloud_function_MCP'
const mcp_version = '1.0.0'
const jsonrpc_version = '2.0'

function onRequest() {
  if (ctx.method() !== 'POST') return

  const body = ctx.body()

  try {
    switch (body.method) {
      case 'initialize':
        return {
          jsonrpc: jsonrpc_version,
          id: body.id,
          result: {
            protocolVersion: '2025-11-25',
            capabilities: {
              tools: {
                listChanged: true,
              },
            },
            serverInfo: {
              name: mcp_name,
              version: mcp_version,
            },
          },
        }

      case 'notifications/initialized':
        // 可以做一些初始化后的动作（可选）
        console.log('MCP initialized')
        return

      case 'tools/list':
        return {
          jsonrpc: jsonrpc_version,
          id: body.id,
          result: {
            tools: [
              {
                name: 'get_user',
                description: '获取用户信息',
                inputSchema: {
                  type: 'object',
                  properties: {
                    id: { type: 'string' },
                  },
                  required: ['id'],
                },
              },
              // 更多 MCP Tools...
            ],
          },
        }

      case 'tools/call':
        const { name, arguments: args } = body.params || {}

        if (name === 'get_user') {
          const getUserResult = mongo
            .collection('users')
            .findOne({ _id: mongo.toObjectID(args.id) })
          return {
            jsonrpc: jsonrpc_version,
            id: body.id,
            result: {
              content: [
                {
                  type: 'text',
                  text: JSON.stringify(getUserResult),
                },
              ],
            },
          }
        }

        // 更多 MCP Tools 调用...

        return {
          jsonrpc: jsonrpc_version,
          id: body.id,
          error: {
            code: -32601,
            message: 'Tool not found',
          },
        }

      case 'ping':
        return {
          jsonrpc: jsonrpc_version,
          id: body.id,
          result: 'pong',
        }

      default:
        return {
          jsonrpc: jsonrpc_version,
          id: body.id,
          error: {
            code: -32601,
            message: 'Method not found',
          },
        }
    }
  } catch (err) {
    return {
      jsonrpc: jsonrpc_version,
      id: body.id,
      error: {
        code: -32603,
        message: err.message,
      },
    }
  }
}
```

以上只是一个最基础的 MCP Server 示例，包含一个 `get_user` MCP 工具，可以从云数据库中查询指定用户信息，如需更复杂的 MCP Server，可以将代码进行封装获得更好的体验。
