package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/teamscanworks/codebreaker/server"
)

const (
	defaultUpdateFreq = "@daily"
)

func main() {
	listenAddr, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	registryUrl := os.Getenv("REGISTRY_URL")

	if registryUrl == "" {
		log.Fatal(errors.New("Error loading REGISTRY_URL env variable \n"))
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	err = server.Serve(ctx, registryUrl, listenAddr, defaultUpdateFreq)
	if err != nil {
		fmt.Print(err)
	}
}

func parseArgs() (string, error) {
	switch len(os.Args) {
	case 2:
		return os.Args[1], nil
	case 1:
		return ":8080", nil
	default:
		return "", errors.New("expected 0 or 1 argument. \n\nUsage: codebreaker [listen-addr]")
	}
}
