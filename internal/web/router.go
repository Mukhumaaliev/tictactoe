package web

import "net/http"

func NewRouter(handler *GameHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /game", handler.CreateGame)
	mux.HandleFunc("POST /game/{uuid}", handler.MakeMove)
	return mux
}
