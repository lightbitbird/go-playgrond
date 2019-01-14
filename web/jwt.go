package main

import (
    "github.com/labstack/echo"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"
    "time"
    "fmt"
)

func main() {
    e := echo.New()
    route(e)
    e.Logger.Fatal(e.Start("localhost:1323"))
}

func route(e *echo.Echo) {
    secretKey := "abcdefg12345"

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    e.GET("/jwt/signature", func(c echo.Context) (err error) {
        u := new(JwtUser)
        if err = c.Bind(u); err != nil {
            return
        }
        if u.Name == "" {
            return c.JSON(400, map[string]string{"message":   "Name is required"})
        }
        fmt.Println("user -> ", u)

        // set an algorithm
        token := jwt.New(jwt.GetSigningMethod("HS256"))
        token.Claims = jwt.MapClaims{
            "user": u.Name,
            "exp": time.Now().Add(time.Hour * 1).Unix(),
        }

        // Add a signature to token
        tokenString, err := token.SignedString([]byte(secretKey))
        if err != nil {
            return c.JSON(500, map[string]string{"message":   "Could not generate token"})
        }
        return c.JSON(200, tokenString)
    })

    e.GET("/jwt/private", func(c echo.Context) error {
        // inspection the signature
        token, err := request.ParseFromRequest(c.Request(), request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
            b := []byte(secretKey)
            return b, nil
        })
        if err == nil {
            claims := token.Claims.(jwt.MapClaims)
            msg := fmt.Sprintf("Hello, %s", claims["user"])
            return c.JSON(200, map[string]string{"message": msg})
        } else {
            return c.JSON(401, map[string]string{"error": fmt.Sprint(err)})
        }
    })
}

type JwtUser struct {
    Name string `json:"name" query:"name"`
    Email string `json:"email" query:"email"`
}
