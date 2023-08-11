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
)

// GetMarketPrice returns the price of Bitcoin in the given currency
func (c *Client) GetMarketPrice(ctx context.Context, curr string) (float64, error) {
	if c.conn == nil {
		return 0.0, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
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

// GetTrade returns the offer information for a trade with given ID
func (c *Client) GetTrade(ctx context.Context, ID string) (*TradeInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
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

// GetTrades returns offers:
// mode = 0 (Open), 1 (Closed), 2 (Failed)
func (c *Client) GetTrades(ctx context.Context, mode int) ([]*TradeInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &GetTradesRequest{
		Category: GetTradesRequest_Category(mode),
	}
	resp, err := c.tc.GetTrades(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Trades, nil
}

// TakeOffer accepts an offer with given ID
func (c *Client) TakeOffer(ctx context.Context, amount int64, offerID, accountID, takerFeeCurrency string) (*TradeInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &TakeOfferRequest{
		OfferId:              offerID,
		PaymentAccountId:     accountID,
		TakerFeeCurrencyCode: takerFeeCurrency,
		Amount:               uint64(amount),
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
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
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
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &ConfirmPaymentReceivedRequest{
		TradeId: tradeID,
	}
	_, err := c.tc.ConfirmPaymentReceived(ctx, req)
	return err
}

// FailTrade cancels a trade
func (c *Client) FailTrade(ctx context.Context, tradeID string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &FailTradeRequest{
		TradeId: tradeID,
	}
	_, err := c.tc.FailTrade(ctx, req)
	return err
}

// UnFailTrade revives a failed trade
func (c *Client) UnFailTrade(ctx context.Context, tradeID string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &UnFailTradeRequest{
		TradeId: tradeID,
	}
	_, err := c.tc.UnFailTrade(ctx, req)
	return err
}

// CloseTrade closes a trade
func (c *Client) CloseTrade(ctx context.Context, tradeID string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &CloseTradeRequest{
		TradeId: tradeID,
	}
	_, err := c.tc.CloseTrade(ctx, req)
	return err
}

// WithdrawFunds cancels a trade and withdraws Bitcoins to an address
func (c *Client) WithdrawFunds(ctx context.Context, tradeID, address, memo string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &WithdrawFundsRequest{
		TradeId: tradeID,
		Address: address,
		Memo:    memo,
	}
	_, err := c.tc.WithdrawFunds(ctx, req)
	return err
}
