package crayfi

type FXService struct {
	client *Client
}

func (s *FXService) Rates(data interface{}) (interface{}, error) {
	return s.client.post("/v2/merchant/rates", data)
}

func (s *FXService) RatesByDestination(data interface{}) (interface{}, error) {
	return s.client.post("/v2/merchant/rates/destination", data)
}

func (s *FXService) Quote(data interface{}) (interface{}, error) {
	return s.client.post("/v2/merchant/quote", data)
}

func (s *FXService) Convert(data interface{}) (interface{}, error) {
	return s.client.post("/v2/merchant/conversions/convert", data)
}

func (s *FXService) Conversions() (interface{}, error) {
	return s.client.get("/v2/merchant/conversions", nil)
}
