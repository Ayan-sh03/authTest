package security

import (
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
var jwtLife = []byte(os.Getenv("JWT_LIFETIME"))

// GenerateJWT generates a JSON Web Token (JWT) for the given email and username.
//
// Parameters:
// - email: a string representing the email of the user.
// - username: a string representing the username of the user.
//
// Returns:
// - tokenString: a string representing the generated JWT.
// - err: an error object indicating any error that occurred during JWT generation.
func GenerateJWT(email string, id int64) (string, bool) {
	// Removed expirationTime variable.
	expTime, expErr := strconv.ParseInt(string(jwtLife), 10, 64) //! Added variable
	if expErr != nil {
		return "", false
	}
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(expTime) * time.Hour).Unix(),
		// Removed authorization field
		"email": email,
		// Added id field
		"id": id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", false
	}

	return tokenStr, true
}

// Changed return type
func ValidateJWT(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return false
	}
	return true
}
