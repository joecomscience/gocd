package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/alertmanager/template"
)

const (
	line_uri = "https://notify-api.line.me/api/notify"
)

func main() {
	port := ":" + os.Getenv("PORT")
	// if port == "" {
	// 	port = "80"
	// }

	router := http.NewServeMux()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.HandleFunc("/hook", hookHandler)

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("server start using port: %s", port)

	<-done
	log.Println("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	promData := template.Data{}
	bearer := "Bearer " + os.Getenv("LINE_TOKEN")
	messageAlert := ""

	if err := json.NewDecoder(r.Body).Decode(&promData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, data := range promData.Alerts {
		messageAlert = data.Annotations["description"]
	}

	client := &http.Client{}
	body := url.Values{}
	body.Set("message", messageAlert)

	req, _ := http.NewRequest("POST", line_uri, bytes.NewBufferString(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", bearer)

	result, err := client.Do(req)
	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Send Line notify error"))
		return
	}
	defer result.Body.Close()
	w.WriteHeader(http.StatusOK)
}
