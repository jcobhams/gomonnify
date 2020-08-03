package gomonnify

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"strings"
)

// VerifyTransaction validates that the payload received is actually from monnify. It computes the transaction hash and compares.
// twoStep if set to true will make a request to monnify to confirm that the transaction exists and was successful
// we cannot implement 3 step verification without some knowledge of your persistence layer. Need it? Open an Issue/PR :)
// Docs: https://docs.teamapt.com/pages/viewpage.action?pageId=13828139
func (g *general) VerifyTransaction(payload *GeneralTransaction, twoStep bool) bool {
	rawStr := fmt.Sprintf("%v|%v|%v|%v|%v", g.Config.SecretKey, payload.PaymentReference, payload.AmountPaid, payload.PaidOn, payload.TransactionReference)
	h := sha512.New()
	h.Write([]byte(rawStr))
	hashed := fmt.Sprintf("%x", h.Sum(nil))
	if hashed != payload.TransactionHash {
		return false
	}

	if twoStep {
		t, err := g.GetTransaction(payload.TransactionReference)
		if err != nil {
			return false
		}

		if t.ResponseBody.PaymentStatus == PaymentStatusPaid {
			return true
		}
		return false
	}

	return true
}

// GetTransaction retrieves the transaction specified by reference from the Monnify API.
// Docs: https://docs.teamapt.com/display/MON/Get+Transaction+Status
func (g *general) GetTransaction(reference string) (*GeneralTransactionResponse, error) {
	url := fmt.Sprintf("%v/v2/transactions/%v", g.APIBaseUrl, reference)
	rawResponse, statusCode, err := g.getRequest(url, requestAuthTypeBearer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := GeneralTransactionResponse{}
	err = g.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

// GetBanks fetches a list of banks and their USSD codes from the monnify api.
// Docs: https://docs.teamapt.com/display/MON/Get+Banks
func (g *general) GetBanks() (*BanksResponse, error) {
	url := fmt.Sprintf("%v/v1/banks", g.APIBaseUrl)
	rawResponse, statusCode, err := g.getRequest(url, requestAuthTypeBearer)
	if err != nil {
		return nil, err
	}

	result := BanksResponse{}
	err = g.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

// GetBanksUseCache checks if the bank list is already in the struct
// and simply reuses that instead of making a HTTP call. If the list is not already in the struct, it fetches and saves
// a copy to the struct for future use.
func (g *general) GetBanksUseCache() (*BanksResponse, error) {
	if g.banks != nil {
		return g.banks, nil
	}
	b, err := g.GetBanks()
	if err != nil {
		return nil, err
	}
	g.banks = b
	return b, nil
}

// InvalidateBankCache empties the cache from the struct so the next call will make an API Call
func (g *general) InvalidateBankCache() {
	g.banks = nil
}
