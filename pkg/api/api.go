package api

import (
    "log"
    "net/http"
)

func Reply(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, err := w.Write([]byte("it works!"))
    log.Println(err)
}

func Home(w http.ResponseWriter, r *http.Request)  {
    w.WriteHeader(http.StatusOK)
    _, err := w.Write([]byte("nothing to see here."))
    log.Println(err)
}
