package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kitabisa/kibitalk/config"
	"io"
	"net/http"
)

var PaymentService IPaymentService

type PaymentClient struct {
	Host string
	Port int
	*http.Client
}

func (p PaymentClient) CreatePayment(ctx context.Context, request CreatePaymentRequest) (response CreatePaymentResponse, err error) {
	var reqBody []byte
	reqBody, err = json.Marshal(request)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s:%d/payment", p.Host, p.Port), bytes.NewReader(reqBody))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	resp, err = p.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New("something error occured")
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return CreatePaymentResponse{}, err
	}

	return

}

type CreatePaymentRequest struct {
	Amount          float64 `json:"amount"`
	PaymentMethodId uint64  `json:"payment_method"`
}

type CreatePaymentResponse struct {
	Id                int64   `json:"id"`
	PaymentMethodName string  `json:"payment_method_name"`
	Amount            float64 `json:"amount"`
}

type IPaymentService interface {
	CreatePayment(ctx context.Context, request CreatePaymentRequest) (CreatePaymentResponse, error)
}

func InitPaymentClient() {
	PaymentService = NewPaymentClient()
}

func NewPaymentClient() IPaymentService {
	c := config.AppCfg

	client := &http.Client{}
	return &PaymentClient{
		Host:   c.ApiPayment.Host,
		Port:   c.ApiPayment.Port,
		Client: client,
	}
}
