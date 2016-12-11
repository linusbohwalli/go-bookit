package bookit

import "time"

//This package should serve as the domain package and should not be dependent on any other packages in the app.
//Will be used for domain types

//TODO: HOW TO HANDLE AUTHENTICATION?? Authentication middleware?

//BookingID - We create own type for domain usage. This way we can call BookingID type in structs and services etc.
type BookingID string

//Booking - Defines how a booking looks like, used to seralize data to be passed to arangodb
type Booking struct {
	CreateTime       time.Time
	ModTime          time.Time
	ID               BookingID
	BookingDate      string
	RespContCustomer string
	RespContSeller   string
	ProjectCode      string
	//BookedBy User.Name
}

//User - Defines what information is necessary for a user
type User struct {
	ID           int
	Name         string
	PasswordHash string
	Email        string
}

//Client will be used to create a connection to services. We only got booking service for now.
type Client interface {
	BookingService() BookingService
}

//*** Services ***
// I.e a User service interface that fetch or save user data

//TODO: Create priorotize booking service (sorting service)

//BookingService - Handles everything with booking
type BookingService interface {
	//Booking is used to fetch a certain booking from db
	GetBooking(id BookingID) (*Booking, error)

	//CreateBooking is used when we need to create a new booking
	CreateBooking(booking *Booking) error

	//DeleteBooking takes an booking id and removes it from db.
	DeleteBooking(id BookingID) (*Booking, error)

	//TODO: complete update booking
	//UpdateBooking updates a booking
	UpdateBooking(booking *Booking) error
}
