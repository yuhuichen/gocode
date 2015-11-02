//
//  Pubsub envelope publisher.
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"time"
	"strconv"
	"log"
)

func main() {
	//  Prepare our publisher
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect("tcp://zmqproxy:5559")
	
	var topic = "CustomerRegistration"
	var n = 0
	
	for {
		//  Write two messages, each with an envelope and content
		publisher.Send(topic, zmq.SNDMORE)
		msg :="A new customer registered - customer ID:" + strconv.Itoa(n)
		log.Printf("[%s] %s\n", topic, msg)
		publisher.Send(msg, 0)
		n += 1
		time.Sleep(time.Second*5)
	}
}