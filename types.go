package tegola

import (
	"net/http"
	"gopkg.in/mgo.v2"
)

const CommandRegString = `^\/(?P<command>[^@\s]+)@?(?:(?P<bot>\S+)|)\s?(?P<args>[\s\S]*)$`

type Bot struct {
	// Bot name
	BotName      string
	// Bot username
	BotUser      string
	// Bot token given by the BotFather
	BotToken     string
	// Preferred timezone to convert Unix time to date
	TimeZone     string
	// Admin users
	AdminUsers   []User
	// Admin chats
	AdminChats   []int64
	// Users allowed to interact with the bot (empty means everyone is allowed)
	AllowedUsers []User
	// Chats allowed to interact with the bot (empty means every chat is allowed)
	AllowedChats []int64
	// HTTP client to perform HTTP requests (GET, POST)
	Client       *http.Client
	// Debug flag for verbose logging
	Debug        bool
	// Log level (in case we decide to use it)
	LogLevel string
	// Size of the channel containing the updates to be streamed to the updates handler
	UpdatesChanSize int
	// Channel to stream the updated to the updates handler
	UpdatesChan     chan Update
	// Route where to listen for updates (updates through webhooks)
	ListenRoute     string
	// Port where to listen for updates (updates through webhooks)
	ListenPort      string
	// Base URL where to listen for updates (updates through webhooks)
	ListenURL	string
	MongoAuth mgo.DialInfo
	MongoSession *mgo.Session
	UpdateMode string
	Threads int
	LoopSleep int
}

type Updates struct {
	Ok      bool     `json:"ok"`
	Updates []Update `json:"result"`
}

type SendMessageConfirm struct {
	Ok      bool    `json:"ok"`
	Message Message `json:"result"`
}

// This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	// The update‘s unique identifier. Update identifiers start
	// from a certain positive number and increase sequentially.
	// This ID becomes especially handy if you’re using Webhooks,
	// since it allows you to ignore repeated updates or to restore the correct update sequence,
	// should they get out of order.
	UpdateId int64 `json:"update_id"`
	// Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message Message `json:"message"`
	// Optional. New version of a message that is known to the bot and was edited
	EditedMessage Message `json:"edited_message"`
	// Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	ChannelPost Message `json:"channel_post"`
	// Optional. New version of a channel post that is known to the bot and was edited
	EditedChannelPost Message `json:"edited_channel_post"`
	// Optional. New incoming inline query
	InlineQuery InlineQuery `json:"inline_query"`
	// Optional. The result of an inline query that was chosen by a user and sent to their chat partner.
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	// Optional. New incoming callback query
	CallbackQuery CallbackQuery `json:"callback_query"`
}

type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	Url string `json:"url"`
	// True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// Number of updates awaiting delivery
	PendingUpdateCount int64 `json:"pending_update_count"`
	// Optional.Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorDate int64 `json:"last_error_date"`
	// Optional.Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message"`
	// Optional.Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	MaxConnections int64 `json:"max_connections"`
	// Array of String
	AllowedUpdated []string `json:"allowed_updates"`
}

type InlineQuery struct {
}

type ChosenInlineResult struct {
}

