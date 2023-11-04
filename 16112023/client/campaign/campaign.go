package campaign

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kitabisa/kibitalk/config"
	"io"
	"net/http"
)

var CampaignService ICampaignService

type CampaignClient struct {
	Host string
	Port int
	*http.Client
}

func (p CampaignClient) GetCampaign(ctx context.Context, id uint64) (response GetCampaignResponse, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s:%d/campaign/%d", p.Host, p.Port, id), nil)
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
		return
	}

	return

}

type GetCampaignResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ICampaignService interface {
	GetCampaign(ctx context.Context, id uint64) (response GetCampaignResponse, err error)
}

func InitCampaignClient() {
	CampaignService = NewCampaignClient()
}

func NewCampaignClient() ICampaignService {
	c := config.AppCfg

	client := &http.Client{}
	return &CampaignClient{
		Host:   c.ApiCampaign.Host,
		Port:   c.ApiCampaign.Port,
		Client: client,
	}
}
