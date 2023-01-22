package authentication

import (
	"Bainel/repository/user"
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
)

// Registration
func (uc registerController) register(registerUser user.User) error {
	uc.rw.Header().Set("Content-Type", "application/json")
	var err error

	err = json.NewDecoder(uc.req.Body).Decode(&registerUser)
	if err != nil {
		if err == io.EOF {
			uc.rw.WriteHeader(http.StatusBadRequest)
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
	}

	// Проверяем, существует ли введенный email в бд, если да, то отправляем ошибку, если нет, то мы ничего не делаем
	var existingUser user.User
	err = collection.FindOne(ctx, bson.M{"email": registerUser.Email}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Такого email не существует, значит можно регистрироваться

			// Проверяем, не введены ли пользователем все поля, чтобы избежать пустых значений в дб
			switch {
			case registerUser.NickName == "":
				uc.rw.WriteHeader(http.StatusLengthRequired)
				err = errors.New("invalid userNickName")
				return err

			case registerUser.Email == "":
				uc.rw.WriteHeader(http.StatusLengthRequired)
				err = errors.New("invalid userEmail")
				return err

			case registerUser.Password == "":
				uc.rw.WriteHeader(http.StatusLengthRequired)
				err = errors.New("invalid userPassword")
				return err
			}
		}
	} else {
		http.Error(uc.rw, "User with this email already exists", http.StatusConflict)
		return err
	}

	// Создаем/Регистрируем пользователя
	var result, resultErr = collection.InsertOne(ctx, registerUser)
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

// Authorisation
func (uc authorizationController) authorization(user user.User, login user.Login) error {
	err = json.NewDecoder(uc.req.Body).Decode(&login)
	if err != nil {
		if err == io.EOF {
			uc.rw.WriteHeader(http.StatusBadRequest)
			return err
		}
		uc.rw.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
	}

	if err = collection.FindOne(context.TODO(), bson.M{"email": login.Email}).Decode(&user); err != nil {
		err = errors.New("incorrect email")
		return err
	}

	if user.Password != login.Password {
		err = errors.New("incorrect password")
		return err
	}
	uc.rw.WriteHeader(http.StatusOK)
	_, err = uc.rw.Write([]byte("login ok"))
	if err != nil {
		log.Panic(err)
	}

	return err
}
