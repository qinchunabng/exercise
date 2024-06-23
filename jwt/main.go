package main

import (
	"encoding/base64"
	"fmt"
	"jwt/jwtV1"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Name   string
	Gender int
	Age    int
	jwt.RegisteredClaims
}

func (mc *MyClaims) Valid() error {
	//字段校验
	//token校验逻辑
	return nil
}

func main() {
	myClaims := MyClaims{
		Name:   "nick",
		Gender: 1,
		Age:    18,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	hs := jwtV1.HS{
		Key: "abcdefg123456",
	}
	sign, err := hs.Encode(&myClaims)
	fmt.Println(sign, err)
	bytes, _ := base64.StdEncoding.DecodeString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	fmt.Println(string(bytes))
	bytes, _ = base64.StdEncoding.DecodeString("eyJOYW1lIjoibmljayIsIkdlbmRlciI6MSwiQWdlIjoxOCwiZXhwIjoxNzE4MjA4ODUxLCJuYmYiOjE3MTgyMDUyNTEsImlhdCI6MTcxODIwNTI1MX0")
	fmt.Println(string(bytes))
}
