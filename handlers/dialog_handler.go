package handlers

import (
	"net/http"
	"otus-social-network-service_gen_swagger/repository"
)

type DialogHandler struct {
	rm *repository.RepositoryManager
}

func NewDialogHandler(rm *repository.RepositoryManager) *DialogHandler {
	return &DialogHandler{rm: rm}
}

func (h *DialogHandler) DialogUserIdListGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *DialogHandler) DialogUserIdSendPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
