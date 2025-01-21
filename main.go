// api-gateway/main.go
package main

import (
	"api-gateway/config"
	"api-gateway/services"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	// Check if environment variables are already set
	if os.Getenv("DB_HOST") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWORD") == "" ||
		os.Getenv("DB_NAME") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_SSLMODE") == "" ||
		os.Getenv("SERVICE_A_API_KEY") == "" {
		// Load .env file if environment variables are not set
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	db := config.GetDBInstance()
	r := mux.NewRouter()

	serviceA := config.GetServiceAKey()
	// Group v1 endpoints
	v1 := r.PathPrefix("/api/v1").Subrouter()

	goong := v1.PathPrefix("/service-a").Subrouter()

	// Route for Service A
	goongController := services.NewAServiceApiController(db, serviceA)
	goong.HandleFunc("/endpoint-1", goongController.EndPoint1).Methods("GET")
	goong.HandleFunc("/endpoint-2", goongController.EndPoint2).Methods("GET")

	// Start server
	fmt.Sprint("Server is running on port 8081")
	http.ListenAndServe(":8081", r)
}
