package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing blood donations
type SmartContract struct {
	contractapi.Contract
}

// Donation describes basic details of a blood donation
type Donation struct {
	ID           string `json:"ID"`
	DonorID      string `json:"DonorID"`
	RecipientID  string `json:"RecipientID"`
	BloodType    string `json:"BloodType"`
	DonationDate string `json:"DonationDate"`
	Expiration   string `json:"Expiration"`
	Status       string `json:"Status"`
}

// InitLedger adds a base set of donations to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	donations := []Donation{
		{ID: "donation1", DonorID: "donor1", RecipientID: "recipient1", BloodType: "O+", DonationDate: "2023-04-30", Expiration: "2023-05-15", Status: "available"},
		{ID: "donation2", DonorID: "donor2", RecipientID: "recipient2", BloodType: "A+", DonationDate: "2023-04-28", Expiration: "2023-05-13", Status: "available"},
		{ID: "donation3", DonorID: "donor3", RecipientID: "", BloodType: "B+", DonationDate: "2023-04-25", Expiration: "2023-05-10", Status: "available"},
	}

	for _, donation := range donations {
		donationJSON, err := json.Marshal(donation)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(donation.ID, donationJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateDonation issues a new blood donation to the world state with given details.
func (s *SmartContract) CreateDonation(ctx contractapi.TransactionContextInterface, id string, donorID string, recipientID string, bloodType string, donationDate string, expiration string) error {
	exists, err := s.DonationExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the donation %s already exists", id)
	}

	donation := Donation{
		ID:           id,
		DonorID:      donorID,
		RecipientID:  recipientID,
		BloodType:    bloodType,
		DonationDate: donationDate,
		Expiration:   expiration,
		Status:       "available",
	}
	donationJSON, err := json.Marshal(donation)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, donationJSON)
}

// UpdateDonation updates an existing blood donation in the world state with provided parameters.
func (s *SmartContract) UpdateDonation(ctx contractapi.TransactionContextInterface, id string, donorID string, recipientID string, bloodType string, donationDate string, expiration string, status string) error {
exists, err := s.DonationExists(ctx, id)
if err != nil {
return err
}
if !exists {
return fmt.Errorf("the donation %s does not exist", id)
}
donation := Donation{
	ID:           id,
	DonorID:      donorID,
	RecipientID:  recipientID,
	BloodType:    bloodType,
	DonationDate: donationDate,
	Expiration:   expiration,
	Status:       status,
}
donationJSON, err := json.Marshal(donation)
if err != nil {
	return err
}

return ctx.GetStub().PutState(id, donationJSON)
}

// QueryDonation returns the donation stored in the world state with given id.
func (s *SmartContract) QueryDonation(ctx contractapi.TransactionContextInterface, id string) (*Donation, error) {
donationJSON, err := ctx.GetStub().GetState(id)
if err != nil {
return nil, fmt.Errorf("failed to read from world state: %v", err)
}
if donationJSON == nil {
return nil, fmt.Errorf("the donation %s does not exist", id)
}
var donation Donation
err = json.Unmarshal(donationJSON, &donation)
if err != nil {
	return nil, err
}

return &donation, nil
}

// DonationExists returns true when donation with given id exists in world state.
func (s *SmartContract) DonationExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
donationJSON, err := ctx.GetStub().GetState(id)
if err != nil {
return false, fmt.Errorf("failed to read from world state: %v", err)
}
return donationJSON != nil, nil
}
