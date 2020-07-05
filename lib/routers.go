package lib

import(
	"github.com/gorilla/mux"
	"net/http"
	
)

func Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/home", home).Methods("GET")
	//fmt.Printf("000")
	router.HandleFunc("/signin", signin).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/admin", admin).Methods("POST")
	// router.HandleFunc("/yoganas",yoganas).Methods("GET")
	// router.HandleFunc("/rti",rti).Methods("GET")
	router.Handle("/",router)
	return router
}