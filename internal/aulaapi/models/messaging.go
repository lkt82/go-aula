package models

import "encoding/json"

// RichTextWrapperDto represents rich text content wrapper (HTML body).
type RichTextWrapperDto struct {
	HTML *string `json:"html,omitempty"`
}

// DownloadFileFromAulaArguments represents a file download reference.
type DownloadFileFromAulaArguments struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// RecipientApiModel represents a recipient/mailbox-owner identity.
type RecipientApiModel struct {
	ID               *int64  `json:"id,omitempty"`
	OtpInboxID       *int64  `json:"otpInboxId,omitempty"`
	MailBoxOwnerType *string `json:"mailBoxOwnerType,omitempty"`
	ProfileID        *int64  `json:"profileId,omitempty"`
	IsDeactivated    bool    `json:"isDeactivated"`
	IsDeleted        bool    `json:"isDeleted"`
	PortalRole       *string `json:"portalRole,omitempty"`
}

// MailBox represents a simple mailbox identity.
type MailBox struct {
	ID          *int64          `json:"id,omitempty"`
	Email       *string         `json:"address,omitempty"`
	DisplayName *string         `json:"displayName,omitempty"`
	Relation    *string         `json:"relation,omitempty"`
	ShortName   json.RawMessage `json:"shortName,omitempty"`
}

// SimpleMessageThreadSubscription represents a simple subscription reference.
type SimpleMessageThreadSubscription struct {
	DisplayName *string `json:"displayName,omitempty"`
	Relation    *string `json:"relation,omitempty"`
	ShortName   *string `json:"shortName,omitempty"`
}

// MessageThread represents a core message thread entity.
type MessageThread struct {
	StartedDateTime  *string                            `json:"startedDateTime,omitempty"`
	Subject          *string                            `json:"subject,omitempty"`
	RequiredStepUp   bool                               `json:"requiredStepUp"`
	SensitivityLevel *string                            `json:"sensitivityLevel,omitempty"`
	Creator          json.RawMessage                    `json:"creator,omitempty"`
	OtherRecipients  []SimpleMessageThreadSubscription  `json:"otherRecipients,omitempty"`
	ThreadID         *string                            `json:"threadId,omitempty"`
	IsForwarded      bool                               `json:"isForwarded"`
}

// ThreadEntityLinkDto represents a link between a thread and an external entity.
type ThreadEntityLinkDto struct {
	EntityID   *string `json:"entityId,omitempty"`
	ThreadType *string `json:"threadType,omitempty"`
}

// MessageThreadLatestMessage represents a latest message summary within a thread subscription.
type MessageThreadLatestMessage struct {
	ID             *string             `json:"id,omitempty"`
	ThreadID       *int64              `json:"threadId,omitempty"`
	SendDateTime   *string             `json:"sendDateTime,omitempty"`
	Text           *RichTextWrapperDto `json:"text,omitempty"`
	Sender         *MailBox            `json:"sender,omitempty"`
	NewRecipient   *MailBox            `json:"newRecipient,omitempty"`
	HasAttachments bool                `json:"hasAttachments"`
	PendingMedia   bool                `json:"pendingMedia"`
}

