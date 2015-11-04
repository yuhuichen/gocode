package zmqClientLib

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

func Req_msg (zmq_url_port string, topic string, msg string) (string, error){	
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect(zmq_url_port)
	requester.Send(msg, 0)
	reply, err := requester.Recv(0)
	log.Println(reply, err.Error())
	return reply, err
}

func Req_bytes (zmq_url_port string, msg []byte) ([]byte, error){	
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect(zmq_url_port)
	requester.SendBytes(msg, 0)
	reply, err := requester.RecvBytes(0)
	//log.Println(string(reply), err.Error())
	return reply, err
}