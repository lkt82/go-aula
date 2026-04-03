package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/lkt82/go-aula/internal/cli"
)

// NewDocumentsCmd creates the "documents" command group.
func NewDocumentsCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "documents",
		Short: "Browse and download shared documents",
	}

	// list
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List secure documents",
		Run: func(c *cobra.Command, args []string) {
			profiles, _ := c.Flags().GetInt64Slice("profiles")
			limit, _ := c.Flags().GetInt32("limit")
			unread, _ := c.Flags().GetBool("unread")
			docsHandleList(profiles, int(limit), unread, *jsonFlag, *envFlag)
		},
	}
	listCmd.Flags().Int64Slice("profiles", nil, "Filter by institution profile IDs (comma-separated)")
	listCmd.Flags().Int32P("limit", "n", 20, "Maximum number of documents to show")
	listCmd.Flags().Bool("unread", false, "Show only unread documents")
	cmd.AddCommand(listCmd)

	// show
	showCmd := &cobra.Command{
		Use:   "show <document-id>",
		Short: "Show document details by ID",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			documentID := mustParseInt64(args[0])
			docType, _ := c.Flags().GetString("doc-type")
			docsHandleShow(documentID, docType, *jsonFlag, *envFlag)
		},
	}
	showCmd.Flags().String("doc-type", "internal", "Document type: 'internal' (default) or 'external'")
	cmd.AddCommand(showCmd)

	return cmd
}

func docsHandleList(profiles []int64, limit int, unread bool, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	profileIDs := profiles
	if len(profileIDs) == 0 {
		if err := session.EnsureContextInitialized(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to initialize session: %v\n", err)
			os.Exit(1)
		}
		profileIDs = session.ChildrenInstProfileIDs()
	}

	var filterUnread *bool
	if unread {
		t := true
		filterUnread = &t
	}
	lim := limit
	idx := 0
	args := &models.GetSecureDocumentsArguments{
		FilterInstitutionProfileIDs: profileIDs,
		FilterUnread:                filterUnread,
		Index:                       &idx,
		Limit:                       &lim,
	}

	result, err := services.GetSecureDocuments(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		total := 0
		if result.TotalCount != nil {
			total = int(*result.TotalCount)
		}
		if len(result.Documents) > 0 {
			table := cli.NewTable([]cli.Column{
				{Header: "ID", Width: 8},
				{Header: "TITLE", Width: 30},
				{Header: "CATEGORY", Width: 15},
				{Header: "LOCKED", Width: 10},
				{Header: "UPDATED", Width: 16},
			})
			table.PrintHeader()
			for _, doc := range result.Documents {
				id := ""
				if doc.ID != nil {
					id = fmt.Sprintf("%d", *doc.ID)
				}
				title := "(untitled)"
				if doc.Title != nil {
					title = *doc.Title
				}
				category := ""
				if doc.Category != nil {
					category = *doc.Category
				}
				locked := "no"
				if doc.IsLocked {
					locked = "yes"
				}
				updated := ""
				if doc.UpdatedAt != nil {
					updated = *doc.UpdatedAt
				}
				table.PrintRow([]string{
					id,
					cli.Truncate(title, 30),
					cli.Truncate(category, 15),
					locked,
					cli.FormatDatetime(updated),
				})
			}
			fmt.Printf("\nTotal: %d document(s)\n", total)
		} else {
			fmt.Println("No documents found.")
		}
	}
}

func docsHandleShow(documentID int64, docType string, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	if docType == "external" {
		doc, err := services.GetExternalDocumentDetails(ctx, session, documentID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if jsonOut {
			cli.PrintJSON(doc)
		} else {
			fmt.Println(cli.Bold(fmt.Sprintf("External Document #%d", documentID)))
			cli.PrintJSON(doc)
		}
	} else {
		doc, err := services.GetInternalDocumentDetails(ctx, session, documentID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if jsonOut {
			cli.PrintJSON(doc)
		} else {
			title := "(untitled)"
			if doc.Title != nil {
				title = *doc.Title
			}
			fmt.Println(cli.Bold(fmt.Sprintf("Document: %s", title)))
			if doc.Category != nil {
				fmt.Printf("  Category: %s\n", *doc.Category)
			}
			if doc.Creator != nil && doc.Creator.Name != nil {
				fmt.Printf("  Creator: %s\n", *doc.Creator.Name)
			}
			if doc.CreatedAt != nil {
				fmt.Printf("  Created: %s\n", cli.FormatDatetime(*doc.CreatedAt))
			}
			if doc.UpdatedAt != nil {
				fmt.Printf("  Updated: %s\n", cli.FormatDatetime(*doc.UpdatedAt))
			}
			if doc.Version != nil {
				fmt.Printf("  Version: %d\n", *doc.Version)
			}
			fmt.Println(cli.Dim("=" + repeatChar('=', 71)))
			if doc.Content != nil && doc.Content.HTML != nil {
				fmt.Println(cli.StripHTMLTags(*doc.Content.HTML))
			}
		}
	}
}
