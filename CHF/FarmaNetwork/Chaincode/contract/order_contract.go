package contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// OrderContract contract for managing Medicine Orders
type OrderContract struct {
	contractapi.Contract
}

type Order struct {
	AssetType    string `json:"assetType"`
	OrderID      string `json:"orderID"`
	MedicineName string `json:"medicineName"`
	Quantity     string `json:"quantity"`
	Distributor  string `json:"distributor"`
}

// getOrderCollection returns private data collection name
func getOrderCollection() string {
	return "OrderCollection"
}

// OrderExists checks if an order exists in private data collection
func (o *OrderContract) OrderExists(ctx contractapi.TransactionContextInterface, orderID string) (bool, error) {
	collectionName := getOrderCollection()
	data, err := ctx.GetStub().GetPrivateDataHash(collectionName, orderID)
	if err != nil {
		return false, err
	}
	return data != nil, nil
}

// CreateOrder creates a new private order — Org2MSP only
func (o *OrderContract) CreateOrder(ctx contractapi.TransactionContextInterface, orderID string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}
	if clientOrgID != "Org2MSP" {
		return "", fmt.Errorf("organisation with MSPID %v cannot create order", clientOrgID)
	}

	exists, err := o.OrderExists(ctx, orderID)
	if err != nil {
		return "", err
	}
	if exists {
		return "", fmt.Errorf("the order %v already exists", orderID)
	}

	transientData, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", err
	}
	if len(transientData) == 0 {
		return "", fmt.Errorf("please provide transient data of medicineName, quantity, distributor")
	}

	medicineName, exists := transientData["medicineName"]
	if !exists {
		return "", fmt.Errorf("medicineName not specified in transient data")
	}
	quantity, exists := transientData["quantity"]
	if !exists {
		return "", fmt.Errorf("quantity not specified in transient data")
	}
	distributor, exists := transientData["distributor"]
	if !exists {
		return "", fmt.Errorf("distributor not specified in transient data")
	}

	order := Order{
		AssetType:    "Order",
		OrderID:      orderID,
		MedicineName: string(medicineName),
		Quantity:     string(quantity),
		Distributor:  string(distributor),
	}

	orderBytes, _ := json.Marshal(order)
	collectionName := getOrderCollection()

	return fmt.Sprintf("Order %v created successfully", orderID), ctx.GetStub().PutPrivateData(collectionName, orderID, orderBytes)
}

// ReadOrder retrieves a private order by ID
func (o *OrderContract) ReadOrder(ctx contractapi.TransactionContextInterface, orderID string) (*Order, error) {
	exists, err := o.OrderExists(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("order %v does not exist", orderID)
	}
	return readPrivateOrderState(ctx, orderID)
}

// DeleteOrder removes an order from private data — Org1MSP, Org2MSP only
func (o *OrderContract) DeleteOrder(ctx contractapi.TransactionContextInterface, orderID string) error {
	clientOrgID, _ := ctx.GetClientIdentity().GetMSPID()
	if clientOrgID != "Org1MSP" && clientOrgID != "Org2MSP" {
		return fmt.Errorf("organisation with MSPID %v cannot delete orders", clientOrgID)
	}

	exists, err := o.OrderExists(ctx, orderID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("order %v does not exist", orderID)
	}

	collectionName := getOrderCollection()
	return ctx.GetStub().DelPrivateData(collectionName, orderID)
}

// GetAllOrders fetches all orders with assetType=Order
func (o *OrderContract) GetAllOrders(ctx contractapi.TransactionContextInterface) ([]*Order, error) {
	collectionName := getOrderCollection()
	queryString := `{"selector":{"assetType":"Order"}}`
	iter, err := ctx.GetStub().GetPrivateDataQueryResult(collectionName, queryString)
	if err != nil {
		return nil, err
	}
	defer iter.Close()
	return orderResultIteratorFunction(iter)
}

// GetOrdersByRange retrieves orders based on start and end keys
func (o *OrderContract) GetOrdersByRange(ctx contractapi.TransactionContextInterface, startKey string, endKey string) ([]*Order, error) {
	collectionName := getOrderCollection()
	iter, err := ctx.GetStub().GetPrivateDataByRange(collectionName, startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer iter.Close()
	return orderResultIteratorFunction(iter)
}

// Helper function — read private order data
func readPrivateOrderState(ctx contractapi.TransactionContextInterface, orderID string) (*Order, error) {
	collectionName := getOrderCollection()
	bytes, err := ctx.GetStub().GetPrivateData(collectionName, orderID)
	if err != nil {
		return nil, err
	}
	var order Order
	err = json.Unmarshal(bytes, &order)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}
	return &order, nil
}

// Iterator function — converts query results to []*Order
func orderResultIteratorFunction(iter shim.StateQueryIteratorInterface) ([]*Order, error) {
	var orders []*Order
	for iter.HasNext() {
		queryResult, err := iter.Next()
		if err != nil {
			return nil, err
		}
		var order Order
		err = json.Unmarshal(queryResult.Value, &order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}
