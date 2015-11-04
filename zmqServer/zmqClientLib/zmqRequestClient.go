package zmqClientLib

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

func req_thread (zmq_url_port string, topic string, msg string) (string, error){	
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect(zmq_url_port)
	requester.Send(msg, 0)
	reply, err := requester.Recv(0)
	log.Println(reply, err.Error())
	return reply, err
}