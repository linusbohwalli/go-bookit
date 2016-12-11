package bolt_test

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/linusbohwalli/go-bookit/backend/db/bolt"
)

//Now is the mocked current time for testing
var Now = time.Date(2016, time.December, 5, 1, 0, 0, 0, time.UTC)

//Client wraps a bolt.Client
type Client struct {
	*bolt.Client
}

//NewClient return a Client pointing to the temp file
func NewClient() *Client {

	//generate temp file name
	f, err := ioutil.TempFile("", "bolt-client-test")
	if err != nil {
		panic(err)
	}
	f.Close()

	//Create client wrapper
	c := &Client{
		Client: bolt.NewClient(),
	}
	c.Path = f.Name()
	c.Now = func() time.Time { return Now }

	return c
}

//MustOpenClient returns new open client instance
func MustOpenClient() *Client {

	c := NewClient()
	if err := c.Open(); err != nil {
		panic(err)
	}

	return c
}

func (c *Client) Close() error {
	defer os.Remove(c.Path)
	return c.Client.Close()
}
