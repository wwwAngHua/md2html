# t1yOS 云数据库

使用 t1yOS 云数据库之前，你需要先前往：[控制台（console.t1y.net）](https://console.t1y.net/)创建一个应用，随后你可以在`云操作系统页面`中使用"`云数据库 App`"管理云数据库。t1yOS 云数据库采用 `NoSQL` 数据库，你可以使用 `const db = require('mongo')` 在云函数中操作云数据库。

t1yOS 云数据库不仅支持原生的 `BSON` 风格操作（类似 `MongoDB` 原生驱动），还自研了一套 `SQL` 转 `BSON` 的系统，让你可以使用熟悉的 `SQL` 语句替代一些低频快速的使用场景，非常方便。

# 快速开始

在云函数中引入 `mongo` 模块，即可操作云数据库：

```js
const db = require('mongo')
```

# 获取集合

操作云数据库的第一步是获取要操作的集合（`Collection`）对象：

```js
const db = require('mongo')
const users = db.collection('users')
```

获取到集合对象后，即可对其进行增删改查等操作。

注意：所有集合操作方法都需要传入正确的参数类型，否则会抛出类型错误。

# 插入文档

文档插入成功后，系统会自动创建 `createdAt` 以及 `updatedAt` 字段。若往不存在的集合中插入文档，系统会自动创建该集合。

# insertOne

使用 `users.insertOne` 函数向集合中插入一条文档，返回插入文档的 `_id`（`ObjectID` 值）。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')
  const objectId = users.insertOne({
    name: '王华',
    age: 23,
    email: 'wwwanghua@outlook.com',
    tags: ['developer', 't1yOS'],
  })
  return { code: 200, message: '插入成功', data: { _id: objectId } }
}
```

```bash
响应 200 OK：
{
  "code": 200,
  "message": "插入成功",
  "data": { "_id": "60d5f7c8a1b2c3d4e5f6a7b8" }
}
```

# insertMany

使用 `users.insertMany` 函数向集合中批量插入多条文档，返回插入文档的 `_id` 数组。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')
  const objectIds = users.insertMany([
    { name: '张三', age: 25, role: 'admin' },
    { name: '李四', age: 30, role: 'editor' },
    { name: '王五', age: 22, role: 'viewer' },
  ])
  return { code: 200, message: '批量插入成功', data: objectIds }
}
```

```bash
响应 200 OK：
{
  "code": 200,
  "message": "批量插入成功",
  "data": [
    "60d5f7c8a1b2c3d4e5f6a7b8",
    "60d5f7c8a1b2c3d4e5f6a7b9",
    "60d5f7c8a1b2c3d4e5f6a7c0"
  ]
}
```

# insertBySQL

你也可以使用 `db.insertBySQL` 函数配合熟悉的 `INSERT` 语句直接插入文档。单条插入返回该文档的 `_id`，多条插入返回 `_id` 数组。

```js
const db = require('mongo')

function onRequest() {
  // 插入单条 — 操作 users 集合
  const id1 = db.insertBySQL(`INSERT INTO users (name, age, city) VALUES ('赵六', 27, '深圳')`)

  // 批量插入多条 — 操作 products 集合
  const ids = db.insertBySQL(
    `INSERT INTO products (name, price, stock) VALUES ('商品A', 99, 100), ('商品B', 199, 50)`
  )

  return { code: 200, message: '插入成功', data: { single: id1, batch: ids } }
}
```

# 查询文档

# findOne

