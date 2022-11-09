package test

import (
	"go_project/src/IM_telegram/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGorm(t *testing.T) {
	db, err := gorm.Open(mysql.Open("pytest:123456@tcp(127.0.0.1:3306)/gin_chat?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//迁移 schema
	db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(models.Message{})
	db.AutoMigrate(models.Contact{})
	db.AutoMigrate(models.GroupBasic{})

	// Create
	//oneUser := &models.UserBasic{
	//	Name:         "刘一豆",
	//	LoginTime:    time.Now(),
	//	LoginOutTime: time.Now(),
	//	HeatBeatTime: time.Now(),
	//}
	//db.Create(oneUser)

	// Read
	//var product Product
	//db.First(oneUser, 1) // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	//db.Model(oneUser).Update("Password", "14523")
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
