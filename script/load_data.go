package main

import (
	"encoding/json"
	"fmt"
	"os"
	"phantom_mask_bu2/config"
	"phantom_mask_bu2/model"
	"strings"
	"time"

	"gorm.io/datatypes"
)

// 讀入藥局相關資料的原始結構
type MaskRaw struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stockQuantity"`
}

type PharmacyRaw struct {
	Name         string    `json:"name"`
	CashBalance  float64   `json:"cashBalance"`
	OpeningHours string    `json:"openingHours"`
	Masks        []MaskRaw `json:"masks"`
}

// 讀入用戶相關資料的原始結構
type PurchasesRaw struct {
	PharmacyName        string  `json:"pharmacyName"`
	MaskName            string  `json:"maskName"`
	TransactionAmount   float64 `json:"transactionAmount"`
	TransactionQuantity int     `json:"transactionQuantity"`
	TransactionDatetime string  `json:"transactionDatetime"`
}
type UserRaw struct {
	Name        string         `json:"name"`
	CashBalance float64        `json:"cashBalance"`
	Purchases   []PurchasesRaw `json:"purchaseHistories"`
}

// 星期轉換成整數資料
var DayOfWeek map[string]int

func init() {
	DayOfWeek = map[string]int{
		"Mon":  1,
		"Tue":  2,
		"Wed":  3,
		"Thur": 4,
		"Fri":  5,
		"Sat":  6,
		"Sun":  7,
	}
}

func parseOpeningHours(raw string) []model.PharmacyOpenHour {
	var hours []model.PharmacyOpenHour
	parts := strings.Split(raw, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		tokens := strings.Split(part, " ")
		if len(tokens) < 4 {
			continue
		}

		dayStr := tokens[0]
		startStr := tokens[1]
		endStr := tokens[3]
		if endStr == "24:00" {
			endStr = "00:00"
		}

		startTime, err1 := time.Parse("15:04", startStr)
		endTime, err2 := time.Parse("15:04", endStr)
		if err1 != nil || err2 != nil {
			continue
		}

		if startTime.Equal(endTime) && startTime.Hour() == 0 && startTime.Minute() == 0 {
			hours = append(hours, model.PharmacyOpenHour{
				DayOfWeek: DayOfWeek[dayStr],
				StartTime: datatypes.NewTime(startTime.Hour(), startTime.Minute(), startTime.Second(), startTime.Nanosecond()),
				EndTime:   datatypes.NewTime(23, 59, 59, 0),
			})
		} else if !startTime.Before(endTime) {
			hours = append(hours, model.PharmacyOpenHour{
				DayOfWeek: DayOfWeek[dayStr],
				StartTime: datatypes.NewTime(startTime.Hour(), startTime.Minute(), startTime.Second(), startTime.Nanosecond()),
				EndTime:   datatypes.NewTime(23, 59, 59, 0),
			})
			if endTime.Hour() != 0 || endTime.Minute() != 0 {
				hours = append(hours, model.PharmacyOpenHour{
					DayOfWeek: DayOfWeek[dayStr]%7 + 1,
					StartTime: datatypes.NewTime(0, 0, 0, 0),
					EndTime:   datatypes.NewTime(endTime.Hour(), endTime.Minute(), endTime.Second(), endTime.Nanosecond()),
				})
			}
		} else {
			hours = append(hours, model.PharmacyOpenHour{
				DayOfWeek: DayOfWeek[dayStr],
				StartTime: datatypes.NewTime(startTime.Hour(), startTime.Minute(), startTime.Second(), startTime.Nanosecond()),
				EndTime:   datatypes.NewTime(endTime.Hour(), endTime.Minute(), endTime.Second(), endTime.Nanosecond()),
			})
		}

	}

	return hours
}

// 寫入藥局相關資料
func loadPharmacies() {
	fmt.Println("匯入藥局資料開始")
	f, err := os.Open("data/pharmacies.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var rawPharmacies []PharmacyRaw
	if err := json.NewDecoder(f).Decode(&rawPharmacies); err != nil {
		panic(err)
	}

	for _, raw := range rawPharmacies {
		p := model.Pharmacy{
			Name:         raw.Name,
			CashBalance:  raw.CashBalance,
			OpeningHours: raw.OpeningHours,
		}

		config.DB.Create(&p)

		for _, hour := range parseOpeningHours(raw.OpeningHours) {
			hour.PharmacyID = p.ID
			config.DB.Create(&hour)
		}

		for _, mask := range raw.Masks {
			config.DB.Create(&model.Mask{
				PharmacyID: p.ID,
				Name:       mask.Name,
				Price:      mask.Price,
				Stock:      mask.Stock,
			})
		}
	}
	fmt.Println("匯入藥局資料結束")
}

// 寫入用戶相關資料
func loadUsers() {
	fmt.Println("匯入用戶資料開始")
	f, err := os.Open("data/users.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var rawUsers []UserRaw
	if err := json.NewDecoder(f).Decode(&rawUsers); err != nil {
		panic(err)
	}

	for _, user := range rawUsers {
		u := model.User{
			Name:        user.Name,
			CashBalance: user.CashBalance,
		}

		config.DB.Create(&u)

		for _, purchase := range user.Purchases {
			date, err := time.Parse("2006-01-02 15:04:05", purchase.TransactionDatetime)
			if err != nil {
				continue
			}

			config.DB.Create(&model.Purchase{
				UserID:       u.ID,
				PharmacyName: purchase.PharmacyName,
				MaskName:     purchase.MaskName,
				Amount:       purchase.TransactionAmount,
				Quantity:     purchase.TransactionQuantity,
				Date:         date,
			})
		}
	}
	fmt.Println("匯入用戶資料結束")
}

func main() {
	// 資料庫連線
	config.InitDB()
	// 自動建立相關資料表
	config.DB.AutoMigrate(&model.Pharmacy{}, &model.PharmacyOpenHour{}, &model.Mask{}, &model.User{}, &model.Purchase{})
	// 載入藥局相關資料
	loadPharmacies()
	// 載入用戶相關資料
	loadUsers()
}
