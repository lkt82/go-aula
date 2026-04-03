package commands

import "context"

// cmd_context returns a background context for CLI commands.
func cmd_context() context.Context {
	return context.Background()
}
