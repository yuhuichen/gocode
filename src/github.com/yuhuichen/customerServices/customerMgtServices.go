//
//  Pubsub envelope publisher.
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"fmt"
	"time"
)

const max_num_msg = 1000	//set threshold before blocking

func main() {
	
	//initPublisher()
	//  Prepare our subscriber
	
	sub_topics := []string{"CustomerRegistration", "b", "c", "d"}
	pub_topic := "CustomerNotification"	
	 	
	msg_chan := make(chan string, max_num_msg)
	go sub_thread(sub_topics, msg_chan)
	go sendNotification_thread(pub_topic, msg_chan)
	ch := make(chan bool)
	ch <- true			//block the application from existing
}

func sendNotification_thread (topic string, c chan string){
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect("tcp://localhost:5563")

	msg := "Thank you for your business"
	
	for{ 
		v:= <- c
		fmt.Printf("[%s] %s\n", "Got: ", v)
		publisher.Send(topic, zmq.SNDMORE)
		publisher.Send(msg, 0)
		Printf("[%s] %s\n", "Sent message: ", msg)
		time.Sleep(time.Second)
	}
}

func sub_thread(topics []string, c chan string){
	
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5564")
	
	for _, topic := range topics {
		subscriber.SetSubscribe(topic)
	}
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
		c <- contents
	}
	
}

