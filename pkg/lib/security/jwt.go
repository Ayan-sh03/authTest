package security

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateJWT generates a JSON Web Token (JWT) for the given email and username.
//
// Parameters:
// - email: a string representing the email of the user.
// - username: a string representing the username of the user.
//
// Returns:
// - tokenString: a string representing the generated JWT.
// - err: an error object indicating any error that occurred during JWT generation.
func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(240 * time.Hour)

	claims := jwt.MapClaims{
		"exp":        expirationTime.Unix(),
		"authorized": true,
		"email":      email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(jwtKey)

	return tokenStr, err
}
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		return err
	}

}
