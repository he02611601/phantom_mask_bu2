package controller

import (
	"net/http"
	"phantom_mask_bu2/dto"
	"phantom_mask_bu2/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 星期轉換成整數資料
var DayOfWeek = map[string]int{
	"Mon":  1,
	"Tue":  2,
	"Wed":  3,
	"Thur": 4,
	"Fri":  5,
	"Sat":  6,
	"Sun":  7,
}

// 自行測試用
func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}

// 1. 列出藥局，可選擇依據特定時間和/或星期進行篩選
func GetPharmacies(c *gin.Context) {
	dayRaw := c.Query("day")
	time := c.Query("time")
	pharmacies := repository.FilterPharmacies(DayOfWeek[dayRaw], time)

	result := make([]dto.ResponsePharmacy, 0, len(pharmacies))
	for _, p := range pharmacies {
		result = append(result, dto.ResponsePharmacy{
			ID:           p.ID,
			Name:         p.Name,
			OpeningHours: p.OpeningHours,
		})
	}
	c.JSON(http.StatusOK, result)
}

// 2. 列出指定藥局販售的所有口罩產品，可選擇依名稱或價格排序
func GetPharmacyMasks(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parameter Error!"})
		return
	}
	sort := c.Query("sort")
	order := c.Query("order")
	masks := repository.FindMasksByPharmacy(uint(id), sort, order)

	result := make([]dto.ResponseMask, 0, len(masks))
	for _, m := range masks {
		result = append(result, dto.ResponseMask{
			ID:         m.ID,
			PharmacyID: m.PharmacyID,
			Name:       m.Name,
			Price:      m.Price,
			Stock:      m.Stock,
		})
	}

	c.JSON(http.StatusOK, result)
}

// func FilterPharmaciesByMaskCount(c *gin.Context) {}
// func GetTopSpenders(c *gin.Context){}
// func CreateMultiPurchase(c *gin.Context)
// func UpdateMaskStock(c *gin.Context)
// func BulkUpsertMasks(c *gin.Context)
// func SearchPharmaciesAndMasks(c *gin.Context)
