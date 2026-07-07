package server

import (
	"encoding/json"
	"net/http"

	"strings"

	"github.com/abrshDev/raft-kv/internal/kvstore"
)

type Server struct {
	store *kvstore.Store
}

func New(store *kvstore.Store) *Server {
	return &Server{store: store}
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/kv/", s.handleKV)
}

func (s *Server) handleKV(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/kv/")
	if key == "" {
		http.Error(w, "key required", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		val, ok := s.store.Get(key)
		if !ok {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"key": key, "value": val})
	case http.MethodPut:
		var body struct{ Value string }
		json.NewDecoder(r.Body).Decode(&body)
		s.store.Set(key, body.Value)
		w.WriteHeader(http.StatusNoContent)
	case http.MethodDelete:
		s.store.Delete(key)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
