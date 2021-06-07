## 开始
sql to xorm go model

这个工具将运行一个网络服务，可以将create sql语句转换为一个go的结构体。

主要功能用来自动生成相关表的增删改查xorm go实体代码。

go run main.go

打开浏览器

http://127.0.0.1:8080/sql2go/xorm

然后，输入create的sql语句，就会生成相关的xorm结构

示例：https://www.superpig.win/sql2go/xorm

注意，可能有BUG。可能不能满足所有类型的create sql。
