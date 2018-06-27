package main

import (
	"log"
	"time"

	"github.com/semihalev/redis-balancer"
)

func main() {
	clients := balancer.New(
		[]balancer.Options{
			{Network: "tcp", Addr: "localhost:6379", CheckInterval: 600 * time.Millisecond},
			{Network: "unix", Addr: "/tmp/redis.sock"},
			{Network: "tcp", Addr: "host-2:6379", CheckInterval: 800 * time.Millisecond},
			{Network: "tcp", Addr: "host-2:6380"},
		},
		balancer.ModeLeastConn,
	)
	defer clients.Close()

	client := clients.Next()
	conn := client.Get()
	defer conn.Close()

	res, err := conn.Do("PING")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
