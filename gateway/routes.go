package gateway

import(
	"github.com/go-chi/chi"

	"cinemaxxi_layout/cinemaSeat"
)

type Routes struct {
	handler *HttpDelivery
}

func NewRoutes(handler *HttpDelivery) *Routes {
	return &Routes{handler: handler}
}

func InjectRoutes() *Routes {
	services := cinemaSeat.InjectServiceMovie()
	delivery := NewHttpDelivery(services)

	return NewRoutes(delivery)
}

func (routes Routes) RegisterRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Post("/cinema/config/set", routes.handler.SetSeatConfiguration)
		r.Get("/cinema/seat/status", routes.handler.GetSeatStatus)
		r.Get("/cinema/transaction/status", routes.handler.GetTransactionStatus)
		r.Post("/cinema/seat/booking/set", routes.handler.SetBoookingSeat)
		r.Post("/cinema/seat/booking/remove", routes.handler.RemoveBoookingSeat)
	})
}