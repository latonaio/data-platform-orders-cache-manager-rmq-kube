package main

import (
	"context"
	"data-platform-orders-cache-manager-rmq-kube/cache"
	"data-platform-orders-cache-manager-rmq-kube/config"
	"data-platform-orders-cache-manager-rmq-kube/controller"
	rmqsessioncontroller "data-platform-orders-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"
	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()
	conf := config.NewConf()
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), "", nil, 0)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	rmqCtrl, err := rmqsessioncontroller.NewRMQSessionCtrl(conf, l)
	if err != nil {
		l.Fatal(err.Error())
	}
	cache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l)
	cacheCtrl := controller.NewController(ctx, cache, rmqCtrl, l)

	l.Info("READY!")
	for msg := range iter {
		go func(msg rabbitmq.RabbitmqMessage) {
			err = callProcess(msg, cacheCtrl, conf)
			if err != nil {
				l.Error(err)
			}
			msg.Success()
		}(msg)
	}
}

func recovery(l *logger.Logger, err *error) {
	if e := recover(); e != nil {
		*err = xerrors.Errorf("error occurred: %w", e)
		l.Error("%+v", *err)
		return
	}
}

func callProcess(msg rabbitmq.RabbitmqMessage, cacheCtrl *controller.Controller, conf *config.Conf) (err error) {
	l := logger.NewLogger()
	defer recovery(l, &err)
	input := map[string]interface{}{}
	err = json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		return err
	}
	switch input["ui_function"] {
	case "OrdersDetail":
		err = cacheCtrl.OrdersDetail.OrdersDetail(msg)
	case "OrdersListBuyer":
		err = cacheCtrl.OrdersList.OrdersList(msg)
	case "OrdersListSeller":
		err = cacheCtrl.OrdersList.OrdersList(msg)
	case "OrdersDetailList":
		err = cacheCtrl.OrdersDetailList.OrdersDetailList(msg)
	default:
		l.Info("Unknow ui-function %v", input["ui_function"])
	}

	// sID, err := getSessionID(msg.Data())
	return err
}

func getSessionID(req interface{}) (string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", err
	}
	rawSID, ok := m["runtime_session_id"]
	if !ok {
		return "", xerrors.Errorf("runtime_session_id not included")
	}

	return fmt.Sprintf("%v", rawSID), nil
}
