package gomonnify

import (
	"errors"
	"fmt"
	"github.com/jcobhams/gomonnify/params"
	"time"
)

const (
	EnvSandbox Environment = "sandbox" //Sandbox environment for development
	EnvLive    Environment = "live"    //Live environmrnt
	EnvTest    Environment = "test"    //Test environment used during unit/integration testing

	SandBoxAPIKey       string = "MK_TEST_SAF7HR5F3F"
	SandBoxSecretKey    string = "4SY6TNL8CK3VPRSBTHTRG2N8XXEGC6NL"
	DefaultContractCode string = "4934121686"

	APIBaseUrlSandbox string = "https://sandbox.monnify.com/api"
	APIBaseUrlLive    string = "https://api.monnify.com/api"

	RequestTimeout time.Duration = 5 * time.Second

	requestAuthTypeBasic  requestAuthType = "basic"
	requestAuthTypeBearer requestAuthType = "bearer"

	CurrencyNGN = params.CurrencyNGN

	ValidationFailedContinue = params.ValidationFailedContinue
	ValidationFailedBreak    = params.ValidationFailedBreak

	NotificationInterval10  = params.NotificationInterval10
	NotificationInterval20  = params.NotificationInterval20
	NotificationInterval50  = params.NotificationInterval50
	NotificationInterval100 = params.NotificationInterval100

	PaymentStatusPaid          string = "PAID"
	PaymentStatusPending       string = "PENDING"
	PaymentStatusOverpaid      string = "OVERPAID"
	PaymentStatusPartiallyPaid string = "PARTIALLY_PAID"
	PaymentStatusExpired       string = "EXPIRED"
	PaymentStatusFailed        string = "FAILED"
	PaymentStatusCancelled     string = "CANCELLED"
)

var (
	DefaultConfig = Config{
		Environment:         EnvSandbox,
		APIKey:              SandBoxAPIKey,
		SecretKey:           SandBoxSecretKey,
		RequestTimeout:      RequestTimeout,
		DefaultContractCode: DefaultContractCode,
	}

	AuthToken string
	ExpiresIn time.Time
)

//New create a new instance of the Monnify struct based on provided config.
//Returns a pointer to the struct and nil error if successful or a nil pointer and an error
func New(config Config) (*Monnify, error) {
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	base := newBase(config)

	m := &Monnify{
		General:          &general{base, nil},
		Disbursements:    &disbursements{base},
		ReservedAccounts: &reservedAccounts{base},
	}
	return m, nil
}

//validateConfig checks the provided config to ensure it's well formed
func validateConfig(config Config) error {
	if config.Environment != EnvSandbox && config.Environment != EnvLive {
		return errors.New(fmt.Sprintf("malformed config - provided enviroment is not supported. - Only %v or %v is allowed", EnvLive, EnvSandbox))
	}

	if config.APIKey == "" {
		return errors.New("malformed config - APIKey is required")
	}

	if config.SecretKey == "" {
		return errors.New("malformed config - secret key if required")
	}

	if config.Environment == EnvLive && config.APIKey == SandBoxAPIKey {
		return errors.New("malformed config - using sandbox api key in live mode not permitted")
	}

	if config.Environment == EnvLive && config.SecretKey == SandBoxSecretKey {
		return errors.New("malformed config - using sandbox secret key in live mode not permitted")
	}
	return nil
}

func failedRequestMessage(httpCode int, monnifyCode interface{}, message string) error {
	return errors.New(fmt.Sprintf("Request Failed - HTTP Status Code: %v | Monnify Status Code: %v | Message: %v", httpCode, monnifyCode, message))
}
