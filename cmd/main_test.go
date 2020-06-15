package cmd

import (
	"github.com/markcheno/go-talib"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/okex"
	"github.com/nntaoli-project/goex_talib"
	"net/http"
	"testing"
)

var api = okex.NewOKEx(&goex.APIConfig{
	HttpClient: http.DefaultClient,
	Endpoint:   "https://www.okex.me",
})

func Test_talib(t *testing.T) {
	data, _ := api.GetKlineRecords(goex.BTC_USDT, goex.KLINE_PERIOD_1H, 300, 0)

	t.Log(goex_talib.Ma(data, 60, talib.EMA, goex_talib.InClose))
	t.Log(goex_talib.Atr(data, 20))

	//boll
	up, middle, low := goex_talib.Boll(data, 20, 2)
	t.Log(up[len(up)-1], middle[len(middle)-1], low[len(low)-1])

	// macd
	dif, dea, macd := goex_talib.Macd(data, 12, 26, 9, goex_talib.InClose)
	t.Log("dif=", dif[len(dif)-1], ",dea=", dea[len(dea)-1], ",macd=", macd[len(macd)-1])
}
