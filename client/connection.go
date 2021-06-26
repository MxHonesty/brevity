package client

import (
	"brevity/command"
	"brevity/response"
	"encoding/gob"
	"errors"
	"net"
	"strconv"
)

// Function for registering all the concrete types of command.Command
// with gob.
func registerGobTypes() {
	//gob.Register()
}

// TODO: Main client loop, command-request dynamic.

// Struct responsible for forwarding commands to the server.
type Client struct {
	host       string
	port       uint64
	connection net.Conn
	connected  bool  // True if there is an ongoing connection.
}

// Creates a new Client instance. It is initially not connected.
func NewClient(host string, port uint64) *Client {
	registerGobTypes()
	return &Client{host: host, port: port, connected: false}
}

// Method establishes the Connection. Returns a net.Conn object and an error. If
// error is not nil, the connection failed.
//
// 	Errors:
//		Already connected
//		Connection error
func (c *Client) Connect() error {
	if !c.connected {
		port := strconv.FormatUint(c.port, 10)  // Convert to string.
		var err error
		c.connection, err = net.Dial("tcp", c.host + ":" + port)
		if err != nil {
			return errors.New("failed to establish connection")
		}
	}
	return errors.New("a connection is already ongoing")
}

// Sends the command.Command instance to the server using gob encoding. Returns a
// non-nil error if the send was not completed successfully
func (c *Client) SendCommand(com command.Command) (response.Response, error) {
	if c.connected {
		errSend := gob.NewEncoder(c.connection).Encode(com)
		if errSend != nil {
			return response.Response{}, errors.New("could not send command")
		} else {
			var data response.Response
			errReceive := gob.NewDecoder(c.connection).Decode(&data)
			if errReceive != nil {
				return response.Response{}, errors.New("decoding error")
			} else {
				return data, nil
			}
		}
	}
	return response.Response{}, errors.New("no connection started")
}

// Method for closing the connection.
func (c *Client) Close() error {
	c.connected = false
	return c.connection.Close()
}
