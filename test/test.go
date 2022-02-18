package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/finversetech/sdk-go/finverse"
)

// obtain these from https://dashboard.finverse.com
// for the test, these will be overriden in initMockServer()
var (
	apiHost      = os.Getenv("FINVERSE_API")
	clientId     = os.Getenv("FINVERSE_CLIENTID")
	clientSecret = os.Getenv("FINVERSE_SECRET")
	redirectUri  = os.Getenv("REDIRECT_URI")
)

func NewTestClient() *finverse.APIClient {
	initMockServer()

	configuration := finverse.NewConfiguration()
	configuration.Servers = finverse.ServerConfigurations{
		{URL: apiHost},
	}
	configuration.HTTPClient = &http.Client{
		Timeout: 60 * time.Second,
	}
	client := finverse.NewAPIClient(configuration)

	return client
}

func initMockServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/customer/token", mockCustomerToken)
	mux.HandleFunc("/link/token", mockLinkToken)
	mux.HandleFunc("/auth/token", mockLoginIdentityToken)
	mux.HandleFunc("/login_identity", mockLoginIdentity)
	mux.HandleFunc("/accounts", mockAccounts)
	mux.HandleFunc("/transactions", mockTransactions)
	mux.HandleFunc("/statements", mockStatements)
	mux.HandleFunc("/statements/", mockStatementDownload)

	ts := httptest.NewServer(mux)
	apiHost = ts.URL
}

func mockCustomerToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`
{
  "access_token": "customer_token",
  "expires_in": 3600,
  "token_type": "Bearer"
}
	`))
}

func mockLinkToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`
{
  "access_token": "link_token",
  "expires_in": 300,
  "link_url": "https://link.sandbox.finverse.net/onboarding?token=link_token",
  "token_type": "Bearer"
}
	`))
}

func mockLoginIdentityToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`
{
  "access_token": "login_identity_token",
  "expires_in": 3600,
  "login_identity_id": "login_identity_id",
  "refresh_token": "login_identity_refresh_token",
  "token_type": "Bearer"
}
	`))
}

func mockLoginIdentity(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`
{
  "institution": {
    "countries": [
      "HKG",
      "SGP",
      "PHL"
    ],
    "institution_id": "testbank",
    "institution_name": "TestBank HK",
    "portal_name": "Test Bank Personal Account"
  },
  "login_identity": {
    "authentication_status": {
      "last_successful_update": "2021-12-03T02:08:23.087Z",
      "last_update": "2021-12-03T02:08:23.087Z",
      "status": "AUTHENTICATED",
      "status_details": "AUTHENTICATED"
    },
    "billing_details": {
      "billed_products": null
    },
    "created_at": "2021-12-03T02:07:58.103Z",
    "customer_app_id": "any_customer_app_id",
    "error": {},
    "first_success": "2021-12-03T02:07:58.103Z",
    "institution_id": "testbank",
    "last_success": "2021-12-03T02:08:23.034Z",
    "login_identity_id": "01FNZ0RP2MWD0SNFBXT8VJ65PE",
    "login_methods_available": {},
    "permissions": null,
    "permissions_expiry_date": "2021-12-03T02:07:58.103Z",
    "permissions_grant_date": "2021-12-03T02:07:58.103Z",
    "product_status": {
      "account_numbers": {
        "last_successful_update": "2021-12-03T02:08:25.189Z",
        "last_update": "2021-12-03T02:08:25.189Z",
        "status": "SUCCESS",
        "status_details": "ACCOUNT_NUMBERS_RETRIEVED"
      },
      "accounts": {
        "last_successful_update": "2021-12-03T02:08:25.050Z",
        "last_update": "2021-12-03T02:08:25.050Z",
        "status": "SUCCESS",
        "status_details": "ACCOUNTS_RETRIEVED"
      },
      "balance_history": {
        "status": "SUCCESS",
        "status_details": "BALANCE_HISTORY_RETRIEVED"
      },
      "historical_transactions": {
        "last_update": "2021-12-03T02:08:33.698Z",
        "status": "SUCCESS",
        "status_details": "HISTORICAL_TRANSACTIONS_RETRIEVED"
      },
      "identity": {
        "last_successful_update": "2021-12-03T02:08:27.898Z",
        "last_update": "2021-12-03T02:08:27.898Z",
        "status": "SUCCESS",
        "status_details": "IDENTITY_RETRIEVED"
      },
      "online_transactions": {
        "last_successful_update": "2021-12-03T02:08:26.790Z",
        "last_update": "2021-12-03T02:08:26.790Z",
        "status": "SUCCESS",
        "status_details": "ONLINE_TRANSACTIONS_RETRIEVED"
      },
      "statements": {
        "last_successful_update": "2021-12-03T02:08:29.092Z",
        "last_update": "2021-12-03T02:08:29.092Z",
        "status": "SUCCESS",
        "status_details": "STATEMENTS_RETRIEVED"
      }
    },
    "status": "DATA_RETRIEVAL_COMPLETE",
    "status_details": {
      "event_date": "2021-12-03T02:08:33.794Z",
      "event_name": "BALANCE_HISTORY_RETRIEVED"
    },
    "updated_at": "2021-12-03T02:08:33.811Z",
    "user_id": "01FN8TJ92X38QF7B4N5HRP8X3K",
    "webhook": "https://example.com/callback"
  }
}
	`))
}

func mockAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`
{
  "accounts": [
    {
      "account_currency": "HKD",
      "account_id": "01F7MP3XTNX36K9N66JPKH131P",
      "account_name": "HKD Checking",
      "balance": {
        "currency": "HKD",
        "raw": "70013.12",
        "value": 70013.12
      },
      "created_at": "2021-06-08T02:09:42.000Z",
      "group_id": "5578758a56843bbe7104f0915f807eae4231aa0e5ce2bc0950c4fcb372b85948",
      "is_closed": false,
      "is_excluded": false,
      "is_parent": false,
      "statement_balance": {
        "currency": "HKD",
        "raw": "70013.12",
        "value": 70013.12
      },
      "updated_at": "2021-06-08T02:09:42.000Z"
    },
    {
      "account_currency": "HKD",
      "account_id": "01F7MP3XTQ4Y8AFQ9KDFQXF14Y",
      "account_name": "HKD Credit Card",
      "balance": {
        "currency": "HKD",
        "raw": "-1833.22",
        "value": -1833.22
      },
      "created_at": "2021-06-08T02:09:42.000Z",
      "group_id": "4f776329e40d2aafaeb33415283aea9e0dae84467aeaa51c5030e4b0311b6040",
      "is_closed": false,
      "is_excluded": false,
      "is_parent": false,
      "statement_balance": {
        "currency": "HKD",
        "raw": "-1833.22",
        "value": -1833.22
      },
      "updated_at": "2021-06-08T02:09:42.000Z"
    },
    {
      "account_currency": "USD",
      "account_id": "01F7MP3XTSBSTG6G68FESC66B1",
      "account_name": "USD FX",
      "balance": {
        "currency": "USD",
        "raw": "1923.22",
        "value": 1923.22
      },
      "created_at": "2021-06-08T02:09:42.000Z",
      "group_id": "582c1d9912586672ef9d028d7ed521ffedf9bb11d2392ddf1c2ec2e231c49ce5",
      "is_closed": false,
      "is_excluded": false,
      "is_parent": false,
      "statement_balance": {
        "currency": "HKD",
        "raw": "15001.116",
        "value": 15001.116
      },
      "updated_at": "2021-06-08T02:09:42.000Z"
    },
    {
      "account_currency": "BTC",
      "account_id": "01F7MP3XTV38VSZMVF4C65NHT3",
      "account_name": "Bitcoin",
      "balance": {
        "currency": "BTC",
        "raw": "420.69",
        "value": 420.69
      },
      "created_at": "2021-06-08T02:09:42.000Z",
      "group_id": "582c1d9912586672ef9d028d7ed521ffedf9bb11d2392ddf1c2ec2e231c49ce5",
      "is_closed": false,
      "is_excluded": false,
      "is_parent": false,
      "statement_balance": {
        "currency": "HKD",
        "raw": "106468292.05",
        "value": 106468292.05
      },
      "updated_at": "2021-06-08T02:09:42.000Z"
    }
  ],
  "institution": {
    "countries": [
      "HKG",
      "SGP",
      "PHL"
    ],
    "institution_id": "testbank",
    "institution_name": "TestBank HK",
    "portal_name": "Test Bank Personal Account"
  },
  "login_identity": {
    "login_identity_id": "01F7MP3J3H485QSDQC0FS15KE7",
    "status": "DATA_RETRIEVAL_COMPLETE"
  }
}
	`))
}

func mockTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`
{
  "accounts": [
    {
      "account_currency": "HKD",
      "account_id": "01FRSAGG1AZ1Z48EQSJ1D16M5V",
      "account_name": "HKD Checking",
      "balance": {
        "currency": "HKD",
        "raw": "70013.12",
        "value": 70013.12
      },
      "created_at": "2022-01-07T03:49:54.000Z",
      "group_id": "01FRSAGG1YCFA2ABEGVATG97SE",
      "is_closed": false,
      "is_excluded": false,
      "is_parent": false,
      "statement_balance": {
        "currency": "HKD",
        "raw": "70013.12",
        "value": 70013.12
      },
      "updated_at": "2022-01-07T03:49:54.000Z"
    },
    {
      "account_currency": "HKD",
      "account_id": "01FRSAGG1E6Y337H4SD0NKD0XN",
      "account_name": "HKD Credit Card",
      "balance": {
        "currency": "HKD",
        "raw": "-1833.22",
        "value": -1833.22
      },
      "created_at": "2022-01-07T03:49:54.000Z",
      "group_id": "01FRSAGG21GJEW1TFMETFGX9FB",
      "is_closed": false,
      "is_excluded": false,
      "is_parent": false,
      "statement_balance": {
        "currency": "HKD",
        "raw": "-1833.22",
        "value": -1833.22
      },
      "updated_at": "2022-01-07T03:49:54.000Z"
    },
    {
      "account_currency": "USD",
      "account_id": "01FRSAGG1J5MJDERMX31EYPVN0",
      "account_name": "USD FX",
      "balance": {
        "currency": "USD",
        "raw": "1923.22",
        "value": 1923.22
      },
      "created_at": "2022-01-07T03:49:54.000Z",
      "group_id": "01FRSAGG24FAGDW039T06XFMSX",
      "is_closed": false,
      "is_excluded": false,
      "is_parent": false,
      "statement_balance": {
        "currency": "HKD",
        "raw": "15001.116",
        "value": 15001.116
      },
      "updated_at": "2022-01-07T03:49:54.000Z"
    },
    {
      "account_currency": "BTC",
      "account_id": "01FRSAGG1NC70799ZNJKJPKK8R",
      "account_name": "Bitcoin",
      "balance": {
        "currency": "BTC",
        "raw": "420.69",
        "value": 420.69
      },
      "created_at": "2022-01-07T03:49:54.000Z",
      "group_id": "01FRSAGG24FAGDW039T06XFMSX",
      "is_closed": false,
      "is_excluded": false,
      "is_parent": false,
      "statement_balance": {
        "currency": "HKD",
        "raw": "106468292.05",
        "value": 106468292.05
      },
      "updated_at": "2022-01-07T03:49:54.000Z"
    },
    {
      "account_currency": "HKD",
      "account_id": "01FRSAGVH6M8ER4WH4XS9GV50R",
      "account_name": "HKD Statement Savings",
      "balance": {
        "currency": "HKD",
        "raw": "19,806.58",
        "value": 19806.58
      },
      "created_at": "2022-01-07T03:50:06.000Z",
      "group_id": "01FRSAGVHCCAVDHY96P0SGJENT",
      "is_closed": false,
      "is_excluded": false,
      "is_parent": false,
      "statement_balance": {
        "currency": "HKD",
        "raw": "19,806.58",
        "value": 19806.58
      },
      "updated_at": "2022-01-07T03:50:06.000Z"
    }
  ],
  "institution": {
    "countries": [
      "HKG",
      "SGP",
      "PHL"
    ],
    "institution_id": "testbank",
    "institution_name": "TestBank HK",
    "portal_name": "Test Bank Personal Account"
  },
  "login_identity": {
    "login_identity_id": "01FRSAFRNSXVKR05KHCASHSN3F",
    "status": "DATA_RETRIEVAL_COMPLETE"
  },
  "total_transactions": 18,
  "transactions": [
    {
      "account_id": "01FRSAGG1AZ1Z48EQSJ1D16M5V",
      "amount": {
        "currency": "HKD",
        "raw": "15000.00",
        "value": 15000
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "Salary",
      "is_pending": false,
      "posted_date": "2020-12-01",
      "transaction_id": "01FRSAGN9VNWZXD7FY6A180XC2",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1E6Y337H4SD0NKD0XN",
      "amount": {
        "currency": "HKD",
        "raw": "-233.00",
        "value": -233
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "NETFLICKS PAYMENT THANKS",
      "is_pending": false,
      "posted_date": "2020-11-22",
      "transaction_id": "01FRSAGNAA2CPE5HFN9A91QMTJ",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1AZ1Z48EQSJ1D16M5V",
      "amount": {
        "currency": "HKD",
        "raw": "-388.00",
        "value": -388
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "BLACK SHEEP DINNER",
      "is_pending": false,
      "posted_date": "2020-11-21",
      "transaction_id": "01FRSAGN9SHA1ZVBW406VNZGAR",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1AZ1Z48EQSJ1D16M5V",
      "amount": {
        "currency": "HKD",
        "raw": "-88.00",
        "value": -88
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "BURGER",
      "is_pending": false,
      "posted_date": "2020-11-17",
      "transaction_id": "01FRSAGN9Q0VC5671A12FX5R9B",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1E6Y337H4SD0NKD0XN",
      "amount": {
        "currency": "HKD",
        "raw": "-150.00",
        "value": -150
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "LUNCH",
      "is_pending": false,
      "posted_date": "2020-11-16",
      "transaction_id": "01FRSAGNA86MWAAQS07F6DN8M5",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1AZ1Z48EQSJ1D16M5V",
      "amount": {
        "currency": "HKD",
        "raw": "-45.00",
        "value": -45
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "COFFEE",
      "is_pending": false,
      "posted_date": "2020-11-13",
      "transaction_id": "01FRSAGN9NTCAJ0TRMK3CSGACK",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1AZ1Z48EQSJ1D16M5V",
      "amount": {
        "currency": "HKD",
        "raw": "-120.12",
        "value": -120.12
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "PIZZA PIZZA",
      "is_pending": false,
      "posted_date": "2020-11-11",
      "transaction_id": "01FRSAGN9KPMZ8KPTGV56QDQPM",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1AZ1Z48EQSJ1D16M5V",
      "amount": {
        "currency": "HKD",
        "raw": "-399.99",
        "value": -399.99
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "HKBN SVC CHR OCT2020",
      "is_pending": false,
      "posted_date": "2020-11-11",
      "transaction_id": "01FRSAGN9G7D6CWB73BYK8EEH5",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1E6Y337H4SD0NKD0XN",
      "amount": {
        "currency": "HKD",
        "raw": "-48.00",
        "value": -48
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "COFFEE",
      "is_pending": false,
      "posted_date": "2020-11-09",
      "transaction_id": "01FRSAGNA66XXHBX7S5Z8H1CWY",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1NC70799ZNJKJPKK8R",
      "amount": {
        "currency": "BTC",
        "raw": "-13.37",
        "value": -13.37
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "Gamestonks!!!",
      "is_pending": false,
      "posted_date": "2020-11-07",
      "transaction_id": "01FRSAGNAGRCA45QZEP330Y4N4",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1J5MJDERMX31EYPVN0",
      "amount": {
        "currency": "USD",
        "raw": "248.00",
        "value": 248
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "USD TRANSFER IN",
      "is_pending": false,
      "posted_date": "2020-11-07",
      "transaction_id": "01FRSAGNAEFRNDCXBHB379SR68",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1E6Y337H4SD0NKD0XN",
      "amount": {
        "currency": "HKD",
        "raw": "-48.00",
        "value": -48
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "COFFEE",
      "is_pending": false,
      "posted_date": "2020-11-07",
      "transaction_id": "01FRSAGNA4PND73E5NQZ7XNE6G",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1E6Y337H4SD0NKD0XN",
      "amount": {
        "currency": "HKD",
        "raw": "-48.00",
        "value": -48
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "COFFEE",
      "is_pending": false,
      "posted_date": "2020-11-07",
      "transaction_id": "01FRSAGNA100YS9ZF10HX7RMD4",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1E6Y337H4SD0NKD0XN",
      "amount": {
        "currency": "HKD",
        "raw": "-48.00",
        "value": -48
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "COFFEE",
      "is_pending": false,
      "posted_date": "2020-11-07",
      "transaction_id": "01FRSAGN9ZVNRWYWP139R9QFH4",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1J5MJDERMX31EYPVN0",
      "amount": {
        "currency": "USD",
        "raw": "123.00",
        "value": 123
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "USD TRANSFER IN",
      "is_pending": false,
      "posted_date": "2020-11-02",
      "transaction_id": "01FRSAGNACT0ZFEG182SJH2EJM",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGG1E6Y337H4SD0NKD0XN",
      "amount": {
        "currency": "HKD",
        "raw": "1839.99",
        "value": 1839.99
      },
      "created_at": "2022-01-07T03:49:32.620Z",
      "description": "PAYMENT THANK YOU",
      "is_pending": false,
      "posted_date": "2020-11-01",
      "transaction_id": "01FRSAGN9XQ51HJH4X5ACTSKQS",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:49:32.620Z"
    },
    {
      "account_id": "01FRSAGVH6M8ER4WH4XS9GV50R",
      "amount": {
        "currency": "HKD",
        "raw": "15.55",
        "value": 15.55
      },
      "created_at": "2022-01-07T03:50:05.921Z",
      "description": "DEP PROMO REBATE\nINT REBATE",
      "is_pending": false,
      "posted_date": "2020-09-10",
      "transaction_id": "01FRSAGVHKTJP9C21ZRES2S889",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:50:05.921Z"
    },
    {
      "account_id": "01FRSAGVH6M8ER4WH4XS9GV50R",
      "amount": {
        "currency": "HKD",
        "raw": "0.01",
        "value": 0.01
      },
      "created_at": "2022-01-07T03:50:05.921Z",
      "description": "CREDIT INTEREST",
      "is_pending": false,
      "posted_date": "2020-08-24",
      "transaction_id": "01FRSAGVHJCG2BXKKJ3EF184DK",
      "transfer_details": null,
      "updated_at": "2022-01-07T03:50:05.921Z"
    }
  ]
}
	`))
}

func mockStatements(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(`
{
  "institution": {
    "countries": [
      "HKG",
      "SGP",
      "PHL"
    ],
    "institution_id": "testbank",
    "institution_name": "TestBank HK",
    "portal_name": "Test Bank Personal Account"
  },
  "login_identity": {
    "login_identity_id": "01F11PR4Z0D85W43PREJ5RXP3K",
    "status": "DATA_RETRIEVAL_COMPLETE"
  },
  "statements": [
    {
      "created_at": "2021-03-18T03:41:08.666Z",
      "date": "2020-09-01",
      "id": "01F11PRDNQQV3HE6K57BVPEH09",
      "name": "Test Bank Statement"
    }
  ]
}
	`))
}

func mockStatementDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/plain") // in real, it could be pdf, csv, etc
	w.Write([]byte(`
this is a statement
	`))
}
