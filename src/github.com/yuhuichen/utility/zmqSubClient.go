package utility

import (
	zmq "github.com/pebbe/zmq4"
)

func sub_thread(zmq_url_port string, topics []string, c chan string){
	
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect(zmq_url_port)
	
	for _, topic := range topics {
		subscriber.SetSubscribe(topic)
	}
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		Trace.Println("[%s] %s\n", address, contents)
		c <- contents
	}
	
}