# Finverse API - Go SDK

## Installation
```
go get -u github.com/finversetech/sdk-go
```


## Getting started

### 1. Authenticate with Finverse API: Obtain Customer Access Token
```go
	// obtain these from https://dashboard.finverse.com
	apiHost := "https://api.sandbox.finverse.net"
	clientId := os.Getenv("FINVERSE_CLIENTID")
	clientSecret := os.Getenv("FINVERSE_SECRET")
	redirectUri := os.Getenv("REDIRECT_URI")

	configuration := finverse.NewConfiguration()
	configuration.Servers = finverse.ServerConfigurations{
		{URL: apiHost},
	}
	configuration.HTTPClient = &http.Client{
		Timeout: 60 * time.Second,
	}
	client := finverse.NewAPIClient(configuration)

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
```

### 2. Link new institution: Obtain Link Token and Link URL to launch Finverse Link UI
```go
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
```

### 3. Finalize linking: Exchange code for Login Identity Access Token
```go
	// when Finverse Link UI is successful, obtain the code from Finverse Link and exchange it for a Login Identity Access Token
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
```

### Retrieve data: Get data using Login Identity Access Token
```go
	loginIdentityCtx := context.WithValue(context.Background(), finverse.ContextAccessToken, loginIdentityToken)

	// get LoginIdentity
	loginIdentityResp, _, err := client.LoginIdentityApi.GetLoginIdentity(loginIdentityCtx).Execute()

	t.Logf("login identity: %+v", loginIdentityResp.LoginIdentity)
	
	// get other products (Accounts, Account Numbers, Transactions)
```
