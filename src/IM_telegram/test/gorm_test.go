package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Community struct {
	gorm.Model
	Name    string
	OwnerId uint
	Img     string
	Desc    string
}

func TestGorm(t *testing.T) {
	db, err := gorm.Open(mysql.Open("pytest:123456@tcp(127.0.0.1:3306)/gin_chat?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//迁移 schema
	//err = db.AutoMigrate(Community{})
	//if err != nil {
	//	t.Log(err)
	//}

	// Create
	//Comms:=[]Community{
	//	{Name: "okkil"},
	//	{Name: "Yossig"},
	//	{Name: "oossig"},
	//}
	//db.Create(&Comms)
	//for _,comm :=range Comms{
	//	fmt.Println(comm.ID)
	//}
	//db.Create(oneComm)
	//db.Select("Name").Create(oneComm)
	//db.Omit("OwnerId","ID").Create(oneComm)
	//db.Model(&Community{}).Create([]map[string]interface{}{
	//	{"Name": "Okkil"},
	//	{"Name": "yossig"},
	//	{"Name": "oosig"},
	//})

	// Read
	//conn:=&Community{Name: "aaaaa"}
	//db.First(conn)
	//fmt.Println(conn)
	//conn=&Community{Name: "aaaaa"}
	//db.Last(conn)
	//fmt.Println(conn)
	//conn=&Community{Name: "aaaaa"}
	//db.Take(conn)
	//fmt.Println(conn)
	var conn []Community
	db.Where("name = ?", "aaaaa").Find(&conn)
	fmt.Println(conn)

	// Update - 将 product 的 price 更新为 200
	//db.Model(oneUser).Update("Password", "14523")
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