// MessageThreadSubscriptionRelatedChild represents a child related to a thread subscription.
type MessageThreadSubscriptionRelatedChild struct {
	ID          *int64  `json:"id,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

// MessageThreadSubscriptionRelatedInstitution represents an institution related to a thread subscription.
type MessageThreadSubscriptionRelatedInstitution struct {
	InstitutionCode *string `json:"institutionCode,omitempty"`
	Name            *string `json:"name,omitempty"`
}

// MessageParticipantDto represents a participant in a message thread.
type MessageParticipantDto struct {
	MailBoxOwner             *RecipientApiModel `json:"mailBoxOwner,omitempty"`
	FullName                 *string            `json:"fullName,omitempty"`
	Metadata                 *string            `json:"metadata,omitempty"`
	LastReadMessageID        *string            `json:"lastReadMessageId,omitempty"`
	LastReadMessageTimestamp  *string            `json:"lastReadMessageTimestamp,omitempty"`
	ShortName                *string            `json:"shortName,omitempty"`
	ProfilePicture           *ProfilePictureDto `json:"profilePicture,omitempty"`
}

// MessageRegardingChildren represents children tagged on a thread.
type MessageRegardingChildren struct {
	ProfileID      *int64                         `json:"profileId,omitempty"`
	DisplayName    *string                        `json:"displayName,omitempty"`
	ProfilePicture *DownloadFileFromAulaArguments `json:"profilePicture,omitempty"`
}

// MessageDraft represents draft message state persisted on a subscription.
type MessageDraft struct {
	Text          *string `json:"text,omitempty"`
	AttachmentIDs []int64 `json:"attachmentIds,omitempty"`
}

// MessageThreadSubscription represents a user's subscription to a message thread.
type MessageThreadSubscription struct {
	ID                              *int64                    `json:"id,omitempty"`
	LeaveTime                       *string                   `json:"leaveTime,omitempty"`
	Muted                           bool                      `json:"muted"`
	Marked                          bool                      `json:"marked"`
	Read                            bool                      `json:"read"`
	Sensitive                       bool                      `json:"sensitive"`
	LastReadMessageID               *string                   `json:"lastReadMessageId,omitempty"`
	InstitutionCode                 *string                   `json:"institutionCode,omitempty"`
	Creator                         *MessageParticipantDto    `json:"creator,omitempty"`
	Recipients                      []MessageParticipantDto   `json:"recipients,omitempty"`
	RegardingChildren               []MessageRegardingChildren `json:"regardingChildren,omitempty"`
	LatestMessage                   *MessageThreadLatestMessage `json:"latestMessage,omitempty"`
	Subject                         *string                   `json:"subject,omitempty"`
	MessageDraft                    *MessageDraft             `json:"messageDraft,omitempty"`
	MailBoxOwner                    *RecipientApiModel        `json:"mailBoxOwner,omitempty"`
	CurrentFolder                   *Folder                   `json:"currentFolder,omitempty"`
	SubscriptionID                  *int64                    `json:"subscriptionId,omitempty"`
	IsThreadOrSubscriptionDeleted   bool                      `json:"isThreadOrSubscriptionDeleted"`
	SubscriptionType                *string                   `json:"subscriptionType,omitempty"`
	NumberOfBundleItems             *int64                    `json:"numberOfBundleItems,omitempty"`
	ExtraRecipientsCount            *int64                    `json:"extraRecipientsCount,omitempty"`
	BundleID                        *int64                    `json:"bundleId,omitempty"`
	ThreadEntityLinkDto             *ThreadEntityLinkDto      `json:"threadEntityLinkDto,omitempty"`
	PrimarySubscriptionID           *int64                    `json:"primarySubscriptionId,omitempty"`
}

// MessageThreadSubscriptionList represents a paginated list of thread subscriptions.
type MessageThreadSubscriptionList struct {
	Threads           []MessageThreadSubscription `json:"threads,omitempty"`
	Page              *int                        `json:"page,omitempty"`
	BundleID          *int64                      `json:"bundleId,omitempty"`
	MoreMessagesExist bool                        `json:"moreMessagesExist"`
}

// MessageDto represents a single message within a thread.
type MessageDto struct {
	ID                *string             `json:"id,omitempty"`
	MessageType       *string             `json:"messageType,omitempty"`
	SendDateTime      *string             `json:"sendDateTime,omitempty"`
	Text              *RichTextWrapperDto `json:"text,omitempty"`
	Sender            *MessageRecipient   `json:"sender,omitempty"`
	CanReplyToMessage bool                `json:"canReplyToMessage"`
	Attachments       []json.RawMessage   `json:"attachments,omitempty"`
	NewRecipient      *MessageRecipient   `json:"newRecipient,omitempty"`
	NewRecipients     []MessageRecipient  `json:"newRecipients,omitempty"`
	OriginalRecipients []MessageRecipient `json:"originalRecipients,omitempty"`
	LeaverName        *string             `json:"leaverName,omitempty"`
	InviterName       *string             `json:"inviterName,omitempty"`
	LeaverNames       []string            `json:"leaverNames,omitempty"`
}

// MessageRecipient represents a recipient on a message.
type MessageRecipient struct {
	ShortName         *string                        `json:"shortName,omitempty"`
	FullName          *string                        `json:"fullName,omitempty"`
	AnswerDirectlyName *string                       `json:"answerDirectlyName,omitempty"`
	MailBoxOwner      *RecipientApiModel             `json:"mailBoxOwner,omitempty"`
	ProfilePicture    *DownloadFileFromAulaArguments  `json:"profilePicture,omitempty"`
	Metadata          *string                        `json:"metadata,omitempty"`
}

// GetMessageInfoLightDto represents lightweight message info.
type GetMessageInfoLightDto struct {
	ThreadID    *int        `json:"threadId,omitempty"`
	Subject     *string     `json:"subject,omitempty"`
	IsSensitive bool        `json:"isSensitive"`
	Message     *MessageDto `json:"message,omitempty"`
}

// MessagesInThreadDto represents a full thread with messages.
type MessagesInThreadDto struct {
	ID                      *int64                        `json:"id,omitempty"`
	FirstMessage            *MessageDto                   `json:"firstMessage,omitempty"`
	Messages                []MessageDto                  `json:"messages,omitempty"`
	IsMarked                bool                          `json:"isMarked"`
	ThreadCreator           *MessagesInThreadRecipientDto `json:"threadCreator,omitempty"`
	ThreadStartedDateTime   *string                       `json:"threadStartedDateTime,omitempty"`
	Recipients              []MessagesInThreadRecipientDto `json:"recipients,omitempty"`
	MoreMessagesExist       bool                          `json:"moreMessagesExist"`
	TotalMessageCount       *int                          `json:"totalMessageCount,omitempty"`
	Page                    *int                          `json:"page,omitempty"`
	Subject                 *string                       `json:"subject,omitempty"`
	Muted                   bool                          `json:"muted"`
	Marked                  bool                          `json:"marked"`
	IsThreadForwarded       bool                          `json:"isThreadForwarded"`
	Sensitive               bool                          `json:"sensitive"`
	HasSecureDocuments      bool                          `json:"hasSecureDocuments"`
	MailboxOwner            *RecipientApiModel            `json:"mailboxOwner,omitempty"`
	ThreadEntityLinkDto     *ThreadEntityLinkDto          `json:"threadEntityLinkDto,omitempty"`
	FolderName              *string                       `json:"folderName,omitempty"`
}

// MessagesInThreadRecipientDto represents a recipient within a thread detail view.
type MessagesInThreadRecipientDto struct {
	LastReadMessageID  *string            `json:"lastReadMessageId,omitempty"`
	LastReadTimeStamp  *string            `json:"lastReadTimeStamp,omitempty"`
	LeaveTime          *string            `json:"leaveTime,omitempty"`
	DeletedAt          *string            `json:"deletedAt,omitempty"`
	FullName           *string            `json:"fullName,omitempty"`
	ShortName          *string            `json:"shortName,omitempty"`
	MailBoxOwner       *RecipientApiModel `json:"mailBoxOwner,omitempty"`
	Metadata           *string            `json:"metadata,omitempty"`
	ProfilePicture     *ProfilePictureDto `json:"profilePicture,omitempty"`
}

// MessagesStubbedChild represents a stubbed child reference in message recipient relations.
type MessagesStubbedChild struct {
	ID              *int64  `json:"id,omitempty"`
	DisplayName     *string `json:"displayName,omitempty"`
	Class           *string `json:"class,omitempty"`
	InstitutionName *string `json:"institutionName,omitempty"`
}

// MessageRecipientRelationDto represents relation info on a message recipient.
type MessageRecipientRelationDto struct {
	RelationType    *string                `json:"type,omitempty"`
	Class           *string                `json:"class,omitempty"`
	Children        []MessagesStubbedChild `json:"children,omitempty"`
	InstitutionName *string                `json:"institutionName,omitempty"`
}

// MessageParticipantRelationDto represents participant relation in message thread view.
type MessageParticipantRelationDto struct {
	RelationType    *string                `json:"type,omitempty"`
	InstitutionName *string                `json:"institutioName,omitempty"`
	Class           *string                `json:"class,omitempty"`
	Children        []MessagesStubbedChild `json:"children,omitempty"`
}

// MessagingParticipantDto represents a messaging participant.
type MessagingParticipantDto struct {
	AnswerDirectlyName *string            `json:"answerDirectlyName,omitempty"`
	ProfilePicture     *ProfilePictureDto `json:"profilePicture,omitempty"`
	MailBoxOwner       *RecipientApiModel `json:"mailBoxOwner,omitempty"`
	FullName           *string            `json:"fullName,omitempty"`
	Metadata           *string            `json:"metadata,omitempty"`
}

// DeleteMessageDto represents a deleted message marker.
type DeleteMessageDto struct {
	DeletedAt *string `json:"deletedAt,omitempty"`
}

// UpdateBundleMessageDto represents a bundle update notification DTO.
type UpdateBundleMessageDto struct {
	IsMarked           bool                       `json:"isMarked"`
	IsSensitive        bool                       `json:"isSensitive"`
	IsUnread           bool                       `json:"isUnread"`
	Thread             *MessageThreadSubscription `json:"thread,omitempty"`
	LastestMessageDate *string                    `json:"lastestMessageDate,omitempty"`
	IsMuted            bool                       `json:"isMuted"`
}

// MessageFileUrl represents a message file URL.
type MessageFileUrl struct {
	URL *string `json:"url,omitempty"`
}

// Folder represents a message folder.
type Folder struct {
	ID         *int    `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	FolderType *string `json:"type,omitempty"`
}

