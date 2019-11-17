package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo"
	"crypto/rand"
	"math/big"
	"unicode"
)

//var record map[string]string

type Payload struct {
	ShortUrl string
	Message string
}

var db Storage
var conf Configuration

func main ()  {
	conf = GetConfiguration()
	e := echo.New()
	s, err := NewStorage()
	if err != nil {
		e.Logger.Fatal("Can NOT connect to DB")
	}
	db = s
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Load Balancer!")
	})
	e.GET("/:key", redirect)
	e.POST("/shorten", shorten)
	e.Logger.Fatal(e.Start(":8080"))
	db.Close()
}

func redirect(c echo.Context) error {
	key := c.Param("key")
	fmt.Println("[redirect] Receive Request for Shorten URL: ", key)
	origin, err := db.GetEntry(key)
	if err != nil {
		fmt.Println("[redirect] GetEntry failed: ", err)
		return c.String(404, "URL NOT FOUND")
	}
	fmt.Println("[redirect] Origin URL: ", origin)
	return c.Redirect(http.StatusMovedPermanently, origin)
}

func generateRandomString(length int) (string, error) {
	var result string
	for len(result) < length {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		if unicode.IsLetter(rune(n)) {
			result += string(n)
		}
	}
	return result, nil
}

func shorten(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	originUrl := m["url"].(string)
	fmt.Println("[shorten] Receive request for URL: ", originUrl)
	var randomStr string = ""
	var err error = nil
	for i := 1; i <= 10; i++ {
		if randomStr, err = generateRandomString(7); err != nil {
			fmt.Println("[shorten] Generate Random String failed")
			err = nil
			continue
		}
		if err = db.CreateEntry(randomStr, originUrl); err != nil {
			fmt.Println("[shorten] CreateEntry failed: ", err)
			err = nil
			continue
		}
		break;
	}
	
	if err != nil {
		payload := &Payload {
			ShortUrl: "",
			Message: "Can NOT shorten the URL",
		}
		return c.JSON(500, payload)
	}

	shortenedUrl := conf.BaseURL + randomStr
	fmt.Println("[shorten] Shortened URL: ", shortenedUrl)
	//fmt.Println(record[shortenedUrl])
	payload := &Payload {
		ShortUrl: shortenedUrl,
		Message: "",
	}
	return c.JSON(200, payload)
}