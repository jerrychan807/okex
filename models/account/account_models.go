package account

import (
	"github.com/amir-the-h/okex"
)

type (
	Balance struct {
		TotalEq     okex.JSONFloat64  `json:"totalEq"`               // 美金层面权益
		IsoEq       okex.JSONFloat64  `json:"isoEq"`                 // 美金层面逐仓仓位权益, 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
		AdjEq       okex.JSONFloat64  `json:"adjEq,omitempty"`       // 美金层面有效保证金, 适用于跨币种保证金模式和组合保证金模式
		OrdFroz     okex.JSONFloat64  `json:"ordFroz,omitempty"`     // 美金层面全仓挂单占用保证金, 适用于跨币种保证金模式和组合保证金模式
		Imr         okex.JSONFloat64  `json:"imr,omitempty"`         // 美金层面占用保证金, 适用于跨币种保证金模式和组合保证金模式
		Mmr         okex.JSONFloat64  `json:"mmr,omitempty"`         // 美金层面维持保证金, 适用于跨币种保证金模式和组合保证金模式
		MgnRatio    okex.JSONFloat64  `json:"mgnRatio,omitempty"`    // 美金层面保证金率, 适用于跨币种保证金模式 和组合保证金模式
		NotionalUsd okex.JSONFloat64  `json:"notionalUsd,omitempty"` // 以美金价值为单位的持仓数量, 即仓位美金价值,适用于跨币种保证金模式和组合保证金模式
		Details     []*BalanceDetails `json:"details,omitempty"`     // 各币种资产详细信息
		UTime       okex.JSONTime     `json:"uTime"`                 // 账户信息的更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	}
	BalanceDetails struct {
		Ccy           string           `json:"ccy"`                     // 币种
		Eq            okex.JSONFloat64 `json:"eq"`                      // 币种总权益
		CashBal       okex.JSONFloat64 `json:"cashBal"`                 // 币种余额
		IsoEq         okex.JSONFloat64 `json:"isoEq,omitempty"`         // 币种逐仓仓位权益, 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
		AvailEq       okex.JSONFloat64 `json:"availEq,omitempty"`       // 可用保证金, 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
		DisEq         okex.JSONFloat64 `json:"disEq"`                   // 美金层面币种折算权益
		AvailBal      okex.JSONFloat64 `json:"availBal"`                // 可用余额, 适用于简单交易模式
		FrozenBal     okex.JSONFloat64 `json:"frozenBal"`               // 币种占用金额
		OrdFrozen     okex.JSONFloat64 `json:"ordFrozen"`               // 挂单冻结数量
		Liab          okex.JSONFloat64 `json:"liab,omitempty"`          // 币种负债额, 适用于跨币种保证金模式和组合保证金模式
		Upl           okex.JSONFloat64 `json:"upl,omitempty"`           // 未实现盈亏, 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
		UplLib        okex.JSONFloat64 `json:"uplLiab,omitempty"`       // 由于仓位未实现亏损导致的负债, 适用于跨币种保证金模式和组合保证金模式
		CrossLiab     okex.JSONFloat64 `json:"crossLiab,omitempty"`     // 币种全仓负债额, 适用于跨币种保证金模式和组合保证金模式
		IsoLiab       okex.JSONFloat64 `json:"isoLiab,omitempty"`       // 币种逐仓负债额, 适用于跨币种保证金模式和组合保证金模式
		MgnRatio      okex.JSONFloat64 `json:"mgnRatio,omitempty"`      // 保证金率, 适用于单币种保证金模式
		Interest      okex.JSONFloat64 `json:"interest,omitempty"`      // 计息, 适用于跨币种保证金模式和组合保证金模式
		Twap          okex.JSONFloat64 `json:"twap,omitempty"`          // 当前负债币种触发系统自动换币的风险,0、1、2、3、4、5其中之一，数字越大代表您的负债币种触发自动换币概率越高,适用于跨币种保证金模式和组合保证金模式
		MaxLoan       okex.JSONFloat64 `json:"maxLoan,omitempty"`       // 币种最大可借, 适用于跨币种保证金模式和组合保证金模式
		EqUsd         okex.JSONFloat64 `json:"eqUsd"`                   // 币种权益美金价值
		NotionalLever okex.JSONFloat64 `json:"notionalLever,omitempty"` // 币种杠杆倍数, 适用于单币种保证金模式
		StgyEq        okex.JSONFloat64 `json:"stgyEq"`                  // 策略权益
		IsoUpl        okex.JSONFloat64 `json:"isoUpl,omitempty"`        // 逐仓未实现盈亏, 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
		UTime         okex.JSONTime    `json:"uTime"`                   //
	}
	Position struct {
		InstID      string              `json:"instId"`             //
		PosCcy      string              `json:"posCcy,omitempty"`   // 仓位资产币种，仅适用于币币杠杆仓位
		LiabCcy     string              `json:"liabCcy,omitempty"`  // 负债币种，仅适用于币币杠杆
		OptVal      string              `json:"optVal,omitempty"`   // 期权市值，仅适用于期权
		Ccy         string              `json:"ccy"`                // 占用保证金的币种
		PosID       string              `json:"posId"`              // 持仓ID
		TradeID     string              `json:"tradeId"`            // 最新成交ID
		Pos         okex.JSONFloat64    `json:"pos"`                // 持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为0的仓位
		AvailPos    okex.JSONFloat64    `json:"availPos,omitempty"` // 可平仓数量，适用于 币币杠杆,交割/永续（开平仓模式），期权（交易账户及保证金账户逐仓）。
		AvgPx       okex.JSONFloat64    `json:"avgPx"`              // 开仓平均价
		Upl         okex.JSONFloat64    `json:"upl"`                // 未实现收益
		UplRatio    okex.JSONFloat64    `json:"uplRatio"`           // 未实现收益率
		Lever       okex.JSONFloat64    `json:"lever"`              // 杠杆倍数,不适用于期权
		LiqPx       okex.JSONFloat64    `json:"liqPx,omitempty"`    // 预估强平价,不适用于期权
		Imr         okex.JSONFloat64    `json:"imr,omitempty"`      // 初始保证金，仅适用于全仓
		Margin      okex.JSONFloat64    `json:"margin,omitempty"`   // 保证金余额，可增减，仅适用于逐仓
		MgnRatio    okex.JSONFloat64    `json:"mgnRatio"`           // 保证金率
		Mmr         okex.JSONFloat64    `json:"mmr"`                // 维持保证金
		Liab        okex.JSONFloat64    `json:"liab,omitempty"`     // 负债额，仅适用于币币杠杆
		Interest    okex.JSONFloat64    `json:"interest"`           // 利息，已经生成的未扣利息
		NotionalUsd okex.JSONFloat64    `json:"notionalUsd"`        // 以美金价值为单位的持仓数量
		ADL         okex.JSONFloat64    `json:"adl"`                // 信号区,分为5档，从1到5，数字越小代表adl强度越弱
		Last        okex.JSONFloat64    `json:"last"`               // 最新成交价
		DeltaBS     okex.JSONFloat64    `json:"deltaBS"`            //
		DeltaPA     okex.JSONFloat64    `json:"deltaPA"`            //
		GammaBS     okex.JSONFloat64    `json:"gammaBS"`            //
		GammaPA     okex.JSONFloat64    `json:"gammaPA"`            //
		ThetaBS     okex.JSONFloat64    `json:"thetaBS"`            //
		ThetaPA     okex.JSONFloat64    `json:"thetaPA"`            //
		VegaBS      okex.JSONFloat64    `json:"vegaBS"`             //
		VegaPA      okex.JSONFloat64    `json:"vegaPA"`             //
		PosSide     okex.PositionSide   `json:"posSide"`            // 持仓方向,long：双向持仓多头,short：双向持仓空头,net：单向持仓（交割/永续/期权：pos为正代表多头，pos为负代表空头。币币杠杆：posCcy为交易货币时，代表多头；posCcy为计价货币时，代表空头。）
		MgnMode     okex.MarginMode     `json:"mgnMode"`            // 保证金模式 cross：全仓 isolated：逐仓
		InstType    okex.InstrumentType `json:"instType"`           // 产品类型
		CTime       okex.JSONTime       `json:"cTime"`              // 持仓创建时间，Unix时间戳的毫秒数格式，如 1597026383085
		UTime       okex.JSONTime       `json:"uTime"`              // 最近一次持仓更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	}
	BalanceAndPosition struct {
		EventType okex.EventType    `json:"eventType"`
		PTime     okex.JSONTime     `json:"pTime"`
		UTime     okex.JSONTime     `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   okex.JSONFloat64                     `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      okex.JSONTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string           `json:"ccy"`
		Eq    okex.JSONFloat64 `json:"eq"`
		DisEq okex.JSONFloat64 `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		Ccy         string              `json:"ccy"`
		NotionalCcy okex.JSONFloat64    `json:"notionalCcy"`
		Pos         okex.JSONFloat64    `json:"pos"`
		NotionalUsd okex.JSONFloat64    `json:"notionalUsd"`
		PosSide     okex.PositionSide   `json:"posSide"`
		InstType    okex.InstrumentType `json:"instType"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string              `json:"ccy"`
		InstID    string              `json:"instId"`
		Notes     string              `json:"notes"`
		BillID    string              `json:"billId"`
		OrdID     string              `json:"ordId"`
		BalChg    okex.JSONFloat64    `json:"balChg"`
		PosBalChg okex.JSONFloat64    `json:"posBalChg"`
		Bal       okex.JSONFloat64    `json:"bal"`
		PosBal    okex.JSONFloat64    `json:"posBal"`
		Sz        okex.JSONFloat64    `json:"sz"`
		Pnl       okex.JSONFloat64    `json:"pnl"`
		Fee       okex.JSONFloat64    `json:"fee"`
		From      okex.AccountType    `json:"from,string"`
		To        okex.AccountType    `json:"to,string"`
		InstType  okex.InstrumentType `json:"instType"`
		MgnMode   okex.MarginMode     `json:"MgnMode"`
		Type      okex.BillType       `json:"type,string"`
		SubType   okex.BillSubType    `json:"subType,string"`
		TS        okex.JSONTime       `json:"ts"`
	}
	Config struct {
		Level      string            `json:"level"`
		LevelTmp   string            `json:"levelTmp"`
		AcctLv     string            `json:"acctLv"`
		AutoLoan   bool              `json:"autoLoan"`
		UID        string            `json:"uid"`
		GreeksType okex.GreekType    `json:"greeksType"`
		PosMode    okex.PositionType `json:"posMode"`
	}
	PositionMode struct {
		PosMode okex.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstID  string            `json:"instId"`
		Lever   okex.JSONFloat64  `json:"lever"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string           `json:"instId"`
		Ccy     string           `json:"ccy"`
		MaxBuy  okex.JSONFloat64 `json:"maxBuy"`
		MaxSell okex.JSONFloat64 `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string           `json:"instId"`
		AvailBuy  okex.JSONFloat64 `json:"availBuy"`
		AvailSell okex.JSONFloat64 `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string            `json:"instId"`
		Amt     okex.JSONFloat64  `json:"amt"`
		PosSide okex.PositionSide `json:"posSide,string"`
		Type    okex.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string           `json:"instId"`
		MgnCcy  string           `json:"mgnCcy"`
		Ccy     string           `json:"ccy"`
		MaxLoan okex.JSONFloat64 `json:"maxLoan"`
		MgnMode okex.MarginMode  `json:"mgnMode"`
		Side    okex.OrderSide   `json:"side,string"`
	}
	Fee struct {
		Level    string              `json:"level"`
		Taker    okex.JSONFloat64    `json:"taker"`
		Maker    okex.JSONFloat64    `json:"maker"`
		Delivery okex.JSONFloat64    `json:"delivery,omitempty"`
		Exercise okex.JSONFloat64    `json:"exercise,omitempty"`
		Category okex.FeeCategory    `json:"category,string"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string           `json:"instId"`
		Ccy          string           `json:"ccy"`
		Interest     okex.JSONFloat64 `json:"interest"`
		InterestRate okex.JSONFloat64 `json:"interestRate"`
		Liab         okex.JSONFloat64 `json:"liab"`
		MgnMode      okex.MarginMode  `json:"mgnMode"`
		TS           okex.JSONTime    `json:"ts"`
	}
	InterestRate struct {
		Ccy          string           `json:"ccy"`
		InterestRate okex.JSONFloat64 `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string           `json:"ccy"`
		MaxWd okex.JSONFloat64 `json:"maxWd"`
	}
)
