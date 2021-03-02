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
	"time"
)

// GetOffer returns the offer for a given ID
func (c *Client) GetOffer(ctx context.Context, ID string) (*OfferInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &GetOfferRequest{
		Id: ID,
	}
	resp, err := c.oc.GetOffer(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Offer, nil
}

// GetMyOffer returns our offer for a given ID
func (c *Client) GetMyOffer(ctx context.Context, ID string) (*OfferInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &GetMyOfferRequest{
		Id: ID,
	}
	resp, err := c.oc.GetMyOffer(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Offer, nil
}

// GetOffers returns all offers for given criteria
func (c *Client) GetOffers(ctx context.Context, dir, curr string) ([]*OfferInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &GetOffersRequest{
		Direction:    dir,
		CurrencyCode: curr,
	}
	resp, err := c.oc.GetOffers(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Offers, nil
}

// GetMyOffers returns all of our offers for given criteria
func (c *Client) GetMyOffers(ctx context.Context, dir, curr string) ([]*OfferInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &GetMyOffersRequest{
		Direction:    dir,
		CurrencyCode: curr,
	}
	resp, err := c.oc.GetMyOffers(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Offers, nil
}

// CreateOffer to create a new offering
func (c *Client) CreateOffer(ctx context.Context, req *CreateOfferRequest) (*OfferInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	resp, err := c.oc.CreateOffer(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Offer, nil
}

// CancelOffer to terminate an active offering
func (c *Client) CancelOffer(ctx context.Context, ID string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &CancelOfferRequest{
		Id: ID,
	}
	_, err := c.oc.CancelOffer(ctx, req)
	return err
}
