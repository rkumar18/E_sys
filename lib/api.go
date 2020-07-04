package lib

import(
	"net/http"
	"E_PANCHAYAT/helpers"
	"fmt"
)

func home(w http.ResponseWriter, r *http.Request){
	helpers.WriteResponse(w, "This is Home Page of e-panchayat" , http.StatusOK)
	fmt.Printf("0")
}

// func login(w http.ResponseWriter, r *http.Request){

// }

// func yoganas(w http.ResponseWriter, r *http.Request){

// }

// func rti(w http.ResponseWriter, r *http.Request){
	
// }