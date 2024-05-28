package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	dto "otus-social-network-service_gen_swagger/dto/login"
	"otus-social-network-service_gen_swagger/repository"
	"otus-social-network-service_gen_swagger/service"
)

type LoginHandler struct {
	rm *repository.RepositoryManager
}

func NewLoginHandler(rm *repository.RepositoryManager) *LoginHandler {
	return &LoginHandler{rm: rm}
}

func (h *LoginHandler) LoginPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// todo обработка ошибок
		return
	}
	loginReq := new(dto.LoginBody)
	err = json.Unmarshal(body, loginReq)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	authServiceInstance := service.NewAuthService(h.rm)
	token, err := authServiceInstance.GenerateToken(loginReq.Email, loginReq.Password)

	// token, err := h.rm.LoginRepositoryInstance().GenerateToken(loginReq.Email, loginReq.Password)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	resp := make(map[string]string)
	resp["token"] = token
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
