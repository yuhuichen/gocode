//
//  Pubsub envelope publisher.
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"time"
	"strconv"
)

func main() {
	//  Prepare our publisher
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect("tcp://localhost:5563")
	
	var topic = "CustomerRegistration"
	var n = 0
	
	for {
		//  Write two messages, each with an envelope and content
		publisher.Send(topic, zmq.SNDMORE)
		msg :="A new customer registered, " + strconv.Itoa(n)
		publisher.Send(msg, 0)
		n += 1
		time.Sleep(time.Second*5)
	}
}