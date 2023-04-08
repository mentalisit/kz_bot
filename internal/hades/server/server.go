package server

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"kz_bot/internal/models"
	"net/http"
	"sync"
)

type Server struct {
	toGame        []models.Message
	NewToGame     chan models.Message
	NewToMessager chan models.Message
	mu            sync.Mutex
}

func NewServer(togame chan models.Message, tomess chan models.Message) *Server {
	s := &Server{
		toGame:        []models.Message{},
		NewToGame:     togame,
		NewToMessager: tomess,
	}

	go s.inbox() // получаю сообщение с канала
	http.HandleFunc("/togame/", s.sendToGame)
	http.HandleFunc("/tomessager", s.sendToMessager)
	fmt.Println("Сервер загружен")
	go s.start()

	return s
}
func (s *Server) start() {
	err := http.ListenAndServe(fmt.Sprintf(":7777"), nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Server) sendToGame(w http.ResponseWriter, r *http.Request) {
	corpName := r.URL.Path[len("/togame/"):]
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(s.SortMessage(corpName))
		//s.toGame = []models.Message{}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func (s *Server) sendToMessager(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var m models.Message
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.NewToMessager <- m
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func (s *Server) SortMessage(CorpName string) []models.Message {
	//s.mu.Lock()
	if len(s.toGame) == 0 {
		return []models.Message{}
	}
	CurrentCorp := []models.Message{}
	OtherCorp := []models.Message{}
	for _, message := range s.toGame {
		if message.Corporation == CorpName {
			CurrentCorp = append(CurrentCorp, message)
		} else {
			OtherCorp = append(OtherCorp, message)
		}
	}
	s.toGame = OtherCorp
	//s.mu.Unlock()
	return CurrentCorp
}
