package crayfi

type WalletsService struct {
	client *Client
}

func (s *WalletsService) Balances() (interface{}, error) {
	return s.client.get("/wallet/api/v1/wallet/balances", nil)
}

func (s *WalletsService) Subaccounts() (interface{}, error) {
	return s.client.get("/wallet/api/v1/wallet/sub-accounts", nil)
}
