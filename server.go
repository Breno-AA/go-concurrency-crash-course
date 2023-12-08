package main

import (
	"go-concurrency-crash-course/controller"
	router "go-concurrency-crash-course/http"
	"go-concurrency-crash-course/service"
)

var (
	carDetailsService    service.CarDetailsService       = service.NewCarDetailsService()
	CarDetailsController controller.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
	httpRouter           router.Router                   = router.NewChiRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/carDetails", CarDetailsController.GetCarDetails)
	httpRouter.SERVE(port)
}
