# Librarian

一个[GoAdmin](https://www.go-admin.cn)插件，帮助您快速生成一个Markdown文档站点。

![](http://quick.go-admin.cn/docs/librarian_interface_2.png)

## 快速开始

### 使用CLI

- 下载对应系统的二进制执行文件 [cli](https://github.com/GoAdminGroup/librarian/releases/tag/v0.0.1)
- 下载 [数据库](https://github.com/GoAdminGroup/librarian/releases/download/v0.0.1/librarian.db)，并放在项目文件夹下面
- 新建一个 nav.yml 在你的文档文件夹下，这里有个[例子](https://github.com/GoAdminGroup/librarian/blob/master/cli/docs/nav.yml)
- 运行以下命令：

```bash
./librarian --port=9055 --path=./docs
```

- 访问：[http://localhost:9055/docs](http://localhost:9055/docs)


### 在你GoAdmin系统中使用

```
cd example
GO111MODULE=on go run main.go
```

访问：[http://localhost:9033/admin/librarian](http://localhost:9033/admin/librarian)

## TODO

- [ ] 文档加密
- [ ] 编辑新建文档
- [ ] 实时监听nav.yml