package producer

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	log    *log.Logger
}

func HandleRequests() *Server {
	s := &Server{
		log:    log.Default(),
		router: mux.NewRouter(),
	}

	s.router.HandleFunc("/", s.defaultResponse)
	s.router.HandleFunc("/node", s.defaultResponse).Methods("GET")
	s.router.HandleFunc("/nodes", s.getAllNodes)
	s.router.HandleFunc("/node/{id}", s.getNodeById).Methods("GET")
	s.router.HandleFunc("/node", s.addNewNode).Methods("POST")
	s.router.HandleFunc("/node/{id}", s.deleteNode).Methods("DELETE")
	s.router.HandleFunc("/node/{id}", s.updateNode).Methods("PUT")

	Nodes = []Node{
		{Id: 1, Name: "Node 1", Location: "Athlone", Latitude: 123, Longitude: 321},
		{Id: 2, Name: "Node 2", Location: "Cork", Latitude: 345, Longitude: 675},
	}

	return s
}

func (s *Server) ListenAndServe(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.router))
}
