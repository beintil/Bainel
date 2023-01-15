package user_controllers

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

type service interface {
	userService() error
}

type registerUserController struct {
	user user.User
	rw   http.ResponseWriter
	req  *http.Request
}

type getUserByIDController struct {
	users *user.User
	rw    http.ResponseWriter
	req   *http.Request
}

func (uc registerUserController) userService() error {
	uc.rw.Header().Set("Content-Type", "application/json")
	var err error

	err = json.NewDecoder(uc.req.Body).Decode(&uc.user)
	if err != nil {
		if err == io.EOF {
			uc.rw.WriteHeader(http.StatusBadRequest)
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
	}

	var existingUser user.User
	err = collection.FindOne(ctx, bson.M{"email": uc.user.Email}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {

		}
	} else {
		http.Error(uc.rw, "User with this email already exists", http.StatusConflict)
		return err
	}

	switch {
	case uc.user.NickName == "":
		uc.rw.WriteHeader(http.StatusLengthRequired)
		err = errors.New("invalid userNickName")
		return err

	case uc.user.Email == "":
		uc.rw.WriteHeader(http.StatusLengthRequired)
		err = errors.New("invalid userEmail")
		return err

	case uc.user.Password == "":
		uc.rw.WriteHeader(http.StatusLengthRequired)
		err = errors.New("invalid userPassword")
		return err
	}

	var result, resultErr = collection.InsertOne(ctx, uc.user)
	if resultErr != nil {
		if err == mongo.ErrNoDocuments {
			uc.rw.WriteHeader(http.StatusNotFound)
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		return err
	}

	uc.rw.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(uc.rw).Encode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			uc.rw.WriteHeader(http.StatusNotFound)
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
	}

	return err
}

func (uc getUserByIDController) userService() error {
	uc.rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(uc.req)

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

	if err = collection.FindOne(ctx, query).Decode(&uc.users); err != nil {
		if err == mongo.ErrNoDocuments {
			uc.rw.WriteHeader(http.StatusNoContent)
			return err
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
	}

	err = json.NewDecoder(uc.req.Body).Decode(uc.users)
	if err != nil {
		if err == io.EOF {
			uc.rw.WriteHeader(http.StatusInternalServerError)
			log.Panic(err)
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
	}

	err = json.NewEncoder(uc.rw).Encode(uc.users)
	if err != nil {
		uc.rw.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
	}

	return err
}
