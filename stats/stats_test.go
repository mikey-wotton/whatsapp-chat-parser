package stats

import (
	"github.com/mikey-wotton/whatsapp-chat-parser/participants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMostUsedSpecificWord(t *testing.T) {
	tests := map[string]struct {
		input      participants.Chatters
		searchWord string

		expected participants.Gamer
		expCount int
	}{
		"No Search term": {
			input: participants.Chatters{
				participants.GamerMike:  []string{"Blah blah: <Media omitted> blah blah", "👩 👩‍ ‍👦 ‍👦"},
				participants.GamerPaul:  []string{"👩‍👩‍👦‍👦❤️☠️㊗️❤️☠️㊗️", "hello👩‍👩‍👦‍👦"},
				participants.GamerDayle: []string{"❤️☠️㊗️", "❤️☠️㊗️"},
			},
			searchWord: "SomeRareWord",
			expected:   participants.GamerUnknown,
		},
		"Has search term": {
			input: participants.Chatters{
				participants.GamerMike:  []string{"Word", "Word", "Word", "word"},
				participants.GamerPaul:  []string{"Blah blah: <Media omitted> blah blah", "Blah blah: <Media omitted> blah blah"},
				participants.GamerDayle: []string{"Blah word blah", "two", "Word"},
			},
			searchWord: "WORD",
			expected:   participants.GamerMike,
			expCount:   4,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			output, count := MostUsedSpecificWord(test.input, test.searchWord)

			assert.Equal(t, test.expected, output)
			assert.Equal(t, test.expCount, count)
		})
	}
}

func TestCountEmoji(t *testing.T) {
	tests := map[string]struct {
		input participants.Chatters

		expected participants.Gamer
		expCount int
	}{
		"No Emojis": {
			input: participants.Chatters{
				participants.GamerMike:  []string{"Blah blah: <Media omitted> blah blah"},
				participants.GamerPaul:  []string{"Blah blah: <Media omitted> blah blah", "Blah blah: <Media omitted> blah blah"},
				participants.GamerDayle: []string{"Blah blah: <Media omitted> blah blah", "two", "three"},
			},
			expected: participants.GamerUnknown,
		},
		"Has Emojis": {
			input: participants.Chatters{
				participants.GamerMike:  []string{"Blah blah: <Media omitted> blah blah", "👩 👩‍ ‍👦 ‍👦"},
				participants.GamerPaul:  []string{"👩‍👩‍👦‍👦❤️☠️㊗️❤️☠️㊗️", "hello👩‍👩‍👦‍👦"},
				participants.GamerDayle: []string{"❤️☠️㊗️", "❤️☠️㊗️"},
			},
			expected: participants.GamerPaul,
			expCount: 14,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			output, count := MostEmojis(test.input)

			assert.Equal(t, test.expected, output)
			assert.Equal(t, test.expCount, count)
		})
	}
}

func TestMostMedia(t *testing.T) {
	tests := map[string]struct {
		input participants.Chatters

		expected participants.Gamer
		expCount int
	}{
		"No one sends media means Unknown is most media sender": {
			input: participants.Chatters{
				participants.GamerMike:  []string{"one"},
				participants.GamerPaul:  []string{"one", "two"},
				participants.GamerDayle: []string{"one", "two", "three"},
			},
			expected: participants.GamerUnknown,
		},
		"Dayle is most media sender": {
			input: participants.Chatters{
				participants.GamerMike:  []string{"Blah blah: <Media omitted> blah blah"},
				participants.GamerPaul:  []string{"Blah blah: <Media omitted> blah blah", "Blah blah: <Media omitted> blah blah"},
				participants.GamerDayle: []string{"Blah blah: <Media omitted> blah blah", "two", "three"},
			},
			expected: participants.GamerPaul,
			expCount: 2,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			output, count := MostMedia(test.input)

			assert.Equal(t, test.expected, output)
			assert.Equal(t, test.expCount, count)
		})
	}
}

func TestMostMessages(t *testing.T) {
	tests := map[string]struct {
		input participants.Chatters

		expected participants.Gamer
		expCount int
	}{
		"Dayle is most chatty": {
			input: participants.Chatters{
				participants.GamerMike:  []string{"one"},
				participants.GamerPaul:  []string{"one", "two"},
				participants.GamerDayle: []string{"one", "two", "three"},
			},
			expected: participants.GamerDayle,
			expCount: 3,
		},
		"No one chats means Unknown is most chatty": {
			input: participants.Chatters{
				participants.GamerMike:  []string{},
				participants.GamerPaul:  []string{},
				participants.GamerDayle: []string{},
			},
			expected: participants.GamerUnknown,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			output, count := MostMessages(test.input)

			assert.Equal(t, test.expected, output)
			assert.Equal(t, test.expCount, count)
		})
	}
}
