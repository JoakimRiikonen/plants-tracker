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

func (h *ReadingsHandler) GetNewestReadings(w http.ResponseWriter, r *http.Request) {
	resources, err := h.store.Newest()
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

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal server error"))
}

func main() {
	store := readings.NewReadingDbStore("./db.db")
	readingsHandler := NewReadingsHandler(store)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /reading", readingsHandler.GetReadings)
	mux.HandleFunc("GET /newest", readingsHandler.GetNewestReadings)

	fmt.Println("Server starting")
	http.ListenAndServe("localhost:8000", mux)
}
