package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	// VAULT login server path
	VAULTSERVER = "http://localhost:8200/v1/auth/contra/login"
)

func init() {
	newRoute := os.Getenv("VAULT_AUTH_SERVER")
	if len(newRoute) > 0 {
		VAULTSERVER = newRoute
	}
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", HelloIndex)
	router.GET("/ping", HelloWeb)
	router.POST("/ping", HelloWeb)
	router.Run()
}

//HelloWeb init point
func HelloWeb(c *gin.Context) {

	if c.Request.Method == "POST" {

		user := c.DefaultPostForm("username", "none")
		pass := c.DefaultPostForm("password", "none")

		token, err := CheckUser(user, pass)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		}

		if len(token) == 0 {
			c.JSON(400, gin.H{
				"message": "Prohibido !!!!!",
			})
			return
		}

		c.HTML(http.StatusOK, "internal.html", nil)
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func HelloIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CheckUser(user, pass string) (token string, err error) {

	// initialize http client
	client := &http.Client{}

	data := struct {
		Password string `json:"password"`
	}{
		Password: pass,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	userAuth := VAULTSERVER + "/" + user

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, userAuth, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if resp.StatusCode == 400 {
		return "", fmt.Errorf("Invalid user")
	}

	response, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	holder := struct {
		Auth struct {
			ClientToken string `json:"client_token"`
		} `json:"auth"`
	}{}

	err = json.Unmarshal(response, &holder)

	if err != nil {
		return "", err
	}

	return holder.Auth.ClientToken, nil
}
