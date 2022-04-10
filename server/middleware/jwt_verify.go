package middleware

import (
	"context"
	"fmt"
	"strings"

	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	apiHandler "payment-gateway-backend/server/handler"
	"payment-gateway-backend/usecase"
)

type jwtClaims struct {
	jwt.StandardClaims
}

// VerifyMiddlewareInit ...
type VerifyMiddlewareInit struct {
	*usecase.ContractUC
}

// VerifyPermissionInit ...
type VerifyPermissionInit struct {
	*usecase.ContractUC
	Menu string
}

func userContextInterface(ctx context.Context, req *http.Request, subject string, body map[string]interface{}) context.Context {
	return context.WithValue(ctx, subject, body)
}

// VerifyJwtTokenCredential ...
func (m VerifyMiddlewareInit) VerifyJwtTokenCredential(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := &jwtClaims{}

		tokenAuthHeader := r.Header.Get("Authorization")
		if !strings.Contains(tokenAuthHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}
		tokenAuth := strings.Replace(tokenAuthHeader, "Bearer ", "", -1)

		_, err := jwt.ParseWithClaims(tokenAuth, claims, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			secret := m.ContractUC.EnvConfig["TOKEN_SECRET"]
			return []byte(secret), nil
		})
		if err != nil {
			apiHandler.RespondWithJSON(w, 401, 401, "Invalid Token!", []map[string]interface{}{}, []map[string]interface{}{})
			return
		}

		if claims.ExpiresAt < time.Now().Unix() {
			apiHandler.RespondWithJSON(w, 401, 401, "Expired Token!", []map[string]interface{}{}, []map[string]interface{}{})
			return
		}

		// Decrypt payload
		jweRes, err := m.ContractUC.Jwe.Rollback(claims.Id)
		if err != nil {
			apiHandler.RespondWithJSON(w, 401, 401, "Error when load the payload!", []map[string]interface{}{}, []map[string]interface{}{})
			return
		}

		var deviceID string
		err = m.ContractUC.GetFromRedis("userDeviceID"+jweRes["id"].(string), &deviceID)
		if err != nil {
			apiHandler.RespondWithJSON(w, 401, 401, "Invalid Device!", []map[string]interface{}{}, []map[string]interface{}{})
			return
		}
		if deviceID != jweRes["device_id"].(string) {
			apiHandler.RespondWithJSON(w, 401, 401, "Expired Device Token!", []map[string]interface{}{}, []map[string]interface{}{})
			return
		}

		// Check id in user table
		userUc := usecase.UserUC{ContractUC: m.ContractUC}
		user, err := userUc.FindByIDWithoutTX(jweRes["id"].(string), false)
		if err != nil {
			apiHandler.RespondWithJSON(w, 401, 401, "Invalid user token!", []map[string]interface{}{}, []map[string]interface{}{})
			return
		}

		claimRes := map[string]interface{}{
			"id": user.ID,
		}

		ctx := userContextInterface(r.Context(), r, "user", claimRes)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
