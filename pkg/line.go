package pkg

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
)

type Line struct {
	Message string
}

func (l *Line) Send() error {
	uri := "https://notify-api.line.me/api/notify"
	token := ""

	c := &http.Client{}
	d := url.Values{}
	d.Set("message", l.Message)

	req, e := http.NewRequest("POST", uri, bytes.NewBufferString(d.Encode()))
	if e != nil {
		return e
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	res, e := c.Do(req)
	if e != nil {
		return e
	}

	log.Println(res)

	return nil
}
