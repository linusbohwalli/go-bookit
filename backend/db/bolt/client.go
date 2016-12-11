package bolt

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/linusbohwalli/go-bookit/backend/bookit"
)

// Client represents a client to the underlying BoltDB data store.
type Client struct {
	// Filename to the BoltDB database.
	Path string

	// Returns the current time.
	Now func() time.Time

	// Services
	bookingService BookingService

	db *bolt.DB
}

//NewClient initiates a new client and it to the bookingService
func NewClient() *Client {
	c := &Client{Now: time.Now}

	c.bookingService.client = c
	return c
}

//Open client connection to database
func (c *Client) Open() error {

	//Open connection to database
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	c.db = db

	//Init buckets
	tr, err := c.db.Begin(true)
	if err != nil {
		return err
	}
	defer tr.Rollback()

	//check if Bucket exist, if not, create it. TODO: add logic to send buck to Open() in bolt package?
	if _, err := tr.CreateBucketIfNotExists([]byte("Bookings")); err != nil {
		return err
	}
	return tr.Commit()
}

//Close closes client connection to db
func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

//BookingService returns the service associated with the client
//TODO: generalize so it is not bound specifically to booking service? or one for each service?
func (c *Client) BookingService() bookit.BookingService {
	return &c.bookingService
}
