package bookit

//General application errors

//Booking errors
const (
	ErrBookingNotFound   = Error("Booking not found")
	ErrBookingExists     = Error("Booking already exists, cannot create same booking twice")
	ErrBookingIDRequired = Error("No ID was received in booking_service, something must have gone during ID creation")
	ErrBookingUnmarshal  = Error("An unexpected error occured when booking was read in database, please try again or contact system administrator")
	ErrBookingNoChange   = Error("No change was done")
)

//Error -  error type is used for bookit errors.
type Error string

//Error returns the error message
func (e Error) Error() string { return string(e) }
