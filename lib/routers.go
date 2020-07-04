package lib

import(
	"github.com/gorilla/mux"
	"net/http"
)

func Router() http.Handler {
	router := mux.NewRouter()
	// router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/home", home).Methods("GET")
	// router.HandleFunc("/yoganas",yoganas).Methods("GET")
	// router.HandleFunc("/rti",rti).Methods("GET")
	router.Handle("/",router)
	return router
}