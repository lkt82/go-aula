package enums

// MessageType is the type of message in a thread.
type MessageType string

const (
	MessageTypeAllMessageRelatedType      MessageType = "allMessageRelatedType"
	MessageTypeMessage                    MessageType = "message"
	MessageTypeRecipientAdded             MessageType = "recipientAdded"
	MessageTypeRecipientRemoved           MessageType = "recipientRemoved"
	MessageTypeAutoReply                  MessageType = "autoReply"
	MessageTypeSystemForward              MessageType = "systemForward"
	MessageTypeSystemReply                MessageType = "systemReply"
	MessageTypeForward                    MessageType = "forward"
	MessageTypeOther                      MessageType = "other"
	MessageTypeRecipientsAdded            MessageType = "recipientsAdded"
	MessageTypeRecipientsRemoved          MessageType = "recipientsRemoved"
	MessageTypeMessageDeleted             MessageType = "messageDeleted"
	MessageTypeMessageEdited              MessageType = "messageEdited"
	MessageTypeSystemForwardSingleMessage MessageType = "systemForwardSingleMessage"
)

// SensitivityLevel is a sensitivity level for messages.
type SensitivityLevel string

const (
	SensitivityLevelLevel1 SensitivityLevel = "level1"
	SensitivityLevelLevel2 SensitivityLevel = "level2"
	SensitivityLevelLevel3 SensitivityLevel = "level3"
)

// SubscriptionStatus is the read/unread status of a thread subscription.
type SubscriptionStatus string

const (
	SubscriptionStatusRead   SubscriptionStatus = "read"
	SubscriptionStatusUnread SubscriptionStatus = "unread"
)

// RecipientType is the type of message recipient.
type RecipientType string

const (
	RecipientTypeProfile     RecipientType = "profile"
	RecipientTypeGroup       RecipientType = "group"
	RecipientTypeCommonInbox RecipientType = "commonInbox"
	RecipientTypeUnknown     RecipientType = "unknown"
)

// RecipientsTarget is the target area for recipients.
type RecipientsTarget string

const (
	RecipientsTargetMessageRecipients    RecipientsTarget = "messageRecipients"
	RecipientsTargetMessageBccRecipients RecipientsTarget = "messageBccRecipients"
	RecipientsTargetCalendarEvent        RecipientsTarget = "calendarEvent"
	RecipientsTargetSecureDocument       RecipientsTarget = "secureDocument"
	RecipientsTargetPost                 RecipientsTarget = "post"
)

// MessageThreadClickType is a click action on a message thread.
type MessageThreadClickType string

const (
	MessageThreadClickTypeItemClick   MessageThreadClickType = "itemClick"
	MessageThreadClickTypeMove        MessageThreadClickType = "move"
	MessageThreadClickTypeMark        MessageThreadClickType = "mark"
	MessageThreadClickTypeDelete      MessageThreadClickType = "delete"
	MessageThreadClickTypeMultiMove   MessageThreadClickType = "multiMove"
	MessageThreadClickTypeMultiMark   MessageThreadClickType = "multiMark"
	MessageThreadClickTypeMultiDelete MessageThreadClickType = "multiDelete"
)

// BundledMessageType is a bundled message display type.
type BundledMessageType string

const (
	BundledMessageTypeIsRegularMessage      BundledMessageType = "isRegularMessage"
	BundledMessageTypeFirstMessage          BundledMessageType = "firstMessage"
	BundledMessageTypeMiddleMessage         BundledMessageType = "middleMessage"
	BundledMessageTypeLastMessage           BundledMessageType = "lastMessage"
	BundledMessageTypePrimaryMessage        BundledMessageType = "primaryMessage"
	BundledMessageTypeSecondaryMessage      BundledMessageType = "secondaryMessage"
	BundledMessageTypeLastOfSecondaryMessage BundledMessageType = "lastOfSecondaryMessage"
)

// CommonInboxType is the type of common inbox.
type CommonInboxType string

const (
	CommonInboxTypeInstitutional     CommonInboxType = "institutional"
	CommonInboxTypeCrossInstitutional CommonInboxType = "crossInstitutional"
)

