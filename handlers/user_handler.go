package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"otus-social-network-service_gen_swagger/app_error"
	dto "otus-social-network-service_gen_swagger/dto/user"
	"otus-social-network-service_gen_swagger/repository"
	"otus-social-network-service_gen_swagger/service"
)

type UserHandler struct {
	rm *repository.RepositoryManager
}

func NewUserHandler(rm *repository.RepositoryManager) *UserHandler {
	return &UserHandler{rm: rm}
}

func (h *UserHandler) UserGetIdGet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userUUID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, httpError := h.rm.UserRepositoryInstance().GetUserById(userUUID)

	if httpError != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	jsonResp, err := json.Marshal(user)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func (h *UserHandler) UserRegisterPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// todo обработка ошибок
		return
	}
	userRegReq := new(dto.UserRegisterBody)
	err = json.Unmarshal(body, userRegReq)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	authServiceInstance := service.NewAuthService(h.rm)
	uuid, err := authServiceInstance.CreateUser(userRegReq)

	if err != nil {
		httpError, _ := err.(*app_error.HttpError)
		w.WriteHeader(httpError.Status())
		return
	}

	w.WriteHeader(http.StatusOK)

	resp := make(map[string]string)
	resp["user_id"] = uuid.String()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func (h *UserHandler) UserSearchGet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	firstName := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")

	users, httpError := h.rm.UserRepositoryInstance().FindUsers(firstName, last_name)

	if httpError != nil {
		jsonResp, _ := json.Marshal(httpError)
		w.Write(jsonResp)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	jsonResp, err := json.Marshal(users)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
