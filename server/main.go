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
	mux.HandleFunc("GET /assets/", GetWebClientAsset)
	mux.HandleFunc("GET /{$}", GetWebClient)

	fmt.Println("Server starting")
	http.ListenAndServe("0.0.0.0:8000", corsMiddleware(mux))
}
