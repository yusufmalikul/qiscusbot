package api

import (
    "log"
    "net/http"
)

func Reply(w http.ResponseWriter, r *http.Request) {
    log.Println(r.Header)
    getBody, err := r.GetBody()
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    log.Println(getBody)
}

func Home(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, err := w.Write([]byte("nothing to see here."))
    log.Println(err)
}
