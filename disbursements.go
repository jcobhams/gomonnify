package gomonnify

import (
	"fmt"
	"github.com/jcobhams/gomonnify/params"
	"net/http"
	"strings"
)

// SingleTransfer sends money to a single recipient.
// Docs: https://docs.teamapt.com/display/MON/Initiate+Transfer
func (d *disbursements) SingleTransfer(params params.SingleTransferParam) (*SingleTransferResponse, error) {
	url := fmt.Sprintf("%v/v1/disbursements/single", d.APIBaseUrl)
	rawResponse, statusCode, err := d.postRequest(url, requestAuthTypeBasic, params)
	if err != nil {
		return nil, err
	}

	result := SingleTransferResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

// BulkTransfer sends money to a list of recipients.
// Docs: https://docs.teamapt.com/display/MON/Initiate+Transfer
func (d *disbursements) BulkTransfer(params params.BulkTransferParam) (*BulkTransferResponse, error) {
	url := fmt.Sprintf("%v/v1/disbursements/batch", d.APIBaseUrl)
	rawResponse, statusCode, err := d.postRequest(url, requestAuthTypeBasic, params)
	if err != nil {
		return nil, err
	}

	result := BulkTransferResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

// AuthorizeSingleTransfer validates the OTP for the transaction
// Docs: https://docs.teamapt.com/pages/viewpage.action?pageId=4587995
func (d *disbursements) AuthorizeSingleTransfer(reference, authorizationCode string) (*SingleTransferResponse, error) {
	url := fmt.Sprintf("%v/v1/disbursements/single/validate-otp", d.APIBaseUrl)
	param := struct {
		Reference         string `json:"reference"`
		AuthorizationCode string `json:"authorizationCode"`
	}{
		Reference:         reference,
		AuthorizationCode: authorizationCode,
	}
	rawResponse, statusCode, err := d.postRequest(url, requestAuthTypeBasic, param)
	if err != nil {
		return nil, err
	}

	result := SingleTransferResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

// AuthorizeBulkTransfer validates the OTP for the transaction
// Docs: https://docs.teamapt.com/pages/viewpage.action?pageId=4587995
func (d *disbursements) AuthorizeBulkTransfer(reference, authorizationCode string) (*BulkTransferResponse, error) {
	url := fmt.Sprintf("%v/v1/disbursements/batch/validate-otp", d.APIBaseUrl)
	param := struct {
		Reference         string `json:"reference"`
		AuthorizationCode string `json:"authorizationCode"`
	}{
		Reference:         reference,
		AuthorizationCode: authorizationCode,
	}
	rawResponse, statusCode, err := d.postRequest(url, requestAuthTypeBasic, param)
	if err != nil {
		return nil, err
	}
	result := BulkTransferResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

// SingleTransferDetails gets a single transfer detail
// Docs: https://docs.teamapt.com/display/MON/Get+Transfer+Details
func (d *disbursements) SingleTransferDetails(reference string) (*SingleTransferDetailsResponse, error) {
	url := fmt.Sprintf("%v/v1/disbursements/single/summary?reference=%v", d.APIBaseUrl, reference)

	rawResponse, statusCode, err := d.getRequest(url, requestAuthTypeBasic)
	if err != nil {
		return nil, err
	}

	result := SingleTransferDetailsResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

// BulkTransferDetails gets a bulk transfer detail
// Docs: https://docs.teamapt.com/display/MON/Get+Transfer+Details
func (d *disbursements) BulkTransferDetails(batchReference string) (*BulkTransferDetailsResponse, error) {
	url := fmt.Sprintf("%v/v1/disbursements/batch/summary?reference=%v", d.APIBaseUrl, batchReference)
	rawResponse, statusCode, err := d.getRequest(url, requestAuthTypeBasic)
	if err != nil {
		return nil, err
	}

	result := BulkTransferDetailsResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

// BulkTransferTransactions returns a list of transactions in a bulk transfer batch
// Docs: https://docs.teamapt.com/display/MON/Get+Bulk+Transfer+Transactions
func (d *disbursements) BulkTransferTransactions(batchReference string, pageNo, pageSize int) (*TransferTransactionsResponse, error) {
	url := fmt.Sprintf("%v/v1/disbursements/bulk/%v/transactions?pageNo=%v&pageSize=%v", d.APIBaseUrl, batchReference, pageNo, pageSize)
	rawResponse, statusCode, err := d.getRequest(url, requestAuthTypeBasic)
	if err != nil {
		return nil, err
	}

	result := TransferTransactionsResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

func (d *disbursements) SingleTransferTransactions(pageNo, pageSize int) (*TransferTransactionsResponse, error) {
	url := fmt.Sprintf("%v/v1/disbursements/single/transactions?pageNo=%v&pageSize=%v", d.APIBaseUrl, pageNo, pageSize)
	rawResponse, statusCode, err := d.getRequest(url, requestAuthTypeBasic)
	if err != nil {
		return nil, err
	}

	result := TransferTransactionsResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

// ValidateAccountNumber This allows you check if an account number is a valid NUBAN, get the account name if valid.
// Docs: https://docs.teamapt.com/display/MON/Validate+Bank+Account
func (d *disbursements) ValidateAccountNumber(accountNumber, bankCode string) (*ValidAccountNumberResponse, error) {
	url := fmt.Sprintf("%v/v1/disbursements/account/validate?accountNumber=%v&bankCode=%v", d.APIBaseUrl, accountNumber, bankCode)
	rawResponse, statusCode, err := d.getRequest(url, requestAuthTypeBasic)
	if err != nil {
		return nil, err
	}

	result := ValidAccountNumberResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

// WalletBalance returns the available balance in the monnify wallet
// Docs: https://docs.teamapt.com/display/MON/Get+Wallet+Balance
func (d *disbursements) WalletBalance(walletId string) (*WalletBalanceResponse, error) {
	url := fmt.Sprintf("%v/v1/disbursements/wallet-balance?walletId=%v", d.APIBaseUrl, walletId)
	rawResponse, statusCode, err := d.getRequest(url, requestAuthTypeBasic)
	if err != nil {
		return nil, err
	}

	result := WalletBalanceResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

func (d *disbursements) ResendOTP(reference string) (*ResendOTPResponse, error) {
	param := struct {
		Reference string `json:"reference"`
	}{
		Reference: reference,
	}

	url := fmt.Sprintf("%v/v1/disbursements/single/resend-otp", d.APIBaseUrl)
	rawResponse, statusCode, err := d.postRequest(url, requestAuthTypeBasic, param)
	if err != nil {
		return nil, err
	}

	result := ResendOTPResponse{}
	err = d.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}
