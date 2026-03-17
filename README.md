# Finverse API - Go SDK

This SDK enables backend integration with the Finverse API, including authentication, institution linking, data retrieval, and payment operations.

## Installation

```
go get -u github.com/finversetech/sdk-go
```

## Client Setup

Create and configure the API client. Use a single shared client instance across your application.

```go
import (
	"net/http"
	"os"
	"time"

	"github.com/finversetech/sdk-go/finverse"
)

// Obtain credentials from https://dashboard.finverse.com
apiHost := "https://api.prod.finverse.net"
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
```

---

## Payment API

The Payment API supports mandates, payments, payment accounts, payment links, payouts, and related operations. Use the **Customer Access Token** for authentication.

### Authentication

Obtain a Customer Access Token (required for Payment API):

```go
ctx := context.Background()
customerTokenResp, _, err := client.PublicApi.GenerateCustomerAccessToken(ctx).TokenRequest(finverse.TokenRequest{
	ClientId:     clientId,
	ClientSecret: clientSecret,
	GrantType:    "client_credentials",
}).Execute()
if err != nil {
	panic(err)
}
customerAccessToken := customerTokenResp.AccessToken
customerCtx := context.WithValue(ctx, finverse.ContextAccessToken, customerAccessToken)
```

**Example: Create a payment**

```go
paymentResp, _, err := client.CustomerApi.CreatePayment(customerCtx).CreatePaymentRequest(finverse.CreatePaymentRequest{
	Amount:    10000, // in minor units (e.g., 100.00 HKD = 10000)
	Currency:  "HKD",
	// ... other required fields
}).Execute()
if err != nil {
	panic(err)
}
```

**Example: List payment accounts**

```go
accountsResp, _, err := client.CustomerApi.ListPaymentAccounts(customerCtx).Execute()
if err != nil {
	panic(err)
}
```

### Online Payment Flow (Payment Link)

Use payment links to collect one-time or recurring payments via redirect. The flow: **Create** → **Redirect** → **Callback** → **Poll** → **Result**.

#### Flow Overview

1. **Create** – Create a payment link via `DefaultAPI.CreatePaymentLink`, get `url` and `payment_link_id`. Save `payment_link_id` and `unique_reference_id`.
2. **Redirect** – Redirect the user to the payment link URL (use HTTP 303 so the redirect is followed with GET)
3. **Callback** – Finverse redirects back with `payment_link_id` and `unique_reference_id` query params. Verify the values match what was stored earlier.
4. **Poll** – Poll `GetPaymentLink` every 2 seconds for up to 30 seconds until `session_status` is `"COMPLETE"`
5. **Result** – If `COMPLETE` → success page; otherwise → error page. Log `payment_id` and `payment_link_id` on success.

#### Create Payment Link (PAYMENT mode)

For one-time payments, use `mode: "PAYMENT"` with `amount` in minor units (e.g., 100.00 HKD = 10000):

```go
callbackURL := os.Getenv("CALLBACK_URL") // e.g. https://yoursite.com/callback
uniqueRefID := "order-12345"             // use a globally unique ID (e.g. UUID)
amount := int32(10000) // 100.00 HKD
uiMode := "redirect"
email := "john@example.com"

linkResp, _, err := client.DefaultApi.CreatePaymentLink(customerCtx).CreatePaymentLinkRequest(finverse.CreatePaymentLinkRequest{
	Mode:              "PAYMENT",
	Amount:            &amount,
	Currency:          "HKD",
	UniqueReferenceId: uniqueRefID,
	Sender: finverse.PaymentLinkSender{
		ExternalUserId: "your-internal-user-id",
		Name:           "John Doe",
		Email:          &email,
	},
	PaymentDetails: &finverse.PaymentLinkDetails{
		Description:                  "Order #12345",
		ExternalTransactionReference: "ORD-12345", // max 35 chars
	},
	LinkCustomizations: &finverse.PaymentLinkCustomizations{
		UiMode:      &uiMode,
		RedirectUri: &callbackURL,
	},
}).Execute()
if err != nil {
	panic(err)
}

// Save payment_link_id and unique_reference_id for callback validation
paymentLinkID := linkResp.GetPaymentLinkId()
// Redirect user to linkResp.GetUrl() (use HTTP 303 for GET redirect)
```

#### Create Payment Link (SETUP mode)

For saving a payment method for future use (e.g., Click to Pay), use `mode: "SETUP"`. Do **not** pass `amount`:

