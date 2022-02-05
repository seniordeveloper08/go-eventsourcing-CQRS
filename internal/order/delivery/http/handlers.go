package http

import (
	"github.com/AleksK1NG/es-microservice/config"
	"github.com/AleksK1NG/es-microservice/internal/metrics"
	"github.com/AleksK1NG/es-microservice/internal/order/aggregate"
	"github.com/AleksK1NG/es-microservice/internal/order/events"
	"github.com/AleksK1NG/es-microservice/internal/order/models"
	"github.com/AleksK1NG/es-microservice/internal/order/queries"
	"github.com/AleksK1NG/es-microservice/internal/order/service"
	"github.com/AleksK1NG/es-microservice/pkg/constants"
	httpErrors "github.com/AleksK1NG/es-microservice/pkg/http_errors"
	"github.com/AleksK1NG/es-microservice/pkg/logger"
	"github.com/AleksK1NG/es-microservice/pkg/middlewares"
	"github.com/AleksK1NG/es-microservice/pkg/tracing"
	"github.com/AleksK1NG/es-microservice/pkg/utils"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type OrderHandlers interface {
	CreateOrder() echo.HandlerFunc
	PayOrder() echo.HandlerFunc
	SubmitOrder() echo.HandlerFunc
	UpdateOrder() echo.HandlerFunc

	GetOrderByID() echo.HandlerFunc
	Search() echo.HandlerFunc
}

type orderHandlers struct {
	group   *echo.Group
	log     logger.Logger
	mw      middlewares.MiddlewareManager
	cfg     *config.Config
	v       *validator.Validate
	os      *service.OrderService
	metrics *metrics.ESMicroserviceMetrics
}

func NewOrderHandlers(
	group *echo.Group,
	log logger.Logger,
	mw middlewares.MiddlewareManager,
	cfg *config.Config,
	v *validator.Validate,
	os *service.OrderService,
	metrics *metrics.ESMicroserviceMetrics,
) *orderHandlers {
	return &orderHandlers{group: group, log: log, mw: mw, cfg: cfg, v: v, os: os, metrics: metrics}
}

// CreateOrder
// @Tags Orders
// @Summary Create order
// @Description Create new order
// @Param order body events.OrderCreatedEventData true "create order"
// @Accept json
// @Produce json
// @Success 201 {string} id ""
// @Router /orders [post]
func (h *orderHandlers) CreateOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.StartHttpServerTracerSpan(c, "orderHandlers.CreateOrder")
		defer span.Finish()
		h.metrics.CreateOrderHttpRequests.Inc()

		eventData := events.OrderCreatedEventData{}
		if err := c.Bind(&eventData); err != nil {
			h.log.Errorf("(Bind) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		if err := h.v.StructCtx(ctx, eventData); err != nil {
			h.log.Errorf("(validate) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		id := uuid.NewV4().String()
		command := aggregate.NewCreateOrderCommand(eventData, id)
		err := h.os.Commands.CreateOrder.Handle(ctx, command)
		if err != nil {
			h.log.Errorf("(CreateOrder.Handle) id: {%s}, err: {%v}", id, err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.log.Infof("(order created) id: {%s}", id)
		return c.JSON(http.StatusCreated, id)
	}
}

// PayOrder
// @Tags Orders
// @Summary Pay order
// @Description Pay existing order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {string} id ""
// @Router /orders/pay/{id} [put]
func (h *orderHandlers) PayOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.StartHttpServerTracerSpan(c, "orderHandlers.PayOrder")
		defer span.Finish()
		h.metrics.PayOrderHttpRequests.Inc()

		orderID, err := uuid.FromString(c.Param(constants.ID))
		if err != nil {
			h.log.Errorf("(uuid.FromString) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		var payment models.Payment
		if err := c.Bind(&payment); err != nil {
			h.log.Errorf("(Bind) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		command := aggregate.NewOrderPaidCommand(payment, orderID.String())
		if err := h.v.StructCtx(ctx, command); err != nil {
			h.log.Errorf("(validate) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		err = h.os.Commands.OrderPaid.Handle(ctx, command)
		if err != nil {
			h.log.Errorf("(OrderPaid.Handle) id: {%s}, err: {%v}", orderID.String(), err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.log.Infof("(order paid) id: {%s}", orderID.String())
		return c.JSON(http.StatusOK, orderID.String())
	}
}

// SubmitOrder
// @Tags Orders
// @Summary Submit order
// @Description Submit existing order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {string} id ""
// @Router /orders/submit/{id} [put]
func (h *orderHandlers) SubmitOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.StartHttpServerTracerSpan(c, "orderHandlers.SubmitOrder")
		defer span.Finish()
		h.metrics.SubmitOrderHttpRequests.Inc()

		orderID, err := uuid.FromString(c.Param(constants.ID))
		if err != nil {
			h.log.Errorf("(uuid.FromString) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		command := aggregate.NewSubmitOrderCommand(orderID.String())
		if err := h.v.StructCtx(ctx, command); err != nil {
			h.log.Errorf("(validate) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		err = h.os.Commands.SubmitOrder.Handle(ctx, command)
		if err != nil {
			h.log.Errorf("(SubmitOrder.Handle) id: {%s}, err: {%v}", orderID.String(), err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.log.Infof("(order submitted) id: {%s}", orderID.String())
		return c.JSON(http.StatusOK, orderID.String())
	}
}

// CancelOrder
// @Tags Orders
// @Summary Cancel order
// @Description Cancel existing order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {string} id ""
// @Router /orders/cancel/{id} [post]
func (h *orderHandlers) CancelOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.StartHttpServerTracerSpan(c, "orderHandlers.CancelOrder")
		defer span.Finish()
		h.metrics.SubmitOrderHttpRequests.Inc()

		orderID, err := uuid.FromString(c.Param(constants.ID))
		if err != nil {
			h.log.Errorf("(uuid.FromString) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		var data events.OrderCanceledEventData
		if err := c.Bind(&data); err != nil {
			h.log.Errorf("(Bind) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		command := aggregate.NewOrderCanceledCommand(data, orderID.String())
		if err := h.v.StructCtx(ctx, command); err != nil {
			h.log.Errorf("(validate) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		err = h.os.Commands.CancelOrder.Handle(ctx, command)
		if err != nil {
			h.log.Errorf("(CancelOrder.Handle) id: {%s}, err: {%v}", orderID.String(), err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.log.Infof("(order canceled) id: {%s}", orderID.String())
		return c.JSON(http.StatusOK, orderID.String())
	}
}

// DeliverOrder
// @Tags Orders
// @Summary Deliver order
// @Description Deliver existing order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {string} id ""
// @Router /orders/delivery/{id} [post]
func (h *orderHandlers) DeliverOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.StartHttpServerTracerSpan(c, "orderHandlers.DeliverOrder")
		defer span.Finish()
		h.metrics.SubmitOrderHttpRequests.Inc()

		orderID, err := uuid.FromString(c.Param(constants.ID))
		if err != nil {
			h.log.Errorf("(uuid.FromString) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		var data events.OrderDeliveredEventData
		if err := c.Bind(&data); err != nil {
			h.log.Errorf("(Bind) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		command := aggregate.NewOrderDeliveredCommand(data, orderID.String())
		if err := h.v.StructCtx(ctx, command); err != nil {
			h.log.Errorf("(validate) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		err = h.os.Commands.DeliveryOrder.Handle(ctx, command)
		if err != nil {
			h.log.Errorf("(DeliveryOrder.Handle) id: {%s}, err: {%v}", orderID.String(), err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.log.Infof("(order delivered) id: {%s}", orderID.String())
		return c.JSON(http.StatusOK, orderID.String())
	}
}

// ChangeDeliveryAddress
// @Tags Orders
// @Summary Change delivery address order
// @Description Deliver existing order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {string} id ""
// @Router /orders/address/{id} [put]
func (h *orderHandlers) ChangeDeliveryAddress() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.StartHttpServerTracerSpan(c, "orderHandlers.ChangeDeliveryAddress")
		defer span.Finish()
		h.metrics.SubmitOrderHttpRequests.Inc()

		orderID, err := uuid.FromString(c.Param(constants.ID))
		if err != nil {
			h.log.Errorf("(uuid.FromString) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		var data events.OrderChangeDeliveryAddress
		if err := c.Bind(&data); err != nil {
			h.log.Errorf("(Bind) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		command := aggregate.NewOrderChangeDeliveryAddressCommand(data, orderID.String())
		if err := h.v.StructCtx(ctx, command); err != nil {
			h.log.Errorf("(validate) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		err = h.os.Commands.ChangeOrderDeliveryAddress.Handle(ctx, command)
		if err != nil {
			h.log.Errorf("(ChangeOrderDeliveryAddress.Handle) id: {%s}, err: {%v}", orderID.String(), err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.log.Infof("(ChangeDeliveryAddress) id: {%s}", orderID.String())
		return c.JSON(http.StatusOK, orderID.String())
	}
}

// UpdateOrder
// @Tags Orders
// @Summary Update order
// @Description Update existing order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body events.OrderUpdatedEventData true "update order"
// @Success 200 {string} id ""
// @Router /orders/{id} [put]
func (h *orderHandlers) UpdateOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.StartHttpServerTracerSpan(c, "orderHandlers.UpdateOrder")
		defer span.Finish()
		h.metrics.UpdateOrderHttpRequests.Inc()

		orderID, err := uuid.FromString(c.Param(constants.ID))
		if err != nil {
			h.log.Errorf("(uuid.FromString) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		eventData := events.OrderUpdatedEventData{}
		if err := c.Bind(&eventData); err != nil {
			h.log.Errorf("(Bind) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		if err := h.v.StructCtx(ctx, eventData); err != nil {
			h.log.Errorf("(validate) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		command := aggregate.NewOrderUpdatedCommand(eventData, orderID.String())
		err = h.os.Commands.UpdateOrder.Handle(ctx, command)
		if err != nil {
			h.log.Errorf("(UpdateOrder.Handle) id: {%s}, err: {%v}", orderID.String(), err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.log.Infof("(order updated) id: {%s}", orderID.String())
		return c.JSON(http.StatusOK, orderID.String())
	}
}

// GetOrderByID
// @Tags Orders
// @Summary Get order
// @Description Get order by id
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.OrderProjection
// @Router /orders/{id} [get]
func (h *orderHandlers) GetOrderByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.StartHttpServerTracerSpan(c, "orderHandlers.GetOrderByID")
		defer span.Finish()
		h.metrics.GetOrderByIdHttpRequests.Inc()

		orderID, err := uuid.FromString(c.Param(constants.ID))
		if err != nil {
			h.log.Errorf("(uuid.FromString) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		query := queries.NewGetOrderByIDQuery(orderID.String())
		if err := h.v.StructCtx(ctx, query); err != nil {
			h.log.Errorf("(validate) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		orderProjection, err := h.os.Queries.GetOrderByID.Handle(ctx, query)
		if err != nil {
			h.log.Errorf("(GetOrderByID.Handle) id: {%s}, err: {%v}", orderID.String(), err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.log.Infof("(get order by id) orderID: {%s}", orderID.String())
		return c.JSON(http.StatusOK, orderProjection)
	}
}

// Search
// @Tags Orders
// @Summary Search orders
// @Description Full text search by title and description
// @Accept json
// @Produce json
// @Param search query string false "search text"
// @Param page query string false "page number"
// @Param size query string false "number of elements"
// @Success 200 {object} orderService.SearchRes
// @Router /orders/search [get]
func (h *orderHandlers) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.StartHttpServerTracerSpan(c, "orderHandlers.Search")
		defer span.Finish()
		h.metrics.SearchOrderHttpRequests.Inc()

		pq := utils.NewPaginationFromQueryParams(c.QueryParam(constants.Size), c.QueryParam(constants.Page))

		query := queries.NewSearchOrdersQuery(c.QueryParam(constants.Search), pq)
		if err := h.v.StructCtx(ctx, query); err != nil {
			h.log.Errorf("(validate) err: {%v}", err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		searchRes, err := h.os.Queries.SearchOrders.Handle(ctx, query)
		if err != nil {
			h.log.Errorf("(SearchOrders.Handle): Search: {%s}, err: {%v}", c.QueryParam(constants.Search), err)
			tracing.TraceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.log.Infof("(search) result: {%s}", searchRes.GetPagination().String())
		return c.JSON(http.StatusOK, searchRes)
	}
}
