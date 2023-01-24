package user_menagment

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
)

// Get User from id
func (uc searchUserByIDController) getUserById() error {
	uc.rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(uc.req)

	// получаем введенный id
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		if err == primitive.ErrInvalidHex {
			uc.rw.WriteHeader(http.StatusLengthRequired)
			err = errors.New("invalid id")
			return err
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		return err
	}

	query := bson.M{"_id": id}

	// Проверяем, существует ли пользователь с таким id, если не существует, то выводим ошибку, если существует, то
	// отправляем данные пользователя
	if err = collection.FindOne(ctx, query).Decode(&uc.user); err != nil {
		if err == mongo.ErrNoDocuments {
			uc.rw.WriteHeader(http.StatusNoContent)
			return err
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		return err
	}

	err = json.NewDecoder(uc.req.Body).Decode(uc.user)
	if err != nil {
		if err == io.EOF {
			uc.rw.WriteHeader(http.StatusInternalServerError)
			return err
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		return err
	}

	err = json.NewEncoder(uc.rw).Encode(uc.user)
	if err != nil {
		uc.rw.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return err
}
