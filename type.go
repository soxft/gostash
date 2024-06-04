package gologstash

import "github.com/silenceper/pool"

type GoStash struct {
	Pool pool.Pool
}

type Protocol string

const (
	TCP Protocol = "tcp"
)
