package grid

import (
	"github.com/amir-the-h/okex/models/grid"
)

type (
	GridOrderList struct {
		//responses.Basic
		Code       string            `json:"code"`
		Msg        string            `json:"msg"`
		GridOrders []*grid.GridOrder `json:"data"`
	}
)
