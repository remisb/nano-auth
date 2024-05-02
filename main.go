package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/remisb/nano-auth/auth"
	"github.com/remisb/nano-auth/user"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var ErrLoginError = errors.New("auth: name, password or both are incorrect")

func main() {
	var handler = http.DefaultServeMux
	http.HandleFunc("GET /", indexHandler)
	//http.HandleFunc("POST /httpBasic", postHttpBasicHandler)
	http.HandleFunc("POST /login", postLoginHandler)
	http.HandleFunc("POST /loginJwt", loginJWTHandler)
	//http.HandleFunc("POST /loginBearer", loginBearerHandler)
	http.HandleFunc("POST /registerBearer", registerBearerHandler)
	http.HandleFunc("POST /resource", resourceHandler)
	http.HandleFunc("POST /resourceBearer", resourceBearerHandler)
	http.HandleFunc("POST /resourceJwtBearer", resourceBearerJwtHandler)
	http.HandleFunc("POST /resetPasswordBearer", resetPasswordHandler)
	const port = ":8000"
	http.ListenAndServe(port, handler)
}

type userReset struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type userLogin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}

type userRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}

func extractUserReset(r *http.Request) (userReset, error) {
	var uReset userReset
	err := json.NewDecoder(r.Body).Decode(&uReset)
	if err != nil {
		return uReset, err
	}

	return uReset, nil
}

func extractUserLogin(r *http.Request) (userLogin, error) {
	var uLogin userLogin
	err := json.NewDecoder(r.Body).Decode(&uLogin)
	if err != nil {
		return uLogin, err
	}

	return uLogin, nil
}

func extractBearerToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		return "", errors.New("Missing bearer token")
	}

	reqToken := strings.Split(bearerToken, " ")[1]
	return reqToken, nil
}

func extractUserRegister(r *http.Request) (userRegister, error) {
	var uRegister userRegister
	err := json.NewDecoder(r.Body).Decode(&uRegister)
	if err != nil {
		return uRegister, err
	}

	return uRegister, nil
}

func resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	userReset, err := extractUserReset(r)
	if err != nil {
		log.Error(err) //logs will provide for devops to
		// probably traceId should be added to all rest API error in the future
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := user.ByNameOrEmail(userReset.Name, userReset.Email)
	if user == nil {
		// REMIS QUESTION Should log output name or password for Sec Penetration prevention
		// maybe remote ip and other relevant information should be collected and provided
		// to Security team for prevention analysis.
		// This is definitely not a system error and could be logged out for security analysis only
		log.Warning("auth: attempt to reset password for unknown user")
		return
	}

	err = auth.ResetPassword(user.Email)
	if err != nil {
		JsonErrorResponse(w, err)
		return
	}

	// REMIS QUESTION SEC Provided information could be used for security penetration.
	// Provided information reveals username or email being a valid system user.
	SendJson(w, H{"message": "auth: userReset message was sent to user's mailbox"}, http.StatusOK)
}

func registerBearerHandler(w http.ResponseWriter, r *http.Request) {
	userRegister, err := extractUserRegister(r)
	if err != nil {
		log.Info(err) // probably invalid data format was provided
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userA := user.ByNameOrEmail(userRegister.Name, userRegister.Email)
	if err != nil {
		// Unknown username or email is provided.
		// Register user workflow hs to be canceled.
		log.Info(err)
		http.Error(w, "auth: user with such name/email already exist, please use login", http.StatusBadRequest)
		return
	}

	valid, err := user.ValidateUnencodedUser(userA.Name, userRegister.Password)
	if err != nil {
		log.Info(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !valid {
		SendJsonMsg(w, http.StatusBadRequest, "Invalid user credentials")
		return
	}

	// Create a new User object
	newUser := user.New(userRegister.Name, userRegister.Email, userRegister.Password)
	err = user.Add(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SendJsonMsg(w, http.StatusOK, "User registered successfully")
}

func resourceBearerHandler(w http.ResponseWriter, r *http.Request) {
	bearerToken, err := extractBearerToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	if !auth.ValidateToken(bearerToken) {
		SendJson(w, H{"message": "unauthorized"}, http.StatusUnauthorized)
		return
	}

	data := H{
		"data": "resource data",
	}
	SendJson(w, data, http.StatusOK)
}

func resourceBearerJwtHandler(w http.ResponseWriter, r *http.Request) {
	bearerToken, err := extractBearerToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	auth.JwtVerify(bearerToken)
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !auth.ValidateToken(bearerToken) {
		SendJson(w, H{"message": "unauthorized"}, http.StatusUnauthorized)
		return
	}

	data := H{
		"data": "resource data",
	}
	SendJson(w, data, http.StatusOK)
}

//func loginBearerHandler(w http.ResponseWriter, r *http.Request) {
//	userLogin, err := extractUserLogin(r)
//	if err != nil {
//		JsonErrorResponse(w, ErrLoginError)
//		return
//	}
//
//	user := user.ByNameOrEmail(userLogin.Name, userLogin.Password)
//	if user == nil {
//		SendJsonMsg(w, http.StatusBadRequest, "auth: Invalid user credentials")
//		return
//	}
//
//	encodedPassword := auth.EncodePassword(userLogin.Password)
//	if user.PassEncoded == userLogin.Password {
//
//	}
//	token, _ := auth.randomHex(20)
//	auth.tokens = append(auth.tokens, token)
//
//	res := struct {
//		Token string
//	}{
//		Token: token,
//	}
//	SendJson(w, res, http.StatusOK)
//	return
//}

func postLoginHandler(w http.ResponseWriter, r *http.Request) {
	//rUser, rPass, ok := r.BasicAuth()
	rUser, _, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user := user.ByName(rUser)
	if user == nil {
		http.Error(w, "auth: ", http.StatusUnauthorized)
		return
	}

	token, _ := auth.RandomHex(20)
	auth.AddToken(token)

	SendJson(w, H{
		"token": token,
	}, http.StatusOK)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {

}

func loginJWTHandler(w http.ResponseWriter, r *http.Request) {
	type userLogin struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"pass"`
	}

	var uLogin userLogin
	err := json.NewDecoder(r.Body).Decode(&uLogin)
	if err != nil {
		SendJsonMsg(w, http.StatusBadRequest, "Invalid user credentials")
		return
	}

	// does the username already used?
	// No connection to mongoDB. Local db connection for development should be used and mongodb should be started.
	// TODO Local docker containers started but dev env settings to connect to local db is not active.

	foundUser := user.ByNameOrEmail(uLogin.Name, uLogin.Email)
	if foundUser == nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		SendJsonMsg(w, http.StatusBadRequest, "Invalid user credentials")
		return
	}

	uName := uLogin.Name
	if len(uName) == 0 {
		uName = uLogin.Email
	}

	token, err := auth.NewSignedToken(uName)
	if err != nil {
		log.Error(err)
		SendJson(w, H{"message": "auth: failed to generate token"}, http.StatusBadRequest)
		return
	}

	auth.AddToken(token)
	SendJson(w, H{"token": token}, http.StatusOK)
}
