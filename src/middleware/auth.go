package middleware

import "github.com/golang-jwt/jwt/v5"

var secretKey = []byte("ia&W^fa7")

func CreateToken(role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			role: role,
		},
	)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
