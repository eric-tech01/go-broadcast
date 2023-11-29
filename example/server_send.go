package main

import (
	"fmt"
	"time"

	json "github.com/eric-tech01/simple-json"

	broadcast "github.com/eric-tech01/go-broadcast"
)

func main() {
	fmt.Println("test")
	broadcastIp := "239.0.0.1:12345"
	b, err := broadcast.NewBroadcast(broadcastIp, "192.168.0.48")
	if err != nil {
		fmt.Println(err)
		return
	}

	j := json.New()
	j.Set("server", "192.168.0.48")
	j.Set("url", "https://baidu.com")
	for {
		n, err := b.Send(j.ToBytes())
		fmt.Println("send len: ", n)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(2 * time.Second)
	}

}
