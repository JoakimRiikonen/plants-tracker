package main

import (
	"encoding/json"
	"fmt"
	"jr/plants-tracker/pkg/readings"
	"net/http"
)

type ReadingsHandler struct {
	store readings.ReadingStore
}

func NewReadingsHandler(s readings.ReadingStore) *ReadingsHandler {
	return &ReadingsHandler{
		store: s,
	}
}

func (h *ReadingsHandler) GetReadings(w http.ResponseWriter, r *http.Request) {
	resources, err := h.store.List()
	if err != nil {
		fmt.Println(err)
		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(resources)
	if err != nil {
		fmt.Println(err)
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *ReadingsHandler) GetLatestReadings(w http.ResponseWriter, r *http.Request) {
	resources, err := h.store.Latest()
	if err != nil {
		fmt.Println(err)
		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(resources)
	if err != nil {
		fmt.Println(err)
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *ReadingsHandler) SetSensorName(w http.ResponseWriter, r *http.Request) {
	match := r.PathValue("id")

	var requestModel SetSensorNameRequest
	if err := json.NewDecoder(r.Body).Decode(&requestModel); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	if requestModel.SensorName == "" {
		InternalServerErrorHandler(w, r)
		return
	}

	_, err := h.store.GetSensor(match)

	if err != nil {
		fmt.Print(err)
		NotFoundHandler(w, r)
		return
	}

	if err := h.store.SetSensorName(match, requestModel.SensorName); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not found"))
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal server error"))
}

func GetWebClient(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webclient/index.html")
}

func GetWebClientAsset(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webclient/"+r.URL.Path)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func main() {
	store := readings.NewReadingDbStore("./db.db")
	readingsHandler := NewReadingsHandler(store)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/reading", readingsHandler.GetReadings)
	mux.HandleFunc("GET /api/latest", readingsHandler.GetLatestReadings)
	mux.HandleFunc("POST /api/sensors/{id}/setName", readingsHandler.SetSensorName)
	mux.HandleFunc("GET /assets/", GetWebClientAsset)
	mux.HandleFunc("GET /{$}", GetWebClient)

	fmt.Println("Server starting")
	http.ListenAndServe("localhost:8000", corsMiddleware(mux))
}
