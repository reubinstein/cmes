package util

import (
	"fmt"
	"net/http"
)

type AdminLevel int

const (
	National AdminLevel = iota
	Regional
	District
	Ward
)

type User struct {
	Username   string
	Password   string
	AdminLevel AdminLevel
}

type AdminLevelHandler struct {
	level AdminLevel
}

func (ah *AdminLevelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Welcome to the %s dashboard!", ah.level)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
	}
}

func admin() {
	// Simulating some data for registered users
	users := []User{
		{
			Username:   "national_user",
			Password:   "123456",
			AdminLevel: National,
		},
		{
			Username:   "regional_user",
			Password:   "123456",
			AdminLevel: Regional,
		},
		{
			Username:   "district_user",
			Password:   "123456",
			AdminLevel: District,
		},
		{
			Username:   "ward_user",
			Password:   "123456",
			AdminLevel: Ward,
		},
	}

	// Creating handlers for each admin level
	nationalHandler := &AdminLevelHandler{level: National}
	regionalHandler := &AdminLevelHandler{level: Regional}
	districtHandler := &AdminLevelHandler{level: District}
	wardHandler := &AdminLevelHandler{level: Ward}

	// Registering handlers to different endpoints
	http.Handle("/national", validateUser(National, users, nationalHandler))
	http.Handle("/regional", validateUser(Regional, users, regionalHandler))
	http.Handle("/district", validateUser(District, users, districtHandler))
	http.Handle("/ward", validateUser(Ward, users, wardHandler))

	// Starting server
	fmt.Println("Server started")
	http.ListenAndServe(":8080", nil)
}

// Middleware to validate user credentials and permissions
func validateUser(level AdminLevel, users []User, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
			return
		}

		for _, user := range users {
			if user.Username == username && user.Password == password && user.AdminLevel == level {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
	})
}
