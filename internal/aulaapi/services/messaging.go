package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lkt82/go-aula/internal/aulaapi"
	"github.com/lkt82/go-aula/internal/aulaapi/models"
)

// SendEventReminderRequest is the request body for SendEventReminder.
type SendEventReminderRequest struct {
	ThreadID *int64  `json:"threadId,omitempty"`
	EntityID *string `json:"entityId,omitempty"`
}

// enumToQueryValue converts a string enum value for use in query parameters.
func enumToQueryValue(value string) string {
	return value
}

// GetThreadList lists message threads (inbox view) with filtering, sorting, and pagination.
func GetThreadList(ctx context.Context, s *aulaapi.Session, args *models.GetThreadListArguments) (models.MessageThreadSubscriptionList, error) {
	params := []string{"method=messaging.getThreads"}
	page := 0
	if args.Page != nil {
		page = *args.Page
	}
	params = append(params, fmt.Sprintf("page=%d", page))
	if args.FolderID != nil {
		params = append(params, fmt.Sprintf("folderId=%d", *args.FolderID))
	}
	if args.FilterType != nil {
		params = append(params, fmt.Sprintf("filterType=%s", enumToQueryValue(*args.FilterType)))
	}
	if args.SortType != nil {
		params = append(params, fmt.Sprintf("sortType=%s", enumToQueryValue(*args.SortType)))
	}
	if args.SortOrder != nil {
		params = append(params, fmt.Sprintf("sortOrder=%s", enumToQueryValue(*args.SortOrder)))
	}
	if args.MailBoxOwnerType != nil {
		params = append(params, fmt.Sprintf("mailBoxOwnerType=%s", enumToQueryValue(*args.MailBoxOwnerType)))
	}
	if args.MailBoxOwners != nil {
		for _, id := range args.MailBoxOwners {
			params = append(params, fmt.Sprintf("mailBoxOwners=%d", id))
		}
	}
	if args.ActiveChildren != nil {
		for _, id := range args.ActiveChildren {
			params = append(params, fmt.Sprintf("activeChildren=%d", id))
		}
	}
	if args.ThreadIDs != nil {
		for _, id := range args.ThreadIDs {
			params = append(params, fmt.Sprintf("threadIds=%d", id))
		}
	}
	path := "?" + strings.Join(params, "&")
	return aulaapi.SessionGet[models.MessageThreadSubscriptionList](ctx, s, path)
}

// GetThreadByID gets messages for a thread.
func GetThreadByID(ctx context.Context, s *aulaapi.Session, args *models.GetMessagesForThreadArguments) (models.MessagesInThreadDto, error) {
	params := []string{"method=messaging.getMessagesForThread"}
	if args.ThreadID != nil {
		params = append(params, fmt.Sprintf("threadId=%d", *args.ThreadID))
	}
	if args.Page != nil {
		params = append(params, fmt.Sprintf("page=%d", *args.Page))
	}
	if args.CommonInboxID != nil {
		params = append(params, fmt.Sprintf("commonInboxId=%d", *args.CommonInboxID))
	}
	path := "?" + strings.Join(params, "&")
	return aulaapi.SessionGet[models.MessagesInThreadDto](ctx, s, path)
}

// StartNewThread starts a new message thread.
func StartNewThread(ctx context.Context, s *aulaapi.Session, args *models.StartNewThreadRequestArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.startNewThread", args)
}

// ReplyToThread replies to an existing thread.
func ReplyToThread(ctx context.Context, s *aulaapi.Session, args *models.ReplyMessageArgument) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.reply", args)
}

// DeleteThreads deletes one or more threads.
func DeleteThreads(ctx context.Context, s *aulaapi.Session, args *models.DeleteThreadArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.deleteThreads", args)
}

// LeaveThread leaves a single thread.
func LeaveThread(ctx context.Context, s *aulaapi.Session, args *models.LeaveThreadArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.leaveThread", args)
}

// LeaveThreads leaves multiple threads at once.
func LeaveThreads(ctx context.Context, s *aulaapi.Session, args *models.LeaveThreadsRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.leaveThreads", args)
}

