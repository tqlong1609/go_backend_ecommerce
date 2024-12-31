package services

import "github.com/tqlong1609/go_backend_ecommerce/internal/repo"

type Pong1Service struct {
	pong1Repo *repo.Pong1Repo
}

func InitPong1Service() *Pong1Service {
	return &Pong1Service{
		pong1Repo: repo.InitPong1Repo(),
	}
}

func (ps *Pong1Service) GetPongService() string {
	return ps.pong1Repo.GetPong1Repo()
}
