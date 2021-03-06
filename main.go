package main

import (
    "fmt"
     "io/ioutil"
    "net/http"
    "log"
    "github.com/gorilla/mux"
     "encoding/json"
)

// func homePage(w http.ResponseWriter, r *http.Request){
//     fmt.Fprintf(w, "Welcome to the HomePage!")
//     fmt.Println("Endpoint Hit: homePage")
// }

// func handleRequests() {
//     http.HandleFunc("/", homePage)
//     log.Fatal(http.ListenAndServe(":10000", nil))
// }

// func main() {
//     handleRequests()
// }
type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", createEvent)
	log.Fatal(http.ListenAndServe(":8080", router))
}