package controller

import (
	"context"
	"data-platform-orders-cache-manager-rmq-kube/cache"
	"data-platform-orders-cache-manager-rmq-kube/controller/ordersdetail"
	"data-platform-orders-cache-manager-rmq-kube/controller/ordersdetaillist"
	"data-platform-orders-cache-manager-rmq-kube/controller/orderslist"
	rmqsessioncontroller "data-platform-orders-cache-manager-rmq-kube/rmq_session_controller"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type Controller struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger

	OrdersDetail     *ordersdetail.OrdersDetailCtrl
	OrdersList       *orderslist.OrdersListCtrl
	OrdersDetailList *ordersdetaillist.OrdersDetailListCtrl
}

func NewController(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *Controller {
	return &Controller{
		cache:            c,
		rmq:              rmq,
		ctx:              ctx,
		log:              log,
		OrdersDetail:     ordersdetail.NewOrdersDetailCtrl(ctx, c, rmq, log),
		OrdersList:       orderslist.NewOrdersListCtrl(ctx, c, rmq, log),
		OrdersDetailList: ordersdetaillist.NewOrdersDetailListCtrl(ctx, c, rmq, log),
	}
}
