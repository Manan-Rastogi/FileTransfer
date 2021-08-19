package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)


func viewData(w http.ResponseWriter, r *http.Request){
	fmt.Println("Request Received")
	
	b, err := ioutil.ReadAll(r.Body)

	if(err != nil){
		fmt.Println("Error 1 : ", err)
	}
	defer r.Body.Close()

	var msg map[string]string

	err = json.Unmarshal(b, &msg)

	if(err != nil){
		fmt.Println("Error 2 : ", err)
	}

	fmt.Println(msg)

	filepath := msg["path"] + "\\" + msg["name"]

	data, err := ioutil.ReadFile(filepath)

	if(err != nil){
		fmt.Println("Error 3 : ", err)
	}

	fmt.Println("Data : ",string(data))
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/view",viewData).Methods("POST")
	http.ListenAndServe(":8081", r)
}