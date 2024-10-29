package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Asset struct {
	DealerID    string `json:"dealerId"`
	MSISDN      string `json:"msisdn"`
	MPIN        string `json:"mpin"`
	Balance     int    `json:"balance"`
	Status      string `json:"status"`
	TransAmount int    `json:"transAmount"`
	TransType   string `json:"transType"`
	Remarks     string `json:"remarks"`
}

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealerId string, msisdn string, mpin string, balance int, status string, transAmount int, transType string, remarks string) error {
	asset := Asset{
		DealerID:    dealerId,
		MSISDN:      msisdn,
		MPIN:        mpin,
		Balance:     balance,
		Status:      status,
		TransAmount: transAmount,
		TransType:   transType,
		Remarks:     remarks,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal asset: %w", err)
	}
	return ctx.GetStub().PutState(dealerId, assetJSON)
}

func (s *SmartContract) QueryAsset(ctx contractapi.TransactionContextInterface, dealerId string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(dealerId)
	if err != nil {
		return nil, fmt.Errorf("failed to read asset: %w", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("asset %s does not exist", dealerId)
	}
	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal asset: %w", err)
	}
	return &asset, nil
}

func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, dealerId string, balance int, status string) error {
	asset, err := s.QueryAsset(ctx, dealerId)
	if err != nil {
		return err
	}
	asset.Balance = balance
	asset.Status = status
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal updated asset: %w", err)
	}
	return ctx.GetStub().PutState(dealerId, assetJSON)
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic(fmt.Sprintf("failed to create chaincode: %v", err))
	}
	if err := chaincode.Start(); err != nil {
		panic(fmt.Sprintf("failed to start chaincode: %v", err))
	}
}
