package auth

type JwtHelper interface {
	GenerateToken(userId int, roles ...string) (string, error)
}
