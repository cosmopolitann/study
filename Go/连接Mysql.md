# GO连接Mysql

###### 1.连接Mysql数据库。

###### golang其实官方不提供连接mysql实现，先下载第三方的实现.

```go
go get github.com/go-sql-driver/mysql
```

###### 2.连接

```go
conn,err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/database")
	defer conn.Close()

root :  ip地址   /数据库名称
```

###### 3.

```go
package main
 
import (
       "fmt"
    _ "github.com/go-sql-driver/mysql" //导入mysql驱动包
)
 
func init() {
 
}
 
func main() {
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
    if err != nil {
        panic(err)
    }
    //增加数据
        stmt, err := db.Prepare(`INSERT student (name,age) values (?,?)`)
        res, err := stmt.Exec()
        id, err := res.LastInsertId()
        fmt.Println("自增id=", id)
    //修改数据
        stmt, err := db.Prepare(`UPDATE student SET age=? WHERE id=?`)
        res, err := stmt.Exec(, )
        num, err := res.RowsAffected() //影响行数
        fmt.Println(num)
    //删除数据
        stmt, err := db.Prepare(`DELETE FROM student WHERE id=?`)
        res, err := stmt.Exec()
        num, err := res.RowsAffected()
        fmt.Println(num)
    //查询数据
    rows, err := db.Query("SELECT * FROM student")
 
    //--------简单一行一行输出---start
    //    for rows.Next() { //满足条件依次下一层
    //        var id int
    //        var name string
    //        var age int
    //        rows.Columns()
 
    //        err = rows.Scan(&id, &name, &age)
    //        fmt.Println(id)
    //        fmt.Println(name)
    //        fmt.Println(age)
    //    }
    //--------简单一行一行输出---end
 
    //--------遍历放入map----start
    //构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
    columns, _ := rows.Columns()
    scanArgs := make([]interface{}, len(columns))
    values := make([]interface{}, len(columns))
 
    for i := range values {
        scanArgs[i] = &values[i]
    }
 
    for rows.Next() {
        //将行数据保存到record字典
        err = rows.Scan(scanArgs...)
        record := make(map[string]string)
        for i, col := range values {
            if col != nil {
                record[columns[i]] = string(col.([]byte))
            }
        }
        fmt.Println(record)
    }
    //--------遍历放入map----end
}
```

