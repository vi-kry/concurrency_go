package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vi-kry/concurrency_go/internal/compute"
	"github.com/vi-kry/concurrency_go/internal/database"
	"github.com/vi-kry/concurrency_go/internal/storage"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Failed to create logger")
		os.Exit(1)
	}
	defer func() {
		_ = logger.Sync()
	}()

	queryParser := compute.NewParser(logger)
	storageEngine := storage.NewEngine(logger)

	computerLayer := compute.NewCompute(queryParser, logger)
	storageLayer := storage.NewStorage(storageEngine, logger)

	db := database.NewDatabase(computerLayer, storageLayer, logger)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		query := strings.TrimSpace(scanner.Text())
		if query == "" {
			fmt.Print("> ")

			continue
		}

		if up := strings.ToUpper(query); up == "EXIT" || up == "QUIT" {
			fmt.Println("Shutting down database...")

			break
		}

		result, dbErr := db.HandleQuery(query)
		if dbErr != nil {
			fmt.Printf("Error: %v\n", dbErr)
			fmt.Print("> ")

			continue
		}

		fmt.Printf("Result: %s\n", result)
		fmt.Print("> ")
	}
}
