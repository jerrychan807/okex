package public_data

import (
	"github.com/jerrychan807/okex/models/publicdata"
)

type (
	GetInstruments struct {
		// responses.Basic
		Code        string                   `json:"code"`
		Msg         string                   `json:"msg"`
		Instruments []*publicdata.Instrument `json:"data,omitempty"`
	}
	GetDeliveryExerciseHistory struct {
		//responses.Basic
		Code      string `json:"code"`
		Msg       string
		Histories []*publicdata.DeliveryExerciseHistory `json:"data,omitempty"`
	}
	GetOpenInterest struct {
		//responses.Basic
		Code          string `json:"code"`
		Msg           string
		OpenInterests []*publicdata.OpenInterest `json:"data,omitempty"`
	}
	GetFundingRate struct {
		//responses.Basic
		Code         string `json:"code"`
		Msg          string
		FundingRates []*publicdata.FundingRate `json:"data,omitempty"`
	}
	GetLimitPrice struct {
		//responses.Basic
		Code        string `json:"code"`
		Msg         string
		LimitPrices []*publicdata.LimitPrice `json:"data,omitempty"`
	}
	GetOptionMarketData struct {
		//responses.Basic
		Code             string `json:"code"`
		Msg              string
		OptionMarketData []*publicdata.OptionMarketData `json:"data,omitempty"`
	}
	GetEstimatedDeliveryExercisePrice struct {
		//responses.Basic
		Code                            string `json:"code"`
		Msg                             string
		EstimatedDeliveryExercisePrices []*publicdata.EstimatedDeliveryExercisePrice `json:"data,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		//responses.Basic
		Code                                 string `json:"code"`
		Msg                                  string
		GetDiscountRateAndInterestFreeQuotas []*publicdata.GetDiscountRateAndInterestFreeQuota `json:"data,omitempty"`
	}
	GetSystemTime struct {
		//responses.Basic
		Code        string `json:"code"`
		Msg         string
		SystemTimes []*publicdata.SystemTime `json:"data,omitempty"`
	}
	GetLiquidationOrders struct {
		//responses.Basic
		Code              string `json:"code"`
		Msg               string
		LiquidationOrders []*publicdata.LiquidationOrder `json:"data,omitempty"`
	}
	GetMarkPrice struct {
		//responses.Basic
		Code       string `json:"code"`
		Msg        string
		MarkPrices []*publicdata.MarkPrice `json:"data,omitempty"`
	}
	GetPositionTiers struct {
		//responses.Basic
		Code          string `json:"code"`
		Msg           string
		PositionTiers []*publicdata.PositionTier `json:"data,omitempty"`
	}
	GetInterestRateAndLoanQuota struct {
		//responses.Basic
		Code                      string `json:"code"`
		Msg                       string
		InterestRateAndLoanQuotas []*publicdata.InterestRateAndLoanQuota `json:"data,omitempty"`
	}
	GetUnderlying struct {
		//responses.Basic
		Code       string `json:"code"`
		Msg        string
		Underlings [][]string `json:"data,omitempty"`
	}
	Status struct {
		//responses.Basic
		Code   string `json:"code"`
		Msg    string
		States []publicdata.State `json:"data,omitempty"`
	}
)
