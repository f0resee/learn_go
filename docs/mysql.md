## MySql
0. mariadb安装及设置
### 注意mysql和mariadb使用配置文件目录不同
### 主从备份需要手动创建数据库及数据表。
sudo apt-get install mariadb-server
切换root用户
使用mysql_secure_installation命令设置密码
设置新用户，并为其授权远程登录：
修改sudo vim /etc/mysql/mariadb.conf.d/50-server.cnf
重启服务  sudo systemctl restart mariadb.service
1. mysql 只能使用root登录
   alter user 'root'@'localhost' identified with mysql_native_password by '123456';
2. mysql 修改密码 
mysql -u root update mysql.user set authentication_string=PASSWORD('123456') where User='root'; flush privileges;
   set password for root@localhost = password('123456'); flush privileges;
3. mysql 创建用户并授权某个数据库:
```sql
CREATE USER 'test'@'localhost' IDENTIFIED BY '123456';
update user set host='%' where user='test';
grant all privileges on test.* to 'test'@'%'; --（两次）

```

4. MySql执行sql文件，a. source xxx.sql  b. mysql -u用户名 -p用户密码 < xxx.sql 
   