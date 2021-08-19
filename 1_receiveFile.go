package main

import (
	//"bytes"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)




func uploadFile(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(10<<20)

	file, _, err := r.FormFile("someFile")

	if(err != nil){
		fmt.Println("Error 1 : ", err)
		return
	}

	defer file.Close()
	// fmt.Printf("Uploaded File : %+v\n", handler.Filename)
	// fmt.Printf("File Size : %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)


	tempFile, err := ioutil.TempFile("temp-files","upload-*.txt")

	if(err!=nil){
		fmt.Println("Error 2 : ",err)
		return
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if(err != nil){
		fmt.Println("Error 3 : ", err)
		return
	}
	
	tempFile.Write(fileBytes)

	fmt.Println(w, "Successfully Uploaded File!")


	url := "http://localhost:8081/view"
	// req, err := http.NewRequest(r.Method, url, bytes.NewReader(fileBytes))

	// res := req.Response.Body

	// fmt.Println("Response : ",res)
	name := tempFile.Name()
	thepath, err := filepath.Abs(filepath.Dir(name))

	fmt.Println("PATH : ", thepath)
	
	filename := strings.Split(name, "\\")
	postBody, _ := json.Marshal(map[string]string{
		"path" : string(thepath),
		"name" : filename[1],
	})
	respBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(url, "application/json", respBody)

	if(err != nil){
		fmt.Println("Error 4 : ", err)
		return
	}
	defer resp.Body.Close()
}




func setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", uploadFile).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080",r))
}



// func sendFile(){
	
// }



func main() {
	setupRoutes()
}