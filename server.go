package kn

import (
	udp "github.com/kxrr/udp-server"
	log "github.com/kxrr/klog"
	"strings"
)

func StartUdpServer(host string, port int, bufsize int, mangonel TelegramMangonel) (chan bool, error) {
	messages, err := udp.ListenUdp(host, port, bufsize, 100)
	if err != nil {
		return nil, err
	}
	log.Printf("UDP server listening on %s:%d", host, port)

	stop := make(chan bool, 1)
	go func() {
		for {
			select {
			case message := <-messages:
				if message.Error != nil {
					log.Errorf("Error when reading message: %s", message.Error.Error())
					continue
				}
				log.Debugf("Forwarding udp message to telegram from %s", message.Remote)
				text := strings.Replace(string(message.Data), "\\n", "\n", -1)
				mangonel.SendMdMessageNoWait(text)
			case <-stop:
				return
			}
		}
	}()
	return stop, nil
}
