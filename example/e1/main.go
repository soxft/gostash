package main

import (
	"github.com/soxft/gostash"
	"log"
)

func main() {
	pool, err := gostash.New("logstash.com", 50000, gostash.TCP, gostash.DefaultPoolConfig())
	if err != nil {
		log.Fatalf("logstash connect failed: %s", err)
	}

	pool.StringLog("{'name':'xcsoft','url':'https://github.com/soxft/gostash'}")
}
