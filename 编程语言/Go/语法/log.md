# log日志

```go
package log

import (
   "io"
   "log"
   "os"
)

var (
   Info    *log.Logger
   Warning *log.Logger
   Error   *log.Logger
)

func Initlog() {
   errFile, err := os.OpenFile("./log/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
   if err != nil {
      log.Fatalln("打开日志文件失败：", err)
   }
   infoFile, err := os.OpenFile("./log/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
   if err != nil {
      log.Fatalln("打开日志文件失败2：", err)
   }

   Info = log.New(io.MultiWriter(os.Stderr, infoFile), "Info", log.Ldate|log.Ltime|log.Lshortfile)
   Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
   Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
   //Warning.Println("Warning")
   //Error.Println("Error")
}
```



创建 日志log 包  文件



使用  log.Info.println

​		 









