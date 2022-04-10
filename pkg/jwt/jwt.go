package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// Credential ...
type Credential struct {
	Secret           string
	ExpSecret        int
	RefreshSecret    string
	RefreshExpSecret int
}

// jwtClaims ...
type jwtClaims struct {
	jwt.StandardClaims
}

// GetToken ...
func (cred *Credential) GetToken(id string) (string, string, error) {
	expirationTime := time.Now().Add(time.Duration(cred.ExpSecret) * time.Hour).Unix()

	unixTimeUTC := time.Unix(expirationTime, 0)
	unitTimeInRFC3339 := unixTimeUTC.UTC().Format(time.RFC3339)

	claims := &jwtClaims{
		jwt.StandardClaims{
			Id:        id,
			ExpiresAt: expirationTime,
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(cred.Secret))

	return token, unitTimeInRFC3339, err
}

// GetRefreshToken ...
func (cred *Credential) GetRefreshToken(id string) (string, string, error) {
	expirationTime := time.Now().Add(time.Duration(cred.RefreshExpSecret) * time.Hour).Unix()

	unixTimeUTC := time.Unix(expirationTime, 0)
	unitTimeInRFC3339 := unixTimeUTC.UTC().Format(time.RFC3339) // converts utc time to RFC3339 format

	claims := &jwtClaims{
		jwt.StandardClaims{
			Id:        id,
			ExpiresAt: expirationTime,
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(cred.RefreshSecret))

	return token, unitTimeInRFC3339, err
}
