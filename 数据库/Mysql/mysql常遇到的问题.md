## Mysql 常遇到的问题

#### 1.数据库字符集不支持插入中文，设置utf8字符集就可以了

```mysql
解决方式：
找到mysqld.cnf文件
###在[mysqld]最下面加入下面几句话
  default-storage-engine=INNODB  
  character-set-server=utf8 
  collation-server=utf8_general_ci
  
  
重启mysql，输入show VARIABLES like 'character%';查询结果
```



2.

