# t1yOS Swift SDK

## 系统要求

- iOS 15.0+ / macOS 12.0+ / watchOS 8.0+ / tvOS 15.0+
- Swift 5.7+
- Xcode 14.0+

## 安装

### Swift Package Manager（SPM）

在 `Package.swift` 中添加 T1YOS 依赖：

```swift
dependencies: [
    .package(url: "https://github.com/t1yOS/t1y-sdk-swift.git", from: "1.0.1")
]
```

或在 Xcode 中：**File → Add Packages...** 然后输入仓库地址。

## 快速开始

```swift
import T1YOS

// 1. 创建客户端
let client = try T1YOS(config: T1YOSConfig(
    appId: 1001,                                           // 必填：应用 ID（>= 1001）
    apiKey: "4fd7448cdc684431a62d8a0111dc6973",           // 必填：32 位 API Key
    secretKey: "17b784e359c946ffa65eebbf9ce29752",         // 必填：32 位 Secret Key
    // 以下参数可选（均有默认值）：
    // baseUrl: "https://myapp.t1y.net",
    // version: 0,
    // isSafeMode: false,
    // timeFormat: "YYYY-MM-DD HH:mm:ss",
    // offset: 0,
))

// 2. 初始化（与服务器同步时间偏移和安全模式）
await client.init()

// 3. 开始使用数据库！
try await client.db.collection("users").insertOne([
    "name": "张三",
    "age": 25,
    "active": true,
    "customTimeAt": T1YType.time.now(),
])
```

## 数据库操作

### 单条操作

```swift
let db = client.db.collection("users")

// 插入一条
try await db.insertOne(["name": "张三", "age": 25])
// 返回：InsertResult(objectId: "507f1f77bcf86cd799439011")

// 通过 ObjectID 查询
let result = try await db.findById("507f1f77bcf86cd799439011")
// result.result = ["_id": "507f1f77...", "name": "张三", ...]

// 通过 ObjectID 更新
try await db.updateById("507f1f77bcf86cd799439011", ["$set": ["age": 26]])

// 通过 ObjectID 删除
try await db.deleteById("507f1f77bcf86cd799439011")
```

### 条件操作

```swift
// 条件查询一条
let result = try await db.findOne(["name": "张三"])

// 条件更新一条
try await db.updateOne(
    ["name": "张三"],             // 查询条件
    body: ["$set": ["age": 27]]   // 更新内容
)

// 条件删除一条
try await db.deleteOne(["name": "张三"])
```

### 批量操作

```swift
// 插入多条
try await db.insertMany([
    ["name": "张三", "age": 25],
    ["name": "李四", "age": 30],
])
// 返回：InsertManyResult(objectIds: [...], insertedCount: 2)

// 删除多条
try await db.deleteMany(["age": ["$lt": 18]])

// 更新多条
try await db.updateMany(
    ["status": "inactive"],
    body: ["$set": ["status": "archived"]]
)
```

### 高级查询

```swift
// 分页查询
let result = try await db.find(
    page: 1,                      // 页码（从 1 开始）
    size: 20,                     // 每页条数（最大 100）
    sort: ["createdAt": -1],      // 排序（按创建时间倒序）
    filter: ["age": ["$gte": 18]] // 查询条件
)
// result.results     — 文档数组
// result.pagination  — Pagination(totalItems: 42, totalPages: 3)

// 聚合查询
let aggResult = try await db.aggregate([
    ["$match": ["status": "completed"]],
    ["$group": ["_id": "$category", "total": ["$sum": "$amount"]]],
    ["$sort": ["total": -1]],
])

// 计数
let count = try await db.count(["status": "active"])
// count = 42

// 去重查询
let values = try await db.distinct("city")
// 带条件过滤
let filteredValues = try await db.distinct("city", filter: ["country": "China"])
```

### 表管理

```swift
// 获取所有表
let collections = try await client.db.getCollections()
// collections = ["users", "orders", "products"]

// 创建表
try await client.db.collection("posts").create()

// 清空表
let deletedCount = try await client.db.collection("posts").clear()

// 删除表
try await client.db.collection("posts").drop()
```

