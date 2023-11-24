package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getAllStudent(w http.ResponseWriter, req *http.Request) {
	students := []Student{
		{Id: 1, Name: "Nguyễn Văn A"},
		{Id: 2, Name: "Trịnh Văn B"},
		{Id: 3, Name: "Ngô Thị C"},
	}
	studentsJson, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(studentsJson)
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Host)
		fmt.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/students", getAllStudent).Methods("GET")
	router.Use(loggingMiddleware)
	port := "8080"
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	s := fmt.Sprintf("Server listen on:http://localhost:%s", port)
	fmt.Println(s)
	srv.ListenAndServe()
}
