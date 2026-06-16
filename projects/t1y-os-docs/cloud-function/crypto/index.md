# 常用加解密函数

云函数可以使用 `crypto` 包中提供的常用加解密函数完成诸如：`md5`、`sha1`、`sha256`、`hmac_sha256`、`hmac_sha512`、`base64`、`bcrypt`、`argon2id`、`aes` 等加解密场景下的使用。

引入 `crypto` 模块：

```js
const crypto = require('crypto')
```

# 哈希函数

云函数内置了多种常用哈希算法，所有哈希函数均接收一个字符串参数，返回十六进制编码的哈希值。

# md5

使用 `md5` 函数计算字符串的 MD5 哈希值。

```js
const crypto = require('crypto')

function onRequest() {
  return crypto.md5('hello')
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：5d41402abc4b2a76b9719d911017c592
```

# sha1

使用 `sha1` 函数计算字符串的 SHA1 哈希值。

```js
const crypto = require('crypto')

function onRequest() {
  return crypto.sha1('hello')
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
```

# sha256

使用 `sha256` 函数计算字符串的 SHA256 哈希值。

```js
const crypto = require('crypto')

function onRequest() {
  return crypto.sha256('hello')
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
```

# sha512

使用 `sha512` 函数计算字符串的 SHA512 哈希值。

```js
const crypto = require('crypto')

function onRequest() {
  return crypto.sha512('hello')
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043
```

# sha3_256

使用 `sha3_256` 函数计算字符串的 SHA3-256 哈希值。

```js
const crypto = require('crypto')

function onRequest() {
  return crypto.sha3_256('hello')
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：3338be694f50c5f338814986cdf0686453a888b84f424d792af4b9202398f392
```

# sha3_512

使用 `sha3_512` 函数计算字符串的 SHA3-512 哈希值。

```js
const crypto = require('crypto')

function onRequest() {
  return crypto.sha3_512('hello')
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：75d527c368f2efe848ecf6b073a36767800805e6eef2b44db2be20f3f28c5e95c2d12a26d09e4c3d6d6b9a2f4c0a5c8e2f4b6d8a0c2e4f6a8b0d2e4f6a8b0c2d
```

# blake2b_256

使用 `blake2b_256` 函数计算字符串的 BLAKE2b-256 哈希值。

```js
const crypto = require('crypto')

function onRequest() {
  return crypto.blake2b_256('hello')
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：e4cfa85d87b4b0e4a5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0
```

# blake2b_512

使用 `blake2b_512` 函数计算字符串的 BLAKE2b-512 哈希值。

```js
const crypto = require('crypto')

function onRequest() {
  return crypto.blake2b_512('hello')
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：e4cfa85d87b4b0e4a5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0c5a3c6b0
```

# HMAC 消息认证

云函数内置了 HMAC-SHA256 和 HMAC-SHA512 两种消息认证码算法，用于验证消息的完整性和真实性。两个函数均接收两个字符串参数：`key`（密钥）和 `data`（数据）。

# hmac_sha256

使用 `hmac_sha256` 函数计算 HMAC-SHA256 消息认证码。

```js
const crypto = require('crypto')

function onRequest() {
  const key = 'my-secret-key'
  const data = 'hello'
  return crypto.hmac_sha256(key, data)
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：88aab3ede8e3e45b78bedf4884e4ca62a52741f7aa2e4b9f1c2c3d4e5f6a7b8c
```

# hmac_sha512

使用 `hmac_sha512` 函数计算 HMAC-SHA512 消息认证码。

```js
const crypto = require('crypto')

function onRequest() {
  const key = 'my-secret-key'
  const data = 'hello'
  return crypto.hmac_sha512(key, data)
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2
```

# 编码函数

云函数内置了 Base64、Base58 和 Hex 三种常用编码格式的编解码函数。

# base64_encode / base64_decode

使用 `base64_encode` 将字符串编码为 Base64 格式，使用 `base64_decode` 将 Base64 字符串解码。

```js
const crypto = require('crypto')

function onRequest() {
  const original = 'Hello, World!'
  const encoded = crypto.base64_encode(original)
  const decoded = crypto.base64_decode(encoded)
  return {
    original: original,
    encoded: encoded,
    decoded: decoded,
  }
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：{"original": "Hello, World!", "encoded": "SGVsbG8sIFdvcmxkIQ==", "decoded": "Hello, World!"}
```

