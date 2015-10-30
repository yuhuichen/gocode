//
//  Pubsub envelope subscriber.
//

package zmqClients

import (
	zmq "github.com/pebbe/zmq4"
	"fmt"
	gcfg "gopkg.in/gcfg.v1"
)

const configFileName = "/home/vagrant/dockercfg/zmq.cfg"


type subClient struct {
	zmqProxy_url_port string
	topic	string
}

func main() {
	
	//read config file	
	var cfg subClient
	gcfg.ReadFileInto(&cfg, configFileName)
	
	//  Prepare our subscriber
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect(cfg.zmqProxy_url_port)
	subscriber.SetSubscribe(cfg.topic)

	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
	}
}



