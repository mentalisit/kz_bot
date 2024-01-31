package transmitter

import (
	"errors"
	"fmt"
	"kz_bot/pkg/logger"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

// A Transmitter represents a message manager for a single guild.
type Transmitter struct {
	session    *discordgo.Session
	guild      string
	title      string
	autoCreate bool

	// channelWebhooks maps from a channel ID to a webhook instance
	channelWebhooks map[string]*discordgo.Webhook

	mutex sync.RWMutex

	log *logger.Logger
}

// ErrWebhookNotFound is returned when a valid webhook for this channel/message combination does not exist
var ErrWebhookNotFound = errors.New("webhook for this channel and message does not exist")

// ErrPermissionDenied is returned if the bot does not have permission to manage webhooks.
//
// Bots can be granted a guild-wide permission and channel-specific permissions to manage webhooks.
// Despite potentially having guild-wide permission, channel specific overrides could deny a bot's permission to manage webhooks.
var ErrPermissionDenied = errors.New("missing 'Manage Webhooks' permission")

// New returns a new Transmitter given a Discord session, guild ID, and title.
func New(session *discordgo.Session, guild string, title string, autoCreate bool, log *logger.Logger) *Transmitter {
	return &Transmitter{
		session:    session,
		guild:      guild,
		title:      title,
		autoCreate: autoCreate,

		channelWebhooks: make(map[string]*discordgo.Webhook),

		log: log,
	}
}

// Send transmits a message to the given channel with the provided webhook data, and waits until Discord responds with message data.
func (t *Transmitter) Send(channelID string, params *discordgo.WebhookParams) (*discordgo.Message, error) {
	wh, err := t.getOrCreateWebhook(channelID)
	if err != nil {
		t.log.ErrorErr(err)
		return nil, err
	}

	msg, err := t.session.WebhookExecute(wh.ID, wh.Token, true, params)
	if err != nil {
		return nil, fmt.Errorf("execute failed: %w", err)
	}

	return msg, nil
}

// Edit will edit a message in a channel, if possible.
func (t *Transmitter) Edit(channelID string, messageID string, params *discordgo.WebhookParams) error {
	wh := t.getWebhook(channelID)

	if wh == nil {
		return ErrWebhookNotFound
	}

	uri := discordgo.EndpointWebhookToken(wh.ID, wh.Token) + "/messages/" + messageID
	_, err := t.session.RequestWithBucketID("PATCH", uri, params, discordgo.EndpointWebhookToken("", ""))
	if err != nil {
		return err
	}

	return nil
}

// HasWebhook checks whether the transmitter is using a particular webhook.
func (t *Transmitter) HasWebhook(id string) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	for _, wh := range t.channelWebhooks {
		if wh.ID == id {
			return true
		}
	}

	return false
}

// AddWebhook allows you to register a channel's webhook with the transmitter.
func (t *Transmitter) AddWebhook(channelID string, webhook *discordgo.Webhook) bool {
	t.log.Debug("Manually added webhook to channel " + channelID)
	t.mutex.Lock()
	defer t.mutex.Unlock()

	_, replaced := t.channelWebhooks[channelID]
	t.channelWebhooks[channelID] = webhook
	return replaced
}

// RefreshGuildWebhooks loads "relevant" webhooks into the transmitter, with careful permission handling.
//
// Notes:
//
// - A webhook is "relevant" if it was created by this bot -- the ApplicationID should match the bot's ID.
// - The term "having permission" means having the "Manage Webhooks" permission. See ErrPermissionDenied for more information.
// - This function is additive and will not unload previously loaded webhooks.
// - A nil channelIDs slice is treated the same as an empty one.
//
// If the bot has guild-wide permission:
//
// 1. it will load any "relevant" webhooks from the entire guild
// 2. the given slice is ignored
//
// If the bot does not have guild-wide permission:
//
// 1. it will load any "relevant" webhooks in each channel
// 2. a single error will be returned if any error occurs (incl. if there is no permission for any of these channels)
//
// If any channel has more than one "relevant" webhook, it will randomly pick one.
func (t *Transmitter) RefreshGuildWebhooks(channelIDs []string) error {
	t.log.Debug("Refreshing guild webhooks")

	botID, err := getDiscordUserID(t.session)
	if err != nil {
		return fmt.Errorf("could not get current user: %w", err)
	}

	// Get all existing webhooks
	hooks, err := t.session.GuildWebhooks(t.guild)
	if err != nil {
		switch {
		case isDiscordPermissionError(err):
			// We fallback on manually fetching hooks from individual channels
			// if we don't have the "Manage Webhooks" permission globally.
			// We can only do this if we were provided channelIDs, though.
			if len(channelIDs) == 0 {
				return ErrPermissionDenied
			}
			t.log.Debug("Missing global 'Manage Webhooks' permission, falling back on per-channel permission")
			return t.fetchChannelsHooks(channelIDs, botID)
		default:
			return fmt.Errorf("could not get webhooks: %w", err)
		}
	}

	t.log.Debug("Refreshing guild webhooks using global permission")
	t.assignHooksByAppID(hooks, botID, false)
	return nil
}

