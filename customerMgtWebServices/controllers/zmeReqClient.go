package controllers

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

func zmq_req(zmq_url_port string, data []byte) ([]byte, error){	
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect(zmq_url_port)
	requester.SendBytes(data, 0)
	//requester.Send(msg, 0)
	//reply, err := requester.Recv(0)
	reply, err := requester.RecvBytes(0)
	log.Println(reply, err.Error())
	return reply, err
}