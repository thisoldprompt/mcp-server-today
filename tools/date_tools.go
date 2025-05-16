package tools

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

const dateFormat = "01-02-2006"

// TodayTool definition
var TodayTool = mcp.NewTool(
	"today",
	mcp.WithDescription("Returns the current date in MM-DD-YYYY format."),
)

// HandleToday processes the request for the today tool
func HandleToday(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	today := time.Now().Format(dateFormat)
	return mcp.NewToolResultText(fmt.Sprintf("Today's date is %s", today)), nil
}

// TomorrowTool definition
var TomorrowTool = mcp.NewTool(
	"tomorrow",
	mcp.WithDescription("Returns tomorrow's date in MM-DD-YYYY format."),
)

// HandleTomorrow processes the request for the tomorrow tool
func HandleTomorrow(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	tomorrow := time.Now().AddDate(0, 0, 1).Format(dateFormat)
	return mcp.NewToolResultText(fmt.Sprintf("Tomorrow's date is %s", tomorrow)), nil
}

// YesterdayTool definition
var YesterdayTool = mcp.NewTool(
	"yesterday",
	mcp.WithDescription("Returns yesterday's date in MM-DD-YYYY format."),
)

// HandleYesterday processes the request for the yesterday tool
func HandleYesterday(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	yesterday := time.Now().AddDate(0, 0, -1).Format(dateFormat)
	return mcp.NewToolResultText(fmt.Sprintf("Yesterday's date is %s", yesterday)), nil
}

// TimeTool definition
var TimeTool = mcp.NewTool(
	"time",
	mcp.WithDescription("Returns the current time with timezone. Defaults to 12-hour format (e.g., 03:04:05 PM PST). Can use 24-hour format if specified."),
	mcp.WithBoolean("use_24_hour_format",
		mcp.Description("If true, returns time in 24-hour format (e.g., 15:04:05). Defaults to false (12-hour format)."),
	),
)

// HandleTime processes the request for the time tool
func HandleTime(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	use24HourFormat := false
	if val, ok := req.Params.Arguments["use_24_hour_format"].(bool); ok {
		use24HourFormat = val
	}

	format := "03:04:05 PM MST"
	if use24HourFormat {
		format = "15:04:05 MST"
	}

	currentTime := time.Now().Format(format)
	return mcp.NewToolResultText(fmt.Sprintf("The current time is %s", currentTime)), nil
}
