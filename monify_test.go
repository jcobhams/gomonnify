package gomonnify

import (
	"fmt"
	"github.com/jcobhams/gomonnify/params"
	"github.com/jcobhams/gomonnify/testhelpers"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var client *Monnify

func TestMain(m *testing.M) {
	mockAPIServer := testhelpers.MockAPIServer()

	os.Setenv("GOMONNIFY_TESTMODE", "ON")
	os.Setenv("GOMONNIFY_TESTURL", mockAPIServer.URL)
	client, _ = New(DefaultConfig)

	os.Exit(m.Run())
}

//Reserve Account Tests
func TestReservedAccounts_ReserveAccount(t *testing.T) {
	opts := params.ReserveAccountParam{
		AccountReference:      testhelpers.AccountReference,
		AccountName:           testhelpers.AccountName,
		CurrencyCode:          CurrencyNGN,
		ContractCode:          testhelpers.ContractCode,
		CustomerEmail:         testhelpers.CustomerEmail,
		CustomerName:          testhelpers.CustomerName,
		RestrictPaymentSource: false,
	}
	r, err := client.ReservedAccounts.ReserveAccount(opts)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.AccountName, r.ResponseBody.AccountName)
	assert.Equal(t, testhelpers.AccountNumber, r.ResponseBody.AccountNumber)
	assert.Equal(t, testhelpers.AccountNumber, r.ResponseBody.AccountNumber)
	assert.Equal(t, testhelpers.ContractCode, r.ResponseBody.ContractCode)
}

func TestReservedAccounts_Details(t *testing.T) {
	r, err := client.ReservedAccounts.Details(testhelpers.AccountReference)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.AccountName, r.ResponseBody.AccountName)
	assert.Equal(t, testhelpers.AccountNumber, r.ResponseBody.AccountNumber)
	assert.Equal(t, testhelpers.AccountNumber, r.ResponseBody.AccountNumber)
	assert.Equal(t, testhelpers.ContractCode, r.ResponseBody.ContractCode)
}

func TestReservedAccounts_Deallocate(t *testing.T) {
	err := client.ReservedAccounts.Deallocate(testhelpers.AccountNumber)
	assert.Nil(t, err)
}

func TestReservedAccounts_Transactions(t *testing.T) {
	tx, err := client.ReservedAccounts.Transactions(testhelpers.AccountReference, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tx.ResponseBody.Content))
	assert.Equal(t, testhelpers.CustomerEmail, tx.ResponseBody.Content[0].CustomerDTO.Email)
	assert.Equal(t, testhelpers.CustomerName, tx.ResponseBody.Content[0].CustomerDTO.Name)
	assert.Equal(t, testhelpers.Amount, tx.ResponseBody.Content[0].Amount)
}

//Disbursement Tests
func TestDisbursements_SingleTransfer(t *testing.T) {
	opts := params.SingleTransferParam{
		Amount:        testhelpers.Amount,
		Reference:     testhelpers.TransferReference,
		Narration:     "TEST",
		BankCode:      testhelpers.BankCode,
		AccountNumber: testhelpers.AccountNumber,
		Currency:      CurrencyNGN,
		WalletId:      testhelpers.WalletId,
	}
	s, err := client.Disbursements.SingleTransfer(opts)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.Amount, s.ResponseBody.Amount)
}

func TestDisbursements_BulkTransfer(t *testing.T) {
	opts := params.BulkTransferParam{
		Title:                "TEST BATCH",
		BatchReference:       testhelpers.BatchReference,
		Narration:            "TEST",
		WalletId:             testhelpers.WalletId,
		OnValidationFailure:  ValidationFailedContinue,
		NotificationInterval: NotificationInterval20,
		TransactionList: []params.SingleTransferParam{{
			Amount:        testhelpers.Amount,
			Reference:     testhelpers.TransferReference,
			Narration:     "TEST",
			BankCode:      testhelpers.BankCode,
			AccountNumber: testhelpers.AccountNumber,
			Currency:      CurrencyNGN,
		}},
	}
	b, err := client.Disbursements.BulkTransfer(opts)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.Amount, b.ResponseBody.TotalAmount)
	assert.Equal(t, testhelpers.BatchReference, b.ResponseBody.BatchReference)
}

