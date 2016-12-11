package arango

import (
	//standard lib packages
	//non-default packages
	"fmt"
	"log"

	"github.com/solher/arangolite"
	//go-bookit packages
)

//***PACKAGE VARIABLES***
//***PACKAGE STRUCTS***

//ReadfromDB - ...
func ReadfromDB(searchVal string) ([]byte, error) {

	db := arangolite.New().
		LoggerOptions(true, false, false).
		Connect("http://localhost:8529", "bookit", "root", "")

	q := arangolite.NewQuery(`FOR b in bookings FILTER b._key == "%v" RETURN b`, searchVal).BatchSize(100)

	r, err := db.Run(q)
	if err != nil {
		log.Fatal(err)
	}
	return r, nil
}

//WriteToDB - Accepts a []byte which must be marshalled as JSON data
//TODO: Maybe better to do JSON handling in db writer
func WriteToDB(ch chan []byte, done chan bool) {

	for {
		work, moreWork := <-ch
		if moreWork {
			//Convert to string To get JSON data from []byte
			//To Write to DB the []byte in this function must be JSON
			data := string(work)
			//Connect to DB
			db := arangolite.New().
				LoggerOptions(true, false, false).
				Connect("http://localhost:8529", "bookit", "root", "")

			//Create new write transaction
			t := arangolite.NewTransaction(nil, []string{"bookings"}).
				AddQuery("resultVar", `INSERT %v INTO bookings`, data)

			if _, err := db.Run(t); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("all done with db")
			done <- true
			return
		}

	}

}
