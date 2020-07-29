package params

const (
	ValidationFailedContinue ValidationFailedOption = "CONTINUE"
	ValidationFailedBreak    ValidationFailedOption = "BREAK"
	NotificationInterval10   NotificationInterval   = 10
	NotificationInterval20   NotificationInterval   = 20
	NotificationInterval50   NotificationInterval   = 50
	NotificationInterval100  NotificationInterval   = 100

	CurrencyNGN Currency = "NGN"
)

type (
	Currency               string
	ValidationFailedOption string
	NotificationInterval   int
	ReserveAccountParam    struct {
		AccountReference      string   `json:"accountReference, omitempty"`
		AccountName           string   `json:"accountName, omitempty"`
		CurrencyCode          Currency `json:"currencyCode, omitempty"`
		ContractCode          string   `json:"contractCode, omitempty"`
		CustomerEmail         string   `json:"customerEmail, omitempty"`
		CustomerName          string   `json:"customerName, omitempty"`
		RestrictPaymentSource bool     `json:"restrictPaymentSource, omitempty"`
		incomeSplitConfig     []IncomeSplitConfigParam
		AllowedPaymentSources AllowedPaymentSourcesParam
	}

	IncomeSplitConfigParam struct {
		SubAccountCode  string  `json:"subAccountCode, omitempty"`
		FeePercentage   float64 `json:"feePercentage, omitempty"`
		SplitPercentage float64 `json:"splitPercentage, omitempty"`
		FeeBearer       bool    `json:"feeBearer, omitempty"`
	}

	AllowedPaymentSourcesParam struct {
		BankAccounts []struct {
			AccountNumber string `json:"accountNumber, omitempty"`
			BankCode      string `json:"bankCode, omitempty"`
		} `json:"bankAccounts, omitempty"`

		AccountNames []string `json:"accountNames, omitempty"`
	}

	SingleTransferParam struct {
		Amount        float64  `json:"amount"`
		Reference     string   `json:"reference"`
		Narration     string   `json:"narration"`
		BankCode      string   `json:"bankCode"`
		AccountNumber string   `json:"accountNumber"`
		Currency      Currency `json:"currency"`
		WalletId      string   `json:"walletId,omitempty"`
	}

	BulkTransferParam struct {
		Title                string                 `json:"title"`
		BatchReference       string                 `json:"batchReference"`
		Narration            string                 `json:"narration"`
		WalletId             string                 `json:"walletId"`
		OnValidationFailure  ValidationFailedOption `json:"onValidationFailure"`
		NotificationInterval NotificationInterval   `json:"notificationInterval"`
		TransactionList      []SingleTransferParam  `json:"transactionList"`
	}
)
