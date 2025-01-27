package dal

import (
	"github.com/stigwang-gang/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
