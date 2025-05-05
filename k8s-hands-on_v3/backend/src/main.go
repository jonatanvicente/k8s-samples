package main

import (
	//"fmt"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type HandsOn struct {
	Time 		time.Time 	`json:"time"`
	Hostname 	string		`json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//w.Write([]byte(`{"message": "Hello World"}`))
	//resp := fmt.Sprintf("La hora es %v y hostname es %v", time.Now(), os.Getenv("HOSTNAME"))

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := HandsOn{
		Time: 		time.Now(),
		Hostname: 	os.Getenv("HOSTNAME"),
	}
	jsonResp, err := json.Marshal(&resp)
	if(err != nil){
		w.Write([]byte("Error"))
		return
	}
	//w.Write([]byte(resp))
	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonResp)
}

func main(){
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":9090", nil)
}