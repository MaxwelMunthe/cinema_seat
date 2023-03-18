package cinemaSeat

import (
	"fmt"

	"cinemaxxi_layout/infrastructure"
)

func transformSetCinemaConfigResp(seatID string, seatAmount int) CinemaConfigResponse{
	response := CinemaConfigResponse{
		Message : fmt.Sprintf(SEAT_AVAILABLE,seatID,seatAmount),
	}
	return response
}

func transformGetSeatStatusResp(result map[string]infrastructure.CinemaConfigParam) SeatStatusResponse{
	var response []string 
	for i, val := range result{
		response = append(response, fmt.Sprintf(SEAT_STATUS, i, val.Status))
	}

	return SeatStatusResponse{
		Data : response,
	}
}

func transformGetTransactionStatusResp(result map[string]infrastructure.CinemaConfigParam) TransactionStatusResponse{
	var soldData []string 
	free,sold := 0 , 0

	for i, val := range result{
		if val.BookedDate == "" {
			free++
		} else {
			sold++
			soldData = append(soldData, fmt.Sprintf(TRANSACTION_STATUS, i, val.BookedDate))
		}

		
	}

	return TransactionStatusResponse{
		SeatStatus: SeatStatus{
			Free: free,
            Sold: sold,
		},
		Data : soldData,
	}
}

func transformBookingSeatResp(seatID string) SeatBookingResponse{
	return SeatBookingResponse{
		Message : fmt.Sprintf(BOOKING_SEAT_SUCCESS, seatID),
	}
}

func transformRemoveBookingSeatResp(seatID string) SeatBookingResponse{
	return SeatBookingResponse{
		Message : fmt.Sprintf(REMOVE_BOOKING_SEAT_SUCCESS, seatID),
	}
}