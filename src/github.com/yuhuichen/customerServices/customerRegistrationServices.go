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


const	pub_zmq_url_port = "tcp://zmqproxy:5559"
var	pub_topic = "CustomerRegistration"	
//sub_zmq_url_port := "tcp://zmqproxy:5560"
//sub_topics := []string{"CustomerRegistration"}


func main() {
	//  Prepare our publisher
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect(pub_zmq_url_port)
	
	var n = 0
	
	for {
		//  Write two messages, each with an envelope and content
		publisher.Send(pub_topic, zmq.SNDMORE)
		msg :="A new customer registered - customer ID:" + strconv.Itoa(n)
		log.Printf("[%s] %s\n", pub_topic, msg)
		publisher.Send(msg, 0)
		n += 1
		time.Sleep(time.Second*5)
	}
}