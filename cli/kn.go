package main

import (
	"os"
	"fmt"
	"flag"
	"github.com/kxrr/kn"
)


func main() {
	token := flag.String("token", "", "Telegram bot token")
	userId := flag.Int64("userid", 0, "Telegram user id")
	host := flag.String("host", "0.0.0.0", "UDP Server host")
	port := flag.Int("port", 9996, "UDP Server port")
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
		1024,
		*tm,
	)
	if err != nil {
		panic(err)
	}
	<-stop
}
