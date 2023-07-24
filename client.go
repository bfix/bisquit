//----------------------------------------------------------------------
// This file is part of bisquit.
// Copyright (C) 2021 Bernd Fix >Y<
//
// bisquit is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// bisquit is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// SPDX-License-Identifier: AGPL3.0-or-later
//----------------------------------------------------------------------

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
	conn    *grpc.ClientConn   // active connection (close on exit)
	rpcHost string             // host:port spec for Bisq gRPC daemon
	creds   PasswordCredential // credential used in RPC call
	timeout time.Duration      // RPC timeout deadline in seconds

	// list of supported clients
	dac DisputeAgentsClient
	hc  HelpClient
	oc  OffersClient
	pac PaymentAccountsClient
	pc  PriceClient
	tc  TradesClient
	wc  WalletsClient
}

// NewClient instaniates a new Bisq client
func NewClient(host, passwd string) *Client {
	return &Client{
		conn:    nil,
		rpcHost: host,
		creds:   PasswordCredential(passwd),
		timeout: 5 * time.Second,
	}

}

// SetTimeout for RPC requests
func (c *Client) SetTimeout(t int) error {
	if t < 1 || t > 300 {
		return fmt.Errorf("invalid timeout value (%d)", t)
	}
	c.timeout = time.Duration(t) * time.Second
	return nil
}

// Connect to Bisq gRPC server
func (c *Client) Connect(ctx context.Context) (err error) {
	// check if client is already connected
	if c.conn != nil {
		return ErrClientConnected
	}
	// dial gRPC server with given credentials
	c.conn, err = grpc.DialContext(ctx, c.rpcHost,
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(c.creds),
		grpc.WithBlock())
	if err == nil {
		// instantiate all supported sub-clients
		c.dac = NewDisputeAgentsClient(c.conn)
		c.hc = NewHelpClient(c.conn)
		c.oc = NewOffersClient(c.conn)
		c.pac = NewPaymentAccountsClient(c.conn)
		c.pc = NewPriceClient(c.conn)
		c.tc = NewTradesClient(c.conn)
		c.wc = NewWalletsClient(c.conn)
	}
	return
}

// Close connection to Bisq gRPC server
func (c *Client) Close() error {
	// check if client is connected
	if c.conn == nil {
		return ErrClientNotConnected
	}
	err := c.conn.Close()
	// reset instances
	c.conn = nil
	c.dac = nil
	c.hc = nil
	c.oc = nil
	c.pac = nil
	c.pc = nil
	c.tc = nil
	c.wc = nil
	return err
}

// GetVersion returns the version of the Bisq server
func (c *Client) GetVersion(ctx context.Context) (string, error) {
	if c.conn == nil {
		return "", ErrClientNotConnected
	}
	srv := NewGetVersionClient(c.conn)
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	r, err := srv.GetVersion(ctx, &GetVersionRequest{})
	if err != nil {
		return "", err
	}
	return r.GetVersion(), nil
}
