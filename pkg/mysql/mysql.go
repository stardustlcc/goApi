package mysql

import (
	"dwd-api/global"
	"dwd-api/pkg/setting"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDBEngine(databaseConfig *setting.DatabaseSetting) (*gorm.DB, error) {
	str := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local"
	db, err := gorm.Open(databaseConfig.DBType, fmt.Sprintf(str,
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.DBName,
		databaseConfig.Charset,
		databaseConfig.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxOpenConns(databaseConfig.MaxOpenConns)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.DB().SetMaxIdleConns(databaseConfig.MaxIdleConns)

	// 设置最大连接超时
	db.DB().SetConnMaxLifetime(time.Minute * databaseConfig.MaxLifeTime)

	//注册回调行为
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	return db, nil
}
