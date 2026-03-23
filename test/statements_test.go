package test

import (
	"context"
	"testing"
)

func TestStatements(t *testing.T) {
	client := NewTestClient()
	loginIdentityToken := "loginIdentityToken"
	loginIdentityCtx := contextWithAccessToken(context.Background(), loginIdentityToken)

	// Get Statements metadata
	statementsResp, _, err := client.LoginIdentityAPI.GetStatements(loginIdentityCtx).Execute()
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
		t.Errorf("expected statement link URL in response")
	}
	t.Logf("statement download URL: %s", links[0].GetUrl())
}
