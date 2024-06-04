package gologstash

import (
	"log"
	"net"
)

func (l *GoStash) StringLog(str string) {
	connection, err := l.Pool.Get()
	if err != nil {
		log.Println("[Logstash Logger] Failed to get connection from pool: ", err)
		return
	}

	_, connWriteError := connection.(net.Conn).Write([]byte(str))
	if connWriteError != nil {
		log.Printf("[Logstash Logger] Failed to send log String message: %s\n", connWriteError)
	}

	_ = l.Pool.Put(connection)
}
