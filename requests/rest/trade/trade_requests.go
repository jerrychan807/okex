package trade

import (
	"github.com/amir-the-h/okex"
)

type (
	PlaceOrder struct {
		ID         string            `json:"-"`
		InstID     string            `json:"instId"`               // 产品ID，如 BTC-USD-190927-5000-C
		Ccy        string            `json:"ccy,omitempty"`        // 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		ClOrdID    string            `json:"clOrdId,omitempty"`    // 客户自定义订单ID,字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间
		Tag        string            `json:"tag,omitempty"`        // 订单标签,字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		ReduceOnly bool              `json:"reduceOnly,omitempty"` // 是否只减仓，true 或 false，默认false
		Sz         float64           `json:"sz,string"`            // 委托数量
		Px         float64           `json:"px,omitempty,string"`  // 委托价格，仅适用于limit、post_only、fok、ioc类型的订单
		TdMode     okex.TradeMode    `json:"tdMode"`               // 交易模式,保证金模式：isolated：逐仓 ；cross：全仓,非保证金模式：cash：非保证金
		Side       okex.OrderSide    `json:"side"`                 // 订单方向,buy：买， sell：卖
		PosSide    okex.PositionSide `json:"posSide,omitempty"`    // 持仓方向,在双向持仓模式下必填，且仅可选择 long 或 short。 仅适用交割、永续
		OrdType    okex.OrderType    `json:"ordType"`              // 订单类型,market：市价单,limit：限价单
		TgtCcy     okex.QuantityType `json:"tgtCcy,omitempty"`
	}
	CancelOrder struct {
		ID      string `json:"-"`
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	AmendOrder struct {
		ID        string  `json:"-"`
		InstID    string  `json:"instId"`
		OrdID     string  `json:"ordId,omitempty"`
		ClOrdID   string  `json:"clOrdId,omitempty"`
		ReqID     string  `json:"reqId,omitempty"`
		NewSz     int64   `json:"newSz,omitempty,string"`
		NewPx     float64 `json:"newPx,omitempty,string"`
		CxlOnFail bool    `json:"cxlOnFail,omitempty"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId"`
		Ccy     string            `json:"ccy,omitempty"`
		PosSide okex.PositionSide `json:"posSide,omitempty"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
	}
	OrderDetails struct {
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	OrderList struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		After    float64             `json:"after,omitempty,string"`
		Before   float64             `json:"before,omitempty,string"`
		Limit    float64             `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
		OrdType  okex.OrderType      `json:"ordType,omitempty"`
		State    okex.OrderState     `json:"state,omitempty"`
	}
	TransactionDetails struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		OrdID    string              `json:"ordId,omitempty"`
		After    float64             `json:"after,omitempty,string"`
		Before   float64             `json:"before,omitempty,string"`
		Limit    float64             `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
	}
	PlaceAlgoOrder struct {
		InstID     string             `json:"instId"`
		TdMode     okex.TradeMode     `json:"tdMode"`
		Ccy        string             `json:"ccy,omitempty"`
		Side       okex.OrderSide     `json:"side"`
		PosSide    okex.PositionSide  `json:"posSide,omitempty"`
		OrdType    okex.AlgoOrderType `json:"ordType"`
		Sz         int64              `json:"sz,string"`
		ReduceOnly bool               `json:"reduceOnly,omitempty"`
		TgtCcy     okex.QuantityType  `json:"tgtCcy,omitempty"`
		StopOrder
		TriggerOrder
		IcebergOrder
		TWAPOrder
	}
	StopOrder struct {
		TpTriggerPx float64 `json:"tpTriggerPx,string,omitempty"`
		TpOrdPx     float64 `json:"tpOrdPx,string,omitempty"`
		SlTriggerPx float64 `json:"slTriggerPx,string,omitempty"`
		SlOrdPx     float64 `json:"slOrdPx,string,omitempty"`
	}
	TriggerOrder struct {
		TriggerPx float64 `json:"triggerPx,string,omitempty"`
		OrdPx     float64 `json:"ordPx,string,omitempty"`
	}
	IcebergOrder struct {
		PxVar    float64 `json:"pxVar,string,omitempty"`
		PxSpread float64 `json:"pxSpread,string,omitempty"`
		SzLimit  int64   `json:"szLimit,string"`
		PxLimit  float64 `json:"pxLimit,string"`
	}
	TWAPOrder struct {
		IcebergOrder
		TimeInterval string `json:"timeInterval"`
	}
	CancelAlgoOrder struct {
		InstID string `json:"instId"`
		AlgoID string `json:"AlgoId"`
	}
	AlgoOrderList struct {
		InstType okex.InstrumentType `json:"instType,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		After    float64             `json:"after,omitempty,string"`
		Before   float64             `json:"before,omitempty,string"`
		Limit    float64             `json:"limit,omitempty,string"`
		OrdType  okex.AlgoOrderType  `json:"ordType,omitempty"`
		State    okex.OrderState     `json:"state,omitempty"`
	}
)
