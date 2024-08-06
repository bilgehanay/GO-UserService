package main

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid request payload", "", nil)
		return
	}
	result, err := userDB.UpdateOne(ctx, bson.M{"email": bson.M{"$ne": newUser.Email}}, bson.M{"$set": newUser}, options.Update().SetUpsert(true))
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Internal Server Error", "", nil)
		return
	}

	if result.MatchedCount == 0 {
		JsonResponse(w, http.StatusBadRequest, "Email is exists, try to login", "", nil)
		return
	}
	JsonResponse(w, http.StatusOK, "OK", "", nil)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser User

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid request payload", "", nil)
		return
	}
	var user User
	err := userDB.FindOne(ctx, bson.M{"email": loginUser.Email, "password": loginUser.Password}).Decode(&user)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, "User not found", "", nil)
		return
	}
	token, err := GenerateToken(user.ID)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Internal Server Error", "", nil)
		return
	}
	JsonResponse(w, http.StatusOK, "OK", token, user)
	return
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == "" {
		JsonResponse(w, http.StatusInternalServerError, "Internal Server Error", "", nil)
		return
	}
	var user User
	err := userDB.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Internal Server Error", "", nil)
		return
	}
	JsonResponse(w, http.StatusOK, "OK", "", user)
	return
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user User
	id := r.Context().Value("id")
	if id == "" {
		JsonResponse(w, http.StatusInternalServerError, "Internal Server Error", "", nil)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid request payload", "", nil)
		return
	}
	_, err := userDB.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user}, options.Update().SetUpsert(true))
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Internal Server Error", "", nil)
		return
	}
	JsonResponse(w, http.StatusOK, "OK", "", nil)
	return
}
