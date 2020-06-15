package goex_talib

import (
	"github.com/markcheno/go-talib"
	"github.com/nntaoli-project/goex"
)

type PriceType int

const (
	InClose PriceType = iota + 1
	InHigh
	InLow
	InOpen
)

func Ma(data []goex.Kline, inTimePeriod int, maType talib.MaType, priceTy PriceType) []float64 {
	return talib.Ma(realData(data, priceTy), inTimePeriod, maType)
}

func Atr(data []goex.Kline, inTimePeriod int) []float64 {
	var (
		inHigh  []float64
		inLow   []float64
		inClose []float64
	)

	for i := len(data) - 1; i >= 0; i-- {
		k := data[i]
		inHigh = append(inHigh, k.High)
		inLow = append(inLow, k.Low)
		inClose = append(inClose, k.Close)
	}

	return talib.Atr(inHigh, inLow, inClose, inTimePeriod)
}

func Macd(data []goex.Kline, inFastPeriod int, inSlowPeriod int, inSignalPeriod int, priceTy PriceType) (outMACD, outMACDSignal, outMACDHist []float64) {
	return talib.Macd(realData(data, priceTy), inFastPeriod, inSlowPeriod, inSignalPeriod)
}

func realData(data []goex.Kline, priceTy PriceType) []float64 {
	var inReal []float64
	for i := len(data) - 1; i >= 0; i-- {
		k := data[i]
		switch priceTy {
		case InClose:
			inReal = append(inReal, k.Close)
		case InHigh:
			inReal = append(inReal, k.High)
		case InLow:
			inReal = append(inReal, k.Low)
		case InOpen:
			inReal = append(inReal, k.Open)
		default:
			panic("please set ema type")
		}
	}
	return inReal
}

func Boll(data []goex.Kline, inTimePeriod int, deviation float64) (up, middle, low []float64) {
	return talib.BBands(realData(data, InClose), inTimePeriod, deviation, deviation, 0)
}
