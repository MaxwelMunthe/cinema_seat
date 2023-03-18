package cinemaSeat

import (
	"net/http"
)


type serviceImpl struct {
	Repository Repository
}

type CinemaConfigResponse struct {
	Message string
}

type SeatStatusResponse struct {
	Data []string
}

type TransactionStatusResponse struct {
	SeatStatus SeatStatus 
	Data []string
}

type SeatStatus struct {
    Free int
    Sold int
}

type SeatBookingResponse struct {
	Message string
}

func NewService(repository Repository) *serviceImpl {
	return &serviceImpl{
		Repository : repository,
	}
}
 
func (s *serviceImpl) ProcessSetCinemaConfig(seatID string, seatAmount int) (result interface{}, errStatus int, err error) {
	err = s.Repository.SetCinemaConfigInMemory(seatID, seatAmount)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return transformSetCinemaConfigResp(seatID,seatAmount), http.StatusOK, nil
}


func (s *serviceImpl) ProcessGetCinemaSeatStatus() (result interface{}, errStatus int, err error){
	data, err := s.Repository.GetCinemaSeatStatusInMemory()
    if err!= nil {
        return nil, http.StatusInternalServerError, err
    }

    return transformGetSeatStatusResp(data), http.StatusOK, nil
}

func (s *serviceImpl) ProcessGetCinemaTransactionStatus() (result interface{}, errStatus int, err error){
	data, err := s.Repository.GetCinemaSeatStatusInMemory()
    if err!= nil {
        return nil, http.StatusInternalServerError, err
    }

    return transformGetTransactionStatusResp(data), http.StatusOK, nil
}

func (s *serviceImpl) ProcessSetBookingSeat(seatID string) (result interface{}, errStatus int, err error){
	err = s.Repository.SetBookingSeatInMemory(seatID)
    if err!= nil {
        return nil, http.StatusBadRequest, err
    }

    return transformBookingSeatResp(seatID), http.StatusOK, nil
}

func (s *serviceImpl) ProcessRemoveBookingSeat(seatID string) (result interface{}, errStatus int, err error){
	err = s.Repository.RemoveBookingSeatInMemory(seatID)
    if err!= nil {
        return nil, http.StatusBadRequest, err
    }

    return transformRemoveBookingSeatResp(seatID), http.StatusOK, nil
}