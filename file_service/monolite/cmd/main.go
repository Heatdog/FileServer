package main

import (
	"log/slog"
	"os"

	"google.golang.org/grpc"
)

func main() {
	opt := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))

	logger.Info("create a grpc connection")
	s := grpc.NewServer()
}