// CommonInboxesDto represents a common (shared) inbox.
type CommonInboxesDto struct {
	ID              *int64                    `json:"id,omitempty"`
	Name            *string                   `json:"name,omitempty"`
	Address         *string                   `json:"address,omitempty"`
	Folders         []Folder                  `json:"folders,omitempty"`
	Participants    []MessagingParticipantDto `json:"participants,omitempty"`
	InstitutionCode *string                   `json:"institutionCode,omitempty"`
	InstitutionName *string                   `json:"institutionName,omitempty"`
	CommonInboxType *string                   `json:"commonInboxType,omitempty"`
}

// SetAutoReplyArguments represents auto-reply configuration.
type SetAutoReplyArguments struct {
	ReplyText     *string `json:"replyText,omitempty"`
	EndDateTime   *string `json:"endDateTime,omitempty"`
	StartDateTime *string `json:"startDateTime,omitempty"`
}

// MessageAutoReplyResult represents auto-reply result.
type MessageAutoReplyResult struct {
	ID            *int64              `json:"id,omitempty"`
	ReplyText     *RichTextWrapperDto `json:"replyText,omitempty"`
	EndDateTime   *string             `json:"endDateTime,omitempty"`
	StartDateTime *string             `json:"startDateTime,omitempty"`
}

