package cinemaSeat

import (
	"fmt"
	"time"

	"cinemaxxi_layout/infrastructure"
)

type RepositoryImpl struct {
	Config map[string]infrastructure.CinemaConfigParam
}

func NewRepository(config map[string]infrastructure.CinemaConfigParam) *RepositoryImpl {
	return &RepositoryImpl{
		Config: config,
	}
}

func (r RepositoryImpl) SetCinemaConfigInMemory(seatID string, seatAmount int) (err error) {
	for i := 1; i <= seatAmount; i++ {
		r.Config[fmt.Sprintf(SEAT_ID, seatID, i)] = infrastructure.CinemaConfigParam{
			Status: SEAT_FREE,
			CreatedDate: time.Now().Format("02-01-2006 15:04:05"),
		}
	}

	return
}

func (r RepositoryImpl) GetCinemaSeatStatusInMemory() (result map[string]infrastructure.CinemaConfigParam, err error) {
	return r.Config, nil
}


func (r RepositoryImpl) SetBookingSeatInMemory(seatID string) (err error) {
	val, ok := r.Config[seatID]
	if !ok {
		return fmt.Errorf(ERR_SEAT_NOT_FOUND)
	}

	val.Status = SEAT_SOLD
	val.BookedDate = time.Now().Format("02-01-2006 15:04:05")

	r.Config[seatID] = val
	return
}

func (r RepositoryImpl) RemoveBookingSeatInMemory(seatID string) (err error) {
	val, ok := r.Config[seatID]
	if !ok {
		return fmt.Errorf(ERR_SEAT_NOT_FOUND)
	}

	val.Status = SEAT_FREE
	val.BookedDate = ""

	r.Config[seatID] = val
	return
}