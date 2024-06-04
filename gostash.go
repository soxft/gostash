package gologstash

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/silenceper/pool"
	"net"
	"time"
)

func DefaultPoolConfig() pool.Config {
	return pool.Config{
		InitialCap:  5,  // 资源池初始连接数
		MaxIdle:     20, // 最大空闲连接数
		MaxCap:      30, // 最大并发连接数
		IdleTimeout: 15 * time.Second,
	}
}

func New(
	Host string,
	Port int,
	Protocol Protocol,
	config pool.Config,
) (*GoStash, error) {

	if Protocol != TCP {
		return nil, fmt.Errorf("unsupported protocol: %s", Protocol)
	}

	factory := func() (interface{}, error) { return net.Dial("tcp", fmt.Sprintf("%s:%d", Host, Port)) }
	clo := func(v interface{}) error { return v.(net.Conn).Close() }

	//ping 检测连接的方法
	//ping := func(v interface{}) error { return nil }

	//创建一个连接池： 初始化5，最大空闲连接是20，最大并发连接30
	config.Factory = factory
	config.Close = clo
	p, err := pool.NewChannelPool(&config)
	if err != nil {
		return nil, errors.Wrap(err, "pool init failed")
	}

	return &GoStash{
		Pool: p,
	}, nil
}
