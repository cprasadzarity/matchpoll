package services

import (
	"fmt"
	"gopkg.in/antage/eventsource.v1"
	"log"
	redis "matchpoll/config"
	"net/http"
	"strconv"
	"time"
)

func StartPublishService() {
	es := eventsource.New(nil, nil)
	defer es.Close()
	http.Handle("/events", es)
	go func() {
		score, err := redis.GetValue("score")
		if err != nil {
			log.Fatal(err)
		}
		id := 1
		for {
			fmt.Println("Sending events to subs")
			es.SendEventMessage(score, "score", strconv.Itoa(id))
			id++
			time.Sleep(2 * time.Second)
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
