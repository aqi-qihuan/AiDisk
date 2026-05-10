# 📚 AqiCloud-AiPan 前端 API 文档

> 一款基于 Vue 3 + TypeScript 的智能云盘系统前端 API 接口文档

---

## 📋 目录

- [📌 项目信息](#📌-项目信息)
- [📤 通用响应格式](#📤-通用响应格式)
- [👤 一、用户管理模块](#一用户管理模块)
- [📁 二、文件管理模块](#二文件管理模块)
- [🔗 三、分享管理模块](#三分享管理模块)
- [🗑️ 四、回收站模块](#四回收站模块)
- [🤖 五、AI 模块](#五ai-模块)
- [📝 数据类型定义](#📝-数据类型定义)
- [📊 状态码说明](#📊-状态码说明)
- [⚠️ 注意事项](#⚠️-注意事项)

---

## 📌 项目信息

| 项目 | 说明 |
|------|------|
| 🏷️ **项目名称** | AqiCloud-AiPan |
| 📦 **版本** | v1.2.0 |
| ⚛️ **前端框架** | Vue 3.2.13 + TypeScript 4.9.5 |
| 🎨 **UI 框架** | Element Plus 2.8.6 + 自定义设计系统 |
| 🌐 **HTTP 客户端** | Axios 1.7.7 |
| 🔧 **构建工具** | Vue CLI 5.0 |
| 📝 **状态管理** | Pinia 2.2.4 |
| 🗂️ **路由管理** | Vue Router 4.0.3 |
| 📱 **响应式设计** | 移动优先，4个断点适配 |

---

## 📤 通用响应格式

所有 API 返回的数据格式如下：

```typescript
interface JsonData<T = any> {
  code: number;       // 状态码 (0 表示成功)
  data: T;           // 响应数据
  message: string;   // 响应消息
}
```

### 响应状态码

| 状态码 | 说明 |
|--------|------|
| 0 | 成功 |
| 400001 | 参数错误 |
| 500001 | 服务器内部错误 |
| 260101 | 分享不存在或已过期 |
| 260102 | 分享码错误 |
| 260103 | 分享已过期 |
| 260104 | 分享已取消 |
| 260401 | 文件不存在 |
| 260402 | 文件已存在 |
| 260405 | 文件类型不合规 |
| 260408 | 分享的文件不合规 |

---

## 👤 一、用户管理模块

### 1.1 获取当前用户信息 🧑

获取当前登录用户的详细信息

**接口**: `GET /account/v1/detail`

**请求参数**: 无

**响应示例**:
```json
{
  "code": 200,
  "success": true,
  "data": {
    "id": 1,
    "username": "用户名",
    "avatarUrl": "头像URL",
    "rootFileId": "根文件ID",
    "rootFileName": "根文件夹名",
    "email": "邮箱",
    "capacity": 10737418240,
    "storageDTO": {
      "totalSize": "10737418240",
      "usedSize": "5368709120"
    },
    "role": 1,
    "createTime": "2024-01-01 00:00:00",
    "updateTime": "2024-01-01 00:00:00"
  }
}
```

---

### 1.2 用户登录 🔐

用户登录接口

**接口**: `POST /account/v1/login`

**请求参数**:
```typescript
interface UserLoginRequest {
  phone: string;    // 用户名/手机号
  password: string; // 用户密码
}
```

**响应示例**:
```json
{
  "code": 200,
  "success": true,
  "data": "token_value",
  "msg": "登录成功"
}
```

---

### 1.3 用户注册 📝

用户注册接口

**接口**: `POST /account/v1/register`

**请求参数**:
```typescript
interface UserRegisterRequest {
  username?: string;   // 用户名
  phone?: string;     // 手机号
  password?: string;   // 密码
  avatarUrl?: string;  // 用户头像
}
```

---

### 1.4 更新用户信息（管理员）⚙️

管理员更新用户信息

**接口**: `POST /v1/user/update`

**请求参数**:
```typescript
interface UserUpdateRequest {
  userId?: number;     // 用户ID
  username?: string;   // 用户名
  avatarUrl?: string;  // 头像URL
  email?: string;      // 邮箱
  capacity?: number;    // 容量
  role?: number;       // 角色
}
```

---

### 1.5 更新当前用户信息 ✏️

用户更新自己的信息

**接口**: `POST /account/v1/update/my`

**请求参数**:
```typescript
// Query 参数
interface updateUserInfoParams {
  file?: string;  // 文件
}

// Body 参数
interface UserUpdateMyRequest {
  userId?: number;     // 用户ID
  username?: string;   // 用户名
  avatarUrl?: string;  // 头像URL
}
```

---

### 1.6 上传用户头像 📷

上传用户头像

**接口**: `POST /account/v1/upload_avatar`

**请求参数**: `multipart/form-data`

---

### 1.7 分页获取用户列表（管理员）📋

管理员分页查询用户列表

**接口**: `POST /v1/user/list`

**请求参数**:
```typescript
interface UserQueryRequest {
  current?: number;      // 当前页号
  pageSize?: number;    // 页面大小
  sortField?: string;    // 排序字段
  sortOrder?: string;    // 排序顺序（默认升序）
  userId?: number;       // 用户ID
  username?: string;     // 用户名
  role?: string;        // 角色
  createTime?: string;   // 创建时间
  updateTime?: string;   // 更新时间
}
```

---

## 📁 二、文件管理模块

### 2.1 分页查询文件列表 🔍

查询指定父目录下的文件列表

**接口**: `GET /file/v1/list`

**请求参数**:
```typescript
{
  parent_id: string;  // 父文件夹ID
}
```

**响应示例**:
```json
{
  "code": 200,
  "success": true,
  "data": [
    {
      "id": "file_id",
      "fileName": "文件名",
      "fileType": "文件类型",
      "fileSuffix": "文件后缀",
      "fileSize": 1024,
      "filePath": "/path/to/file",
      "fileUrl": "文件URL",
      "isDirectory": 0,
      "updateTime": "2024-01-01 00:00:00"
    }
  ]
}
```

---

### 2.2 根据条件查询文件列表 🔎

根据文件名搜索文件

**接口**: `GET /file/v1/search`

**请求参数**:
```typescript
{
  search: string;  // 搜索关键词
}
```

---

### 2.3 获取文件树 🌳

获取文件夹树结构

**接口**: `GET /file/v1/folder/tree`

**响应示例**:
```json
{
  "code": 200,
  "data": [
    {
      "id": "folder_id",
      "label": "文件夹名",
      "depth": 1,
      "children": []
    }
  ]
}
```

---

### 2.4 创建文件夹 📂

创建新文件夹

**接口**: `POST /file/v1/create_folder`

**请求参数**:
```typescript
interface CreateFileRequest {
  filePath?: string;    // 文件路径
  fileName?: string;     // 文件名
  folderName?: string;   // 文件夹名称
  parentId?: string;     // 父文件夹ID
  accountId?: number;    // 用户ID
}
```

---

### 2.5 分片上传相关 📤

#### 2.5.1 获取分片上传进度 📊

获取文件分片上传进度

**接口**: `GET /file/v1/chunk_upload_progress/{identifier}`

**请求参数**:
```typescript
{
  identifier: string;  // 文件唯一标识
}
```

**响应示例**:
```json
{
  "code": 200,
  "data": {
    "finished": false,
    "path": "/path/to/file",
    "taskRecord": {
      "uploadId": "upload_id",
      "fileIdentifier": "md5_value",
      "fileName": "文件名",
      "totalSize": 1024000,
      "chunkSize": 524288,
      "chunkNum": 2,
      "exitPartList": [
        {
          "partNumber": 1,
          "eTag": "etag_value",
          "size": 524288
        }
      ]
    }
  }
}
```

---

#### 2.5.2 创建分片上传任务 ▶️

初始化分片上传任务

**接口**: `POST /file/v1/init_file_chunk_task`

**请求参数**:
```typescript
interface ChunkInitTaskRequest {
  identifier: string;   // 文件唯一标识（MD5）
  totalSize: number;    // 文件总大小
  chunkSize: number;   // 分片大小
  filename: string;    // 文件名
  accountId: number;    // 用户ID
}
```

---

#### 2.5.3 获取分片上传地址 🔗

获取分片的预签名上传 URL

**接口**: `GET /file/v1/get_file_chunk_upload_url/{identifier}/{partNumber}`

**请求参数**:
```typescript
{
  identifier: string;   // 文件唯一标识
  partNumber: number;   // 分片编号
}
```

---

#### 2.5.4 合并分片 🔄

合并所有分片为完整文件

**接口**: `POST /file/v1/merge_file_chunk`

**请求参数**:
```typescript
interface FileMergeRequest {
  parentId?: number;     // 父文件夹ID
  identifier?: string;   // 文件唯一标识
  accountId?: number;    // 用户ID
}
```

---

### 2.6 批量删除文件 🗑️

批量删除文件到回收站

**接口**: `POST /file/v1/del_batch`

**请求参数**:
```typescript
interface BatchFilesRequest {
  fileIds?: string[];      // 文件ID列表
  targetParentId?: string; // 目标父ID
  accountId?: number;      // 用户ID
}
```

---

### 2.7 批量移动文件 ➡️

批量移动文件到指定目录

**接口**: `POST /file/v1/move_batch`

**请求参数**: `BatchFilesRequest`

---

### 2.8 批量复制文件 📋

批量复制文件到指定目录

**接口**: `POST /file/v1/copy_batch`

**请求参数**: `BatchFilesRequest`

---

### 2.9 文件重命名 ✏️

重命名文件或文件夹

**接口**: `POST /file/v1/rename_file`

**请求参数**:
```typescript
interface FileUpdateRequest {
  fileId?: string;        // 文件ID
  newFilename?: string;    // 新文件名
  accountId?: number;     // 用户ID
  filePath?: string;      // 文件路径
  isDirectory?: number;  // 是否为文件夹（0否，1是）
}
```

---

### 2.10 获取批量下载地址 ⬇️

获取文件的下载链接

**接口**: `POST /file/v1/batch_download_url`

**请求参数**: `BatchFilesRequest`

**响应示例**:
```json
{
  "code": 200,
  "data": [
    {
      "downloadUrl": "https://...",
      "fileName": "文件名.pdf"
    }
  ]
}
```

---

### 2.11 单文件下载 📥

下载单个文件

**接口**: `POST /file/v1/download`

**请求参数**:
```typescript
interface downloadUrlParam {
  fileIds?: string[];  // 文件ID列表
  accountId?: number;   // 用户ID
}
```

**响应类型**: `blob`

---

### 2.12 文件预览 👁️

在线预览文件

**接口**: `GET /v1/file/preview`

**请求参数**:
```typescript
interface previewParams {
  fileId: string;  // 文件ID
}
```

---

### 2.13 文件秒传 ⚡

使用文件 MD5 进行秒传

**接口**: `POST /file/v1/second_upload`

**请求参数**:
```typescript
interface SecUploadRequest {
  parentId?: string;   // 父文件夹ID
  filename?: string;  // 文件名
  accountId?: string;  // 用户ID
  identifier?: string; // 文件唯一标识（MD5）
}
```

---

### 2.14 单文件上传 ⬆️

上传单个文件

**接口**: `POST /file/v1/upload`

**请求参数**: `multipart/form-data`

```typescript
interface UploadFile {
  filename?: string;   // 文件名
  identifier?: string; // 文件标识
  accountId?: number;  // 用户ID
  parentId?: string;  // 父文件夹ID
  fileSize?: number;  // 文件大小
  file?: File;       // 文件对象
}
```

---

## 🔗 三、分享管理模块

### 3.1 访问分享链接 🔗

访问分享链接查看分享信息

**接口**: `GET /share/v1/visit`

**请求参数**:
```typescript
{
  shareId: string;  // 分享ID
}
```

---

### 3.2 创建分享链接 ➕

创建文件/文件夹分享链接

**接口**: `POST /share/v1/create`

**请求参数**:
```typescript
interface CreateShareRequest {
  shareName?: string;   // 分享名称
  shareType?: string;   // 分享类型
  shareDayType?: number; // 分享有效期（0永久，1:7天，2:30天）
  fileIds?: string[];   // 分享的文件ID列表
  accountId?: number;   // 用户ID
  host?: string;        // 分享域名
}
```

---

### 3.3 校验分享码 🔐

校验分享提取码

**接口**: `POST /share/v1/check_share_code`

**请求参数**:
```typescript
interface CheckShareRequest {
  shareId: string;   // 分享ID
  shareCode?: string; // 分享码
}
```

---

### 3.4 获取分享详情 📋

获取当前用户的分享列表

**接口**: `GET /share/v1/detail`

**响应示例**:
```json
{
  "code": 200,
  "data": [
    {
      "shareId": "share_id",
      "shareName": "分享名称",
      "shareUrl": "https://...",
      "shareCode": "1234",
      "shareStatus": 1,
      "shareType": "FILE",
      "shareDayType": 0,
      "shareEndTime": "2024-12-31 23:59:59",
      "createTime": "2024-01-01 00:00:00"
    }
  ]
}
```

---

### 3.5 获取分享详情（需要校验）🔐

获取需要校验的分享详情

**接口**: `GET /v1/share/detail-with-code`

---

### 3.6 分享文件列表查询 📁

查询分享的文件列表

**接口**: `POST /share/v1/list_share_file`

**请求参数**:
```typescript
interface ShareFileQueryRequest {
  parentId?: string;   // 父文件夹ID
  current?: number;    // 当前页号
  pageSize?: number;   // 页面大小
  sortField?: string;  // 排序字段
  sortOrder?: string;  // 排序顺序
  shareId?: string;    // 分享ID
  fileName?: string;   // 文件名
  fileType?: number;   // 文件类型
  isDirectory?: number; // 是否为文件夹
}
```

---

### 3.7 分享文件列表查询（需要校验）🔐

查询需要校验的分享文件列表

**接口**: `POST /v1/share/share-list-with-code`

---

### 3.8 文件转存 💾

将分享的文件保存到自己的网盘

**接口**: `POST /share/v1/transfer`

**请求参数**:
```typescript
interface SaveShareRequest {
  shareId?: string;    // 分享ID
  fileIds?: string[];  // 要保存的文件ID列表
  accountId?: number;   // 用户ID
  parentId?: string;   // 保存位置
}
```

---

### 3.9 文件转存（需要校验）💾

将需要校验的分享文件保存到网盘

**接口**: `POST /v1/share/save-with-code`

---

### 3.10 取消分享 ❌

取消分享链接

**接口**: `POST /share/v1/cancel`

**请求参数**:
```typescript
interface CancelShareRequest {
  shareIds?: string[]; // 分享ID列表
  accountId?: number;  // 用户ID
}
```

---

### 3.11 获取分享列表 📋

获取当前用户的所有分享

**接口**: `GET /share/v1/list`

---

## 🗑️ 四、回收站模块

### 4.1 查询回收站文件列表 📋

获取回收站中的文件列表

**接口**: `GET /recycle/v1/list`

---

### 4.2 批量彻底删除 🗑️

彻底删除回收站中的文件

**接口**: `POST /recycle/v1/delete`

**请求参数**: `BatchFilesRequest`

---

### 4.3 批量还原文件 ↩️

将回收站中的文件还原到原位置

**接口**: `POST /recycle/v1/restore`

**请求参数**: `BatchFilesRequest`

---

## 五、AI 模块

当前 AI 模块暂无接口实现，保留扩展功能。

---

## 📝 数据类型定义

### 用户相关

```typescript
interface UserDTO {
  id?: number;              // 用户ID
  rootFileId?: string;      // 根文件ID
  rootFileName?: string;    // 根文件夹名
  username?: string;        // 用户名
  avatarUrl?: string;       // 用户头像URL
  email?: string;           // 用户邮箱
  capacity?: number;        // 用户容量大小
  storageDTO: sDTO;        // 用户存储信息
  role?: number;            // 用户角色
  createTime?: string;      // 创建时间
  updateTime?: string;      // 更新时间
}

interface sDTO {
  accountId?: string;  // 账户ID
  id?: string;        // ID
  totalSize: string;  // 总大小
  usedSize: string;   // 已使用大小
}
```

### 文件相关

```typescript
interface FileDTO {
  isDir?: number;                  // 是否为目录
  id?: string;                     // 文件ID
  accountId?: number;              // 用户ID
  parentId?: string;               // 父文件夹ID
  fileId?: string;                 // 文件ID
  fileName?: string;               // 文件名称
  fileType?: string;               // 文件类型
  fileSuffix?: string;             // 文件后缀
  fileSize?: number;               // 文件大小
  filePath?: string;               // 文件路径
  fileUrl?: string;               // 文件URL
  fileIdentifier?: string;          // 唯一标识（MD5）
  fileBucketName?: string;         // 存储桶名
  objectKey?: string;             // MinIO对象键
  filePreviewContentType?: string;  // 预览Content-Type
  updateTime?: string;             // 更新时间
  gmtModified?: string;           // 修改时间
  isDirectory?: number;           // 是否为文件夹（0否，1是）
}
```

### 分片上传相关

```typescript
interface FileChunkInfoDTO {
  finished?: boolean;             // 是否完成上传
  path?: string;                 // 文件地址
  taskRecord?: FileChunkRecordDTO; // 上传记录
}

interface FileChunkRecordDTO {
  exitPartList?: PartSummary[]; // 已上传的分片
  uploadId?: string;            // 上传ID
  fileIdentifier?: string;       // 文件唯一标识
  fileName?: string;            // 文件名
  bucketName?: string;          // 桶名
  objectKey?: string;          // 对象键
  totalSize?: number;           // 文件大小
  chunkSize?: number;           // 分片大小
  chunkNum?: number;            // 分片数量
  userId?: number;              // 用户ID
}

interface PartSummary {
  partNumber?: number;    // 分片编号
  lastModified?: string;  // 最后修改时间
  eTag?: string;        // ETag
  size?: number;        // 分片大小
}
```

### 分享相关

```typescript
interface ShareDTO {
  shareId?: string;         // 分享ID
  id?: string;             // ID
  shareName?: string;       // 分享名称
  shareUrl?: string;       // 分享链接URL
  shareCode?: string;      // 分享码
  shareStatus?: number;     // 分享状态
  shareType?: string;      // 分享类型
  shareDayType?: number;    // 分享有效期类型
  shareEndTime?: string;    // 分享过期时间
  createTime?: string;      // 创建时间
}

interface ShareDetailDTO {
  shareId?: string;                // 分享ID
  shareName?: string;              // 分享名称
  shareUrl?: string;              // 分享链接URL
  shareCode?: string;             // 分享码
  shareStatus?: number;            // 分享状态
  shareType?: string;             // 分享类型 (no_code/need_code)
  shareDayType?: number;           // 分享有效期类型 (0永久/1七天/2三十天)
  shareEndTime?: string;           // 分享过期时间
  createTime?: string;             // 创建时间
  fileDTOList?: FileDTO[];        // 分享的文件列表
  shareAccountDTO?: shareUserDTO;  // 分享者信息
}

interface ShareSimpleDTO {
  shareId?: string;                // 分享ID
  shareName?: string;              // 分享名称
  shareUrl?: string;              // 分享链接URL
  shareStatus?: number;            // 分享状态
  shareType?: string;             // 分享类型
  shareDayType?: number;           // 分享有效期类型
  shareEndTime?: string;           // 分享过期时间
  createTime?: string;             // 创建时间
  fileCount?: number;             // 分享的文件数量 (v1.2.0新增)
  shareAccountDTO?: shareUserDTO;  // 分享者信息
  shareToken?: string;            // 分享令牌 (无提取码时返回)
}

interface shareUserDTO {
  id?: string;           // 用户ID
  username?: string;     // 用户名
  avatarUrl?: string;    // 头像URL
}
```

### 树形结构相关

```typescript
interface TreeNodeDTO {
  id: string;                         // 节点ID
  label?: string;                      // 节点名称
  depth?: number;                      // 深度
  state?: string;                      // 状态
  attributes?: Record<string, string>;   // 属性集合
  children?: TreeNodeDTO[];            // 子节点列表
}
```

---

## 📊 状态码说明

| 状态码 | 说明 | 图标 |
|--------|------|------|
| 200 | 请求成功 | ✅ |
| 400 | 请求参数错误 | ❌ |
| 401 | 未授权，需要登录 | 🔐 |
| 403 | 无权限访问 | 🚫 |
| 404 | 资源不存在 | ❓ |
| 500 | 服务器内部错误 | ⚠️ |

---

## ⚠️ 注意事项

1. 🔑 **认证**: 大部分接口需要在请求头中携带认证 Token
2. 📤 **分片上传**: 大文件上传建议使用分片上传，提高上传成功率
3. ⚡ **秒传**: 如果服务器已存在相同 MD5 的文件，可以使用秒传快速上传
4. 📅 **分享有效期**: `shareDayType` 值为 0 表示永久，1 表示 7 天，2 表示 30 天
5. 📁 **文件类型**: `isDirectory` 字段，0 表示文件，1 表示文件夹
6. 🕐 **时间格式**: 所有时间字段均为字符串格式，如 `"2024-01-01 00:00:00"`
7. 🔐 **分享访问**: 访问分享详情需要在请求头中携带 `Share-Token`
8. 🔢 **提取码**: 提取码长度支持 4-6 位字母数字组合

---

## 📅 更新日志

| 版本 | 日期 | 说明 |
|------|------|------|
| v1.2.0 | 2026-03-14 | 🎨 新增 ShareSimpleDTO.fileCount 字段，优化分享页面文件数量显示，修复移动端表格问题 |
| v1.1.0 | 2026-03-12 | ✅ 设计系统迁移完成，统一UI组件库 |
| v1.0.0 | 2026-02-01 | 🎉 初始版本，包含用户、文件、分享、回收站模块 |

---

<div align="center">

**© 2026 AqiCloud-AiPan. All rights reserved.**

</div>
