package dto

type ResponsePharmacy struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	OpeningHours string `json:"opening_hours"`
}

type ResponseMask struct {
	ID         uint    `json:"id"`
	PharmacyID uint    `json:"pharmacy_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
}