// MessageContentRequest represents message content for composing/replying.
type MessageContentRequest struct {
	AttachmentIDs []int64 `json:"attachmentIds,omitempty"`
	Text          *string `json:"text,omitempty"`
}

// StartNewThreadRequestArguments represents starting a new message thread.
type StartNewThreadRequestArguments struct {
	Message       *MessageContentRequest `json:"message,omitempty"`
	Subject       *string                `json:"subject,omitempty"`
	Recipients    []RecipientApiModel    `json:"recipients,omitempty"`
	BccRecipients []RecipientApiModel    `json:"bccRecipients,omitempty"`
	Sensitive     bool                   `json:"sensitive"`
	Creator       *RecipientApiModel     `json:"creator,omitempty"`
}

// ForwardInfoRequestArguments represents forward info when forwarding a thread.
type ForwardInfoRequestArguments struct {
	ForwardedThreadID       *int64   `json:"forwardedThreadId,omitempty"`
	ForwardedMessageIDs     []string `json:"forwardedMessageIds,omitempty"`
	DirectReply             bool     `json:"directReply"`
	ForwardSingleMessage    bool     `json:"forwardSingleMessage"`
	DirectReplyToCreator    bool     `json:"directReplyToCreator"`
}

// ForwardThreadRequestArguments represents forwarding a thread.
type ForwardThreadRequestArguments struct {
	Message       *MessageContentRequest       `json:"message,omitempty"`
	Subject       *string                      `json:"subject,omitempty"`
	Recipients    []RecipientApiModel          `json:"recipients,omitempty"`
	BccRecipients []RecipientApiModel          `json:"bccRecipients,omitempty"`
	Sensitive     bool                         `json:"sensitive"`
	Creator       *RecipientApiModel           `json:"creator,omitempty"`
	ForwardInfo   *ForwardInfoRequestArguments `json:"forwardInfo,omitempty"`
}

