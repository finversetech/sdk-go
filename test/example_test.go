package test

import (
	"context"
	"testing"
	"time"

	"github.com/finversetech/sdk-go/finverse"
)

func TestAll(t *testing.T) {
	ctx := context.Background()
	client := NewTestClient()

	// obtain customer access token
	customerTokenResp, _, err := client.PublicAPI.GenerateCustomerAccessToken(ctx).TokenRequest(finverse.TokenRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}).Execute()
	if err != nil {
		panic(err)
	}
	customerAccessToken := customerTokenResp.AccessToken

	// generate a link token
	userId := "someUserId"     // reference back to your system userId, finverse does not use this
	state := "someUniqueState" // this will be sent in the redirectUri callback, can be used to identify the state
	customerCtx := contextWithAccessToken(context.Background(), customerAccessToken)
	linkTokenResp, _, err := client.LinkAPI.GenerateLinkToken(customerCtx).LinkTokenRequest(finverse.LinkTokenRequest{
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

	// the linkUrl can be used to initiate Finverse Link
	t.Logf("linkUrl: %s", linkTokenResp.LinkUrl)

	// When Finverse link is done, obtain the code and use it to exchange for login identity access token
	code := "obtainAfterLink"
	loginIdentityTokenResp, _, err := client.LinkAPI.Token(customerCtx).
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
	loginIdentityCtx := contextWithAccessToken(context.Background(), loginIdentityToken)

	// Get LoginIdentity
	loginIdentityResp, _, err := client.LoginIdentityAPI.GetLoginIdentity(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	t.Logf("login identity: %+v", loginIdentityResp.LoginIdentity)

	// Wait until loginIdentityStatus is ready
	finalStatuses := map[string]struct{}{"ERROR": struct{}{}, "DATA_RETRIEVAL_PARTIALLY_SUCCESSFUL": struct{}{}, "DATA_RETRIEVAL_COMPLETE": struct{}{}}
	var i int
	var loginIdentity *finverse.LoginIdentity
	for i = 0; i < 20; i++ {
		loginIdentityResp, _, err := client.LoginIdentityAPI.GetLoginIdentity(loginIdentityCtx).Execute()
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
		panic("loginIdentity status error")
	}

	// Get Accounts
	accountsResp, _, err := client.LoginIdentityAPI.ListAccounts(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	t.Logf("accounts: %+v", accountsResp.Accounts)

	// Get Transactions with pagination using offset and limit
	offset := 0
	for {
		transactionsResp, _, err := client.LoginIdentityAPI.ListTransactionsByLoginIdentityId(loginIdentityCtx).Offset(int32(offset)).Execute()
		if err != nil {
			panic(err)
		}

		t.Logf("total: %d, transactions: %+v", transactionsResp.TotalTransactions, transactionsResp.Transactions)
		offset += len(transactionsResp.Transactions)

		if offset >= int(transactionsResp.TotalTransactions) {
			break
		}
	}

	// Get Statements metadata
	statementsResp, _, err := client.LoginIdentityAPI.GetStatements(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	t.Logf("statements: %+v", statementsResp.Statements)

	// Get statement binary
	// assume there is one statement
	statementId := statementsResp.Statements[0].Id

	stmtLinkResp, httpResp, err := client.LoginIdentityAPI.GetStatement(loginIdentityCtx, *statementId).Execute()
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode != 200 {
		panic("bad http status")
	}
	defer httpResp.Body.Close()

	links := stmtLinkResp.GetStatementLinks()
	if len(links) == 0 || links[0].GetUrl() == "" {
		panic("expected statement link URL in response")
	}
	t.Logf("statement download URL: %s", links[0].GetUrl())
}