```go
callbackURL := os.Getenv("CALLBACK_URL")
uniqueRefID := "setup-001" // globally unique
uiMode := "redirect"
futurePayments := "CLICK_TO_PAY"
email := "john@example.com"

linkResp, _, err := client.DefaultApi.CreatePaymentLink(customerCtx).CreatePaymentLinkRequest(finverse.CreatePaymentLinkRequest{
	Mode:              "SETUP",
	Currency:           "HKD",
	UniqueReferenceId:  uniqueRefID,
	Sender: finverse.PaymentLinkSender{
		ExternalUserId: "your-internal-user-id",
		Name:           "John Doe",
		Email:          &email,
	},
	PaymentDetails: &finverse.PaymentLinkDetails{
		Description:                  "Save payment method",
		ExternalTransactionReference: "SETUP-001",
	},
	LinkCustomizations: &finverse.PaymentLinkCustomizations{
		UiMode:      &uiMode,
		RedirectUri: &callbackURL,
	},
	PaymentSetupOptions: &finverse.PaymentSetupOptionsRequest{
		FuturePayments: &futurePayments,
	},
}).Execute()
```

#### Callback Handler and Polling

In your callback handler:

1. Read `payment_link_id` and `unique_reference_id` from query params (required)
2. Optionally validate against your stored records
3. Poll until `session_status` is `"COMPLETE"` or timeout

```go
// In your callback handler (e.g. /callback?payment_link_id=xxx&unique_reference_id=yyy)
paymentLinkID := r.URL.Query().Get("payment_link_id")
if paymentLinkID == "" {
	http.Error(w, "missing payment_link_id", http.StatusBadRequest)
	return
}
uniqueRefID := r.URL.Query().Get("unique_reference_id")
// Optionally validate paymentLinkID and uniqueRefID match what you stored when creating the link

// Poll for up to 30 seconds
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

var linkResp *finverse.PaymentLinkResponse
for i := 0; i < 15; i++ {
	linkResp, _, err = client.DefaultApi.GetPaymentLink(customerCtx, paymentLinkID).Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if linkResp.GetSessionStatus() == "COMPLETE" {
		// Success: log paymentId and paymentLinkId, then redirect
		paymentID := ""
		if linkResp.Payment != nil {
			paymentID = linkResp.Payment.GetPaymentId()
		}
		log.Printf("Payment complete: payment_id=%s payment_link_id=%s", paymentID, paymentLinkID)
		http.Redirect(w, r, "/success?payment_id="+paymentID, http.StatusSeeOther)
		return
	}
	time.Sleep(2 * time.Second)
}

// Timeout: redirect to error page
http.Redirect(w, r, "/error", http.StatusSeeOther)
```
---

## Data API

The Data API supports retrieving financial data (accounts, transactions, statements, identity, income) from linked institutions. Use the **Login Identity Access Token** for authentication, obtained after completing the institution linking flow.

### Authentication & Institution Linking

Data retrieval requires a Login Identity Access Token. Complete the following flow:

**1. Obtain Customer Access Token**

```go
ctx := context.Background()
customerTokenResp, _, err := client.PublicApi.GenerateCustomerAccessToken(ctx).TokenRequest(finverse.TokenRequest{
	ClientId:     clientId,
	ClientSecret: clientSecret,
	GrantType:    "client_credentials",
}).Execute()
if err != nil {
	panic(err)
}
customerAccessToken := customerTokenResp.AccessToken
customerCtx := context.WithValue(ctx, finverse.ContextAccessToken, customerAccessToken)
```

**2. Generate Link Token and launch Finverse Link UI**

```go
userId := "someUserId"     // reference to your system userId
state := "someUniqueState" // sent in redirectUri callback
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
// Use linkTokenResp.LinkUrl to initiate Finverse Link
```

**3. Exchange code for Login Identity Access Token**

```go
// After Finverse Link completes, obtain the code from the redirectUri callback
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
loginIdentityToken := loginIdentityTokenResp.AccessToken
loginIdentityCtx := context.WithValue(ctx, finverse.ContextAccessToken, loginIdentityToken)
```

### Data API Usage Examples

**Poll login identity status until ready**

```go
finalStatuses := map[string]struct{}{
	"ERROR": {},
	"DATA_RETRIEVAL_PARTIALLY_SUCCESSFUL": {},
	"DATA_RETRIEVAL_COMPLETE": {},
}
for i := 0; i < 20; i++ {
	loginIdentityResp, _, err := client.LoginIdentityApi.GetLoginIdentity(loginIdentityCtx).Execute()
	if err != nil {
		panic(err)
	}
	loginIdentity := loginIdentityResp.LoginIdentity
	if _, ok := finalStatuses[loginIdentity.GetStatus()]; ok {
		break
	}
	time.Sleep(3 * time.Second)
}
if loginIdentity.GetStatus() == "ERROR" {
	panic("loginIdentity status error")
}
```

