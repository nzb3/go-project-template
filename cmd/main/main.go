package main

import (
	"context"
	"log"

	"github.com/nzb3/go-project-template/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err.Error())
	}

	err = a.Start(ctx)
	if err != nil {
		log.Fatalf("failed to start app: %v", err.Error())
	}
}
