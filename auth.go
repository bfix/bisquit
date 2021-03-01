package bisquit

import (
	"context"
)

// PasswordCredential for authentication
// Must match "apiPassword=..." in "bisq.properties"
type PasswordCredential struct {
	passwd string
}

// GetRequestMetadata for authentication
func (c PasswordCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"password": c.passwd,
	}, nil
}

// RequireTransportSecurity not mandatory
func (c PasswordCredential) RequireTransportSecurity() bool {
	return false
}
