# Go Word Server

这是一个基于Go语言开发的单词学习服务器应用，提供用户管理、单词数据管理、云端生词本、打印导出和学习功能的后端API服务。

## 功能特性

- 用户注册和登录
- JWT身份验证
- 单词数据管理
- 用户单词学习记录
- RESTful API接口
- MySQL数据库支持

## 技术栈

- **后端**: Go (Gin框架)
- **数据库**: MySQL
- **认证**: JWT
- **前端**: HTML/CSS/JavaScript (静态文件)

## 安装要求

- Go 1.19 或更高版本
- MySQL 5.7 或更高版本
- Git

## 安装步骤

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd go-test
   ```

2. **安装依赖**
   ```bash
   go mod download
   ```

3. **设置数据库**
   - 安装并启动MySQL服务
   - 创建数据库：
     ```sql
     CREATE DATABASE wordserver;
     ```
   - 导入数据库结构：
     ```bash
     mysql -u username -p wordserver < mysql.sql
     ```

4. **配置数据库连接**
   - 编辑 `internal/config/database.go` 文件，修改数据库连接参数：
     ```go
     dsn := "username:password@tcp(localhost:3306)/wordserver?charset=utf8mb4&parseTime=True&loc=Local"
     ```

## 运行应用

1. **启动服务器**
   ```bash
   go run main.go
   ```

2. **访问应用**
   - 打开浏览器访问 `http://localhost:8080`
   - API基础路径：`http://localhost:8080/api/v1`

## API使用说明

详细的API文档请参考 [API.md](API.md) 文件。

### 主要API端点

#### 用户相关
- `POST /api/v1/user/register` - 用户注册
- `POST /api/v1/user/login` - 用户登录
- `GET /api/v1/user/profile` - 获取用户信息

#### 单词相关
- `GET /api/v1/words` - 获取单词列表
- `GET /api/v1/words/{id}` - 获取单词详情
- `POST /api/v1/words` - 添加新单词

#### 用户单词学习
- `GET /api/v1/user/words` - 获取用户学习记录
- `POST /api/v1/user/words` - 添加用户单词
- `PUT /api/v1/user/words/{id}` - 更新学习进度

### 请求示例

#### 用户注册
```bash
curl -X POST http://localhost:8080/api/v1/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123",
    "email": "test@example.com"
  }'
```

#### 获取单词列表
```bash
curl -X GET http://localhost:8080/api/v1/words \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 项目结构

```
go-test/
├── API.md                 # API文档
├── go.mod                 # Go模块文件
├── main.go                # 应用入口
├── mysql.sql              # 数据库初始化脚本
├── wordserver             # 编译后的可执行文件
├── data/                  # 数据文件目录
├── internal/              # 内部代码
│   ├── config/            # 配置相关
│   │   └── database.go    # 数据库配置
│   ├── handler/           # 请求处理器
│   │   ├── user_handler.go
│   │   ├── word_handler.go
│   │   └── user_new_word_handler.go
│   ├── middleware/        # 中间件
│   │   ├── CROS.go        # 跨域处理
│   │   └── JWTAuth.go     # JWT认证
│   ├── models/            # 数据模型
│   │   ├── user.go
│   │   ├── word.go
│   │   ├── word_detail.go
│   │   └── user_new_word.go
│   ├── routers/           # 路由配置
│   │   └── routers.go
│   └── tool/              # 工具函数
│       ├── JsonTool.go
│       └── PageTool.go
├── static/                # 静态文件
│   ├── css/
│   │   └── main.css
│   └── js/
│       ├── axios.min.js
│       └── word_data.js
└── templates/             # HTML模板
    └── index.html
```

## 开发说明

### 构建应用
```bash
go build -o wordserver main.go
```

### 运行测试
```bash
go test ./...
```

### 代码格式化
```bash
go fmt ./...
```

## 部署说明

1. 编译为可执行文件：
   ```bash
   GOOS=linux GOARCH=amd64 go build -o wordserver main.go
   ```

2. 上传到服务器并配置数据库

3. 启动服务：
   ```bash
   ./wordserver
   ```

## 常见问题

### 数据库连接失败
- 检查MySQL服务是否启动
- 验证数据库连接字符串是否正确
- 确保数据库和表已创建

### JWT认证失败
- 检查token是否过期
- 验证token格式是否正确
- 确认用户凭据是否有效

### 跨域问题
- 确保前端请求包含正确的CORS头
- 检查中间件配置