package main

import (
	"github.com/joho/godotenv"
	"github.com/stigwang-gang/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/stigwang-gang/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/stigwang-gang/biz-demo/gomall/demo/demo_proto/biz/model"
)

func main() {
	//env get
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	//CURD
	mysql.DB.Create(&model.User{Email: "test@gmail.com", Password: "123456"})
	//
	//mysql.DB.Model(&model.User{}).Where("email=?", "test@gmail.com").
	//	Update("password", "654321")
	//var row model.User
	//mysql.DB.Model(&model.User{}).Where("email=?", "test@gmail.com").First(&row)
	//fmt.Println(row)

	//mysql.DB.Where("email=?", "test@gmail.com").Delete(&model.User{})
	//mysql.DB.Unscoped().Where("email=?", "test@gmail.com").Delete(&model.User{})

}
