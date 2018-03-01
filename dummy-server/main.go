package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, genRandomString(responseSize))
}

var responseSize int

func main() {
	responseSizePtr := flag.Int("n", 200, "number of calls to execute")
	flag.Parse()

	responseSize = *responseSizePtr
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
