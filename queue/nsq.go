package queue

import (
	"github.com/nsqio/go-nsq"
	"log"
)


type NSQ struct {
	producer *nsq.Producer
}

func New(addr string) *NSQ {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(addr, config)
	if err != nil {
		log.Panic(err)
	}
	return &NSQ{producer: p}
}

func (queue *NSQ) Publish(topic string, data []byte) error  {
	
	err := queue.producer.Publish(topic,data)
	if err != nil {
		return err
	}
	return nil
}

