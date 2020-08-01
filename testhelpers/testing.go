package testhelpers

import (
	"crypto/sha512"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

//Testhelpers contains utility methods to ease third party unit/integration testing
const (
	AccessToken          string  = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsibW9ubmlmeS1wYXltZW50LWVuZ2luZSJdLCJzY29wZSI6WyJwcm9maWxlIl0sImV4cCI6MTU5NTc5MDc1NiwiYXV0aG9yaXRpZXMiOlsiTVBFX01BTkFHRV9MSU1JVF9QUk9GSUxFIiwiTVBFX1VQREFURV9SRVNFUlZFRF9BQ0NPVU5UIiwiTVBFX0lOSVRJQUxJWkVfUEFZTUVOVCIsIk1QRV9SRVNFUlZFX0FDQ09VTlQiLCJNUEVfQ0FOX1JFVFJJRVZFX1RSQU5TQUNUSU9OIiwiTVBFX1JFVFJJRVZFX1JFU0VSVkVEX0FDQ09VTlQiLCJNUEVfREVMRVRFX1JFU0VSVkVEX0FDQ09VTlQiLCJNUEVfUkVUUklFVkVfUkVTRVJWRURfQUNDT1VOVF9UUkFOU0FDVElPTlMiXSwianRpIjoiODE0NzI1YWItOTUwZi00MzBkLWI2NWQtMzliNzg5OTI5Njg2IiwiY2xpZW50X2lkIjoiTUtfVEVTVF9MVzJRV0YyMjJVIn0.r5zZXaomV6AfnRq14VjA1eSyOT2h3owN8TfSFmJ9lhRroitiFOFtuX8LztwKFswvQj5wOL_dyoInmORIBOboiV_ukPPHM3J1Nco9OqG8SNj8wx2-DgI5JhxoIpAmABjw5zbCYXyu9YB92KrICzrVqPaIhn_IprSAResufzNa92n91_qc5a0NdCWyA-WIQflMfBYeunCuIjhC91Yf2HfAjZZ-_2l2GsZdjwfXq4RldHFKuWJf3lWp4r6K9yJmZKLce2syph7kj9Kh0CYwFxQzVKfC1EQkCkFfLYJBgXkqQUEtLdtvA2JApvom4wvWEiM4vmJE29z_O2CiUKf_SJvjOQ"
	ContractCode         string  = "4934121686"
	CustomerName         string  = "John Doe"
	AccountName          string  = "Test Account"
	AccountReference     string  = "TEST_ACCT_REF"
	CurrencyCode         string  = "NGN"
	CustomerEmail        string  = "test@tester.com"
	AccountNumber        string  = "3000017736"
	BankName             string  = "Providus Bank"
	BankCode             string  = "101"
	CollectionChannel    string  = "RESERVED_ACCOUNT"
	ReservationReference string  = "EPC3RD5ULJFN5QB8A8AT"
	ReservedAccountType  string  = "GENERAL"
	StatusActive         string  = "ACTIVE"
	CreatedOn            string  = "2020-07-26 19:24:39.113"
	MerchantCode         string  = "ALJKHDALASD"
	ProviderAmount       float64 = 0.21
	PaymentMethod        string  = "ACCOUNT_TRANSFER"
	Amount               float64 = 100.00
	ProviderCode         string  = "98271"
	TransferReference    string  = "TEST_TRF_REF"
	WalletId             string  = "TEST_WLT_ID"
	BatchReference       string  = "TEST_BCH_REF"
	ValidOTP             string  = "111111"
	AvailableBalance     float64 = 500.99
	LedgerBalance        float64 = 500.98
	PaymentReference     string  = "330854835"
	PaidOn               string  = "26/02/2020 09:38:13 AM"
	SecretKey            string  = "SECRET_KEY"
)

func mockLoginResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "accessToken": "%v",
        "expiresIn": 3599
    }
}`, AccessToken)
}

func mockReserveAccountResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "contractCode": "%v",
        "accountReference": "%v",
        "accountName": "%v",
        "currencyCode": "%v",
        "customerEmail": "%v",
        "customerName": "%v",
        "accountNumber": "%v",
        "bankName": "%v",
        "bankCode": "%v",
        "collectionChannel": "%v",
        "reservationReference": "%v",
        "reservedAccountType": "%v",
        "status": "%v",
        "createdOn": "%v",
        "incomeSplitConfig": [],
        "restrictPaymentSource": false,
		"contract": {
            "name": "Default Contract",
            "code": "%v",
            "description": null,
            "supportsAdvancedSettlementAccountSelection": false,
            "sweepToExternalAccount": false
        }
    }
}`, ContractCode, AccountReference, AccountName, CurrencyCode, CustomerEmail, CustomerName, AccountNumber, BankName,
		BankCode, CollectionChannel, ReservationReference, ReservedAccountType, StatusActive, CreatedOn, ContractCode)
}

func mockTransactionsResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "content": [
            {
                "customerDTO": {
                    "email": "%v",
                    "name": "%v",
                    "merchantCode": "%v"
                },
                "providerAmount": %v,
                "paymentMethod": "%v",
                "createdOn": "%v",
                "amount": %v,
                "flagged": false,
                "providerCode": "%v",
                "fee": 0.79,
                "currencyCode": "NGN",
                "completedOn": "2019-07-24T14:12:28.000+0000",
                "paymentDescription": "Test Reserved Account",
                "paymentStatus": "PAID",
                "transactionReference": "MNFY|20190724141227|003374",
                "paymentReference": "MNFY|20190724141227|003374",
                "merchantCode": "ALJKHDALASD",
                "merchantName": "Test Limited",
                "payableAmount": 100.00,
                "amountPaid": 100.00,
                "completed": true
            }
        ],
        "pageable": {
            "sort": {
                "sorted": true,
                "unsorted": false,
                "empty": false
            },
            "pageSize": 10,
            "pageNumber": 0,
            "offset": 0,
            "unpaged": false,
            "paged": true
        },
        "totalElements": 1,
        "totalPages": 1,
        "last": true,
        "sort": {
            "sorted": true,
            "unsorted": false,
            "empty": false
        },
        "first": true,
        "numberOfElements": 1,
        "size": 10,
        "number": 0,
        "empty": false
    }
}`, CustomerEmail, CustomerName, MerchantCode, ProviderAmount, PaymentMethod, CreatedOn, Amount, ProviderCode)
}

func mockSingleTransferResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "amount": %v,
        "reference": "%v",
        "status": "SUCCESS",
        "dateCreated": "13/11/2019 09:34:32 PM"
    }
}`, Amount, TransferReference)
}

func mockBulkTransferResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "totalAmount": %v,
        "totalFee": 8.48,
        "batchReference": "%v",
        "batchStatus": "COMPLETED",
        "totalTransactions": 1,
        "dateCreated": "13/11/2019 09:42:06 PM"
    }
}`, Amount, BatchReference)
}

func mockSingleTransferDetailsResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "amount": %v,
        "reference": "%v",
        "narration": "911 Transaction",
        "bankCode": "%v",
        "accountNumber": "%v",
        "currency": "NGN",
        "accountName": "MEKILIUWA, SMART CHINONSO",
        "bankName": "%v",
        "dateCreated": "13/11/2019 09:42:07 PM",
        "fee": 1.00,
        "status": "SUCCESS"
    }
}`, Amount, TransferReference, BankCode, AccountNumber, BankName)
}

func mockBulkTransferDetailsResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "title": "Final Batch - Continue on Failure",
        "totalAmount": %v,
        "totalFee": 8.48,
        "batchReference": "%v",
        "totalTransactions": 3,
        "failedCount": 0,
        "successfulCount": 0,
        "pendingCount": 3,
        "batchStatus": "AWAITING_PROCESSING",
        "dateCreated": "13/11/2019 10:45:08 PM"
    }
}`, Amount, BatchReference)
}

func mockTransferTransactionsResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "content": [
            {
                "amount": %v,
                "reference": "%v",
                "narration": "911 Transaction",
                "bankCode": "%v",
                "accountNumber": "%v",
                "currency": "NGN",
                "accountName": "MEKILIUWA, SMART CHINONSO",
                "bankName": "%v",
                "dateCreated": "13/11/2019 10:45:08 PM",
                "fee": 5.20,
                "status": "PENDING"
            }
        ],
        "pageable": {
            "sort": {
                "sorted": false,
                "unsorted": true,
                "empty": true
            },
            "pageSize": 10,
            "pageNumber": 0,
            "offset": 0,
            "paged": true,
            "unpaged": false
        },
        "totalPages": 1,
        "totalElements": 3,
        "last": true,
        "sort": {
            "sorted": false,
            "unsorted": true,
            "empty": true
        },
        "first": true,
        "numberOfElements": 3,
        "size": 10,
        "number": 0,
        "empty": false
    }
}`, Amount, TransferReference, BankCode, AccountNumber, BankName)
}

