package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"server/packages/config"
	"server/packages/utils"
	"net/http"
	"strings"
	"strconv"
)

func VerifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(config.Config[config.JWT_KEY])
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func JWTAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", config.Config[config.CLIENT_URL])
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")
			return
		}

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			utils.NewErrorResponse(w, http.StatusUnauthorized, "Authentication failure")
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			utils.NewErrorResponse(w, http.StatusUnauthorized, "Error verifying JWT token: " + err.Error())
			return
		}

		//pass userId claim to req
		//todo: find a better way to convert the claim to string
		userId := strconv.FormatFloat(claims.(jwt.MapClaims)["user_id"].(float64), 'g', 1, 64)
		r.Header.Set("userID", userId)
		
		next.ServeHTTP(w, r)
	})
}