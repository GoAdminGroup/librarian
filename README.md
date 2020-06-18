# Librarian

A GoAdmin plugin which see markdown docs.

[中文文档](./README_CN.md)

![](http://quick.go-admin.cn/docs/librarian_interface_2.png)

## How To

### CLI

- download the [cli](https://github.com/GoAdminGroup/librarian/releases/tag/v0.0.1)
- download the [database](https://github.com/GoAdminGroup/librarian/releases/download/v0.0.1/librarian.db) and put the db in your project folder
- set up nav.yml in your docs folder, here is an [example](https://github.com/GoAdminGroup/librarian/blob/master/cli/docs/nav.yml)
- run:

```bash
./librarian --port=9055 --path=./docs
```

- visit: [http://localhost:9055/docs](http://localhost:9055/docs)


### Use in your GoAdmin app

An [example](https://github.com/GoAdminGroup/librarian/blob/master/example/main.go): 

```
cd example
GO111MODULE=on go run main.go
```

visit: [http://localhost:9033/admin/librarian](http://localhost:9033/admin/librarian)

## TODO

- [ ] Document encryption
- [ ] Edit the new document
- [ ] Real-time monitor nav. Yml