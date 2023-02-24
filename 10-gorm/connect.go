package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 不存在databaseName就创建
func check(databaseName string) bool {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("create database if not exists " + databaseName + " default character set utf8 default collate utf8_general_ci")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("use " + databaseName)
	if err != nil {
		panic(err)
	}
	fmt.Println("use " + databaseName + " ok ")
	return true
}
func createTable() {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}
	// 检测User结构体对应的表是否存在
	fmt.Println(db.Migrator().HasTable(&User{}))

	// 检测表名users是否存在
	fmt.Println(db.Migrator().HasTable("users"))
}
func ReadyDB() {
	//想要正确的处理 time.Time ，您需要带上 parseTime 参数，
	//要支持完整的 UTF-8 编码 您需要将 charset=utf8 更改为 charset=utf8mb4
	dsn := "root:123456@tcp(127.0.0.1:3306)/test2?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// GORM 使用 database/sql 维护连接池
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("db is ready")
}
