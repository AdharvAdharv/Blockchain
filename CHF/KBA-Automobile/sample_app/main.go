package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Car struct {
CarId        string `json:"carId"`
Make        string `json:"make"`
Model       string `json:"model"`
Color        string `json:"color"`
Date         string `json:"dateOfManufacture"`
// Manufacturer string `json:"manufacturerName"`
ManufacturerName string `json:"manufacturerName"`
}

func main() {
router := gin.Default()

router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:5173"},
    AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
    AllowCredentials: true,
}))
router.Static("/public", "./public")
// router.LoadHTMLGlob("templates/*")

router.GET("/", func(ctx *gin.Context) {
ctx.HTML(http.StatusOK, "index.html", gin.H{})
})

router.POST("/api/car", func(ctx *gin.Context) {
var req Car
if err := ctx.BindJSON(&req);
 err != nil {
ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
return
}

fmt.Printf("car response: %s", req)
// submitTxnFn("org1", "autochannel", "KBA-Automobile", "CarContract", "invoke", make(map[string][]byte), "CreateCar", req.CarId, req.Make, req.Model, req.Color, req.Manufacturer, req.Date)
submitTxnFn("org1", "autochannel", "KBA-Automobile", "CarContract", "invoke", make(map[string][]byte), "CreateCar", req.CarId, req.Make, req.Model, req.Color, req.Date,req.ManufacturerName)

ctx.JSON(http.StatusOK, req)
})

// router.GET("/api/car/:id", func(ctx *gin.Context) {
// carId := ctx.Param("id")
// result := submitTxnFn("org1", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "ReadCar", carId)

// ctx.JSON(http.StatusOK, gin.H{"data": result})
// })
router.GET("/api/car/:id", func(ctx *gin.Context) {
	carId := ctx.Param("id")
	result := submitTxnFn("org1", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "ReadCar", carId)

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(result), &parsed); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse car data"})
		return
	}

	ctx.JSON(http.StatusOK, parsed)
})

router.GET("/api/GetAllCars", func(ctx *gin.Context) {

	result := submitTxnFn("org1", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "GetAllCars")

	ctx.JSON(http.StatusOK, gin.H{"data": result})
})


router.Run("localhost:3001")
}