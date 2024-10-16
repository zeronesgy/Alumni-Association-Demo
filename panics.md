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
