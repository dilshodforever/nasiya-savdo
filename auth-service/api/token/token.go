package token

import (
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/dilshodforever/nasiya-savdo/config"
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
)

type JWTHandler struct {
	Sub        string
	Exp        string
	Iat        string
	Role       string
	SigningKey string
	Token      string
	Timeout    int
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

var tokenKey = config.Load().TokenKey

func GenereteAccsessJWTToken(user *pb.UserLoginRes) *Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["full_name"] = user.FullName
	claims["email"] = user.Email
	claims["phone_number"] = user.PhoneNumber
	claims["role"] = "admin"
	claims["storage_id"] = user.StorageId
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	access, err := accessToken.SignedString([]byte(tokenKey))
	if err != nil {
		log.Fatal("error while genereting access token : ", err)
	}

	return &Tokens{
		AccessToken: access,
	}
}

func GenereteJWTToken(user *pb.UserLoginRes) *Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["full_name"] = user.FullName
	claims["email"] = user.Email
	claims["phone_number"] = user.PhoneNumber
	claims["role"] = "admin"
	claims["storage_id"] = user.StorageId
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	access, err := accessToken.SignedString([]byte(tokenKey))
	if err != nil {
		log.Fatal("error while genereting access token : ", err)
	}

	rftclaims := refreshToken.Claims.(jwt.MapClaims)
	rftclaims["id"] = user.Id
	rftclaims["full_name"] = user.FullName
	rftclaims["email"] = user.Email
	rftclaims["phone_number"] = user.PhoneNumber
	rftclaims["role"] = "admin"
	claims["storage_id"] = user.StorageId
	rftclaims["iat"] = time.Now().Unix()
	rftclaims["exp"] = time.Now().Add(30 * 24 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(tokenKey))
	if err != nil {
		log.Fatal("error while genereting refresh token : ", err)
	}

	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func ExtractClaim(cfg *config.Config, tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.TokenKey), nil
	}

	token, err = jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return claims, nil
}

func (jwtHandler *JWTHandler) ExtractClaims() (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtHandler.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtHandler.SigningKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		slog.Error("invalid jwt token")
		return nil, err
	}
	return claims, nil
}

func GetIdFromToken(r *http.Request, cfg *config.Config) (string, int) {
	var softToken string
	token := r.Header.Get("Authorization")

	if token == "" {
		return "unauthorized", http.StatusUnauthorized
	} else if strings.Contains(token, "Bearer") {
		softToken = strings.TrimPrefix(token, "Bearer ")
	} else {
		softToken = token
	}

	claims, err := ExtractClaim(cfg, softToken)
	if err != nil {
		return "unauthorized", http.StatusUnauthorized
	}

	return cast.ToString(claims["id"]), 0
}

func GetEmailFromToken(r *http.Request, cfg *config.Config) (string, int) {
	var softToken string
	token := r.Header.Get("Authorization")

	if token == "" {
		return "unauthorized", http.StatusUnauthorized
	} else if strings.Contains(token, "Bearer") {
		softToken = strings.TrimPrefix(token, "Bearer ")
	} else {
		softToken = token
	}

	claims, err := ExtractClaim(cfg, softToken)
	if err != nil {
		return "unauthorized", http.StatusUnauthorized
	}

	return cast.ToString(claims["email"]), 0
}
