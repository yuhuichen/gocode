//
//  Pubsub envelope subscriber.
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

func main() {
	//  Prepare our subscriber
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5560")
	
	topics := []string{"CustomerNotification"}
	for _, topic := range topics {
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