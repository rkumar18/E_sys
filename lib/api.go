package lib

import(
	"net/http"
	"E_PANCHAYAT/helpers"
	"encoding/json"
	"fmt"
)

func home(w http.ResponseWriter, r *http.Request){
	helpers.WriteResponse(w, "This is Home Page of E-panchayat" , http.StatusOK)
}

func login(w http.ResponseWriter, r *http.Request){
	var input Users 
	json.NewDecoder(r.Body).Decode(&input)

}

func signin(w http.ResponseWriter, r *http.Request){
	var input Users 
	fmt.Printf("1")
	json.NewDecoder(r.Body).Decode(&input)
	result := dbsignin(&input)
	if result.Email==""{
		helpers.WriteResponse(w, "User sign up successfully" , http.StatusOK)
	}else{
		helpers.WriteResponse(w, "Record already exist" , http.StatusBadRequest)
	}
}

// func yoganas(w http.ResponseWriter, r *http.Request){

// }

// func rti(w http.ResponseWriter, r *http.Request){
	
// }