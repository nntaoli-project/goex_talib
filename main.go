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
	inHigh := realData(data, InHigh)
	inLow := realData(data, InLow)
	inClose := realData(data, InClose)
	return talib.Atr(inHigh, inLow, inClose, inTimePeriod)
}

func Macd(data []goex.Kline, inFastPeriod int, inSlowPeriod int, inSignalPeriod int, priceTy PriceType) (DIF, DEA, MACD []float64) {
	var macd []float64
	dif, dea, hist := talib.Macd(realData(data, priceTy), inFastPeriod, inSlowPeriod, inSignalPeriod)
	for _, item := range hist {
		macd = append(macd, item*2)
	}
	return dif, dea, macd
}

func Boll(data []goex.Kline, inTimePeriod int, deviation float64, priceTy PriceType) (up, middle, low []float64) {
	return talib.BBands(realData(data, priceTy), inTimePeriod, deviation, deviation, 0)
}

func Rsi(data []goex.Kline, inTimePeriod int, priceTy PriceType) []float64 {
	return talib.Rsi(realData(data, priceTy), inTimePeriod)
}

func Stoch(data []goex.Kline, fastKPeriod, slowKPeriod int, slowKMaty talib.MaType, slowDPeriod int, slowDMAty talib.MaType) (outSlowK []float64, outSlowD []float64) {
	inHigh := realData(data, InHigh)
	inLow := realData(data, InLow)
	inClose := realData(data, InClose)
	outSlowK, outSlowD = talib.Stoch(inHigh, inLow, inClose, fastKPeriod, slowKPeriod, slowKMaty, slowDPeriod, slowDMAty)
	return
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
