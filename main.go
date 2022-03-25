package main

import (
	"fmt"
	"github.com/memcachier/mc/v3"
	"os"
)

func main() {

	m := mc.NewMC(os.Getenv("MEMCACHED_SERVER"), os.Getenv("MEMCACHED_USERNAME"), os.Getenv("MEMCACHED_PASSWORD"))
	defer m.Quit()

	s, err := m.Stats()
	if err != nil {
		fmt.Println("Couldn't connect to memcached:", err)
		os.Exit(1)
	}

	for _, d := range s {
		if v, ok := d["version"]; ok {
			fmt.Println("Connection to memcached is ok, memcached version:", v)
		}
	}
}
