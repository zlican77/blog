# 个人技术博客系统

一个使用 Vue3 + Go 开发的现代化个人技术博客系统，具有清晰的文章管理、用户互动和后台管理功能。

## 技术栈

### 前端
- Vue 3
- Vite
- Element Plus
- Markdown 编辑器 (md-editor-v3)
- Prism.js 代码高亮
- Vue Router
- Pinia 状态管理

### 后端
- Go
- Gin 框架
- Gorm
- Redis (Session 存储)
- MySQL

## 主要功能

### 文章系统
- 文章的 CRUD 操作
- Markdown 编辑器支持
- 代码高亮显示
- 文章分类与标签管理
- 文章目录自动生成

### 用户系统
- 用户注册/登录
- 手机验证码登录
- 图片验证码支持
- 用户评论功能

### 其他功能
- 响应式设计
- 深色/浅色主题切换
- 文章搜索
- 留言板功能
- 访问统计

## 项目结构
```tree
blog/
├── blog-frontend/ # 前端项目
│ ├── src/
│ │ ├── components/ # 组件
│ │ ├── views/ # 页面
│ │ ├── router/ # 路由配置
│ │ ├── assets/ # 静态资源
│ │ └── blogapi/ # API 接口
│ └── ...
│
└── blog-backend/ # 后端项目
├── router/ # 路由处理
├── controller/ # 控制器
├── model/ # 数据模型
├── middleware/ # 中间件
└── ...
```

## 开发环境要求

- Node.js >= 16
- Go >= 1.18
- MySQL >= 8.0
- Redis >= 6.0

## 快速开始

1. 克隆项目
```bash
git clone https://github.com/yourusername/blog.git
cd blog
```


2. 前端设置
```bash
cd blog-frontend
npm install
npm run dev
```

4. 后端设置
```bash
cd blog-backend
go mod tidy
go run main.go
```

5. 数据库设置
- 创建 MySQL 数据库
- 导入 `mysql/init.sql` 初始化数据库结构
- 配置 Redis

## 部署

### 前端部署
```bash
cd blog-frontend
npm run build
```

### 后端部署
```bash
cd blog-backend
go build
```


## 配置说明

### 前端配置
- 在 `.env` 文件中配置后端 API 地址
- 在 `vite.config.js` 中配置开发服务器设置

### 后端配置
- 数据库连接配置
- Redis 配置
- 跨域设置
- 会话管理配置

## 贡献指南

1. Fork 本项目
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开一个 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 联系方式

- 作者：陈子陵
- 邮箱：zlican77@gmail.com
- GitHub：[zlican77](https://github.com/zlican77)