type Message struct {
	// Unique message identifier inside this chat
	MessageId int64 `json:"message_id"`
	// Conversation the message belongs to
	Chat Chat `json:"chat"`
	// Date the message was sent in Unix time
	Date int64 `json:"date"`
	// Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Text string `json:"text"`
	// Optional. Message is a shared location, information about the location
	Location Location `json:"location"`
	// Optional. Sender, can be empty for messages sent to channels
	From User `json:"from"`
	// Optional. For forwarded messages, sender of the original message
	ForwardFrom User `json:"forward_from"`
	// Optional. For messages forwarded from a channel, information about the original channel
	ForwardFromChat Chat `json:"forward_from_chat"`
	// Optional. For forwarded channel posts, identifier of the original message in the channel
	ForwardFromMessageId int64 `json:"forward_from_message_id"`
	// Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardDate int64 `json:"forward_date"`
	// Optional. For replies, the original message.
	// Note that the Message object in this field will not contain
	// further reply_to_message fields even if it itself is a reply.
	ReplyToMessage SecondLevelMessage `json:"reply_to_message"`
	// Optional. Date the message was last edited in Unix time
	EditDate int64 `json:"edit_date"`
	// Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities"`
	// Optional. Message is an audio file, information about the file
	Audio Audio `json:"audio"`
	// Optional. Message is a general file, information about the file
	Document Document `json:"document"`
	// Optional. Message is a game, information about the game. More about games »
	Game Game `json:"game"`
	// Optional. Message is a photo, available sizes of the photo
	Photo []PhotoSize `json:"photo"`
	// Optional. Message is a sticker, information about the sticker
	Sticker Sticker `json:"sticker"`
	// Optional. Message is a video, information about the video
	Video Video `json:"video"`
	// Optional. Message is a voice message, information about the file
	Voice Voice `json:"voice"`
	// Optional. Caption for the document, photo or video, 0-200 characters
	Caption string `json:"caption"`
	// Optional. Message is a shared contact, information about the contact
	Contact Contact `json:"contact"`
	// Optional. Message is a venue, information about the venue
	Venue Venue `json:"venue"`
	// Optional. A new member was added to the group, information about them (this member may be the bot itself)
	NewChatMember User `json:"new_chat_member"`
	// Optional. A member was removed from the group, information about them (this member may be the bot itself)
	LeftChatMember User `json:"left_chat_member"`
	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title"`
	// Optional. A chat photo was change to this value
	NewChatPhoto []PhotoSize `json:"new_chat_photo"`
	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo"`
	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created"`
	// Optional. Service message: the supergroup has been created.
	// This field can‘t be received in a message coming through updates,
	// because bot can’t be a member of a supergroup when it is created.
	// It can only be found in reply_to_message if someone replies to a very
	// first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created"`
	// Optional. Service message: the channel has been created.
	// This field can‘t be received in a message coming through updates,
	// because bot can’t be a member of a channel when it is created.
	// It can only be found in reply_to_message if someone replies to a very
	// first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created"`
	// Optional. The group has been migrated to a supergroup with the specified identifier.
	// This number may be greater than 32 bits and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it smaller than 52 bits, so a signed 64 bit integer or double-precision
	// float type are safe for storing this identifier.
	MigrateToChatId int64 `json:"migrate_to_chat_id"`
	// Optional. The supergroup has been migrated from a group with the specified identifier.
	// This number may be greater than 32 bits and some programming languages may
	// have difficulty/silent defects in interpreting it.
	// But it smaller than 52 bits, so a signed 64 bit integer or double-precision
	// float type are safe for storing this identifier.
	MigrateFromChatId int64 `json:"migrate_from_chat_id"`
	// Optional. Specified message was pinned.
	// Note that the Message object in this field will not contain further
	// reply_to_message fields even if it is itself a reply.
	PinnedMessage SecondLevelMessage `json:"pinned_message"`
}

// Just like Message, but without ReplyToMessage, PinnedMessage to avoid recursion error on them
type SecondLevelMessage struct {
	MessageId             int64           `json:"message_id"`
	Chat                  Chat            `json:"chat"`
	Date                  int64           `json:"date"`
	Text                  string          `json:"text"`
	Location              Location        `json:"location"`
	From                  User            `json:"from"`
	ForwardFrom           User            `json:"forward_from"`
	ForwardFromChat       Chat            `json:"forward_from_chat"`
	ForwardFromMessageId  int64           `json:"forward_from_message_id"`
	ForwardDate           int64           `json:"forward_date"`
	EditDate              int64           `json:"edit_date"`
	Entities              []MessageEntity `json:"entities"`
	Audio                 Audio           `json:"audio"`
	Document              Document        `json:"document"`
	Game                  Game            `json:"game"`
	Photo                 []PhotoSize     `json:"photo"`
	Sticker               Sticker         `json:"sticker"`
	Video                 Video           `json:"video"`
	Voice                 Voice           `json:"voice"`
	Caption               string          `json:"caption"`
	Contact               Contact         `json:"contact"`
	Venue                 Venue           `json:"venue"`
	NewChatMember         User            `json:"new_chat_member"`
	LeftChatMember        User            `json:"left_chat_member"`
	NewChatTitle          string          `json:"new_chat_title"`
	NewChatPhoto          []PhotoSize     `json:"new_chat_photo"`
	DeleteChatPhoto       bool            `json:"delete_chat_photo"`
	GroupChatCreated      bool            `json:"group_chat_created"`
	SupergroupChatCreated bool            `json:"supergroup_chat_created"`
	ChannelChatCreated    bool            `json:"channel_chat_created"`
	MigrateToChatId       int64           `json:"migrate_to_chat_id"`
	MigrateFromChatId     int64           `json:"migrate_from_chat_id"`
}

// This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity. Can be mention (@username), hashtag, bot_command, url,
	// email, bold (bold text), italic (italic text), code (monowidth string),
	// pre (monowidth block), text_link (for clickable text URLs), text_mention
	// (for users without usernames)
	Type string
	// Offset in UTF-16 code units to the start of the entity
	Offset int64
	// Length of the entity in UTF-16 code units
	Length int64
	// Optional. For “text_link” only, url that will be opened after user taps on the text
	Url string
	// Optional. For “text_mention” only, the mentioned user
	User User
}

// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Photo width
	Width int64 `json:"width"`
	// Photo height
	Height int64 `json:"height"`
	// Optional. File size
	FileSize int64 `json:"file_size"`
}

