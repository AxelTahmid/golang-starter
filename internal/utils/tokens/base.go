package tokens

import (
	"crypto/ecdsa"
	"errors"
	"log"
	"os"
	"sync"

	"github.com/AxelTahmid/golang-starter/config"
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
	once sync.Once

	// global secret
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey

	// errors
	errTokenExpired     = errors.New("token expired")
	errSignatureInvalid = errors.New("invalid token signature")
	errParsingToken     = errors.New("error parsing token")
	errUserEmpty        = errors.New("user cannot be empty")
	errTokenCreate      = errors.New("error creating token")
)

// function to parse the private and public keys
func SetJwtKeyPair(conf config.Jwt) {

	once.Do(func() {
		jwtPrivateKey, err := os.ReadFile(conf.JwtPvtKeyPath)
		if err != nil {
			log.Fatalf("error reading private key: %v", err)
		}

		privateKey, err = jwt.ParseECPrivateKeyFromPEM(jwtPrivateKey)
		if err != nil {
			log.Fatalf("error parsing private key: %v", err)
		}

		jwtPublicKey, err := os.ReadFile(conf.JwtPubKeyPath)
		if err != nil {
			log.Fatalf("error reading public key: %v", err)
		}

		publicKey, err = jwt.ParseECPublicKeyFromPEM(jwtPublicKey)
		if err != nil {
			log.Fatalf("error parsing public key: %v", err)
		}

		log.Println("Parsed private and public keys for JWT")
	})

}

func newToken(claims jwt.RegisteredClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return accessToken.SignedString(privateKey)
}
