package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/BradPreston/blog-backend/pkg/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func ResponseJSON(w http.ResponseWriter, data interface{}, statusMessaage string, statusCode int) {
	response := make(map[string]interface{})
	response["status"] = statusMessaage
	response["data"] = data
	resJSON, _ := json.MarshalIndent(response, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resJSON)
}

func ENV() (map[string]string, error) {
	env, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}
	return env, nil
}

func CreateJWTString(user api.User) (string, error) {
	env, err := ENV()
	if err != nil {
		log.Println(err)
	}

	var jwtKey = []byte(env["JWT_SECRET"])

	type Claims struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}

	// set the expiration time to 1 day
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			// Set the time to Unix for JWT
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// sign the token and add claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// create the token string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
