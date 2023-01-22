package user_menagment

import (
	"RegisterUser/repository/user"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
)

// Get User from id
func (uc searchUserByIDController) getUserById(getUser *user.User) error {
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
		log.Panic(err)
	}

	query := bson.M{"_id": id}

	// Проверяем, существует ли пользователь с таким id, если не существует, то выводим ошибку, если существует, то
	// отправляем данные пользователя
	if err = collection.FindOne(ctx, query).Decode(&getUser); err != nil {
		if err == mongo.ErrNoDocuments {
			uc.rw.WriteHeader(http.StatusNoContent)
			return err
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
	}

	err = json.NewDecoder(uc.req.Body).Decode(getUser)
	if err != nil {
		if err == io.EOF {
			uc.rw.WriteHeader(http.StatusInternalServerError)
			log.Panic(err)
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
	}

	err = json.NewEncoder(uc.rw).Encode(getUser)
	if err != nil {
		uc.rw.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
	}

	return err
}