使用 `users.findOne` 函数根据过滤条件查询集合中的一条文档。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')
  const user = users.findOne({ name: '王华' })
  return { code: 200, data: user }
}
```

```bash
响应 200 OK：
{
  "code": 200,
  "data": {
    "_id": "60d5f7c8a1b2c3d4e5f6a7b8",
    "name": "王华",
    "age": 23,
    "email": "wwwanghua@outlook.com",
    "tags": ["developer", "t1yOS"]
  }
}
```

# find

使用 `users.find` 函数分页查询集合中的文档，需要传入 `page`（页码）、`size`（每页数量）、`filter`（过滤条件）、`sort`（排序规则）四个参数。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')

  // 查询第 1 页，每页 10 条，按 age 降序排列
  const result = users.find(1, 10, {}, { age: -1 })

  return {
    code: 200,
    data: result.results, // 当前页数据
    pagination: result.pagination, // 分页信息：{ totalItems, totalPages }
  }
}
```

```bash
响应 200 OK：
{
  "code": 200,
  "data": [
    { "_id": "...", "name": "孙七", "age": 32, "city": "北京" },
    { "_id": "...", "name": "李四", "age": 30, "role": "editor" },
    { "_id": "...", "name": "王华", "age": 23, "email": "wwwanghua@outlook.com" }
  ],
  "pagination": {
    "page": 1,
    "size": 10,
    "totalItems": 25,
    "totalPages": 3
  }
}
```

**带过滤条件的分页查询：**

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')

  // 查询年龄大于 25，按创建时间升序排列
  const result = users.find(
    1, // page
    20, // size
    { age: { $gt: 25 } }, // filter
    { createdAt: 1 } // sort
  )

  return {
    code: 200,
    data: result.results,
    pagination: result.pagination,
  }
}
```

# selectBySQL

使用 `db.selectBySQL` 函数配合熟悉的 `SELECT` 语句直接查询文档。

```js
const db = require('mongo')

function onRequest() {
  // 基础查询 — 查询 users 集合
  const allUsers = db.selectBySQL('SELECT * FROM users')

  // 带条件的查询 — 查询 users 集合
  const adults = db.selectBySQL(
    'SELECT name, age, city FROM users WHERE age > 18 ORDER BY age DESC'
  )

  // 分页查询 — 查询 orders 集合
  const paged = db.selectBySQL('SELECT * FROM orders LIMIT 10 OFFSET 20')

  return { code: 200, data: { all: allUsers, adults: adults, paged: paged } }
}
```

注意：`SQL` 查询方式适合低频快速场景，复杂聚合场景请使用 `aggregate` 方法。

# 更新文档

# updateOne

使用 `users.updateOne` 函数更新集合中匹配过滤条件的第一条文档，返回受影响的文档数量。若 `_id` 不存在，系统不会执行任何操作。更新成功后，系统会自动更新 `updatedAt` 字段。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')

  // 将王华的年龄更新为 24
  const modifiedCount = users.updateOne(
    { name: '王华' }, // filter — 过滤条件
    { age: 24, city: '深圳' } // data — 要更新的字段
  )

  return { code: 200, message: `更新了 ${modifiedCount} 条文档` }
}
```

# updateMany

使用 `users.updateMany` 函数更新集合中所有匹配过滤条件的文档，返回受影响的文档数量。若 `_id` 不存在，系统不会执行任何操作。更新成功后，系统会自动更新 `updatedAt` 字段。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')

  // 将所有 viewer 角色的用户升级为 member
  const modifiedCount = users.updateMany({ role: 'viewer' }, { role: 'member' })

  return { code: 200, message: `批量更新了 ${modifiedCount} 条文档` }
}
```

# updateBySQL

使用 `db.updateBySQL` 函数配合 `UPDATE` 语句直接更新文档。更新成功后，系统会自动更新 `updatedAt` 字段。

```js
const db = require('mongo')

function onRequest() {
  const modifiedCount = db.updateBySQL(
    "UPDATE users SET status = 'active', verified = true WHERE age >= 18"
  )

  return { code: 200, message: `更新了 ${modifiedCount} 条文档` }
}
```

# 删除文档

# deleteOne

使用 `users.deleteOne` 函数删除集合中匹配过滤条件的第一条文档，返回删除的文档数量。若 `_id` 不存在，系统不会执行任何操作。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')

  const deletedCount = users.deleteOne({ name: '王五' })

  return { code: 200, message: `删除了 ${deletedCount} 条文档` }
}
```

