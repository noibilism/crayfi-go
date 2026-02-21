package crayfi

import "fmt"

type FXService struct {
	client *Client
}

func (s *FXService) Rates(data interface{}) (interface{}, error) {
	return s.client.post("/api/rates", data)
}

func (s *FXService) RatesByDestination(data interface{}) (interface{}, error) {
	return s.client.post("/api/rates/destination", data)
}

func (s *FXService) Quote(data interface{}) (interface{}, error) {
	return s.client.post("/api/quote", data)
}

func (s *FXService) Convert(data interface{}) (interface{}, error) {
	return s.client.post("/api/conversions", data)
}

func (s *FXService) Conversions() (interface{}, error) {
	return s.client.get("/api/conversions", nil)
}

func (s *FXService) DisputeConversion(conversionID string, data interface{}) (interface{}, error) {
	return s.client.post(fmt.Sprintf("/api/conversions/%s/dispute", conversionID), data)
}
