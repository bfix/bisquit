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
	"testing"
)

/*
rpc error: code = Unknown desc = balance is not yet available

func TestGetBalances(t *testing.T) {
	ctx := context.Background()
	balances, err := testClient.GetBalances(ctx, "BTC")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Balances: %v\n", balances)
}
*/

func TestGetUnusedBsqAddress(t *testing.T) {
	ctx := context.Background()
	addr, err := testClient.GetUnusedBsqAddress(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("BSQ address: %s\n", addr)
}

func TestGetTxFeeRate(t *testing.T) {
	ctx := context.Background()
	fee, err := testClient.GetTxFeeRate(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Fee rate: %v\n", fee)
}

func TestGetFundingAddresses(t *testing.T) {
	ctx := context.Background()
	addrs, err := testClient.GetFundingAddresses(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for i, addr := range addrs {
		t.Logf("Funding addr#%d: %v\n", i, addr)
	}
}
