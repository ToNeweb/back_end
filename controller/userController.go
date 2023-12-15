package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server04/ent"
	"server04/utils"
	"strings"
	"time"

	"server04/service"

	"github.com/golang-jwt/jwt"
)

func UserCreateController(w http.ResponseWriter, r *http.Request) {

	var newUser ent.UserSec
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}
	newUser.Password = utils.Sha3Hash(newUser.Password)
	user, err := service.NewUserOps(r.Context()).UserCreate(newUser)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, user)
}

var secretKey = []byte("eeSecretYouShouldHide")

func generateJWT(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":     "bar",
		"user_id": userId,
		"exp":     time.Now().Add(10 * time.Minute),
	})

	// claims := token.Claims.(jwt.MapClaims)
	// claims["exp"] = time.Now().Add(10 * time.Minute)
	// claims["authorized"] = true
	// claims["user"] = "username"
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return "Signing Error", err
	}

	return tokenString, nil
}

func UserLoginController(w http.ResponseWriter, r *http.Request) {

	var newUser ent.UserSec
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}
	user, err := service.NewUserOps(r.Context()).UserGetByEmail(newUser.Email)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}
	if utils.Sha3Hash(newUser.Password) == user.Password {
		jwtLog, _ := generateJWT(user.ID)
		utils.Return(w, true, http.StatusOK, nil, jwtLog)
	} else {
		utils.Return(w, false, http.StatusNotAcceptable, err, nil)
	}
}

func UserValidateController(w http.ResponseWriter, r *http.Request) {
	//var tokenString string //:= "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJ1c2VyIjoic2VwZWhybW5wIn0.Xam-9R5pSPWhilCfVHt_pYE_WnAoeGNFvjR0bhm-rpY"
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	token, _ := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", nil)
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secretKey, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.Return(w, true, http.StatusConflict, fmt.Errorf("jwt parsed wrong"), nil)
	}
	utils.Return(w, true, http.StatusOK, nil, claims["exp"])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// claims, ok := token.Claims.(jwt.MapClaims)
	// if ok {
	// 	fmt.Println(claims["foo"])
	// } else {
	// 	fmt.Println(err, "wtf")
	// }
	// utils.Return(w, true, http.StatusOK, nil, claims["userId"])
}
