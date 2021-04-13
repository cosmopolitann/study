# 解决Navicat连接Mysql问题

###### 1.error 2059: Authentication plugin ‘caching_sha2_password’ cannot be loaded”错误。

###### 问题解决方法



#### 我连接不上的解决方法

直接把最后的那几个命令直接 复制粘贴在 登陆之后的mysql 数据库中执行  然后进行连接 就可以了。



## 问题描述

在 VMware 虚拟机的客户机 CentOS7 里面安装运行有 Docker 的 MySQL 8.0，由于当前 CentOS7 默认的 MySQL 客户端版本太低（5.5.60）（低版本的客户端认 mysql_native_password 认证插件，而高版本认 caching_sha2_password 插件） ，导致连接服务器时出现以下的错误：





| 12   | error 2059: Authentication plugin 'caching_sha2_password' cannot be loaded: /usr/lib64/mysql/plugin/caching_sha2_password.so: cannot open shared **object** file: No such file **or** directory |
| ---- | ------------------------------------------------------------ |
|      |                                                              |



## 解决办法

1. 用高版本的 MySQL，或者进入该 Docker 容器，登录 MySQL 服务器。

2. 执行 MySQL shell 命令查看服务器的版本：

   

   

   | 12   | > select version(); |
   | ---- | ------------------- |
   |      |                     |

   

   执行结果：

   

   

   | 1234567 | +-----------+\| version() \|+-----------+\| 8.0.16    \|+-----------+1 row **in** set (0.00 sec) |
   | ------- | ------------------------------------------------------------ |
   |         |                                                              |

   

3. 查看当前默认的密码认证插件：

   

   

   | 12   | > show variables like 'default_authentication_plugin'; |
   | ---- | ------------------------------------------------------ |
   |      |                                                        |

   

   ```go
   
   > show variables like 'default_authentication_plugin';
   
   
   > show variables like 'default_authentication_plugin';
    
   ```

   

   ```go
   
   +-------------------------------+-----------------------+
   | Variable_name                 | Value                 |
   +-------------------------------+-----------------------+
   | default_authentication_plugin | caching_sha2_password |
   +-------------------------------+-----------------------+
   1 row in set (0.01 sec)
    
   ```

   

   

   查看当前所有用户绑定的认证插件：

   

   ```go
   	> select host,user,plugin from mysql.user;
   ```

   ```go
   +-----------+------------------+-----------------------+
   | host      | user             | plugin                |
   +-----------+------------------+-----------------------+
   | %         | root             | caching_sha2_password |
   | localhost | healthchecker    | caching_sha2_password |
   | localhost | mysql.infoschema | caching_sha2_password |
   | localhost | mysql.session    | caching_sha2_password |
   | localhost | mysql.sys        | caching_sha2_password |
   +-----------+------------------+-----------------------+
   5 rows in set (0.00 sec)
   ```

   

4. 假如想更改 root 用户的认证方式

   ```go
   
   # 修改加密规则
   > ALTER USER 'root'@'%' IDENTIFIED BY 'root' PASSWORD EXPIRE NEVER;
   # 更新用户密码
   > ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
   # 赋予 root 用户最高权限
   > grant all privileges on *.* to root@'%' with grant option;
   # 刷新权限
   > flush privileges;
   ```

   完成解决方案。