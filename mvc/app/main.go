package main

import (
	v "app/package/view"
	"encoding/json"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
	"crypto/sha512"
	"errors"
	"strings"
	"strconv"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Content struct {
	Id int `json:"id"`
	Content string `json:"content"`
}

func main() {
	v.PrintHomeStatement()
	rand.Seed(time.Now().Unix())
	hash, err := getRandomString(20)
	if err != nil {
		fmt.Println("Error genarating password -> [getRandomString]")
	}
	fmt.Println("Random String here: ", hash)
	fmt.Println(GetHashedString(hash))
}

func GetHashedString(stringToHash string) string {
	b := sha512.New()
	b.Write([]byte(stringToHash))
	bytes := b.Sum(nil)
	return	hex.EncodeToString(bytes)
}

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	res := map[string]string {
		"CODE": "200",
		"STATUS": "OK",
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(jsonData)
}

// Variadic function in Golang
func getVariadicFunc(str ...string) {
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
}

func getBase10Int() (int) {
	var base10 = []rune("0123456789")
	now := time.Now()
	getRand, err := strconv.Atoi(string(base10[rand.Intn(int(now.Second() * now.Minute()) % len(base10))]))
	if err != nil {
		fmt.Println("Error at *[getBase10int]")
	}
	return getRand
}

func getRandomChar() (string) {
	now := time.Now()
	var runes = []rune("abcdefghijlklmnopqstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ")
	randChar := runes[rand.Intn(int(now.Second() * now.Minute()) % len(runes))]
	return string(randChar)
}

func getRandomSymbol() (string) {
	now := time.Now()
	var runes = []rune("!@#$%^&*():?>,.;[]{}|<")
	randSymbol := runes[rand.Intn(int(now.Second() * now.Minute()) % len(runes))]
	return string(randSymbol)
}

func getRandomString(length int) (string, error) {
	var stringsArray []string

	if length < 12 {
		return "", errors.New("password must at least 12 character")
	}
	if length > 20 {
		return "", errors.New("password length too long")
	}
	for i := 0; i < length; i++ {
		if i >= 7 {
			stringsArray = append(stringsArray, getRandomChar())
		} else if i >= 5 {
			stringsArray = append(stringsArray, getRandomSymbol())
		} else {
			stringsArray = append(stringsArray, strconv.Itoa(getBase10Int()))
		}
	}
	return strings.Join(stringsArray, ""), nil
}
