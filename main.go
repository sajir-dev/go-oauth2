package main

import (
	"net/http"
	"oauth-2-youtube/controllers"
)

func main() {
	http.HandleFunc("/google/login", controllers.GoogleLogin)
	http.HandleFunc("/google/callback", controllers.GoogleValidityCallback)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`pong`))
	})
	http.ListenAndServe(":3080", nil)
}
