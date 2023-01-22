package authentication_tests

import (
	"RegisterUser/internal/app/authentication"
	"RegisterUser/repository/user"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	// Создаем тестовый сервер
	ts := httptest.NewServer(http.HandlerFunc(authentication.Registration))
	defer ts.Close()

	// Создаем тестовый запрос, но тк email не должен повторятся, то меняйте название email для каждого теста
	testUser := user.User{NickName: "test", Email: "testmail888", Password: "testpassword"}
	jsonUser, err := json.Marshal(testUser)
	if err != nil {
		t.Fatal(err)
	}
	// Кодируем наше тело client в виде json и используем метод POST, проверяя ошибку
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
