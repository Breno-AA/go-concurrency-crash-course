package controller

import (
	"encoding/json"
	"go-concurrency-crash-course/service"
	"net/http"
)

var (
	carDetailsService service.CarDetailsService
)

type CarDetailsController interface {
	GetCarDetails(response http.ResponseWriter, request *http.Request)
}

type controller struct{}

func NewCarDetailsController(service service.CarDetailsService) CarDetailsController {
	carDetailsService = service
	return &controller{}
}

func (*controller) GetCarDetails(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	result := carDetailsService.GetDetails()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
