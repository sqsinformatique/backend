package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sqsinformatique/backend/srv"
	"github.com/sqsinformatique/backend/utils"
)

var closeSignal chan os.Signal

// Exit handler
func exitLoop() {
	exit := make(chan struct{})
	closeSignal = make(chan os.Signal, 1)
	signal.Notify(closeSignal, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-closeSignal
		srv.Stop()
		utils.Infoln("Exit program")
		close(exit)
	}()

	// Exit app if chan is closed
	<-exit
}