**Get accounts**

```go
accountsResp, _, err := client.LoginIdentityApi.ListAccounts(loginIdentityCtx).Execute()
if err != nil {
	panic(err)
}
```

**Get transactions (with pagination)**

```go
offset := 0
for {
	transactionsResp, _, err := client.LoginIdentityApi.ListTransactionsByLoginIdentityId(loginIdentityCtx).Offset(int32(offset)).Execute()
	if err != nil {
		panic(err)
	}
	offset += len(transactionsResp.Transactions)
	if offset >= int(transactionsResp.TotalTransactions) {
		break
	}
}
```

**Get statements**

```go
statementsResp, _, err := client.LoginIdentityApi.GetStatements(loginIdentityCtx).Execute()
if err != nil {
	panic(err)
}

// Download statement binary
statementId := statementsResp.Statements[0].Id
_, httpResp, err := client.LoginIdentityApi.GetStatement(loginIdentityCtx, *statementId).Execute()
if err != nil {
	panic(err)
}
defer httpResp.Body.Close()
statementRaw, err := io.ReadAll(httpResp.Body)
```

**Get institution details**

```go
institutionResp, _, err := client.CustomerApi.GetInstitution(customerCtx, "institutionId").Execute()
if err != nil {
	panic(err)
}
```

---

## Best Practices

### Context and Timeouts

- Pass `context.Context` to all API calls for cancellation and deadlines.
- Use `context.WithTimeout` for long-running operations (e.g., polling, batch fetches):

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
defer cancel()
resp, _, err := client.LoginIdentityApi.GetLoginIdentity(ctx).Execute()
```

### HTTP Client Configuration

- Set an appropriate `Timeout` on `http.Client` to avoid hanging requests. The default in examples is 60 seconds.
- Reuse a single `APIClient` instance; it is safe for concurrent use.

### Error Handling

- Always check `err != nil` before using response data.
- Inspect `*http.Response` when needed (e.g., status codes for raw downloads).
- Handle API-specific error types for retries or user-facing messages.

### Response Body Cleanup

- When reading raw response bodies (e.g., `GetStatement`), always close the body to avoid leaks:

```go
_, httpResp, err := client.LoginIdentityApi.GetStatement(ctx, statementId).Execute()
if err != nil {
	return err
}
defer httpResp.Body.Close()
data, err := io.ReadAll(httpResp.Body)
```

### Token Handling

- Use `context.WithValue(ctx, finverse.ContextAccessToken, token)` to pass tokens.
- Never log or expose access tokens.
- Store credentials (client ID, secret) in environment variables or a secrets manager.

### Pagination

- Use `Offset` and `Limit` for paginated endpoints (e.g., `ListTransactionsByLoginIdentityId`).
- Iterate until `offset >= totalTransactions` or the response is empty.

### Concurrent Usage

- `APIClient` is designed for concurrent use. Share one client across goroutines.
- Use separate contexts per request when running operations in parallel.

### Payment Idempotency

Use idempotency keys to safely retry payment operations without creating duplicates. Generate a unique key per logical operation (e.g., per order or transaction) and reuse it for retries.

**CreatePayment** and **CreateMandate** support optional `IdempotencyKey`:

```go
idempotencyKey := "order-12345-" + uuid.New().String() // unique per logical payment

paymentResp, _, err := client.CustomerApi.CreatePayment(customerCtx).
	IdempotencyKey(idempotencyKey).
	CreatePaymentRequest(finverse.CreatePaymentRequest{
		Amount:   &amount,
		Currency: "HKD",
		// ... other fields
	}).Execute()
```

**CreateMandateForExistingSender** and **CreateScheduledPayout** require `IdempotencyKey`:

```go
mandateResp, _, err := client.DefaultApi.CreateMandateForExistingSender(customerCtx).
	IdempotencyKey(idempotencyKey).
	CreateMandateWithSenderAccountRequest(req).Execute()
```

**Payment links** use `unique_reference_id` as the idempotency identifier—use a globally unique value (e.g., order ID or UUID) and store it to validate callbacks.

- Use the same idempotency key for all retries of the same logical operation.
- Keys should be unique per operation; do not reuse across different payments or mandates.

---

## Resources

- **[Finverse AI Skills](https://github.com/finversetech/ai)** – Cursor Agent Skills for Finverse integrations. Contains implementation guides for payment flows, API patterns, and best practices. AI agents can use these skills when helping build Finverse integrations.
