package contracts

import (
	"encoding/json"
	"fmt"
	

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// MedicineContract contract for managing CRUD for Medicine
type MedicineContract struct {
	contractapi.Contract
}

type Medicine struct {
	AssetType   string `json:"assetType"`
	MedicineID  string `json:"medicineId"`
	Name        string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	MFD         string `json:"mfd"`
	Expiry      string `json:"expiry"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
	Status      string `json:"status"`
}

type HistoryQueryResult struct {
	Record     *Medicine `json:"record"`
	TxId       string    `json:"txId"`
	Timestamp  string    `json:"timestamp"`
	IsDelete   bool      `json:"isDelete"`
}

// MedicineExists checks if medicine exists in world state
func (m *MedicineContract) MedicineExists(ctx contractapi.TransactionContextInterface, medID string) (bool, error) {
	data, err := ctx.GetStub().GetState(medID)
	if err != nil {
		return false, err
	}
	return data != nil, nil
}

// CreateMedicine by Org1MSP
func (m *MedicineContract) CreateMedicine(ctx contractapi.TransactionContextInterface, medID, name, manufacturer, mfd, expiry, price, quantity string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID != "Org1MSP" {
		return "", fmt.Errorf("MSP %v not permitted to create medicine", clientOrgID)
	}

	exists, err := m.MedicineExists(ctx, medID)
	if err != nil {
		return "", err
	}
	if exists {
		return "", fmt.Errorf("medicine %v already exists", medID)
	}

	medicine := Medicine{
		AssetType:    "medicine",
		MedicineID:   medID,
		Name:         name,
		Manufacturer: manufacturer,
		MFD:          mfd,
		Expiry:       expiry,
		Price:        price,
		Quantity:     quantity,
		Status:       "Manufactured",
	}

	bytes, err := json.Marshal(medicine)
	if err != nil {
		return "", fmt.Errorf("failed to marshal medicine data: %v", err)
	}

	err = ctx.GetStub().PutState(medID, bytes)
	if err != nil {
		return "", fmt.Errorf("failed to store medicine in ledger: %v", err)
	}

	return fmt.Sprintf("Added medicine %v", medID), nil
}


// ReadMedicine retrieves medicine from state
func (m *MedicineContract) ReadMedicine(ctx contractapi.TransactionContextInterface, medID string) (*Medicine, error) {
	bytes, err := ctx.GetStub().GetState(medID)
	if err != nil {
		return nil, err
	}
	if bytes == nil {
		return nil, fmt.Errorf("medicine %v not found", medID)
	}
	var medicine Medicine
	_ = json.Unmarshal(bytes, &medicine)
	return &medicine, nil
}

// DeleteMedicine removes the instance of Medicine from the world state — Org1MSP only
func (m *MedicineContract) DeleteMedicine(ctx contractapi.TransactionContextInterface, medID string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID != "Org1MSP" {
		return "", fmt.Errorf("MSP %v is not authorized to perform this action", clientOrgID)
	}
	
	// Check if medicine exists
	exists, err := m.MedicineExists(ctx, medID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state: %v", err)
	} else if !exists {
		return "", fmt.Errorf("the medicine %s does not exist", medID)
	}

	// Delete medicine state
	err = ctx.GetStub().DelState(medID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Medicine with ID %v has been deleted from the world state.", medID), nil
}


// GetAllMedicines
func (m *MedicineContract) GetAllMedicines(ctx contractapi.TransactionContextInterface) ([]*Medicine, error) {
	query := `{"selector":{"assetType":"medicine"}}`
	iter, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer iter.Close()
	return medicineResultIteratorFunction(iter)
}



func medicineResultIteratorFunction(iter shim.StateQueryIteratorInterface) ([]*Medicine, error) {
	var meds []*Medicine
	for iter.HasNext() {
		queryResult, _ := iter.Next()
		var med Medicine
		_ = json.Unmarshal(queryResult.Value, &med)
		meds = append(meds, &med)
	}
	return meds, nil
}
