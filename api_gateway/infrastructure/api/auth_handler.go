package api

import (
	"context"
	"encoding/json"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/services"
	auth "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/auth_service"
	userInfo "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type AuthHandler struct {
	userInfoClientAdress string
	authClientAdress     string
}

func NewAuthHandler(userInfoClientAdress string, authClientAdress string) Handler {
	return &AuthHandler{
		userInfoClientAdress: userInfoClientAdress,
		authClientAdress:     authClientAdress,
	}
}

func (handler *AuthHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/auth/Register", handler.Register)
	if err != nil {
		panic(err)
	}
}

func (handler *AuthHandler) Register(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp1, err1 := handler.saveCredentials(&user)
	var userResp = domain.User{
		Name:     "",
		Surname:  "",
		Email:    resp1.User.Email,
		Password: "",
	}

	if resp1.User.Email == "error" {

		jsonBytes, err := json.Marshal(userResp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
	}

	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(userResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Id = resp1.User.Id
	_, err2 := handler.saveUserInfo(&user)

	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (handler *AuthHandler) saveCredentials(user *domain.User) (*auth.CreateUserResponse, error) {
	authClient := services.NewAuthClient(handler.authClientAdress)
	newUser := &auth.NewUser{
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
	response, err := authClient.CreateUser(context.TODO(), &auth.CreateUserRequest{NewUser: newUser})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (handler *AuthHandler) saveUserInfo(user *domain.User) (*userInfo.CreateUserResponse, error) {
	userInfoClient := services.NewUserClient(handler.userInfoClientAdress)
	newUser := &userInfo.NewUser{
		Id:      user.Id,
		Name:    user.Name,
		Surname: user.Surname,
	}
	response, err := userInfoClient.CreateUser(context.TODO(), &userInfo.CreateUserRequest{NewUser: newUser})
	if err != nil {
		return nil, err
	}
	return response, nil
}
