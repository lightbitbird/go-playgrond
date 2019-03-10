package main

import (
    "fmt"
    "github.com/labstack/echo/middleware"
    "io/ioutil"
    "net/http"
    "os"
    "runtime"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo"
    "github.com/labstack/gommon/log"
)

func main() {
    e := echo.New()
    eLog = e.Logger
    eLog.SetLevel(log.DEBUG)

    rotation = logRotation{complete: true}
    // rotation.lastTime = time.Date(2019, 3, 10, 0, 0, 0, 0, time.Local)

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

var eLog echo.Logger
var rotation logRotation

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
    Name  string `json:"name"`
    Admin bool   `json:"admin"`
    jwt.StandardClaims
}

type logRotation struct {
    lastTime time.Time
    complete bool
}

func jwt_login(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    // 期間以前のログファイル削除
    removeLogFiles(&rotation)

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

        time.Sleep(1000 * time.Millisecond)

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

func removeLogFiles(rotation *logRotation) {
    go func() {

        fmt.Println("LogRotation = ", rotation.lastTime, rotation.complete)
        today, _ := midnightTime(time.Now())

        fmt.Println("today = ", *today, today)
        if rotation.lastTime.IsZero() || rotation.lastTime.Before(*today) {
            fmt.Println("lastTime = ", rotation.lastTime, rotation.lastTime.Before(*today))
            rotation.complete = false
        }
        if !rotation.complete {
            if files, err := ioutil.ReadDir("./logs"); err == nil {
                if files == nil || len(files) == 0 {
                    rotation.lastTime = *today
                    rotation.complete = true
                }
                var i = 0
                maxAge := today.AddDate(0, 0, -10)
                var isCompleted = true
                for _, file := range files {
                    t := file.ModTime()
                    mt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
                    // mt := time.Date(t.Year(), t.Month(), t.Day() - i, 0, 0, 0, 0, time.Local)
                    fmt.Println("mt = ", file.ModTime(), mt, maxAge)
                    if maxAge.After(mt) {
                        fmt.Println("file", file.Name(), maxAge, mt)
                        if err := os.Remove(fmt.Sprintf("./logs/%s", file.Name())); err != nil {
                            fmt.Println("Error occured")
                            isCompleted = false
                        } else {
                            fmt.Println("file.ModTime = ", file.ModTime())
                        }
                    }
                    i++
                }
                if isCompleted {
                    rotation.lastTime = *today
                    rotation.complete = true
                    fmt.Println("complete!!", rotation.lastTime, rotation.complete)
                }
            }
        }

    }()
}

func midnightTime(target time.Time) (*time.Time, error) {
    midnight := time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, time.Local)
    return &midnight, nil
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
