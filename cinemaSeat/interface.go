package cinemaSeat

import (
    "cinemaxxi_layout/infrastructure"
)

type Service interface {
	ProcessSetCinemaConfig(seatID string, seatAmount int) (result interface{}, errStatus int, err error)
	ProcessGetCinemaSeatStatus() (result interface{}, errStatus int, err error)
	ProcessGetCinemaTransactionStatus() (result interface{}, errStatus int, err error)
	ProcessSetBookingSeat(seatID string) (result interface{}, errStatus int, err error)
	ProcessRemoveBookingSeat(seatID string) (result interface{}, errStatus int, err error)
}

type Repository interface {
	SetCinemaConfigInMemory(eatID string, seatAmount int) (err error)
	GetCinemaSeatStatusInMemory() (result map[string]infrastructure.CinemaConfigParam, err error)
	SetBookingSeatInMemory(seatID string) (err error)
	RemoveBookingSeatInMemory(seatID string) (err error)
}