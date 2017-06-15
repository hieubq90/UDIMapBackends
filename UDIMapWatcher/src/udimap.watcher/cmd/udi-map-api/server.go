package main

import (
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"udimap.watcher/models"
)

// Server is an http server that handles REST requests.
type Server struct {
	db *gorm.DB
}

// NewServer creates a new instance of a Server.
func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

// RegisterRouter registers a router onto the Server.
func (s *Server) RegisterRouter(router *httprouter.Router) {
	router.GET("/ping", s.ping)

	router.GET("/current_info", s.getCustomers)
}

func (s *Server) ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	writeTextResult(w, "go/gorm")
}

func (s *Server) getCustomers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	currentInfo := &models.CurrentInfo{
		Err:              0,
		ListCameras:      make([]*models.Camera, 0),
		ListRainTracker:  make([]*models.TramDoMua, 0),
		ListTideTracker:  make([]*models.TramDoTrieu, 0),
		ListFloodTracker: make([]*models.TramQuanTracNgap, 0),
		ListFloodPoint:   make([]*models.FloodPoint, 0),
	}
	listCameras := make([]*models.Camera, 0)
	if errCamera := s.db.Find(&listCameras).Error; errCamera == nil {
		currentInfo.ListCameras = listCameras
	}

	listRainTracker := make([]*models.TramDoMua, 0)
	if errCamera := s.db.Find(&listRainTracker).Error; errCamera == nil {
		currentInfo.ListRainTracker = listRainTracker
	}

	listTideTracker := make([]*models.TramDoTrieu, 0)
	if errCamera := s.db.Find(&listTideTracker).Error; errCamera == nil {
		currentInfo.ListTideTracker = listTideTracker
	}

	listFloodTracker := make([]*models.TramQuanTracNgap, 0)
	if errCamera := s.db.Find(&listFloodTracker).Error; errCamera == nil {
		currentInfo.ListFloodTracker = listFloodTracker
	}

	listFloodPoint := make([]*models.FloodPoint, 0)
	if errCamera := s.db.Find(&listFloodPoint).Error; errCamera == nil {
		currentInfo.ListFloodPoint = listFloodPoint
	}

	writeJSONResult(w, currentInfo)
}

func writeTextResult(w http.ResponseWriter, res string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, res)
}

func writeJSONResult(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func writeMissingParamError(w http.ResponseWriter, paramName string) {
	http.Error(w, fmt.Sprintf("missing query param %q", paramName), http.StatusBadRequest)
}

func errToStatusCode(err error) int {
	switch err {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
