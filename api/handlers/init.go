package handlers

import (
	"api-gateway/config/logger"
	pbp "api-gateway/genproto/payment"
	pbr "api-gateway/genproto/reservation"

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	Reservation pbr.ReservationServiceClient
	Restaurant  pbr.RestaurantServiceClient
	Payment     pbp.PaymentServiceClient
	Logger      logger.Logger
}

func NewHandler(connR, connP *grpc.ClientConn, l logger.Logger) *HTTPHandler {
	return &HTTPHandler{
		Reservation: pbr.NewReservationServiceClient(connR),
		Restaurant:  pbr.NewRestaurantServiceClient(connR),
		Payment:     pbp.NewPaymentServiceClient(connR),
		Logger:      l,
	}
}
