package main

import (
	"phantom_mask_bu2/config"
	"phantom_mask_bu2/router"
)

func main() {
	config.InitDB()
	r := router.SetupRouter()
	r.Run(":8080") // 或改為從環境變數讀取
}
