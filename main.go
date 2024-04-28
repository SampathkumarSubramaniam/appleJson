package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

// Example JSON data structure (adjust to match your file)
type jsonData struct {
	Servers []string `json:"Servers"`
	Version string   `json:"Version"`
	BaseURL string   `json:"BaseURL"`
}

func main() {
	http.HandleFunc("/.well-known/com.apple.remotemanagement.json", renderJSON)
	err := http.ListenAndServeTLS(":443", "./fullchain.pem", "./privkey.pem", nil)
	// err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func renderJSON(w http.ResponseWriter, r *http.Request) {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Println("Host:", host)

	// Read JSON file
	file, err := os.Open("abc.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	// Parse JSON data
	var data jsonData
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Encode and write JSON data to the response
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}

}
