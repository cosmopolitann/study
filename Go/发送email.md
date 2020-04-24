# 发送email邮件

## 简介

程序中时常有发送邮件的需求。有异常情况了需要通知管理员和负责人，用户下单后可能需要通知订单信息，电商平台、中国移动和联通都有每月账单，这些都可以通过邮件来推送。还有我们平时收到的垃圾邮件大都也是通过这种方式发送的。那么如何在 Go 语言发送邮件？本文我们介绍一下`email`库的使用。

## 快速使用

先安装库

```
$ go get github.com/jordan-wright/email
```

我们需要额外一些工作。我们知道邮箱使用`SMTP/POP3/IMAP`等协议从邮件服务器上拉取邮件。邮件并不是直接发送到邮箱的，而是邮箱请求拉取的。所以，我们需要配置`SMTP/POP3/IMAP`服务器。从头搭建固然可行，而且也有现成的开源库，但是比较麻烦。现在一般的邮箱服务商都开放了`SMTP/POP3/IMAP`服务器。我这里拿 126 邮箱来举例，使用`SMTP`服务器。当然，用 QQ 邮箱也可以。

- 首先，登录邮箱；
- 点开顶部的设置，选择`POP3/SMTP/IMAP`；
- 点击开启`IMAP/SMTP`服务，按照步骤开启即可，有个密码设置，记住这个密码，后面有用。

然后就可以编码了：

```
package main

import (
  "log"
  "net/smtp"

  "github.com/jordan-wright/email"
)

func main() {
  e := email.NewEmail()
  e.From = "<xxx@126.com>"
  e.To = []string{"	邮箱地址xxx @qq.com"}
  e.Subject = "Awesome web"
  e.Text = []byte("Text Body is, of course, supported!")
  err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "xxx@126.com", "yyy", "smtp.126.com"))
  if err != nil {
    log.Fatal(err)
  }
}
```

代码中`xxx`替换成你的邮箱账号，`yyy`替换成上面设置的密码。

代码步骤比较简单清晰：

- 先调用`NewEmail`创建一封邮件；
- 设置`From`发送方，`To`接收者，`Subject`邮件主题（标题），`Text`设置邮件内容；
- 然后调用`Send`发送，参数1是 SMTP 服务器的地址，参数2为验证信息。

运行程序将会向我的 QQ 邮箱发送一封邮件：

有的邮箱会把这种邮件放在垃圾箱中，例如 QQ,如果收件箱找不到，记得到垃圾箱瞅瞅。

平常我们发邮件的时候可能会抄送给一些人，还有一些人要秘密抄送，即 CC（Carbon Copy）和 BCC （Blind Carbon Copy）。`email`我们也可以设置这两个参数：

```go
package main

import (
  "log"
  "net/smtp"

  "github.com/jordan-wright/email"
)

func main() {
  e := email.NewEmail()
  e.From = "dj <xxx@126.com>"
  e.To = []string{"xxx@qq.com"}
  e.Cc = []string{"test1@126.com", "test2@126.com"}
  e.Bcc = []string{"xxx@126.com"}
  e.Subject = "Test"
  e.Text = []byte("Text Body is, of course, supported!")
  err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "xxx@126.com", "yyy", "smtp.126.com"))
  if err != nil {
    log.Fatal(err)
  }
}
```



### 这个是我自己测试的代码

```go
package main

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	e.From = "winter <jiang_yunzhen@163.com>"
	e.To = []string{"xxx@qq.com"}
	e.Subject = "Test Email"
	e.Text = []byte("Text Body is, of course, supported!")
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "jiang_yunzhen@163.com", "CIRYEDMWRIVMKLDL", "smtp.163.com"))
	if err != nil {
		log.Fatal(err)
	}
}

```

