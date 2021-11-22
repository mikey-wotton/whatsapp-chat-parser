package stats

import (
	"fmt"
	"github.com/mikey-wotton/whatsapp-chat-parser/participants"
	"strings"
)

type Individual struct {
	TotalMessages         int
	TotalMedia            int
	TotalSpecificWordUses int
}

func Acquire(chatters participants.Chatters, searchWord string) map[string]Individual {
	stats := make(map[string]Individual, len(chatters))

	for chatter, messages := range chatters {
		stats[chatter.String()] = Individual{
			TotalMessages: len(messages),
			TotalMedia:    countMediaMessages(messages),
			TotalSpecificWordUses: countWordUse(messages, searchWord),
		}
	}

	return stats
}

func MostUsedSpecificWord(chatters participants.Chatters, searchWord string) (participants.Gamer, int) {
	searchWord = strings.ToLower(searchWord)
	mostCount := 0
	mostWordUser := participants.GamerUnknown

	for chatter, messages := range chatters {
		usedCount := countWordUse(messages, searchWord)

		if usedCount > mostCount {
			mostCount = usedCount
			mostWordUser = chatter
		}
	}

	return mostWordUser, mostCount
}

func MostEmojis(chatters participants.Chatters) (participants.Gamer, int) {
	mostCount := 0
	mostEmojiUser := participants.GamerUnknown

	for chatter, messages := range chatters {
		emojiCount := countEmojis(messages)

		if emojiCount > mostCount {
			mostCount = emojiCount
			mostEmojiUser = chatter
		}
	}

	return mostEmojiUser, mostCount
}

func MostMessages(chatters participants.Chatters) (participants.Gamer, int) {
	mostCount := 0
	mostChatty := participants.GamerUnknown

	for chatter, messages := range chatters {
		count := len(messages)

		if count > mostCount {
			mostCount = count
			mostChatty = chatter
		}

	}

	return mostChatty, mostCount
}

func MostMedia(chatters participants.Chatters) (participants.Gamer, int) {
	mostCount := 0
	mostMediaSender := participants.GamerUnknown

	for chatter, messages := range chatters {
		count := countMediaMessages(messages)
		if count > mostCount {
			mostMediaSender = chatter
			mostCount = count
		}
	}

	return mostMediaSender, mostCount
}

func countEmojis(messages []string) int {
	totalEmojis := 0
	for _, message := range messages {
		totalEmojis += countEmoji(message)
	}

	return totalEmojis
}

func countMediaMessages(messages []string) int {
	count := 0
	for _, message := range messages {
		if isMediaMessage(message) {
			count++
		}
	}

	return count
}

func isMediaMessage(s string) bool {
	return strings.Contains(s, "<Media omitted>")
}

func countWordUse(messages []string, word string) int {
	count := 0
	for _, message := range messages {
		count += strings.Count(strings.ToLower(message), word)
	}

	return count
}
