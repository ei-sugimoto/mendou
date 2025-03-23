package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"1.0.0",
	)

	// Add tool
	tool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)

	tatekae := mcp.NewTool("tatekae",
		mcp.WithDescription("tatekaeは、立替による取引回数を最適化するためのツールです"),
		mcp.WithArray("members", mcp.Items(map[string]interface{}{
			"type": "string",
		})),
		mcp.WithArray("transactions", mcp.Items(map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"from": map[string]interface{}{
					"type": "string",
				},
				"to": map[string]interface{}{
					"type": "string",
				},
				"amount": map[string]interface{}{
					"type": "number",
				},
			},
		})),
	)

	// Add tool handler
	s.AddTool(tool, helloHandler)
	s.AddTool(tatekae, tatekaeHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.Params.Arguments["name"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}

func tatekaeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", members)), nil
}
