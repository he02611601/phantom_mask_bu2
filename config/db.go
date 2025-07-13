package config

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Step 1: 建立資料庫
	sqlDB, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/")
	if err != nil {
		log.Fatalf("❌ 無法連線 MySQL: %v", err)
	}
	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS phantom_mask_bu2")
	if err != nil {
		log.Fatalf("❌ 建立資料庫失敗: %v", err)
	}

	// Step 2: 連線到該資料庫
	DB, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/phantom_mask_bu2?parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ GORM 無法連接資料庫: %v", err)
	}
}
