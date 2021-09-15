//interface是只要你实现了接口定义的类型（方法），就可以作为接口的实例拿来用，在语言层面上不再需要其它的约束。

//在结构体中嵌入匿名类型成员，就可以继承匿名类型的方法。
//
//不仅可以嵌入匿名类型成员，还可以嵌入匿名的接口，或者匿名的指针对象，是一种纯虚继承，继承的只是接口的规范，而实现还得靠自己。

struct 成员小写导致rpc的结果全为默认值

二维切片中反复添加同一个切片，修改一个切片的值会影响二维切片中的所有值

安装go包：可以直接在程序中引用，然后运行使用go mod（GO111MODULE=on）

# mysql
1. mysql 只能root登录
   alter user 'root'@'localhost' identified with mysql_native_password by '123456';
2. mysql 修改密码 
3. mysql 创建用户
4. 