## 特殊类型

SDK 提供了一系列辅助函数，用于生成服务端可识别的类型标记：

```swift
import T1YOS

try await db.insertOne([
    // ObjectID 引用
    "userId": try T1YType.objectID("507f1f77bcf86cd799439011"),

    // 日期类型
    "birthday": T1YType.date("2000-01-01T00:00:00Z"),
    "eventTime": T1YType.dateTime("2024-06-15T14:30:00Z"),
    "loginAt": T1YType.timestamp(1705312200),

    // 数值类型
    "active": T1YType.boolean(true),
    "quantity": T1YType.integer(42),
    "bigNumber": T1YType.bigint(9007199254740991),
    "rating": T1YType.float(4.5),
    "preciseValue": T1YType.double(3.141592653589793),

    // 结构化类型
    "tags": T1YType.array(["swift", "ios"]),
    "metadata": T1YType.map(["theme": "dark", "lang": "zh"]),
    "history": T1YType.mapArray([["action": "login"], ["action": "logout"]]),

    // 空值
    "deletedAt": T1YType.null,    // 服务端转为 nil
    "middleName": T1YType.none,   // 服务端转为 nil

    // 服务端时间辅助
    "customTimeAt": T1YType.time.now(),       // 服务端的 time.Now()
    "unixCreatedAt": T1YType.time.nowUnix(),  // 服务端的 time.Now().Unix()
])
```

## 元数据

```swift
// 获取全部元数据
let meta = try await client.getMeta()
// meta["results"] = ["version": 1, "collections": [...], ...]

// 获取指定字段
let versionMeta = try await client.getMeta("version")
// versionMeta["result"] = 1

// 检查更新
let hasUpdate = try await client.checkUpdate()
```

## 云函数

```swift
// 调用 .jsc 云函数
let result = try await client.callFunc("hello", params: ["name": "World"])

// 为此调用单独启用安全模式
let safeResult = try await client.callFunc("secureFunc", params: params, enableSafeMode: true)
```

## 安全机制

### 认证请求头

每个请求都会携带以下请求头：

- `X-T1Y-Application-ID` — 应用 ID
- `X-T1Y-API-Key` — 32 位 API Key
- `X-T1Y-Safe-Timestamp` — Unix 时间戳（UTC + 初始化时获取的时间偏移）
- `X-T1Y-Safe-Sign` — HMAC-SHA256 签名（64 位十六进制）

### 签名算法

```
message = METHOD + "\n" + URL_PATH + "\n" + SHA256(body) + "\n" + appId + "\n" + timestamp
signature = HMAC-SHA256(secretKey, message)
```

### 安全模式（AES-256-GCM）

当启用安全模式时（通过 `isSafeMode: true` 或初始化时自动检测），请求体将使用 AES-256-GCM 加密，密钥为应用的 SecretKey，服务端响应也会自动解密。

## API 参考

### T1YOS

| 方法                                         | 说明                                        |
| -------------------------------------------- | ------------------------------------------- |
| `T1YOS(config:)`                             | 创建客户端（校验 appId、apiKey、secretKey） |
| `init()`                                     | 与服务端同步时间偏移和安全模式              |
| `db.collection(name)`                        | 获取集合操作实例（链式调用）                |
| `db.toObjectID(id)`                          | 创建 ObjectID 标记字符串                    |
| `db.getCollections()`                        | 获取所有表                                  |
| `getMeta(_ field:)`                          | 获取应用元数据                              |
| `checkUpdate()`                              | 检查是否存在新版本                          |
| `callFunc(_ name:params:enableSafeMode:)`    | 调用云函数                                  |
| `request(method: path: params: encryption:)` | 原始认证请求                                |

### T1Collection

