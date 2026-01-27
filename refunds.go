package crayfi

import "fmt"

type RefundsService struct {
	client *Client
}

func (s *RefundsService) Initiate(data interface{}) (interface{}, error) {
	return s.client.post("/api/refunds/initiate", data)
}

func (s *RefundsService) Query(reference string) (interface{}, error) {
	return s.client.get(fmt.Sprintf("/api/refunds/query/%s", reference), nil)
}
