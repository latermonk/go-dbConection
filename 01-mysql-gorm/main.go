package main

import (
	"fmt"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var (
	DB *gorm.DB
)


func main() {
	fmt.Println("==============")

	//db, err := gorm.Open("mysql", "root:root1234@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	//defer db.Close()

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root1234@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//fmt.Printf("%s", db)
	fmt.Println("%s", db)

	//err = db.Ping()
	if err != nil {
		//return 0
		fmt.Println("===== Failure failure !!!======")
	}else {
		fmt.Println("========Ok , It's ok ======================")
	}
	//return nil


}
