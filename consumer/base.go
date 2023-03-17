package consumer

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/eminoz/graceful/broker"
	"github.com/eminoz/graceful/model"
	"github.com/eminoz/graceful/service"
)

func RunConsumer() {
	f := createGracefulShutdown()
	u := make(chan model.User, 10)
	userService := service.UserService{}

	ctx, cancel := context.WithCancel(context.Background())
	go broker.User(ctx, u)
	go userService.UserProcess(u)
	<-f
	cancel()
}

func createGracefulShutdown() chan os.Signal {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM)
	signal.Notify(gracefulShutdown, os.Interrupt)
	signal.Notify(gracefulShutdown, syscall.SIGINT)
	return gracefulShutdown
}
