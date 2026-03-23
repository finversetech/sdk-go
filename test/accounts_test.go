package test

import (
	"context"
	"testing"
)

func TestAccounts(t *testing.T) {
	client := NewTestClient()
	loginIdentityToken := "loginIdentityToken"
	loginIdentityCtx := contextWithAccessToken(context.Background(), loginIdentityToken)

	// Get Accounts
	accountsResp, _, err := client.LoginIdentityAPI.ListAccounts(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	if accountsResp.Accounts == nil {
		t.Errorf("empty accounts")
	}
}
