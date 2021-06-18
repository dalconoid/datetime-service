package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// server time delta (can be changed with [/time/correct])
var delta time.Duration = 0

// Server - server struct
type Server struct {
	router *mux.Router
}

// New - creates and returns new server
func New() *Server {
	return &Server{router: mux.NewRouter()}
}

// Start - starts server on specified address
func (s *Server) Start(address string) error{
	return http.ListenAndServe(address, s.router)
}

// ConfigureRouter - binds handles to routes
func (s *Server) ConfigureRouter() {
	s.router.Handle("/alive", handleAlive()).Methods("GET")
	s.router.Handle("/time/now", handleGetTimeNow()).Methods("GET")
	s.router.Handle("/time/string", handleGetTimeString()).Methods("GET")
	s.router.Handle("/time/add", handleAddTime()).Methods("GET")
	s.router.Handle("/time/correct", handleCorrectServerTime()).Methods("POST")
}