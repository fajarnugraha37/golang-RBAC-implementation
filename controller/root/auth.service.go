package root

type AuthService interface {
	Login(username string, password string) string
	Register(username string, password string, confirmPassword string) string
}
