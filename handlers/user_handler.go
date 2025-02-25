package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/faizkhan-06/go-auth/config"
	"github.com/faizkhan-06/go-auth/models"
	"github.com/faizkhan-06/go-auth/utils"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data interface{} `json:"data"`
}

func Register(w http.ResponseWriter, r *http.Request){
	var user models.User
	w.Header().Set("Content-Type","application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, `{"message": "Invalid json formate"}`, http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "Email and Passoword are required",
		})
		return
	}

	hashedPassword, _ := utils.GenerateHash(user.Password)

	user.Password = hashedPassword

	result := config.DB.Create(&user)
	
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "1062") {
			w.WriteHeader(http.StatusConflict) 
			json.NewEncoder(w).Encode(Response{
				Message: "User already exists",
				Status:  http.StatusConflict,
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Message: "Internal server error",
			Status: http.StatusInternalServerError,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Message: "User Registered",
		Status: http.StatusCreated,
		Data: user,
	})
}

func Login(w http.ResponseWriter, r *http.Request){

}