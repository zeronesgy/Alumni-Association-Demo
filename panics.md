# panics and questions

#### 引用import ("github.com/gin-gonic/gin""github.com/jinzhu/gorm")
    最后还是引用官方文档的库（"gorm.io/driver/mysql" "gorm.io/gorm"）
#### 如何在Go语言中使用NewSeed()替代已弃用的rand.Seed(SEED)方法
    rand.New(rand.NewSource(time.Now().UnixNano())).Read(result)
#### net start mysql 发生系统错误 5。  拒绝访问。
    cmd管理员身份启用
#### panic: Error 1146: Table 'mysql.users' doesn't exist
    如今的gorm已经没有了Close方法，删去
    如今使用Open方法的写法：
    db, err := gorm.Open(mysql.New(mysql.Config{DriverName: driverName, DSN: args}), &gorm.Config{})
### gorm:
    sql语句，对于type size unique等，别写错了
#### gorm log信息 record not found
    [0.854ms] [rows:0] SELECT * FROM `users` WHERE telephone = '12345678913' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
    gorm中logger.go中修改 IgnoreRecordNotFoundError: true （忽略记录未找到错误）

#### yarn
    yarn安装完后执行yarn的命令遇到问题：yarn : 无法加载文件…因为在此系统上禁止运行脚本。有关详细信息，请参阅 https:/go.microsoft.com/fwlink/?LinkID=135170 中的 about_Execution_Policies。
    解决方法：如下
    1.首先在windows搜索windows PowerSell，然后以管理员身份运行ISE
    2.执行命令：set-ExecutionPolicy RemoteSigned，没有报错就说明ok了，之后就可以正常运行yarn的命令