package donation

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kitabisa/kibitalk/client/campaign"
	"github.com/kitabisa/kibitalk/client/payment"
	"github.com/kitabisa/kibitalk/config/broker/rabbitmq"
	"github.com/kitabisa/kibitalk/config/cache"
	"github.com/kitabisa/kibitalk/config/database"
	plog "github.com/kitabisa/perkakas/log"
	"github.com/kitabisa/perkakas/log/ctxkeys"
	"net/http"
	"time"
)

func CreateDonationHandler(rw http.ResponseWriter, r *http.Request) {
	var req CreateDonationRequest
	ctx := r.Context()
	if ctx.Value(middleware.RequestIDKey) != nil {
		logger := plog.Zlogger(ctx).With().Str("request_id", ctx.Value(middleware.RequestIDKey).(string)).Logger()
		ctx = context.WithValue(ctx, ctxkeys.CtxLogger, logger)
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		plog.Zlogger(ctx).Err(err).Msgf("Error occurred during decoding")
		byteRes, _ := json.Marshal(Error{Error: "please validate your input"})
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(byteRes)
		return
	}

	plog.Zlogger(ctx).Info().Msgf("Creating Donation . .")

	campaignResponse, err := campaign.CampaignService.GetCampaign(ctx, req.CampaignId)
	if err != nil {
		plog.Zlogger(ctx).Err(err).Msgf("Error occurred during getting campaign")
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
		plog.Zlogger(ctx).Err(err).Msgf("Error occurred during creating payment")
		byteRes, _ := json.Marshal(Error{Error: err.Error()})
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(byteRes)
		return
	}

	res, err := database.MySqlDB.Exec(ctx, "INSERT INTO donations (amount, payment_method_id, campaign_id) VALUES (?, ?, ?)", req.Amount, req.PaymentMethodId, req.CampaignId)
	if err != nil {
		plog.Zlogger(ctx).Err(err).Msgf("Error creating donation")
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
		plog.Zlogger(ctx).Err(errCache).Msgf("Error creating donation cache")
	}

	err = rabbitmq.RabbitPublish.Publish("donation_created", fmt.Sprintf("%d", donationId))
	if err != nil {
		plog.Zlogger(ctx).Err(err).Msgf("Error publishing donation")
	}

	cacheResult, _ := json.Marshal(resp)

	plog.Zlogger(ctx).Info().Interface("donation_data", resp).Msgf("Donation Created")

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(cacheResult)
	rw.WriteHeader(http.StatusOK)
}
