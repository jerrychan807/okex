package grid

import "github.com/jerrychan807/okex"

type (
	GridOrder struct {
		AlgoID              string              `json:"algoId"`
		InstType            okex.InstrumentType `json:"instType"`
		InstID              string              `json:"instId"`
		CTime               okex.JSONTime       `json:"cTime"`
		UTime               okex.JSONTime       `json:"uTime"`
		AlgoOrdType         string              `json:"algoOrdType"` // 策略订单类型
		State               okex.OrderState     `json:"state"`
		MaxPx               okex.JSONFloat64    `json:"maxPx"`
		MinPx               okex.JSONFloat64    `json:"minPx"`
		GridNum             string              `json:"gridNum"`
		RunType             string              `json:"runType"`
		TpTriggerPx         okex.JSONFloat64    `json:"tpTriggerPx"`
		SlTriggerPx         okex.JSONFloat64    `json:"slTriggerPx"`
		ArbitrageNum        string              `json:"arbitrageNum"`
		TotalPnl            okex.JSONFloat64    `json:"totalPnl"`
		PnlRatio            okex.JSONFloat64    `json:"pnlRatio"`
		Investment          okex.JSONFloat64    `json:"investment"`
		GridProfit          okex.JSONFloat64    `json:"gridProfit"`
		FloatProfit         okex.JSONFloat64    `json:"floatProfit"`
		TotalAnnualizedRate okex.JSONFloat64    `json:"totalAnnualizedRate"`
		AnnualizedRate      okex.JSONFloat64    `json:"annualizedRate"`
		CancelType          string              `json:"cancelType"`
		StopType            string              `json:"stopType"`
		QuoteSz             okex.JSONFloat64    `json:"quoteSz"`
		BaseSz              okex.JSONFloat64    `json:"baseSz"`
		Direction           string              `json:"direction"`
		BasePos             bool                `json:"basePos"`
		Sz                  okex.JSONFloat64     `json:"sz"`
		Lever               okex.JSONFloat64    `json:"lever"`
		ActualLever         okex.JSONFloat64    `json:"actualLever"`
		LiqPx               okex.JSONFloat64    `json:"liqPx"`
		Uly                 string              `json:"uly"`
		Tag                 string              `json:"tag"`
	}
)
