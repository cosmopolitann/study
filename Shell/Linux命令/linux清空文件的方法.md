# Linux文件清空的几种方法

### 备注一下  个人觉得 好用的 方法是 第 2 个 和 第4个.



## 1、使用重定向的方法

```go
[root@centos7 ~]# du -h test.txt 
4.0K    test.txt
[root@centos7 ~]# > test.txt 
[root@centos7 ~]# du -h test.txt 
0    test.txt

//这里需要注意的是   当使用 > test.txt 的时候   如果 不按 control z 退出的话   在下面

输入的内容 就存入 test。txt 文件中了  
```

## 2、使用true命令重定向清空文件

```
[root@centos7 ~]# du -h test.txt 
4.0K    test.txt
[root@centos7 ~]# true > test.txt 
[root@centos7 ~]# du -h test.txt 
0    test.txt
```

## 3、使用cat/cp/dd命令及/dev/null设备来清空文件

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
[root@centos7 ~]# du -h test.txt 
4.0K    test.txt
[root@centos7 ~]# cat /dev/null >  test.txt 
[root@centos7 ~]# du -h test.txt 
0    test.txt
###################################################
[root@centos7 ~]# echo "Hello World" > test.txt 
[root@centos7 ~]# du -h test.txt 
4.0K    test.txt
[root@centos7 ~]# cp /dev/null test.txt 
cp：是否覆盖"test.txt"？ y
[root@centos7 ~]# du -h test.txt 
0    test.txt
##################################################
[root@centos7 ~]# echo "Hello World" > test.txt 
[root@centos7 ~]# du -h test.txt 
4.0K    test.txt
[root@centos7 ~]# dd if=/dev/null of=test.txt 
记录了0+0 的读入
记录了0+0 的写出
0字节(0 B)已复制，0.00041594 秒，0.0 kB/秒
[root@centos7 ~]# du -h test.txt 
0    test.txt
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

## 4、使用echo命令清空文件

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
[root@centos7 ~]# echo "Hello World" > test.txt 
[root@centos7 ~]# du -h test.txt 
4.0K    test.txt
[root@centos7 ~]# echo -n "" > test.txt    #要加上"-n"参数，默认情况下是"\n"，就是回车符
[root@centos7 ~]# du -h test.txt  
0    test.txt
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

## 5、使用truncate命令清空文件

```
[root@centos7 ~]# du -h test.txt 
4.0K    test.txt
[root@centos7 ~]# truncate -s 0 test.txt   -s参数用来设定文件的大小，清空文件，就设定为0；
[root@centos7 ~]# du -h test.txt 
0    test.txt
```