package token

import (
	"golang-jwt/db/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	privateKeyPath = "keys/app.rsa"
	publicKeyPath  = "keys/app.rsa.pub"
)

func InitJWt() error {
	//reading private key from the file
	signBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return err
	}

	//parse the private key
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}

	//reading public key
	verifyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}
	return nil
}

func CreateNewToken(uuid string, role string) (authToken, refreshToken, csrfToken string, err error) {
	// genrearate csrf secret
	csrfSecret, err := models.GenrerateCSRFSecret()
	if err != nil {
		return
	}

	//create refresh token
	refreshToken, err = createRefreshToken(uuid, role, csrfSecret)
	if err != nil {
		return
	}
	//create auth token
	authToken, err = createAuthToken(uuid, role, csrfSecret)
	if err != nil {
		return
	}
	return
}

func CheckAndRefreshTokens() {

}

func createAuthToken(uuid, role, csrfSecret string) (string, error) {
	authTokenExp := time.Now().Add(models.AuthTokenvalidTime)
	authClaims := models.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: uuid,
			ExpiresAt: &jwt.NumericDate{
				Time: authTokenExp,
			},
		},
		Role: role,
		Csrf: csrfSecret,
	}
	auhtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, authClaims)

	return auhtToken.SignedString(signKey)
}

func createRefreshToken(uuid, role, csrfSecret string) (string, error) {

	refreshTokenExp := time.Now().Add(models.RefreshTokenValidTime)
	refreshJti, err := db.StoreRefreshToken()
	if err != nil {
		return "", err
	}
	refreshClaims := models.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:      refreshJti,
			Subject: uuid,
			ExpiresAt: &jwt.NumericDate{
				Time: refreshTokenExp,
			},
		},
		Role: role,
		Csrf: csrfSecret,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)

	return refreshToken.SignedString(signKey)
}

func updateRefreshTokenExp() {

}

func UpdateAuthToken() {}

func RevokeRefreshTOken() error {

}

func UpdateRefreshTokenCsrf() {

}

func GrabUUID() {

}
