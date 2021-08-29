package main

import (
	"context"
	"log"
	"time"
)

type Client struct {
	ch1 chan interface{}
	ch2 chan interface{}
	ctx context.Context
}

func (c *Client) loop1() {
	for {
		select {
		case <-c.ctx.Done():
			log.Printf("loop1 exit")
			return
		case <-c.ch1:
			log.Printf("recvice loop1 event")
		}
	}
}

func (c *Client) loop2() {
	for {
		select {
		case <-c.ctx.Done():
			log.Printf("loop2 exit")
			return
		case <-c.ch2:
			log.Printf("recvice loop2 event")
		}
	}
}

func main() {
	rootCtx := context.Background()
	ctx, cancelFn := context.WithCancel(rootCtx)

	c := &Client{
		ch1: make(chan interface{}, 3),
		ch2: make(chan interface{}, 3),
		ctx: ctx,
	}

	go c.loop1()
	go c.loop2()

	c.ch1 <- struct{}{}
	c.ch2 <- struct{}{}
	c.ch1 <- struct{}{}
	c.ch2 <- struct{}{}
	c.ch1 <- struct{}{}
	c.ch2 <- struct{}{}

	<-time.After(time.Second)
	cancelFn()
	<-time.After(time.Second)
}
