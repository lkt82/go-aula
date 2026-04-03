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

// NewPostsCmd creates the "posts" command group.
func NewPostsCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "posts",
		Short: "View and manage posts in the institution feed",
	}

	// list
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List recent posts",
		Run: func(c *cobra.Command, args []string) {
			limit, _ := c.Flags().GetInt32("limit")
			group, _ := c.Flags().GetInt64("group")
			important, _ := c.Flags().GetBool("important")
			unread, _ := c.Flags().GetBool("unread")
			bookmarked, _ := c.Flags().GetBool("bookmarked")
			var groupPtr *int64
			if c.Flags().Changed("group") {
				groupPtr = &group
			}
			postsHandleList(int(limit), groupPtr, important, unread, bookmarked, *jsonFlag, *envFlag)
		},
	}
	listCmd.Flags().Int32P("limit", "n", 20, "Maximum number of posts to show")
	listCmd.Flags().Int64("group", 0, "Filter by group ID")
	listCmd.Flags().Bool("important", false, "Show only important posts")
	listCmd.Flags().Bool("unread", false, "Show only unread posts")
	listCmd.Flags().Bool("bookmarked", false, "Show only bookmarked posts")
	cmd.AddCommand(listCmd)

	// show
	cmd.AddCommand(&cobra.Command{
		Use:   "show <post-id>",
		Short: "Show a single post by ID",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			postID := mustParseInt64(args[0])
			postsHandleShow(postID, *jsonFlag, *envFlag)
		},
	})

	// create
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new post",
		Run: func(c *cobra.Command, args []string) {
			title, _ := c.Flags().GetString("title")
			body, _ := c.Flags().GetString("body")
			instCode, _ := c.Flags().GetString("institution-code")
			profile, _ := c.Flags().GetInt64("profile")
			allowComments, _ := c.Flags().GetBool("allow-comments")
			important, _ := c.Flags().GetBool("important")
			postsHandleCreate(title, body, instCode, profile, allowComments, important, *jsonFlag, *envFlag)
		},
	}
	createCmd.Flags().StringP("title", "t", "", "Post title")
	createCmd.Flags().StringP("body", "b", "", "Post body (HTML or plain text)")
	createCmd.Flags().String("institution-code", "", "Institution code")
	createCmd.Flags().Int64("profile", 0, "Creator institution profile ID")
	createCmd.Flags().Bool("allow-comments", true, "Allow comments on the post")
	createCmd.Flags().Bool("important", false, "Mark post as important")
	_ = createCmd.MarkFlagRequired("title")
	_ = createCmd.MarkFlagRequired("body")
	_ = createCmd.MarkFlagRequired("institution-code")
	_ = createCmd.MarkFlagRequired("profile")
	cmd.AddCommand(createCmd)

	return cmd
}

func postsHandleList(limit int, groupID *int64, important, unread, bookmarked bool, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	if err := session.EnsureContextInitialized(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to initialize session: %v\n", err)
		os.Exit(1)
	}
	instProfileIDs := session.AllInstitutionProfileIDs()

	parent := "profile"
	var isImportant *bool
	if important {
		t := true
		isImportant = &t
	}
	lim := limit
	idx := 0
	params := &models.GetPostApiParameters{
		Parent:                &parent,
		GroupID:               groupID,
		IsImportant:           isImportant,
		InstitutionProfileIDs: instProfileIDs,
		IsUnread:              unread,
		IsBookmarked:          bookmarked,
		Limit:                 &lim,
		Index:                 &idx,
	}

	result, err := services.GetPosts(ctx, session, params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		printPostList(&result, limit)
	}
}

func printPostList(result *models.GetPostApiResult, limit int) {
	if len(result.Posts) == 0 {
		fmt.Println("No posts found.")
		return
	}

	table := cli.NewTable([]cli.Column{
		{Header: "ID", Width: 8},
		{Header: "TITLE", Width: 30},
		{Header: "AUTHOR", Width: 20},
		{Header: "DATE", Width: 20},
	})
	table.PrintHeader()

	for i, post := range result.Posts {
		if i >= limit {
			break
		}
		id := ""
		if post.ID != nil {
			id = fmt.Sprintf("%d", *post.ID)
		}
		title := "(no title)"
		if post.Title != nil {
			title = *post.Title
		}
		author := "(unknown)"
		if post.OwnerProfile != nil && post.OwnerProfile.FullName != nil {
			author = *post.OwnerProfile.FullName
		}
		date := ""
		if post.TimeStamp != nil {
			date = *post.TimeStamp
		}
		flags := ""
		if post.IsImportant {
			flags += "!"
		}
		if post.IsBookmarked {
			flags += "*"
		}

		table.PrintRow([]string{
			id + flags,
			cli.Truncate(title, 30),
			cli.Truncate(author, 20),
			cli.FormatDatetime(date),
		})
	}

	if result.HasMorePosts {
		fmt.Fprintln(os.Stderr, "\n(more posts available)")
	}
}

func postsHandleShow(postID int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	post, err := services.GetPostByID(ctx, session, postID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(post)
	} else {
		printPostDetail(&post)
	}
}

func printPostDetail(post *models.PostApiDto) {
	title := "(no title)"
	if post.Title != nil {
		title = *post.Title
	}
	author := "(unknown)"
	if post.OwnerProfile != nil && post.OwnerProfile.FullName != nil {
		author = *post.OwnerProfile.FullName
	}
	date := ""
	if post.TimeStamp != nil {
		date = *post.TimeStamp
	}

	fmt.Println(cli.Bold(fmt.Sprintf("Post: %s", title)))
	fmt.Printf("  Author: %s\n", author)
	fmt.Printf("  Date: %s\n", cli.FormatDatetime(date))

	if post.IsImportant {
		fmt.Printf("  %s", cli.Yellow("[IMPORTANT]"))
	}
	if post.IsBookmarked {
		fmt.Print("  [BOOKMARKED]")
	}
	fmt.Println()

	if len(post.SharedWithGroups) > 0 {
		var names []string
		for _, g := range post.SharedWithGroups {
			if g.Name != nil {
				names = append(names, *g.Name)
			}
		}
		if len(names) > 0 {
			fmt.Printf("  Groups: %s\n", joinStrings(names, ", "))
		}
	}

	fmt.Println(cli.Dim("=" + repeatChar('=', 71)))

	if post.Content != nil && post.Content.HTML != nil {
		fmt.Println(cli.StripHTMLTags(*post.Content.HTML))
	}

	if len(post.Attachments) > 0 {
		fmt.Printf("\n  %s\n", cli.Dim(fmt.Sprintf("[%d attachment(s)]", len(post.Attachments))))
	}

	if post.CommentCount != nil && *post.CommentCount > 0 {
		fmt.Printf("\n  %s\n", cli.Dim(fmt.Sprintf("[%d comment(s)]", *post.CommentCount)))
	}
}

func postsHandleCreate(title, body, instCode string, profile int64, allowComments, important bool, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	params := &models.CreatePostApiParameter{
		Title:                          &title,
		Content:                        &body,
		InstitutionCode:                &instCode,
		CreatorInstitutionProfileID:    &profile,
		AllowComments:                  allowComments,
		IsImportant:                    important,
	}

	result, err := services.CreatePost(ctx, session, params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		fmt.Println("Post created.")
	}
}
