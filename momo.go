package crayfi

import "fmt"

type MoMoService struct {
	client *Client
}

func (s *MoMoService) Initiate(data interface{}) (interface{}, error) {
	return s.client.post("/momo/api/v1/momo/initiate", data)
}

func (s *MoMoService) Requery(customerReference string) (interface{}, error) {
	return s.client.get(fmt.Sprintf("/momo/api/v1/momo/requery/%s", customerReference), nil)
}
