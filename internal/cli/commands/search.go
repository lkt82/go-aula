package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/lkt82/go-aula/internal/cli"
)

// NewSearchCmd creates the "search" command.
func NewSearchCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search <query>",
		Short: "Search across Aula content",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			query := args[0]
			limit, _ := c.Flags().GetInt("limit")
			page, _ := c.Flags().GetInt("page")
			counts, _ := c.Flags().GetBool("counts")
			probe, _ := c.Flags().GetBool("probe")
			docType, _ := c.Flags().GetString("type")

			var pagePtr *int
			if c.Flags().Changed("page") {
				pagePtr = &page
			}
			searchHandle(query, limit, pagePtr, counts, probe, docType, *jsonFlag, *envFlag)
		},
	}
	cmd.Flags().IntP("limit", "n", 20, "Maximum number of results")
	cmd.Flags().Int("page", 0, "Page number")
	cmd.Flags().Bool("counts", false, "Include document type counts in output")
	cmd.Flags().Bool("probe", false, "Probe all search endpoints and report which ones work")
	cmd.Flags().String("type", "ThreadMessage", "Content type: ThreadMessage, Post, Profile, Event")
	return cmd
}

func searchHandle(query string, limit int, page *int, counts, probe bool, docType string, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	if probe {
		searchProbeEndpoints(ctx, session, query, limit)
		return
	}

	if err := session.EnsureContextInitialized(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to initialize session: %v\n", err)
		os.Exit(1)
	}

	offset := 0
	params := &models.GlobalSearchParameters{
		Text:                                &query,
		Limit:                               &limit,
		Offset:                              &offset,
		PageLimit:                           &limit,
		PageNumber:                          page,
		DocTypeCount:                        counts,
		DocType:                             &docType,
		InstitutionProfileIDs:               session.InstitutionProfileIDs(),
		ActiveChildrenInstitutionProfileIDs: session.ChildrenInstProfileIDs(),
	}
	result, err := services.GlobalSearch(ctx, session, params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: search failed: %v\n", err)
		os.Exit(1)
	}

	if jsonOut {
		cli.PrintJSON(result)
	} else {
		searchDisplayResults(query, counts, &result)
	}
}

func searchTryGlobal(ctx context.Context, session *aulaapi.Session, query string, limit int, page *int, counts bool) (*models.SearchResponse, error) {
	offset := 0
	params := &models.GlobalSearchParameters{
		Text:                                &query,
		Limit:                               &limit,
		Offset:                              &offset,
		PageLimit:                           &limit,
		PageNumber:                          page,
		DocTypeCount:                        counts,
		InstitutionProfileIDs:               session.InstitutionProfileIDs(),
		ActiveChildrenInstitutionProfileIDs: session.ChildrenInstProfileIDs(),
	}
	result, err := services.GlobalSearch(ctx, session, params)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func searchAcrossDocTypes(ctx context.Context, session *aulaapi.Session, query string, limit int, page *int, counts bool) (*models.SearchResponse, error) {
	docTypes := []string{"ThreadMessage", "Post", "Profile", "Event"}
	var allItems []models.SearchResultItem
	total := 0
	succeeded := 0

	offset := 0
	for _, dt := range docTypes {
		docType := dt
		params := &models.GlobalSearchParameters{
			Text:                                &query,
			Limit:                               &limit,
			Offset:                              &offset,
			PageLimit:                           &limit,
			PageNumber:                          page,
			DocTypeCount:                        counts,
			DocType:                             &docType,
			InstitutionProfileIDs:               session.InstitutionProfileIDs(),
			ActiveChildrenInstitutionProfileIDs: session.ChildrenInstProfileIDs(),
		}
		result, err := services.GlobalSearch(ctx, session, params)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[search] findGeneric(%s) failed: %v\n", dt, err)
			continue
		}
		succeeded++
		if result.TotalSize != nil {
			total += *result.TotalSize
		}
		allItems = append(allItems, result.Results...)
	}

	if succeeded == 0 {
		return nil, fmt.Errorf("all search endpoints failed")
	}

	return &models.SearchResponse{
		TotalSize: &total,
		Results:   allItems,
	}, nil
}

