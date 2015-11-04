//
//  Weather proxy device.
//
//  NOT TESTED
//

package main

import (
	zmq "github.com/pebbe/zmq4"	
	"log"
)

const debug = true
const frontend_url_port ="tcp://*:5555"
const backend_url_port = "tcp://*:5556"

func main() {

	//config := ReadConfig("")
	//  Prepare our sockets
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	frontend.Bind(frontend_url_port)
	backend.Bind(backend_url_port)

	//  Initialize poll set
	poller := zmq.NewPoller()
	poller.Add(frontend, zmq.POLLIN)
	poller.Add(backend, zmq.POLLIN)
	
	log.Println("0MQ broker started!")
	log.Println("Frontend protocl/url/port:", frontend_url_port)
	log.Println("Backend protocol/url/port:", backend_url_port)

	//  Switch messages between sockets
	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case frontend:
				for {
					msg, _ := s.Recv(0)
					if more, _ := s.GetRcvmore(); more {
						backend.Send(msg, zmq.SNDMORE)
					} else {
						backend.Send(msg, 0)
						break
					}
					if debug {log.Printf("relayed request: [%s]\n", msg)}
					
				}
			case backend:
				for {
					msg, _ := s.Recv(0)
					if more, _ := s.GetRcvmore(); more {
						frontend.Send(msg, zmq.SNDMORE)
					} else {
						frontend.Send(msg, 0)
						break
					}
					if debug {log.Printf("relayed reply: [%s]\n", msg)}
				}
				
			}
		}
	}
	//log.Fatalln("Proxy interrupted:", err)
}