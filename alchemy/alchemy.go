package alchemy

import (
	"context"
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

// Block tags: https://docs.alchemy.com/reference/transfers-api-quickstart#what-are-the-different-types-of-block-tags

type AlchemyClient struct {
	ctx       context.Context
	rawurl    string
	rpcClient *rpc.Client
}

func NewAlchemyClient(ctx context.Context, rawurl string) (*AlchemyClient, error) {
	rpcClient, err := rpc.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return &AlchemyClient{
		ctx:       ctx,
		rpcClient: rpcClient,
		rawurl:    rawurl,
	}, nil
}

func (ac *AlchemyClient) Reconnect() error {
	// Close the existing connection
	if ac.rpcClient != nil {
		ac.rpcClient.Close()
	}

	// Create a new connection
	rpcClient, err := rpc.Dial(ac.rawurl)
	if err != nil {
		return err
	}

	// Replace the old client with the new one
	ac.rpcClient = rpcClient

	return nil
}

type GetAssetTransfersArgs struct {
	FromBlock         string   `json:"fromBlock"`                   //Inclusive from block (hex string, int, latest, or indexed). Defaults to 0x0.
	ToBlock           string   `json:"toBlock"`                     //Inclusive to block (hex string, int, latest, or indexed). Defaults to latest.
	FromAddress       string   `json:"fromAddress,omitempty"`       //From address (hex string). Default wildcard - any address
	ToAddress         string   `json:"toAddress,omitempty"`         //To address (hex string). Default wildcard - any address
	ContractAddresses []string `json:"contractAddresses,omitempty"` //List of contract addresses (hex strings) to filter for - only applies to "erc20", "erc721", "erc1155" transfers. Default wildcard - any address
	Category          []string `json:"category"`                    //'Array of categories, can be any of the following: "external", "internal", "erc20", "erc721", "erc1155", or "specialnft". See the table above for supported categories on each network.'
	Order             string   `json:"order"`                       //Whether to return results in ascending (asc) or descending (desc) order. Ascending order is from oldest to newest transactions, descending order is from newest to oldest. Defaults to asc.
	WithMetadata      bool     `json:"withMetadata"`                //Whether or not to include additional metadata about each transfer event. Defaults to false.
	ExcludeZeroValue  bool     `json:"excludeZeroValue"`            //A boolean to exclude transfers with zero value - zero value is not the same as null value. Defaults to true.
	MaxCount          string   `json:"maxCount"`                    //Max hex string number of results to return per call. Defaults to 0x3e8 (1000).
	PageKey           string   `json:"pageKey,omitempty"`           //UUID for pagination. If more results are available, a uuid pageKey will be returned in the response. Pass that uuid into pageKey to fetch the next 1000 or maxCount.
}

// NewGetAssetTransfersArgs creates a new instance of GetAssetTransfersArgs with default values.
func NewGetAssetTransfersArgs(params ...func(*GetAssetTransfersArgs)) *GetAssetTransfersArgs {
	// Set default values
	req := &GetAssetTransfersArgs{
		FromBlock:         "0x0",
		ToBlock:           "latest",
		FromAddress:       "",  // Default wildcard
		ToAddress:         "",  // Default wildcard
		ContractAddresses: nil, // Default wildcard
		Category:          []string{"external"},
		Order:             "asc",
		WithMetadata:      false,
		ExcludeZeroValue:  true,
		MaxCount:          "0x3e8",
		PageKey:           "",
	}

	// Override default values with user-defined values
	for _, param := range params {
		param(req)
	}

	return req
}

// SetFromBlock sets the FromBlock field.
func SetFromBlock(value string) func(*GetAssetTransfersArgs) {
	return func(req *GetAssetTransfersArgs) {
		req.FromBlock = value
	}
}

// SetToBlock sets the ToBlock field.
func SetToBlock(value string) func(*GetAssetTransfersArgs) {
	return func(req *GetAssetTransfersArgs) {
		req.ToBlock = value
	}
}

// SetFromAddress sets the FromAddress field.
func SetFromAddress(value string) func(*GetAssetTransfersArgs) {
	return func(req *GetAssetTransfersArgs) {
		req.FromAddress = value
	}
}

func SetToAddress(value string) func(*GetAssetTransfersArgs) {
	return func(req *GetAssetTransfersArgs) {
		req.ToAddress = value
	}
}

func SetContractAddresses(value []string) func(*GetAssetTransfersArgs) {
	return func(req *GetAssetTransfersArgs) {
		req.ContractAddresses = value
	}
}

func SetCategory(value []string) func(*GetAssetTransfersArgs) {
	return func(req *GetAssetTransfersArgs) {
		req.Category = value
	}
}

func SetOrder(value string) func(*GetAssetTransfersArgs) {
	return func(req *GetAssetTransfersArgs) {
		req.Order = value
	}
}

func SetWithMetadata(value bool) func(*GetAssetTransfersArgs) {
	return func(req *GetAssetTransfersArgs) {
		req.WithMetadata = value
	}
}

func SetExcludeZeroValue(value bool) func(*GetAssetTransfersArgs) {
	return func(req *GetAssetTransfersArgs) {
		req.ExcludeZeroValue = value
	}
}

func SetMaxCount(value string) func(*GetAssetTransfersArgs) {
	return func(req *GetAssetTransfersArgs) {
		req.MaxCount = value
	}
}

type RawContract struct {
	Value   string `json:"value"`
	Address string `json:"address"`
	Decimal string `json:"decimal"`
}

type AssetTransfer struct {
	BlockNum        string      `json:"blockNum"`
	UniqueId        string      `json:"uniqueId"`
	Hash            string      `json:"hash"`
	From            string      `json:"from"`
	To              string      `json:"to"`
	Value           float64     `json:"value"`
	ERC721TokenId   string      `json:"erc721TokenId"`
	ERC1155Metadata string      `json:"erc1155Metadata"`
	TokenId         string      `json:"tokenId"`
	Asset           string      `json:"asset"`
	Category        string      `json:"category"`
	Contract        RawContract `json:"rawContract"`
}
type GetAssetTransfersResponse struct {
	Transfers []AssetTransfer `json:"transfers"`
	PageKey   string          `json:"pageKey"`
}

func (ac *AlchemyClient) GetAllAssetTransfers(ctx context.Context, params *GetAssetTransfersArgs) ([]AssetTransfer, error) {
	transfers := make([]AssetTransfer, 0)
	firstCall := true

	for params.PageKey != "" || firstCall {
		newTransfers, newPageKey, err := ac.GetAssetTransfers(ctx, params)
		if err != nil {
			return nil, err
		}
		params.PageKey = newPageKey
		firstCall = false
		transfers = append(newTransfers, transfers...)
	}
	return transfers, nil
}

func (ac *AlchemyClient) GetAssetTransfers(ctx context.Context, params *GetAssetTransfersArgs) ([]AssetTransfer, string, error) {
	var raw json.RawMessage
	retry := 0
	for {
		err := ac.rpcClient.CallContext(ctx, &raw, "alchemy_getAssetTransfers", params)
		if err != nil {
			if !strings.Contains(err.Error(), "429") {
				retry++
				ac.Reconnect()
			}
			if retry > 5 {
				return nil, "", errors.Wrap(err, "alchemy_getAssetTransfers failed after 5 retries")
			}
			waitTime := time.Duration(rand.Intn(250)+1000) * time.Millisecond
			time.Sleep(waitTime)
			continue
		}
		break
	}

	var response GetAssetTransfersResponse
	err := json.Unmarshal(raw, &response)
	if err != nil {
		return nil, "", err
	}
	return response.Transfers, response.PageKey, nil
}
