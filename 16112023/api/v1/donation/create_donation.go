package donation

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kitabisa/kibitalk/client/campaign"
	"github.com/kitabisa/kibitalk/client/payment"
	"github.com/kitabisa/kibitalk/config/broker/rabbitmq"
	"github.com/kitabisa/kibitalk/config/cache"
	"github.com/kitabisa/kibitalk/config/database"
	zlog "github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func CreateDonationHandler(rw http.ResponseWriter, r *http.Request) {
	var req CreateDonationRequest
	ctx := zlog.With().Str("request_id", r.Context().Value(middleware.RequestIDKey).(string)).Logger().WithContext(r.Context())

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		zlog.Ctx(ctx).Err(err).Msgf("Error occurred during decoding")
		byteRes, _ := json.Marshal(Error{Error: "please validate your input"})
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(byteRes)
		return
	}

	zlog.Ctx(ctx).Info().Msgf("Creating Donation . .")

	campaignResponse, err := campaign.CampaignService.GetCampaign(ctx, req.CampaignId)
	if err != nil {
		zlog.Ctx(ctx).Err(err).Msgf("Error occurred during getting campaign")
		byteRes, _ := json.Marshal(Error{Error: err.Error()})
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(byteRes)
		return
	}

	paymentResp, err := payment.PaymentService.CreatePayment(ctx, payment.CreatePaymentRequest{
		Amount:          req.Amount,
		PaymentMethodId: req.PaymentMethodId,
	})

	if err != nil {
		zlog.Ctx(ctx).Err(err).Msgf("Error occurred during creating payment")
		byteRes, _ := json.Marshal(Error{Error: err.Error()})
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(byteRes)
		return
	}

	res, err := database.MySqlDB.Exec(ctx, "INSERT INTO donations (amount, payment_method_id, campaign_id) VALUES (?, ?, ?)", req.Amount, req.PaymentMethodId, req.CampaignId)
	if err != nil {
		zlog.Ctx(ctx).Err(err).Msgf("Error creating donation")
		byteRes, _ := json.Marshal(Error{Error: err.Error()})
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(byteRes)
		return
	}

	var resp CreateDonationResponse
	donationId, _ := res.LastInsertId()
	resp.Id = donationId
	resp.Amount = paymentResp.Amount
	resp.PaymentMethod = paymentResp.PaymentMethodName
	resp.Campaign = campaignResponse.Name

	errCache := cache.ClientInstance.Set(ctx, fmt.Sprintf("%s:%d", "cache_donation_id", donationId), resp, 600*time.Second)
	if errCache != nil {
		zlog.Ctx(ctx).Err(errCache).Msgf("Error creating donation cache")
	}

	err = rabbitmq.RabbitPublish.Publish("donation_created", fmt.Sprintf("%d", donationId))
	if err != nil {
		zlog.Ctx(ctx).Err(err).Msgf("Error publishing donation")
	}

	cacheResult, _ := json.Marshal(resp)

	zlog.Ctx(ctx).Info().Interface("donation_data", resp).Msgf("Donation Created")

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(cacheResult)
	rw.WriteHeader(http.StatusOK)
}
