package crayfi

import "fmt"

type MoMoService struct {
	client *Client
}

func (s *MoMoService) Initiate(data interface{}) (interface{}, error) {
	return s.client.post("/api/v2/momo/initiate", data)
}

func (s *MoMoService) Requery(customerReference string) (interface{}, error) {
	return s.client.get(fmt.Sprintf("/api/v2/momo/requery/%s", customerReference), nil)
}
