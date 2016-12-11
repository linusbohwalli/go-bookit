package http

import (
	"github.com/linusbohwalli/go-bookit/backend/bookit"

	"net/http"
)

type booking struct {
	bookit.Booking
}

//in start of func we can implement a service(receiver) which we should apply serveHTTP on, THIS IS NOT DONE JSUT TO EXEMPLIFY HOW WE WRAP STANDARD LIB
func (b *booking) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//handle request here
}
