package parser

import (
	"github.com/mikey-wotton/whatsapp-chat-parser/participants"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetChatterName(t *testing.T) {
	chatter := getChatterName("08/11/2020, 05:48 - Mike: Rip Santosh, he fought a good fight")

	assert.Equal(t, participants.GamerMike, chatter)
}

func TestParse_totalMessage(t *testing.T) {
	//Arrange
	file, err := os.Open("./tests/chat.txt")
	assert.Nil(t, err)

	//Act
	output, err := Parse(file)
	assert.Nil(t, err)

	//Assess
	assert.Equal(t, 4, len(output[participants.GamerSantosh]))
}

func TestParse_singleLine(t *testing.T) {
	//Arrange
	file, err := os.Open("./tests/singleLine.txt")
	assert.Nil(t, err)

	expected := getAllChatters()
	expected[participants.GamerWezz] = append(expected[participants.GamerWezz], "24/07/2020, 21:03 - Wezz: Holy fuck. The express are printing it? Must be true. FUCKYOUALIENS")
	expected[participants.GamerDayle] = append(expected[participants.GamerDayle], "25/07/2020, 10:02 - Dayle: BANNED")
	expected[participants.GamerPaul] = append(expected[participants.GamerPaul], "25/07/2020, 10:25 - Paul: How many bananas could you eat in 1 sitting")

	//Act
	output, err := Parse(file)
	assert.Nil(t, err)

	//Assess
	assert.Equal(t, expected, output)
}

func TestParse_parseMultiLine(t *testing.T) {
	//Arrange
	file, err := os.Open("./tests/multiLine.txt")
	assert.Nil(t, err)

	expected := getAllChatters()
	expected[participants.GamerMike] = append(expected[participants.GamerMike], `24/07/2020, 20:14 - Mike: "So, what right do we have to claim the Sun as our own?"

Alien simp detected`)
	//Act
	output, err := Parse(file)
	assert.Nil(t, err)

	//Assess
	assert.Equal(t, expected, output)
}
