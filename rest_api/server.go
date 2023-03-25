package rest_api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter((prometheus.CounterOpts{
		Name: "message_service_ops_total",
		Help: "Number of events processed",
	}))
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
	s.router.Path("/metrics").Handler(promhttp.Handler())

	Nodes = []Node{
		{Id: 1, Name: "Node 1", Location: "Athlone", Latitude: 123, Longitude: 321},
		{Id: 2, Name: "Node 2", Location: "Cork", Latitude: 345, Longitude: 675},
	}

	return s
}

func (s *Server) ListenAndServe(addr string) {
	recordMetrics()
	log.Fatal(http.ListenAndServe(addr, s.router))
}
