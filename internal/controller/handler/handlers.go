package handlers

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/Arseniy-Polyakov/GoCourse/cmd/config"
)

var (
	links     = make(map[string]string)
	urlYandex = "https://practicum.yandex.ru"
	fullLink  string
	shortLink string
)

func HandlerPost(w http.ResponseWriter, r *http.Request) {
	cfg := config.NewConfig()
	symbols := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	shortPart := []string{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		randomNumber := rand.Intn(len(symbols))
		randomSymbol := string(symbols[randomNumber])
		shortPart = append(shortPart, randomSymbol)
	}

	shortPartStr := strings.Join(shortPart, "")
	shortLink = "http://localhost:" + cfg.BaseURL + "/" + shortPartStr

	links[shortPartStr] = urlYandex

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(shortLink))
}

func HandlerGet(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Path[len("/"):]
	fullLink = links[shortLink]
	http.Redirect(w, r, fullLink, http.StatusTemporaryRedirect)
}
