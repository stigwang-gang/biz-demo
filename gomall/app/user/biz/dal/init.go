package dal

import (
	"github.com/stigwang-gang/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/stigwang-gang/biz-demo/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
