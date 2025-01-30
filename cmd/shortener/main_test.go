package main

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	handlers "github.com/Arseniy-Polyakov/GoCourse/internal/controller/handler"
)

func Test_main(t *testing.T) {
	reqPost, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	reqGet, err := http.NewRequest("GET", "/{shortLink}", nil)
	if err != nil {
		t.Fatal(err)
	}

	rrPost := httptest.NewRecorder()
	rrGet := httptest.NewRecorder()
	handlerPost := http.HandlerFunc(handlers.HandlerPost)
	handlerGet := http.HandlerFunc(handlers.HandlerGet)

	handlerPost.ServeHTTP(rrPost, reqPost)
	if status := rrPost.Code; status != 201 {
		t.Errorf("Handler вернул неверный статус код для запроса POST: получен %v, ожидался 201",
			status)
	}

	handlerGet.ServeHTTP(rrGet, reqGet)
	if status := rrGet.Code; status != 307 {
		t.Errorf("Handler вернул неверный статус код для запроса GET: получен %v, ожидался 307",
			status)
	}

	matched, _ := regexp.MatchString(`[A-Za-z0-9]{8}`, rrPost.Body.String())
	if !matched {
		t.Error("Короткая ссылка не соответствует стандарту")
	}
}
