# 常量

云函数常量包括 HTTP 方法、常用 MIME 类型、HTTP 状态码、HTTP 头，很大一部分来自 `Go` 语言 `net/http` 标准库中，在云函数中，你也只需要导入 `http` 包即可使用，例如：

```js
const ctx = require('context')
const http = require('http')

function onRequest() {
  const method = ctx.method()
  if (method == http.MethodGet) {
    const ip = ctx.get(http.HeaderXForwardedFor)
    ctx.sendStatus(http.StatusOK) // HTTP 200 OK
    return `你的 IP 地址是：${ip}`
  }
  return null
}
```

除此之外，`http` 包中还有 `http.statusText` 函数，可以将 `HTTP` 状态码（如 `200`、`404`）转换为标准 `HTTP` 状态描述文本（如 "`OK`"、"`Not Found`"），例如：

```js
const http = require('http')

function onRequest() {
  return http.statusText(500) // Internal Server Error
}
```

**newError：**

```js
const http = require('http')

function onRequest() {
  return http.newError(782, 'Custom error message')
}
```

## HTTP 方法

```js
export const MethodGet = 'GET' // RFC 7231, 4.3.1
export const MethodHead = 'HEAD' // RFC 7231, 4.3.2
export const MethodPost = 'POST' // RFC 7231, 4.3.3
export const MethodPut = 'PUT' // RFC 7231, 4.3.4
export const MethodPatch = 'PATCH' // RFC 5789
export const MethodDelete = 'DELETE' // RFC 7231, 4.3.5
export const MethodConnect = 'CONNECT' // RFC 7231, 4.3.6
export const MethodOptions = 'OPTIONS' // RFC 7231, 4.3.7
export const MethodTrace = 'TRACE' // RFC 7231, 4.3.8
export const MethodUse = 'USE'
```

## 常用 MIME 类型

```js
export const MIMETextXML = 'text/xml'
export const MIMETextHTML = 'text/html'
export const MIMETextPlain = 'text/plain'
export const MIMETextJavaScript = 'text/javascript'
export const MIMETextCSS = 'text/css'
export const MIMEApplicationXML = 'application/xml'
export const MIMEApplicationJSON = 'application/json'
export const MIMEApplicationCBOR = 'application/cbor'
export const MIMEApplicationForm = 'application/x-www-form-urlencoded'
export const MIMEOctetStream = 'application/octet-stream'
export const MIMEMultipartForm = 'multipart/form-data'

export const MIMETextXMLCharsetUTF8 = 'text/xml; charset=utf-8'
export const MIMETextHTMLCharsetUTF8 = 'text/html; charset=utf-8'
export const MIMETextPlainCharsetUTF8 = 'text/plain; charset=utf-8'
export const MIMETextJavaScriptCharsetUTF8 = 'text/javascript; charset=utf-8'
export const MIMETextCSSCharsetUTF8 = 'text/css; charset=utf-8'
export const MIMEApplicationXMLCharsetUTF8 = 'application/xml; charset=utf-8'
export const MIMEApplicationJSONCharsetUTF8 = 'application/json; charset=utf-8'
```

## HTTP 状态码（`镜像 Go net/http 包`）

