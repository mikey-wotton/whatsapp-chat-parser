package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mikey-wotton/whatsapp-chat-parser/parser"
	"github.com/mikey-wotton/whatsapp-chat-parser/stats"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	searchWord := flag.String("word", "noob", "look for the Gamer who uses this word the most")
	flag.Parse()
	//load file, smallish so memory not an issue
	file, err := os.Open("./parser/tests/chat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	chatters, err := parser.Parse(file)
	if err != nil {
		panic(err)
	}

	allStats := stats.Acquire(chatters, *searchWord)

	f, _ := json.MarshalIndent(allStats, "", " ")
	_ = ioutil.WriteFile("test.json", f, 0644)


	mostMedia, mostMediaCount := stats.MostMedia(chatters)
	mostMessages, mostMessagesCount := stats.MostMessages(chatters)
	mostEmojis, mostEmojisCount := stats.MostEmojis(chatters)
	mostSpecificWord, count := stats.MostUsedSpecificWord(chatters, *searchWord)

	fmt.Println(fmt.Sprintf("Most media posted: %s (%d times)", mostMedia, mostMediaCount))
	fmt.Println(fmt.Sprintf("Most messages sent: %s (%d times)", mostMessages, mostMessagesCount))
	fmt.Println(fmt.Sprintf("Most emojis sent: %s (%d times)", mostEmojis, mostEmojisCount))
	fmt.Println(fmt.Sprintf("Most %s used: %s (%d times)", *searchWord, mostSpecificWord, count))
}
