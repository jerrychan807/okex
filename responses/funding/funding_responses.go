package funding

import (
	models "github.com/jerrychan807/okex/models/funding"
)

type (
	GetCurrencies struct {
		//responses.Basic
		Code       string `json:"code"`
		Msg        string
		Currencies []*models.Currency `json:"data"`
	}
	GetBalance struct {
		//responses.Basic
		Code     string `json:"code"`
		Msg      string
		Balances []*models.Balance `json:"data"`
	}
	FundsTransfer struct {
		//responses.Basic
		Code      string `json:"code"`
		Msg       string
		Transfers []*models.Transfer `json:"data"`
	}
	AssetBillsDetails struct {
		//responses.Basic
		Code  string `json:"code"`
		Msg   string
		Bills []*models.Bill `json:"data"`
	}
	GetDepositAddress struct {
		//responses.Basic
		Code             string `json:"code"`
		Msg              string
		DepositAddresses []*models.DepositAddress `json:"data"`
	}
	GetDepositHistory struct {
		//responses.Basic
		Code             string `json:"code"`
		Msg              string
		DepositHistories []*models.DepositHistory `json:"data"`
	}
	Withdrawal struct {
		//responses.Basic
		Code        string `json:"code"`
		Msg         string
		Withdrawals []*models.Withdrawal `json:"data"`
	}
	GetWithdrawalHistory struct {
		//responses.Basic
		Code                string `json:"code"`
		Msg                 string
		WithdrawalHistories []*models.WithdrawalHistory `json:"data"`
	}
	PiggyBankPurchaseRedemption struct {
		//responses.Basic
		Code       string `json:"code"`
		Msg        string
		PiggyBanks []*models.PiggyBank `json:"data"`
	}
	GetPiggyBankBalance struct {
		//responses.Basic
		Code     string `json:"code"`
		Msg      string
		Balances []*models.PiggyBankBalance `json:"data"`
	}
)
