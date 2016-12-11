package internal

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/linusbohwalli/go-bookit/backend/bookit"
)

//The go generate comment will run the command after go generate keywords. In this case we will compile the pb.go file for protobuffer usage.

//go:generate protoc --gogo_out=. internal.proto

//MarshalBooking TODO: write stuff
func MarshalBooking(b *bookit.Booking) ([]byte, error) {

	return proto.Marshal(&Booking{
		//UnixNano is used to convert time.Time  to an int64 for type safe storage
		CreateTime: b.CreateTime.UnixNano(),
		ModTime:    b.ModTime.UnixNano(),
		//Convert to string since this is a BookingID which is a domain type
		ID:               string(b.ID),
		BookingDate:      b.BookingDate,
		RespContCustomer: b.RespContCustomer,
		RespContSeller:   b.RespContSeller,
		ProjectCode:      b.ProjectCode,
	})
}

//UnmarshalBooking TODO: write stuff
func UnmarshalBooking(data []byte, b *bookit.Booking) error {
	var pb Booking
	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}

	//Convert back to UTC time
	b.CreateTime = time.Unix(0, pb.CreateTime).UTC()
	b.ModTime = time.Unix(0, pb.ModTime).UTC()
	//Convert back to domain type
	b.ID = bookit.BookingID(pb.ID)
	b.BookingDate = pb.BookingDate
	b.RespContCustomer = pb.RespContCustomer
	b.RespContSeller = pb.RespContSeller
	b.ProjectCode = pb.ProjectCode

	return nil

}
