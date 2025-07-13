package repository

import (
	"fmt"
	"phantom_mask_bu2/config"
	"phantom_mask_bu2/model"
)

func FilterPharmacies(day int, time string) []model.Pharmacy {
	var pharmacies []model.Pharmacy
	db := config.DB.Joins("JOIN pharmacy_open_hours ON pharmacies.id = pharmacy_open_hours.pharmacy_id")

	if day != 0 {
		db = db.Where("pharmacy_open_hours.day_of_week = ?", day)
	}
	if time != "" {
		db = db.Where("pharmacy_open_hours.start_time <= ? AND pharmacy_open_hours.end_time >= ?", time, time)
	}

	if err := db.Find(&pharmacies).Error; err != nil {
		return []model.Pharmacy{}
	}

	return pharmacies
}

func FindMasksByPharmacy(id uint, sort string, order string) []model.Mask {
	var masks []model.Mask
	db := config.DB.Where("pharmacy_id = ?", id)

	if sort != "" && order != "" {
		db = db.Order(fmt.Sprintf("%s %s", sort, order))
	}

	if err := db.Find(&masks).Error; err != nil {
		return []model.Mask{}
	}

	return masks
}

func FindPharmaciesByMaskCount(minPrice, maxPrice float64, minStock, maxStock int64) []model.Pharmacy {
	var pharmacies []model.Pharmacy
	subQuery := config.DB.Model(&model.Pharmacy{}).Select("pharmacies.*", "SUM(masks.stock) AS sum_stock").Joins("JOIN masks ON pharmacies.id = masks.pharmacy_id")

	if minPrice != 0 && maxPrice != 0 {
		subQuery = subQuery.Where("masks.price >= ? AND masks.price <= ?", minPrice, maxPrice)
	}
	subQuery = subQuery.Group("pharmacies.id")
	query := config.DB.Table("(?) AS sub", subQuery)

	if minStock != 0 && maxStock != 0 {
		query = query.Where("sum_stock >= ? AND sum_stock <= ?", minStock, maxStock)
	} else if minStock != 0 && maxStock == 0 {
		query = query.Where("sum_stock >= ?", minStock)
	} else if minStock == 0 && maxStock != 0 {
		query = query.Where("sum_stock <= ?", maxStock)
	}

	if err := query.Find(&pharmacies).Error; err != nil {
		return []model.Pharmacy{}
	}

	return pharmacies
}

// func GetTopUsersBySpending(start, end time.Time, limit int) []TopUserDTO {}
// func ProcessMultiPurchase(req PurchaseRequestDTO) error {}
// func ChangeStock(maskID uint, amount int) error {}
// func CreateOrUpdateMasks(pharmacyID uint, masks []MaskDTO) error {}
// func SearchByName(term string) []SearchResultDTO {}
