package crayfi

import "fmt"

type PayoutsService struct {
	client *Client
}

func (s *PayoutsService) PaymentMethods(countryCode string) (interface{}, error) {
	return s.client.get(fmt.Sprintf("/api/payout/payment-methods/%s", countryCode), nil)
}

func (s *PayoutsService) Banks(countryCode string) (interface{}, error) {
	params := map[string]string{}
	if countryCode != "" {
		params["countryCode"] = countryCode
	}
	return s.client.get("/api/payout/banks", params)
}

func (s *PayoutsService) ValidateAccount(data interface{}) (interface{}, error) {
	return s.client.post("/api/payout/accounts/validate", data)
}

func (s *PayoutsService) Disburse(data interface{}) (interface{}, error) {
	return s.client.post("/api/payout/disburse", data)
}

func (s *PayoutsService) Requery(transactionId string) (interface{}, error) {
	return s.client.get(fmt.Sprintf("/api/payout/requery/%s", transactionId), nil)
}
