package middleware

import (
	"backend/internal/config"
	logging "backend/internal/logger"
	"backend/internal/utils"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		//Check if the user is authenticated
		// Grab the auth token from the request header
		authorizationTokens := req.Cookies()
		var authToken string
		var refreshToken string

		// Grab the auth token from cookies

		for _, cookie := range authorizationTokens {
			if cookie.Name == "auth_token" {
				authToken = cookie.Value
			} else if cookie.Name == "refresh_token" {
				refreshToken = cookie.Value
			}
			break
		}

		userCreds, err := utils.DecodeJwt(authToken, &utils.AuthTokenClaims{}, config.Env.AuthToken)
		if err != nil {
			return
		}

		refreshCreds, err := utils.DecodeJwt(refreshToken, &utils.RefreshTokenClaims{}, config.Env.RefreshToken)
		if err != nil {
			err := utils.Encode(res, req, http.StatusUnauthorized, "Unauthorized")
			if err != nil {
				return
			}
		}

		logging.Logger.Info().Msg("User is authenticated")
		logging.Logger.LogInfo().Msgf("User is authenticated with email: %s", userCreds.UserID)
		logging.Logger.LogInfo().Msgf("User is authenticated with email: %s", refreshCreds.UserID)

		//If not, return an error
		
		//If the user is authenticated, call the next handler
		next.ServeHTTP(res, req)
	}
}
