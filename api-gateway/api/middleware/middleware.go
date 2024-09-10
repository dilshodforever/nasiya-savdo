package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"log"

	"github.com/casbin/casbin/v2"
	"github.com/dilshodforever/nasiya-savdo/api/token"
	"github.com/dilshodforever/nasiya-savdo/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JwtRoleAuth struct {
	enforcer   *casbin.Enforcer
	jwtHandler token.JWTHandler
}

func NewAuth(enforce *casbin.Enforcer) gin.HandlerFunc {
	auth := JwtRoleAuth{
		enforcer: enforce,
	}
	return func(ctx *gin.Context) {
		allow, err := auth.CheckPermission(ctx)
		if err != nil {
			valid, _ := err.(jwt.ValidationError)
			if valid.Errors == jwt.ValidationErrorExpired {
				ctx.AbortWithStatusJSON(http.StatusForbidden, "Invalid token !!!")

			} else {
				ctx.AbortWithStatusJSON(401, "Access token expired")
			}
		} else if !allow {
			ctx.AbortWithStatusJSON(http.StatusForbidden, "Permission denied")

		}
	}

}

func (a *JwtRoleAuth) GetRole(r *gin.Context) (string, error) {
	var (
		claims jwt.MapClaims
		err    error
	)
	jwtToken := r.Request.Header.Get("Authorization")

	if jwtToken == "" {
		return "unauthorized", nil
	} else if strings.Contains(jwtToken, "Basic") {
		return "unauthorized", nil
	}
	a.jwtHandler.Token = jwtToken
	a.jwtHandler.SigningKey = config.Load().TokenKey
	claims, err = a.jwtHandler.ExtractClaims()

	if err != nil {
		log.Println("Error while extracting claims: ", err)
		return "unauthorized", err
	}
	fmt.Println("role: ", claims["role"])
	return claims["role"].(string), nil
}

func (a *JwtRoleAuth) CheckPermission(r *gin.Context) (bool, error) {
	role, err := a.GetRole(r)
	if err != nil {
		log.Println("Error while getting role from token: ", err)
		return false, err
	}
	method := r.Request.Method
	path := r.FullPath()
	allowed, err := a.enforcer.Enforce(role, path, method)

	if err != nil {
		log.Println("Error while comparing role from csv list: ", err)
		return false, err
	}

	return allowed, nil
}

func GetUserId(r *gin.Context) string {
	jwtToken := r.Request.Header.Get("Authorization")

	claims, err := token.ExtractClaim(jwtToken)
	if err != nil {
		panic(err)
	}
	id := claims["id"].(string)
	return id
}


func GetStorageId(r *gin.Context) string {
	jwtToken := r.Request.Header.Get("Authorization")

	claims, err := token.ExtractClaim(jwtToken)
	if err != nil {
		panic(err)
	}
	id := claims["storage_id"].(string)
	return id
}