// This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileId    string `json:"file_id"`   // Unique identifier for this file
	Duration  int64  `json:"duration"`  // Duration of the audio in seconds as defined by sender
	Title     string `json:"title"`     // Title of the audio as defined by sender or by audio tags
	Performer string `json:"performer"` // Optional. Performer of the audio as defined by sender or by audio tags
	MimeType  string `json:"mime_type"` // Optional. MIME type of the file as defined by sender
	File_Size int64  `json:"file_size"` // Optional. File size
}

// This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileId    string    `json:"file_id"`   // Unique file identifier
	Thumb     PhotoSize `json:"thumb"`     // Optional. Document thumbnail as defined by sender
	FileName  string    `json:"file_name"` // Optional. Original filename as defined by sender
	MimeType  string    `json:"mime_type"` // Optional. MIME type of the file as defined by sender
	File_Size int64     `json:"file_size"` // Optional. File size
}

// This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	//https://core.telegram.org/bots/api#game
}

// This object represents a sticker.
type Sticker struct {
	FileId    string    `json:"file_id"`   // Unique file identifier
	Width     int64     `json:"width"`     // Sticker width
	Height    int64     `json:"height"`    // Sticker height
	Thumb     PhotoSize `json:"thumb"`     // Optional. Sticker thumbnail in .webp or .jpg format
	Emoji     string    `json:"emoji"`     // Optional. Emoji associated with the sticker
	File_Size int64     `json:"file_size"` // Optional. File size
}

// This object represents a voice note.
type Voice struct {
	FileId    string    `json:"file_id"`   // Unique file identifier for this file
	Width     int64     `json:"width"`     // Video width as defined by sender
	Height    int64     `json:"height"`    // Video height as defined by sender
	Duration  int64     `json:"duration"`  // Duration of the video in seconds as defined by sender
	Thumb     PhotoSize `json:"thumb"`     // Optional. Video thumbnail
	MimeType  string    `json:"mime_type"` // Optional. Mime type of a file as defined by sender
	File_Size int64     `json:"file_size"` // Optional. File size
}

// This object represents a video file.
type Video struct {
	FileId    string `json:"file_id"`   // Unique identifier for this file
	Duration  int64  `json:"duration"`  // Duration of the audio in seconds as defined by sender
	MimeType  string `json:"mime_type"` // Optional. MIME type of the file as defined by sender
	File_Size int64  `json:"file_size"` // Optional. File size
}

// This object represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"` // Contact's phone number
	FirstName   string `json:"first_name"`   // Contact's first name
	LastName    string `json:"last_name"`    // Optional. Contact's last name
	UserId      int64  `json:"user_id"`      // Contact's user identifier in Telegram
}

// This object represents a venue.
type Venue struct {
	Location     Location `json:"location"`      // Venue location
	Title        string   `json:"title"`         //Name of the venue
	Address      string   `json:"address"`       // Address of the venue
	FoursquareId string   `json:"foursquare_id"` // Optional. Foursquare identifier of the venue
}

// This object represents a Telegram user or bot.
type User struct {
	Id        int64  `json:"id"`         // Unique identifier for this user or bot
	FirstName string `json:"first_name"` // User‘s or bot’s first name
	LastName  string `json:"last_name"`  // Optional. User‘s or bot’s last name
	Username  string `json:"username"`   // Optional. User‘s or bot’s username
}

// This object represents a chat.
type Chat struct {
	Id                          int64  `json:"id"`
	Username                    string `json:"username"`                       // Optional. Username, for private chats, supergroups and channels if available
	FirstName                   string `json:"first_name"`                     // Optional. First name of the other party in a private chat
	LastName                    string `json:"last_name"`                      // Optional. Last name of the other party in a private chat
	Type                        string `json:"type"`                           // can be either “private”, “group”, “supergroup” or “channel”
	Title                       string `json:"title"`                          // Optional. Title, for supergroups, channels and group chats
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"` // Optional. True if a group has ‘All Members Are Admins’ enabled.
}

// This object represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile.
type File struct {
	FileId   string `json:"file_id"`   // Unique identifier for this file
	FileSize int64  `json:"file_size"` //Optional. File size, if known
	FilePath string `json:"file_path"` //Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

// This object represents a point on the map.
type Location struct {
	Latitude  float64 `json:"latitude"`  // Latitude as defined by sender
	Longitude float64 `json:"longitude"` // Longitude as defined by sender
}

// This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct{}

// This object represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button.
// Optional fields are mutually exclusive.
type KeyboardButton struct{}

// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct{}

// This object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct{}

// This object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct{}

// This object represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot, the field message will be present.
// If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present.
// Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct{}

// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct{}

// This object contains information about one member of the chat.
type ChatMember struct{}

// Contains information about why a request was unsuccessfull.
type ResponseParameters struct{}

// This object represents the contents of a file to be uploaded.
// Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile struct{}
