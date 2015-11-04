package zmqClientLib

import (
	zmq "github.com/pebbe/zmq4"
	"log"
	
)


// not really useful in request/reply mode. But worth having this code, perhaps as an example. 
func reply_thread(zmq_url_port string, request_chan chan string, reply_chan chan string){
		
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Connect(zmq_url_port)
	
	for {

		request, _ := responder.Recv(0)
		log.Printf("Received request: [%s]\n", request)

		request_chan <- request				//unbuffered channel to expect a request/reply pair,
											//however unless the invoker very careful, mismatch may happen.
		reply_msg := <- reply_chan				//careful! may not be the reply to the corresponding request
		
		log.Printf("Sent reply: [%s]\n", reply_msg)
		responder.Send(reply_msg, 0)
	}
}

// not really useful in request/reply mode. But worth having this code, perhaps as an example. 
func reply_bytes_thread(zmq_url_port string, request_chan chan []byte, reply_chan chan []byte){
		
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Connect(zmq_url_port)
	
	for {

		request, _ := responder.RecvBytes(0)
		log.Printf("Received request: [%s]\n", request)

		request_chan <- request			//unbuffered channel to expect a request/reply pair,
										//however unless the invoker very careful, mismatch may happen.
		reply_msg := <- reply_chan		//careful! may not be the reply to the corresponding request
		
		log.Printf("Sent reply: [%s]\n", reply_msg)
		responder.SendBytes(reply_msg, 0)
	}
}