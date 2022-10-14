package rest

import (
	"encoding/json"
	"github.com/jerrychan807/okex"
	requests "github.com/jerrychan807/okex/requests/rest/grid"
	responses "github.com/jerrychan807/okex/responses/grid"
	"net/http"
)

type Grid struct {
	client *ClientRest
}

// NewTrade returns a pointer to a fresh Trade
func NewGrid(c *ClientRest) *Grid {
	return &Grid{c}
}

func (c *Grid) GetGridHistoryOrderList(req requests.GridHistoryOrderList) (response responses.GridOrderList, err error) {
	p := "/api/v5/tradingBot/grid/orders-algo-history"

	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}
