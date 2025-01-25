package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	links     = make(map[string]string)
	urlYandex = "https://practicum.yandex.ru"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
			return
		}
		url := string(body)

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
		fmt.Println(url)
		fmt.Println(http.StatusCreated)
		fmt.Println(shortLink)
		fmt.Println(links)
	})

	http.HandleFunc("/{shortlink}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "The method GET is not allowed", http.StatusMethodNotAllowed)
			return
		}
		shortLink := r.URL.Path[len("/"):]
		fullLink := links[shortLink]
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fullLink))
		fmt.Println(fullLink)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print("Server Error")
	} else {
		fmt.Println("OK")
	}
}
