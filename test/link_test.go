package test

import (
	"context"
	"testing"

	"github.com/finversetech/sdk-go/finverse"
)

func TestLinkToken(t *testing.T) {
	userId := "someUserId"     // reference back to your system userId, finverse does not use this
	state := "someUniqueState" // this will be sent in the redirectUri callback, can be used to identify the state
	customerAccessToken := "customerAccessToken"

	client := NewTestClient()
	customerCtx := context.WithValue(context.Background(), finverse.ContextAccessToken, customerAccessToken)
	linkTokenResp, _, err := client.CustomerApi.GenerateLinkToken(customerCtx).LinkTokenRequest(finverse.LinkTokenRequest{
		ClientId:     clientId,
		UserId:       &userId,
		RedirectUri:  redirectUri,
		State:        &state,
		ResponseMode: "form_post",
		ResponseType: "code",
		GrantType:    "client_credentials",
	}).Execute()
	if err != nil {
		panic(err)
	}

	// the LinkUrl can be used to initialize Finverse Link
	// when Finverse Link is finished, a code will be sent to redirectUri and it can be used to exchange for LoginIdentity Token
	if linkTokenResp.LinkUrl == "" {
		t.Errorf("empty linkUrl")
	}
}

func TestLoginIdentityToken(t *testing.T) {
	customerAccessToken := "customerAccessToken"
	client := NewTestClient()
	customerCtx := context.WithValue(context.Background(), finverse.ContextAccessToken, customerAccessToken)

	// When Finverse link is done, obtain the code and use it to exchange for login identity access token
	code := "obtainAfterLink"
	loginIdentityTokenResp, _, err := client.LinkApi.Token(customerCtx).
		Code(code).
		ClientId(clientId).
		RedirectUri(redirectUri).
		GrantType("authorization_code").
		Execute()
	if err != nil {
		panic(err)
	}

	// the loginIdentityToken can be used to retrieve data
	loginIdentityToken := loginIdentityTokenResp.AccessToken
	if loginIdentityToken == "" {
		t.Errorf("empty loginIdentityToken")
	}
}
