package contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// PharmacyContract handles medicine assignments to pharmacies
type PharmacyContract struct {
	contractapi.Contract
}

// PharmacyAssignment represents a medicine assigned to a pharmacy
type PharmacyAssignment struct {
	AssetType    string `json:"assetType"`
	MedicineID   string `json:"medicineID"`
	PharmacyName string `json:"pharmacyName"`
	Quantity     string `json:"quantity"`
}

// MedicineExists checks whether a medicine ID exists in world state
func (p *PharmacyContract) MedicineExists(ctx contractapi.TransactionContextInterface, medicineID string) (bool, error) {
	bytes, err := ctx.GetStub().GetState(medicineID)
	if err != nil {
		return false, err
	}
	return bytes != nil, nil
}

// AssignMedicineToPharmacy assigns an existing medicine to a pharmacy — Org3MSP only
func (p *PharmacyContract) AssignMedicineToPharmacy(ctx contractapi.TransactionContextInterface, medicineID string, pharmacyName string, quantity string) (string, error) {
	// MSP validation
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}
	if clientOrgID != "Org3MSP" {
		return "", fmt.Errorf("MSP %v is not authorized to assign medicine to pharmacy", clientOrgID)
	}

	// Check medicine exists
	exists, err := p.MedicineExists(ctx, medicineID)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("Medicine %v does not exist", medicineID)
	}

	// Create assignment struct
	assignment := PharmacyAssignment{
		AssetType:    "PharmacyAssignment",
		MedicineID:   medicineID,
		PharmacyName: pharmacyName,
		Quantity:     quantity,
	}

	// Marshal and store in world state
	bytes, _ := json.Marshal(assignment)
	err = ctx.GetStub().PutState(medicineID, bytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Medicine %v assigned to %v with quantity %v", medicineID, pharmacyName, quantity), nil
}

// ReadPharmacyAssignment reads an assignment for a medicine
func (p *PharmacyContract) ReadPharmacyAssignment(ctx contractapi.TransactionContextInterface, medicineID string) (*PharmacyAssignment, error) {
	// Check medicine exists
	exists, err := p.MedicineExists(ctx, medicineID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("Medicine %v does not exist", medicineID)
	}

	// Retrieve assignment
	bytes, err := ctx.GetStub().GetState(medicineID)
	if err != nil {
		return nil, err
	}
	if bytes == nil {
		return nil, fmt.Errorf("assignment for medicine %v does not exist", medicineID)

	}

	// Unmarshal and return
	var assignment PharmacyAssignment
	err = json.Unmarshal(bytes, &assignment)
	if err != nil {
		return nil, err
	}

	return &assignment, nil
}
