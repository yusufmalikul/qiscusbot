package api

import (
    "io/ioutil"
    "log"
    "net/http"
)

func Reply(w http.ResponseWriter, r *http.Request) {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    getBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    log.Printf("%s", getBody)
}

func Home(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, err := w.Write([]byte("nothing to see here."))
    log.Println(err)
}
