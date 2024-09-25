package main

import (
	redis "matchpoll/config"
	"matchpoll/cron"
	"matchpoll/services"
)

func main() {
	//Init
	redis.CreateConnection()

	// Start Polling Service
	cron.Start()

	// Start Writer Service
	services.StartPublishService()

}
