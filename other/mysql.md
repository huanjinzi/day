# MySQL Master-Slave

## MySQL的安装
1.下载镜像
```
docker pull mysql:5.7
```

2.创建容器
```
docker run -p 3339:3306 --name mysql-master -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7
```

1. `-p 3339:3306`：将容器的`3306`端口映射到宿主的`3339`端口
2. `-e MYSQL_ROOT_PASSWORD=123456`：设置`root`用户的默认密码
3. `-d mysql:5.7`：指定镜像

3.查看所有的`container`
```
docker container ls -a
```

## Master配置
1.进入容器的`bash`终端
```
docker exec -it a19c5678dc95 bash
```

2.编辑`/etc/mysql/my.cnf`
```
[mysqld]
## 同一局域网内注意要唯一
server-id=1
## 开启二进制日志功能，可以随便取（关键）
log-bin=mysql-bin
```

3.重启
```
service mysql restart
docker start mysql-master
```

4.配置`master`的数据库
```
CREATE USER 'slave'@'%' IDENTIFIED BY '123456';
GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'slave'@'%';
```
查看状态`show master status;`

```
+------------------+----------+--------------+------------------+-------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+------------------+----------+--------------+------------------+-------------------+
| mysql-bin.000001 |      433 |              |                  |                   |
+------------------+----------+--------------+------------------+-------------------+
```

## Slave配置
1.配置文件
```
[mysqld]
## 设置server_id,注意要唯一
server-id=101  
## 开启二进制日志功能，以备Slave作为其它Slave的Master时使用
log-bin=mysql-slave-bin   
## relay_log配置中继日志
relay_log=edu-mysql-relay-bin 
```

2.进入`mysql`，执行下面的命令：
```
change master to master_host='172.17.0.2',master_user='slave',master_password='root',master_port=3306,master_log_file='mysql-bin.000001',master_log_pos= 433,master_connect_retry=30;
```

`master_log_pos`在上面的`show master status;`命令中可以查看。

3.查看`Slave`状态
```
show slave status \G;
```
```
*************************** 1. row ***************************
               Slave_IO_State: 
                  Master_Host: 172.17.0.2
                  Master_User: slave
                  Master_Port: 3306
                Connect_Retry: 30
              Master_Log_File: mysql-bin.000001
          Read_Master_Log_Pos: 433
               Relay_Log_File: edu-mysql-relay-bin.000002
                Relay_Log_Pos: 283
        Relay_Master_Log_File: mysql-bin.000001
             Slave_IO_Running: No
            Slave_SQL_Running: No
              Replicate_Do_DB: 
          Replicate_Ignore_DB: 
           Replicate_Do_Table: 
       Replicate_Ignore_Table: 
      Replicate_Wild_Do_Table: 
  Replicate_Wild_Ignore_Table: 
                   Last_Errno: 0
                   Last_Error: 
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 433
              Relay_Log_Space: 460
              Until_Condition: None
               Until_Log_File: 
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File: 
           Master_SSL_CA_Path: 
              Master_SSL_Cert: 
            Master_SSL_Cipher: 
               Master_SSL_Key: 
        Seconds_Behind_Master: NULL
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error: 
               Last_SQL_Errno: 0
               Last_SQL_Error: 
  Replicate_Ignore_Server_Ids: 
             Master_Server_Id: 1
                  Master_UUID: 1deb5735-722e-11e9-8b68-0242ac110002
             Master_Info_File: /var/lib/mysql/master.info
                    SQL_Delay: 0
          SQL_Remaining_Delay: NULL
      Slave_SQL_Running_State: 
           Master_Retry_Count: 86400
                  Master_Bind: 
      Last_IO_Error_Timestamp: 
     Last_SQL_Error_Timestamp: 
               Master_SSL_Crl: 
           Master_SSL_Crlpath: 
           Retrieved_Gtid_Set: 
            Executed_Gtid_Set: 
                Auto_Position: 0
1 row in set (0.00 sec)
```

4.启动`Master-Slave` Replication
```
start slave;
```

