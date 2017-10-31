package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Data struct {
	Field1	string
	Field2	[]string
	Field3	int
}

func main() {
	fs := http.FileServer(http.Dir("/Users/JW/WebstormProjects/ML/public"))

	//routes
	http.Handle("/", fs)
	http.HandleFunc("/get-data", getData)

	fmt.Println("Go webserver listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getData(w http.ResponseWriter, r *http.Request) {
	data := Data{"this is a string", []string{"hello", "world you mofo", "how are you"}, 59}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}