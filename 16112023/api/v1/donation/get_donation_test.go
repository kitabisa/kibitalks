package donation

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/kitabisa/kibitalk/config/cache"
	"github.com/kitabisa/kibitalk/config/cache/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDonationByIdCacheHit(t *testing.T) {
	// Create a request with a sample query parameter
	req, err := http.NewRequest("GET", "/v1/donation/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rw := httptest.NewRecorder()

	cacheMock := &mocks.ICache{}
	byteRes, _ := json.Marshal(&CreateDonationResponse{
		Id:            1,
		Amount:        50000,
		PaymentMethod: "Kantong Donasi",
		Campaign:      "Bantuan Kemanusiaan Gaza",
	})

	cacheMock.EXPECT().Get(mock.Anything, mock.Anything).Return(byteRes, nil)

	cache.ClientInstance = cacheMock

	r := chi.NewRouter()
	r.Mount("/v1", V1DonationRoutes())

	r.ServeHTTP(rw, req)

	assert.Equal(t, 200, rw.Code)
}
