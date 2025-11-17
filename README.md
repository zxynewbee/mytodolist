# TODO List 项目说明文档

## 1. 技术选型
- **编程语言**：Go 1.13，理由：语言熟练，开发方便，性能优秀。  
- **Web框架**：Gin v1.7.7，理由：高性能HTTP框架，生态成熟，中间件丰富。  
- **ORM框架**：GORM v1.9.16，理由：功能强大的Go语言ORM库，支持多种数据库。  
- **数据库**：MySQL，理由：关系型数据库，学习成本低，使用方便。  
- **前端技术**：Vue.js + Element UI，理由：现代化前端框架，组件丰富，开发效率高。

## 2. 项目结构设计 
- 目录结构示例：  
  ```
  conf/ 配置文件
  cmd/ 数据库表创建脚本
  controller/ 
  dao/ 数据库操作
  models/ 结构体集合
  routers/ 路由
  setting/ 配置文件设置
  static/ 静态文件
  templates/ 模版文件
  main.go 主程序运行入口
  ```  

## 3. 功能特性

### 核心功能
- 任务增删改查（CRUD）操作
- 任务状态管理（完成/未完成）
- RESTful API接口
- 响应式前端界面
- 数据库自动迁移


## 4. 数据库设计

### todos表结构
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | INT AUTO_INCREMENT | 主键，自增ID |
| title | VARCHAR(255) | 任务标题，必填 |
| status | TINYINT | 完成状态，默认0 |

## 5. 运行与测试方式

### 环境要求
- Go 1.13+ 
- MySQL 5.7+
- 现代浏览器（Chrome/Firefox/Safari/Edge）

### 本地运行步骤

1. **克隆项目**
   ```bash
   git clone <项目地址>
   cd bubble
   ```

2. **安装依赖**
   ```bash
   go mod download
   ```

3. **配置数据库**
   - 创建MySQL数据库：`CREATE DATABASE bubble;`
   - 修改配置文件 `conf/config.ini`：
     ```ini
     [mysql]
     user = root
     password = your_password
     host = 127.0.0.1
     port = 3306
     dbname = bubble
     ```

4. **创建数据库表**
   ```bash
   cd cmd
   go run sql_create.go
   ```

5. **启动应用**
   ```bash
   go run main.go
   ```

6. **访问应用**
   打开浏览器访问：http://localhost:8080


### API接口测试
使用curl或Postman测试API：

```bash
# 获取所有任务
curl http://localhost:8080/api/v1/todo

# 创建新任务
curl -X POST http://localhost:8080/api/v1/todo -d 'title=学习Go语言'

# 更新任务状态
curl -X PUT http://localhost:8080/api/v1/todo/1 -d 'status=1'

# 删除任务
curl -X DELETE http://localhost:8080/api/v1/todo/1
```


## 6.改进方向
如果有更多时间，可以考虑以下改进：
   - 添加用户认证和授权系统
   - 实现任务分类和标签功能
   - 添加任务优先级和过期时间管理
   - 支持任务搜索和过滤
   - 添加批量操作功能