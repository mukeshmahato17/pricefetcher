package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mukeshmahato17/pricefetcher/proto"
	"github.com/mukeshmahato17/pricefetcher/types"
	"google.golang.org/grpc"
)

func NewGRPCClient(remoteAddr string) (proto.PriceFetcherClient, error) {
	conn, err := grpc.NewClient(remoteAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := proto.NewPriceFetcherClient(conn)
	return c, nil
}

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)

	req, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}

		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("service respond with non OK status code: %s", httpErr["error"])
	}

	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}

	return priceResp, nil
}
