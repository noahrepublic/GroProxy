package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/h2non/gentleman.v2"
)

var errInvalidAuth error = errors.New("invalid API key")
var errInvalidURL error = errors.New("invalid URL")
var errFailedConnection error = errors.New("failed to connect to server")

var accounts = make(map[string]string)

func processRequest(c *gin.Context) {
	if accounts["API_KEY"] != "" {
		if key := c.GetHeader("API_KEY"); key != "" {

			if key != accounts["API_KEY"] {
				c.Error(errInvalidAuth)
				c.Status(http.StatusUnauthorized)
				return
			}
		} else {
			c.Error(errInvalidAuth)
			c.Status(http.StatusUnauthorized)
			return
		}
	}

	req := c.Request

	url := strings.SplitN(req.RequestURI[1:], "/", 2)

	if len(url) < 2 {
		c.Error(errInvalidURL)
		c.Status(http.StatusBadRequest)
		return
	}

	reqUrl, err := req.URL.Parse("https://" + url[0] + ".roblox.com/" + url[1])

	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	req.URL = reqUrl
	requestServer(req, c)
}

func requestServer(req *http.Request, c *gin.Context) {

	client := gentleman.New()
	client.Method(req.Method)
	client.URL(req.URL.String())

	res, err := client.Request().Send()

	if err != nil {
		c.Error(errFailedConnection)
		c.Status(http.StatusBadGateway)
		return
	}

	body := res.Context.Request.Body

	for key, value := range res.Header {
		for _, v := range value {
			c.Header(key, v)
		}
	}

	defer body.Close()

	// write the body into the response

	c.Writer.Write(res.Bytes())
	c.Status(res.StatusCode)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	err := godotenv.Load()

	if err == nil {
		apiKey, ok := os.LookupEnv("API_KEY")

		if ok {
			accounts["API_KEY"] = apiKey
			fmt.Println("API Key: " + apiKey)
		}
	}

	router.GET("*subdomain", processRequest)

	if err := router.Run(); err != nil {
		panic(err)
	}

	fmt.Println("Server started")

}
