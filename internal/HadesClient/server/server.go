package server

//
//type Server struct {
//	toGame        []models.MessageHadesClient
//	NewToGame     chan models.MessageHadesClient
//	NewToMessager chan models.MessageHadesClient
//	mu            sync.Mutex
//}
//
//func NewServer(togames chan models.MessageHadesClient, tomess chan models.MessageHadesClient) *Server {
//	s := &Server{
//		toGame:        []models.MessageHadesClient{},
//		NewToGame:     togames,
//		NewToMessager: tomess,
//	}
//
//	go s.inbox() // получаю сообщение с канала
//	http.HandleFunc("/togame/", s.sendToGame)
//	http.HandleFunc("/tomessager", s.sendToMessager)
//	fmt.Println("Сервер загружен")
//	go s.start()
//
//	return s
//}
//func (s *Server) start() {
//	err := http.ListenAndServe(fmt.Sprintf(":8888"), nil)
//	if err != nil {
//		fmt.Println("start", err)
//		return
//	}
//}
//
//func (s *Server) sendToGame(w http.ResponseWriter, r *http.Request) {
//	corpName := r.URL.Path[len("/togame/"):]
//	if r.Method == http.MethodGet {
//		json.NewEncoder(w).Encode(s.sortMessage(corpName))
//		//s.toGame = []models.Message{}
//	} else {
//		w.WriteHeader(http.StatusMethodNotAllowed)
//	}
//}
//func (s *Server) sendToMessager(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodPost {
//		var m models.MessageHadesClient
//		err := json.NewDecoder(r.Body).Decode(&m)
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		s.NewToMessager <- m
//		w.WriteHeader(http.StatusOK)
//	} else {
//		w.WriteHeader(http.StatusMethodNotAllowed)
//	}
//}
//func (s *Server) sortMessage(CorpName string) []models.MessageHadesClient {
//	//s.mu.Lock()
//	if len(s.toGame) == 0 {
//		return []models.MessageHadesClient{}
//	}
//	CurrentCorp := []models.MessageHadesClient{}
//	OtherCorp := []models.MessageHadesClient{}
//	for _, message := range s.toGame {
//		if message.Corporation == CorpName {
//			CurrentCorp = append(CurrentCorp, message)
//		} else {
//			OtherCorp = append(OtherCorp, message)
//		}
//	}
//	s.toGame = OtherCorp
//	//s.mu.Unlock()
//	return CurrentCorp
//}
