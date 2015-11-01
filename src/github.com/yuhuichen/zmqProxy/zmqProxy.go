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

const frontend_url_port ="tcp://*:5563"
const backend_url_port = "tcp://*:5564"

func listener_thread() { 
	pipe, _ := zmq.NewSocket(zmq.PAIR) 
 	pipe.Bind("inproc://pipe") 
 
 
 	//  Print everything that arrives on pipe 
	for { 
		msg, err := pipe.RecvMessage(0) 
		if err != nil { 
			break //  Interrupted 
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
	frontend.Bind(frontend_url_port)
	
	//  This is our public endpoint for subscribers
	backend, _ := zmq.NewSocket(zmq.XPUB)
	defer backend.Close()
	backend.Bind(backend_url_port)
	
	listener, _ := zmq.NewSocket(zmq.PAIR) 
	listener.Connect("inproc://pipe") 

	//  Run the proxy until the user interrupts us
	err := zmq.Proxy(frontend, backend, listener)
	log.Fatalln("Proxy interrupted:", err)
}