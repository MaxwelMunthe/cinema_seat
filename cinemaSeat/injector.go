package cinemaSeat

import	(
	"cinemaxxi_layout/infrastructure"
)

func InjectServiceMovie() *serviceImpl {
	repository := NewRepository(infrastructure.CinemaConfig)
	service := NewService(repository)

	return service
}