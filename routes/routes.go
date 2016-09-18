package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	s1 := "fasdfasfds00erd3未找到t3rfhaausdfasdjkfdsf"
	s2 := ""
	for i := 0; i < 10000; i++ {
		s2 += s1

	}
	io.WriteString(w, s2)
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.PostForm)
	io.WriteString(w, "post\n")
}

type Result struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

func HandleJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(Result{"tee", "dub"})
	io.WriteString(w, string(result))
}
