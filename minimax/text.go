package minimax

import (
	"context"
	"errors"
	"io"

	textv1 "github.com/nercoeus/go-minimax/gen/go/minimax/text/v1"
	"github.com/nercoeus/go-minimax/internal"
	"google.golang.org/grpc"
)

var _ textv1.MinimaxServiceClient = new(Client)

func (cli *Client) ChatCompletions(ctx context.Context, in *textv1.ChatCompletionsRequest, opts ...grpc.CallOption) (*textv1.ChatCompletionsResponse, error) {
	res := new(struct {
		textv1.ChatCompletionsResponse
	})

	in.Stream = false
	resp, err := cli.client.R().
		SetBody(in).
		SetSuccessResult(res).
		Post(cli.apiPath)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, err
	}
	res.ChatCompletionsResponse.TraceId = resp.Header.Get("Trace-Id")
	return &res.ChatCompletionsResponse, err
}

func (cli *Client) ChatCompletionStream(ctx context.Context, in *textv1.ChatCompletionsRequest, opts ...grpc.CallOption) (textv1.MinimaxService_ChatCompletionStreamClient, error) {

	in.Stream = true
	in.UseStandardSse = true
	resp, err := cli.client.R().
		DisableAutoReadResponse().
		SetBody(in).
		Post(cli.apiPath)

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}
	result := internal.NewStreamReader[*textv1.ChatCompletionsResponse](resp.Body)
	result.SetHeader("Trace-Id", resp.Header.Get("Trace-Id"))
	return result, err
}
