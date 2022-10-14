package trade

import (
	"github.com/jerrychan807/okex/models/trade"
)

type (
	PlaceOrder struct {
		//responses.Basic
		Code        string              `json:"code"`
		Msg         string              `json:"msg"`
		PlaceOrders []*trade.PlaceOrder `json:"data"`
	}
	CancelOrder struct {
		//responses.Basic
		Code         string               `json:"code"`
		Msg          string               `json:"msg"`
		CancelOrders []*trade.CancelOrder `json:"data"`
	}
	AmendOrder struct {
		//responses.Basic
		Code        string              `json:"code"`
		Msg         string              `json:"msg"`
		AmendOrders []*trade.AmendOrder `json:"data"`
	}
	ClosePosition struct {
		//responses.Basic
		Code           string                 `json:"code"`
		Msg            string                 `json:"msg"`
		ClosePositions []*trade.ClosePosition `json:"data"`
	}
	OrderList struct {
		//responses.Basic
		Code   string         `json:"code"`
		Msg    string         `json:"msg"`
		Orders []*trade.Order `json:"data"`
	}
	TransactionDetail struct {
		//responses.Basic
		Code               string                     `json:"code"`
		Msg                string                     `json:"msg"`
		TransactionDetails []*trade.TransactionDetail `json:"data"`
	}
	PlaceAlgoOrder struct {
		//responses.Basic
		Code            string                  `json:"code"`
		Msg             string                  `json:"msg"`
		PlaceAlgoOrders []*trade.PlaceAlgoOrder `json:"data"`
	}
	CancelAlgoOrder struct {
		//responses.Basic
		Code             string                   `json:"code"`
		Msg              string                   `json:"msg"`
		CancelAlgoOrders []*trade.CancelAlgoOrder `json:"data"`
	}
	AlgoOrderList struct {
		//responses.Basic
		Code       string             `json:"code"`
		Msg        string             `json:"msg"`
		AlgoOrders []*trade.AlgoOrder `json:"data"`
	}
)
