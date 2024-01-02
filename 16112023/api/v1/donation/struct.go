package donation

type CreateDonationRequest struct {
	Amount          float64 `json:"amount"`
	PaymentMethodId uint64  `json:"payment_method_id"`
	CampaignId      uint64  `json:"campaign_id"`
}

type CreateDonationResponse struct {
	Id            int64   `json:"id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	Campaign      string  `json:"campaign"`
}

type Donation struct {
	Id              int64   `db:"id"`
	Amount          float64 `db:"amount"`
	PaymentMethodId string  `db:"payment_method_id"`
	CampaignId      string  `db:"campaign_id"`
}

type Error struct {
	Error string `json:"error"`
}