// ReplyMessageArgument represents replying to an existing thread.
type ReplyMessageArgument struct {
	ThreadID      *int64                 `json:"threadId,omitempty"`
	Message       *MessageContentRequest `json:"message,omitempty"`
	CommonInboxID *int64                 `json:"commonInboxId,omitempty"`
	BundleID      *int64                 `json:"bundleId,omitempty"`
}

// EditMessageRequest represents editing an existing message.
type EditMessageRequest struct {
	ThreadID       *int64                 `json:"threadId,omitempty"`
	CommonInboxID  *int64                 `json:"commonInboxId,omitempty"`
	BundleID       *int64                 `json:"bundleId,omitempty"`
	MessageID      *string                `json:"messageId,omitempty"`
	MessageRequest *MessageContentRequest `json:"messageRequest,omitempty"`
}

// AddRecipientArguments represents adding recipients to an existing thread.
type AddRecipientArguments struct {
	ThreadID      *int64              `json:"threadId,omitempty"`
	Recipients    []RecipientApiModel `json:"recipients,omitempty"`
	CommonInboxID *int64              `json:"commonInboxId,omitempty"`
}

// DeleteMessageRequest represents deleting a single message.
type DeleteMessageRequest struct {
	MessageID *string `json:"messageId,omitempty"`
	ThreadID  *int64  `json:"threadId,omitempty"`
}

// DeleteThreadArguments represents deleting thread(s).
type DeleteThreadArguments struct {
	SubscriptionIDs []int64 `json:"subscriptionIds,omitempty"`
	ThreadIDs       []int64 `json:"threadIds,omitempty"`
	CommonInboxID   *int64  `json:"commonInboxId,omitempty"`
}

// LeaveThreadArguments represents leaving a thread.
type LeaveThreadArguments struct {
	ThreadID *int64 `json:"threadId,omitempty"`
}

// LeaveThreadsRequest represents leaving multiple threads.
type LeaveThreadsRequest struct {
	SubscriptionIDs []int64 `json:"subscriptionIds,omitempty"`
}

// MarkThreadsRequest represents marking/unmarking threads.
type MarkThreadsRequest struct {
	Marked          bool    `json:"marked"`
	ThreadIDs       []int64 `json:"threadIds,omitempty"`
	SubscriptionIDs []int64 `json:"subscriptionIds,omitempty"`
	CommonInboxID   *int64  `json:"commonInboxId,omitempty"`
}

// MuteThreadRequestArguments represents muting/unmuting thread(s).
type MuteThreadRequestArguments struct {
	Muted           bool                `json:"muted"`
	Owner           *RecipientApiModel  `json:"MailBoxOwner,omitempty"`
	SubscriptionIDs []int64             `json:"subscriptionIds,omitempty"`
	CommonInboxID   *int64              `json:"commonInboxId,omitempty"`
	ThreadIDs       []int64             `json:"threadIds,omitempty"`
}

// SetLastMessageRequestArguments represents setting last-read message on a thread.
type SetLastMessageRequestArguments struct {
	MessageID     *string `json:"messageId,omitempty"`
	ThreadID      *int64  `json:"threadId,omitempty"`
	CommonInboxID *int64  `json:"commonInboxId,omitempty"`
}

