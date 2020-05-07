# Librarian

A GoAdmin plugin which see markdown docs.

![](http://quick.go-admin.cn/docs/librarian_interface_2.png)

## How To

### CLI

- download the cli
- download the database
- run

```bash
./librarian --port=9055 --path=./docs
```

- visit: [http://localhost:9055/docs](http://localhost:9055/docs)


### Use in your GoAdmin app

An example: 

```
cd example
GO111MODULE=on go run main.go
```

visit: [http://localhost:9033/admin/librarian](http://localhost:9033/admin/librarian)

## TODO

- [ ] 文档加密
- [ ] 编辑新建文档
- [ ] 实时监听nav.yml