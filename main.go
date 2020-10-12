package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// func main() {

// 	token := createToken()
// 	createSession(token)
// }

func createToken() string {
	secret := ""
	mySigningKey := []byte(secret)
	apiKey := ""
	//currentTime := time.Now().Unix()
	//expiryTime:= time.Until().Unix()

	// Create the Claims
	claims := &jwt.StandardClaims{
		Issuer:    apiKey,
		IssuedAt:  1602486301,
		ExpiresAt: 1602521084,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", tokenString, err)
	return tokenString
}

func createSession(token string) {
	client := &http.Client{}
	var data = strings.NewReader(`$datastr`)
	req, err := http.NewRequest("POST", "https://api.opentok.com/session/create", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-OPENTOK-AUTH", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

}
