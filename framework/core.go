package framework

import "net/http"

type MicroService struct {
}

func NewMicroService() *MicroService {
	return &MicroService{}
}

func (c *MicroService) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}
