package dal

import (
	"github.com/stigwang-gang/biz-demo/gomall/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
