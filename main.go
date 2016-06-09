package main
import (
    "encoding/json"
    "log"
    "net/http"
	"io/ioutil"    
    "strconv"
    "fmt"


    "github.com/gorilla/mux"
)


//struct used to send the id of the message in json format
type messageID struct {
	Id int `json:"id"`
}


//Store for the Messages
var msgStore = make(map[string]string)

//int used as key into message store
var idSeq int = 0

//==============================================================================
// Struct Error handling
type ErrorMessage struct {
	Status string `json:"status"`
	Message string `json:"error_message"`
}

func genericErrorHandler(w http.ResponseWriter, r *http.Request) {
	if p := recover(); p != nil {
		s := fmt.Sprintf("internal error>> %s", p)
        e := &ErrorMessage{Status:"error occured",Message: s}
		writeErrorResponse(w, e,http.StatusInternalServerError)
    }

}
func writeErrorResponse(w http.ResponseWriter,e *ErrorMessage, status int) {
    j, err := json.Marshal(e)
	if err != nil {
			panic(err)
	}
	fmt.Printf(e.Status,e.Message)
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}
//==============================================================================


//POST method handler
func PostMsgHandler(w http.ResponseWriter, r *http.Request) {
	defer genericErrorHandler(w,r)

	msgReader := http.MaxBytesReader(w,r.Body,4096) //Limit the max number of chars in message
	reqDump, err := ioutil.ReadAll(msgReader)
	
	if err != nil {
		panic (fmt.Sprintf("PostMsgHandler:[Max content length 4096] err[%s]", err))
	}
    
    msgContent := string(reqDump[:])
    idSeq++
    k := strconv.Itoa(idSeq)
    msgStore[k] =msgContent

    log.Printf("Received message %d => %s", idSeq, msgContent)
    resp := &messageID{idSeq}
    j, err := json.Marshal(resp)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(j)
}


//HTTP GET - /messages/{id}
func GetMsgHandler(w http.ResponseWriter, r *http.Request) {
	defer genericErrorHandler(w,r)
    vars := mux.Vars(r)
    k := vars["id"]
    log.Printf("key %s",k)
    if msgContent, ok := msgStore[k]; ok {
    	log.Printf("Key found [%s]=>[%s]",k,msgContent)
    	w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, msgContent) // Escape HTML special chars
    } else {
    	log.Printf("Key Not found %s",k)
        w.WriteHeader(http.StatusNotFound)
    }

}

//Entry point of the program
func main() {
	
    r := mux.NewRouter().StrictSlash(false)
    r.HandleFunc("/messages/{id}", GetMsgHandler).Methods("GET")
    r.HandleFunc("/messages", PostMsgHandler).Methods("POST")
    
    server := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }
    log.Println("Listening...")
    server.ListenAndServe()
}