```js
export const StatusContinue = 100 // RFC 7231, 6.2.1
export const StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
export const StatusProcessing = 102 // RFC 2518, 10.1
export const StatusEarlyHints = 103 // RFC 8297
export const StatusOK = 200 // RFC 7231, 6.3.1
export const StatusCreated = 201 // RFC 7231, 6.3.2
export const StatusAccepted = 202 // RFC 7231, 6.3.3
export const StatusNonAuthoritativeInformation = 203 // RFC 7231, 6.3.4
export const StatusNoContent = 204 // RFC 7231, 6.3.5
export const StatusResetContent = 205 // RFC 7231, 6.3.6
export const StatusPartialContent = 206 // RFC 7233, 4.1
export const StatusMultiStatus = 207 // RFC 4918, 11.1
export const StatusAlreadyReported = 208 // RFC 5842, 7.1
export const StatusIMUsed = 226 // RFC 3229, 10.4.1
export const StatusMultipleChoices = 300 // RFC 7231, 6.4.1
export const StatusMovedPermanently = 301 // RFC 7231, 6.4.2
export const StatusFound = 302 // RFC 7231, 6.4.3
export const StatusSeeOther = 303 // RFC 7231, 6.4.4
export const StatusNotModified = 304 // RFC 7232, 4.1
export const StatusUseProxy = 305 // RFC 7231, 6.4.5
export const StatusSwitchProxy = 306 // RFC 9110, 15.4.7 (Unused)
export const StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
export const StatusPermanentRedirect = 308 // RFC 7538, 3
export const StatusBadRequest = 400 // RFC 7231, 6.5.1
export const StatusUnauthorized = 401 // RFC 7235, 3.1
export const StatusPaymentRequired = 402 // RFC 7231, 6.5.2
export const StatusForbidden = 403 // RFC 7231, 6.5.3
export const StatusNotFound = 404 // RFC 7231, 6.5.4
export const StatusMethodNotAllowed = 405 // RFC 7231, 6.5.5
export const StatusNotAcceptable = 406 // RFC 7231, 6.5.6
export const StatusProxyAuthRequired = 407 // RFC 7235, 3.2
export const StatusRequestTimeout = 408 // RFC 7231, 6.5.7
export const StatusConflict = 409 // RFC 7231, 6.5.8
export const StatusGone = 410 // RFC 7231, 6.5.9
export const StatusLengthRequired = 411 // RFC 7231, 6.5.10
export const StatusPreconditionFailed = 412 // RFC 7232, 4.2
export const StatusRequestEntityTooLarge = 413 // RFC 7231, 6.5.11
export const StatusRequestURITooLong = 414 // RFC 7231, 6.5.12
export const StatusUnsupportedMediaType = 415 // RFC 7231, 6.5.13
export const StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
export const StatusExpectationFailed = 417 // RFC 7231, 6.5.14
export const StatusTeapot = 418 // RFC 7168, 2.3.3
export const StatusMisdirectedRequest = 421 // RFC 7540, 9.1.2
export const StatusUnprocessableEntity = 422 // RFC 4918, 11.2
export const StatusLocked = 423 // RFC 4918, 11.3
export const StatusFailedDependency = 424 // RFC 4918, 11.4
export const StatusTooEarly = 425 // RFC 8470, 5.2.
export const StatusUpgradeRequired = 426 // RFC 7231, 6.5.15
export const StatusPreconditionRequired = 428 // RFC 6585, 3
export const StatusTooManyRequests = 429 // RFC 6585, 4
export const StatusRequestHeaderFieldsTooLarge = 431 // RFC 6585, 5
export const StatusUnavailableForLegalReasons = 451 // RFC 7725, 3
export const StatusInternalServerError = 500 // RFC 7231, 6.6.1
export const StatusNotImplemented = 501 // RFC 7231, 6.6.2
export const StatusBadGateway = 502 // RFC 7231, 6.6.3
export const StatusServiceUnavailable = 503 // RFC 7231, 6.6.4
export const StatusGatewayTimeout = 504 // RFC 7231, 6.6.5
export const StatusHTTPVersionNotSupported = 505 // RFC 7231, 6.6.6
export const StatusVariantAlsoNegotiates = 506 // RFC 2295, 8.1
export const StatusInsufficientStorage = 507 // RFC 4918, 11.5
export const StatusLoopDetected = 508 // RFC 5842, 7.2
export const StatusNotExtended = 510 // RFC 2774, 7
export const StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
```

## HTTP 头（`镜像 Go net/http 包`）

