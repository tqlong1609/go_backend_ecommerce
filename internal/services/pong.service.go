package services

import "github.com/tqlong1609/go_backend_ecommerce/internal/repo"

type PongService struct {
	pongRepo *repo.PongRepo
}

func InitPongService() *PongService {
	return &PongService{
		pongRepo: repo.InitPongRepo(),
	}
}

func (ps *PongService) GetPongService() string {
	return ps.pongRepo.GetPong()
}
