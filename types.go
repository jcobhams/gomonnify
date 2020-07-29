package gomonnify

import (
	"net/http"
	"time"
)

type (
	Environment string

	requestAuthType string

	base struct {
		HTTPClient *http.Client
		APIBaseUrl string
		Config     *Config
	}

	disbursements struct {
		*base
	}

	reservedAccounts struct {
		*base
	}

	invoicing struct {
		*base
	}

	general struct {
		*base
	}

	Monnify struct {
		//General          *general
		//Invoicing        *invoicing
		Disbursements    *disbursements
		ReservedAccounts *reservedAccounts
	}

	//Config is used to initialize the Monnify client.
	//Environment - sets the current environment. Sandbox or Live
	//APIKey - well pretty obvious :)
	//SecretKey - same as above
	//RequestTimeout - used to set a deadline on the HTTP requests made. defaults to 5seconds.
	//setting it to 0 to ignores timeout and could make request wait indefinitely (not recommended).
	//DefaultContractCode - used by some endpoints. is not provided in the endpoint method params. Not required.
	Config struct {
		Environment         Environment
		APIKey              string
		SecretKey           string
		RequestTimeout      time.Duration
		DefaultContractCode string
	}

	// Endpoint Responses || Method Return Values
	apiResponseMeta struct {
		RequestSuccessful bool   `json:"requestSuccessful"`
		ResponseMessage   string `json:"responseMessage"`
		ResponseCode      string `json:"responseCode"`
	}

	LoginResponse struct {
		apiResponseMeta
		ResponseBody struct {
			AccessToken string `json:"accessToken"`
			ExpiresIn   int    `json:"expiresIn"`
		}
	}

	ReserveAccountResponse struct {
		apiResponseMeta
		ResponseBody struct {
			ContractCode         string `json:"contractCode"`
			AccountReference     string `json:"accountReference"`
			AccountName          string `json:"accountName"`
			CurrencyCode         string `json:"currencyCode"`
			CustomerEmail        string `json:"customerEmail"`
			CustomerName         string `json:"customerName"`
			AccountNumber        string `json:"accountNumber"`
			BankName             string `json:"bankName"`
			BankCode             string `json:"bankCode"`
			CollectionChannel    string `json:"collectionChannel"`
			ReservationReference string `json:"reservationReference"`
			ReservedAccountType  string `json:"reservedAccountType"`
			Status               string `json:"status"`
			CreatedOn            string `json:"createdOn"`
			IncomeSplitConfig    []struct {
				SubAccountCode  string  `json:"subAccountCode"`
				FeePercentage   float64 `json:"feePercentage"`
				FeeBearer       bool    `json:"feeBearer"`
				SplitPercentage float64 `json:"splitPercentage"`
			} `json:"incomeSplitConfig"`
			RestrictPaymentSource bool `json:"restrictPaymentSource"`
			Contract              struct {
				Name                                       string `json:"name"`
				Code                                       string `json:"code"`
				Description                                string `json:"description"`
				SupportsAdvancedSettlementAccountSelection bool   `json:"supportsAdvancedSettlementAccountSelection"`
				SweepToExternalAccount                     bool   `json:"sweepToExternalAccount"`
			} `json:"contract"`
		} `json:"responseBody"`
	}

	TransactionsResponse struct {
		apiResponseMeta
		ResponseBody struct {
			Content  []Transaction `json:"content"`
			Pageable struct {
				Sort struct {
					Sorted   bool `json:"sorted"`
					Unsorted bool `json:"unsorted"`
					Empty    bool `json:"empty"`
				} `json:"sort"`
				PageSize   int  `json:"pageSize"`
				PageNumber int  `json:"pageNumber"`
				Offset     int  `json:"offset"`
				Unpaged    bool `json:"unpaged"`
				Paged      bool `json:"paged"`
			} `json:"pageable"`
			TotalElements int  `json:"totalElements"`
			TotalPages    int  `json:"totalPages"`
			Last          bool `json:"last"`
			Sort          struct {
				Sorted   bool `json:"sorted"`
				Unsorted bool `json:"unsorted"`
				Empty    bool `json:"empty"`
			} `json:"sort"`
			First            bool `json:"first"`
			NumberOfElements int  `json:"numberOfElements"`
			Size             int  `json:"size"`
			Number           int  `json:"number"`
			Empty            bool `json:"empty"`
		} `json:"responseBody"`
	}

	Transaction struct {
		CustomerDTO struct {
			Email        string `json:"email"`
			Name         string `json:"name"`
			MerchantCode string `json:"merchantCode"`
		} `json:"customerDTO"`
		ProviderAmount       float64 `json:"providerAmount"`
		PaymentMethod        string  `json:"paymentMethod"`
		CreatedOn            string  `json:"createdOn"`
		Amount               float64 `json:"amount"`
		Flagged              bool    `json:"flagged"`
		ProviderCode         string  `json:"providerCode"`
		Fee                  float64 `json:"fee"`
		CurrencyCode         string  `json:"currencyCode"`
		CompletedOn          string  `json:"completedOn"`
		PaymentDescription   string  `json:"paymentDescription"`
		PaymentStatus        string  `json:"paymentStatus"`
		TransactionReference string  `json:"transactionReference"`
		PaymentReference     string  `json:"paymentReference"`
		MerchantCode         string  `json:"merchantCode"`
		MerchantName         string  `json:"merchantName"`
		PayableAmount        float64 `json:"payableAmount"`
		AmountPaid           float64 `json:"amountPaid"`
		Completed            bool    `json:"completed"`
	}

	SingleTransferResponse struct {
		apiResponseMeta
		ResponseBody struct {
			Amount      float64 `json:"amount"`
			Reference   string  `json:"reference"`
			Status      string  `json:"status"`
			DateCreated string  `json:"dateCreated"`
		}
	}

	BulkTransferResponse struct {
		apiResponseMeta
		ResponseBody struct {
			TotalAmount       float64 `json:"totalAmount"`
			TotalFee          float64 `json:"totalFee"`
			BatchReference    string  `json:"batchReference"`
			BatchStatus       string  `json:"batchStatus"`
			TotalTransactions int     `json:"totalTransactions"`
			DateCreated       string  `json:"date_created"`
		}
	}

	SingleTransferDetailsResponse struct {
		apiResponseMeta
		ResponseBody SingleTransferDetails `json:"responseBody"`
	}

	SingleTransferDetails struct {
		Amount        float64 `json:"amount"`
		Reference     string  `json:"reference"`
		Narration     string  `json:"narration"`
		BankCode      string  `json:"bankCode"`
		AccountNumber string  `json:"accountNumber"`
		Currency      string  `json:"currency"`
		AccountName   string  `json:"accountName"`
		BankName      string  `json:"bankName"`
		DateCreated   string  `json:"dateCreated"`
		Fee           float64 `json:"fee"`
		Status        string  `json:"status"`
	}

	BulkTransferDetailsResponse struct {
		apiResponseMeta
		ResponseBody struct {
			Title             string  `json:"title"`
			TotalAmount       float64 `json:"totalAmount"`
			TotalFee          float64 `json:"totalFee"`
			BatchReference    string  `json:"batchReference"`
			TotalTransactions int     `json:"totalTransactions"`
			FailedCount       int     `json:"failedCount"`
			SuccessfulCount   int     `json:"successfulCount"`
			PendingCount      int     `json:"pendingCount"`
			BatchStatus       string  `json:"batchStatus"`
			DateCreated       string  `json:"dateCreated"`
		} `json:"responseBody"`
	}

	TransferTransactionsResponse struct {
		apiResponseMeta
		ResponseBody struct {
			Content  []SingleTransferDetails `json:"content"`
			Pageable struct {
				Sort struct {
					Sorted   bool `json:"sorted"`
					Unsorted bool `json:"unsorted"`
					Empty    bool `json:"empty"`
				} `json:"sort"`
				PageSize   int  `json:"pageSize"`
				PageNumber int  `json:"pageNumber"`
				Offset     int  `json:"offset"`
				Unpaged    bool `json:"unpaged"`
				Paged      bool `json:"paged"`
			} `json:"pageable"`
			TotalElements int  `json:"totalElements"`
			TotalPages    int  `json:"totalPages"`
			Last          bool `json:"last"`
			Sort          struct {
				Sorted   bool `json:"sorted"`
				Unsorted bool `json:"unsorted"`
				Empty    bool `json:"empty"`
			} `json:"sort"`
			First            bool `json:"first"`
			NumberOfElements int  `json:"numberOfElements"`
			Size             int  `json:"size"`
			Number           int  `json:"number"`
			Empty            bool `json:"empty"`
		} `json:"responseBody"`
	}

	ValidAccountNumberResponse struct {
		apiResponseMeta
		ResponseBody struct {
			AccountNumber string `json:"accountNumber"`
			AccountName   string `json:"accountName"`
			BankCode      string `json:"bankCode"`
		} `json:"responseBody"`
	}

	WalletBalanceResponse struct {
		apiResponseMeta
		ResponseBody struct {
			AvailableBalance float64 `json:"availableBalance"`
			LedgerBalance    float64 `json:"ledgerBalance"`
		} `json:"responseBody"`
	}

	ResendOTPResponse struct {
		apiResponseMeta
		ResponseBody struct {
			Message string `json:"message"`
		} `json:"responseBody"`
	}
)
