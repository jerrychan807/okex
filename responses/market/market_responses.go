package market

import (
	"github.com/amir-the-h/okex/models/market"
)

type (
	Ticker struct {
		//responses.Basic
		Code    string           `json:"code"`
		Msg     string           `json:"msg"`
		Tickers []*market.Ticker `json:"data,omitempty"`
	}
	IndexTicker struct {
		//responses.Basic
		Code         string                `json:"code"`
		Msg          string                `json:"msg"`
		IndexTickers []*market.IndexTicker `json:"data,omitempty"`
	}
	OrderBook struct {
		//responses.Basic
		Code       string              `json:"code"`
		Msg        string              `json:"msg"`
		OrderBooks []*market.OrderBook `json:"data,omitempty"`
	}
	Candle struct {
		//responses.Basic
		Code    string           `json:"code"`
		Msg     string           `json:"msg"`
		Candles []*market.Candle `json:"data,omitempty"`
	}
	IndexCandle struct {
		//responses.Basic
		Code    string                `json:"code"`
		Msg     string                `json:"msg"`
		Candles []*market.IndexCandle `json:"data,omitempty"`
	}
	CandleMarket struct {
		//responses.Basic
		Code    string                `json:"code"`
		Msg     string                `json:"msg"`
		Candles []*market.IndexCandle `json:"data,omitempty"`
	}
	Trade struct {
		//responses.Basic
		Code   string          `json:"code"`
		Msg    string          `json:"msg"`
		Trades []*market.Trade `json:"data,omitempty"`
	}
	TotalVolume24H struct {
		//responses.Basic
		Code            string                   `json:"code"`
		Msg             string                   `json:"msg"`
		TotalVolume24Hs []*market.TotalVolume24H `json:"data,omitempty"`
	}
	IndexComponent struct {
		//responses.Basic
		Code            string                 `json:"code"`
		Msg             string                 `json:"msg"`
		IndexComponents *market.IndexComponent `json:"data,omitempty"`
	}
)
