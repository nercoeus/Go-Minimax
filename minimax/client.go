package minimax

import (
	"log"
	"os"

	"github.com/imroc/req/v3"
	"github.com/moul/http2curl"
)

const minimaxBaseUrl = "https://api.minimax.chat"
const minimaxBaseApiPath = "/v1/text/chatcompletion"

type Client struct {
	client   *req.Client
	apiToken string
	groupId  string
	apiPath  string
}

func New(opts ...Option) (*Client, error) {
	cli := &Client{
		client:  req.C().SetBaseURL(minimaxBaseUrl),
		apiPath: minimaxBaseApiPath,
	}
	//curl --location "https://api.minimax.chat/v1/text/chatcompletion?GroupId=${group_id}" \

	cli.client.OnBeforeRequest(func(client *req.Client,
		req *req.Request) error {
		if len(cli.apiToken) > 0 {
			req.SetHeader("Authorization", cli.apiToken)
		}
		req.SetHeader("Content-Type", "application/json")
		if len(cli.groupId) > 0 {
			req.SetQueryParam("GroupId", cli.groupId)
		}
		return nil
	}).OnAfterResponse(func(client *req.Client, resp *req.Response) error {
		if os.Getenv("APP_ENV") == "debug" {
			curl, err := http2curl.GetCurlCommand(resp.Request.RawRequest)
			if err != nil {
				return err
			}
			log.Println(curl)
		}
		return nil
	})
	for _, opt := range opts {
		err := opt(cli)
		if err != nil {
			return nil, err
		}
	}
	return cli, nil
}
