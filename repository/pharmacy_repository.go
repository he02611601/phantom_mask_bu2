package repository

import (
	"phantom_mask_bu2/config"
	"phantom_mask_bu2/model"
)

func FilterPharmacies() []model.Pharmacy {
	var pharmacies []model.Pharmacy
	config.DB.Preload("Masks").Preload("OpenHours").Find(&pharmacies)
	return pharmacies
}

func FindMasksByPharmacy(id uint) model.Pharmacy {
	var pharmacy model.Pharmacy
	config.DB.Preload("Masks").Preload("OpenHours").First(&pharmacy, id)
	return pharmacy
}

// func FilterPharmacies(day string, time string) []model.Pharmacy {}
// func FindMasksByPharmacy(pharmacyID uint, sortBy string, order string) []model.Mask {}
// func FindPharmaciesByMaskCount(priceMin, priceMax float64, countOp string, countVal int) []model.Pharmacy {}
// func GetTopUsersBySpending(start, end time.Time, limit int) []TopUserDTO {}
// func ProcessMultiPurchase(req PurchaseRequestDTO) error {}
// func ChangeStock(maskID uint, amount int) error {}
// func CreateOrUpdateMasks(pharmacyID uint, masks []MaskDTO) error {}
// func SearchByName(term string) []SearchResultDTO {}
