package consumer

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eminoz/graceful/broker"
	"github.com/eminoz/graceful/model"
	"github.com/eminoz/graceful/service"
)

func RunConsumer() {
	f := createGracefulShutdown()
	u := make(chan model.User) //if chan size is gave, chan continue to take data after shut down
	userService := service.UserService{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go broker.User(ctx, u)
	go userService.UserProcess(u)

	<-f
	cancel()
	time.Sleep(time.Second * 10)
}

func createGracefulShutdown() chan os.Signal {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM)
	signal.Notify(gracefulShutdown, os.Interrupt)
	signal.Notify(gracefulShutdown, syscall.SIGINT)
	return gracefulShutdown
}
