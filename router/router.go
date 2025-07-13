package router

import (
	"phantom_mask_bu2/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// 1. 列出藥局，可選擇依據特定時間和/或星期進行篩選
		api.GET("/pharmacies", controller.GetPharmacies)
		// 2. 列出指定藥局販售的所有口罩產品，可選擇依名稱或價格排序
		api.GET("/pharmacies/:id/masks", controller.GetPharmacyMasks)
		// 3. 列出販售一定數量口罩（在指定價格範圍內）的藥局，可設定門檻為大於、小於或介於區間之間
		api.GET("/pharmacies/mask-filter", controller.HelloWorld)
		// 4. 顯示在特定日期區間內花費最多的前 N 名使用者
		api.GET("/users/top-spenders", controller.HelloWorld)
		// 5. 處理使用者一次向多間藥局購買口罩的請求
		api.POST("/purchases/multi", controller.HelloWorld)
		// 6. 更新現有口罩產品的庫存數量，可以增加或減少
		api.PATCH("/masks/:id/stock", controller.HelloWorld)
		// 7. 一次為指定藥局建立或更新多個口罩產品，包含名稱、價格與庫存數量
		api.PUT("/pharmacies/:id/masks/bulk", controller.HelloWorld)
		// 8. 根據名稱搜尋藥局或口罩，並依與搜尋字詞的相關性進行排序
		api.GET("/search", controller.HelloWorld)
	}

	return r
}
