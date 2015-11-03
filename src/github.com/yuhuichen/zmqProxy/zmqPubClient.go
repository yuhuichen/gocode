package zmqProxy

import (
	zmq "github.com/pebbe/zmq4"
	"time"
)

func pub_thread (zmq_url_port string, topic string, c chan string){
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect(zmq_url_port)
		
	for{ 
		msg := <- c
		publisher.Send(topic, zmq.SNDMORE)
		publisher.Send(msg, 0)
		Trace.Println("[%s] %s\n", "Sent message: ", msg)
		time.Sleep(time.Second)
	}
}