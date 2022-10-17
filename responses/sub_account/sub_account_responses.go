package sub_account

import (
	"github.com/jerrychan807/okex/models/account"
	models "github.com/jerrychan807/okex/models/subaccount"
)

type (
	ViewList struct {
		//responses.Basic
		Code        string `json:"code"`
		Msg         string
		SubAccounts []*models.SubAccount `json:"data,omitempty"`
	}
	APIKey struct {
		//responses.Basic
		Code    string `json:"code"`
		Msg     string
		APIKeys []*models.APIKey `json:"data,omitempty"`
	}
	GetBalance struct {
		//responses.Basic
		Code     string `json:"code"`
		Msg      string
		Balances []*account.Balance `json:"data,omitempty"`
	}
	HistoryTransfer struct {
		//responses.Basic
		Code             string `json:"code"`
		Msg              string
		HistoryTransfers []*models.HistoryTransfer `json:"data,omitempty"`
	}
	ManageTransfer struct {
		//responses.Basic
		Code      string `json:"code"`
		Msg       string
		Transfers []*models.Transfer `json:"data,omitempty"`
	}
)
