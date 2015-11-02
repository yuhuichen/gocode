//
//  Pubsub envelope subscriber.
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)


const sub_zmq_url_port = "tcp://zmqproxy:5560"
var sub_topics = []string{"CustomerNotification",}

func main() {
	//  Prepare our subscriber
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect(sub_zmq_url_port)

	for _, topic := range sub_topics {
		subscriber.SetSubscribe(topic)
	}

	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		log.Printf("[%s] %s\n", address, contents)		
	}
}