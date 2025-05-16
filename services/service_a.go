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

type ServiceAController struct {
	db       *gorm.DB
	aService string
}

func NewServiceAController(db *gorm.DB, aService string) *ServiceAController {
	return &ServiceAController{
		db:       db,
		aService: aService,
	}
}

func (g *ServiceAController) handleEndpoint(w http.ResponseWriter, r *http.Request, url string) {
	startTime := time.Now()
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)

	paramsJSON, err := json.Marshal(queryParams)
	if err != nil {
		fmt.Printf("Error marshaling params: %v\n", err)
		return
	}

	logEntry := models.Log{
		Endpoint:  url,
		Method:    "GET",
		Params:    postgres.Jsonb{RawMessage: paramsJSON},
		Results:   postgres.Jsonb{RawMessage: body},
		IPAddress: r.RemoteAddr,
		Status:    resp.StatusCode,
		Elapsed:   time.Since(startTime).Seconds(),
	}

	if err = g.db.Table("logs").Create(&logEntry).Error; err != nil {
		fmt.Printf("Error saving log: %v\n", err)
	}
}

func (g *ServiceAController) EndPoint1(w http.ResponseWriter, r *http.Request) {
	g.handleEndpoint(w, r, "https://rsapi.xxx.io/xxx1")
}

func (g *ServiceAController) EndPoint2(w http.ResponseWriter, r *http.Request) {
	g.handleEndpoint(w, r, "https://rsapi.xxx.io/xxx2")
}
