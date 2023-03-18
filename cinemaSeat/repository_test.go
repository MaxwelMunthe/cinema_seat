package cinemaSeat

import (
	"testing"
	"errors"

    "github.com/stretchr/testify/assert"

	"cinemaxxi_layout/infrastructure"
)

type RepoMock struct {
	Repository map[string]infrastructure.CinemaConfigParam
}

func initRepo_mock(t *testing.T) RepoMock {
	return RepoMock{
		Repository : make(map[string]infrastructure.CinemaConfigParam),
	}
}


func Test_SetCinemaConfigInMemory(t *testing.T){
	type args struct {
		seatID string
		seatAmount int
		errorMsg error
	}

	tests := []struct {
		testCase string
		args args
		wantErr bool
	 }{
		{
			testCase : "success set configuration",
			args: args{
				seatID: "B",
				seatAmount: 10,
			},
		},
	 }

	 for _, tt := range tests {
		mock := initRepo_mock(t)

		repo := RepositoryImpl{
			Config: mock.Repository,
		}

		err := repo.SetCinemaConfigInMemory(tt.args.seatID, tt.args.seatAmount)
		assert.Equal(t, err, tt.args.errorMsg)
	 }
}

func Test_GetCinemaSeatStatusInMemory(t *testing.T){
	type args struct {
		errorMsg error
	}

	tests := []struct {
		testCase string
		args args
		wantErr bool
	 }{
		{
			testCase : "success get seat status",
		},
	 }

	 for _, tt := range tests {
		mock := initRepo_mock(t)

		repo := RepositoryImpl{
			Config: mock.Repository,
		}

		_, err := repo.GetCinemaSeatStatusInMemory()
		assert.Equal(t, err, tt.args.errorMsg)
	 }
}

func Test_SetBookingSeatInMemory(t *testing.T){
	type args struct {
		seatID string
		errorMsg error
	}

	tests := []struct {
		testCase string
		args args
		wantErr bool
	 }{
		{
			testCase : "success set seat booking",
			args: args{
				seatID: "B",
			},
		},
		{
			testCase : "error get seat not found",
			args: args{
				seatID: "B1",
				errorMsg : errors.New("seat not found"),
			},
		},
	 }

	 for _, tt := range tests {
		repo := RepositoryImpl{
			Config: map[string]infrastructure.CinemaConfigParam{
				"B": infrastructure.CinemaConfigParam{
					Status: "Free",
				},
			},
		}

		err := repo.SetBookingSeatInMemory(tt.args.seatID)
		assert.Equal(t, err, tt.args.errorMsg)
	 }
}

func Test_RemoveBookingSeatInMemory(t *testing.T){
	type args struct {
		seatID string
		errorMsg error
	}

	tests := []struct {
		testCase string
		args args
		wantErr bool
	 }{
		{
			testCase : "success set seat booking",
			args: args{
				seatID: "B",
			},
		},
		{
			testCase : "error get seat not found",
			args: args{
				seatID: "B1",
				errorMsg : errors.New("seat not found"),
			},
		},
	 }

	 for _, tt := range tests {
		repo := RepositoryImpl{
			Config: map[string]infrastructure.CinemaConfigParam{
				"B": infrastructure.CinemaConfigParam{
					Status: "Free",
				},
			},
		}

		err := repo.RemoveBookingSeatInMemory(tt.args.seatID)
		assert.Equal(t, err, tt.args.errorMsg)
	 }
}