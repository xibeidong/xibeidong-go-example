package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

var db *gorm.DB

type User struct {
	ID        int64
	Name      string
	Age       uint8
	Birthday  time.Time
	Location  Location  //自定义
	CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
	UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
}

// Location 通过自定义类型创建记录
type Location struct {
	X, Y int
}

// Scan 方法实现了 sql.Scanner 接口
func (loc *Location) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	return nil
}

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
	}
}

// AfterSave Hook BeforeSave, BeforeCreate, AfterSave, AfterCreate
func (u *User) AfterSave(tx *gorm.DB) (err error) {

	fmt.Println("AfterSave", u.ID)
	return
}

func main() {
	if !check("test2") {
		panic("db is not exists and create fail!")
	}
	ReadyDB()
	createTable()
	insertWithType()
	insertWithMap()
	insertWithSQl()

}
