package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type JWTService interface {
	GenerateToken(id string) (string, int)
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtCustomClaim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}
type jwtService struct {
	secretKey string
}

func getSecretKey() string {

	err := godotenv.Load(".env")
	if err != nil {
		panic("can't load env file")
	}
	return os.Getenv("SECRET_KEY")
}

func NewJwtService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}
func (j *jwtService) GenerateToken(id string) (string, int) {
	var customClaim jwtCustomClaim
	customClaim.Id = id
	ttl := 60 * time.Minute
	customClaim.StandardClaims.ExpiresAt = time.Now().UTC().Add(ttl).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &customClaim)
	signedToken, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		panic(err)
	}
	return signedToken, int(ttl)
}
func (jwtSrv *jwtService) ValidateToken(tokenFromHeader string) (*jwt.Token, error) {
	return jwt.Parse(tokenFromHeader, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(jwtSrv.secretKey), nil
	})
}
