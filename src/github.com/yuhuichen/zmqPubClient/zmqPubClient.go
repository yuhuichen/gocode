//
//  Pubsub envelope publisher.
//

package zmqPubClient

import (
	zmq "github.com/pebbe/zmq4"
	// "fmt"
)

var publisher, _ = zmq.NewSocket(zmq.PUB)
type pubClient struct {
	zmqProxy_url_port	string
}

func sendMsg(topic string, msg string) {
	defer publisher.Close()
	publisher.Send(topic, zmq.SNDMORE)
	publisher.Send(msg, 0)
}

func initZmqPubClient(zmqProxy_url_port string){

	//pubClient.zmqProxy_url_port = zmqProxy_url_port
	
	//  Prepare our publisher
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect(zmqProxy_url_port)
}

func clostZmqClient(){
	defer publisher.Close()
}


