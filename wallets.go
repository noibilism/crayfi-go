package crayfi

type WalletsService struct {
	client *Client
}

func (s *WalletsService) Balances() (interface{}, error) {
	return s.client.get("/api/balance", nil)
}

func (s *WalletsService) Subaccounts() (interface{}, error) {
	return s.client.get("/api/get-subaccount", nil)
}
