package bisquit

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

// Error codes
var (
	ErrClientConnected    = fmt.Errorf("Client already connected")
	ErrClientNotConnected = fmt.Errorf("Client not connected")
)

// Client for Bisq API calls
type Client struct {
	conn    *grpc.ClientConn
	rpcHost string
	creds   PasswordCredential
}

// NewClient instaniates a new Bisq client
func NewClient(host, passwd string) *Client {
	return &Client{
		conn:    nil,
		rpcHost: host,
		creds:   PasswordCredential{passwd},
	}
}

// Connect to Bisq gRPC server
func (c *Client) Connect() (err error) {
	if c.conn != nil {
		return ErrClientConnected
	}
	c.conn, err = grpc.Dial(c.rpcHost, grpc.WithInsecure(), grpc.WithPerRPCCredentials(c.creds))
	return
}

// Close connection to Bisq gRPC server
func (c *Client) Close() error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	err := c.conn.Close()
	c.conn = nil
	return err
}

// GetVersion returns the version of the Bisq server
func (c *Client) GetVersion() (string, error) {
	if c.conn == nil {
		return "", ErrClientNotConnected
	}
	srv := NewGetVersionClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := srv.GetVersion(ctx, &GetVersionRequest{})
	if err != nil {
		return "", err
	}
	return r.GetVersion(), nil
}
