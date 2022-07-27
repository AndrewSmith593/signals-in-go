package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// In this example, you start the program, then kill it with CTRL + C,
// which sends the SIGINT signal
// which then unblocks the goroutine, and lets it send "true" to the done channel
// then we recieve from the done channel and the app closes gracefully

func main() {
	// make a channel to recieve some kind of os signal
	sigs := make(chan os.Signal, 1)

	// "register" (aka forward on some incoming signals to) the sigs channel so it receives a certain kind of signal
	// here, we forward on SIGINT and SIGKILL signals to the sigs channel
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGKILL)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		log.Printf("\nReceived a signal on sigs: %s\nWill now send to and recieve true from Done channel", sig)
		done <- true
	}()

	log.Println("Awaiting signal...")
	<-done
	log.Println("Got signal true from Done channel, shutting down...")

}
