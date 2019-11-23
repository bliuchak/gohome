package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

var (
	name = "my-awesome-app"
	version = "dev"
	commitFull = "n/a"
	buildTime = "n/a"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", HomeHandler)
	r.Get("/v1.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/v1.yaml")
	})
	fs := http.StripPrefix("/swaggerui", http.FileServer(http.Dir("dist")))
	r.Get("/swaggerui/*", fs.ServeHTTP)

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}

type homeResponse struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	CommitFull  string `json:"commit_full"`
	BuildTime   string `json:"build_time"`
	SwaggerUI   string `json:"swagger_ui"`
	SwaggerYAML string `json:"swagger_yaml"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp := homeResponse{
		Name:        name,
		Version:     version,
		CommitFull:  commitFull,
		BuildTime:   buildTime,
		SwaggerUI:   "http://" + r.Host + "/swaggerui/",
		SwaggerYAML: "http://" + r.Host + "/v1.yaml",
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