func mockValidateAccountNumberResponseData() string {
	return fmt.Sprintf(`
{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "accountNumber": "%v",
        "accountName": "%v",
        "bankCode": "%v"
    }
}`, AccountNumber, AccountName, BankCode)
}

func mockWalletBalanceResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "availableBalance": %v,
        "ledgerBalance": %v
    }
}`, AvailableBalance, LedgerBalance)
}

func mockResendOTPResponseData() string {
	return `{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "message": "Authorization code will be processed and sent to predefined email addresses(s)"
    }
}`
}

func mockTransactionStatusResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": {
        "transactionReference": "%v",
        "paymentReference": "%v",
        "amountPaid": "%v",
        "totalPayable": "%v",
        "settlementAmount": "99.21",
        "paidOn": "%v",
		"transactionHash": "%v",
        "paymentStatus": "PAID",
        "paymentDescription": "LahrayWeb",
        "currency": "NGN",
        "paymentMethod": "ACCOUNT_TRANSFER",
        "product": {
            "type": "WEB_SDK",
            "reference": "330854835"
        },
        "cardDetails": null,
        "accountDetails": {
            "accountName": "DAMILARE SAMUEL OGUNNAIKE",
            "accountNumber": "******7503",
            "bankCode": "000001",
            "amountPaid": "100.00"
        },
        "accountPayments": [
            {
                "accountName": "%v",
                "accountNumber": "******7503",
                "bankCode": "000001",
                "amountPaid": "%v"
            }
        ],
        "customer": {
            "email": "%v",
            "name": "%v"
        },
        "metaData": {
            "name": "Damilare",
            "age": "45"
        }
    }
}`, TransferReference, PaymentReference, Amount, Amount, PaidOn, GenerateTransactionHash(SecretKey), AccountName, Amount, CustomerEmail, CustomerName)
}

func mockGetBanksResponseData() string {
	return fmt.Sprintf(`{
    "requestSuccessful": true,
    "responseMessage": "success",
    "responseCode": "0",
    "responseBody": [
        {
            "name": "Access bank",
            "code": "044",
            "ussdTemplate": "*901*Amount*AccountNumber#",
            "baseUssdCode": "*901#",
            "transferUssdTemplate": "*901*AccountNumber#"
        },
        {
            "name": "Coronation Bank",
            "code": "559",
            "ussdTemplate": null,
            "baseUssdCode": null,
            "transferUssdTemplate": null
        }
	]
}`)
}

