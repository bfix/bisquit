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
	"encoding/json"
)

// CreatePaymentAccount creates a new payment account
func (c *Client) CreatePaymentAccount(ctx context.Context, form string) (*PaymentAccount, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &CreatePaymentAccountRequest{
		PaymentAccountForm: form,
	}
	resp, err := c.pac.CreatePaymentAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.PaymentAccount, nil
}

// GetPaymentAccounts returns a list of payment accounts
func (c *Client) GetPaymentAccounts(ctx context.Context) ([]*PaymentAccount, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := c.pac.GetPaymentAccounts(ctx, &GetPaymentAccountsRequest{})
	if err != nil {
		return nil, err
	}
	return resp.PaymentAccounts, nil
}

// GetPaymentMethods returns all available payment methods
func (c *Client) GetPaymentMethods(ctx context.Context) ([]*PaymentMethod, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	resp, err := c.pac.GetPaymentMethods(ctx, &GetPaymentMethodsRequest{})
	if err != nil {
		return nil, err
	}
	return resp.PaymentMethods, nil
}

// GetPaymentAccountForm returns a template for payment accounts
func (c *Client) GetPaymentAccountForm(ctx context.Context, mthdID string) (map[string]interface{}, error) {
	if c.conn == nil {
		return nil, ErrClientNotConnected
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req := &GetPaymentAccountFormRequest{
		PaymentMethodId: mthdID,
	}
	resp, err := c.pac.GetPaymentAccountForm(ctx, req)
	if err != nil {
		return nil, err
	}
	form := resp.PaymentAccountFormJson
	res := make(map[string]interface{})
	err = json.Unmarshal([]byte(form), &res)
	return res, err
}
