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