func TestDisbursements_AuthorizeSingleTransfer(t *testing.T) {
	s, err := client.Disbursements.AuthorizeSingleTransfer(testhelpers.TransferReference, testhelpers.ValidOTP)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.Amount, s.ResponseBody.Amount)
}

func TestDisbursements_AuthorizeBulkTransfer(t *testing.T) {
	s, err := client.Disbursements.AuthorizeBulkTransfer(testhelpers.TransferReference, testhelpers.ValidOTP)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.Amount, s.ResponseBody.TotalAmount)
}

func TestDisbursements_SingleTransferDetails(t *testing.T) {
	d, err := client.Disbursements.SingleTransferDetails(testhelpers.TransferReference)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.Amount, d.ResponseBody.Amount)
}

func TestDisbursements_BulkTransferDetails(t *testing.T) {
	d, err := client.Disbursements.BulkTransferDetails(testhelpers.BatchReference)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.Amount, d.ResponseBody.TotalAmount)
}

func TestDisbursements_BulkTransferTransactions(t *testing.T) {
	tx, err := client.Disbursements.BulkTransferTransactions(testhelpers.BatchReference, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.BankCode, tx.ResponseBody.Content[0].BankCode)
	assert.Equal(t, testhelpers.AccountNumber, tx.ResponseBody.Content[0].AccountNumber)
	assert.Equal(t, testhelpers.Amount, tx.ResponseBody.Content[0].Amount)
}

func TestDisbursements_SingleTransferTransactions(t *testing.T) {
	tx, err := client.Disbursements.SingleTransferTransactions(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.BankCode, tx.ResponseBody.Content[0].BankCode)
	assert.Equal(t, testhelpers.Amount, tx.ResponseBody.Content[0].Amount)
}

func TestDisbursements_ValidateAccountNumber(t *testing.T) {
	d, err := client.Disbursements.ValidateAccountNumber(testhelpers.AccountNumber, testhelpers.BankCode)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.AccountNumber, d.ResponseBody.AccountNumber)
	assert.Equal(t, testhelpers.AccountName, d.ResponseBody.AccountName)
	assert.Equal(t, testhelpers.BankCode, d.ResponseBody.BankCode)
}

func TestDisbursements_WalletBalance(t *testing.T) {
	w, err := client.Disbursements.WalletBalance(testhelpers.WalletId)
	assert.Nil(t, err)
	assert.Equal(t, testhelpers.AvailableBalance, w.ResponseBody.AvailableBalance)
	assert.Equal(t, testhelpers.LedgerBalance, w.ResponseBody.LedgerBalance)
}

func TestDisbursements_ResendOTP(t *testing.T) {
	r, err := client.Disbursements.ResendOTP(testhelpers.TransferReference)
	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestGeneral_GetTransaction(t *testing.T) {
	tx, err := client.General.GetTransaction(testhelpers.TransferReference)
	assert.Nil(t, err)
	assert.Equal(t, PaymentStatusPaid, tx.ResponseBody.PaymentStatus)
	assert.Equal(t, fmt.Sprintf("%v", testhelpers.Amount), tx.ResponseBody.AmountPaid)
}

func TestGeneral_VerifyTransaction(t *testing.T) {
	//Set Secret Key To Be Test Environemnt Secret Key
	client.General.Config.SecretKey = testhelpers.SecretKey
	tx, _ := client.General.GetTransaction(testhelpers.TransferReference)
	assert.True(t, client.General.VerifyTransaction(&tx.ResponseBody, false))
	assert.True(t, client.General.VerifyTransaction(&tx.ResponseBody, true))
}

func TestGeneral_GetBanks(t *testing.T) {
	b, err := client.General.GetBanks()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(b.ResponseBody))
	assert.NotEmpty(t, b.ResponseBody[0].Name)

	bc, err := client.General.GetBanksUseCache()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(bc.ResponseBody))
	assert.NotEmpty(t, bc.ResponseBody[0].Name)

	client.General.InvalidateBankCache()
	bc, err = client.General.GetBanksUseCache()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(bc.ResponseBody))
}
