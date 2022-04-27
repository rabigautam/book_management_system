package routes

import(
	"github.com/gorilla/mux"
	"github.com/rabigautam/book_management_system/pkg/controllers"
)


var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")

}