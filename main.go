package cmes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Policy struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func createMP(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the MP data
	var mp MP
	err := json.NewDecoder(r.Body).Decode(&mp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func main() {
	router := mux.NewRouter()

	// register routes for CRUD APIs
	router.HandleFunc("/mps", createMP).Methods("POST")
	router.HandleFunc("/mps", getMPs).Methods("GET")
	router.HandleFunc("/mps/{id}", getMP).Methods("GET")
	router.HandleFunc("/mps/{id}", updateMP).Methods("PUT")
	router.HandleFunc("/mps/{id}", deleteMP).Methods("DELETE")

	router.HandleFunc("/counselors", createCounselor).Methods("POST")
	router.HandleFunc("/counselors", getCounselors).Methods("GET")
	router.HandleFunc("/counselors/{id}", getCounselor).Methods("GET")
	router.HandleFunc("/counselors/{id}", updateCounselor).Methods("PUT")
	router.HandleFunc("/counselors/{id}", deleteCounselor).Methods("DELETE")

	router.HandleFunc("/challenges", createChallenge).Methods("POST")
	router.HandleFunc("/challenges", getChallenge).Methods("GET")
	router.HandleFunc("/challenges/{id}", getChallenge).Methods("GET")
	router.HandleFunc("/challenges/{id}", updateChallenge).Methods("PUT")
	router.HandleFunc("/challenges/{id}", deleteChallenge).Methods("DELETE")

	router.HandleFunc("/projects", createProject).Methods("POST")
	router.HandleFunc("/projects", getProjects).Methods("GET")
	router.HandleFunc("/projects/{id}", getProject).Methods("GET")
	router.HandleFunc("/projects/{id}", updateProject).Methods("PUT")
	router.HandleFunc("/projects/{id}", deleteProject).Methods("DELETE")

	router.HandleFunc("/policies", createPolicy).Methods("POST")
	router.HandleFunc("/policies", getAllPolicies).Methods("GET")
	router.HandleFunc("/policies/{id}", getPolicy).Methods("GET")
	router.HandleFunc("/policies/{id}", updatePolicy).Methods("PUT")
	router.HandleFunc("/policies/{id}", deletePolicy).Methods("DELETE")

	// register route for standards benchmarking API
	router.HandleFunc("/benchmark", benchmark).Methods("GET")

	// start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
