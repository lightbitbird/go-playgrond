package main

import (
    "net/http"
    "runtime"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/labstack/gommon/log"
)

var eLog echo.Logger

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
    Name  string `json:"name"`
    Admin bool   `json:"admin"`
    jwt.StandardClaims
}

func jwt_login(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    if username == "jon" && password == "shhh!" {

        // Set custom claims
        claims := &jwtCustomClaims{
            "Jon Snow",
            true,
            jwt.StandardClaims{
                ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
            },
        }

        // Create token with claims
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

        // Generate encoded token and send it as response.
        t, err := token.SignedString([]byte("secret"))
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, echo.Map{
            "token": t,
        })
    }

    return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {
    var jsn log.JSON
    jsn = make(log.JSON)
    jsn["message"] = "log test"
    if pc, file, line, ok := runtime.Caller(1); ok {
        jsn["pc"] = pc
        jsn["last_file"] = file
        jsn["last_row"] = line
    }
    eLog.Debugj(jsn)
    return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(*jwtCustomClaims)
    name := claims.Name
    return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
    e := echo.New()
    eLog = e.Logger
    eLog.SetLevel(log.DEBUG)

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Login route
    e.POST("/login", jwt_login)

    // Unauthenticated route
    e.GET("/", accessible)

    // Restricted group
    r := e.Group("/restricted")

    // Configure middleware with the custom claims type
    config := middleware.JWTConfig{
        Claims:     &jwtCustomClaims{},
        SigningKey: []byte("secret"),
    }
    r.Use(middleware.JWTWithConfig(config))
    r.GET("", restricted)

    e.Logger.Fatal(e.Start(":1323"))
}