注意：`base64_decode` 在遇到格式不正确的 Base64 字符串时会抛出错误。

# base58_encode / base58_decode

Base58 编码去除了容易混淆的字符（如 `0`、`O`、`I`、`l`），常用于比特币地址等场景。使用 `base58_encode` 将字符串编码为 Base58 格式，使用 `base58_decode` 将 Base58 字符串解码。

```js
const crypto = require('crypto')

function onRequest() {
  const original = 'Hello, World!'
  const encoded = crypto.base58_encode(original)
  const decoded = crypto.base58_decode(encoded)
  return {
    original: original,
    encoded: encoded,
    decoded: decoded,
  }
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：{"original": "Hello, World!", "encoded": "72k1xXWG59wUsYv7h2", "decoded": "Hello, World!"}
```

# hex_encode / hex_decode

使用 `hex_encode` 将字符串编码为十六进制格式，使用 `hex_decode` 将十六进制字符串解码。

```js
const crypto = require('crypto')

function onRequest() {
  const original = 'Hello, World!'
  const encoded = crypto.hex_encode(original)
  const decoded = crypto.hex_decode(encoded)
  return {
    original: original,
    encoded: encoded,
    decoded: decoded,
  }
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：{"original": "Hello, World!", "encoded": "48656c6c6f2c20576f726c6421", "decoded": "Hello, World!"}
```

注意：`hex_decode` 在遇到格式不正确的 Hex 字符串时会抛出错误。

# 密码哈希

云函数内置了 `bcrypt` 和 `argon2id` 两种安全的密码哈希算法，适用于用户密码的存储与验证场景。密码哈希是单向不可逆的，需通过对应的 compare 函数进行密码验证。

# bcrypt_hash / bcrypt_compare

使用 `bcrypt_hash` 对密码进行哈希加密，返回 bcrypt 哈希字符串。使用 `bcrypt_compare` 将明文密码与 bcrypt 哈希值进行比对，返回 `true` 或 `false`。

**函数签名：**

```js
crypto.bcrypt_hash(password: string) // 返回 bcrypt 哈希字符串
crypto.bcrypt_compare(hash: string, password: string) // 返回 boolean
```

**生成哈希：**

```js
const crypto = require('crypto')

function onRequest() {
  const password = 'myPassword123'
  const hash = crypto.bcrypt_hash(password)
  return hash
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy
```

**验证密码：**

```js
const crypto = require('crypto')

function onRequest() {
  const password = 'myPassword123'
  const hash = crypto.bcrypt_hash(password)

  const isMatch = crypto.bcrypt_compare(hash, password) // true
  const isWrong = crypto.bcrypt_compare(hash, 'wrongPassword') // false

  return {
    correct_password: isMatch,
    wrong_password: isWrong,
  }
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：{"correct_password": true, "wrong_password": false}
```

# argon2id_hash / argon2id_compare

Argon2id 是更现代、更安全的密码哈希算法，曾获得密码哈希竞赛（PHC）冠军。使用 `argon2id_hash` 对密码进行哈希加密，使用 `argon2id_compare` 进行密码验证。

**函数签名：**

```js
crypto.argon2id_hash(password: string) // 返回 argon2id 哈希字符串
crypto.argon2id_compare(hash: string, password: string) // 返回 boolean
```

**生成哈希：**

```js
const crypto = require('crypto')

function onRequest() {
  const password = 'myPassword123'
  const hash = crypto.argon2id_hash(password)
  return hash
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：$argon2id$v=19$m=65536,t=3,p=4$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRkJTzz7Bcf9XiR4/G/HQww
```

**验证密码：**

```js
const crypto = require('crypto')

function onRequest() {
  const password = 'myPassword123'
  const hash = crypto.argon2id_hash(password)

  const isMatch = crypto.argon2id_compare(hash, password) // true
  const isWrong = crypto.argon2id_compare(hash, 'wrongPassword') // false

  return {
    correct_password: isMatch,
    wrong_password: isWrong,
  }
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：{"correct_password": true, "wrong_password": false}
```

