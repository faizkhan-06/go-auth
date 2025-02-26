package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/faizkhan-06/go-auth/config"
	"github.com/faizkhan-06/go-auth/models"
	"github.com/faizkhan-06/go-auth/utils"
	"gorm.io/gorm"
)


func Register(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, `{"message": "Invalid json formate"}`, http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Message: "Email and Passoword are required",
		})
		return
	}

	hashedPassword, _ := utils.GenerateHash(user.Password)

	user.Password = hashedPassword

	result := config.DB.Create(&user)
	
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Message: result.Error.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.Response{
		Message: "User Registered",
		Status: http.StatusCreated,
		Data: user,
	})
}

func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user models.User
	var storedUser models.User

	 err := json.NewDecoder(r.Body).Decode(&user)
	 if err != nil{
		http.Error(w, `{"message": "Invalid json formate"}`, http.StatusBadRequest)
		return
	 }

	 if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status: http.StatusBadRequest,
			Message: "Email and password both are required",
		})
		return
	}


	result := config.DB.Where("email = ?",user.Email).First(&storedUser)

	if result.Error != nil{
		if result.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.Response{
				Status: http.StatusNotFound,
				Message: "User not found",
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.Response{
				Status: http.StatusInternalServerError,
				Message: result.Error.Error(),
			})
			return
	}

	if !utils.CompareHashAndPassword(storedUser.Password, user.Password){
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(utils.Response{
			Status: http.StatusUnauthorized,
			Message: "Password does not matched",
		})
		return
	}

	token, err := utils.GenerateJWTToken(storedUser.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Message: "Jwt token generator faild",
			Status: http.StatusInternalServerError,
		})
		return
	}

	responseUser := map[string]interface{}{
		"id": storedUser.ID,
		"email": storedUser.Email,
		"createdAt" : storedUser.CreatedAt,
		"updatedAt" : storedUser.UpdatedAt,
		"deletedAT" : storedUser.DeletedAt,
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{
		Message: "You're logged in",
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"user": responseUser,
			"token": token,
		},
	})
	
}


func Home (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{
		Message: "Response OK",
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"text": "Hello Friend",
		},
	})
}