// createWebhook creates a webhook for a specific channel.
func (t *Transmitter) createWebhook(channel string) (*discordgo.Webhook, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	wh, err := t.session.WebhookCreate(channel, t.title+time.Now().Format(" 3:04:05PM"), "")
	if err != nil {
		t.log.ErrorErr(err)
		return nil, err
	}
	t.channelWebhooks[channel] = wh
	return wh, nil
}

func (t *Transmitter) getWebhook(channel string) *discordgo.Webhook {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	//t.log.Println("gethook 178")
	webhooks, err := t.session.ChannelWebhooks(channel)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var webhook *discordgo.Webhook
	for _, i := range webhooks {
		//fmt.Printf("%s\n", i.User.Username)
		if i.User.Bot && i.User.Username == t.session.State.User.Username {
			webhook = i
			return webhook
		}
	}
	if webhook == nil {
		webhookCreate, err1 := t.session.WebhookCreate(channel, t.title, "")
		if err1 != nil {
			fmt.Println(err1)
			return nil
		}
		fmt.Println("webhookCreateNILL", webhookCreate.Token, webhookCreate.ID)
		return webhookCreate
	}

	if len(webhooks) == 0 {
		webhookCreate, err1 := t.session.WebhookCreate(channel, t.title, "")
		if err1 != nil {
			fmt.Println(err1)
			return nil
		}
		fmt.Println("webhookCreate", webhookCreate.Token, webhookCreate.ID)
		return webhookCreate
	}

	return nil
}

func (t *Transmitter) getOrCreateWebhook(channelID string) (*discordgo.Webhook, error) {
	// If we have a webhook for this channel, immediately return it
	wh := t.getWebhook(channelID)
	if wh != nil {
		return wh, nil
	}
	//t.log.Println(209)
	// Early exit if we don't want to automatically create one
	if !t.autoCreate {
		return nil, ErrWebhookNotFound
	}
	//t.log.Println(214)
	t.log.Info("Creating a webhook for " + channelID)
	wh, err := t.createWebhook(channelID)
	if err != nil {
		t.log.ErrorErr(err)
		return nil, fmt.Errorf("could not create webhook: %w", err)
	}

	return wh, nil
}

// fetchChannelsHooks fetches hooks for the given channelIDs and calls assignHooksByAppID for each channel's hooks
func (t *Transmitter) fetchChannelsHooks(channelIDs []string, botID string) error {
	// For each channel, search for relevant hooks
	var failedHooks []string
	for _, channelID := range channelIDs {
		hooks, err := t.session.ChannelWebhooks(channelID)
		if err != nil {
			failedHooks = append(failedHooks, "\n- "+channelID+": "+err.Error())
			continue
		}
		t.assignHooksByAppID(hooks, botID, true)
	}

	// Compose an error if any hooks failed
	if len(failedHooks) > 0 {
		return errors.New("failed to fetch hooks:" + strings.Join(failedHooks, ""))
	}

	return nil
}

func (t *Transmitter) assignHooksByAppID(hooks []*discordgo.Webhook, appID string, channelTargeted bool) {
	logLine := "Picking up webhook"
	if channelTargeted {
		logLine += " (channel targeted)"
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	for _, wh := range hooks {
		if wh.ApplicationID != appID {
			continue
		}

		t.channelWebhooks[wh.ChannelID] = wh

	}
}
