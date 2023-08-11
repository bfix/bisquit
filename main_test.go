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
	"os"
	"testing"
	"time"
)

var (
	testClient *Client
)

func TestMain(m *testing.M) {
	// read test settings from environment
	host := os.Getenv("BISQ_API_HOST")
	if len(host) == 0 {
		fmt.Println("'BISQ_API_HOST' not defined -- skipping tests...")
		os.Exit(0)
	}
	passwd := os.Getenv("BISQ_API_PASSWORD")
	if len(passwd) == 0 {
		fmt.Println("'BISQ_API_PASSWORD' not defined -- skipping tests...")
		os.Exit(0)
	}
	// connect client to Bisq instance
	ctx := context.Background()
	testClient = NewClient(host, passwd, 30*time.Second)
	if err := testClient.Connect(ctx, time.Minute); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	// Run tests
	rc := m.Run()

	// teardown and cleanup
	testClient.Close()
	os.Exit(rc)
}

func TestClient(t *testing.T) {
	version, err := testClient.GetVersion(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Version is '%s'\n", version)
}
