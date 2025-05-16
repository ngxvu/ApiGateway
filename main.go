package main

import (
	"api-gateway/config"
	"api-gateway/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func loadEnv() {
	required := []string{
		"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT", "DB_SSLMODE", "SERVICE_A_API_KEY",
	}
	for _, key := range required {
		if os.Getenv(key) == "" {
			if err := godotenv.Load(); err != nil {
				log.Fatalf("Error loading .env file")
			}
			break
		}
	}
}

func main() {
	loadEnv()

	db := config.GetDBInstance()
	r := mux.NewRouter()

	serviceA := config.GetServiceAKey()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	goong := v1.PathPrefix("/service-a").Subrouter()

	goongController := services.NewServiceAController(db, serviceA)
	goong.HandleFunc("/endpoint-1", goongController.EndPoint1).Methods("GET")
	goong.HandleFunc("/endpoint-2", goongController.EndPoint2).Methods("GET")

	log.Println("Server is running on port 8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
