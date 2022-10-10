package main

import (
	"ala_web/web"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Llongfile)
	handler := &web.Handler(
		Logger: logger,
	)

	srv := &http.Server(
		handler: handler,
		Addr: ":8080"
	)

	return nil
}
