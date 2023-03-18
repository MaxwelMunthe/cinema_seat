package gateway

import (
	"net/http"
	"strconv"

	"cinemaxxi_layout/shared"
	"cinemaxxi_layout/cinemaSeat"
	
)

type HttpDelivery struct {
	service cinemaSeat.Service
}

func NewHttpDelivery(service cinemaSeat.Service) *HttpDelivery {
	return &HttpDelivery{service: service}
}
 
func (delivery HttpDelivery) SetSeatConfiguration(w http.ResponseWriter, r *http.Request) {
	seatID := r.URL.Query().Get("seat_id")
	seatAmount := r.URL.Query().Get("seat_amount")

	seatAmountInt, err := strconv.Atoi(seatAmount)
	if err!= nil {
        shared.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
        return
    }

	res, status, err := delivery.service.ProcessSetCinemaConfig(seatID, seatAmountInt)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), status)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (delivery HttpDelivery) GetSeatStatus(w http.ResponseWriter, r *http.Request) {
	res, status, err := delivery.service.ProcessGetCinemaSeatStatus()
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), status)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (delivery HttpDelivery) GetTransactionStatus(w http.ResponseWriter, r *http.Request) {
	res, status, err := delivery.service.ProcessGetCinemaTransactionStatus()
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), status)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}


func (delivery HttpDelivery) SetBoookingSeat(w http.ResponseWriter, r *http.Request) {
	seatID := r.URL.Query().Get("seat_id")

	res, status, err := delivery.service.ProcessSetBookingSeat(seatID)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), status)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (delivery HttpDelivery) RemoveBoookingSeat(w http.ResponseWriter, r *http.Request) {
	seatID := r.URL.Query().Get("seat_id")

	res, status, err := delivery.service.ProcessRemoveBookingSeat(seatID)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), status)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}