func GenerateTransactionHash(secretKey string) string {
	rawStr := fmt.Sprintf("%v|%v|%v|%v|%v", secretKey, PaymentReference, Amount, PaidOn, TransferReference)
	h := sha512.New()
	h.Write([]byte(rawStr))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//MockAPIServer initializes a test HTTP server useful for request mocking, Integration tests and Client configuration
func MockAPIServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		switch r.URL.Path {

		case "/v1/auth/login":
			switch r.Method {
			case http.MethodPost:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockLoginResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: POST request expected in Login() method or /auth/login endpoint, Got: %v", r.Method)
			}

		case "/v1/bank-transfer/reserved-accounts":
			switch r.Method {
			case http.MethodPost:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockReserveAccountResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: POST request expected in reservedAccounts.ReserveAccount() method or /bank-transfer/reserved-accounts endpoint, Got: %v", r.Method)
			}

		case fmt.Sprintf("/v1/bank-transfer/reserved-accounts/%v", AccountReference):
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockReserveAccountResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: GET request expected in reservedAccounts.Details() method or /bank-transfer/reserved-accounts/{{accountReference}} endpoint, Got: %v", r.Method)
			}

		case fmt.Sprintf("/v1/bank-transfer/reserved-accounts/%v", AccountNumber):
			switch r.Method {
			case http.MethodDelete:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockReserveAccountResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: DELETE request expected in reservedAccounts.Deallocate() method or /bank-transfer/reserved-accounts/{{accountNumber}} endpoint, Got: %v", r.Method)
			}

		case "/v1/bank-transfer/reserved-accounts/transactions":
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockTransactionsResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: GET request expected in reservedAccounts.Transactions() method or /bank-transfer/reserved-accounts/transactions endpoint, Got: %v", r.Method)
			}

		case "/v1/disbursements/single":
			switch r.Method {
			case http.MethodPost:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockSingleTransferResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: POST request expected in disbursements.SingleTransfer() method or /disbursements/single endpoint, Got: %v", r.Method)
			}

		case "/v1/disbursements/batch":
			switch r.Method {
			case http.MethodPost:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockBulkTransferResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: POST request expected in disbursements.BulkTransfer() method or /disbursements/batch endpoint, Got: %v", r.Method)
			}

		case "/v1/disbursements/single/validate-otp":
			switch r.Method {
			case http.MethodPost:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockSingleTransferResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: POST request expected in disbursements.AuthorizeSingleTransfer() method or /disbursements/single/validate-otp endpoint, Got: %v", r.Method)
			}

		case "/v1/disbursements/batch/validate-otp":
			switch r.Method {
			case http.MethodPost:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockBulkTransferResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: POST request expected in disbursements.AuthorizeBulkTransfer() method or /disbursements/batch/validate-otp endpoint, Got: %v", r.Method)
			}

		case "/v1/disbursements/single/summary":
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockSingleTransferDetailsResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: GET request expected in disbursements.SingleTransferDetails() method or /disbursements/single/summary endpoint, Got: %v", r.Method)
			}

		case "/v1/disbursements/batch/summary":
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockBulkTransferDetailsResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: GET request expected in disbursements.SingleTransferDetails() method or /disbursements/single/summary endpoint, Got: %v", r.Method)
			}

		case fmt.Sprintf("/v1/disbursements/bulk/%v/transactions", BatchReference):
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockTransferTransactionsResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: GET request expected in disbursements.BulkTransferDetails() method or /disbursements/bulk/{{batchReference}}/transactions endpoint, Got: %v", r.Method)
			}

		case "/v1/disbursements/single/transactions":
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockTransferTransactionsResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: GET request expected in disbursements.SingleTransferDetails() method or /disbursements/single/transactions endpoint, Got: %v", r.Method)
			}

		case "/v1/disbursements/account/validate":
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockValidateAccountNumberResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: GET request expected in disbursements.ValidateAccountNumber() method or /disbursements/account/validate endpoint, Got: %v", r.Method)
			}

		case "/v1/disbursements/wallet-balance":
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockWalletBalanceResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: GET request expected in disbursements.WalletBalance() method or /disbursements/wallet-balance endpoint, Got: %v", r.Method)
			}

		case "/v1/disbursements/single/resend-otp":
			switch r.Method {
			case http.MethodPost:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockResendOTPResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: POST request expected in disbursements.ResendOTP() method or /disbursements/single/resend-otp endpoint, Got: %v", r.Method)
			}

		case fmt.Sprintf("/v2/transactions/%v", TransferReference):
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockTransactionStatusResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: GET request expected in general.GetTransaction() method or /v2/transactions/{{reference}}, Got: %v", r.Method)
			}

		case "/v1/banks":
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(200)
				fmt.Fprintf(w, mockGetBanksResponseData())
			default:
				log.Fatalf("gomonnify.testhelpers: GET request expected in general.GetBanks() method or /v1/banks, Got: %v", r.Method)
			}

		}

	}))
	return server
}
