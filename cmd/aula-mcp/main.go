package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lkt82/go-aula/internal/aulaapi"
	aulamcp "github.com/lkt82/go-aula/internal/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	store, err := aulaapi.DefaultTokenStore()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: could not determine token store location")
		os.Exit(1)
	}

	if !store.Exists() {
		fmt.Fprintln(os.Stderr, "Not logged in. Run 'aula-cli auth login' first.")
		os.Exit(1)
	}

	client, err := aulaapi.NewAulaClientWithConfig(aulaapi.AulaClientConfig{
		Environment: aulaapi.EnvProduction,
		APIVersion:  23,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to create client: %v\n", err)
		os.Exit(1)
	}

	session, err := aulaapi.NewSession(client, store, aulaapi.DefaultSessionConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to create session: %v\n", err)
		os.Exit(1)
	}

	mcpServer := aulamcp.NewAulaServer(session)

	if err := server.ServeStdio(mcpServer); err != nil {
		log.Fatalf("MCP server error: %v", err)
	}
}
