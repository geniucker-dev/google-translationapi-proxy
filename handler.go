package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Handler(c *gin.Context) {
	url := "https://translate.googleapis.com"

	path := c.Param("path")
	params := c.Request.URL.Query()
	println("path:", path)

	// create a new request
	req, err := http.NewRequest(c.Request.Method, url+path, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// copy query params
	q := req.URL.Query()
	for k, v := range params {
		for _, vv := range v {
			q.Add(k, vv)
		}
	}
	req.URL.RawQuery = q.Encode()

	// copy headers
	for k, v := range c.Request.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}

	// send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// copy headers
	for k, v := range resp.Header {
		for _, vv := range v {
			c.Writer.Header().Add(k, vv)
		}
	}

	// copy status code
	c.Status(resp.StatusCode)

	// copy body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Writer.Write(body)
}
