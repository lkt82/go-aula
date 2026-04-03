# go-aula

Go client for Denmark's Aula school communication platform. Ported from the Rust [aulalibre](https://github.com/eisbaw/aulalibre) project.

## Disclaimer

This is an **unofficial**, community-driven project. It is **not affiliated with, endorsed by, or associated with** Aula, [KOMBIT](https://kombit.dk/), [Netcompany](https://netcompany.com/), or any Danish municipality.

- This software accesses Aula's API using **your own credentials** to retrieve **your own data**.
- It does not scrape, store, or redistribute data belonging to other users.
- Use at your own risk. The authors are not responsible for any consequences arising from the use of this software, including but not limited to account suspension or violation of Aula's terms of service.
- This project is provided "as is" under the [MIT License](LICENSE), with no warranty of any kind.

## Components

- **`aula-cli`** -- Command-line interface for reading messages, calendar, posts, gallery, documents, presence, and more
- **`aula-mcp`** -- MCP server exposing Aula data as tools for Claude and other LLM clients

## CLI commands

All commands support `--json` for machine-readable output.

| Command | Description |
|---------|-------------|
| `aula-cli auth login` | Browser-based OIDC login (UniLogin or MitID) |
| `aula-cli auth status` | Show current login state and token expiry |
| `aula-cli auth refresh` | Refresh an expired access token |
| `aula-cli auth logout` | Clear session and tokens |
| `aula-cli messages list` | List message threads |
| `aula-cli messages read <id>` | Read a specific thread |
| `aula-cli messages folders` | List message folders |
| `aula-cli calendar today` | Today's calendar events |
| `aula-cli calendar week` | This week's events |
| `aula-cli calendar show <id>` | Show a specific event |
| `aula-cli presence status` | Children's current presence |
| `aula-cli presence schedule` | This week's presence schedule |
| `aula-cli posts list` | List institution feed posts |
| `aula-cli posts show <id>` | Show a specific post |
| `aula-cli gallery list` | List photo albums |
| `aula-cli gallery show <id>` | Show media in an album |
| `aula-cli documents list` | List secure documents |
| `aula-cli documents show <id>` | Show a specific document |
| `aula-cli notifications list` | List recent notifications |
| `aula-cli search <query>` | Search profiles and groups |
| `aula-cli groups list` | List groups for an institution profile |
| `aula-cli groups show <id>` | Show group details |
| `aula-cli groups members <id>` | Show group members |
| `aula-cli profile me` | Show your profile |
| `aula-cli profile master-data` | Show name, email, phone |
| `aula-cli config policy` | Show data policy documents per institution |
| `aula-cli config privacy` | Show the platform privacy policy |
| `aula-cli config path` | Show the configuration file path |

Global flags: `--json`, `--env <environment>`, `--verbose`, `--profile <name>`.

## Getting started

```bash
# Build
make build

# Log in (opens browser for UniLogin/MitID)
./aula-cli auth login

# Check your session
./aula-cli auth status

# Try some commands
./aula-cli messages list
./aula-cli calendar today
./aula-cli presence status
```

## Configuration

Configuration file location:
- **Linux**: `~/.config/aula/config.toml`
- **macOS**: `~/Library/Application Support/aula/config.toml`
- **Windows**: `%APPDATA%\aula\config.toml`

```toml
default_environment = "production"
default_format = "text"
default_profile = "MyChild"
verbose = false
```

All settings can be overridden with CLI flags.

## Authentication

OIDC Authorization Code + PKCE flow through `login.aula.dk`.

1. `aula-cli auth login` opens your browser for UniLogin/MitID authentication.
2. Complete the login. The browser will end up on Aula's website — that's expected.
3. Press Enter in the terminal. The browser opens again and lands on a page that won't load.
4. Copy the URL from the browser's address bar and paste it into the terminal.

Tokens are stored in the platform data directory (`%APPDATA%\aula` on Windows, `~/.local/share/aula` on Linux, `~/Library/Application Support/aula` on macOS) and refreshed automatically.

## MCP Server

The `aula-mcp` binary is an MCP (Model Context Protocol) server that exposes Aula data as tools for Claude Desktop, Claude Code, and other MCP-compatible clients.

### Available tools

| Tool | Description |
|------|-------------|
| `list_messages` | List message threads |
| `read_message` | Read messages in a thread |
| `list_events` | List calendar events |
| `show_event` | Show event details |
| `list_posts` | List institution posts |
| `show_post` | Show a specific post |
| `presence_status` | Children's presence status |
| `daily_overview` | Today's presence overview |
| `list_notifications` | List notifications |
| `search` | Search across all content |
| `list_albums` | List photo albums |
| `list_documents` | List secure documents |
| `profile` | Show your profile |

### Setup

1. Build: `make build`
2. Log in first: `./aula-cli auth login`
3. The repo includes a `.mcp.json` that configures the MCP server for Claude Code automatically using `go run`

### Claude Desktop

Add to your Claude Desktop config:

```json
{
  "mcpServers": {
    "aula": {
      "command": "/path/to/aula-mcp"
    }
  }
}
```

Replace `/path/to/aula-mcp` with the actual path to the built binary.

## Project structure

```
cmd/
  aula-cli/           CLI binary entry point
  aula-mcp/           MCP server entry point
internal/
  aulaapi/            API client library
    enums/            Enum types from Aula .NET assemblies
    models/           Data model structs
    services/         API service functions (one per domain)
  cli/                CLI support (config, output)
    commands/         Cobra command implementations
  mcp/                MCP server (tool definitions + handlers)
```

## License

MIT License.
