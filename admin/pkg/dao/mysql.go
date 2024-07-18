package dao

import (
	"fmt"
	"gjob-admin/pkg/model"
	"gjob-admin/pkg/utils"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {
	db, err := gorm.Open(mysql.Open(utils.Config.GetString("mysql.base.dsn")), &gorm.Config{})
	if err != nil {
		fmt.Printf("Init database error,%v \n", err)
		os.Exit(1)
	}
	DB = db
	db.AutoMigrate(&model.SqlJob{})
}
