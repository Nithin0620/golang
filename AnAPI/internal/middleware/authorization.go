package middleware

import (
	"errors"
	"net/http"

	"github.com/Nithin0620/golang/api"
	"github.com/Nithin0620/golang/internal/tools"
	log "github.com/sirupsen/logrus"
)

var unAuthorizedError = errors.New("Invalid Username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" {
			api.RequestErrorHandler(w, errors.New("username is required"))
			return
		}

		var database tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = database.GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
