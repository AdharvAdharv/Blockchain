package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Medicine structure matching your chaincode
type Medicine struct {
	MedicineID   string `json:"medicineId"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	MFD          string `json:"mfd"`
	Expiry       string `json:"expiry"`
	Price        string `json:"price"`
	Quantity     string `json:"quantity"`
}

func main() {
	router := gin.Default()

	router.Static("/public", "./public")
	router.LoadHTMLGlob("templates/*")

	// Home route
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/api/medicine", func(ctx *gin.Context) {
		var req Medicine
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}
	
		fmt.Printf("Received medicine creation request: %+v\n", req)
	
		// Collect arguments
		args := []string{
			req.MedicineID,
			req.Name,
			req.Manufacturer,
			req.MFD,
			req.Expiry,
			req.Price,
			req.Quantity,
		}
	
		fmt.Printf("Arguments for CreateMedicine txn: %v\n", args)
	
		// Submit transaction
		result := submitTxnFn(
			"org1", "farmanetwork", "Farma-Network",
			"MedicineContract", "invoke", make(map[string][]byte),
			"CreateMedicine",
			args...,
		)
	
		fmt.Printf("Transaction result: %v\n", result)
	
		ctx.JSON(http.StatusOK, req)
	})
	
	// GET /api/medicine/:id — read medicine by ID
	router.GET("/api/medicine/:id", func(ctx *gin.Context) {
		medicineId := ctx.Param("id")

		// Call chaincode query transaction
		result := submitTxnFn(
			"org1", "farmanetwork", "Farma-Network",
			"MedicineContract", "query", make(map[string][]byte),
			"ReadMedicine", medicineId,
		)
		fmt.Printf("Result from chaincode: %v\n", result)


		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	// GET /api/medicines — get all medicines
	router.GET("/api/medicines", func(ctx *gin.Context) {
		// Call chaincode query transaction
		result := submitTxnFn(
			"org1", "farmanetwork", "Farma-Network",
			"MedicineContract", "query", make(map[string][]byte),
			"GetAllMedicines",
		)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	// Run server
	router.Run("localhost:3001")
}
 