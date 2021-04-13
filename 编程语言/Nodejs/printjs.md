## 开始学习Nodejs

#### 1.查看 node 版本

```go
node --version
v 14.16 .0
```

#### 2.打印输出  hello

1.新建一个 server.js 文件

```go

var foo ='hello'

console.log(foo)


输出 ： hello
```



#### 3.读取文件

```js
//读取 文件  
// fs 就是 file-system  缩写 
// 在Node中 如果要读取文件 就要引入  fs 模块
//在  fs  这个核心模块中  提供了 操作文件的 API 接口
//列如 读取 一个  文件  test.txt  fs.readFile
const fs =require('fs');

// 第一个 参数  是  文件路径 
// 第二个 参数  是  回调函数
fs.readFile('./test.txt',function(error,data){
    console.log(data.toString());
})


```

#### 4.http 请求

```js
//使用 http 模块 去构建一个 web 服务器 
//引入 http 

var https = require('https')


var server=https.createServer()

server.on('request',function(){
    console.log('收到客户端请求了')
})

server.listen(3001,function(){
    console.log("可以访问了")
})
```



#### 5.