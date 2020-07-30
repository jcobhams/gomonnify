# GoMonnify - TeamAPT's Monnify API Wrapper
[![Build Status](https://travis-ci.org/jcobhams/gomonnify.svg?branch=master)](https://travis-ci.org/jcobhams/gomonnify)
[![codecov](https://codecov.io/gh/jcobhams/gomonnify/branch/master/graph/badge.svg)](https://codecov.io/gh/jcobhams/gomonnify)
[![GoDoc](https://godoc.org/github.com/jcobhams/gomonnify?status.svg)](https://godoc.org/github.com/jcobhams/gomonnify)


### Installation
`$ go get github.com/jcobhams/gomonnify`

### Usage
```go
package main

import (
	"github.com/jcobhams/gomonnify"
    "fmt"
)

func main() {
    monnifyCfg := Config{
                 		Environment:         gomonnify.EnvLive,
                 		APIKey:              "Your_KEY",
                 		SecretKey:           "Your_Secret",
                 		RequestTimeout:      5 * time.Second,
                 		DefaultContractCode: "Your_Contract_Code",
                 	}
    
    monnify, err := gomonnify.New(monnifyCfg)
    if err != nil {
        //Handle error
    }
    
    res, err := monnify.Disbursements.WalletBalance("your_wallet_id")
    if err != nil {
        //handle error
    }
    
    fmt.Println(res)
}
```

### Modules
1. Disbursements (All EndPoints) - https://docs.teamapt.com/display/MON/Monnify+Disbursements

2. ReservedAccounts (Except `UpdateIncomeSplitConfig()` and `UpdatePaymentSourceFilter()` )

3. Invoice - Coming soon or open a PR :)

4. General - Coming soon or open a PR :)

### Test Helpers
GoMonnify ships with nifty test helpers to ease unit and integration testing your code that import or relies on gomonnify.
Set the following environment variables: 

`GOMONNIFY_TESTMODE` this tells the module to assume it's running in a test context.

`GOMONNIFY_TESTURL` a full integration server is bundled with the module, you can use it or write one if you need. 

Example Test File `my_controller_test.go`
```go
package gomonnify

import (
	"github.com/jcobhams/gomonnify/testhelpers"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	mockAPIServer := testhelpers.MockAPIServer()

	os.Setenv("GOMONNIFY_TESTMODE", "ON")
	os.Setenv("GOMONNIFY_TESTURL", mockAPIServer.URL)
	
	os.Exit(m.Run())
}

func TestMyContollerMethodThatUsesGoMonnify(t *testing.T) {
    ...
    // test body
    ...

}



```


### Run Tests
`$ go test -race -v -coverprofile cover.out`

### View Coverage
`$ go tool cover -html=cover.out`