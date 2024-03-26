package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	fmt.Println("NTP Time:", ntpTime.Format(time.RFC3339))

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stop:
			fmt.Println("\nExiting")
			os.Exit(0)
		case <-ticker.C:
			ntpTime, err := ntp.Time("pool.ntp.org")
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			fmt.Println("NTP Time:", ntpTime.Format(time.RFC3339))
		}
	}
}
