package service

import (
	"os"
	"os/signal"
	"syscall"
)

type Runner interface {
	Start() error
	Stop() error
	Log(msg string)
}

func Run(name string, r Runner) error {
	if CurrentRunMode() == RunModeNone {
		return runForeground(r)
	}
	return runService(name, r)
}

func runForeground(r Runner) error {
	if err := r.Start(); err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGTERM, os.Interrupt)
	<-sig
	return r.Stop()
}