# deleteMany

使用 `users.deleteMany` 函数删除集合中所有匹配过滤条件的文档，返回删除的文档数量。若 `_id` 不存在，系统不会执行任何操作。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')

  // 删除所有 role 为 viewer 的用户
  const deletedCount = users.deleteMany({ role: 'viewer' })

  return { code: 200, message: `批量删除了 ${deletedCount} 条文档` }
}
```

# deleteBySQL

使用 `db.deleteBySQL` 函数配合 `DELETE` 语句直接删除文档。

```js
const db = require('mongo')

function onRequest() {
  const deletedCount = db.deleteBySQL("DELETE FROM users WHERE status = 'inactive'")

  return { code: 200, message: `删除了 ${deletedCount} 条文档` }
}
```

# 聚合查询

使用 `aggregate` 方法可以执行复杂的聚合管道查询，支持 `MongoDB` 聚合管道的所有阶段操作。

```js
const db = require('mongo')

function onRequest() {
  const orders = db.collection('orders')

  // 按 city 分组统计用户数量
  const pipeline = [{ $group: { _id: '$city', count: { $sum: 1 } } }, { $sort: { count: -1 } }]

  const results = orders.aggregate(pipeline)

  return { code: 200, data: results }
}
```

```bash
响应 200 OK：
{
  "code": 200,
  "data": [
    { "_id": "深圳", "count": 156 },
    { "_id": "北京", "count": 132 },
    { "_id": "上海", "count": 98 }
  ]
}
```

**更复杂的聚合示例：**

```js
const db = require('mongo')

function onRequest() {
  const orders = db.collection('orders')

  const pipeline = [
    // 1. 只查询已完成的订单
    { $match: { status: 'completed' } },
    // 2. 按用户分组，统计总金额和订单数
    {
      $group: {
        _id: '$userId',
        totalAmount: { $sum: '$amount' },
        orderCount: { $sum: 1 },
        avgAmount: { $avg: '$amount' },
      },
    },
    // 3. 只返回总金额大于 1000 的用户
    { $match: { totalAmount: { $gt: 1000 } } },
    // 4. 按总金额降序排列
    { $sort: { totalAmount: -1 } },
    // 5. 只取前 10 名
    { $limit: 10 },
  ]

  const results = orders.aggregate(pipeline)

  return { code: 200, data: results }
}
```

注意：当聚合结果为空时，会抛出"没有任何匹配文档"的错误。你可以在调用前使用 `count` 方法预先判断文档是否存在。

# 计数与去重

# count

使用 `users.count` 函数根据过滤条件统计集合中的文档数量。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')

  // 统计所有用户
  const total = users.count({})

  // 统计深圳的用户数量
  const shenzhenUsers = users.count({ city: '深圳' })

  // 统计年龄大于 25 的用户
  const adults = users.count({ age: { $gt: 25 } })

  return {
    code: 200,
    data: { total: total, shenzhen: shenzhenUsers, adults: adults },
  }
}
```

# distinct

使用 `users.distinct` 函数查询某个字段的所有不重复值。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')

  // 查询所有不重复的城市
  const cities = users.distinct('city', {})

  // 查询 role 为 admin 的用户来自哪些城市
  const adminCities = users.distinct('city', { role: 'admin' })

  return { code: 200, data: { cities: cities, adminCities: adminCities } }
}
```

# 集合管理

# create

使用 `db.collection('xxx').create` 函数手动创建一个集合。

```js
const db = require('mongo')

function onRequest() {
  const logs = db.collection('logs')
  logs.create()

  return { code: 200, message: '集合创建成功' }
}
```

注意：通常情况下，在插入第一条文档时集合会自动创建，无需手动调用 `create` 方法。但在需要提前创建集合并设置索引等场景下，可以手动创建。

# getCollections

使用 `db.getCollections` 函数获取当前云数据库中所有的集合名称列表。

```js
const db = require('mongo')

