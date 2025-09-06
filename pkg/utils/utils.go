package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseBody(r *http.Request , x interface{}){
	// if body , err := ioutil.ReadAll(r.Body); err == nil {
	// 	if err := json.Unmarshal([]byte(body), x) ; err != nil {
	// 		return 
	// 	}
	// }

	err := json.NewDecoder(r.Body).Decode(x)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	

}