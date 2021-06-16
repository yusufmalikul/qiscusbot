package api

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "os"
    "strings"
)

type PayloadData struct {
    Payload struct {
        From struct {
            Id   int    `json:"id"`
            Name string `json:"name"`
        } `json:"from"`
        Message struct {
            Text string `json:"text"`
        } `json:"message"`
        Room struct {
            Id int `json:"id"`
        } `json:"room"`
    } `json:"payload"`
}

func Reply(w http.ResponseWriter, r *http.Request) {
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    getBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    log.Printf("%s", getBody)
    var payloadData PayloadData

    err = json.Unmarshal(getBody, &payloadData)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    postComment(payloadData.Payload.Room.Id)

}

func postComment(roomId int) {
    url := "https://api.qiscus.com/api/v2.1/rest/post_comment"
    method := "POST"

    comment := []string{
        "senang bertemu denganmu!",
        "saya sering lupa.",
        "bagaimana kabarmu?",
        "waktunya istirahat",
        "namaku qiscusbot :)",
    }

    payload := strings.NewReader(fmt.Sprintf(`{
      "user_id": "bot",
      "room_id": "%v",
      "message": "%s"
    }`, roomId, comment[rand.Intn(len(comment))]))

    client := &http.Client{
    }
    req, err := http.NewRequest(method, url, payload)

    if err != nil {
        fmt.Println(err)
        return
    }
    req.Header.Add("QISCUS-SDK-APP-ID", "qchatsdkd-fsdm3yu2oip")
    req.Header.Add("QISCUS_SDK_SECRET", os.Getenv("QISCUS_SDK_SECRET"))
    req.Header.Add("Content-Type", "application/json")

    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
}

func Home(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    _, err := w.Write([]byte("nothing to see here."))
    log.Println(err)
}
