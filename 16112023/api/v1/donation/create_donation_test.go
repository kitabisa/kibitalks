package donation

import (
	"github.com/go-chi/chi/v5"
	"github.com/kitabisa/kibitalk/client/campaign"
	mocks3 "github.com/kitabisa/kibitalk/client/campaign/mocks"
	"github.com/kitabisa/kibitalk/client/payment"
	mocks4 "github.com/kitabisa/kibitalk/client/payment/mocks"
	"github.com/kitabisa/kibitalk/config/broker/rabbitmq"
	mocks5 "github.com/kitabisa/kibitalk/config/broker/rabbitmq/mocks"
	"github.com/kitabisa/kibitalk/config/cache"
	"github.com/kitabisa/kibitalk/config/cache/mocks"
	"github.com/kitabisa/kibitalk/config/database"
	mocks2 "github.com/kitabisa/kibitalk/config/database/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateDonationByIdHandler(t *testing.T) {
	// Create a request with a sample query parameter
	req, err := http.NewRequest("POST", "/v1/donation", strings.NewReader("{\"amount\": 50000,\"payment_method_id\": 1,\"campaign_id\": 1}"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rw := httptest.NewRecorder()

	// mock campaign
	var campaignService mocks3.ICampaignService
	campaignService.EXPECT().GetCampaign(mock.Anything, mock.Anything).Return(campaign.GetCampaignResponse{
		Id:   1,
		Name: "Bantuan Kemanusiaan Gaza",
	}, nil)

	campaign.CampaignService = &campaignService

	// mock payment
	var paymentService mocks4.IPaymentService
	paymentService.EXPECT().CreatePayment(mock.Anything, mock.Anything).Return(payment.CreatePaymentResponse{
		Id:                1,
		PaymentMethodName: "Kantong Donasi",
		Amount:            50000,
	}, nil)

	payment.PaymentService = &paymentService

	// mock cache
	cacheMock := &mocks.ICache{}
	cacheMock.EXPECT().Set(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	cache.ClientInstance = cacheMock

	// mock db exec
	var sqlResult SqlResult
	dbMock := &mocks2.IMySQL{}
	dbMock.EXPECT().Exec(mock.Anything, mock.Anything, float64(50000), uint64(1), uint64(1)).Return(sqlResult, nil)
	database.MySqlDB = dbMock

	// mock rabbitmq publish
	var rabbitMQMock mocks5.IRabbitPublisher
	rabbitMQMock.EXPECT().Publish(mock.Anything, mock.Anything).Return(nil)
	rabbitmq.RabbitPublish = &rabbitMQMock

	r := chi.NewRouter()
	r.Mount("/v1", V1DonationRoutes())

	r.ServeHTTP(rw, req)

	assert.Equal(t, 200, rw.Code)
}

type SqlResult struct {
}

func (s SqlResult) RowsAffected() (int64, error) {
	return 1, nil
}

func (s SqlResult) LastInsertId() (int64, error) {
	return 1, nil
}