function onRequest() {
  const collections = db.getCollections()

  return { code: 200, data: collections }
}
```

```bash
响应 200 OK：
{
  "code": 200,
  "data": ["users", "orders", "logs", "products"]
}
```

# clear

使用 `users.clear` 函数清空集合中的所有文档，保留集合索引。返回删除的文档数量。

```js
const db = require('mongo')

function onRequest() {
  const logs = db.collection('logs')
  const deletedCount = logs.clear()

  return { code: 200, message: `清空了 ${deletedCount} 条日志` }
}
```

# drop

使用 `users.drop` 函数删除整个集合，包括集合中的所有文档和索引。

```js
const db = require('mongo')

function onRequest() {
  const tempData = db.collection('temp_data')
  tempData.drop()

  return { code: 200, message: '集合已删除' }
}
```

注意：`drop` 操作不可逆，删除后集合中的所有文档将永久丢失，请谨慎操作。`clear` 与 `drop` 的区别在于：`clear` 仅清空文档保留集合，`drop` 则彻底删除整个集合。

# 工具方法

# toObjectID

使用 `db.toObjectID` 函数将字符串格式的 `ObjectID` 转换为标准的 `ObjectID` 类型，常用于基于 `_id` 的查询操作。

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')

  // 将字符串转换为 ObjectID
  const objectId = db.toObjectID('60d5f7c8a1b2c3d4e5f6a7b8')

  // 使用 ObjectID 查询文档
  const user = users.findOne({ _id: objectId })

  return { code: 200, data: user }
}
```

注意：`toObjectID` 要求传入有效的 24 位十六进制字符串，否则会抛出类型错误。

# 常见查询操作符

t1yOS 云数据库支持 `MongoDB` 标准的查询操作符，以下是一些常用示例：

```js
const db = require('mongo')

function onRequest() {
  const users = db.collection('users')

  // 等于
  const eq = users.find(1, 10, { status: 'active' }, {})

  // 不等于
  const ne = users.find(1, 10, { status: { $ne: 'banned' } }, {})

  // 大于 / 大于等于
  const gt = users.find(1, 10, { age: { $gt: 18 } }, {})
  const gte = users.find(1, 10, { age: { $gte: 18 } }, {})

  // 小于 / 小于等于
  const lt = users.find(1, 10, { age: { $lt: 60 } }, {})
  const lte = users.find(1, 10, { age: { $lte: 60 } }, {})

  // 包含于（in）
  const inQuery = users.find(1, 10, { role: { $in: ['admin', 'editor'] } }, {})

  // 不包含于（nin）
  const nin = users.find(1, 10, { role: { $nin: ['banned'] } }, {})

  // 逻辑与（and）
  const andQuery = users.find(
    1,
    10,
    {
      $and: [{ age: { $gte: 18 } }, { age: { $lte: 60 } }],
    },
    {}
  )

  // 逻辑或（or）
  const orQuery = users.find(
    1,
    10,
    {
      $or: [{ role: 'admin' }, { role: 'editor' }],
    },
    {}
  )

  // 字段存在性
  const exists = users.find(1, 10, { email: { $exists: true } }, {})

  // 正则匹配
  const regex = users.find(1, 10, { name: { $regex: '^王' } }, {})

  return { code: 200, message: '查询完成' }
}
```

# 完整示例

以下是一个完整的云函数示例，展示了如何在云函数中使用云数据库完成一个用户管理接口：

