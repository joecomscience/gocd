package pkg

import (
	"net/http"
)

type Line struct{}

func (l *Line) send(msg string) error {
	url := ""
	if _, e := http.Get(url); e != nil {
		return e
	}

	return nil
}
