# gostash
> send logs to logstash from golang, connection pool supported

# install
> go get github.com/soxft/gostash

# usage

```go
    pool, err := gostash.New("logstash.com", 50000, gostash.TCP, gostash.DefaultPoolConfig())
    if err != nil {
        log.Fatalf("logstash connect failed: %s", err)
    }
    
    pool.StringLog("{'name':'xcsoft','url':'https://github.com/soxft/gostash'}")
```

