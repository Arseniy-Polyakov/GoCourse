package main

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	links     = make(map[string]string)
	urlYandex = "https://practicum.yandex.ru"
)

func HandlerPost(w http.ResponseWriter, r *http.Request) {
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
	// 	return
	// }
	// url := string(body)

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
	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte(fullLink))
	http.Redirect(w, r, fullLink, http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", HandlerPost)
	http.HandleFunc("/{shortLink}", HandlerGet)
	http.ListenAndServe(":8080", nil)
}
