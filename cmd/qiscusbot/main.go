package main

import (
    "github.com/gorilla/mux"
    "github.com/yusufmalikul/qiscusbot/pkg/handler"
    "log"
    "net/http"
    "os"
)

type Env struct {
    Router *mux.Router
    Port   string
}

func main() {
    env := Setup()
    log.Println("Listening on ", env.Port)
    env.Run(":" + env.Port)
}

func Setup() Env {
    port := "8080" // default port

    // needed in heroku environment
    if os.Getenv("PORT") != "" {
        port = os.Getenv("PORT")
    }

    env := Env{
        Router: mux.NewRouter(),
        Port:   port,
    }

    return env
}

func (env *Env) Run(addr string) {
    handler.InitRoutes(env.Router)
    log.Fatal(http.ListenAndServe(addr, env.Router))
}
