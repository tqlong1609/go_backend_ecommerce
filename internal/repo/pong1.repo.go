package repo

type Pong1Repo struct{}

func InitPong1Repo() *Pong1Repo {
	return &Pong1Repo{}
}

func (pr *Pong1Repo) GetPong1Repo() string {
	return "pong1"
}
