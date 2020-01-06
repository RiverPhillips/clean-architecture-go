package interfaces

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/RiverPhillips/clean-architecture-go/usecases"
)

type OrderInteractor interface {
	Items(userId, orderId int) ([]usecases.Item, error)
	Add(userId, orderId, itemId int) error
}

type WebServiceHandler struct {
	OrderInteractor OrderInteractor
}

func (handler WebServiceHandler) ShowOrders(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))
	orderId, _ := strconv.Atoi(req.FormValue("orderId"))
	items, _ := handler.OrderInteractor.Items(userId, orderId)
	for _, item := range items {
		io.WriteString(res, fmt.Sprintf("item id: %d\n", item.Id))
		io.WriteString(res, fmt.Sprintf("item name: %d\n", item.Name))
		io.WriteString(res, fmt.Sprintf("item value: %d\n", item.Value))
	}
}
