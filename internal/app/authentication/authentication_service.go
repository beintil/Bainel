package authentication

import (
	"Bainel/repository/user"
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
)

// Registration
func (uc registerController) register() error {
	uc.rw.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(uc.req.Body).Decode(&uc.registerUser)
	if err != nil {
		if err == io.EOF {
			uc.rw.WriteHeader(http.StatusBadRequest)
			return err
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		return err
	}

	// Проверяем, существует ли введенный email в бд, если да, то отправляем ошибку, если нет, то мы ничего не делаем
	var existingUser user.User
	err = collection.FindOne(ctx, bson.M{"email": uc.registerUser.Email}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Такого email не существует, значит можно регистрироваться

			// Проверяем, не введены ли пользователем все поля, чтобы избежать пустых значений в дб
			switch {
			case uc.registerUser.NickName == "":
				uc.rw.WriteHeader(http.StatusLengthRequired)
				err = errors.New("invalid userNickName")
				return err
			case uc.registerUser.Email == "":
				uc.rw.WriteHeader(http.StatusLengthRequired)
				err = errors.New("invalid userEmail")
				return err
			case uc.registerUser.Password == "":
				uc.rw.WriteHeader(http.StatusLengthRequired)
				err = errors.New("invalid userPassword")
				return err
			}
		}
	} else {
		http.Error(uc.rw, "User with this email already exists", http.StatusConflict)
		return err
	}

	// Create new user
	var result, resultErr = collection.InsertOne(ctx, uc.registerUser)
	if resultErr != nil {
		if err == mongo.ErrNoDocuments {
			uc.rw.WriteHeader(http.StatusNotFound)
			return err
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		return err
	}

	uc.rw.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(uc.rw).Encode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			uc.rw.WriteHeader(http.StatusNotFound)
			return err
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return err
}

// Authorisation
func (uc authorizationController) authorization() error {
	err := json.NewDecoder(uc.req.Body).Decode(&uc.login)
	if err != nil {
		if err == io.EOF {
			uc.rw.WriteHeader(http.StatusBadRequest)
			return err
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		return err
	}

	if err = collection.FindOne(context.TODO(), bson.M{"email": uc.login.Email}).Decode(&uc.user); err != nil {
		err = errors.New("incorrect email")
		return err
	}

	if uc.user.Password != uc.login.Password {
		err = errors.New("incorrect password")
		return err
	}
	uc.rw.WriteHeader(http.StatusOK)
	_, err = uc.rw.Write([]byte("uc.login ok"))
	if err != nil {
		return err
	}

	return err
}
