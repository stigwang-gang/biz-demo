module github.com/stigwang-gang/biz-demo/gomall/app/user

go 1.23.4

replace (github.com/apache/thrift => github.com/apache/thrift v0.13.0
		github.com/stigwang-gang/biz-demo/gomall/rpc_gen => ../../rpc_gen
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.14.0 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
