package bolt

import (
	"fmt"
	"log"
	"reflect"

	"github.com/boltdb/bolt"
	"github.com/linusbohwalli/go-bookit/backend/bookit"
	"github.com/linusbohwalli/go-bookit/backend/db/bolt/internal"
)

//db handling for booking service, build db specific service? or general handling?

//Implement bookit.BookingService
var _ bookit.BookingService = &BookingService{}

//BookingService to set client
type BookingService struct {
	client *Client
}

//GetBooking - Receiver is a pointer to booking service struct which implements client
//GetBooking takes a bookit.BookingID to use in call to db, return the bookit.Booking
func (s *BookingService) GetBooking(id bookit.BookingID) (*bookit.Booking, error) {

	//start db session to get booking, read-only
	tx, err := s.client.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	//look for booking and unmarshal with protobuf (internal package)
	var b bookit.Booking

	//v is used to check for value in Bookings bucket
	if v := tx.Bucket([]byte("Bookings")).Get([]byte(id)); v == nil {
		// TODO: look over error handling
		//if v is empty, no booking with ID was found
		return nil, bookit.ErrBookingNotFound
	} else if err := internal.UnmarshalBooking(v, &b); err != nil {
		//else if we found booking in Bookings bucket, check for err during Unmarshal to protobuf.
		// TODO: look over error handling
		return nil, err
	}

	//return value to variable which points to bookit.Booking (DSL) struct
	return &b, nil
}

//CreateBooking - creates a new booking and saves to bookings bucket
func (s *BookingService) CreateBooking(b *bookit.Booking) error {

	//Most likely unecessary since we generate Id prior to this method call
	if b.ID == "" {
		return bookit.ErrBookingIDRequired
	}

	//start writing to database
	tx, err := s.client.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	//Check that booking do not exist
	//Should not happen since each booking got unique id
	//TODO: Evaluate if’’ below check is needed
	bu := tx.Bucket([]byte("Bookings"))
	if v := bu.Get([]byte(b.ID)); v != nil {
		return bookit.ErrBookingExists
	}

	//Set create and modified time
	t := s.client.Now()
	b.ModTime = t
	b.CreateTime = t

	//Marshal protobuf and add to bolt db
	if v, err := internal.MarshalBooking(b); err != nil {
		return err
	} else if err := bu.Put([]byte(b.ID), v); err != nil {
		return err
	}

	return tx.Commit()

}

//DeleteBooking deletes the booking with given id
func (s *BookingService) DeleteBooking(id bookit.BookingID) (*bookit.Booking, error) {

	var b bookit.Booking

	//Start get transaction
	if err := s.client.db.Update(func(tx *bolt.Tx) error {
		//Check id in bookings bucket
		v := tx.Bucket([]byte("Bookings")).Get([]byte(id))

		if v == nil {
			//check if booking exists else return
			return bookit.ErrBookingNotFound
		} else if err := internal.UnmarshalBooking(v, &b); err != nil {
			//check for unmarshalling error
			return bookit.ErrBookingUnmarshal
		}

		return nil
	}); err != nil {
		//TODO - fix error handling
		log.Fatal(err)
	}
	//Start delete transaction
	//Delete the booking
	if err := s.client.db.Update(func(tx *bolt.Tx) error {

		return tx.Bucket([]byte("Bookings")).Delete([]byte(id))
	}); err != nil {
		//fix error handling
		log.Fatal(err)
	}

	return &b, nil

}

//UpdateBooking takes a booking with updated values and replace the one in db. Returns an id to confirm what booking was updated and an error
func (s *BookingService) UpdateBooking(b *bookit.Booking) error {
	// TODO: Create UpdateBooking service function

	tx, err := s.client.db.Begin(true)
	if err != nil {
		//fix error handling
		log.Fatal(err)
	}
	defer tx.Rollback()

	var bk bookit.Booking
	//Get db bucket
	bu := tx.Bucket([]byte("Bookings"))

	//Fetch booking, key is ID
	v := bu.Get([]byte(b.ID))
	if v == nil {
		return bookit.ErrBookingNotFound
	} else if err := internal.UnmarshalBooking(v, &bk); err != nil {
		//check for unmarshalling error
		return bookit.ErrBookingUnmarshal
	}

	b1 := reflect.ValueOf(b)
	b2 := reflect.ValueOf(&bk)

	fmt.Println(b1)
	fmt.Println(b2)

	//Compare if any change was done from what is stored in database
	if b1 == b2 {
		return bookit.ErrBookingNoChange
	}

	//We check if bookings are different above, since bookings were different, we now update Modified time in the new booking
	t := s.client.Now()
	b.ModTime = t

	//After modified time is update, we put new booking back in bucket
	//Marshal protobuf and add to bolt db
	if v, err := internal.MarshalBooking(b); err != nil {
		return err
	} else if err := bu.Put([]byte(b.ID), v); err != nil {
		return err
	}

	return tx.Commit()

}
