# mcp-server-today MCP Server

This server provides tools to get today's, tomorrow's, and yesterday's date, as well as the current time with timezone information.

## Prerequisites

- Go (version 1.20 or higher recommended)

## Getting Started

1. Clone the repository (if applicable).
2. Navigate to the project directory.
3. Run `go mod tidy` to install dependencies if you haven't already or if `go.mod` changed.

## Building the Server

To compile the server, run the following command in the project root:

```bash
go build -o mcp_server_today main.go
```

This will create an executable named `mcp_server_today` in the project's root directory.

## Running the Server

This executable is an MCP server. It's designed to be run by an MCP client (e.g., Claude Desktop). You can configure your MCP client to run the compiled `mcp_server_today` executable or use `go run main.go` for development.

### Example `.cursor/mcp.json` configuration for Claude Desktop:

To use with Claude Desktop, you can create or update your `claude_desktop_config.json` file (usually found in `~/Library/Application Support/Claude/` on macOS or `%APPDATA%\Claude\` on Windows).

**Option 1: Using the compiled executable**
```json
{
  "mcpServers": {
    "mcp-server-today-compiled": {
      "command": ["/full/path/to/mcp-server-today/mcp_server_today"],
      "workingDirectory": "/full/path/to/mcp-server-today"
    }
  }
}
```
**Note:** Replace `/full/path/to/mcp-server-today` with the actual absolute path to your project directory.

**Option 2: Running with `go run` (for development)**
```json
{
  "mcpServers": {
    "mcp-server-today-dev": {
      "command": ["go", "run", "main.go"],
      "debugCommand": ["dlv", "debug", "main.go", "--", "stdio"], // Optional: for debugging with Delve
      "workingDirectory": "/full/path/to/mcp-server-today"
    }
  }
}
```
**Note:** Ensure Go is in your system's PATH. Replace `/full/path/to/mcp-server-today` with the actual absolute path.

## Tools

This MCP server exposes the following tools:

- **`today`**:
    - Description: Returns the current date.
    - Output Format: `MM-DD-YYYY`
- **`tomorrow`**:
    - Description: Returns tomorrow's date.
    - Output Format: `MM-DD-YYYY`
- **`yesterday`**:
    - Description: Returns yesterday's date.
    - Output Format: `MM-DD-YYYY`
- **`time`**:
    - Description: Returns the current time with timezone.
    - Parameters:
        - `use_24_hour_format` (boolean, optional): If `true`, returns time in 24-hour format (e.g., `15:04:05 MST`). Defaults to `false` (12-hour format, e.g., `03:04:05 PM MST`).
    - Output Format: Time string with timezone (e.g., `03:04:05 PM EST` or `15:04:05 EST`).

## Development Best Practices

### .gitignore

Ensure you have a comprehensive `.gitignore` file in your project directory before committing to a code repository. This helps prevent sensitive information and unnecessary files (like build artifacts or local editor configurations) from being tracked.

A typical Go project `.gitignore` might include:
```
# Compiled Object files, Static Libraries, Shared Libraries and Executables
*.o
*.a
*.so
*.exe
*.dll
mcp_server_today # Or your specific executable name

# Go build cache
*.test
*.prof

# IDE / Editor specific files
.vscode/
.idea/
*.swp
*~
.cursor/
```

### Security

- **Never commit API Keys, Secrets, or Private Keys** to a code repository.
- Instead, make sure they are managed through environment variables or secure `.env` files (which should also be listed in `.gitignore`).
- While this specific server currently doesn't handle external API keys, it's a crucial practice for any project.

## Contributing

(Placeholder for contributing guidelines) 