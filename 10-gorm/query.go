package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// more:  https://gorm.io/zh_CN/docs/query.html
func find() {
	var user User
	// 获取第一条记录（主键升序）
	db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	result := db.First(&user)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(result.Error)        // returns error or nil

	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)

}