5.再次检查`Slave`状态
```
*************************** 1. row ***************************
               Slave_IO_State: Waiting for master to send event
                  Master_Host: 172.17.0.2
                  Master_User: slave
                  Master_Port: 3306
                Connect_Retry: 30
              Master_Log_File: mysql-bin.000001
          Read_Master_Log_Pos: 433
               Relay_Log_File: edu-mysql-relay-bin.000003
                Relay_Log_Pos: 283
        Relay_Master_Log_File: mysql-bin.000001
             Slave_IO_Running: Yes
            Slave_SQL_Running: Yes
              Replicate_Do_DB: 
          Replicate_Ignore_DB: 
           Replicate_Do_Table: 
       Replicate_Ignore_Table: 
      Replicate_Wild_Do_Table: 
  Replicate_Wild_Ignore_Table: 
                   Last_Errno: 0
                   Last_Error: 
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 433
              Relay_Log_Space: 623
              Until_Condition: None
               Until_Log_File: 
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File: 
           Master_SSL_CA_Path: 
              Master_SSL_Cert: 
            Master_SSL_Cipher: 
               Master_SSL_Key: 
        Seconds_Behind_Master: 0
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error: 
               Last_SQL_Errno: 0
               Last_SQL_Error: 
  Replicate_Ignore_Server_Ids: 
             Master_Server_Id: 1
                  Master_UUID: 1deb5735-722e-11e9-8b68-0242ac110002
             Master_Info_File: /var/lib/mysql/master.info
                    SQL_Delay: 0
          SQL_Remaining_Delay: NULL
      Slave_SQL_Running_State: Slave has read all relay log; waiting for the slave I/O thread to update it
           Master_Retry_Count: 86400
                  Master_Bind: 
      Last_IO_Error_Timestamp: 
     Last_SQL_Error_Timestamp: 
               Master_SSL_Crl: 
           Master_SSL_Crlpath: 
           Retrieved_Gtid_Set: 
            Executed_Gtid_Set: 
                Auto_Position: 0
1 row in set (0.00 sec)

```

可以看到，启动`Master-Slave` Replication 之后的改变：
```
Slave_IO_Running: Yes
Slave_SQL_Running: Yes
```
如果有错误发生，请查看下面的字段：
```
Last_IO_Errno: 0
Last_IO_Error: 
Last_SQL_Errno: 0
Last_SQL_Error: 
```

## mysql
```
sudo vim /etc/mysql/debian.cnf
mysql -udebian-sys-maint -p

mysql -u root -p
ALTER USER 'root'@'localhost' IDENTIFIED BY 'MyNewPass4!';

use mysql;
select host,user,authentication_string from user;
select user,plugin from user; // auth_socket
update user set authentication_string =password('root'),plugin='mysql_native_password' where user='root'; //将auth_socket改为msyql_native_password
sudo service mysql restart //以上修改需要重启生效
grant all privileges on appstore.* to 'appstore'@'%' identified by 'appstore'; // 创建appstore用户，并且分配权限
show grants for appstore; // 查看appstore的权限
show databases;
UPDATE user SET authentication_string=PASSWORD('root') where USER='root'; //修改密码


revoke insert on appstore.* from 'appstore'@'%'; //收回insert权限
flush privileges; //刷新权限

help contents //帮助文档

mysql -h 192.168.1.172 -P 3306 -u appstore -p

// 中文问题
sudo vim /etc/mysql/conf.d/mysql.cnf
　　
[mysql]
default-character-set=utf8
[mysqld]
character-set-server=utf8
show variables like '%char%' //查看数据库编码
show create database appstore; //查看数据库创建指令
show create table appstore_category; //查看数据表创建命令
alter database appstore character set utf8; //修改数据库编码
alter table appstore_category charset=utf8; //修改数据表编码
alter table appstore_category convert to character set utf8; //修改数据表编码
alter table appstore_category [column]... character set utf8;
delete from appstore_apk_res where apk_key like '%';//数据库删除

desc [table]; //查看表结构
show create table [table] //查看表创建命令
select * from cms_r_video_info where channel=0 order by id asc; //查询结果排序
UPDATE runoob_tbl SET runoob_title='学习 C++' WHERE runoob_id=3;

// 一定要注意，这里是utf8，不是utf-8

drop database xxxx; //删除数据库
drop table xxxx; //删除数据表格

set password for 用户名@localhost = password('新密码');
mysqladmin -uroot -p123456 password 123

//远程连接问题
sudo vim  /etc/mysql/mysql.conf.d/mysqld.cnf
// #bind-address		= 127.0.0.1 //注释
SELECT DISTINCT CONCAT('User: ''',user,'''@''',host,''';') AS query FROM mysql.user;
GRANT ALL PRIVILEGES ON *.* TO 'username'@'192.168.10.83' IDENTIFIED BY 'password' WITH GRANT OPTION;
flush privileges;
show global variables like 'port'; // 常看端口号
CREATE DATABASE appstore; // 创建数据库

show create table appstore_app_info;
alter table appstore_app_info default character set utf8;

//修改mysql默认端口
vi /etc/my.cnf
port=3306

mysqldump -uroot -proot -P3306 -A > all.sql
mysqldump -uroot -proot -P3306 -A -t > all_table.sql
mysqldump -uroot -proot -P3306 -A -d > all_data.sql
mysqldump -uroot -proot -P3306 sakila  > sakila.sql
mysqldump -uroot -proot -P3306 sakila -t > sakila_table.sql
mysqldump -uroot -proot -P3306 sakila -d > sakila_data.sql
mysqldump -uroot -proot--databases db1 db2 > db1_db2.sql

mysqladmin -uroot -p123456 create sakila 
mysql -uroot -proot  sakila < sakila.sql
```









