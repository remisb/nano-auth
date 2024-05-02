package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type SignatureAlgorithm string

type JWTAuth struct {
	alg       SignatureAlgorithm
	signKey   interface{} // private-key
	verifyKey interface{} // public-key, only used by RSA and ECDSA algorithms
	verifier  jwt.ParseOption
}

var AuthToken *JWTAuth

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//func GenerateJWT() (string, error) {
//	expirationTime := time.Now().Add(5 * time.Minute)
//	claims := &Claims{
//		Username: "username",
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(expirationTime),
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString(jwtKey)
//}

// NewSignedToken generates a signed JWT token for the given subject ID.
//
// The token contains the following claims:
//   - "sub": the provided ID as the subject of the token
//   - "exp": the expiration time of the token, set to 15 minutes from the current time
//   - "iss": the issuer of the token, set to "test"
//
// The function returns the generated token string and an error if any occurred.
func NewSignedToken(id string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": id,
			"exp": jwt.NewNumericDate(expirationTime),
			"iss": "test",
		})
	return token.SignedString(jwtKey)
}

//func GenerateUserToken(id string) (string, error) {
//	expirationTime := time.Now().Add(30 * time.Minute)
//	_, tokenString, err := Encode(jwt.MapClaims{"user_id": id, "expiration": expirationTime})
//	if err != nil {
//		return "", err
//	}
//	return tokenString, nil
//}

//func (ja *JWTAuth) Encode(claims map[string]any) (t jwt.Token, tokenString string, err error) {
//	t = jwt.New()
//	for k, v := range claims {
//		t.Set(k, v)
//	}
//	payload, err := ja.sign(t)
//	if err != nil {
//		return nil, "", err
//	}
//	tokenString = string(payload)
//	return
//}

//func CreteJwtToken(uname, pass string) (string, error) {
//
//	// Declare the expiration time of the token
//	expirationTime := time.Now().Add(5 * time.Minute)
//	// Create the JWT claims, which includes the username and expiry time
//	claims := &Claims{
//		Username: uname,
//		RegisteredClaims: jwt.RegisteredClaims{
//			// In JWT, the expiry time is expressed as unix milliseconds
//			ExpiresAt: jwt.NewNumericDate(expirationTime),
//		},
//	}
//
//	// Declare the token with the algorithm used for signing, and the claims
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	// Create the JWT string
//	return token.SignedString(jwtKey)
//}
