package cmd

import (
	"context"
	"echo-wedge/backend/config"
	er "echo-wedge/backend/router/echo"
	serv "echo-wedge/backend/services/tcpServer"
	cl "echo-wedge/backend/tcpClient"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tasks := []func() error{
		setLogLevel,
		printStartMessage,
		setupServer,
		setupClient,
		setupAPI,
	}

	for _, t := range tasks {
		if err := t(); err != nil {
			log.Fatal(err)
		}
	}

	sigChan := make(chan os.Signal)
	exitChan := make(chan struct{})
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.WithField("signal", <-sigChan).Info("signal received")
	go func() {
		log.Warning("Stopping Echo-wedge rest service!")
		exitChan <- struct{}{}
	}()
	select {
	case <-exitChan:
	case s := <-sigChan:
		log.WithField("signal", s).Info("signal received, stopping immediately")
	}

	return nil
}

func setLogLevel() error {
	log.SetLevel(log.Level(uint8(config.C.General.LogLevel)))
	return nil
}

func printStartMessage() error {
	log.WithFields(log.Fields{
		"version": version,
		"docs":    "https://www.seluxit.com/docs",
	}).Info("Rest service for gateway. ")
	return nil
}

func setupServer() error {
	if err := serv.Setup(config.C); err != nil {
		return errors.Wrap(err, "setup tcp server error")
	}
	return nil
}

func setupClient() error {
	if err := cl.Setup(config.C); err != nil {
		return errors.Wrap(err, "setup tcp gateway client error")
	}
	return nil
}

func setupAPI() error {
	if err := er.RunApp(config.C); err != nil {
		return errors.Wrap(err, "setup Rest service api error")
	}
	return nil
}
