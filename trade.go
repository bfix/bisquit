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

	"google.golang.org/grpc"
)

// GetMarketPrice returns the price of Bitcoin in the given currency
func (c *Client) GetMarketPrice(ctx context.Context, curr string) (float64, error) {
	if c.conn == nil {
		return 0.0, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &MarketPriceRequest{
		CurrencyCode: curr,
	}
	resp, err := c.pc.GetMarketPrice(ctx, req)
	if err != nil {
		return 0.0, err
	}
	return resp.Price, nil
}

// GetTradeStatistics returns a list of past trades
func (c *Client) GetTradeStatistics(ctx context.Context) ([]*TradeStatistics3, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	resp, err := c.tsc.GetTradeStatistics(
		ctx, &GetTradeStatisticsRequest{},
		grpc.MaxCallRecvMsgSize(52428800))
	if err != nil {
		return nil, err
	}
	return resp.TradeStatistics, nil
}

// GetTrade returns the offer information for a trade with given ID
func (c *Client) GetTrade(ctx context.Context, ID string) (*TradeInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &GetTradeRequest{
		TradeId: ID,
	}
	resp, err := c.tc.GetTrade(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Trade, nil
}

// TakeOffer accepts an offer with given ID
func (c *Client) TakeOffer(ctx context.Context, offerID, accountID, takerFeeCurrency string) (*TradeInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &TakeOfferRequest{
		OfferId:              offerID,
		PaymentAccountId:     accountID,
		TakerFeeCurrencyCode: takerFeeCurrency,
	}
	resp, err := c.tc.TakeOffer(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Trade, nil
}

// ConfirmPaymentStarted starts the arbitration process for payments
func (c *Client) ConfirmPaymentStarted(ctx context.Context, tradeID string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &ConfirmPaymentStartedRequest{
		TradeId: tradeID,
	}
	_, err := c.tc.ConfirmPaymentStarted(ctx, req)
	return err
}

// ConfirmPaymentReceived closes an arbitration process for payments
func (c *Client) ConfirmPaymentReceived(ctx context.Context, tradeID string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &ConfirmPaymentReceivedRequest{
		TradeId: tradeID,
	}
	_, err := c.tc.ConfirmPaymentReceived(ctx, req)
	return err
}

// KeepFunds cancels a payment process
func (c *Client) KeepFunds(ctx context.Context, tradeID string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &KeepFundsRequest{
		TradeId: tradeID,
	}
	_, err := c.tc.KeepFunds(ctx, req)
	return err
}

// WithdrawFunds cancels a trade and withdraws Bitcoins to an address
func (c *Client) WithdrawFunds(ctx context.Context, tradeID, address, memo string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req := &WithdrawFundsRequest{
		TradeId: tradeID,
		Address: address,
		Memo:    memo,
	}
	_, err := c.tc.WithdrawFunds(ctx, req)
	return err
}
