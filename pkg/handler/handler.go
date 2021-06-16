package handler

import (
    "github.com/gorilla/mux"
    "github.com/yusufmalikul/qiscusbot/pkg/api"
)

func InitRoutes(router *mux.Router) {
    router.HandleFunc("/", api.Home).Methods("GET")
    router.HandleFunc("/webhook", api.Reply).Methods("POST")
}
