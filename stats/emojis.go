package stats

type Emoji struct {
	Slug        string `json:"slug"`
	Character   string `json:"character"`
	UnicodeName string `json:"unicode_name"`
	CodePoint   string `json:"code_point"`
	Group       string `json:"group"`
	SubGroup    string `json:"sub_group"`
}

func countEmoji(s string) int {
	count := 0
	for _, r := range s {
		if isEmoji(r) {
			count++
		}
	}

	return count
}

func isEmoji(r rune) bool {
	_, ok := emojiMap[r]
	return ok
}
