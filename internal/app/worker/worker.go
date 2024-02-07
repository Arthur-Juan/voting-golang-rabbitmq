package worker

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

func Run(start time.Time) {
	log.Debug("Running")

	duration := time.Until(start.Add(3 * time.Hour))

	log.Debug(duration)
	if duration > 0 {
		time.Sleep(duration)
	} else {
		log.Debug("Start time is in the past. Not sleeping.")
	}

	fmt.Println("[*] STARTED!!!!")
	log.Debugf("Debug message")
}
