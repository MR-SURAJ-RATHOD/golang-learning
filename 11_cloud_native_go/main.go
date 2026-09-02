package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Lesson: 12-Factor App — config from environment
type Config struct {
	AppName string
	Port    string
	Region  string
}

func loadConfig() Config {
	return Config{
		AppName: getEnv("APP_NAME", "cloud-native-go"),
		Port:    getEnv("PORT", "8080"),
		Region:  getEnv("AWS_REGION", "us-east-1"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	fmt.Println("=== 11 Cloud Native Go ===")

	cfg := loadConfig()
	fmt.Printf("\n--- Lesson 1: 12-Factor Config ---\n")
	fmt.Printf("  App: %s | Port: %s | Region: %s\n", cfg.AppName, cfg.Port, cfg.Region)

	fmt.Println("\n--- Lesson 2: Graceful Shutdown ---")
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		fmt.Println("  Shutdown signal received — draining connections...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  Server stopped gracefully.")
		os.Exit(0)
	}()

	fmt.Println("\n--- Lesson 3: AWS S3 (optional) ---")
	runS3Demo()

	fmt.Println("\nPress Ctrl+C to test graceful shutdown...")
	<-ctx.Done()
}

func runS3Demo() {
	bucket := os.Getenv("BUCKET_NAME")
	if bucket == "" {
		fmt.Println("  ⚠️  Set BUCKET_NAME to test real S3. Skipping.")
		fmt.Println("  Example: set BUCKET_NAME=my-bucket && go run main.go")
		return
	}

	// AWS SDK demo — requires credentials in ~/.aws/credentials or env
	fmt.Printf("  Would list objects in bucket: %s\n", bucket)
	fmt.Println("  (Full S3 code in docs — requires AWS credentials)")
	_ = log.Default()
}
