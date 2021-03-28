package services

import(
	"fmt"
	"net/http"
	"encoding/json"
	"brunoedsm.com/v1/models"
)

type eventServiceHandler struct {
	dbhandler string
}

func NewEventHandler(databasehandler string) *eventServiceHandler {
	fmt.Println("Open db connection: ",databasehandler)
	return &eventServiceHandler{
		dbhandler: databasehandler,
	}
}

func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Search events endpoint")
}

func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	
	obj := models.Hall{"Royal Albert Hall","UK",1000}
	
	fmt.Println(obj)
}

func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	obj := models.Hall{}
	err := json.NewDecoder(r.Body).Decode(&obj)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "error occured while decoding event data %s"}`, err)
		return
	}
	fmt.Println(obj)
}
