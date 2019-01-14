package main

import (
    "log"
    jwt "github.com/dgrijalva/jwt-go"
)

func main() {

    // Decode with StandardClaims struct
    tokenstring := createTokenString()
    log.Println(tokenstring)

    // Decode with StandardClaims struct
    token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
        return []byte("foobar"), nil
    })
    log.Println(token.Claims, err)

    user := User{}
    // Decode with StandardClaims struct
    token, err = jwt.ParseWithClaims(tokenstring, &user, func(token *jwt.Token) (interface{}, error) {
        return []byte("foobar"), nil
    })
    log.Println(token.Valid, user, err)
}

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"'`
    jwt.StandardClaims
}

func createTokenString() string {
    token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
        Name: "max10",
        Age:  30,
    })

    tokenstring, err := token.SignedString([]byte("foobar"))
    if err != nil {
        log.Fatalln(err)
    }
    return tokenstring
}
