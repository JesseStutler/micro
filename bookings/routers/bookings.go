package routers

import (
	"github.com/JesseStutler/micro/bookings/controllers"
	"github.com/gorilla/mux"
)

func SetBookingsRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/bookings", controllers.GetBookings).Methods("GET")
	router.HandleFunc("/bookings", controllers.CreateBooking).Methods("POST")
	return router
}
