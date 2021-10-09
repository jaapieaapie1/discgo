package disgo

import (
	"fmt"
	"strconv"
	"testing"
)

// FDK>L
func TestAllIndividualIntents(t *testing.T) {
	o := ClientOptions{
		Intents: []Intent{
			GuildsIntent,
			GuildsMembersIntent,
			GuildsBansIntent,
			GuildsEmojisAndStickersIntent,
			GuildIntegrationsIntent,
			GuildWebhooksIntent,
			GuildInvitesIntent,
			GuildVoiceStatesIntent,
			GuildPresencesIntent,
			GuildMessagesIntent,
			GuildMessageReactionsIntent,
			GuildMessageTypingIntent,
			DirectMessagesIntent,
			DirectMessagesReactionsIntent,
			DirectMessagesTypingIntent,
		},
	}

	if o.GetIntentInt() != 32767 {
		fmt.Printf("The intent number was %s (%x) not 37767\n", strconv.Itoa(o.GetIntentInt()), o.GetIntentInt())
		t.Fail()
		return
	}
}

func TestAllIntents(t *testing.T) {
	o := ClientOptions{
		Intents: []Intent{
			AllIntent,
		},
	}

	if o.GetIntentInt() != 32767 {
		fmt.Printf("The intent number was %s (%x) not 37767\n", strconv.Itoa(o.GetIntentInt()), o.GetIntentInt())
		t.Fail()
		return
	}
}