// ForwardThread forwards a thread to new recipients.
func ForwardThread(ctx context.Context, s *aulaapi.Session, args *models.ForwardThreadRequestArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.startNewThread", args)
}

// ReplyInNewThread replies to a thread by creating a new thread (quote-reply).
func ReplyInNewThread(ctx context.Context, s *aulaapi.Session, args *models.ForwardThreadRequestArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.startNewThread", args)
}

// GetMessageList gets messages in a thread (paginated).
func GetMessageList(ctx context.Context, s *aulaapi.Session, args *models.GetMessagesForThreadArguments) (models.MessagesInThreadDto, error) {
	return GetThreadByID(ctx, s, args)
}

// GetMessageInfoLight gets lightweight message info (for notifications/previews).
func GetMessageInfoLight(ctx context.Context, s *aulaapi.Session, messageID string, commonInboxID *int64, otpInboxID *int64) (models.GetMessageInfoLightDto, error) {
	params := []string{
		"method=messaging.getMessageInfoLight",
		fmt.Sprintf("messageId=%s", EncodeValue(messageID)),
	}
	if commonInboxID != nil {
		params = append(params, fmt.Sprintf("commonInboxId=%d", *commonInboxID))
	}
	if otpInboxID != nil {
		params = append(params, fmt.Sprintf("otpInboxId=%d", *otpInboxID))
	}
	path := "?" + strings.Join(params, "&")
	return aulaapi.SessionGet[models.GetMessageInfoLightDto](ctx, s, path)
}

// DeleteMessage deletes a message from a thread.
func DeleteMessage(ctx context.Context, s *aulaapi.Session, args *models.DeleteMessageRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.deleteMessage", args)
}

// EditMessage edits an existing message.
func EditMessage(ctx context.Context, s *aulaapi.Session, args *models.EditMessageRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.editMessage", args)
}

// SetLastReadMessage marks a message as the last read in a thread.
func SetLastReadMessage(ctx context.Context, s *aulaapi.Session, args *models.SetLastMessageRequestArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.setLastReadMessage", args)
}

// SetThreadMuted mutes or unmutes a thread.
func SetThreadMuted(ctx context.Context, s *aulaapi.Session, args *models.MuteThreadRequestArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.setThreadsMuted", args)
}

// SetThreadMarked marks or unmarks a thread (star/flag).
func SetThreadMarked(ctx context.Context, s *aulaapi.Session, args *models.MarkThreadsRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.setThreadsMarked", args)
}

// SetSensitiveLevel sets the sensitivity level on a thread.
func SetSensitiveLevel(ctx context.Context, s *aulaapi.Session, args *models.SetSensitivityLevelRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.setSensitivityLevel", args)
}

// AddRecipientsToThread adds recipients to an existing thread.
func AddRecipientsToThread(ctx context.Context, s *aulaapi.Session, args *models.AddRecipientArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.addRecipients", args)
}

// SetAutoReply sets an auto-reply message.
func SetAutoReply(ctx context.Context, s *aulaapi.Session, args *models.SetAutoReplyArguments) (models.MessageAutoReplyResult, error) {
	return aulaapi.SessionPost[models.MessageAutoReplyResult](ctx, s, "?method=messaging.setAutoReply", args)
}

// GetAutoReply gets the current auto-reply configuration.
func GetAutoReply(ctx context.Context, s *aulaapi.Session) (models.MessageAutoReplyResult, error) {
	return aulaapi.SessionGet[models.MessageAutoReplyResult](ctx, s, "?method=messaging.getAutoReply")
}

// DeleteAutoReply deletes the auto-reply configuration.
func DeleteAutoReply(ctx context.Context, s *aulaapi.Session) (json.RawMessage, error) {
	return aulaapi.SessionPostEmpty[json.RawMessage](ctx, s, "?method=messaging.deleteAutoReply")
}