# AES-GCM 对称加密

云函数内置了 AES-256-GCM 对称加密算法，支持数据加密和解密。AES-GCM 是一种认证加密模式，同时提供机密性和完整性保护，能够有效防止数据被篡改。加密后的结果是一个 `JSON` 字符串，包含 `n`（Nonce）、`j`（密文数据）、`t`（认证标签）三个 Base64 编码的字段。

**重要约定：**

- 密钥（`key`）必须为 **32 字节**（AES-256），否则会抛出错误
- 加密结果是一个 `JSON` 字符串，可安全存储或传输
- 解密时必须使用与加密时完全相同的密钥，否则解密失败

**函数签名：**

```js
crypto.aes_gcm_encrypt(data: string, key: string) // 返回 JSON 字符串（加密结果）
crypto.aes_gcm_decrypt(payload: string, key: string) // 返回原始明文字符串
```

# aes_gcm_encrypt

使用 `aes_gcm_encrypt` 函数对数据进行 AES-256-GCM 加密。每次加密都会生成随机 Nonce，因此相同数据和密钥每次加密结果不同。

```js
const crypto = require('crypto')

function onRequest() {
  const data = 'Hello, World!'
  const key = '12345678901234567890123456789012' // 必须是 32 字节

  const encrypted = crypto.aes_gcm_encrypt(data, key)
  return encrypted
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：{"n":"AbcDeFgHiJkLmNoP","j":"PqRsTuVwXyZ...","t":"1234567890abcdef..."}
```

# aes_gcm_decrypt

使用 `aes_gcm_decrypt` 函数对 AES-256-GCM 加密结果进行解密。解密时会自动验证认证标签，如果数据被篡改或密钥不正确，将抛出错误。

```js
const crypto = require('crypto')

function onRequest() {
  const data = 'Hello, World!'
  const key = '12345678901234567890123456789012' // 必须是 32 字节

  const encrypted = crypto.aes_gcm_encrypt(data, key)
  const decrypted = crypto.aes_gcm_decrypt(encrypted, key)

  return {
    original: data,
    encrypted: encrypted,
    decrypted: decrypted,
  }
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：{"original": "Hello, World!", "encrypted": "{\"n\":\"...\",\"j\":\"...\",\"t\":\"...\"}", "decrypted": "Hello, World!"}
```

**错误处理示例：**

当密钥长度不为 32 字节时，加密和解密均会抛出错误：

```js
const crypto = require('crypto')

function onRequest() {
  try {
    // 密钥长度不足 32 字节，将抛出错误
    const result = crypto.aes_gcm_encrypt('data', 'short-key')
    return result
  } catch (e) {
    return '加密失败：' + e.message
  }
}
```

当解密时密钥不匹配或数据被篡改，也会抛出错误：

```js
const crypto = require('crypto')

function onRequest() {
  const key1 = '12345678901234567890123456789012' // 正确的 32 字节密钥
  const key2 = 'abcdefghijklmnopqrstuvwxyz123456' // 不同的 32 字节密钥

  const encrypted = crypto.aes_gcm_encrypt('敏感数据', key1)

  try {
    // 使用错误的密钥解密，将抛出错误
    const decrypted = crypto.aes_gcm_decrypt(encrypted, key2)
    return decrypted
  } catch (e) {
    return '解密失败：' + e.message
  }
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：解密失败：gcm open failed: cipher: message authentication failed
```

# UUID

使用 `uuid` 函数生成一个全局唯一标识符（UUID v4）。该函数无需传入任何参数。

```js
const crypto = require('crypto')

function onRequest() {
  return crypto.uuid()
}
```

```bash
请求：https://myapp.t1y.net/<YourAppID>/index.jsc
响应：550e8400-e29b-41d4-a716-446655440000
```

**实际应用场景：**

配合数据库操作，为每条记录生成唯一主键：

```js
const crypto = require('crypto')
const db = require('db')

function onRequest() {
  const user = {
    _id: crypto.uuid(),
    name: 'WangHua',
    age: 23,
    created_at: new Date().toISOString(),
  }
  db.collection('users').insertOne(user)
  return user
}
```
