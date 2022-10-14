package grid

import "github.com/amir-the-h/okex"

type (
	GridHistoryOrderList struct {
		AlgoOrdType string              `json:"algoOrdType"` // 策略订单类型
		AlgoID      string              `json:"algoId,omitempty"`
		InstID      string              `json:"instId,omitempty"`   // 产品ID
		InstType    okex.InstrumentType `json:"instType,omitempty"` // 产品类型
		After       float64             `json:"after,omitempty,string"`
		Before      float64             `json:"before,omitempty,string"`
		Limit       float64             `json:"limit,omitempty,string"`
	}
)
