package main

import (
	"fmt"
	"go_im/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3307)/ginchat?charset=utf8mb4&parseTime=True&loc=local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	db.Create(&models.UserBasic{UserName: "user1", PassWord: "111"})

	// Read
	var user1 models.UserBasic
	fmt.Printf("db.First(&user1, 1): %v\n", db.First(&user1, 1)) // 根据整型主键查找
	// db.First(&user1, "UserName = ?", "user1") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&user1).Update("Phone", "133xxxx3333")
	// // Update - 更新多个字段
	// db.Model(&user1).Updates(models.UserBasic{UserName: "user2", PassWord: "F42"}) // 仅更新非零值字段
	// db.Model(&user1).Updates(map[string]interface{}{"UserName": "user3", "PassWord": "F43"})

	// Delete - 删除 product
	db.Delete(&user1, 1)
}
