package endpoints

import (
	"regexp"
	"time"

	userDb "github.com/Oscar170/i-rol-back/user-api/db/user"
	"github.com/Oscar170/i-rol-back/user-api/models"
	b64 "github.com/Oscar170/i-rol-back/utils/base64"
	"github.com/Oscar170/i-rol-back/utils/http"
	jwt "github.com/dgrijalva/jwt-go"
)

const jwtSecret = "myfancysecret"

var emailRegExp, _ = regexp.Compile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}`)

// SignIn Check the credentials of the user and if it's valid return JWT
func SignIn(urlParams http.UrlParam, queryParams http.QueryParams, body http.GetBody) interface{} {
	credentials := &models.Credentials{}
	user := &models.User{}

	body(credentials)

	username, uErr := b64.Decode(credentials.Username)
	password, pErr := b64.Decode(credentials.Password)

	if uErr != nil || pErr != nil {
		return models.Status{
			Success: false,
		}
	}

	credentials.Username = username
	credentials.Password = password

	err := userDb.Restore(user, credentials)

	if err != nil {
		return models.Status{
			Success: false,
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["_id"] = user.ID
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecret))

	return &models.Authorization{
		Token: t,
	}
}

// SignUp check if the user is valid. If its valid it is create in the bd.
func SignUp(urlParams http.UrlParam, queryParams http.QueryParams, body http.GetBody) interface{} {
	user := &models.User{}

	body(user)

	emptyFields := user.Email == "" ||
		user.Name == "" ||
		user.Password == ""

	if emptyFields {
		return models.Status{
			Success: false,
		}
	}

	if !emailRegExp.MatchString(user.Email) {
		return models.Status{
			Success: false,
		}
	}

	count, err := userDb.IsValid(*user)

	if err != nil || count > 0 {
		return models.Status{
			Success: false,
		}
	}

	userDb.Store(user)

	return models.Status{
		Success: true,
	}
}
