package user_controllers

import (
	"RegisterUser/repository/user"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	// Создаем тестовый сервер
	ts := httptest.NewServer(http.HandlerFunc(RegisterUser))
	defer ts.Close()

	// Создаем тестовый запрос
	user := user.User{NickName: "test", Email: "teswrwr2", Password: "testpassword"}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}

	// Отправляем запрос и получаем ответ
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Проверяем, статус кода
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, res.StatusCode)
	}
}
