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

// func FindPharmaciesByMaskCount(priceMin, priceMax float64, countOp string, countVal int) []model.Pharmacy {}
// func GetTopUsersBySpending(start, end time.Time, limit int) []TopUserDTO {}
// func ProcessMultiPurchase(req PurchaseRequestDTO) error {}
// func ChangeStock(maskID uint, amount int) error {}
// func CreateOrUpdateMasks(pharmacyID uint, masks []MaskDTO) error {}
// func SearchByName(term string) []SearchResultDTO {}
