package account

import (
	models "github.com/jerrychan807/okex/models/account"
)

type (
	GetBalance struct {
		//responses.Basic
		Code     string            `json:"code"`
		Msg      string            `json:"msg"`
		Balances []*models.Balance `json:"data,omitempty"`
	}
	GetPositions struct {
		//responses.Basic
		Code      string             `json:"code"`
		Msg       string             `json:"msg"`
		Positions []*models.Position `json:"data"`
	}
	GetAccountAndPositionRisk struct {
		//responses.Basic
		Code                    string `json:"code"`
		Msg                     string
		PositionAndAccountRisks []*models.PositionAndAccountRisk `json:"data"`
	}
	GetBills struct {
		//responses.Basic
		Code  string `json:"code"`
		Msg   string
		Bills []*models.Bill `json:"data"`
	}
	GetConfig struct {
		//responses.Basic
		Code    string `json:"code"`
		Msg     string
		Configs []*models.Config `json:"data"`
	}
	SetPositionMode struct {
		//responses.Basic
		Code          string `json:"code"`
		Msg           string
		PositionModes []*models.PositionMode `json:"data"`
	}
	Leverage struct {
		//responses.Basic
		Code      string             `json:"code"`
		Msg       string             `json:"msg"`
		Leverages []*models.Leverage `json:"data"`
	}
	GetMaxBuySellAmount struct {
		//responses.Basic
		Code              string `json:"code"`
		Msg               string
		MaxBuySellAmounts []*models.MaxBuySellAmount `json:"data"`
	}
	GetMaxAvailableTradeAmount struct {
		//responses.Basic
		Code                     string `json:"code"`
		Msg                      string
		MaxAvailableTradeAmounts []*models.MaxAvailableTradeAmount `json:"data"`
	}
	IncreaseDecreaseMargin struct {
		//responses.Basic
		Code                 string `json:"code"`
		Msg                  string
		MarginBalanceAmounts []*models.MarginBalanceAmount `json:"data"`
	}
	GetMaxLoan struct {
		//responses.Basic
		Code  string `json:"code"`
		Msg   string
		Loans []*models.Loan `json:"data"`
	}
	GetFeeRates struct {
		//responses.Basic
		Code string `json:"code"`
		Msg  string
		Fees []*models.Fee `json:"data"`
	}
	GetInterestAccrued struct {
		//responses.Basic
		Code            string `json:"code"`
		Msg             string
		InterestAccrues []*models.InterestAccrued `json:"data"`
	}
	GetInterestRates struct {
		//responses.Basic
		Code      string `json:"code"`
		Msg       string
		Interests []*models.InterestRate `json:"data"`
	}
	SetGreeks struct {
		//responses.Basic
		Code   string `json:"code"`
		Msg    string
		Greeks []*models.Greek `json:"data"`
	}
	GetMaxWithdrawals struct {
		//responses.Basic
		Code           string `json:"code"`
		Msg            string
		MaxWithdrawals []*models.MaxWithdrawal `json:"data"`
	}
)
