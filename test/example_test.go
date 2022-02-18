package test

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/finversetech/sdk-go/finverse"
)

func TestAll(t *testing.T) {
	ctx := context.Background()
	client := NewTestClient()

	// obtain customer access token
	customerTokenResp, _, err := client.PublicApi.GenerateCustomerAccessToken(ctx).TokenRequest(finverse.TokenRequest{
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

	// the linkUrl can be used to initiate Finverse Link
	t.Logf("linkUrl: %s", linkTokenResp.LinkUrl)

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
	loginIdentityCtx := context.WithValue(context.Background(), finverse.ContextAccessToken, loginIdentityToken)

	// Get LoginIdentity
	loginIdentityResp, _, err := client.LoginIdentityApi.GetLoginIdentity(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	t.Logf("login identity: %+v", loginIdentityResp.LoginIdentity)

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
		panic("loginIdentity status error")
	}

	// Get Accounts
	accountsResp, _, err := client.LoginIdentityApi.ListAccounts(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	t.Logf("accounts: %+v", accountsResp.Accounts)

	// Get Transactions with pagination using offset and limit
	offset := 0
	for {
		transactionsResp, _, err := client.LoginIdentityApi.ListTransactionsByLoginIdentityId(loginIdentityCtx).Offset(int32(offset)).Execute()
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
	statementsResp, _, err := client.LoginIdentityApi.GetStatements(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	t.Logf("statements: %+v", statementsResp.Statements)

	// Get statement binary
	// assume there is one statement
	statementId := statementsResp.Statements[0].Id

	httpResp, err := client.LoginIdentityApi.GetStatement(loginIdentityCtx, *statementId).Execute()
	if err != nil {
		panic(err)
	}
	if httpResp.StatusCode != 200 {
		panic("bad http status")
	}

	statementRaw, err := ioutil.ReadAll(httpResp.Body)
	httpResp.Body.Close()

	t.Logf("statement: %s", statementRaw)
}