// SetSensitivityLevelRequest represents setting sensitivity level on a thread.
type SetSensitivityLevelRequest struct {
	ThreadID         *int64 `json:"threadId,omitempty"`
	SensitivityLevel *int   `json:"sensitivityLevel,omitempty"`
	CommonInboxID    *int64 `json:"commonInboxId,omitempty"`
	BundleID         *int64 `json:"bundleId,omitempty"`
}

// UpdateMessageThreadsSubscriptionStatusRequest represents updating read/unread status.
type UpdateMessageThreadsSubscriptionStatusRequest struct {
	SubscriptionIDs []int64 `json:"subscriptionIds,omitempty"`
	IsRead          bool    `json:"isRead"`
}

// GetMessageInfoLightRequest represents getting a lightweight message preview.
type GetMessageInfoLightRequest struct {
	ThreadID      *int64  `json:"threadId,omitempty"`
	MessageID     *string `json:"messageId,omitempty"`
	CommonInboxID *int64  `json:"commonInboxId,omitempty"`
	OtpInboxID    *int64  `json:"otpInboxId,omitempty"`
}

// CreateFolderArguments represents creating a new folder.
type CreateFolderArguments struct {
	FolderName    *string `json:"folderName,omitempty"`
	CommonInboxID *int64  `json:"commonInboxId,omitempty"`
}

// GetFoldersArguments represents getting folders for a mailbox.
type GetFoldersArguments struct {
	IncludeDeletedFolders bool   `json:"includeDeletedFolders"`
	CommonInboxID         *int64 `json:"commonInboxId,omitempty"`
}

// MoveThreadsToFolderRequestArguments represents moving threads to a folder.
type MoveThreadsToFolderRequestArguments struct {
	ThreadIDs       []int64 `json:"threadIds,omitempty"`
	SubscriptionIDs []int64 `json:"subscriptionIds,omitempty"`
	FolderID        *int64  `json:"folderId,omitempty"`
	CommonInboxID   *int64  `json:"commonInboxId,omitempty"`
}

// UpdateFolderArguments represents updating (renaming) a folder.
type UpdateFolderArguments struct {
	FolderID   *int64  `json:"folderId,omitempty"`
	FolderName *string `json:"folderName,omitempty"`
}

// GetCommonInboxesArguments represents getting common inboxes.
type GetCommonInboxesArguments struct {
	InstitutionProfileIDs           []int64 `json:"institutionProfileIds,omitempty"`
	ShouldIncludeProfilePictureURL  bool    `json:"shouldIncludeProfilePictureUrl"`
}

// GetMessagesForThreadArguments represents getting messages in a thread (paginated).
type GetMessagesForThreadArguments struct {
	ThreadID      *int64 `json:"threadId,omitempty"`
	Page          *int   `json:"page,omitempty"`
	CommonInboxID *int64 `json:"commonInboxId,omitempty"`
}

// GetThreadListArguments represents getting thread list.
type GetThreadListArguments struct {
	FolderID         *int64  `json:"folderId,omitempty"`
	FilterType       *string `json:"filterType,omitempty"`
	SortType         *string `json:"sortType,omitempty"`
	SortOrder        *string `json:"sortOrder,omitempty"`
	Page             *int    `json:"page,omitempty"`
	ThreadIDs        []int64 `json:"threadIds,omitempty"`
	MailBoxOwnerType *string `json:"mailBoxOwnerType,omitempty"`
	MailBoxOwners    []int64 `json:"mailBoxOwners,omitempty"`
	ActiveChildren   []int64 `json:"activeChildren,omitempty"`
}

// GetThreadsInBundleArguments represents getting threads within a bundle.
type GetThreadsInBundleArguments struct {
	BundleID *int64 `json:"bundleId,omitempty"`
}

// AttachMessagesToSecureDocumentRequest represents attaching messages to a secure document.
type AttachMessagesToSecureDocumentRequest struct {
	SecureDocumentID *int64   `json:"secureDocumentId,omitempty"`
	MessageIDs       []string `json:"messageIds,omitempty"`
	ThreadID         *int64   `json:"threadId,omitempty"`
	CommonInboxID    *int64   `json:"commonInboxId,omitempty"`
}
