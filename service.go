package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/articles", returnAllArticles)

	myRouter.HandleFunc("/articles", createNewArticles).Methods("POST")

	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page")

	fmt.Print("Endppoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	response := Response{Success: true, Message: "Operation completed successfully", Data: Articles}
	json.NewEncoder(w).Encode(response)
}

func createNewArticles(w http.ResponseWriter, r *http.Request) {
	//Get the body of the post request

	reqBody, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(w, "%+v", string(reqBody))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	index, _ := strconv.ParseInt(key, 0, 64)
	json.NewEncoder(w).Encode(Articles[index])
}
