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

// PasswordCredential for Bisq API authentication:
// Must match "apiPassword=..." in "bisq.properties"
type PasswordCredential string

// GetRequestMetadata for API password authentication
func (c PasswordCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"password": string(c),
	}, nil
}

// RequireTransportSecurity signals that insecure (local) connections are fine
func (c PasswordCredential) RequireTransportSecurity() bool {
	return false
}
