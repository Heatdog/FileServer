package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Heatdog/FileServer/main_server/internal/config"
)

func main() {
	opt := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))

	logger.Info("get config")
	config, err := config.GetConfig(logger)
	if err != nil {
		logger.Error("config file was not founded", slog.String("err", err.Error()))
		panic(err)
	}

	mux := http.NewServeMux()

	host := fmt.Sprintf("%s:%s", config.NetworkConfig.IP, config.NetworkConfig.Port)
	logger.Info("listen on", slog.String("host", host))

	if err := http.ListenAndServe(host, mux); err != nil {
		logger.Error("server listen failed", slog.String("err", err.Error()))
		panic(err)
	}
}
