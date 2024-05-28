package handlers

import (
	"net/http"
	"otus-social-network-service_gen_swagger/repository"
)

type FriendHandler struct {
	rm *repository.RepositoryManager
}

func NewFriendHandler(rm *repository.RepositoryManager) *FriendHandler {
	return &FriendHandler{rm: rm}
}

func (h *FriendHandler) FriendDeleteUserIdPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *FriendHandler) FriendSetUserIdPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
