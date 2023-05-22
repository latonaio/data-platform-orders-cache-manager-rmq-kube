package rmqsessioncontroller

import (
	"data-platform-orders-cache-manager-rmq-kube/config"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type RMQSessionCtrl struct {
	rmq    *rabbitmq.RabbitmqClient
	reqMap map[string]chan rabbitmq.RabbitmqMessage
	log    *logger.Logger
}

func NewRMQSessionCtrl(conf *config.Conf, log *logger.Logger) (*RMQSessionCtrl, error) {
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.SessionControlQueue(), "", nil, 0)
	if err != nil {
		return nil, err
	}
	c := &RMQSessionCtrl{
		rmq:    rmq,
		reqMap: make(map[string]chan rabbitmq.RabbitmqMessage),
		log:    log,
	}
	go c.watchResponse()
	return c, nil
}

func (c *RMQSessionCtrl) watchResponse() {
	iter, err := c.rmq.Iterator()
	if err != nil {
		c.log.Fatal(err)
	}
	for msg := range iter {
		msg.Success()
		d := msg.Data()
		id, ok := d["runtime_session_id"]
		if !ok {
			c.log.Error("unknown message received")
			continue
		}
		sID, ok := id.(string)
		if !ok {
			c.log.Error("unknown message received. runtime_session_id is %v", id)
			continue
		}
		_, ok = c.reqMap[sID]
		if !ok {
			c.log.Error("unknown runtime_session_id %v", id)
			continue
		}
		c.reqMap[sID] <- msg
	}

}

func (c *RMQSessionCtrl) SessionRequest(sendQueue string, payload interface{}, sessionID string) func() rabbitmq.RabbitmqMessage {
	c.rmq.Send(sendQueue, payload)
	c.reqMap[sessionID] = make(chan rabbitmq.RabbitmqMessage)
	return c.receiveResponse(sessionID)
}

func (c *RMQSessionCtrl) receiveResponse(sID string) func() rabbitmq.RabbitmqMessage {
	var msg rabbitmq.RabbitmqMessage = nil
	return func() rabbitmq.RabbitmqMessage {
		ticker := time.NewTicker(10 * time.Second)
		select {
		case msg = <-c.reqMap[sID]:
		case <-ticker.C:
			c.log.Error("could not get response of session id %v", sID)
		}
		return msg
	}
}
