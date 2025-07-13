package controller

import (
	"net/http"
	"phantom_mask_bu2/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}

func GetPharmacies(c *gin.Context) {
	pharmacies := repository.FilterPharmacies()
	c.JSON(http.StatusOK, pharmacies)
}

func GetPharmacyMasks(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pharmacy := repository.FindMasksByPharmacy(uint(id))
	if pharmacy.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, pharmacy)
}

// func FilterPharmaciesByMaskCount(c *gin.Context) {}
// func GetTopSpenders(c *gin.Context){}
// func CreateMultiPurchase(c *gin.Context)
// func UpdateMaskStock(c *gin.Context)
// func BulkUpsertMasks(c *gin.Context)
// func SearchPharmaciesAndMasks(c *gin.Context)
