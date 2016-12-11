package main

import (
	//standard lib packages
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/linusbohwalli/go-bookit/backend/handlers/booking"
	//non default packages - communicaty packages
	//go-bookit packages
)

//Main package shoud tie together dependencies

//*** Package Structs ***
//*** Package Variables ***
//*** Package Methods ***
//*** Package Functions ***

//init always runs when program is started
//func init() {
//}

//ONLY used for testing purposes
func handleNewBooking() {

	done := false
	done2 := false
	done3 := false

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Do you want to make a booking?: ")
	ans := scanner.Scan()
	switch {
	case ans && scanner.Text() == "yes":
		scanner1 := bufio.NewScanner(os.Stdin)
		fmt.Print("How many bookings do you want to make?: ")
		ans1 := scanner1.Scan()
		if ans1 {
			newAns1, err := strconv.Atoi(scanner1.Text())
			if err != nil {
				log.Fatal(err)
			}
			for i := 0; i < newAns1; i++ {
				bookingid, err := handlers.CreateBooking()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Your booking ID is:", bookingid, "save it to check your status")
			}
		}

		done = true
		//Send back to http that booking is completed, with booking info

	case ans && scanner.Text() != "yes":

		done = true
		fmt.Println("no booking done, please try again")
	}

	if done {
		time.Sleep(3 * time.Second)
		scanner2 := bufio.NewScanner(os.Stdin)
		fmt.Print("Do you want to see bookings in db?: ")
		ans2 := scanner2.Scan()
		switch {
		case ans2 && scanner2.Text() == "yes":

			scanner3 := bufio.NewScanner(os.Stdin)
			fmt.Print("what is the search id: ")
			ans3 := scanner3.Scan()
			if ans3 {
				searchID := scanner3.Text()

				res, err := handlers.GetBooking(searchID)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("here is the result: ", string(res))
				done2 = true

				//		fmt.Println(res)
				//To only print JSON data, just print it as a string.

			}
		case ans2 && scanner2.Text() != "yes":
			done2 = true
		}
	}

	if done2 {
		scanner = bufio.NewScanner(os.Stdin)
		fmt.Print("Do you want to delete a booking?: ")
		ans = scanner.Scan()

		if ans && scanner.Text() == "yes" {
			scanner = bufio.NewScanner(os.Stdin)
			fmt.Print("Please enter the id you want to delete: ")
			ans = scanner.Scan()

			if ans {
				deleteID := scanner.Text()

				id, booking, err := handlers.DeleteBooking(deleteID)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("the following booking with ID: %v have been deleted", id)
				fmt.Println("here is the result: ", string(booking))
				done3 = true

			}
		} else if ans && scanner.Text() != "yes" {
			done3 = true
		}
	}
	if done3 {
		scanner = bufio.NewScanner(os.Stdin)
		fmt.Print("Do you want to update a booking?: ")
		ans = scanner.Scan()

		if ans && scanner.Text() == "yes" {
			scanner = bufio.NewScanner(os.Stdin)
			fmt.Print("Please enter the json file you want to update: ")
			ans = scanner.Scan()

			if ans {
				jsonString := scanner.Text()

				data := []byte(jsonString)
				res, err := handlers.UpdateBooking(data)
				if err != nil {
					log.Fatal(err)

				}
				fmt.Println("Answer from res variable: ", res)

			}
		}
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.ExecuteTemplate(w, "index.html", nil)

}
func main() {
	//	fs := http.FileServer(http.Dir("static"))
	//	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//	http.HandleFunc("/", serveIndex)
	//	http.HandleFunc("/bookThis", bkiBooker.CreateBooking)

	//	if err := http.ListenAndServe(":8080", nil); err != nil {
	//		log.Fatal(err)
	//	}

	handleNewBooking()
}
