package main
import (
        "net/http"
        "net/http/httptest"
        "strings"
        "testing"

        "github.com/gorilla/mux"
)


func TestStoreMsg(t *testing.T) {
    r := mux.NewRouter()
    r.HandleFunc("/messages", PostMsgHandler).Methods("POST")
    req, err := http.NewRequest(
            "POST",
            "/messages",
            strings.NewReader(`Hello World`),
    )
    if err != nil {
            t.Error(err)
    }

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    if w.Code != 201 {
            t.Errorf("HTTP Status expected: 201, got: %d", w.Code)
    }
}

func TestStoreTooLongMsg(t *testing.T) {
    r := mux.NewRouter()
    r.HandleFunc("/messages", PostMsgHandler).Methods("POST")
    byteSlice := []byte{0:'a',5000:'b'}
    tooLongString:= string(byteSlice[:])

    req, err := http.NewRequest(
            "POST",
            "/messages",
            strings.NewReader(tooLongString),
    )
    if err != nil {
            t.Error(err)
    }

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    if w.Code != 500 {
            t.Errorf("HTTP Status expected: 500, got: %d", w.Code)
    }
}


func TestGetMsgHandler(t *testing.T) {
    r := mux.NewRouter()
    r.HandleFunc("/messages/{id}", GetMsgHandler).Methods("GET")
    req, err := http.NewRequest("GET", "/messages/1", nil)
    if err != nil {
            t.Error(err)
    }
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    if w.Code != 200 {
            t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
    }
}

func TestGetNonExistentMsgHandler(t *testing.T) {
    r := mux.NewRouter()
    r.HandleFunc("/messages/{id}", GetMsgHandler).Methods("GET")
    req, err := http.NewRequest("GET", "/messages/31", nil)
    if err != nil {
            t.Error(err)
    }
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    if w.Code != 404 {
            t.Errorf("HTTP Status expected: 404, got: %d", w.Code)
    }
}


/*
func TestGetMsgHandlerWithServer(t *testing.T) {
    log.SetPrefix("TestGetMsgHandlerWithServer:")
    log.Println ("Testing message retrieval")
    r := mux.NewRouter()
    r.HandleFunc("/messages/{id}", GetMsgHandler).Methods("GET")
    server := httptest.NewServer(r)
    defer server.Close()
    msgUrl := fmt.Sprintf("%s/messages/1", server.URL)
    log.Printf ("msgUrl %s", msgUrl)
    request, err := http.NewRequest("GET", msgUrl, nil)

    res, err := http.DefaultClient.Do(request)
    if err != nil {
        t.Error(err)
    }

    resDump, err2 := ioutil.ReadAll(res.Body)
    
    
    if err2 != nil {
        t.Error(err2)
    }


    log.Printf("Response is [%s]",string(resDump[:]))
    if res.StatusCode != 200 {
        t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
    }
}*/