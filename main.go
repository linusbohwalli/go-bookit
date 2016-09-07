package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	//non default packages

	"github.com/ventu-io/go-shortid"
)

//import "github.com/astaxie/beego"

//Global structs
//No Global structs yet

//Global functions
//No Global functions yet

//Package structs

//Below struct is part of blocking directory listing
type justFilesFilesystem struct {
	fs http.FileSystem
}

//Below struct is part of blocking directory listing
type neuteredReaddirFile struct {
	http.File
}

//TODO Split booking struct? Create one for each value in parseForm?
type booking struct {
	RespContCustomer string
	RespContSeller   string
	ProjectCode      string
	Bookingid        string
	BookingDate      time.Time
}

type bookingheader struct {
	UUID     string
	Bookings booking
}

type bookingheaders struct {
	bookingheaders []bookingheader
}

//Package functions

//Below function is part of blocking directory listing
func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

//Below function is part of blocking directory listing
func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func writeJSON(dat bookingheader, w http.ResponseWriter, r *http.Request) {

	if _, err := os.Stat("bookings.json"); err == nil && r.Method != "GET" {
		// TODO: fix logic for appending more information to JSON, open and resave?
		//Read JSON
		data, err := ioutil.ReadFile("bookings.json")
		if err != nil {
			log.Fatal(err)
		}

		var booking []bookingheader
		e := json.Unmarshal(data, &booking)
		if e != nil {
			log.Fatal(e)
		}

		booking = append(booking, dat)

		b, err := json.MarshalIndent(booking, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile("bookings.json", b, 0644); err != nil {
			log.Fatal(err)
		}

	} else {

		//Save booking data to JSON file, one file for each booking, this should be transferred to postgres db, but will suffice for now
		b, err := json.MarshalIndent(dat, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile("bookings.json", b, 0644); err != nil {
			log.Fatal(err)
		}
	}
}

//Displaying start page of app
func loadIndex(w http.ResponseWriter, r *http.Request) {

	switch string(r.URL.Path[1:]) {
	case "new:user:booking:connection":
		t, err := template.ParseFiles("index.html", "public/templates/layout.tmpl", "public/templates/drawer.tmpl", "public/templates/content.tmpl", "public/templates/calender.tmpl", "public/templates/connection.tmpl")
		if err != nil {
			log.Fatal(err)
		}
		t.ExecuteTemplate(w, "index", "")

	case "show:user:booking":
		t, err := template.ParseFiles("index.html", "public/templates/layout.tmpl", "public/templates/drawer.tmpl", "public/templates/content.tmpl", "public/templates/calender.tmpl", "public/templates/showbookings.tmpl")
		if err != nil {
			log.Fatal(err)
		}

		b, err := displayBooking()
		if err != nil {
			log.Fatal(err)
		}

		t.ExecuteTemplate(w, "index", b)

	default:
		t, err := template.ParseFiles("index.html", "public/templates/layout.tmpl", "public/templates/drawer.tmpl", "public/templates/content.tmpl", "public/templates/calender.tmpl", "public/templates/leadtime.tmpl")
		if err != nil {
			log.Fatal(err)
		}
		t.ExecuteTemplate(w, "index", "")

	}
}

func handleBooking(w http.ResponseWriter, r *http.Request) {
	//This part will first check if user logs on to url /booked instead of submitting a form
	if r.Method == "GET" {
		//redirect to frontpage
		http.Redirect(w, r, "/", 303)
	} else {

		//This part will parse the booking from the user
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		//Create unique ID to be used as GoBookItID (GB-ID)
		bookingid, err := shortid.Generate()
		bookingid = "GB" + bookingid
		if err != nil {
			log.Fatal(err)
		}

		currDate := time.Now()

		dat := bookingheader{bookingid,
			booking{r.FormValue("respContCust"),
				r.FormValue("respContSeller"),
				r.FormValue("projectCode"),
				bookingid,
				currDate}}

		//Send bookingslice and http.Request to function to create json
		writeJSON(dat, w, r)

		//Load frontPage again
		http.Redirect(w, r, "/", 303)
	}
}

func displayBooking() (booking, error) {

	return booking{Bookingid: "test2"}, nil

}

func main() {

	http.HandleFunc("/", loadIndex)
	http.HandleFunc("/booked", handleBooking)

	fs := justFilesFilesystem{http.Dir("public/")}
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(fs)))
	//Listen and Serve the files
	log.Fatal(http.ListenAndServe(":8080", nil))

}
