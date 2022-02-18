package test

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/finversetech/sdk-go/finverse"
)

func TestStatements(t *testing.T) {
	client := NewTestClient()
	loginIdentityToken := "loginIdentityToken"
	loginIdentityCtx := context.WithValue(context.Background(), finverse.ContextAccessToken, loginIdentityToken)

	// Get Statements metadata
	statementsResp, _, err := client.LoginIdentityApi.GetStatements(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	if statementsResp.Statements == nil {
		t.Errorf("empty statements")
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
	if len(statementRaw) == 0 {
		t.Errorf("empty statement body")
	}
}
