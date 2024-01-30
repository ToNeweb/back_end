package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

func UserExpireController(w http.ResponseWriter, r *http.Request) { /// currently not useable
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
}

func UserValidationSendController(w http.ResponseWriter, r *http.Request) {
	var emailToValidate struct {
		Email string `json:"email,omitempty"`
	}
	err := json.NewDecoder(r.Body).Decode(&emailToValidate)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}
	seed := rand.NewSource(time.Now().UnixNano())
	randomNum := rand.New(seed).Intn(90000) + 10000

	fmt.Println(emailToValidate.Email, ": ", randomNum)
	/// sendEmail(emailToValidate.Email, randomNumber)
	service.NewUserOps(r.Context()).UserAddValidateCode(emailToValidate.Email, randomNum)
	utils.Return(w, true, http.StatusOK, nil, emailToValidate)
}

func UserValidationCheckController(w http.ResponseWriter, r *http.Request) {
	var emailToValidateWithCode struct {
		Email string `json:"email,omitempty"`
		Code  int    `json:"code,omitempty"`
	}
	err := json.NewDecoder(r.Body).Decode(&emailToValidateWithCode)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}
	savedCode := service.NewUserOps(r.Context()).UserGetValidateCode(emailToValidateWithCode.Email) //get error too, if wasn't in here; generate another
	if emailToValidateWithCode.Code == savedCode {
		fmt.Println("validated")
		/// here check if this account already exists, dont create it again
		
		user, err := service.NewUserOps(r.Context()).UserGetByEmail(emailToValidateWithCode.Email)
		if err != nil { // check for only "account does Not Exist!"
			/// giva back a jwt
			user, _ = service.NewUserOps(r.Context()).UserCreateWithValidationEmail(emailToValidateWithCode.Email)
		}
		jwtLog, _ := generateJWT(user.ID)
		utils.Return(w, true, http.StatusOK, nil, jwtLog)

	} else {
		fmt.Println("no", savedCode, " ", emailToValidateWithCode.Code)
	}

}

/// add setting password for this email that created in UserValidationCheckController
