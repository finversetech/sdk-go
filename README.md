# Finverse API - Go SDK 
Usage: This SDK enables a basic end-to-end backend integration with the Finverse API, including API authentication, institution linking, and data retrieval.

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

### 4. Retrieve data: Get data using Login Identity Access Token
```go
loginIdentityCtx := context.WithValue(context.Background(), finverse.ContextAccessToken, loginIdentityToken)

// get LoginIdentity
loginIdentityResp, _, err := client.LoginIdentityApi.GetLoginIdentity(loginIdentityCtx).Execute()
if err != nil {
    panic(err)
}


t.Logf("login identity: %+v", loginIdentityResp.LoginIdentity)

// get other products (Accounts, Account Numbers, Transactions)
```


### 5. Poll loginIdentityStatus until ready
```go
	// Poll until loginIdentityStatus is ready
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
```

### 6. Get Accounts
```go
	// Get Accounts
	accountsResp, _, err := client.LoginIdentityApi.ListAccounts(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}

	t.Logf("accounts: %+v", accountsResp.Accounts)
```

### 7. Get Transactions
```go
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
```

### 8. Get Statements
```go
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
```
