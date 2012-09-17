package main

import (
	"fmt"
	"sync"

	redis "engine/goredis"
)

var (
	client *redis.Client
	mux    sync.Mutex
)

func init() {
	mux.Lock()
	defer mux.Unlock()

	if client != nil {
		return
	}

	client = &redis.Client{
		Addr: "127.0.0.1:6379",
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// string
	// Set
	err := client.Set("name", []byte("viney"))
	checkErr(err)

	// Get
	value, err := client.Get("name")
	checkErr(err)
	fmt.Println(string(value))

	// Del
	ok, err := client.Del("name")
	checkErr(err)

	fmt.Println(ok)

	// list
	var list = []string{"a", "b", "c", "d", "e"}
	for _, v := range list {
		err = client.Rpush("1", []byte(v))
		checkErr(err)
	}

	lrange, err := client.Lrange("1", 0, 4)
	checkErr(err)

	for _, v := range lrange {
		fmt.Println(string(v))
	}

	ok, err = client.Del("1")
	checkErr(err)
	fmt.Println(ok)
}
