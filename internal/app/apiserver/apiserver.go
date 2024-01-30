package apiserver

import (
	"github.com/gorilla/mux"

	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	router *mux.Router
}

// New config
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

// Start server
func (server *APIServer) Start() error {
	server.configureRouter()
	return http.ListenAndServe(server.config.BindAddr, server.router)
}

// Configure router
func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/hello", server.handleHello())
}

func (server *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "Hello"); err != nil {

		}
	}
}
