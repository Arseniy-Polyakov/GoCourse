package main

import (
	"flag"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	links     = make(map[string]string)
	urlYandex = "https://practicum.yandex.ru"
	fullLink  string
	shortLink string
	baseURL   string
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

	if _, ok := os.LookupEnv("BASE_URL"); ok {
		baseURL = os.Getenv("BASE_URL")
		shortLink = "http://localhost:" + baseURL + "/" + shortPartStr
	} else {
		baseURL := flag.String("b", ":8080", "BASE_URL")
		flag.Parse()
		if baseURL != nil {
			shortLink = "http://localhost:" + *baseURL + "/" + shortPartStr
		} else {
			shortLink = "http://localhost:8080/" + shortPartStr
		}
	}

	links[shortPartStr] = urlYandex

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(shortLink))
}

func HandlerGet(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Path[len("/"):]
	fullLink = links[shortLink]
	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte(fullLink))
	http.Redirect(w, r, fullLink, http.StatusTemporaryRedirect)
}

func main() {
	if _, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		port := os.Getenv("SERVER_ADDRESS")
		http.HandleFunc("/", HandlerPost)
		http.HandleFunc("/{shortLink}", HandlerGet)
		http.ListenAndServe(":"+port, nil)
	} else {
		port := flag.String("a", ":8080", "SERVER_ADDRESS")
		flag.Parse()
		if port != nil {
			http.HandleFunc("/", HandlerPost)
			http.HandleFunc("/{shortLink}", HandlerGet)
			http.ListenAndServe(*port, nil)
		} else {
			http.HandleFunc("/", HandlerPost)
			http.HandleFunc("/{shortLink}", HandlerGet)
			http.ListenAndServe(":8080", nil)
		}
	}
}

// else {
// 	port := flag.String("a", ":8080", "Server Address")
// 	flag.Parse()
// 	http.ListenAndServe(port, nil)
// }
// flag.StringVar(&port, "a", ":8080", "Server Address")
// flag.Parse()
// r := chi.NewRouter()
// r.Use(middleware.Recoverer)
// r.Post("/", HandlerPost)
// r.Get("/{shortLink}", HandlerGet)

// package main

// import (
// 	"math/rand"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/go-chi/chi/v5"
// )

// var (
// 	links     = make(map[string]string)
// 	urlYandex = "https://practicum.yandex.ru"
// )

// func HandlerPost(w http.ResponseWriter, r *http.Request) {
// 	// body, err := io.ReadAll(r.Body)
// 	// if err != nil {
// 	// 	http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
// 	// 	return
// 	// }
// 	// url := string(body)

// 	symbols := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
// 	shortPart := []string{}

// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 8; i++ {
// 		randomNumber := rand.Intn(len(symbols))
// 		randomSymbol := string(symbols[randomNumber])
// 		shortPart = append(shortPart, randomSymbol)
// 	}

// 	shortPartStr := strings.Join(shortPart, "")
// 	shortLink := "http://localhost:8080/" + shortPartStr

// 	links[shortPartStr] = urlYandex

// 	w.WriteHeader(http.StatusCreated)
// 	w.Header().Set("Content-Type", "text/plain")
// 	w.Write([]byte(shortLink))
// }

// func HandlerGet(w http.ResponseWriter, r *http.Request) {
// 	shortLink := r.URL.Path[len("/"):]
// 	fullLink := links[shortLink]
// 	// w.WriteHeader(http.StatusCreated)
// 	// w.Write([]byte(fullLink))
// 	http.Redirect(w, r, fullLink, http.StatusTemporaryRedirect)
// }

// func main() {
// 	r := chi.NewRouter()
// 	r.Route("/", func(r chi.Router) {
// 		r.Get("/{shortLink}", HandlerGet)
// 		r.Post("/", HandlerPost)
// 	})
// 	http.ListenAndServe(":8080", nil)
// }
