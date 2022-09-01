package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes(handl *http.ServeMux) {
	handl.HandleFunc("/", h.root)
}

func (h *Handler) root(w http.ResponseWriter, r *http.Request) {

	user, err := h.service.GetUser(r.Context())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error.."))
		fmt.Printf("err occured. %s\n", err.Error())
		return
	}

	if user == nil {
		w.WriteHeader(404)
		w.Write([]byte("not found..."))
		return
	}

	msg, err := json.Marshal(*user)
	if err != nil {
		fmt.Printf("err occured. %s\n", err.Error())
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(msg)
	return
}
