package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"three_girls/model"
)

var db *gorm.DB

func InitDatabase() {
	db, err := gorm.Open("mysql", "root:123@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("error occurred, err is ", err)
		panic(err)
	}
	CreateTable()
	// InsertUserInfo()
	defer db.Close()
}

//if table doesn't exist, create it
func CreateTable() {
	db.AutoMigrate(&model.UserInfo{})
}

func InsertUserInfo() {
	var userInfo model.UserInfo = model.UserInfo{
		PhoneNum: "15878487894",

		UserToken: "fsdf1235fds1321",

		UserId: 1,

		NickName: "Jack",

		City: "Beijing",

		Preference: model.LOLITA,

		LoginTimes: 123,

		IdolNum: 999,

		FansNum: 100000,
	}

	db.Create(&userInfo)
}