// MessageFormType is the form type when composing a message.
type MessageFormType string

const (
	MessageFormTypeStartNewThread                        MessageFormType = "startNewThread"
	MessageFormTypeReplyInThread                         MessageFormType = "replyInThread"
	MessageFormTypeForward                               MessageFormType = "forward"
	MessageFormTypeStartNewThreadWithUser                MessageFormType = "startNewThreadWithUser"
	MessageFormTypeReplyInThreadFromAnswerOptionsButton  MessageFormType = "replyInThreadFromAnswerOptionsButton"
	MessageFormTypeForwardSingleMessage                  MessageFormType = "forwardSingleMessage"
)

// MessageThreadCellMoreMenuActionEnum is a more-menu action on a message thread cell.
type MessageThreadCellMoreMenuActionEnum string

const (
	MessageThreadCellMoreMenuActionEnumMoveToFolder     MessageThreadCellMoreMenuActionEnum = "moveToFolder"
	MessageThreadCellMoreMenuActionEnumMarkAsImportant  MessageThreadCellMoreMenuActionEnum = "markAsImportant"
	MessageThreadCellMoreMenuActionEnumForward          MessageThreadCellMoreMenuActionEnum = "forward"
)

// SendMessageButton is a send message button option.
type SendMessageButton string

const (
	SendMessageButtonReplySingle SendMessageButton = "REPLY_SINGLE"
	SendMessageButtonReplyAll    SendMessageButton = "REPLY_ALL"
)

// SubscriptionType is the subscription type for message threads.
type SubscriptionType string

const (
	SubscriptionTypeBundle     SubscriptionType = "bundle"
	SubscriptionTypeBundleItem SubscriptionType = "bundleItem"
	SubscriptionTypeUnbundled  SubscriptionType = "unbundled"
)

// DropdownActionEnum is a dropdown action in thread details.
type DropdownActionEnum string

const (
	DropdownActionEnumAddRecipient           DropdownActionEnum = "addRecipient"
	DropdownActionEnumForwarding             DropdownActionEnum = "forwarding"
	DropdownActionEnumToggleMute             DropdownActionEnum = "toggleMute"
	DropdownActionEnumLeave                  DropdownActionEnum = "leave"
	DropdownActionEnumToggleSensitive        DropdownActionEnum = "toggleSensitive"
	DropdownActionEnumExportThreadToDocument DropdownActionEnum = "exportThreadToDocument"
	DropdownActionEnumMarkAsImportant        DropdownActionEnum = "markAsImportant"
	DropdownActionEnumMoveToFolder           DropdownActionEnum = "moveToFolder"
	DropdownActionEnumDelete                 DropdownActionEnum = "delete"
	DropdownActionEnumToggleReadStatus       DropdownActionEnum = "toggleReadStatus"
	DropdownActionEnumCreateDocument         DropdownActionEnum = "createDocument"
)

// MessageMoreOption is a more option on a single message.
type MessageMoreOption string

const (
	MessageMoreOptionEdit    MessageMoreOption = "edit"
	MessageMoreOptionDelete  MessageMoreOption = "delete"
	MessageMoreOptionForward MessageMoreOption = "forward"
)

// ThreadType is the thread type.
type ThreadType string

const (
	ThreadTypeThread                   ThreadType = "thread"
	ThreadTypeEventReminder            ThreadType = "eventReminder"
	ThreadTypeVacationRequestReminder  ThreadType = "vacationRequestReminder"
)

// RecipientApiType is a recipient API type.
type RecipientApiType string

const (
	RecipientApiTypeUnknown            RecipientApiType = "unknown"
	RecipientApiTypeInstitutionProfile RecipientApiType = "institutionProfile"
	RecipientApiTypeCommonInbox        RecipientApiType = "commonInbox"
	RecipientApiTypeOtpInbox           RecipientApiType = "otpInbox"
)

// FolderType is a folder type for message folders.
type FolderType string

const (
	FolderTypeNormal     FolderType = "normal"
	FolderTypeDeleted    FolderType = "deleted"
	FolderTypeButtonCell FolderType = "buttonCell"
)
