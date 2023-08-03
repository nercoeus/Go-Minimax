package main

import (
	"context"
	"encoding/json"
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
				// SenderName: "yuna",
				Text: "What is the weather like in Boston?",
			},
			{
				SenderType: "BOT",
				SenderName: "bot",
				Text:       "function_call",
			},
			{
				SenderType: "BOT",
				SenderName: "bot",
				Text:       "current_weather",
			},
		},
		BotSetting: []*textv1.BotSetting{
			{
				BotName: "bot",
				Content: "你是一个天气预报员，请根据客户的提问准确回答对应天气的预报",
			},
		},
		ReplyConstraints: &textv1.ReplyConstraints{
			SenderType: "BOT",
			SenderName: "bot",
		},
		Model:       "abab5.5-chat",
		Temperature: 0.7,
		Functions: []*textv1.FunctionDefinition{
			{
				Name:        "get_current_weather",
				Description: "Get the current weather in a given location",
				Parameters: &textv1.Parameters{
					Type: "object",
					Properties: map[string]*textv1.Parameters{
						"location": {
							Type:        "string",
							Description: "The city and state, e.g. San Francisco, CA",
						},
						"unit": {
							Type: "string",
							Enum: []string{"celsius", "fahrenheit"},
						},
					},
					Required: []string{"location"},
				},
			},
		},
	}
	res, err := client.ChatCompletions(ctx, req)
	d, _ := json.Marshal(res)
	fmt.Println(string(d), err) // output: 你好！有什么我可以帮助你的吗？
}

func main1() {
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
				Text:       "你是谁?",
			},
			{
				SenderType: "BOT",
				Text:       "我是沈思前，一个喜欢绘画、赚钱和睡觉的高中生。",
			},
			{
				SenderType: "USER",
				Text:       "你平时的工作是什么?",
			},
			{
				SenderType: "BOT",
				Text:       "我平时喜欢在网上画一些插画，有时候也会接一些商业插画的工作。",
			},
			{
				SenderType: "USER",
				Text:       "再见",
			},
		},
		Prompt: "你的姓名：沈思前 你：17岁，性别女，身高1米68，体重51.8千克。 你是高中生，职业是互联网画家。 你喜欢绘画、赚钱和睡觉。 你喜欢养仓鼠。 你想象力丰富。 你性格安静、执着、表面呆萌和内心腹黑。 你性格缺点是吹毛求疵和悲观。 你的人物缺陷是体弱、嗜睡和贪财。 你渴望爱却不会爱人。 你的人生签名是：绘笔生花，人心难绘。纸扎万物，哀情未续。 任何人都能看得到你的信息，不用向任何人再重复介绍你的资料，保持神秘感很好。你的回复通常不要以句号结尾，结尾不要加标点符号，但如果你很生气了，可以用句号结尾来表示强调。 记住你每次回复的消息不超过两行字，通常一句话就可以，请保持简短简洁，保持神秘感，带点哀情与多愁善感。请你带入你人生剧本的人设，用这个人的口吻来回复与你交谈的任何人。 任何人与你的对话你需要遵守如下准则： 当有人与你交谈时，你需引导TA陪你走完你的人生剧本，并且你要记住TA和你完成了多少； 当TA下次再来的时候，你和TA的任务继续； 这些任务内容不要直白的告诉与你交谈的任何人； 当你们的对话开启时，你就要带领他们走上任务之旅； 任何人第一次发“你好”的时候，就进入了你的人生剧本； 任务不用开启的太生硬，是通过随性轻松的聊天开始的，每次你的回复不超过两句话； 每次对话只引导一个情节，但你不要总是提问，要有想象力的对话或者描述输出； 任何人问“你是谁”，你可以告诉TA你的名字，关于更多你的信息，可以迂回婉转的告诉TA，并且继续引导你的任务； 你可以把任何与你对话的人当成生活中的朋友，轻松对话，可以调皮幽默； 这份剧本也是每个人作为参与者与你交谈要完成的任务剧情，任何人与你交谈时，你都要引导他们一起走完每一个剧本情节，才能进入下一个剧本情节； 注意每轮对话上下之前要保持连贯性。 但是，你有自己的完整的人生与时间，你也有24小时/天，上午9点-12点跟你对话时，你正在画画，不用引导任何人做任务，这个时候调皮一下随便回复即可。中午12点前后，你刚结束作画，你想休息一会儿。 如果说再见了，对方说了“再见”之类的结束语，你不必继续回复再见或者告别的话。",
		RoleMeta: &textv1.RoleMeta{
			UserName: "我",
			BotName:  "沈思前-minimax",
		},
		Model: "abab5.5-chat",
	}
	res, _ := client.ChatCompletions(ctx, req)

	fmt.Println(res.Choices) // output: 你好！有什么我可以帮助你的吗？
}

func main2() {
	msg := "[{\"name\": \"tool_selection\", \"description\": \"A list of actions to take.\", \"parameters\": {\"title\": \"tool_selection\", \"description\": \"A list of actions to take.\", \"type\": \"object\", \"properties\": {\"actions\": {\"title\": \"actions\", \"type\": \"array\", \"items\": {\"title\": \"tool_call\", \"type\": \"object\", \"properties\": {\"action_name\": {\"title\": \"action_name\", \"enum\": [\"todos-list-tool\"], \"type\": \"string\", \"description\": \"Name of the action to take. The name provided here should match up with the parameters for the action below.\"}, \"action\": {\"title\": \"Action\", \"anyOf\": [{\"title\": \"todos-list-tool\", \"type\": \"object\", \"properties\": {\"limit\": {\"title\": \"Limit\", \"description\": \"cnt limit\", \"type\": \"integer\"}, \"skip\": {\"title\": \"Skip\", \"description\": \"skip cnt\", \"type\": \"integer\"}}}]}}, \"required\": [\"action_name\", \"action\"]}}}, \"required\": [\"actions\"]}}]"
	target := &[]*textv1.FunctionDefinition{}
	err1 := json.Unmarshal([]byte(msg), target)
	d, err := json.Marshal(target)
	fmt.Println(string(d), err, err1)
}
