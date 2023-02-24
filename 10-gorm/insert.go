package main

import (
	"gorm.io/gorm/clause"
	"log"
	"time"
)

func insertWithType() {
	user := &User{
		Name:     "小明",
		Age:      12,
		Birthday: time.Now(),
	}
	tx := db.Create(user)
	log.Println(tx.Error, tx.RowsAffected, user.ID)

	user.Name = "jinzhu"
	user.Age = 18
	//创建记录并更新给出的字段
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
	db.Select("Name", "Age", "CreatedAt").Create(&user)

	user.Birthday = time.Now().Add(-time.Hour * 2)
	//ID 主键设置为0，insert的时候会重新分配自增id
	user.ID = 0
	//创建一个记录且一同忽略传递给略去的字段值。
	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
	db.Omit("Name", "Age", "CreatedAt").Create(&user)

	//批量insert ，CreateInBatches(users, 100) 可以分批指定数量
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

}

// GORM 支持根据 map[string]interface{}
// 和 []map[string]interface{}{} 创建记录
// 根据 map 创建记录时，association、hook不会被调用，且主键也不会自动填充??
func insertWithMap() {
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})
	// batch insert from `[]map[string]interface{}{}`
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_11", "Age": 181},
		{"Name": "jinzhu_21", "Age": 201},
	})
}

// 使用 SQL 表达式、Context Valuer 创建记录
func insertWithSQl() {
	// 通过 map 创建记录
	db.Model(User{}).Create(map[string]interface{}{
		"Name":     "jinzhu",
		"Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
	})
	// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"));

	db.Create(&User{
		Name:     "jinzhu",
		Location: Location{X: 100, Y: 100},
	})
	// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"))

}
