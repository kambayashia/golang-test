package twitter_ads

import (
	"fmt"
	"github.com/dghubble/sling"
	"errors"
)

type AccountService struct {
	sling *sling.Sling
}

func NewAccountService(sling *sling.Sling) *AccountService {
	return &AccountService{
		sling: sling,
	}
}

type Account struct {
}

type AccountsParams struct {

}

func (o *AccountService) Accounts(params *AccountsParams) ([]Account, error) {
	response := new(Response)
	apiError := new(ApiError)
	if resp, _ := ApiVersion(o.sling, 1).Get("accounts").QueryStruct(params).Receive(response, apiError); resp.StatusCode >= 400 {
		return nil, errors.New(fmt.Sprintf("error:%v", apiError.Errors[0]))
	}

	fmt.Printf("errors: %v", apiError.Errors)
	return make([]Account, response.TotalCount), nil
}