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
| `aula auth login` | Browser-based OIDC login (UniLogin or MitID) |
| `aula auth status` | Show current login state and token expiry |
| `aula auth refresh` | Refresh an expired access token |
| `aula auth logout` | Clear session and tokens |
| `aula messages list` | List message threads |
| `aula messages read <id>` | Read a specific thread |
| `aula messages folders` | List message folders |
| `aula calendar today` | Today's calendar events |
| `aula calendar week` | This week's events |
| `aula calendar show <id>` | Show a specific event |
| `aula presence status` | Children's current presence |
| `aula presence schedule` | This week's presence schedule |
| `aula posts list` | List institution feed posts |
| `aula posts show <id>` | Show a specific post |
| `aula gallery list` | List photo albums |
| `aula gallery show <id>` | Show media in an album |
| `aula documents list` | List secure documents |
| `aula documents show <id>` | Show a specific document |
| `aula notifications list` | List recent notifications |
| `aula search <query>` | Search profiles and groups |
| `aula groups list` | List groups for an institution profile |
| `aula groups show <id>` | Show group details |
| `aula groups members <id>` | Show group members |
| `aula profile me` | Show your profile |
| `aula profile master-data` | Show name, email, phone |
| `aula config policy` | Show data policy documents per institution |
| `aula config privacy` | Show the platform privacy policy |
| `aula config path` | Show the configuration file path |

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

`~/.config/aula/config.toml`:

```toml
default_environment = "production"
default_format = "text"
default_profile = "MyChild"
verbose = false
```

All settings can be overridden with CLI flags.

## Authentication

OIDC Authorization Code + PKCE flow through `login.aula.dk`.

1. `aula auth login` opens your browser for UniLogin/MitID authentication.
2. Complete the login. The browser will end up on Aula's website — that's expected.
3. Press Enter in the terminal. The browser opens again and lands on a page that won't load.
4. Copy the URL from the browser's address bar and paste it into the terminal.

Tokens are stored at `~/.local/share/aula/tokens.json` and refreshed automatically.

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
3. Add to your MCP client config (see below)

### Claude Code

Add to `~/.claude/settings.json`:

```json
{
  "mcpServers": {
    "aula": {
      "command": "/path/to/aula-mcp"
    }
  }
}
```

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
