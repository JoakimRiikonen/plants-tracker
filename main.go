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

func (h *ReadingsHandler) AddReading(w http.ResponseWriter, r *http.Request) {
	var reading readings.Reading
	if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.store.Add(reading); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ReadingsHandler) GetReadings(w http.ResponseWriter, r *http.Request) {
	resources, err := h.store.List()

	jsonBytes, err := json.Marshal(resources)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal server error"))
}

func main() {
	store := readings.NewReadingDbStore("./db.db")
	readingsHandler := NewReadingsHandler(store)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /reading", readingsHandler.AddReading)
	mux.HandleFunc("GET /reading", readingsHandler.GetReadings)

	fmt.Println("Server starting")
	http.ListenAndServe("localhost:8000", mux)
}
