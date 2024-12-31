package repo

type PongRepo struct{}

func InitPongRepo() *PongRepo {
	return &PongRepo{}
}

func (p *PongRepo) GetPong() string {
	return "pong"
}
