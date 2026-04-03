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

// NewGalleryCmd creates the "gallery" command group.
func NewGalleryCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gallery",
		Short: "Browse and download gallery media (photos, videos)",
	}

	// list
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List gallery albums",
		Run: func(c *cobra.Command, args []string) {
			institution, _ := c.Flags().GetString("institution")
			limit, _ := c.Flags().GetInt32("limit")
			var instPtr *string
			if c.Flags().Changed("institution") {
				instPtr = &institution
			}
			galHandleList(instPtr, int(limit), *jsonFlag, *envFlag)
		},
	}
	listCmd.Flags().String("institution", "", "Filter by institution code")
	listCmd.Flags().Int32P("limit", "n", 20, "Maximum number of albums to show")
	cmd.AddCommand(listCmd)

	// show
	showCmd := &cobra.Command{
		Use:   "show <album-id>",
		Short: "Show album contents (media items)",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			albumID := mustParseInt64(args[0])
			limit, _ := c.Flags().GetInt32("limit")
			galHandleShow(albumID, int(limit), *jsonFlag, *envFlag)
		},
	}
	showCmd.Flags().Int32P("limit", "n", 20, "Maximum number of media items to show")
	cmd.AddCommand(showCmd)

	return cmd
}

func galHandleList(institution *string, limit int, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	lim := limit
	idx := 0
	sortOn := "createdAt"
	orderDir := "desc"
	filter := &models.GalleryViewFilter{
		SelectedInstitutionCodeForFilter: institution,
		Limit:          &lim,
		Index:          &idx,
		SortOn:         &sortOn,
		OrderDirection: &orderDir,
	}

	albums, err := services.GetAlbums(ctx, session, filter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(albums)
	} else if len(albums) == 0 {
		fmt.Println("No albums found.")
	} else {
		table := cli.NewTable([]cli.Column{
			{Header: "ID", Width: 8},
			{Header: "TITLE", Width: 30},
			{Header: "CREATOR", Width: 20},
			{Header: "ITEMS", Width: 8},
			{Header: "CREATED", Width: 16},
		})
		table.PrintHeader()
		for _, album := range albums {
			id := ""
			if album.ID != nil {
				id = fmt.Sprintf("%d", *album.ID)
			}
			title := "(untitled)"
			if album.Title != nil {
				title = *album.Title
			} else if album.Name != nil {
				title = *album.Name
			}
			creator := "(unknown)"
			if album.Creator != nil && album.Creator.Name != nil {
				creator = *album.Creator.Name
			}
			size := ""
			if album.TotalSize != nil {
				size = fmt.Sprintf("%d", *album.TotalSize)
			} else if album.Size != nil {
				size = fmt.Sprintf("%d", *album.Size)
			}
			date := ""
			if album.CreationDate != nil {
				date = *album.CreationDate
			}
			table.PrintRow([]string{id, title, creator, size, cli.FormatDatetime(date)})
		}
	}
}

func galHandleShow(albumID int64, limit int, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	lim := limit
	idx := 0
	filter := &models.GetMediaInAlbumFilter{
		AlbumID: &albumID,
		Limit:   &lim,
		Index:   &idx,
	}

	result, err := services.GetMediasInAlbum(ctx, session, filter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		if result.Album != nil {
			title := "(untitled)"
			if result.Album.Title != nil {
				title = *result.Album.Title
			} else if result.Album.Name != nil {
				title = *result.Album.Name
			}
			fmt.Println(cli.Bold(fmt.Sprintf("Album: %s", title)))
			if result.Album.Description != nil && *result.Album.Description != "" {
				fmt.Printf("  %s\n", *result.Album.Description)
			}
		}
		if result.MediaCount != nil {
			fmt.Printf("  Total media: %d\n", *result.MediaCount)
		}
		fmt.Println(cli.Dim("=" + repeatChar('=', 71)))

		if len(result.Results) > 0 {
			table := cli.NewTable([]cli.Column{
				{Header: "TITLE", Width: 30},
				{Header: "TYPE", Width: 10},
				{Header: "URL", Width: 40},
			})
			table.PrintHeader()
			for _, m := range result.Results {
				title := "(untitled)"
				if m.Title != nil {
					title = *m.Title
				}
				mediaType := ""
				if m.MediaType != nil {
					mediaType = *m.MediaType
				}
				urlStr := ""
				if m.ThumbnailURL != nil {
					urlStr = *m.ThumbnailURL
				} else if m.File != nil && m.File.URL != nil {
					urlStr = *m.File.URL
				}
				table.PrintRow([]string{title, mediaType, urlStr})
			}
		} else {
			fmt.Println("No media items found.")
		}
	}
}
