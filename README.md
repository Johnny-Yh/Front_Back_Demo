# 前后端分离Demo

目录说明：

- backend：Go后端（Gin）

- frontend：前端（Vue）

- sql：数据表（SQLite）

### 1 运行环境

#### 1.1 Vue【前端】

##### 运行方式：

```shell
npm install

npm run serve
```

#### 1.2 Go【后端】

##### 运行方式：

```shell
go mod tidy

go run main.go
```

### 3 TODO

- 添加员工的更多属性
- 添加用户注册登录退出功能
- 添加用户权限隔离，只有管理员才能修改、删除用户
- 优化前端界面
- 后端返回值需要规范
- Go的DB操作换成ORM可节省大量代码
- API生成和相关文档工具使用
- http常用状态码
- ......

