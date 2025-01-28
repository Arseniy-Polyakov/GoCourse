package main

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	links     = make(map[string]string)
	urlYandex = "https://practicum.yandex.ru"
)

func HandlerPost(w http.ResponseWriter, r *http.Request) {
	symbols := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	shortPart := []string{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		randomNumber := rand.Intn(len(symbols))
		randomSymbol := string(symbols[randomNumber])
		shortPart = append(shortPart, randomSymbol)
	}

	shortPartStr := strings.Join(shortPart, "")
	shortLink := "http://localhost:8080/" + shortPartStr

	links[shortPartStr] = urlYandex

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(shortLink))
}

func HandlerGet(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Path[len("/"):]
	fullLink := links[shortLink]
	http.Redirect(w, r, fullLink, http.StatusTemporaryRedirect)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", HandlerPost)
	r.Get("/{shortLink}", HandlerGet)
	http.ListenAndServe(":8080", r)
}
