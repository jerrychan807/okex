package trade_data

import (
	"github.com/jerrychan807/okex/models/tradedata"
)

type (
	GetSupportCoin struct {
		//responses.Basic
		Code         string `json:"code"`
		Msg          string
		SupportCoins *tradedata.SupportCoin `json:"data,omitempty"`
	}
	GetTakerVolume struct {
		//responses.Basic
		Code         string `json:"code"`
		Msg          string
		TakerVolumes []*tradedata.TakerVolume `json:"data,omitempty"`
	}
	GetRatio struct {
		//responses.Basic
		Code   string `json:"code"`
		Msg    string
		Ratios []*tradedata.Ratio `json:"data,omitempty"`
	}
	GetOpenInterestAndVolume struct {
		//responses.Basic
		Code                    string `json:"code"`
		Msg                     string
		InterestAndVolumeRatios []*tradedata.InterestAndVolumeRatio `json:"data,omitempty"`
	}
	GetPutCallRatio struct {
		//responses.Basic
		Code          string `json:"code"`
		Msg           string
		PutCallRatios []*tradedata.PutCallRatio `json:"data,omitempty"`
	}
	GetOpenInterestAndVolumeExpiry struct {
		//responses.Basic
		Code                     string `json:"code"`
		Msg                      string
		InterestAndVolumeExpires []*tradedata.InterestAndVolumeExpiry `json:"data,omitempty"`
	}
	GetOpenInterestAndVolumeStrike struct {
		//responses.Basic
		Code                     string `json:"code"`
		Msg                      string
		InterestAndVolumeStrikes []*tradedata.InterestAndVolumeStrike `json:"data,omitempty"`
	}
	GetTakerFlow struct {
		//responses.Basic
		Code      string `json:"code"`
		Msg       string
		TakerFlow *tradedata.TakerFlow `json:"data"`
	}
)
