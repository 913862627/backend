package routes

import (
	"net/http"
	"webt/utils"
)

func init() {
	http.HandleFunc("/", HandleIndex)

	// private views
	http.HandleFunc("/post", utils.PostOnly(utils.BasicAuth(HandlePost)))
	http.HandleFunc("/json", utils.GetOnly(utils.BasicAuth(HandleJSON)))
}
