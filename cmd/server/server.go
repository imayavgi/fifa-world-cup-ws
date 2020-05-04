package main

import (
	"net/http"

	"github.com/imayavgi/fifa-world-cup-ws/internal/pkg/data"
	"github.com/imayavgi/fifa-world-cup-ws/internal/pkg/handlers"
)

func main() {
	data.PrintUsage()

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/winners", handlers.WinnersHandler)
	http.ListenAndServe(":8000", nil)
}
