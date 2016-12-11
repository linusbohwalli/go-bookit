package handlers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/linusbohwalli/go-bookit/backend/bookit"
	"github.com/linusbohwalli/go-bookit/backend/db/bolt"
	uuid "github.com/satori/go.uuid"
)

const (
	bookitdb = "bookit.db"
)

//CreateBooking - creates a booking and saves in db
//func CreateBooking(w http.ResponseWriter, r *http.Request) {
func CreateBooking() (string, error) {
	//if r.Method == "POST" {

	//Create new UUID for each booking
	newUUID := uuid.NewV4().String()

	//Create unique ID to be used as GoBookItID (GB-ID)
	/*id, err := shortid.Generate()
	if err != nil {
		log.Fatal(err)
	}
	bookingid := "GB" + id*/

	//Actual booking create time
	currDate := time.Now().String()

	//Fetch Responsible contact for customer from http.request
	//respContCustomer := r.FormValue("RespContCustomer")

	//Fetch Responsible seller contact from http.request
	//respContSeller := r.FormValue("RespContSeller")

	//Fetch project code from http.request
	projectCode := "testValue3"

	bookingData := &bookit.Booking{

		ID:               bookit.BookingID(newUUID),
		BookingDate:      currDate,
		RespContCustomer: "testValue1",
		RespContSeller:   "testValue2",
		ProjectCode:      projectCode,
	}

	c := bolt.NewClient()
	c.Path = bookitdb
	if err := c.Open(); err != nil {
		log.Fatal(err)
	}

	if err := c.BookingService().CreateBooking(bookingData); err != nil {
		log.Fatal(err)
	}

	if err := c.Close(); err != nil {
		// TODO: fix error handling
		log.Fatal(err)
	}
	return newUUID, nil
	//	}// else {
	//TODO: serve index file if other method than post
	//	fmt.Fprintf(w, "Not the method your looking for")
	//	}

}

//GetBooking - Handle to fetch booking in db.
func GetBooking(id string) ([]byte, error) {

	c := bolt.NewClient()
	c.Path = bookitdb
	if err := c.Open(); err != nil {
		return nil, err
	}
	data, err := c.BookingService().GetBooking(bookit.BookingID(id))
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := c.Close(); err != nil {
		// TODO: fix error handling
		log.Fatal(err)
	}

	return b, nil

}

//DeleteBooking handles the call to delete in booking service
//Returns the given id, the booking that was the value of that id and an error interface
func DeleteBooking(id string) (string, []byte, error) {

	c := bolt.NewClient()
	c.Path = bookitdb
	if err := c.Open(); err != nil {
		// TODO: error handling
		return id, nil, err
	}

	data, err := c.BookingService().DeleteBooking(bookit.BookingID(id))
	if err != nil {
		return id, nil, err
	}

	//marshal JSON to present to user
	b, err := json.Marshal(data)
	if err != nil {
		return id, nil, err
	}

	if err := c.Close(); err != nil {
		// TODO: fix error handling
		log.Fatal(err)
	}

	return id, b, nil
}

//UpdateBooking updates a booking by calling bookingservice.UpdateBooking. takes a JSON file, return result and an error.
func UpdateBooking(data []byte) (string, error) {

	//Open new client
	c := bolt.NewClient()
	//Set what db to init
	c.Path = bookitdb
	//Open the above database
	if err := c.Open(); err != nil {
		//TODO: error handling
		return "", err
	}

	var b bookit.Booking
	//UnmarshalJSON
	if err := json.Unmarshal(data, &b); err != nil {
		// TODO: fix error handling
		log.Fatal(err)
	}

	if err := c.BookingService().UpdateBooking(&b); err != nil {
		log.Fatal(err)
	}

	if err := c.Close(); err != nil {
		// TODO: fix error handling
		log.Fatal(err)
	}

	return "Booking successfully updated", nil
}
