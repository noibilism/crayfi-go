package crayfi

import "fmt"

type RefundsService struct {
	client *Client
}

func (s *RefundsService) Initiate(data interface{}) (interface{}, error) {
	return s.client.post("/api/v2/refund/initiate", data)
}

func (s *RefundsService) Query(reference string) (interface{}, error) {
	return s.client.get(fmt.Sprintf("/api/v2/refund/query/%s", reference), nil)
}
