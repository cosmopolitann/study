## 测试用例



1.filepathJoin

```go
var initDocPaths = []string{
	filepath.Join("init-doc", "about"),
	filepath.Join("init-doc", "readme"),
	filepath.Join("init-doc", "help"),
	filepath.Join("init-doc", "contact"),
	filepath.Join("init-doc", "security-notes"),
	filepath.Join("init-doc", "quick-start"),
	filepath.Join("init-doc", "ping"),
}

fmt.Println("initDocPaths:=", initDocPaths)

结果：
initDocPaths:= [init-doc/about init-doc/readme init-doc/help init-doc/contact init-doc/security-notes init-doc/quick-start init-doc/ping]

相当于 拼接起来了参数

```

