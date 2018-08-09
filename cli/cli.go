package main

import (
	"flag"
	"github.com/kxrr/kn"
	"os"
	"fmt"
)

func main() {
	token := flag.String("token", "", "Bot token")
	userId := flag.Int64("userid", 0, "User id")
	buffSize := flag.Int("buff", 2048, "Buffer size")
	host := flag.String("host", "0.0.0.0", "Server host")
	port := flag.Int("port", 9996, "Server port")
	debug := flag.Bool("debug", false, "Turn on debug")
	flag.Parse()

	if *userId == 0 || *token == "" {
		fmt.Printf("Please input user id and token\n")
		os.Exit(2)
	}

	tm, err := kn.NewTelegramMangonel(*userId, *token, *debug)
	if err != nil {
		panic(err)
	}
	stop, err := kn.StartUdpServer(
		*host,
		*port,
		*buffSize,
		*tm,
	)
	if err != nil {
		panic(err)
	}
	<-stop
}
