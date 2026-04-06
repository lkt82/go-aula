package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lkt82/go-aula/internal/aulaapi/models"
	"github.com/lkt82/go-aula/internal/aulaapi/services"
	"github.com/lkt82/go-aula/internal/cli"
)

// NewMessagesCmd creates the "messages" command group.
func NewMessagesCmd(jsonFlag *bool, envFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "messages",
		Short: "Read and send messages (threads)",
	}

	// list
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List message threads (inbox view)",
		Run: func(c *cobra.Command, args []string) {
			limit, _ := c.Flags().GetInt32("limit")
			page, _ := c.Flags().GetInt32("page")
			all, _ := c.Flags().GetBool("all")
			unread, _ := c.Flags().GetBool("unread")
			marked, _ := c.Flags().GetBool("marked")
			folder, _ := c.Flags().GetInt64("folder")
			var pagePtr *int
			if c.Flags().Changed("page") {
				p := int(page)
				pagePtr = &p
			}
			var folderPtr *int64
			if c.Flags().Changed("folder") {
				folderPtr = &folder
			}
			msgHandleList(int(limit), pagePtr, all, unread, marked, folderPtr, *jsonFlag, *envFlag)
		},
	}
	listCmd.Flags().Int32P("limit", "n", 20, "Maximum number of threads to show")
	listCmd.Flags().Int32("page", 0, "Page number for pagination")
	listCmd.Flags().Bool("all", false, "Fetch and display all pages")
	listCmd.Flags().Bool("unread", false, "Show only unread threads")
	listCmd.Flags().Bool("marked", false, "Show only marked (starred) threads")
	listCmd.Flags().Int64("folder", 0, "Filter by folder ID")
	cmd.AddCommand(listCmd)

	// read
	readCmd := &cobra.Command{
		Use:     "read <thread-id>",
		Aliases: []string{"show"},
		Short:   "Show messages in a thread",
		Args:    cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			threadID := mustParseInt64(args[0])
			page, _ := c.Flags().GetInt("page")
			var pagePtr *int
			if c.Flags().Changed("page") {
				pagePtr = &page
			}
			msgHandleRead(threadID, pagePtr, *jsonFlag, *envFlag)
		},
	}
	readCmd.Flags().Int("page", 0, "Page number for pagination")
	cmd.AddCommand(readCmd)

	// send
	sendCmd := &cobra.Command{
		Use:   "send",
		Short: "Send a new message (start a new thread)",
		Run: func(c *cobra.Command, args []string) {
			to, _ := c.Flags().GetInt64Slice("to")
			subject, _ := c.Flags().GetString("subject")
			body, _ := c.Flags().GetString("body")
			var subjectPtr *string
			if c.Flags().Changed("subject") {
				subjectPtr = &subject
			}
			msgHandleSend(to, subjectPtr, body, *jsonFlag, *envFlag)
		},
	}
	sendCmd.Flags().Int64SliceP("to", "t", nil, "Recipient profile IDs (comma-separated)")
	sendCmd.Flags().StringP("subject", "s", "", "Message subject")
	sendCmd.Flags().StringP("body", "b", "", "Message body text")
	_ = sendCmd.MarkFlagRequired("to")
	_ = sendCmd.MarkFlagRequired("body")
	cmd.AddCommand(sendCmd)

	// reply
	replyCmd := &cobra.Command{
		Use:   "reply <thread-id>",
		Short: "Reply to an existing thread",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			threadID := mustParseInt64(args[0])
			body, _ := c.Flags().GetString("body")
			msgHandleReply(threadID, body, *jsonFlag, *envFlag)
		},
	}
	replyCmd.Flags().StringP("body", "b", "", "Reply body text")
	_ = replyCmd.MarkFlagRequired("body")
	cmd.AddCommand(replyCmd)

	// mark-read
	cmd.AddCommand(&cobra.Command{
		Use:   "mark-read <thread-id>",
		Short: "Mark a thread as read",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			threadID := mustParseInt64(args[0])
			msgHandleMarkRead(threadID, *jsonFlag, *envFlag)
		},
	})

	// delete
	cmd.AddCommand(&cobra.Command{
		Use:   "delete <thread-id>",
		Short: "Delete a thread",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			threadID := mustParseInt64(args[0])
			msgHandleDelete(threadID, *jsonFlag, *envFlag)
		},
	})

	// folders
	cmd.AddCommand(&cobra.Command{
		Use:   "folders",
		Short: "List message folders",
		Run: func(c *cobra.Command, args []string) {
			msgHandleFolders(*jsonFlag, *envFlag)
		},
	})

	// move
	moveCmd := &cobra.Command{
		Use:   "move <thread-id>",
		Short: "Move a thread to a folder",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			threadID := mustParseInt64(args[0])
			folder, _ := c.Flags().GetInt64("folder")
			msgHandleMove(threadID, folder, *jsonFlag, *envFlag)
		},
	}
	moveCmd.Flags().Int64("folder", 0, "Target folder ID")
	_ = moveCmd.MarkFlagRequired("folder")
	cmd.AddCommand(moveCmd)

	return cmd
}

