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

func TestGetPaymentAccounts(t *testing.T) {
	ctx := context.Background()
	accnts, err := testClient.GetPaymentAccounts(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for i, accnt := range accnts {
		t.Logf("Account#%d: %v\n", i, accnt)
	}
}

func TestGetPaymentMethods(t *testing.T) {
	ctx := context.Background()
	mthds, err := testClient.GetPaymentMethods(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for i, mthd := range mthds {
		t.Logf("Method#%d: %v\n", i, mthd)
	}
}

func TestGetPaymentAccountForm(t *testing.T) {
	ctx := context.Background()
	form, err := testClient.GetPaymentAccountForm(ctx, "SEPA")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Form: %v\n", form)
}