```js
export const HeaderAuthorization = 'Authorization'
export const HeaderProxyAuthenticate = 'Proxy-Authenticate'
export const HeaderProxyAuthorization = 'Proxy-Authorization'
export const HeaderWWWAuthenticate = 'WWW-Authenticate'
export const HeaderAge = 'Age'
export const HeaderCacheControl = 'Cache-Control'
export const HeaderClearSiteData = 'Clear-Site-Data'
export const HeaderExpires = 'Expires'
export const HeaderPragma = 'Pragma'
export const HeaderWarning = 'Warning'
export const HeaderAcceptCH = 'Accept-CH'
export const HeaderAcceptCHLifetime = 'Accept-CH-Lifetime'
export const HeaderContentDPR = 'Content-DPR'
export const HeaderDPR = 'DPR'
export const HeaderEarlyData = 'Early-Data'
export const HeaderSaveData = 'Save-Data'
export const HeaderViewportWidth = 'Viewport-Width'
export const HeaderWidth = 'Width'
export const HeaderETag = 'ETag'
export const HeaderIfMatch = 'If-Match'
export const HeaderIfModifiedSince = 'If-Modified-Since'
export const HeaderIfNoneMatch = 'If-None-Match'
export const HeaderIfUnmodifiedSince = 'If-Unmodified-Since'
export const HeaderLastModified = 'Last-Modified'
export const HeaderVary = 'Vary'
export const HeaderConnection = 'Connection'
export const HeaderKeepAlive = 'Keep-Alive'
export const HeaderAccept = 'Accept'
export const HeaderAcceptCharset = 'Accept-Charset'
export const HeaderAcceptEncoding = 'Accept-Encoding'
export const HeaderAcceptLanguage = 'Accept-Language'
export const HeaderCookie = 'Cookie'
export const HeaderExpect = 'Expect'
export const HeaderMaxForwards = 'Max-Forwards'
export const HeaderSetCookie = 'Set-Cookie'
export const HeaderAccessControlAllowCredentials = 'Access-Control-Allow-Credentials'
export const HeaderAccessControlAllowHeaders = 'Access-Control-Allow-Headers'
export const HeaderAccessControlAllowMethods = 'Access-Control-Allow-Methods'
export const HeaderAccessControlAllowOrigin = 'Access-Control-Allow-Origin'
export const HeaderAccessControlExposeHeaders = 'Access-Control-Expose-Headers'
export const HeaderAccessControlMaxAge = 'Access-Control-Max-Age'
export const HeaderAccessControlRequestHeaders = 'Access-Control-Request-Headers'
export const HeaderAccessControlRequestMethod = 'Access-Control-Request-Method'
export const HeaderOrigin = 'Origin'
export const HeaderTimingAllowOrigin = 'Timing-Allow-Origin'
export const HeaderXPermittedCrossDomainPolicies = 'X-Permitted-Cross-Domain-Policies'
export const HeaderDNT = 'DNT'
export const HeaderTk = 'Tk'
export const HeaderContentDisposition = 'Content-Disposition'
export const HeaderContentEncoding = 'Content-Encoding'
export const HeaderContentLanguage = 'Content-Language'
export const HeaderContentLength = 'Content-Length'
export const HeaderContentLocation = 'Content-Location'
export const HeaderContentType = 'Content-Type'
export const HeaderForwarded = 'Forwarded'
export const HeaderVia = 'Via'
export const HeaderXForwardedFor = 'X-Forwarded-For'
export const HeaderXForwardedHost = 'X-Forwarded-Host'
export const HeaderXForwardedProto = 'X-Forwarded-Proto'
export const HeaderXForwardedProtocol = 'X-Forwarded-Protocol'
export const HeaderXForwardedSsl = 'X-Forwarded-Ssl'
export const HeaderXUrlScheme = 'X-Url-Scheme'
export const HeaderLocation = 'Location'
export const HeaderFrom = 'From'
export const HeaderHost = 'Host'
export const HeaderReferer = 'Referer'
export const HeaderReferrerPolicy = 'Referrer-Policy'
export const HeaderUserAgent = 'User-Agent'
export const HeaderAllow = 'Allow'
export const HeaderServer = 'Server'
export const HeaderAcceptRanges = 'Accept-Ranges'
export const HeaderContentRange = 'Content-Range'
export const HeaderIfRange = 'If-Range'
export const HeaderRange = 'Range'
export const HeaderContentSecurityPolicy = 'Content-Security-Policy'
export const HeaderContentSecurityPolicyReportOnly = 'Content-Security-Policy-Report-Only'
export const HeaderCrossOriginResourcePolicy = 'Cross-Origin-Resource-Policy'
export const HeaderExpectCT = 'Expect-CT'
export const HeaderFeaturePolicy = 'Feature-Policy'
export const HeaderPublicKeyPins = 'Public-Key-Pins'
export const HeaderPublicKeyPinsReportOnly = 'Public-Key-Pins-Report-Only'
export const HeaderStrictTransportSecurity = 'Strict-Transport-Security'
export const HeaderUpgradeInsecureRequests = 'Upgrade-Insecure-Requests'
export const HeaderXContentTypeOptions = 'X-Content-Type-Options'
export const HeaderXDownloadOptions = 'X-Download-Options'
export const HeaderXFrameOptions = 'X-Frame-Options'
export const HeaderXPoweredBy = 'X-Powered-By'
export const HeaderXXSSProtection = 'X-XSS-Protection'
export const HeaderLastEventID = 'Last-Event-ID'
export const HeaderNEL = 'NEL'
export const HeaderPingFrom = 'Ping-From'
export const HeaderPingTo = 'Ping-To'
export const HeaderReportTo = 'Report-To'
export const HeaderTE = 'TE'
export const HeaderTrailer = 'Trailer'
export const HeaderTransferEncoding = 'Transfer-Encoding'
export const HeaderSecWebSocketAccept = 'Sec-WebSocket-Accept'
export const HeaderSecWebSocketExtensions = 'Sec-WebSocket-Extensions'
export const HeaderSecWebSocketKey = 'Sec-WebSocket-Key'
export const HeaderSecWebSocketProtocol = 'Sec-WebSocket-Protocol'
export const HeaderSecWebSocketVersion = 'Sec-WebSocket-Version'
export const HeaderAcceptPatch = 'Accept-Patch'
export const HeaderAcceptPushPolicy = 'Accept-Push-Policy'
export const HeaderAcceptSignature = 'Accept-Signature'
export const HeaderAltSvc = 'Alt-Svc'
export const HeaderDate = 'Date'
export const HeaderIndex = 'Index'
export const HeaderLargeAllocation = 'Large-Allocation'
export const HeaderLink = 'Link'
export const HeaderPushPolicy = 'Push-Policy'
export const HeaderRetryAfter = 'Retry-After'
export const HeaderServerTiming = 'Server-Timing'
export const HeaderSignature = 'Signature'
export const HeaderSignedHeaders = 'Signed-Headers'
export const HeaderSourceMap = 'SourceMap'
export const HeaderUpgrade = 'Upgrade'
export const HeaderXDNSPrefetchControl = 'X-DNS-Prefetch-Control'
export const HeaderXPingback = 'X-Pingback'
export const HeaderXRequestID = 'X-Request-ID'
export const HeaderXRequestedWith = 'X-Requested-With'
export const HeaderXRobotsTag = 'X-Robots-Tag'
export const HeaderXUACompatible = 'X-UA-Compatible'
export const HeaderAccessControlAllowPrivateNetwork = 'Access-Control-Allow-Private-Network'
export const HeaderAccessControlRequestPrivateNetwork = 'Access-Control-Request-Private-Network'
```
