package disgo

type Intent struct {
	BitCode int
}

var (
	AllIntent                     = Intent{32767}
	GuildsIntent                  = Intent{1 << 0}
	GuildsMembersIntent           = Intent{1 << 1}
	GuildsBansIntent              = Intent{1 << 2}
	GuildsEmojisAndStickersIntent = Intent{1 << 3}
	GuildIntegrationsIntent       = Intent{1 << 4}
	GuildWebhooksIntent           = Intent{1 << 5}
	GuildInvitesIntent            = Intent{1 << 6}
	GuildVoiceStatesIntent        = Intent{1 << 7}
	GuildPresencesIntent          = Intent{1 << 8}
	GuildMessagesIntent           = Intent{1 << 9}
	GuildMessageReactionsIntent   = Intent{1 << 10}
	GuildMessageTypingIntent      = Intent{1 << 11}
	DirectMessagesIntent          = Intent{1 << 12}
	DirectMessagesReactionsIntent = Intent{1 << 13}
	DirectMessagesTypingIntent    = Intent{1 << 14}
)
