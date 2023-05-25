package views

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var projects []Project

func main() {
	router := mux.NewRouter()

	// Create Project
	router.HandleFunc("/projects", createProjectHandler).Methods("POST")

	// Get all Projects
	router.HandleFunc("/projects", getAllProjectsHandler).Methods("GET")

	// Get Project by ID
	router.HandleFunc("/projects/{id}", getProjectByIDHandler).Methods("GET")

	// Update Project by ID
	router.HandleFunc("/projects/{id}", updateProjectHandler).Methods("PUT")

	// Delete Project by ID
	router.HandleFunc("/projects/{id}", deleteProjectHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var project Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	project.ID = len(projects) + 1
	projects = append(projects, project)
	json.NewEncoder(w).Encode(project)
}

func getAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func getProjectByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, project := range projects {
		if project.ID == projectID {
			json.NewEncoder(w).Encode(project)
			return
		}
	}
	http.NotFound(w, r)
}

func updateProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var updatedProject Project
	err = json.NewDecoder(r.Body).Decode(&updatedProject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, project := range projects {
		if project.ID == projectID {
			updatedProject.ID = projectID
			projects[i] = updatedProject
			json.NewEncoder(w).Encode(updatedProject)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	projectID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, project := range projects {
		if project.ID == projectID {
			projects = append(projects[:i], projects[i+1:]...)
			json.NewEncoder(w).Encode(project)
			return
		}
	}
	http.NotFound(w, r)
}
