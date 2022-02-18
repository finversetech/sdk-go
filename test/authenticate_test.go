package test

import (
	"context"
	"testing"

	"github.com/finversetech/sdk-go/finverse"
)

func TestCustomerAccessToken(t *testing.T) {
	client := NewTestClient()

	ctx := context.Background()
	customerTokenResp, err := getCustomerAccessToken(ctx, client)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
		t.FailNow()
	}

	if customerTokenResp.GetAccessToken() == "" {
		t.Errorf("empty access token")
	}
}

func getCustomerAccessToken(ctx context.Context, client *finverse.APIClient) (*finverse.TokenResponse, error) {
	customerTokenResp, _, err := client.PublicApi.GenerateCustomerAccessToken(ctx).TokenRequest(finverse.TokenRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}).Execute()

	return customerTokenResp, err
}
