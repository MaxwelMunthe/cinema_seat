package gateway

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"errors"

    "github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"

	"cinemaxxi_layout/cinemaSeat"
)

type Mock struct {
	Service *cinemaSeat.MockService
}

func init_mock(t *testing.T) Mock {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	return Mock{
		Service : cinemaSeat.NewMockService(ctrl),
	}
}


func Test_SetSeatConfiguration(t *testing.T){
	type args struct {
		seatID string
		seatAmount string
	}

	tests := []struct {
		testCase string
		args args
		wantStatusCode  int
	 }{
		{
			testCase : "success set configuration",
			args: args{
				seatID: "B",
				seatAmount: "10",
			},
			wantStatusCode  : http.StatusOK,
		},
		{
			testCase:    "error set configuration",
            wantStatusCode  : http.StatusBadRequest,
		},
	 }

	 for _, tt := range tests {
		mock := init_mock(t)

		delivery := HttpDelivery{
			service: mock.Service,
		}

		r := httptest.NewRequest(http.MethodGet, "/cinema/config/set?seat_id="+tt.args.seatID+"&seat_amount="+tt.args.seatAmount, nil)
		w := httptest.NewRecorder()

		mock.Service.EXPECT().ProcessSetCinemaConfig(gomock.Any(), gomock.Any()).Return(nil, tt.wantStatusCode, nil)

		delivery.SetSeatConfiguration(w, r)
		assert.Equal(t, tt.wantStatusCode, w.Code)
	 }
}

func Test_GetSeatStatus(t *testing.T){
	type args struct {
		errorMsg error
	}

	tests := []struct {
		testCase string
		args args
		wantStatusCode  int
	 }{
		{
			testCase : "success get seat status",
			wantStatusCode  : http.StatusOK,
		},
		{
			testCase:    "error get seat status",
			args:        args{
				errorMsg: errors.New("error"),
			},
            wantStatusCode  : http.StatusBadRequest,
		},
	 }

	 for _, tt := range tests {
		mock := init_mock(t)

		delivery := HttpDelivery{
			service: mock.Service,
		}

		r := httptest.NewRequest(http.MethodGet, "/cinema/seat/status", nil)
		w := httptest.NewRecorder()

		mock.Service.EXPECT().ProcessGetCinemaSeatStatus().Return(nil, tt.wantStatusCode, tt.args.errorMsg)

		delivery.GetSeatStatus(w, r)
		assert.Equal(t, tt.wantStatusCode, w.Code)
	 }
}

func Test_GetTransactionStatus(t *testing.T){
	type args struct {
		errorMsg error
	}

	tests := []struct {
		testCase string
		args args
		wantStatusCode  int
	 }{
		{
			testCase : "success get transaction status",
			wantStatusCode  : http.StatusOK,
		},
		{
			testCase:    "error get transaction status",
			args:        args{
				errorMsg: errors.New("error"),
			},
            wantStatusCode  : http.StatusBadRequest,
		},
	 }

	 for _, tt := range tests {
		mock := init_mock(t)

		delivery := HttpDelivery{
			service: mock.Service,
		}

		r := httptest.NewRequest(http.MethodGet, "/cinema/transaction/status", nil)
		w := httptest.NewRecorder()

		mock.Service.EXPECT().ProcessGetCinemaTransactionStatus().Return(nil, tt.wantStatusCode, tt.args.errorMsg)

		delivery.GetTransactionStatus(w, r)
		assert.Equal(t, tt.wantStatusCode, w.Code)
	 }
}

func Test_SetBoookingSeat(t *testing.T){
	type args struct {
		errorMsg error
		seatID string
	}

	tests := []struct {
		testCase string
		args args
		wantStatusCode  int
	 }{
		{
			testCase : "success get transaction status",
			args : args{
				seatID: "B",
			},
			wantStatusCode  : http.StatusOK,
		},
		{
			testCase:    "error get transaction status",
			args:        args{
				errorMsg: errors.New("error"),
			},
            wantStatusCode  : http.StatusBadRequest,
		},
	 }

	 for _, tt := range tests {
		mock := init_mock(t)

		delivery := HttpDelivery{
			service: mock.Service,
		}

		r := httptest.NewRequest(http.MethodGet, "/cinema/seat/booking/set?seat_id="+tt.args.seatID, nil)
		w := httptest.NewRecorder()

		mock.Service.EXPECT().ProcessSetBookingSeat(gomock.Any()).Return(nil, tt.wantStatusCode, tt.args.errorMsg)

		delivery.SetBoookingSeat(w, r)
		assert.Equal(t, tt.wantStatusCode, w.Code)
	 }
}

func Test_RemoveBoookingSeat(t *testing.T){
	type args struct {
		errorMsg error
		seatID string
	}

	tests := []struct {
		testCase string
		args args
		wantStatusCode  int
	 }{
		{
			testCase : "success get transaction status",
			args : args{
				seatID: "B",
			},
			wantStatusCode  : http.StatusOK,
		},
		{
			testCase:    "error get transaction status",
			args:        args{
				errorMsg: errors.New("error"),
			},
            wantStatusCode  : http.StatusBadRequest,
		},
	 }

	 for _, tt := range tests {
		mock := init_mock(t)

		delivery := HttpDelivery{
			service: mock.Service,
		}

		r := httptest.NewRequest(http.MethodGet, "/cinema/seat/booking/remove?seat_id="+tt.args.seatID, nil)
		w := httptest.NewRecorder()

		mock.Service.EXPECT().ProcessRemoveBookingSeat(gomock.Any()).Return(nil, tt.wantStatusCode, tt.args.errorMsg)

		delivery.RemoveBoookingSeat(w, r)
		assert.Equal(t, tt.wantStatusCode, w.Code)
	 }
}