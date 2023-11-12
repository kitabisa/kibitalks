package donation

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kitabisa/kibitalk/config/cache"
	"github.com/kitabisa/kibitalk/config/database"
	zlog "github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"time"
)

func GetDonationByIdHandler(rw http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	ctx := zlog.With().Str("request_id", r.Context().Value(middleware.RequestIDKey).(string)).Logger().WithContext(r.Context())

	// Convert the ID parameter to an integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(rw, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	var resp CreateDonationResponse
	cacheResult, err := cache.ClientInstance.Get(ctx, fmt.Sprintf("%s:%d", "cache_donation_id", id))
	if err == nil {
		zlog.Ctx(ctx).Info().Msgf("Get donation from cache")
		err = json.Unmarshal(cacheResult, &resp)
		if err != nil {
			zlog.Ctx(ctx).Err(err).Msgf("Error unmarshalling cache")
			byteRes, _ := json.Marshal(Error{Error: err.Error()})
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(byteRes)
			return
		}
	} else {
		// cache not found, get from DB
		resp, err = getDonationFromDB(ctx, uint64(id))
		if err != nil {
			zlog.Ctx(ctx).Err(err).Msgf("Error getting donation from DB")
			byteRes, _ := json.Marshal(Error{Error: err.Error()})
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(byteRes)
			return
		}
		cacheResult, _ = json.Marshal(resp)
	}

	zlog.Ctx(ctx).Info().Interface("donation_data", cacheResult).Msgf("Donation Found")

	rw.Write(cacheResult)
	rw.WriteHeader(http.StatusOK)

}

func getDonationFromDB(ctx context.Context, id uint64) (resp CreateDonationResponse, err error) {
	var donation Donation
	zlog.Ctx(ctx).Info().Msgf("Get donation from database")
	err = database.MySqlDB.QueryRow(ctx, "SELECT * FROM donations WHERE id = ?", id).Scan(&donation.Id, &donation.Amount, &donation.PaymentMethodId, &donation.CampaignId)
	if err != nil {
		zlog.Ctx(ctx).Err(err).Msgf("Error getting donation")
		return
	}

	errCache := cache.ClientInstance.Set(ctx, fmt.Sprintf("%s:%d", "cache_donation_id", donation.Id), resp, 600*time.Second)
	if errCache != nil {
		zlog.Ctx(ctx).Err(errCache).Msgf("Error set donation cache")
	}

	resp.Id = donation.Id
	resp.Amount = donation.Amount
	resp.PaymentMethod = donation.PaymentMethodId
	resp.Campaign = donation.CampaignId
	return
}
