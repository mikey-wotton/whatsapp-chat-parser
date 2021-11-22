package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mikey-wotton/whatsapp-chat-parser/participants"
)

const maxInt64 = 9223372036854775807

var findDateTimePrefix = regexp.MustCompile("^[0-9]+/[0-9]+/[0-9]+, [0-9]+:[0-9]+")

func Parse(file *os.File) (participants.Chatters, error) {
	chatters := getAllChatters()

	scanner := bufio.NewScanner(file)

	lineNum := 0
	var chatter participants.Gamer
	for scanner.Scan() {
		lineNum++
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("parsing line %d: %v", lineNum, err)
		}

		line := scanner.Text()

		if line == "" {
			line = "\n\n"
		}

		if isSameMessage(line) {
			chatters[chatter] = appendNewLineToLastMessage(chatters[chatter], line)
			continue
		}

		chatter = getChatterName(line)

		messages := chatters[chatter]
		messages = append(messages, line)
		chatters[chatter] = messages
	}

	return chatters, nil
}

func isSameMessage(line string) bool {
	return len(findDateTimePrefix.FindAllString(line, -1)) == 0
}

//Search for the gamer which has the first appearance within the line
//This is because a gamer's name could appear in the message later on, so we must search for first
func getChatterName(line string) participants.Gamer {
	lowestIndex := maxInt64
	gamer := participants.GamerUnknown
	for i := 0; i < len(participants.Gamers); i++ {
		name := participants.Gamers[i].String()

		firstAppearance := strings.Index(line, name)

		if firstAppearance == -1 {
			continue
		}

		if firstAppearance < lowestIndex {
			gamer = participants.Gamers[i]
			lowestIndex = firstAppearance
		}
	}

	return gamer
}

func appendNewLineToLastMessage(arr []string, new string) []string {
	lastMessage := arr[len(arr)-1]
	lastMessage += new
	arr[len(arr)-1] = lastMessage

	return arr
}

func getAllChatters() participants.Chatters {
	chatters := make(participants.Chatters, 0)
	for i := 0; i < len(participants.Gamers); i++ {
		chatters[participants.Gamers[i]] = make([]string, 0)
	}

	return chatters
}
