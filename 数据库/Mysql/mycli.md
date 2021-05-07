### Mycli 工具



1.快速方便的连接mysql 数据库，提示输入。

```go
地址：
https://github.com/dbcli/mycli
```





##### 安装：

```go
$ pip install -U mycli


$ brew update && brew install mycli  # Only on macOS


$ sudo apt-get install mycli # Only on debian or ubuntu
```



注意：

可以用 pip3 安装,安装的时候 是找不到   mycli 的 执行文件的

要将  mycli  移动到  /usr/local/bin 下面



前提是 要找到 mycli    路径 可以使用pip3 查找



```go
pip3 show mycli


就可以找到 安装路径了  
然后 往上一层 返回  找到 python  的  bin目录   mycli 就在这里 然后 
sudo cp mycli /usr/local/bin

```



连接 mycli

```go
可以使用 3种 方式去连接 mycli
查看  mycli --help 
 
第一种方式  
mycli mysql://root@127.0.0.1:3306/test
输入之后 会提示输入密码 然后就可以连接上了


```



