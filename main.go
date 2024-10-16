package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20); not null"`
	Telephone string `gorm:"type:varchar(110); not null; unique"`
	Password  string `gorm:"type:size(255); not null"`
}

func main() {
	db := InitDB()

	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		// 数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}

		// 如果名称没有传，则随机生成一个10位字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		// 判断条件
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已存在"})
			return
		}

		// 创建成功
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)
		// 返回结果

		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run(":8080"))
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

func RandomString(n int) string {
	var letters = []byte("ASDFGHJKLasdfghjkl")
	result := make([]byte, n)
	rand.New(rand.NewSource(time.Now().UnixNano())).Read(result)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "aad"
	user := "root"
	password := "root"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, host, port, database, charset)

	db, err := gorm.Open(mysql.New(mysql.Config{DriverName: driverName, DSN: args}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}
