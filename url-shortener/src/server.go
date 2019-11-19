package main

import (
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	e.Use(middleware.CORS())
	s, err := NewStorage()
	if err != nil {
		logger.Println("Can NOT connect to DB")
	}
	db = s
	e.GET("/", func(c echo.Context) error {
		logger.Println("[main] Health Check from Load Balancer")
		return c.String(http.StatusOK, "Hello, Load Balancer!")
	})
	e.GET("/unshorten/:key", redirect)
	e.POST("/shorten", shorten)
	logger.Println(e.Start(":8080"))
	db.Close()
}

func redirect(c echo.Context) error {
	key := c.Param("key")
	logger.Println("[redirect] Receive Request for Shorten URL: ", key)
	origin, err := db.GetEntry(key)
	if err != nil {
		logger.Println("[redirect] GetEntry failed: ", err)
		return c.String(404, "URL NOT FOUND")
	}
	logger.Println("[redirect] Origin URL: ", origin)
	if err = touchTopService(origin); err != nil {
		logger.Println("[redirect] touchTopService failed: ", err)
		payload := &Payload {
			ShortUrl: "",
			Message: "Can NOT redirect shortened URL",
		}
		return c.JSON(500, payload)
	}
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
	user := m["Username"].(string)
	logger.Println("[shorten] Receive request from User: ", user)
	logger.Println("[shorten] Receive request for URL: ", originUrl)
	var randomStr string = ""
	var err error = nil
	for i := 1; i <= 10; i++ {
		if randomStr, err = generateRandomString(7); err != nil {
			logger.Println("[shorten] Generate Random String failed")
			err = nil
			continue
		}
		if err = db.CreateEntry(randomStr, originUrl); err != nil {
			logger.Println("[shorten] CreateEntry failed: ", err)
			err = nil
			continue
		}
		break
	}
	
	if err != nil {
		payload := &Payload {
			ShortUrl: "",
			Message: "Can NOT shorten the URL",
		}
		return c.JSON(500, payload)
	}

	if err = touchUserService(user, originUrl); err != nil {
		logger.Println("[shorten] touchUserService failed: ", err)
		payload := &Payload {
			ShortUrl: "",
			Message: "Can NOT shorten the URL",
		}
		return c.JSON(500, payload)
	}

	shortenedUrl := conf.BaseURL + randomStr
	logger.Println("[shorten] Shortened URL: ", shortenedUrl)
	//fmt.Println(record[shortenedUrl])
	payload := &Payload {
		ShortUrl: shortenedUrl,
		Message: "",
	}
	return c.JSON(200, payload)
}

func touchUserService(username string, originUrl string) error {
	m := map[string]string{"Username" : username, "url" : originUrl}
	jsonStr, _ := json.Marshal(m)
    req, err := http.NewRequest("POST", conf.UserServiceURL, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		logger.Println("[touchUserService] Response with error: ", err)
		return err
    }
    defer resp.Body.Close()

    logger.Println("[touchUserService] Response Status: ", resp.Status)
    logger.Println("[touchUserService] Response Headers: ", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
	logger.Println("touchUserService] Response Body: ", string(body))
	return nil
}

func touchTopService(originUrl string) error {
	m := map[string]string{"url" : originUrl}
	jsonStr, _ := json.Marshal(m)
    req, err := http.NewRequest("POST", conf.TopServiceURL, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		logger.Println("[touchTopService] Response with error: ", err)
		return err
    }
    defer resp.Body.Close()

    logger.Println("[touchTopService] Response Status: ", resp.Status)
    logger.Println("[touchTopService] Response Headers: ", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
	logger.Println("touchTopService] Response Body: ", string(body))
	return nil
}