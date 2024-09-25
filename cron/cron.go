package cron

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"matchpoll/services"
	"time"
)

func Start() {
	// create a scheduler
	s := gocron.NewScheduler(time.UTC)

	s.Every(10).Seconds().Do(func() {
		fmt.Println("Polling data from servers")
		services.Poll()
	})

	s.StartAsync()
}

func Stop() {

}
