package controllers

import (
    //"sync"
	//zmqClient "github.com/yuhuichen/zmqServer/zmqClientLib"
)


type (
	ZmqClientProxy struct {
		session interface{}
	}
)

func NewZmqClientProxy() *ZmqClientProxy {
	return &ZmqClientProxy{getSession()}
}