//
//  Pubsub envelope publisher.
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"time"
	"log"
	"strings"
	"fmt"
)

const max_num_msg = 1000	//set threshold before blocking

func main() {
	
	//initPublisher()
	//  Prepare our subscriber	
	sub_topics := []string{"CustomerRegistration"}
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
	publisher.Connect("tcp://zmqpoxy:5559")
	
	for{ 
		v:= <- c
		cid := strings.Split(v, ":")
		msg := fmt.Sprintf("%s %s %s", "Dear Mr.", cid[1], ", Thank you for your business!")		
		publisher.Send(topic, zmq.SNDMORE)
		publisher.Send(msg, 0)
		log.Printf("<%s> %s", "Sending greeting message", msg)
		log.Println()
		time.Sleep(time.Second)
	}
}

func sub_thread(topics []string, c chan string){
	
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://zmqproxy:5560")
	
	for _, topic := range topics {
		subscriber.SetSubscribe(topic)
	}
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		log.Printf("[%s] %s\n", address, contents)
		c <- contents
	}
	
}

