//
//  Weather proxy device.
//
//  NOT TESTED
//

package main

import (
	zmq "github.com/pebbe/zmq4"
	"log"
	"time"
)

const proxy_frontend_url_port ="tcp://*:5559"
const proxy_backend_url_port = "tcp://*:5560"

func listener_thread() { 
	pipe, _ := zmq.NewSocket(zmq.PAIR) 
 	pipe.Bind("inproc://pipe") 
 
 	//  Print everything that arrives on pipe 
	for { 
		msg, err := pipe.RecvMessage(0) 
		if err != nil { 
			break 		//  Interrupted 
 		} 
 		log.Println(msg) 		
 	} 
} 

func main() {

	go listener_thread()
	time.Sleep(time.Millisecond * 5)
	
	//  This is where the weather server sits
	frontend, _ := zmq.NewSocket(zmq.XSUB)
	defer frontend.Close()
	frontend.Bind(proxy_frontend_url_port)
	
	//  This is our public endpoint for subscribers
	backend, _ := zmq.NewSocket(zmq.XPUB)
	defer backend.Close()
	backend.Bind(proxy_backend_url_port)
	
	listener, _ := zmq.NewSocket(zmq.PAIR) 
	listener.Connect("inproc://pipe") 

	log.Println("0MQ proxy started!")
	log.Println("Frontend protocl/url/port:", proxy_frontend_url_port)
	log.Println("Backend protocol/url/port:", proxy_backend_url_port)
	
	//  Run the proxy until the user interrupts us
	err := zmq.Proxy(frontend, backend, listener)
	log.Fatalln("Proxy interrupted:", err)
}