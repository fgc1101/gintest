package test

import (
	"fmt"
	"gintest/models/trip"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

func TestGormTest(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/jyoa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "as_",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	data := make([]*trip.Trip, 0)
	err = db.Table("as_market_trip").Limit(10).Find(&data).Error
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, v := range data {
		fmt.Printf("trip=> %v\n", v)
	}
}
