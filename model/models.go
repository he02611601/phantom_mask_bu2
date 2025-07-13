package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// 藥局、結構
type Pharmacy struct {
	gorm.Model
	Name         string             // 藥局名稱
	CashBalance  float64            // 藥局現金餘額
	OpeningHours string             // 營業時間（原始資料）
	OpenHours    []PharmacyOpenHour // 營業時間資料
	Masks        []Mask             // 商品資料
}

// 藥局營業時間結構
type PharmacyOpenHour struct {
	gorm.Model
	PharmacyID uint           // 藥局ID
	DayOfWeek  int            // 營業星期
	StartTime  datatypes.Time // 營業開始時間
	EndTime    datatypes.Time // 營業結束時間
}

// 口罩結構
type Mask struct {
	gorm.Model
	PharmacyID uint    // 藥局ID
	Name       string  // 口罩名稱
	Price      float64 // 價格
	Stock      int
}

// 用戶結構
type User struct {
	gorm.Model
	Name        string     // 用戶名稱
	CashBalance float64    // 用戶現金餘額
	Purchases   []Purchase // 用戶訂單資料
}

// 用戶訂單結構
type Purchase struct {
	gorm.Model
	UserID       uint      // 用戶ID
	PharmacyName string    // 藥局ID
	MaskName     string    // 口罩名稱
	Amount       float64   // 訂單總金額
	Quantity     int       // 訂單口罩數量
	Date         time.Time // 購買日期
}