// ---------------------------------------------------------------------------
// Handlers
// ---------------------------------------------------------------------------

func msgHandleList(limit int, page *int, all, unread, marked bool, folder *int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	var filterType *string
	if unread {
		ft := "FilterUnread"
		filterType = &ft
	} else if marked {
		ft := "FilterMarked"
		filterType = &ft
	}
	sortOn := "date"
	orderDir := "desc"

	if all {
		var allThreads []models.MessageThreadSubscription
		currentPage := 0
		for {
			p := currentPage
			threadArgs := &models.GetThreadListArguments{
				FolderID:       folder,
				FilterType:     filterType,
				SortOn:         &sortOn,
				OrderDirection: &orderDir,
				Page:           &p,
			}
			result, err := services.GetThreadList(ctx, session, threadArgs)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			allThreads = append(allThreads, result.Threads...)
			if !result.MoreMessagesExist {
				break
			}
			currentPage++
		}
		merged := models.MessageThreadSubscriptionList{
			Threads:           allThreads,
			MoreMessagesExist: false,
		}
		if jsonOut {
			cli.PrintJSON(merged)
		} else {
			printThreadList(&merged, limit)
		}
	} else {
		threadArgs := &models.GetThreadListArguments{
			FolderID:       folder,
			FilterType:     filterType,
			SortOn:         &sortOn,
			OrderDirection: &orderDir,
			Page:           page,
		}
		result, err := services.GetThreadList(ctx, session, threadArgs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if jsonOut {
			cli.PrintJSON(result)
		} else {
			printThreadList(&result, limit)
		}
	}
}

func printThreadList(list *models.MessageThreadSubscriptionList, limit int) {
	if len(list.Threads) == 0 {
		fmt.Println("No threads found.")
		return
	}
	table := cli.NewTable([]cli.Column{
		{Header: "ID", Width: 6},
		{Header: "", Width: 1},
		{Header: "FROM", Width: 20},
		{Header: "SUBJECT", Width: 40},
		{Header: "DATE", Width: 20},
	})
	table.PrintHeader()

	for i, thread := range list.Threads {
		if i >= limit {
			break
		}
		id := ""
		if thread.ID != nil {
			id = fmt.Sprintf("%d", *thread.ID)
		}
		marker := cli.UnreadMarker(thread.Read)

		from := "(unknown)"
		if thread.LatestMessage != nil && thread.LatestMessage.Sender != nil && thread.LatestMessage.Sender.DisplayName != nil {
			from = *thread.LatestMessage.Sender.DisplayName
		} else if thread.Creator != nil && thread.Creator.FullName != nil {
			from = *thread.Creator.FullName
		}

		subject := "(no subject)"
		if thread.Subject != nil {
			subject = *thread.Subject
		}

		date := ""
		if thread.LatestMessage != nil && thread.LatestMessage.SendDateTime != nil {
			date = *thread.LatestMessage.SendDateTime
		}

		fmt.Printf("%-6s %s %-20s %-40s %-20s\n",
			id, marker, cli.Truncate(from, 20), cli.Truncate(subject, 40), cli.FormatDatetime(date))
	}

	if list.Page != nil {
		cli.PrintPaginationHint(list.Page, list.MoreMessagesExist, "--page")
	}
}

func msgHandleRead(threadID int64, page *int, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	p := 0
	if page != nil {
		p = *page
	}
	args := &models.GetMessagesForThreadArguments{
		ThreadID: &threadID,
		Page:     &p,
	}

	result, err := services.GetThreadByID(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if jsonOut {
		cli.PrintJSON(result)
	} else {
		printThreadMessages(&result)
	}
}

func printThreadMessages(thread *models.MessagesInThreadDto) {
	subject := "(no subject)"
	if thread.Subject != nil {
		subject = *thread.Subject
	}
	fmt.Println(cli.Bold(fmt.Sprintf("Thread: %s", subject)))
	if thread.FolderName != nil {
		fmt.Printf("  Folder: %s", *thread.FolderName)
	}
	if thread.Muted {
		fmt.Printf("  %s", cli.Dim("[MUTED]"))
	}
	if thread.Marked {
		fmt.Printf("  %s", cli.Yellow("[MARKED]"))
	}
	if thread.Sensitive {
		fmt.Printf("  %s", cli.Red("[SENSITIVE]"))
	}
	fmt.Println()

	if thread.TotalMessageCount != nil {
		fmt.Printf("  Messages: %d\n", *thread.TotalMessageCount)
	}
	fmt.Println(cli.Dim("=" + string(make([]byte, 71))))

	if thread.FirstMessage != nil {
		printMessage(thread.FirstMessage)
	}
	for _, msg := range thread.Messages {
		printMessage(&msg)
	}

	if thread.MoreMessagesExist {
		nextPage := 0
		if thread.Page != nil {
			nextPage = *thread.Page
		}
		cli.PrintPaginationHint(&nextPage, true, "--page")
	}
}

func printMessage(msg *models.MessageDto) {
	sender := "(unknown)"
	if msg.Sender != nil && msg.Sender.FullName != nil {
		sender = *msg.Sender.FullName
	}
	date := ""
	if msg.SendDateTime != nil {
		date = *msg.SendDateTime
	}
	msgType := ""
	if msg.MessageType != nil {
		msgType = *msg.MessageType
	}

	fmt.Println()
	fmt.Printf("--- %s  (%s)  [%s]\n", cli.Bold(sender), cli.Dim(date), cli.Dim(msgType))

	if msg.Text != nil && msg.Text.HTML != nil {
		plain := cli.StripHTMLTags(*msg.Text.HTML)
		fmt.Println(plain)
	}

	if len(msg.Attachments) > 0 {
		fmt.Printf("  %s\n", cli.Dim(fmt.Sprintf("[%d attachment(s)]", len(msg.Attachments))))
	}
}

func msgHandleSend(to []int64, subject *string, body string, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	if len(to) == 0 {
		fmt.Fprintln(os.Stderr, "error: at least one recipient is required (--to)")
		os.Exit(1)
	}

	recipients := make([]models.RecipientApiModel, len(to))
	for i, id := range to {
		ownerType := "InstitutionProfile"
		recipients[i] = models.RecipientApiModel{
			ID:               &id,
			MailBoxOwnerType: &ownerType,
			ProfileID:        &id,
		}
	}

	text := body
	emptyAttachments := []int64{}
	emptyBcc := []models.RecipientApiModel{}
	args := &models.StartNewThreadRequestArguments{
		Message: &models.MessageContentRequest{
			AttachmentIDs: emptyAttachments,
			Text:          &text,
		},
		Subject:       subject,
		Recipients:    recipients,
		BccRecipients: emptyBcc,
	}

	result, err := services.StartNewThread(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		fmt.Println("Message sent.")
		// Try to extract thread ID from raw JSON
		var m map[string]json.RawMessage
		if json.Unmarshal(result, &m) == nil {
			if v, ok := m["id"]; ok {
				fmt.Printf("  Thread: %s\n", string(v))
			} else if v, ok := m["threadId"]; ok {
				fmt.Printf("  Thread: %s\n", string(v))
			}
		}
	}
}

func msgHandleReply(threadID int64, body string, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	text := body
	emptyAttachments := []int64{}
	args := &models.ReplyMessageArgument{
		ThreadID: &threadID,
		Message: &models.MessageContentRequest{
			AttachmentIDs: emptyAttachments,
			Text:          &text,
		},
	}

	result, err := services.ReplyToThread(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		fmt.Printf("Reply sent to thread %d.\n", threadID)
	}
}

func msgHandleMarkRead(threadID int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	p := 0
	fetchArgs := &models.GetMessagesForThreadArguments{
		ThreadID: &threadID,
		Page:     &p,
	}

	thread, err := services.GetThreadByID(ctx, session, fetchArgs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to fetch thread: %v\n", err)
		os.Exit(1)
	}

	var lastMsgID *string
	if len(thread.Messages) > 0 {
		lastMsgID = thread.Messages[len(thread.Messages)-1].ID
	} else if thread.FirstMessage != nil {
		lastMsgID = thread.FirstMessage.ID
	}

	if lastMsgID == nil {
		fmt.Fprintln(os.Stderr, "error: thread has no messages to mark as read")
		os.Exit(1)
	}

	msgID := *lastMsgID
	args := &models.SetLastMessageRequestArguments{
		MessageID: &msgID,
		ThreadID:  &threadID,
	}

	result, err := services.SetLastReadMessage(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		fmt.Printf("Thread %d marked as read.\n", threadID)
	}
}

func msgHandleDelete(threadID int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	args := &models.DeleteThreadArguments{
		ThreadIDs: []int64{threadID},
	}

	result, err := services.DeleteThreads(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		fmt.Printf("Thread %d deleted.\n", threadID)
	}
}

func msgHandleFolders(jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	args := &models.GetFoldersArguments{}

	folders, err := services.GetFolders(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(folders)
	} else if len(folders) == 0 {
		fmt.Println("No folders found.")
	} else {
		table := cli.NewTable([]cli.Column{
			{Header: "ID", Width: 8},
			{Header: "NAME", Width: 20},
			{Header: "TYPE", Width: 10},
		})
		table.PrintHeader()
		for _, f := range folders {
			id := ""
			if f.ID != nil {
				id = fmt.Sprint(*f.ID)
			}
			name := "(unnamed)"
			if f.Name != nil {
				name = *f.Name
			}
			ftype := ""
			if f.FolderType != nil {
				ftype = *f.FolderType
			}
			table.PrintRow([]string{id, name, ftype})
		}
	}
}

func msgHandleMove(threadID, folderID int64, jsonOut bool, envOverride string) {
	session := cli.BuildSession(envOverride)
	ctx := context.Background()

	args := &models.MoveThreadsToFolderRequestArguments{
		ThreadIDs: []int64{threadID},
		FolderID:  &folderID,
	}

	result, err := services.MoveThreadsToFolder(ctx, session, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if jsonOut {
		cli.PrintJSON(result)
	} else {
		fmt.Printf("Thread %d moved to folder %d.\n", threadID, folderID)
	}
}
