package crayfi

import "fmt"

type CardsService struct {
	client *Client
}

func (s *CardsService) Initiate(data interface{}) (interface{}, error) {
	return s.client.post("/api/v2/initiate", data)
}

func (s *CardsService) Charge(data interface{}) (interface{}, error) {
	return s.client.post("/api/v2/charge", data)
}

func (s *CardsService) Query(customerReference string) (interface{}, error) {
	return s.client.get(fmt.Sprintf("/api/query/%s", customerReference), nil)
}