```js
const db = require('mongo')
const ctx = require('context')
const http = require('http')

function onRequest() {
  const users = db.collection('users')

  switch (ctx.method()) {
    case 'GET':
      return handleQuery(users)
    case 'POST':
      return handleCreate(users)
    case 'PUT':
      return handleUpdate(users)
    case 'DELETE':
      return handleDelete(users)
    default:
      return { code: http.StatusMethodNotAllowed, message: '不支持的请求方法' }
  }
}

function handleQuery(users) {
  const page = parseInt(ctx.query('page', '1'))
  const size = parseInt(ctx.query('size', '10'))
  const keyword = ctx.query('keyword', '')

  const filter = keyword ? { name: { $regex: keyword } } : {}
  const result = users.find(page, size, filter, { createdAt: -1 })

  return { code: http.StatusOK, data: result.results, pagination: result.pagination }
}

function handleCreate(users) {
  const body = ctx.body()
  if (!body.name) {
    return { code: http.StatusBadRequest, message: 'name 字段不能为空' }
  }

  const objectId = users.insertOne({
    name: body.name,
    age: body.age || 0,
    email: body.email || '',
    role: body.role || 'user',
  })

  return { code: http.StatusCreated, message: '创建成功', data: { _id: objectId } }
}

function handleUpdate(users) {
  const body = ctx.body()
  const userId = ctx.query('id')

  if (!userId) {
    return { code: http.StatusBadRequest, message: 'id 参数不能为空' }
  }

  const objectId = db.toObjectID(userId)
  const modifiedCount = users.updateOne(
    { _id: objectId },
    {
      name: body.name,
      age: body.age,
      email: body.email,
      role: body.role,
    }
  )

  return { code: http.StatusOK, message: `更新了 ${modifiedCount} 条文档` }
}

function handleDelete(users) {
  const userId = ctx.query('id')

  if (!userId) {
    return { code: http.StatusBadRequest, message: 'id 参数不能为空' }
  }

  const objectId = db.toObjectID(userId)
  const deletedCount = users.deleteOne({ _id: objectId })

  return { code: http.StatusOK, message: `删除了 ${deletedCount} 条文档` }
}
```

```bash
# 查询用户列表（GET）
请求：GET https://myapp.t1y.net/<YourAppID>/users.jsc?page=1&size=10&keyword=王
响应 200 OK：
{
  "code": 200,
  "data": [
    { "_id": "...", "name": "王华", "age": 23, "role": "admin" }
  ],
  "pagination": { "page": 1, "size": 10, "totalItems": 1, "totalPages": 1 }
}

# 创建用户（POST）
请求：POST https://myapp.t1y.net/<YourAppID>/users.jsc
Body：{ "name": "新用户", "age": 25, "email": "new@example.com" }
响应 201 Created：
{ "code": 201, "message": "创建成功", "data": { "_id": "60d5f7c8a1b2c3d4e5f6a7b8" } }

# 更新用户（PUT）
请求：PUT https://myapp.t1y.net/<YourAppID>/users.jsc?id=60d5f7c8a1b2c3d4e5f6a7b8
Body：{ "name": "更新后的名称", "age": 30 }
响应 200 OK：
{ "code": 200, "message": "更新了 1 条文档" }

# 删除用户（DELETE）
请求：DELETE https://myapp.t1y.net/<YourAppID>/users.jsc?id=60d5f7c8a1b2c3d4e5f6a7b8
响应 200 OK：
{ "code": 200, "message": "删除了 1 条文档" }
```

# API 参考

# mongo 全局方法

| 方法                  | 参数                                | 返回值                  | 说明                    |
| --------------------- | ----------------------------------- | ----------------------- | ----------------------- |
| `db.toObjectID(str)`  | `str: String` — 24 位十六进制字符串 | `ObjectID`              | 将字符串转换为 ObjectID |
| `db.collection(name)` | `name: String` — 集合名称           | `Collection`            | 获取集合操作对象        |
| `db.getCollections()` | 无                                  | `String[]`              | 获取所有集合名称列表    |
| `db.insertBySQL(sql)` | `sql: String` — INSERT 语句         | `ObjectID / ObjectID[]` | 使用 SQL INSERT 插入    |
| `db.deleteBySQL(sql)` | `sql: String` — DELETE 语句         | `Number`                | 使用 SQL DELETE 删除    |
| `db.updateBySQL(sql)` | `sql: String` — UPDATE 语句         | `Number`                | 使用 SQL UPDATE 更新    |
| `db.selectBySQL(sql)` | `sql: String` — SELECT 语句         | `Object[]`              | 使用 SQL SELECT 查询    |

