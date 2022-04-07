package cmd

import (
	"github.com/markcheno/go-talib"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/okex/v5"
	"github.com/nntaoli-project/goex_talib"
	"net/http"
	"testing"
)

var api = okex.NewOKExV5Spot(&goex.APIConfig{
	HttpClient: http.DefaultClient,
	Endpoint:   "https://www.okex.com",
})

func Test_talib(t *testing.T) {
	data, err := api.GetKlineRecords(goex.BTC_USDT, goex.KLINE_PERIOD_1H, 100)
	if len(data) == 0 {
		t.Log(err)
		return
	}
	t.Log(goex_talib.Rsi(data, 25, goex_talib.InClose))
	t.Log(goex_talib.Stoch(data, 7, 15, talib.SMA, 15, talib.SMA))
	t.Log(goex_talib.Ma(data, 60, talib.SMA, goex_talib.InClose)) //基于收盘价的简单移动均线
	t.Log(goex_talib.Atr(data, 20))                               // atr

	//boll
	up, middle, low := goex_talib.Boll(data, 20, 2, goex_talib.InClose)
	t.Log(up[len(up)-1], middle[len(middle)-1], low[len(low)-1])

	// macd
	dif, dea, macd := goex_talib.Macd(data, 12, 26, 9, goex_talib.InClose)
	t.Log("dif=", dif[len(dif)-1], ",dea=", dea[len(dea)-1], ",macd=", macd[len(macd)-1])
}
