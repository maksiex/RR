package routes

import (
	"github.com/gorilla/mux"
	"github.com/maksiecluster/booking/pkg/controllers"
)

var RegisterRoomsRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.GetTemplate).Methods("GET")
	router.HandleFunc("/card", controllers.GetTemplate).Methods("GET")
	router.HandleFunc("/login", controllers.GetTemplate).Methods("GET")
	router.HandleFunc("/signup", controllers.GetTemplate).Methods("GET")
	router.HandleFunc("/orders", controllers.GetTemplate).Methods("GET")
	router.HandleFunc("/rent", controllers.GetTemplate).Methods("GET")

	router.HandleFunc("/main", controllers.GetAllRooms).Methods("POST")

	router.HandleFunc("/reg", controllers.Reg).Methods("POST")

	router.HandleFunc("/create", controllers.CreateRent).Methods("POST")

	router.HandleFunc("/room/{roomId}", controllers.GetRoomById).Methods("GET")

	router.HandleFunc("/room/{roomId}", controllers.UpdateRoom).Methods("PUT")

	router.HandleFunc("/room/{roomId}", controllers.DeleteRoom).Methods("DELETE")

	router.HandleFunc("/{[a-z]}", controllers.GetTemplate).Methods("GET")
}
