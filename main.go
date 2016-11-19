package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	//non default packages
	pb "github.com/linusbohwalli/ng-go-bookit/backend/protobuffers/booking"

	"github.com/golang/protobuf/proto"
	"github.com/satori/go.uuid"
	"github.com/ventu-io/go-shortid"
)

//*** Global structs ***

//Defines how a booking looks like, used in protobuf message

//*** Global Variables ***
//*** Global Methods ***

//*** Global Functions ***

//init always runs when program is started
//func init() {
//}

func createBooking() (b []byte, e error) {

	//Create new UUID for each booking
	newUUID := uuid.NewV4().String()

	//Create unique ID to be used as GoBookItID (GB-ID)
	id, err := shortid.Generate()
	if err != nil {
		log.Fatal(err)
	}
	bookingid := "GB" + id

	//Actual booking create time
	currDate := time.Now().String()

	//Fetch Responsible contact for customer from http.request
	respContCustomer := "testValue1"

	//Fetch Responsible seller contact from http.request
	respContSeller := "testValue2"

	//Fetch project code from http.request
	projectCode := "testValue3"

	bookingData := &pb.Booking{
		UUID:             newUUID,
		Bookingid:        bookingid,
		BookingDate:      currDate,
		RespContCustomer: respContCustomer,
		RespContSeller:   respContSeller,
		ProjectCode:      projectCode,
	}

	out, err := proto.Marshal(bookingData)
	if err != nil {
		log.Fatal(err)
	}
	return out, nil

	/*	return &booking{newUUID,
		bookingid,
		currDate,
		respContCustomer,
		respContSeller,
		projectCode}, nil
	*/
}

func main() {

	http.Handle("/src/", http.StripPrefix("/src", http.FileServer(http.Dir("./src"))))
	http.Handle("/node_modules/", http.StripPrefix("/node_modules", http.FileServer(http.Dir("./node_modules"))))
	//TODO: Build http server / http serve static files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/index.html")
	})

	//Start web server
	http.ListenAndServe(":8080", nil)

	//implement API handling for booking / handleFunc
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Do you want to make a booking?: ")
	ans := scanner.Scan()
	if ans && scanner.Text() == "yes" {

		b, err := createBooking()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Slice of bytes: ", b)

		//Send back to http that booking is completed, with booking info
		fmt.Println("booking complete")

		fmt.Print("Do you want to look at the booking result?: ")
		ans2 := scanner.Scan()
		if ans2 && scanner.Text() == "yes" {
			//unmarshal proto
			res := &pb.Booking{}
			err = proto.Unmarshal(b, res)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Booking result: ", res, "     printlninign works aswell")
		}
	}
}
