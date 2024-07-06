package tokens

import (
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtIssuer = "auth-service"
)

type (
	jwtAuthKey string
	UserClaims struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	Tokens struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

const AuthReqCtxKey jwtAuthKey = "authUser"

var (
	// global secret
	privateKey = []byte(os.Getenv("JWT_PVT_KEY"))
	publicKey = []byte(os.Getenv("JWT_PUB_KEY"))

	// errors
	errTokenExpired     = errors.New("token expired")
	errSignatureInvalid = errors.New("invalid token signature")
	errParsingToken     = errors.New("error parsing token")
	errUserEmpty        = errors.New("user cannot be empty")
	errTokenCreate      = errors.New("error creating token")
	errParsingPemKey    = errors.New("error parsing pem key")
)

func newToken(claims jwt.RegisteredClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	secret, err := jwt.ParseECPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Printf("error parsing pem key: %v", err)
		return "", errParsingPemKey
	}

	return accessToken.SignedString(secret)
}

// func pemKeyPair(key *ecdsa.PrivateKey) (privKeyPEM []byte, pubKeyPEM []byte, err error) {
// 	der, err := x509.MarshalECPrivateKey(key)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	privKeyPEM = pem.EncodeToMemory(&pem.Block{
// 		Type:  "EC PRIVATE KEY",
// 		Bytes: der,
// 	})

// 	der, err = x509.MarshalPKIXPublicKey(key.Public())
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	pubKeyPEM = pem.EncodeToMemory(&pem.Block{
// 		Type:  "EC PUBLIC KEY",
// 		Bytes: der,
// 	})

// 	return
// }
