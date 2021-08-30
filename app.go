package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/callicoder/go-docker-compose/model"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome! Hello VTTC. direc to '/quote' get get the quote"))
}

func quoteOfTheDayHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		date := currentTime.Format("21-07-2021")
		log.Println("Cache miss for date ", date)
		quoteResp, err := getQuoteFromAPI()
		if err != nil {
			w.Write([]byte("Sorry! We could not get the Quote of the Day. Please try again."))
			return
		}
		quote := quoteResp.Contents.Quotes[0].Quote
		w.Write([]byte(quote))
	}
}

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/quote", quoteOfTheDayHandler())

	srv := &http.Server{
		Handler:      r,
		Addr:         ":3000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server at :" + srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

func getQuoteFromAPI() (*model.QuoteResponse, error) {
	API_URL := "http://quotes.rest/qod.json"
	resp, err := http.Get(API_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Println("Quote API Returned: ", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		quoteResp := &model.QuoteResponse{}
		json.NewDecoder(resp.Body).Decode(quoteResp)
		return quoteResp, nil
	} else {
		return nil, errors.New("Could not get quote from API")
	}

}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
