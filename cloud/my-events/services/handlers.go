package services

import(
	"fmt"
	"net/http"
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
	fmt.Println("All events endpoint")
}

func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New events endpoint")
}
