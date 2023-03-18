package cinemaSeat

import (
	"testing"
	"net/http"
	"errors"

    "github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"

	"cinemaxxi_layout/infrastructure"
)

type Mock struct {
	Repository *MockRepository
}

func init_mock(t *testing.T) Mock {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	return Mock{
		Repository : NewMockRepository(ctrl),
	}
}


func Test_ProcessSetCinemaConfig(t *testing.T){
	type args struct {
		seatID string
		seatAmount int
		errorMsg error
	}

	tests := []struct {
		testCase string
		args args
		wantStatusCode  int
		wantErr bool
		result CinemaConfigResponse
	 }{
		{
			testCase : "success set configuration",
			args: args{
				seatID: "B",
				seatAmount: 10,
			},
			wantStatusCode  : http.StatusOK,
			result: CinemaConfigResponse{
				Message: "seat available B-10",
			},
		},
		{
			testCase:    "error set configuration",
			args: args{
				seatID: "B",
				seatAmount: 10,
				errorMsg: errors.New("error"),
			},
            wantStatusCode  : http.StatusBadRequest,
			wantErr : true,
		},
	 }

	 for _, tt := range tests {
		mock := init_mock(t)

		service := serviceImpl{
			Repository: mock.Repository,
		}

		mock.Repository.EXPECT().SetCinemaConfigInMemory(gomock.Any(), gomock.Any()).Return(tt.args.errorMsg)

		result, _, err := service.ProcessSetCinemaConfig(tt.args.seatID, tt.args.seatAmount)
		if tt.wantErr {
			assert.Contains(t, err.Error(), tt.args.errorMsg.Error())
		} else {
			assert.Equal(t, tt.result, result)
		}
	 }
}

func Test_ProcessGetCinemaSeatStatus(t *testing.T){
	type args struct {
		errorMsg error
		data map[string]infrastructure.CinemaConfigParam
	}

	tests := []struct {
		testCase string
		args args
		wantStatusCode  int
		wantErr bool
		result SeatStatusResponse
	 }{
		{
			testCase : "success get configuration",
			args: args{
				data: map[string]infrastructure.CinemaConfigParam{
					"B1": infrastructure.CinemaConfigParam{
						Status: "Free",
					},
				},
			},
			wantStatusCode  : http.StatusOK,
			result: SeatStatusResponse{
				Data: []string{"B1 - Free"},
			},
		},
		{
			testCase:    "error get configuration",
			args: args{
				errorMsg: errors.New("error"),
			},
            wantStatusCode  : http.StatusBadRequest,
			wantErr : true,
		},
	 }

	 for _, tt := range tests {
		mock := init_mock(t)

		service := serviceImpl{
			Repository: mock.Repository,
		}

		mock.Repository.EXPECT().GetCinemaSeatStatusInMemory().Return(tt.args.data, tt.args.errorMsg)

		result, _, err := service.ProcessGetCinemaSeatStatus()
		if tt.wantErr {
			assert.Contains(t, err.Error(), tt.args.errorMsg.Error())
		} else {
			assert.Equal(t, tt.result, result)
		}
	 }
}

func Test_ProcessGetCinemaTransactionStatus(t *testing.T){
	type args struct {
		errorMsg error
		data map[string]infrastructure.CinemaConfigParam
	}

	tests := []struct {
		testCase string
		args args
		wantStatusCode  int
		wantErr bool
		result TransactionStatusResponse
	 }{
		{
			testCase : "success get transaction status",
			args: args{
				data: map[string]infrastructure.CinemaConfigParam{
					"B1": infrastructure.CinemaConfigParam{
						Status: "Free",
					},
				},
			},
			wantStatusCode  : http.StatusOK,
			result: TransactionStatusResponse{
				SeatStatus: SeatStatus{
					Free: 1,
					Sold: 0,
				},
			},
		},
		{
			testCase:    "error get trasaction status",
			args: args{
				errorMsg: errors.New("error"),
			},
            wantStatusCode  : http.StatusBadRequest,
			wantErr : true,
		},
	 }

	 for _, tt := range tests {
		mock := init_mock(t)

		service := serviceImpl{
			Repository: mock.Repository,
		}

		mock.Repository.EXPECT().GetCinemaSeatStatusInMemory().Return(tt.args.data, tt.args.errorMsg)

		result, _, err := service.ProcessGetCinemaTransactionStatus()
		if tt.wantErr {
			assert.Contains(t, err.Error(), tt.args.errorMsg.Error())
		} else {
			assert.Equal(t, tt.result, result)
		}
	 }
}


func Test_ProcessSetBookingSeat(t *testing.T){
	type args struct {
		seatID string
		errorMsg error
	}

	tests := []struct {
		testCase string
		args args
		wantStatusCode  int
		wantErr bool
		result SeatBookingResponse
	 }{
		{
			testCase : "success set booking seat",
			args: args{
				seatID: "B10",
			},
			wantStatusCode  : http.StatusOK,
			result: SeatBookingResponse{
				Message: "success booked seat B10",
			},
		},
		{
			testCase:    "error set booking seat",
			args: args{
				seatID: "B",
				errorMsg: errors.New("error"),
			},
            wantStatusCode  : http.StatusBadRequest,
			wantErr : true,
		},
	 }

	 for _, tt := range tests {
		mock := init_mock(t)

		service := serviceImpl{
			Repository: mock.Repository,
		}

		mock.Repository.EXPECT().SetBookingSeatInMemory(gomock.Any()).Return(tt.args.errorMsg)

		result, _, err := service.ProcessSetBookingSeat(tt.args.seatID)
		if tt.wantErr {
			assert.Contains(t, err.Error(), tt.args.errorMsg.Error())
		} else {
			assert.Equal(t, tt.result, result)
		}
	 }
}

func Test_ProcessRemoveBookingSeat(t *testing.T){
	type args struct {
		seatID string
		errorMsg error
	}

	tests := []struct {
		testCase string
		args args
		wantStatusCode  int
		wantErr bool
		result SeatBookingResponse
	 }{
		{
			testCase : "success remove booking",
			args: args{
				seatID: "B10",
			},
			wantStatusCode  : http.StatusOK,
			result: SeatBookingResponse{
				Message: "success canceled booked seat B10",
			},
		},
		{
			testCase:    "error remove booking",
			args: args{
				seatID: "B",
				errorMsg: errors.New("error"),
			},
            wantStatusCode  : http.StatusBadRequest,
			wantErr : true,
		},
	 }

	 for _, tt := range tests {
		mock := init_mock(t)

		service := serviceImpl{
			Repository: mock.Repository,
		}

		mock.Repository.EXPECT().RemoveBookingSeatInMemory(gomock.Any()).Return(tt.args.errorMsg)

		result, _, err := service.ProcessRemoveBookingSeat(tt.args.seatID)
		if tt.wantErr {
			assert.Contains(t, err.Error(), tt.args.errorMsg.Error())
		} else {
			assert.Equal(t, tt.result, result)
		}
	 }
}