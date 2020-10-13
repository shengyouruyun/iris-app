package database

import (
	"github.com/iris-app/model"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func InitDb() *xorm.Engine {

	orm, err := xorm.NewEngine("mysql", "toby:ws520360@/mydb?charset=utf8")

	if err != nil {

		log.Fatal("数据库连接失败:", err)

		fmt.Println(" failed to connect to database")

	}


	fmt.Println("数据库连接成功")

	return orm

}


func InitProducDB() *xorm.Engine {

	db := InitDb()

	db.Sync2(new(model.Product))

	return db

}

