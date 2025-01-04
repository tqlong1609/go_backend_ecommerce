package repo

type IUserRepository interface {
	GetUserByEmail(email string) (string, error)
}

type userRepository struct{}

func InitUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUserByEmail(email string) (string, error) {
	return "user", nil
}
