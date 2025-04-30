package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	logrus.Info("Starting server on :34567")
	if err := http.ListenAndServe(":34567", router); err != nil {
		logrus.Error("Error starting server:", err)
	}
}
