package tatekae

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

type TatekaeHandler struct{}

func NewTatekaeHandler() *TatekaeHandler {
	return &TatekaeHandler{}
}

func (t TatekaeHandler) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 型アサーションで []interface{} を取得
	rawMembers, ok := request.Params.Arguments["members"].([]interface{})
	if !ok {
		return nil, errors.New("members must be an array")
	}

	// []interface{} を []string に変換
	members := make([]string, len(rawMembers))
	for i, v := range rawMembers {
		str, ok := v.(string)
		if !ok {
			return nil, errors.New("all members must be strings")
		}
		members[i] = str
	}

	params := &TatekaeParams{
		Members: members,
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", params.Members)), nil
}
