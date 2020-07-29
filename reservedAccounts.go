package gomonnify

import (
	"errors"
	"fmt"
	"github.com/jcobhams/gomonnify/params"
	"net/http"
	"strings"
)

//ReserveAccount reserves an account number for a customer based on the provided configuration params.
//Docs: https://docs.teamapt.com/display/MON/Reserving+An+Account
func (r *reservedAccounts) ReserveAccount(params params.ReserveAccountParam) (*ReserveAccountResponse, error) {
	url := fmt.Sprintf("%v/bank-transfer/reserved-accounts", r.APIBaseUrl)
	rawResponse, statusCode, err := r.postRequest(url, requestAuthTypeBearer, params)
	if err != nil {
		return nil, err
	}

	result := ReserveAccountResponse{}
	err = r.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

//Details gets the reserved account information for the provided account reference.
//Docs: https://docs.teamapt.com/display/MON/Get+Reserved+Account+Details
func (r *reservedAccounts) Details(accountReference string) (*ReserveAccountResponse, error) {
	if accountReference == "" {
		return nil, errors.New("accountReference is required")
	}

	url := fmt.Sprintf("%v/bank-transfer/reserved-accounts/%v", r.APIBaseUrl, accountReference)
	rawResponse, statusCode, err := r.getRequest(url, requestAuthTypeBearer)
	if err != nil {
		return nil, err
	}

	var result ReserveAccountResponse
	err = r.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

//Deallocate deletes a reserved account.
//Docs: https://docs.teamapt.com/display/MON/Deallocating+a+reserved+account
func (r *reservedAccounts) Deallocate(accountNumber string) error {
	if accountNumber == "" {
		return errors.New("accountNumber is required")
	}

	url := fmt.Sprintf("%v/bank-transfer/reserved-accounts/%v", r.APIBaseUrl, accountNumber)
	rawResponse, statusCode, err := r.deleteRequest(url, requestAuthTypeBearer)
	if err != nil {
		return err
	}

	var result ReserveAccountResponse
	err = r.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return err
	}

	if statusCode != http.StatusOK {
		return failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return nil
}

//Transactions fetches all the transaction on a reserved account for the provided account reference.
//Docs: https://docs.teamapt.com/display/MON/Getting+all+transactions+on+a+reserved+account
func (r *reservedAccounts) Transactions(accountReference string, page, size int) (*TransactionsResponse, error) {
	if accountReference == "" {
		return nil, errors.New("accountNumber is required")
	}

	url := fmt.Sprintf("%v/bank-transfer/reserved-accounts/transactions?accountReference=%v&page=%v&size=%v", r.APIBaseUrl, accountReference, page, size)
	rawResponse, statusCode, err := r.getRequest(url, requestAuthTypeBearer)
	if err != nil {
		return nil, err
	}

	var result TransactionsResponse
	err = r.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}

	return &result, nil
}

//TODO
func (r *reservedAccounts) updateIncomeSplitConfig(accountReference string) {
}

func (r *reservedAccounts) updatePaymentSourceFilter(accountReference string) {
}
