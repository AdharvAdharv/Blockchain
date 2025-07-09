package main

import (
	contracts "farmNetwork/contract"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	medicineContract:= new(contracts.MedicineContract)
	orderContract:=new(contracts.OrderContract)
	pharmacyContract:=new(contracts.PharmacyContract)

	chaincode,err :=contractapi.NewChaincode(medicineContract,orderContract,pharmacyContract)

	if err!=nil{
		log.Panicf("Could not create Chaincode : %v" ,err)
	}
	err = chaincode.Start()
	if err!=nil {
		log.Panicf("Failed to start chaincode : %v", err)
		
	}
}