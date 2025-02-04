package handler

// handler bertugas untuk membaca HTTP Request, melakukan mapping menggunakan mapper, memanggil service, dan mengembalikan response

import (
	"barebone-go-crud/src/models/dto/request"
	requestmapper "barebone-go-crud/src/models/mapper/request_mapper"
	responsemapper "barebone-go-crud/src/models/mapper/response_mapper"
	"barebone-go-crud/src/services"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserHandler struct {
	service services.UserService
}

func NewHandleUser(s services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Url is not valid", http.StatusBadRequest)
		return
	}

	idStr := segments[2]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Id is not valid", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	user, err := h.service.GetUserById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}

	resp := responsemapper.MapUserModelToUserResponse(user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req request.UserRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request is not valid", http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user := requestmapper.MapUserRequestToUserModel(&req)

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	insertUser, err := h.service.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := responsemapper.MapUserModelToUserResponse(insertUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Url is not valid", http.StatusBadRequest)
		return
	}

	idStrs := segments[2]
	id, err := strconv.ParseInt(idStrs, 10, 64)
	if err != nil {
		http.Error(w, "Id is not valid", http.StatusBadRequest)
		return
	}

	var req request.UserRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Body is not valid", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user := requestmapper.MapUserRequestToUserModel(&req)
	user.Id = id

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	updatedUser, err := h.service.UpdateUser(ctx, id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := responsemapper.MapUserModelToUserResponse(updatedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Url is not valid", http.StatusBadRequest)
		return
	}

	idStrs := segments[2]
	id, err := strconv.ParseInt(idStrs, 10, 64)
	if err != nil {
		http.Error(w, "Id is not valid", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := h.service.DeleteUser(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
