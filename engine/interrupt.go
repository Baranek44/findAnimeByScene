package engine

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/muesli/termenv"
)

func Interrupt(s *spinner.Spinner) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT)
	<-channel

	if s.Active() {
		s.FinalMSG = color.HiGreenString("Arrivederci\n")
		s.Stop()
	}

	termenv.ShowCursor()
	os.Exit(0)
}
