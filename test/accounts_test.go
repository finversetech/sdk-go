package test

import (
	"context"
	"testing"

	"github.com/finversetech/sdk-go/finverse"
)

func TestAccounts(t *testing.T) {
	client := NewTestClient()
	loginIdentityToken := "loginIdentityToken"
	loginIdentityCtx := context.WithValue(context.Background(), finverse.ContextAccessToken, loginIdentityToken)

	// Get Accounts
	accountsResp, _, err := client.LoginIdentityApi.ListAccounts(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	if accountsResp.Accounts == nil {
		t.Errorf("empty accounts")
	}
}
