package test

import (
	"context"
	"testing"
)

func TestTransactions(t *testing.T) {
	client := NewTestClient()
	loginIdentityToken := "loginIdentityToken"
	loginIdentityCtx := contextWithAccessToken(context.Background(), loginIdentityToken)

	// Get Transactions with pagination using offset and limit
	offset := 0
	for {
		transactionsResp, _, err := client.LoginIdentityAPI.ListTransactionsByLoginIdentityId(loginIdentityCtx).Offset(int32(offset)).Execute()
		if err != nil {
			panic(err)
		}

		if transactionsResp.Transactions == nil {
			t.Errorf("empty transactions")
		}

		t.Logf("total: %d, transactions: %+v", transactionsResp.TotalTransactions, transactionsResp.Transactions)
		offset += len(transactionsResp.Transactions)

		if offset >= int(transactionsResp.TotalTransactions) {
			break
		}
	}
}
