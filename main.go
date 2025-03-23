package main

import (
	"fmt"

	"github.com/ei-sugimoto/mendou/internal/pkg/hello"
	"github.com/ei-sugimoto/mendou/internal/pkg/tatekae"
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

	tatekaeTool := mcp.NewTool("tatekae",
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

	helloHandler := hello.NewHelloHandler()
	tatekaeHandler := tatekae.NewTatekaeHandler()

	// Add tool handler
	s.AddTool(tool, helloHandler.Handle)
	s.AddTool(tatekaeTool, tatekaeHandler.Handle)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
