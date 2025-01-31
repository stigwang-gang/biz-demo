package dal

import (
	"github.com/stigwang-gang/biz-demo/gomall/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
