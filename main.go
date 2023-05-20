package main

import (
	"fmt"
	"log"
	"time"
)

func sleep() {
	var forever chan struct{}
	<-forever
}

func main() {
	config := MQConfig{
		host: "localhost",
		port: "5672",
		user: "guest",
		pass: "guest",
	}

	mq := MQ{}
	mq.Init(config)

	q := mq.Queue("QUEUUUUE")

	consumer := func(id string) {
		cb := func(msg []byte) {
			log.Printf("Received a message: %s, I am %s btw", msg, id)
		}
		q.Consume(cb)
	}

	go consumer("apollo")
	go consumer("vladimir")

	everyFiveSeconds := time.NewTicker(2 * time.Second)
	defer everyFiveSeconds.Stop()

	go func() {
		i := 0
		for {
			select {
			case <-everyFiveSeconds.C:
				body := "HEYOO! person number " + fmt.Sprintf("%d", i)
				q.Publish([]byte(body))
				i++
			}
		}
	}()

	sleep()
}
