package main

import (
	"context"
	_case "generic-T/case"
	"os"
	"os/signal"
)

func main() {
	_case.SimpleCase()
	_case.CustNumTCase()
	_case.BuiltInCase()
	_case.TTypeCase()
	_case.InterfaceCase()
	_case.ReceiverCase()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
