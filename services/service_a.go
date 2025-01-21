package services

import (
	"api-gateway/middleware"
	"api-gateway/models"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

func NewAServiceApiController(db *gorm.DB, aService string) *AServiceApiController {
	return &AServiceApiController{
		db:       db,
		aService: aService,
	}
}

type AServiceApiController struct {
	db       *gorm.DB
	aService string
}

func (g *AServiceApiController) EndPoint1(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	url := "https://rsapi.xxx.io/xxx1"
	queryParams := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			queryParams[key] = values[0]
		}
	}
	resp, err := middleware.ProxyRequest(url, "GET", queryParams, g.aService)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)

	// Convert queryParams to JSON
	paramsJSON, err := json.Marshal(queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resultsJSON := postgres.Jsonb{RawMessage: body}

	// Log the request and response
	log := models.Log{
		Endpoint:  url,
		Method:    "GET",
		Params:    postgres.Jsonb{RawMessage: paramsJSON},
		Results:   resultsJSON,
		IPAddress: r.RemoteAddr,
		Status:    resp.StatusCode,
		Elapsed:   time.Since(startTime).Seconds(),
	}

	if err = g.db.Table("logs").Create(&log).Error; err != nil {
		// Handle the error, e.g., log it or return it
		fmt.Printf("Error saving log: %v\n", err)
	}
}

func (g *AServiceApiController) EndPoint2(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	url := "https://rsapi.xxx.io/xxx2"
	queryParams := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			queryParams[key] = values[0]
		}
	}
	resp, err := middleware.ProxyRequest(url, "GET", queryParams, g.aService)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)

	// Convert queryParams to JSON
	paramsJSON, err := json.Marshal(queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resultsJSON := postgres.Jsonb{RawMessage: body}

	// Log the request and response
	log := models.Log{
		Endpoint:  url,
		Method:    "GET",
		Params:    postgres.Jsonb{RawMessage: paramsJSON},
		Results:   resultsJSON,
		IPAddress: r.RemoteAddr,
		Status:    resp.StatusCode,
		Elapsed:   time.Since(startTime).Seconds(),
	}

	if err = g.db.Table("logs").Create(&log).Error; err != nil {
		// Handle the error, e.g., log it or return it
		fmt.Printf("Error saving log: %v\n", err)
	}
}
