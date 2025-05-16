package main

import (
	"log"

	"mcp-server-today/tools"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer("mcp-server-today", "0.1.0")

	s.AddTool(tools.TodayTool, tools.HandleToday)
	s.AddTool(tools.TomorrowTool, tools.HandleTomorrow)
	s.AddTool(tools.YesterdayTool, tools.HandleYesterday)
	s.AddTool(tools.TimeTool, tools.HandleTime)

	log.Println("Starting MCP server...")
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Error serving MCP server: %v", err)
	}
}