# Collection 方法

| 方法                              | 参数                                                       | 返回值                                | 说明                 |
| --------------------------------- | ---------------------------------------------------------- | ------------------------------------- | -------------------- |
| `.insertOne(data)`                | `data: Object`                                             | `ObjectID`                            | 插入单条文档         |
| `.insertMany(data)`               | `data: Object[]`                                           | `ObjectID[]`                          | 批量插入文档         |
| `.deleteOne(filter)`              | `filter: Object`                                           | `Number`                              | 删除单条文档         |
| `.deleteMany(filter)`             | `filter: Object`                                           | `Number`                              | 批量删除文档         |
| `.updateOne(filter, data)`        | `filter: Object, data: Object`                             | `Number`                              | 更新单条文档         |
| `.updateMany(filter, data)`       | `filter: Object, data: Object`                             | `Number`                              | 批量更新文档         |
| `.findOne(filter)`                | `filter: Object`                                           | `Object / null`                       | 查询单条文档         |
| `.find(page, size, filter, sort)` | `page: Number, size: Number, filter: Object, sort: Object` | `{ results, page, size, pagination }` | 分页查询             |
| `.aggregate(pipeline)`            | `pipeline: Object[]`                                       | `Object[]`                            | 聚合管道查询         |
| `.count(filter)`                  | `filter: Object`                                           | `Number`                              | 统计文档数量         |
| `.distinct(field, filter)`        | `field: String, filter: Object`                            | `Array`                               | 字段去重查询         |
| `.create()`                       | 无                                                         | `true`                                | 创建集合             |
| `.clear()`                        | 无                                                         | `Number`                              | 清空集合（保留结构） |
| `.drop()`                         | 无                                                         | `true`                                | 删除集合（彻底删除） |

# 结合云函数使用

云数据库最常见的场景是在云函数中使用。云函数收到客户端请求后，通过 `mongo` 模块操作云数据库，然后将结果返回给客户端。

```js
const db = require('mongo')
const ctx = require('context')

function onRequest() {
  const products = db.collection('products')
  const page = parseInt(ctx.query('page', '1'))
  const size = parseInt(ctx.query('size', '20'))
  const category = ctx.query('category', '')

  const filter = category ? { category: category } : {}
  const result = products.find(page, size, filter, { sales: -1 })

  return {
    code: 200,
    data: result.results,
    pagination: result.pagination,
  }
}
```

除此之外，t1yOS 还支持将云数据库自动映射为 `RESTful API`，客户端可直接通过 `HTTP` 请求对云数据库进行增删改查操作，无需编写云函数。详情请参阅：[RESTful API 文档](../restful-api/index.html)。

# 注意事项

1. 集合名称不能为空字符串，必须指定有效的集合名称
2. 所有方法的参数类型必须严格匹配，例如 `filter` 必须为 `Object` 类型，否则会抛出类型错误
3. `aggregate` 方法要求传入的管道必须是数组，且每个阶段必须是有效的对象
4. `find` 方法的 `page` 和 `size` 会自动转换为整数，`page` 从 1 开始
5. `drop` 操作不可逆，执行前请确认
6. `toObjectID` 要求传入的字符串必须是 24 位有效的十六进制 ObjectID 格式
7. `SQL` 风格操作（`insertBySQL`、`deleteBySQL`、`updateBySQL`、`selectBySQL`）是 `mongo` 模块的顶层方法，直接在 `db` 上调用，真实操作的目标集合由 SQL 语句中指定的表名决定
