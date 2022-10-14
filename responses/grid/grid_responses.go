package grid

import (
	"github.com/jerrychan807/okex/models/grid"
)

type (
	GridOrderList struct {
		//responses.Basic
		Code       string            `json:"code"`
		Msg        string            `json:"msg"`
		GridOrders []*grid.GridOrder `json:"data"`
	}
)
