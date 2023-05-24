package register

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Challenge struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

type Promise struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var challenges []Challenge
var promises []Promise

func route() {
	r := mux.NewRouter()

	// challenge routes
	r.HandleFunc("/challenge", createChallenge).Methods("POST")
	r.HandleFunc("/challenge/{id}", getChallenge).Methods("GET")
	r.HandleFunc("/challenges", getAllChallenges).Methods("GET")
	r.HandleFunc("/challenges/{id}", updateChallenge).Methods("PUT")
	r.HandleFunc("/challenges/{id}", deleteChallenge).Methods("DELETE")

	// promise routes
	r.HandleFunc("/promise", createPromise).Methods("POST")
	r.HandleFunc("/promise/{id}", getPromise).Methods("GET")
	r.HandleFunc("/promises", getAllPromises).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func createChallenge(w http.ResponseWriter, r *http.Request) {
	var challenge Challenge
	err := json.NewDecoder(r.Body).Decode(&challenge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	challenges = append(challenges, challenge)
	json.NewEncoder(w).Encode(challenge)
}

func getChallenge(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, challenge := range challenges {
		if challenge.ID == params["id"] {
			json.NewEncoder(w).Encode(challenge)
			return
		}
	}
	http.Error(w, "Challenge not found", http.StatusNotFound)
}

func getAllChallenges(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(challenges)
}

func createPromise(w http.ResponseWriter, r *http.Request) {
	var promise Promise
	err := json.NewDecoder(r.Body).Decode(&promise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	promises = append(promises, promise)
	json.NewEncoder(w).Encode(promise)
}

func getPromise(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, promise := range promises {
		if promise.ID == params["id"] {
			json.NewEncoder(w).Encode(promise)
			return
		}
	}
	http.Error(w, "Promise not found", http.StatusNotFound)
}

func getAllPromises(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(promises)
}
func updateChallenge(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, challenge := range challenges {
		if challenge.ID == params["id"] {
			json.NewEncoder(w).Encode(challenge)
			return
		}
	}
	http.Error(w, "Challenge not updated", http.StatusNotFound)
}

func deleteChallenge(w http.ResponseWriter, r *http.Request) {
	var challenge Challenge
	err := json.NewDecoder(r.Body).Decode(&challenge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	challenges = append(challenges, challenge)
	json.NewEncoder(w).Encode(challenge)
}
