# chiblog backend
## An awesome blogging system backend.
在这里很荣幸告诉大家

这个backend将会用新技术栈

使用GO(Fiber) + MySQL(后续可能会加SQLite,PgSQL和MongoDB)

### 原来不是写Node.js的吗？

是的，但是Go性能更高~

以后js还是会写的~

### 关于fiber的争议问题

fiber本身是个后起之秀（fiber就已经19k了），但是express-friendly的API设计还是令人耳目一新的。

原来是想用gin，但是后面考虑了一下，然后又做了性能测试，发现fiber更胜一筹。所以选了fiber

### How to use?

简单使用

```bash
git clone https://github.com/chihuo2104/chiblog-backend
cd chiblog-backend
go get
go run app.go
```

注意！ 在使用之前先配置好config中的config.sample.go后面改至config.go才可以运行！否则会报错！！！

Docker版本将在稍后推出，敬请期待~