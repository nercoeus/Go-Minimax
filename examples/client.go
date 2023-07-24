package main

import (
	"context"
	"fmt"

	textv1 "github.com/nercoeus/go-minimax/gen/go/minimax/text/v1"
	"github.com/nercoeus/go-minimax/minimax"
)

func main() {
	ctx := context.Background()

	//init client
	client, _ := minimax.New(
		minimax.WithApiToken("Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoibHVjeeWvhumSpSIsIlN1YmplY3RJRCI6IjE2ODk2NTExNDc4MDA5NTUiLCJQaG9uZSI6Ik1UZzFNVGcwT0RVeU1UST0iLCJHcm91cElEIjoiMTY4OTY1MTE0NzA0NzY4MCIsIlBhZ2VOYW1lIjoiIiwiTWFpbCI6InRlbmNlbnRfYXp1cmVfcXFhaUB0ZW5jZW50LmNvbSIsIkNyZWF0ZVRpbWUiOiIyMDIzLTA3LTI0IDEyOjU4OjM3IiwiaXNzIjoibWluaW1heCJ9.Ln6fLQ-tPdNsddhH9eb3InIsfhw6Bg0w3EuPR32PBTCCT2rt50i-XdQSgg-zq-glnZDWXehZw8UW0Ju073X05hnTvwtRj8FqUn5HG0jd9ESq6L8YYiHUz-XBVzpUFIk1tuLJPYO782tf6kB6zL_o7eFC2k2Gmdn7xxKXcC9YYsIJ7ZFWQixNXbXZOCIvFCbP_3ayCelBNac3kspDCwyDLMNwGL6RD1GzJ88Agw2qQW_SrVUyxY2nxSIh7IzYd1lrnQxYB9csNqqFFzHTRfJd1p7bDhZOL0U6q4aj5U2olwgjrF-bswlgpb_ZC1GovStIrokKCiXYeaQz-Tkjmx_6Qg"),
		minimax.WithGroupId("1689651147047680"),
		minimax.WithApiPath("/v1/text/chatcompletion_pro"),
	)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
				SenderType: "USER",
				SenderName: "user",
				Text:       "你好",
			},
			{
				SenderType: "BOT",
				SenderName: "MM助手",
				Text:       "你好",
			},
			{
				SenderType: "USER",
				SenderName: "user",
				Text:       "明天上海天气怎么样?",
			},
		},
		BotSetting: []*textv1.BotSetting{
			{
				BotName: "MM助手",
				Content: "MM智能助理是MiniMax自主研发的大型语言模型，回答问题简洁有条理，没有调用其他产品接口。MiniMax是一家中国科技公司，致力于大模型相关的研究。",
			},
		},
		ReplyConstraints: &textv1.ReplyConstraints{
			SenderType: "BOT",
			SenderName: "MM助手",
		},
		Model:       "abab5.5-chat",
		Temperature: 0.7,
		Functions: []*textv1.FunctionDefinition{
			{
				Name:        "get_weather",
				Description: "获得天气信息",
				Parameters: &textv1.Definition{
					Type: "object",
					Properties: map[string]*textv1.Definition{
						"location": {
							Type:        "string",
							Description: "要获得天气的地点",
						},
					},
					Required: []string{"location"},
				},
			},
		},
	}
	res, err := client.ChatCompletions(ctx, req)

	fmt.Println(res.Choices[0], err) // output: 你好！有什么我可以帮助你的吗？
}

func main2() {
	ctx := context.Background()

	//init client
	client, _ := minimax.New(
		minimax.WithApiToken("Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoibHVjeeWvhumSpSIsIlN1YmplY3RJRCI6IjE2ODk2NTExNDc4MDA5NTUiLCJQaG9uZSI6Ik1UZzFNVGcwT0RVeU1UST0iLCJHcm91cElEIjoiMTY4OTY1MTE0NzA0NzY4MCIsIlBhZ2VOYW1lIjoiIiwiTWFpbCI6InRlbmNlbnRfYXp1cmVfcXFhaUB0ZW5jZW50LmNvbSIsIkNyZWF0ZVRpbWUiOiIyMDIzLTA3LTI0IDEyOjU4OjM3IiwiaXNzIjoibWluaW1heCJ9.Ln6fLQ-tPdNsddhH9eb3InIsfhw6Bg0w3EuPR32PBTCCT2rt50i-XdQSgg-zq-glnZDWXehZw8UW0Ju073X05hnTvwtRj8FqUn5HG0jd9ESq6L8YYiHUz-XBVzpUFIk1tuLJPYO782tf6kB6zL_o7eFC2k2Gmdn7xxKXcC9YYsIJ7ZFWQixNXbXZOCIvFCbP_3ayCelBNac3kspDCwyDLMNwGL6RD1GzJ88Agw2qQW_SrVUyxY2nxSIh7IzYd1lrnQxYB9csNqqFFzHTRfJd1p7bDhZOL0U6q4aj5U2olwgjrF-bswlgpb_ZC1GovStIrokKCiXYeaQz-Tkjmx_6Qg"),
		minimax.WithGroupId("1689651147047680"),
	)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
				SenderType: "USER",
				Text:       "hi~",
			},
		},
		Model:       "abab5-chat",
		Temperature: 0.7,
	}
	res, _ := client.ChatCompletions(ctx, req)

	fmt.Println(res.Choices) // output: 你好！有什么我可以帮助你的吗？
}
