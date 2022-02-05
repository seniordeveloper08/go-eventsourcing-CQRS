package aggregate

import (
	"context"
	"encoding/json"
	"github.com/AleksK1NG/es-microservice/internal/order/events"
	serviceErrors "github.com/AleksK1NG/es-microservice/pkg/service_errors"
	"github.com/AleksK1NG/es-microservice/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

func (a *OrderAggregate) CreateOrder(ctx context.Context, command *CreateOrderCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "OrderAggregate.CreateOrder")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	if a.Order.Created {
		return serviceErrors.ErrAlreadyCreated
	}

	if command.OrderCreatedEventData.ShopItems == nil {
		return serviceErrors.ErrOrderItemsIsRequired
	}

	createdData := &events.OrderCreatedEventData{ShopItems: command.ShopItems, AccountEmail: command.AccountEmail, DeliveryAddress: command.DeliveryAddress}
	createdDataBytes, err := json.Marshal(createdData)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "json.Marshal")
	}

	createOrderEvent := events.NewCreateOrderEvent(a, createdDataBytes)
	if err := createOrderEvent.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(createOrderEvent)
}

func (a *OrderAggregate) PayOrder(ctx context.Context, command *OrderPaidCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "OrderAggregate.PayOrder")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	if !a.Order.Created || a.Order.Canceled {
		return serviceErrors.ErrAlreadyCreatedOrCancelled
	}
	if a.Order.Paid {
		return serviceErrors.ErrAlreadyPaid
	}
	if a.Order.Submitted {
		return serviceErrors.ErrAlreadySubmitted
	}

	payOrderEvent := events.NewPayOrderEvent(a, nil)
	if err := payOrderEvent.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(payOrderEvent)
}

func (a *OrderAggregate) SubmitOrder(ctx context.Context, command *SubmitOrderCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "OrderAggregate.SubmitOrder")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	if !a.Order.Created || a.Order.Canceled {
		return serviceErrors.ErrAlreadyCreatedOrCancelled
	}
	if !a.Order.Paid {
		return serviceErrors.ErrOrderNotPaid
	}
	if a.Order.Submitted {
		return serviceErrors.ErrAlreadySubmitted
	}

	submitOrderEvent := events.NewSubmitOrderEvent(a)
	if err := submitOrderEvent.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(submitOrderEvent)
}

func (a *OrderAggregate) UpdateOrder(ctx context.Context, command *OrderUpdatedCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "OrderAggregate.UpdateOrder")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	if !a.Order.Created || a.Order.Canceled {
		return serviceErrors.ErrAlreadyCreatedOrCancelled
	}
	if a.Order.Submitted {
		return serviceErrors.ErrAlreadySubmitted
	}

	eventData := &events.OrderUpdatedEventData{ShopItems: command.ShopItems}
	eventDataBytes, err := json.Marshal(eventData)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "json.Marshal")
	}

	orderUpdatedEvent := events.NewOrderUpdatedEvent(a, eventDataBytes)
	if err := orderUpdatedEvent.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(orderUpdatedEvent)
}

func (a *OrderAggregate) CancelOrder(ctx context.Context, command *OrderCanceledCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "OrderAggregate.CancelOrder")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	eventData := &events.OrderCanceledEventData{CancelReason: command.CancelReason}
	eventDataBytes, err := json.Marshal(eventData)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "json.Marshal")
	}

	event := events.NewOrderCanceledEvent(a, eventDataBytes)
	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}

func (a *OrderAggregate) DeliverOrder(ctx context.Context, command *OrderDeliveredCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "OrderAggregate.DeliverOrder")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	eventData := &events.OrderDeliveredEventData{DeliveryTimestamp: command.DeliveryTimestamp}
	eventDataBytes, err := json.Marshal(eventData)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "json.Marshal")
	}

	event := events.NewOrderDeliveredEvent(a, eventDataBytes)
	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}

func (a *OrderAggregate) ChangeDeliveryAddress(ctx context.Context, command *OrderChangeDeliveryAddressCommand) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "OrderAggregate.ChangeDeliveryAddress")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	if a.Order.Delivered {
		return errors.New("Order already delivered")
	}

	eventData := &events.OrderChangeDeliveryAddress{DeliveryAddress: command.DeliveryAddress}
	eventDataBytes, err := json.Marshal(eventData)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "json.Marshal")
	}

	event := events.NewOrderDeliveryAddressUpdatedEvent(a, eventDataBytes)
	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
