package tracker

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/syc0x00/trakx/tracker/udp"

	"github.com/syc0x00/trakx/tracker/shared"
	"go.uber.org/zap"
)

var (
	SigStop = os.Interrupt
)

func handleSigs(peerdb *shared.PeerDatabase, udptracker *udp.UDPTracker) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGUSR1)

	for {
		sig := <-c

		switch sig {
		case os.Interrupt, os.Kill, syscall.SIGTERM:
			logger.Info("Exiting")

			peerdb.WriteFull()
			if udptracker != nil {
				udptracker.WriteConns()
			}

			os.Exit(128 + int(sig.(syscall.Signal)))
		default:
			logger.Info("Got unknown sig", zap.Any("Signal", sig))
		}
	}
}
