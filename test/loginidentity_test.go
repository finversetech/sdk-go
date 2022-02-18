package test

import (
	"context"
	"testing"
	"time"

	"github.com/finversetech/sdk-go/finverse"
)

func TestLoginIdentity(t *testing.T) {
	client := NewTestClient()
	loginIdentityToken := "loginIdentityToken"
	loginIdentityCtx := context.WithValue(context.Background(), finverse.ContextAccessToken, loginIdentityToken)

	// Get LoginIdentity
	loginIdentityResp, _, err := client.LoginIdentityApi.GetLoginIdentity(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	if loginIdentityResp.LoginIdentity == nil {
		t.Errorf("empty loginIdentity")
	}
}

func TestLoginIdentityPollUntilReady(t *testing.T) {
	client := NewTestClient()
	loginIdentityToken := "loginIdentityToken"
	loginIdentityCtx := context.WithValue(context.Background(), finverse.ContextAccessToken, loginIdentityToken)

	// Wait until loginIdentityStatus is ready
	finalStatuses := map[string]struct{}{"ERROR": struct{}{}, "DATA_RETRIEVAL_PARTIALLY_SUCCESSFUL": struct{}{}, "DATA_RETRIEVAL_COMPLETE": struct{}{}}
	var i int
	var loginIdentity *finverse.LoginIdentity
	for i = 0; i < 20; i++ {
		loginIdentityResp, _, err := client.LoginIdentityApi.GetLoginIdentity(loginIdentityCtx).Execute()
		if err != nil {
			panic(err)
		}
		loginIdentity = loginIdentityResp.LoginIdentity

		if _, ok := finalStatuses[loginIdentity.GetStatus()]; ok {
			break
		}

		time.Sleep(3 * time.Second)
	}

	if loginIdentity.GetStatus() == "ERROR" {
		t.Errorf("loginIdentity status error: %s", loginIdentity.GetStatus())
	}

}
