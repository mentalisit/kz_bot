package hserver

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Httpserver struct {
	name string
	text string
}

type Response struct {
	Status  int
	Service string
}

func (h *Httpserver) NewServer(name, text string) {
	h.name = name
	h.text = text
	mux := http.NewServeMux()
	mux.HandleFunc("/"+h.name, h.health)
	s := http.Server{
		Addr:         ":8888",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Запуск сервера на %s", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
func (h *Httpserver) health(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Content-Type", "application/json")
	(w).WriteHeader(http.StatusOK)
	dataBytes, _ := json.Marshal(h.text) //Response{Status: http.StatusOK, Service: "health"})
	(w).Write(dataBytes)
}
