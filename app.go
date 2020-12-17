package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	// "github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

type postBody struct {
	Number string `json:"number"`
}

type resultBody struct {
	Number string `json:"number"`
	Factor string `json:"factor"`
	Result int    `json:"result"`
}

// MultiplyBy
func multiplyBy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var postbody postBody
	_ = json.NewDecoder(r.Body).Decode(&postbody)
	numberstr := postbody.Number
	number, _ := strconv.Atoi(numberstr)
	json.NewEncoder(w).Encode(resultBody{Number: numberstr, Factor: "5", Result: number * 5})
}

// DivideBy
func divideBy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var postbody postBody
	_ = json.NewDecoder(r.Body).Decode(&postbody)
	numberstr := postbody.Number
	number, _ := strconv.Atoi(numberstr)
	json.NewEncoder(w).Encode(resultBody{Number: numberstr, Factor: "5", Result: number / 5})
}

// Main function
func main() {
	// // Init router
	// r := mux.NewRouter()

	// // Route handles & endpoints
	// r.HandleFunc("/v1.0/MultiplyBy/5", multiplyBy).Methods("POST")
	// r.HandleFunc("/v1.0/DivideBy/5", divideBy).Methods("POST")

	// // Start server
	// log.Fatal(http.ListenAndServe(":5000", r))

	router := httprouter.New()
	router.POST("/v1.0/MultiplyBy/5", multiplyBy)
	router.GET("/v1.0/DivideBy/5", divideBy)

	log.Fatal(http.ListenAndServe(":5000", router))
}
