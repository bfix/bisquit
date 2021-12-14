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

// GetBalances returns balance info for given currency
func (c *Client) GetBalances(ctx context.Context, curr string) (*BalancesInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &GetBalancesRequest{
		CurrencyCode: curr,
	}
	resp, err := c.wc.GetBalances(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Balances, nil
}

// GetAddressBalance returns the balance for a Bitcoin address
func (c *Client) GetAddressBalance(ctx context.Context, addr string) (*AddressBalanceInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &GetAddressBalanceRequest{
		Address: addr,
	}
	resp, err := c.wc.GetAddressBalance(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.AddressBalanceInfo, nil
}

// GetUnusedBsqAddress returns an unused BSQ address in the wallet
func (c *Client) GetUnusedBsqAddress(ctx context.Context) (string, error) {
	if c.conn == nil {
		return "", ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := c.wc.GetUnusedBsqAddress(ctx, &GetUnusedBsqAddressRequest{})
	if err != nil {
		return "", err
	}
	return resp.Address, nil
}

// SendBsq to transfer given amount of BSQ to address
func (c *Client) SendBsq(ctx context.Context, address, amount, txFeeRate string) (*TxInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &SendBsqRequest{
		Address:   address,
		Amount:    amount,
		TxFeeRate: txFeeRate,
	}
	resp, err := c.wc.SendBsq(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.TxInfo, nil
}

// SendBtc to send given amount of Bitcoin to address
func (c *Client) SendBtc(ctx context.Context, address, amount, txFeeRate, memo string) (*TxInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &SendBtcRequest{
		Address:   address,
		Amount:    amount,
		TxFeeRate: txFeeRate,
		Memo:      memo,
	}
	resp, err := c.wc.SendBtc(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.TxInfo, nil
}

// GetTxFeeRate returns information about the proposed fee rate
func (c *Client) GetTxFeeRate(ctx context.Context) (*TxFeeRateInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := c.wc.GetTxFeeRate(ctx, &GetTxFeeRateRequest{})
	if err != nil {
		return nil, err
	}
	return resp.TxFeeRateInfo, nil
}

// SetTxFeeRatePreference sets the preferred TxFeeRate
func (c *Client) SetTxFeeRatePreference(ctx context.Context, pref uint64) (*TxFeeRateInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &SetTxFeeRatePreferenceRequest{
		TxFeeRatePreference: pref,
	}
	resp, err := c.wc.SetTxFeeRatePreference(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.TxFeeRateInfo, nil
}

// UnsetTxFeeRatePreference unsets any previously specified preferene
func (c *Client) UnsetTxFeeRatePreference(ctx context.Context) (*TxFeeRateInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := c.wc.UnsetTxFeeRatePreference(ctx, &UnsetTxFeeRatePreferenceRequest{})
	if err != nil {
		return nil, err
	}
	return resp.TxFeeRateInfo, nil
}

// GetTransaction with the specified ID
func (c *Client) GetTransaction(ctx context.Context, txID string) (*TxInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &GetTransactionRequest{
		TxId: txID,
	}
	resp, err := c.wc.GetTransaction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.TxInfo, nil
}

// GetFundingAddresses returns a list of available funding addresses
func (c *Client) GetFundingAddresses(ctx context.Context) ([]*AddressBalanceInfo, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := c.wc.GetFundingAddresses(ctx, &GetFundingAddressesRequest{})
	if err != nil {
		return nil, err
	}
	return resp.AddressBalanceInfo, nil
}

// SetWalletPassword sets a new password for the wallet
func (c *Client) SetWalletPassword(ctx context.Context, passwdOld, passwdNew string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &SetWalletPasswordRequest{
		Password:    passwdOld,
		NewPassword: passwdNew,
	}
	_, err := c.wc.SetWalletPassword(ctx, req)
	return err
}

// RemoveWalletPassword removes password protection from the wallet
func (c *Client) RemoveWalletPassword(ctx context.Context, passwd string) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &RemoveWalletPasswordRequest{
		Password: passwd,
	}
	_, err := c.wc.RemoveWalletPassword(ctx, req)
	return err
}

// LockWallet locks a wallet from further usage
func (c *Client) LockWallet(ctx context.Context) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	_, err := c.wc.LockWallet(ctx, &LockWalletRequest{})
	return err
}

// UnlockWallet with password for given period of time
func (c *Client) UnlockWallet(ctx context.Context, passwd string, timeout uint64) error {
	if c.conn == nil {
		return ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &UnlockWalletRequest{
		Password: passwd,
		Timeout:  timeout,
	}
	_, err := c.wc.UnlockWallet(ctx, req)
	return err
}
