package auth

type Auth interface {
	GeneratePasswordHash(password string) string
	GenerateToken(id int) (signedString string, err error)
	ParseToken(signedString string) (id int, err error)
}