| 方法                              | HTTP   | 端点                                |
| --------------------------------- | ------ | ----------------------------------- |
| `insertOne(_ data:)`              | POST   | `/v5/classes/:name`                 |
| `deleteById(_ objectId:)`         | DELETE | `/v5/classes/:name/:objectId`       |
| `updateById(_: _:)`               | PUT    | `/v5/classes/:name/:objectId`       |
| `findById(_ objectId:)`           | GET    | `/v5/classes/:name/:objectId`       |
| `deleteOne(_ filter:)`            | DELETE | `/v5/classes/:name/one`             |
| `updateOne(_: body:)`             | PUT    | `/v5/classes/:name/one`             |
| `findOne(_ filter:)`              | POST   | `/v5/classes/:name/one`             |
| `insertMany(_ dataList:)`         | POST   | `/v5/classes/:name/many`            |
| `deleteMany(_ filter:)`           | DELETE | `/v5/classes/:name/many`            |
| `updateMany(_: body:)`            | PUT    | `/v5/classes/:name/many`            |
| `find(page: size: sort: filter:)` | POST   | `/v5/classes/:name/find`            |
| `aggregate(_ pipeline:)`          | POST   | `/v5/classes/:name/aggregate`       |
| `count(_ filter:)`                | POST   | `/v5/classes/:name/count`           |
| `distinct(_ fieldName:filter:)`   | POST   | `/v5/classes/:name/distinct/:field` |
| `create()`                        | POST   | `/v5/schemas/:name`                 |
| `clear()`                         | PUT    | `/v5/schemas/:name`                 |
| `drop()`                          | DELETE | `/v5/schemas/:name`                 |

### 数据库访问器

| 方法                  | HTTP | 端点          |
| --------------------- | ---- | ------------- |
| `db.getCollections()` | GET  | `/v5/schemas` |

### 特殊类型函数

所有类型标记通过 `T1YType` 命名空间访问：

| 函数                         | 说明                                |
| ---------------------------- | ----------------------------------- |
| `T1YType.objectID(_:)`       | 创建 ObjectID 标记（24 位十六进制） |
| `T1YType.date(_:)`           | 创建 Date 标记                      |
| `T1YType.dateTime(_:)`       | 创建 DateTime 标记                  |
| `T1YType.timestamp(_:)`      | 创建 Timestamp 标记（Unix 秒）      |
| `T1YType.boolean(_:)`        | 创建 Boolean 标记                   |
| `T1YType.integer(_:)`        | 创建 Integer 标记                   |
| `T1YType.bigint(_:)`         | 创建 Bigint 标记                    |
| `T1YType.float(_:)`          | 创建 Float 标记                     |
| `T1YType.double(_:)`         | 创建 Double 标记                    |
| `T1YType.array(_:)`          | 创建 Array 标记                     |
| `T1YType.map(_:)`            | 创建 Map 标记                       |
| `T1YType.mapArray(_:)`       | 创建 MapArray 标记                  |
| `T1YType.null` / `.none`     | 空值标记（服务端转为 nil）          |
| `T1YType.empty`              | 空值标记                            |
| `T1YType.undefined`          | 未定义值标记                        |
| `T1YType.time.now()`         | 服务端 time.Now() 标记              |
| `T1YType.time.nowUnix()`     | 服务端 time.Now().Unix() 标记       |
| `T1YType.time.nowUnixNano()` | 服务端 time.Now().UnixNano() 标记   |

### 加密工具函数

| 函数                                          | 说明                          |
| --------------------------------------------- | ----------------------------- |
| `hmacSHA256Hex(secret:message:)`              | 计算 HMAC-SHA256 十六进制摘要 |
| `verifyHmacSHA256(secret:message:signature:)` | 验证 HMAC-SHA256 签名         |
| `sha256Hex(_:)`                               | 计算 SHA-256 十六进制摘要     |
| `createSignature(input:)`                     | 创建请求 HMAC-SHA256 签名     |
| `encryptAESGCM(data:keyBytes:)`               | AES-256-GCM 加密              |
| `decryptAESGCM(jsonPayload:keyBytes:)`        | AES-256-GCM 解密              |

## 错误处理

SDK 使用结构化的错误类型：

```swift
do {
    try await client.db.collection("users").insertOne(["name": "张三"])
} catch let error as T1YError {
    print("API 错误：\(error.code) - \(error.message)")
    print(error.data ?? "")
} catch let error as ValidationError {
    print("参数校验错误：\(error.message)")
} catch {
    print("未知错误：\(error)")
}
```