// GetFolders gets folders for the current user's mailbox.
func GetFolders(ctx context.Context, s *aulaapi.Session, args *models.GetFoldersArguments) ([]models.Folder, error) {
	params := []string{
		"method=messaging.getFolders",
		fmt.Sprintf("includeDeletedFolders=%t", args.IncludeDeletedFolders),
	}
	if args.CommonInboxID != nil {
		params = append(params, fmt.Sprintf("commonInboxId=%d", *args.CommonInboxID))
	}
	path := "?" + strings.Join(params, "&")
	return aulaapi.SessionGet[[]models.Folder](ctx, s, path)
}

// CreateFolder creates a new message folder.
func CreateFolder(ctx context.Context, s *aulaapi.Session, args *models.CreateFolderArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.createFolder", args)
}

// UpdateFolder renames a folder.
func UpdateFolder(ctx context.Context, s *aulaapi.Session, args *models.UpdateFolderArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.updateFolder", args)
}

// DeleteFolder deletes a folder.
func DeleteFolder(ctx context.Context, s *aulaapi.Session, folderID int64, commonInboxID *int64) (json.RawMessage, error) {
	params := []string{
		"method=messaging.deletefolder",
		fmt.Sprintf("folderId=%d", folderID),
	}
	if commonInboxID != nil {
		params = append(params, fmt.Sprintf("commonInboxId=%d", *commonInboxID))
	}
	path := "?" + strings.Join(params, "&")
	return aulaapi.SessionPostEmpty[json.RawMessage](ctx, s, path)
}

// MoveThreadsToFolder moves threads to a folder.
func MoveThreadsToFolder(ctx context.Context, s *aulaapi.Session, args *models.MoveThreadsToFolderRequestArguments) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.moveThreadsToFolder", args)
}

// GetCommonInboxes gets common (shared) inboxes for the user's institution profiles.
func GetCommonInboxes(ctx context.Context, s *aulaapi.Session, institutionProfileIDs []int64, includeProfilePictureURL bool) ([]models.CommonInboxesDto, error) {
	params := []string{"method=messaging.getCommonInboxes"}
	for _, id := range institutionProfileIDs {
		params = append(params, fmt.Sprintf("institutionProfileIds=%d", id))
	}
	if includeProfilePictureURL {
		params = append(params, "shouldIncludeProfilePictureUrl=true")
	}
	path := "?" + strings.Join(params, "&")
	return aulaapi.SessionGet[[]models.CommonInboxesDto](ctx, s, path)
}

// GetThreadsInBundleList gets threads within a bundle (grouped threads).
func GetThreadsInBundleList(ctx context.Context, s *aulaapi.Session, args *models.GetThreadsInBundleArguments) (models.MessageThreadSubscriptionList, error) {
	bundleID := int64(0)
	if args.BundleID != nil {
		bundleID = *args.BundleID
	}
	path := fmt.Sprintf("?method=messaging.getThreadsInBundle&bundleId=%d", bundleID)
	return aulaapi.SessionGet[models.MessageThreadSubscriptionList](ctx, s, path)
}

// SetSubscriptionStatus updates subscription status (read/unread) for thread subscriptions.
func SetSubscriptionStatus(ctx context.Context, s *aulaapi.Session, args *models.UpdateMessageThreadsSubscriptionStatusRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.updateSubscriptionStatus", args)
}

// CheckRecipientsForBlockedChannels checks whether recipients have blocked messaging channels.
func CheckRecipientsForBlockedChannels(ctx context.Context, s *aulaapi.Session, recipients []models.RecipientApiModel) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=municipalConfiguration.getBlockedCommunicationInstitutionProfilesAndGroups", recipients)
}

// AttachMessagesToSecureDocument attaches messages from a thread to a secure document.
func AttachMessagesToSecureDocument(ctx context.Context, s *aulaapi.Session, args *models.AttachMessagesToSecureDocumentRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.attachMessagesToSecureDocument", args)
}

// SendEventReminder sends an event reminder message.
func SendEventReminder(ctx context.Context, s *aulaapi.Session, args *SendEventReminderRequest) (json.RawMessage, error) {
	return aulaapi.SessionPost[json.RawMessage](ctx, s, "?method=messaging.sendEventReminder", args)
}
