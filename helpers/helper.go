  
package helpers

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(write http.ResponseWriter, response interface{}, statusCode int){
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(statusCode)
	json.NewEncoder(write).Encode(response)
}