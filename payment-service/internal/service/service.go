package service

import (
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/models"
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/repository"
	elk "github.com/DurkaVerder/elk-send-logs/elk"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
)

type Service interface {
	CreatePaymentIntent(order models.Order) (*stripe.PaymentIntent, error)
}

type PaymentService struct {
	repo repository.Repository
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	stripe.Key = os.Getenv("PAYMENT_GATEWAY_SECRET")
}

func NewPaymentService(repo repository.Repository) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}

func (s *PaymentService) CreatePaymentIntent(order models.Order) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(int64(order.Amount)),
		Currency:           stripe.String(string(stripe.CurrencyRUB)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Failed to create payment intent",
			Fields: map[string]interface{}{
				"method":   "CreatePaymentIntent",
				"action":   "create_payment_intent",
				"amount":   order.Amount,
				"user_id":  order.UserID,
				"order_id": order.ID,
				"error":    err.Error(),
			},
		})
		return nil, err
	}

	if err := s.repo.AddPayment(order, pi.ID); err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Failed to add payment to database",
			Fields: map[string]interface{}{
				"method":   "CreatePaymentIntent",
				"action":   "create_payment_intent",
				"amount":   order.Amount,
				"user_id":  order.UserID,
				"order_id": order.ID,
				"error":    err.Error(),
			},
		})

		return nil, err
	}

	elk.Log.SendMsg(elk.LogMessage{
		Level:   'I',
		Message: "Payment intent created",
		Fields: map[string]interface{}{
			"method":   "CreatePaymentIntent",
			"action":   "create_payment_intent",
			"amount":   order.Amount,
			"user_id":  order.UserID,
			"order_id": order.ID,
		},
	})

	return pi, nil
}
