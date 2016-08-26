package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

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
	RespContCustomer []string
	RespContSeller   []string
	ProjectCode      []string
	GoBookItID       string
}

type bookingslice struct {
	Bookings []booking
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

//TODO Func to check for error, if error found, log.Fatal
func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func writeJSON(s bookingslice, r *http.Request, id string) {

	if _, err := os.Stat("bookings.json"); err == nil && r.Method != "GET" {
		//TODO Build logic to add bookings into the JSON file. Use unique id to add a new booking,
		//need to rebuild structs to get a better JSON file
		data, err := ioutil.ReadFile("bookings.json")
		check(err)

		//TODO need to check this, something is happening atleast...
		var readBooking bookingslice
		e := json.Unmarshal(data, &readBooking)
		check(e)

		fmt.Println(s)
		fmt.Println(readBooking)

	} else {
		//Save booking data to JSON file
		b, err := json.MarshalIndent(s, "", "\t")
		check(err)

		/*
			//Not done yet, will indent JSON however, mismatch of types in ioutil.WriteFile and json.Indent.
			var out bytes.Buffer
			json.Indent(&out, b, "", "\t")

			d, err := out.WriteTo(os.Stdout)
			fmt.Println(d)
			check(err)
		*/

		e := ioutil.WriteFile("bookings.json", b, 0644)
		check(e)

	}
}

//Displaying start page of app
func loadIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html", "public/templates/layout.tmpl", "public/templates/drawer.tmpl", "public/templates/welcome.tmpl", "public/templates/calender.tmpl", "public/templates/qb.tmpl")
	check(err)

	t.ExecuteTemplate(w, "index", "")
}

func handleBooking(w http.ResponseWriter, r *http.Request) {
	//This part will first check if user logs on to url /booked instead of submitting a form
	if r.Method == "GET" {
		//redirect to frontpage
		http.Redirect(w, r, "/", 303)
	} else {

		//This part will parse the booking from the user
		err := r.ParseForm()
		check(err)

		//Using structs to structure JSON file
		var s bookingslice

		//Create unique ID to be used as GoBookItID (FT-ID)
		bookingid, err := shortid.Generate()
		check(err)

		s.Bookings = append(s.Bookings, booking{
			RespContCustomer: r.Form["respContCust"],
			RespContSeller:   r.Form["respContSeller"],
			ProjectCode:      r.Form["projectCode"],
			GoBookItID:       "GB-" + bookingid})

		//Send bookingslice and http.Request to function to create json
		writeJSON(s, r, bookingid)

		//Load frontPage again
		loadIndex(w, r)
	}
}

func main() {

	http.HandleFunc("/", loadIndex)
	http.HandleFunc("/booked", handleBooking)

	fs := justFilesFilesystem{http.Dir("public/")}
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(fs)))
	//Listen and Serve the files
	http.ListenAndServe(":8080", nil)
	//beego.Run()
}
