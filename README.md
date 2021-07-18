# Uwall-backend

> 校园万能墙后端

前端地址 `react` : [jsun969/Uwall-frontend](https://github.com/jsun969/Uwall-frontend)

## 如何部署

1. 下载所有文件至服务器 , 并配置好 `golang` 环境
2. 配置数据库
    1. 将`config` 文件夹中的 `database.go.template` 更名为 `database.go`
    2. 按照注释修改 `database.go`
3. 运行程序 `go run main.go`
4. 添加 `nginx` 反向代理 , 推荐配置 (前端api地址也请使用推荐配置)
   ```nginx
   location /api/ {
      proxy_pass http://127.0.0.1:1314/;
   }
   ```

## 审核消息

打开数据库 修改各表中 `status` 字段

## Todo List

- [ ] 图片上传