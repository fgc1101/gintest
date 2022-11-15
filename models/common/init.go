package common

import (
	"fmt"
	"gintest/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB = Init()

func Init() *gorm.DB {
	_config := common.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		_config.DataSource.Username,
		_config.DataSource.Password,
		_config.DataSource.Host,
		_config.DataSource.Port,
		_config.DataSource.Database,
		_config.DataSource.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "as_",
		},
	})

	if err != nil {
		fmt.Println("GORM INIT ERROR:", err)
		return nil
	}

	return db

}