func searchTryProfileSearch(ctx context.Context, session *aulaapi.Session, query string, limit int) (*models.SearchResponse, error) {
	params := &models.SearchForProfilesAndGroupsParameters{
		Text:  &query,
		Limit: &limit,
	}
	result, err := services.SearchForProfiles(ctx, session, params)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func searchTryGroupSearch(ctx context.Context, session *aulaapi.Session, query string, limit int) ([]models.SearchResultItem, error) {
	offset := 0
	params := &models.SearchGroupRequestModel{
		Text:   &query,
		Limit:  &limit,
		Offset: &offset,
	}
	result, err := services.SearchGroups(ctx, session, params)
	if err != nil {
		return nil, err
	}
	// Convert group results to SearchResultItem
	var items []models.SearchResultItem
	for _, g := range result.Results {
		docID := ""
		if g.ID != nil {
			docID = fmt.Sprintf("g-%d", *g.ID)
		}
		docType := "Group"
		items = append(items, models.SearchResultItem{
			DocID:           &docID,
			DocType:         &docType,
			InstitutionCode: g.InstitutionCode,
			InstitutionName: g.InstitutionName,
			Name:            g.Name,
		})
	}
	return items, nil
}

func searchTryCombined(ctx context.Context, session *aulaapi.Session, query string, limit int) (*models.SearchResponse, error) {
	var allItems []models.SearchResultItem
	total := 0
	sourcesOK := 0

	// Profiles
	profileResult, err := searchTryProfileSearch(ctx, session, query, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[search] findProfiles failed: %v\n", err)
	} else {
		if profileResult.TotalSize != nil {
			total += *profileResult.TotalSize
		}
		allItems = append(allItems, profileResult.Results...)
		sourcesOK++
	}

	// Groups
	groupItems, err := searchTryGroupSearch(ctx, session, query, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[search] findGroups failed: %v\n", err)
	} else {
		total += len(groupItems)
		allItems = append(allItems, groupItems...)
		sourcesOK++
	}

	if sourcesOK == 0 {
		return nil, fmt.Errorf("all combined search endpoints failed")
	}

	return &models.SearchResponse{
		TotalSize: &total,
		Results:   allItems,
	}, nil
}

func searchProbeEndpoints(ctx context.Context, session *aulaapi.Session, query string, limit int) {
	fmt.Println(cli.Bold(fmt.Sprintf("Probing search endpoints with query '%s':", query)))
	fmt.Println()

	// findGeneric
	fmt.Print("  search.findGeneric ... ")
	if r, err := searchTryGlobal(ctx, session, query, limit, nil, false); err != nil {
		fmt.Printf("FAIL: %v\n", err)
	} else {
		n := len(r.Results)
		total := 0
		if r.TotalSize != nil {
			total = *r.TotalSize
		}
		fmt.Printf("OK (%d results, %d total)\n", n, total)
	}

	// findProfiles
	fmt.Print("  search.findProfiles ... ")
	if r, err := searchTryProfileSearch(ctx, session, query, limit); err != nil {
		fmt.Printf("FAIL: %v\n", err)
	} else {
		n := len(r.Results)
		total := 0
		if r.TotalSize != nil {
			total = *r.TotalSize
		}
		fmt.Printf("OK (%d results, %d total)\n", n, total)
	}

	// findGroups
	fmt.Print("  search.findGroups ... ")
	if items, err := searchTryGroupSearch(ctx, session, query, limit); err != nil {
		fmt.Printf("FAIL: %v\n", err)
	} else {
		fmt.Printf("OK (%d results)\n", len(items))
	}

	fmt.Println()
	fmt.Println("Probe complete.")
}

func searchDisplayResults(query string, showCounts bool, result *models.SearchResponse) {
	total := 0
	if result.TotalSize != nil {
		total = *result.TotalSize
	}
	fmt.Println(cli.Bold(fmt.Sprintf("Search results for '%s' (%d total):", query, total)))
	fmt.Println()

	if showCounts && len(result.DocTypeCount) > 0 {
		fmt.Println("Content types:")
		for _, c := range result.DocTypeCount {
			name := "(unknown)"
			if c.Name != nil {
				name = *c.Name
			}
			count := 0
			if c.Count != nil {
				count = *c.Count
			}
			fmt.Printf("  %s: %d\n", name, count)
		}
		fmt.Println()
	}

	if len(result.Results) > 0 {
		table := cli.NewTable([]cli.Column{
			{Header: "TYPE", Width: 10},
			{Header: "NAME", Width: 30},
			{Header: "DESCRIPTION", Width: 30},
		})
		table.PrintHeader()
		for _, item := range result.Results {
			docType := ""
			if item.DocType != nil {
				docType = *item.DocType
			}
			name := "(unnamed)"
			if item.Name != nil {
				name = *item.Name
			}
			desc := ""
			if item.Description != nil {
				desc = *item.Description
			}
			table.PrintRow([]string{
				cli.Truncate(docType, 10),
				cli.Truncate(name, 30),
				cli.Truncate(desc, 30),
			})
		}
	} else {
		fmt.Println("No results found.")
	}
}
