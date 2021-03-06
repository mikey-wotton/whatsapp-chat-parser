package stats

var (
	emojiMap = map[int32]Emoji{

		42: {
			Slug:        "keycap",
			Character:   "*️⃣",
			UnicodeName: "keycap: *",
			CodePoint:   "002A FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		48: {
			Slug:        "keycap-0",
			Character:   "0️⃣",
			UnicodeName: "keycap: 0",
			CodePoint:   "0030 FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		49: {
			Slug:        "keycap-1",
			Character:   "1️⃣",
			UnicodeName: "keycap: 1",
			CodePoint:   "0031 FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		50: {
			Slug:        "keycap-2",
			Character:   "2️⃣",
			UnicodeName: "keycap: 2",
			CodePoint:   "0032 FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		51: {
			Slug:        "keycap-3",
			Character:   "3️⃣",
			UnicodeName: "keycap: 3",
			CodePoint:   "0033 FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		52: {
			Slug:        "keycap-4",
			Character:   "4️⃣",
			UnicodeName: "keycap: 4",
			CodePoint:   "0034 FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		53: {
			Slug:        "keycap-5",
			Character:   "5️⃣",
			UnicodeName: "keycap: 5",
			CodePoint:   "0035 FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		54: {
			Slug:        "keycap-6",
			Character:   "6️⃣",
			UnicodeName: "keycap: 6",
			CodePoint:   "0036 FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		55: {
			Slug:        "keycap-7",
			Character:   "7️⃣",
			UnicodeName: "keycap: 7",
			CodePoint:   "0037 FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		56: {
			Slug:        "keycap-8",
			Character:   "8️⃣",
			UnicodeName: "keycap: 8",
			CodePoint:   "0038 FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		57: {
			Slug:        "keycap-9",
			Character:   "9️⃣",
			UnicodeName: "keycap: 9",
			CodePoint:   "0039 FE0F 20E3",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		169: {
			Slug:        "copyright",
			Character:   "©️",
			UnicodeName: "copyright",
			CodePoint:   "00A9 FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		174: {
			Slug:        "registered",
			Character:   "®️",
			UnicodeName: "registered",
			CodePoint:   "00AE FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		8252: {
			Slug:        "double-exclamation-mark",
			Character:   "‼️",
			UnicodeName: "double exclamation mark",
			CodePoint:   "203C FE0F",
			Group:       "symbols",
			SubGroup:    "punctuation",
		},

		8265: {
			Slug:        "exclamation-question-mark",
			Character:   "⁉️",
			UnicodeName: "exclamation question mark",
			CodePoint:   "2049 FE0F",
			Group:       "symbols",
			SubGroup:    "punctuation",
		},

		8482: {
			Slug:        "trade-mark",
			Character:   "™️",
			UnicodeName: "trade mark",
			CodePoint:   "2122 FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		8505: {
			Slug:        "information",
			Character:   "ℹ️",
			UnicodeName: "information",
			CodePoint:   "2139 FE0F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		8596: {
			Slug:        "left-right-arrow",
			Character:   "↔️",
			UnicodeName: "left-right arrow",
			CodePoint:   "2194 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		8597: {
			Slug:        "up-down-arrow",
			Character:   "↕️",
			UnicodeName: "up-down arrow",
			CodePoint:   "2195 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		8598: {
			Slug:        "up-left-arrow",
			Character:   "↖️",
			UnicodeName: "up-left arrow",
			CodePoint:   "2196 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		8599: {
			Slug:        "up-right-arrow",
			Character:   "↗️",
			UnicodeName: "up-right arrow",
			CodePoint:   "2197 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		8600: {
			Slug:        "down-right-arrow",
			Character:   "↘️",
			UnicodeName: "down-right arrow",
			CodePoint:   "2198 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		8601: {
			Slug:        "down-left-arrow",
			Character:   "↙️",
			UnicodeName: "down-left arrow",
			CodePoint:   "2199 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		8617: {
			Slug:        "right-arrow-curving-left",
			Character:   "↩️",
			UnicodeName: "right arrow curving left",
			CodePoint:   "21A9 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		8618: {
			Slug:        "left-arrow-curving-right",
			Character:   "↪️",
			UnicodeName: "left arrow curving right",
			CodePoint:   "21AA FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		8986: {
			Slug:        "e0-6-watch",
			Character:   "⌚",
			UnicodeName: "E0.6 watch",
			CodePoint:   "231A",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		8987: {
			Slug:        "e0-6-hourglass-done",
			Character:   "⌛",
			UnicodeName: "E0.6 hourglass done",
			CodePoint:   "231B",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		9000: {
			Slug:        "keyboard",
			Character:   "⌨️",
			UnicodeName: "keyboard",
			CodePoint:   "2328 FE0F",
			Group:       "objects",
			SubGroup:    "computer",
		},

		9167: {
			Slug:        "eject-button",
			Character:   "⏏️",
			UnicodeName: "eject button",
			CodePoint:   "23CF FE0F",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9193: {
			Slug:        "fast-forward-button",
			Character:   "⏩",
			UnicodeName: "fast-forward button",
			CodePoint:   "23E9",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9194: {
			Slug:        "fast-reverse-button",
			Character:   "⏪",
			UnicodeName: "fast reverse button",
			CodePoint:   "23EA",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9195: {
			Slug:        "fast-up-button",
			Character:   "⏫",
			UnicodeName: "fast up button",
			CodePoint:   "23EB",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9196: {
			Slug:        "fast-down-button",
			Character:   "⏬",
			UnicodeName: "fast down button",
			CodePoint:   "23EC",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9197: {
			Slug:        "next-track-button",
			Character:   "⏭️",
			UnicodeName: "next track button",
			CodePoint:   "23ED FE0F",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9198: {
			Slug:        "last-track-button",
			Character:   "⏮️",
			UnicodeName: "last track button",
			CodePoint:   "23EE FE0F",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9199: {
			Slug:        "play-or-pause-button",
			Character:   "⏯️",
			UnicodeName: "play or pause button",
			CodePoint:   "23EF FE0F",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9200: {
			Slug:        "e0-6-alarm-clock",
			Character:   "⏰",
			UnicodeName: "E0.6 alarm clock",
			CodePoint:   "23F0",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		9201: {
			Slug:        "e1-0-stopwatch",
			Character:   "⏱️",
			UnicodeName: "E1.0 stopwatch",
			CodePoint:   "23F1 FE0F",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		9202: {
			Slug:        "e1-0-timer-clock",
			Character:   "⏲️",
			UnicodeName: "E1.0 timer clock",
			CodePoint:   "23F2 FE0F",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		9203: {
			Slug:        "e0-6-hourglass-not-done",
			Character:   "⏳",
			UnicodeName: "E0.6 hourglass not done",
			CodePoint:   "23F3",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		9208: {
			Slug:        "pause-button",
			Character:   "⏸️",
			UnicodeName: "pause button",
			CodePoint:   "23F8 FE0F",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9209: {
			Slug:        "stop-button",
			Character:   "⏹️",
			UnicodeName: "stop button",
			CodePoint:   "23F9 FE0F",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9210: {
			Slug:        "record-button",
			Character:   "⏺️",
			UnicodeName: "record button",
			CodePoint:   "23FA FE0F",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9410: {
			Slug:        "circled-m",
			Character:   "Ⓜ️",
			UnicodeName: "circled M",
			CodePoint:   "24C2 FE0F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		9642: {
			Slug:        "black-small-square",
			Character:   "▪️",
			UnicodeName: "black small square",
			CodePoint:   "25AA FE0F",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		9643: {
			Slug:        "white-small-square",
			Character:   "▫️",
			UnicodeName: "white small square",
			CodePoint:   "25AB FE0F",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		9654: {
			Slug:        "play-button",
			Character:   "▶️",
			UnicodeName: "play button",
			CodePoint:   "25B6 FE0F",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9664: {
			Slug:        "reverse-button",
			Character:   "◀️",
			UnicodeName: "reverse button",
			CodePoint:   "25C0 FE0F",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		9723: {
			Slug:        "white-medium-square",
			Character:   "◻️",
			UnicodeName: "white medium square",
			CodePoint:   "25FB FE0F",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		9724: {
			Slug:        "black-medium-square",
			Character:   "◼️",
			UnicodeName: "black medium square",
			CodePoint:   "25FC FE0F",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		9725: {
			Slug:        "white-medium-small-square",
			Character:   "◽",
			UnicodeName: "white medium-small square",
			CodePoint:   "25FD",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		9726: {
			Slug:        "black-medium-small-square",
			Character:   "◾",
			UnicodeName: "black medium-small square",
			CodePoint:   "25FE",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		9728: {
			Slug:        "sun",
			Character:   "☀️",
			UnicodeName: "sun",
			CodePoint:   "2600 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9729: {
			Slug:        "cloud",
			Character:   "☁️",
			UnicodeName: "cloud",
			CodePoint:   "2601 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9730: {
			Slug:        "umbrella",
			Character:   "☂️",
			UnicodeName: "umbrella",
			CodePoint:   "2602 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9731: {
			Slug:        "snowman",
			Character:   "☃️",
			UnicodeName: "snowman",
			CodePoint:   "2603 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9732: {
			Slug:        "comet",
			Character:   "☄️",
			UnicodeName: "comet",
			CodePoint:   "2604 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9742: {
			Slug:        "telephone",
			Character:   "☎️",
			UnicodeName: "telephone",
			CodePoint:   "260E FE0F",
			Group:       "objects",
			SubGroup:    "phone",
		},

		9745: {
			Slug:        "check-box-with-check",
			Character:   "☑️",
			UnicodeName: "check box with check",
			CodePoint:   "2611 FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		9748: {
			Slug:        "umbrella-with-rain-drops",
			Character:   "☔",
			UnicodeName: "umbrella with rain drops",
			CodePoint:   "2614",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9749: {
			Slug:        "e0-6-hot-beverage",
			Character:   "☕",
			UnicodeName: "E0.6 hot beverage",
			CodePoint:   "2615",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		9752: {
			Slug:        "e1-0-shamrock",
			Character:   "☘️",
			UnicodeName: "E1.0 shamrock",
			CodePoint:   "2618 FE0F",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		9757: {
			Slug:        "e1-0-index-pointing-up-dark-skin-tone",
			Character:   "☝🏿",
			UnicodeName: "E1.0 index pointing up: dark skin tone",
			CodePoint:   "261D 1F3FF",
			Group:       "people-body",
			SubGroup:    "hand-single-finger",
		},

		9760: {
			Slug:        "e1-0-skull-and-crossbones",
			Character:   "☠️",
			UnicodeName: "E1.0 skull and crossbones",
			CodePoint:   "2620 FE0F",
			Group:       "smileys-emotion",
			SubGroup:    "face-negative",
		},

		9762: {
			Slug:        "radioactive",
			Character:   "☢️",
			UnicodeName: "radioactive",
			CodePoint:   "2622 FE0F",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		9763: {
			Slug:        "biohazard",
			Character:   "☣️",
			UnicodeName: "biohazard",
			CodePoint:   "2623 FE0F",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		9766: {
			Slug:        "orthodox-cross",
			Character:   "☦️",
			UnicodeName: "orthodox cross",
			CodePoint:   "2626 FE0F",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		9770: {
			Slug:        "star-and-crescent",
			Character:   "☪️",
			UnicodeName: "star and crescent",
			CodePoint:   "262A FE0F",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		9774: {
			Slug:        "peace-symbol",
			Character:   "☮️",
			UnicodeName: "peace symbol",
			CodePoint:   "262E FE0F",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		9775: {
			Slug:        "yin-yang",
			Character:   "☯️",
			UnicodeName: "yin yang",
			CodePoint:   "262F FE0F",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		9784: {
			Slug:        "wheel-of-dharma",
			Character:   "☸️",
			UnicodeName: "wheel of dharma",
			CodePoint:   "2638 FE0F",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		9785: {
			Slug:        "e0-7-frowning-face",
			Character:   "☹️",
			UnicodeName: "E0.7 frowning face",
			CodePoint:   "2639 FE0F",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		9786: {
			Slug:        "e0-6-smiling-face",
			Character:   "☺️",
			UnicodeName: "E0.6 smiling face",
			CodePoint:   "263A FE0F",
			Group:       "smileys-emotion",
			SubGroup:    "face-affection",
		},

		9792: {
			Slug:        "female-sign",
			Character:   "♀️",
			UnicodeName: "female sign",
			CodePoint:   "2640 FE0F",
			Group:       "symbols",
			SubGroup:    "gender",
		},

		9794: {
			Slug:        "male-sign",
			Character:   "♂️",
			UnicodeName: "male sign",
			CodePoint:   "2642 FE0F",
			Group:       "symbols",
			SubGroup:    "gender",
		},

		9800: {
			Slug:        "aries",
			Character:   "♈",
			UnicodeName: "Aries",
			CodePoint:   "2648",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9801: {
			Slug:        "taurus",
			Character:   "♉",
			UnicodeName: "Taurus",
			CodePoint:   "2649",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9802: {
			Slug:        "gemini",
			Character:   "♊",
			UnicodeName: "Gemini",
			CodePoint:   "264A",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9803: {
			Slug:        "cancer",
			Character:   "♋",
			UnicodeName: "Cancer",
			CodePoint:   "264B",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9804: {
			Slug:        "leo",
			Character:   "♌",
			UnicodeName: "Leo",
			CodePoint:   "264C",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9805: {
			Slug:        "virgo",
			Character:   "♍",
			UnicodeName: "Virgo",
			CodePoint:   "264D",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9806: {
			Slug:        "libra",
			Character:   "♎",
			UnicodeName: "Libra",
			CodePoint:   "264E",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9807: {
			Slug:        "scorpio",
			Character:   "♏",
			UnicodeName: "Scorpio",
			CodePoint:   "264F",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9808: {
			Slug:        "sagittarius",
			Character:   "♐",
			UnicodeName: "Sagittarius",
			CodePoint:   "2650",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9809: {
			Slug:        "capricorn",
			Character:   "♑",
			UnicodeName: "Capricorn",
			CodePoint:   "2651",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9810: {
			Slug:        "aquarius",
			Character:   "♒",
			UnicodeName: "Aquarius",
			CodePoint:   "2652",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9811: {
			Slug:        "pisces",
			Character:   "♓",
			UnicodeName: "Pisces",
			CodePoint:   "2653",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9823: {
			Slug:        "chess-pawn",
			Character:   "♟️",
			UnicodeName: "chess pawn",
			CodePoint:   "265F FE0F",
			Group:       "activities",
			SubGroup:    "game",
		},

		9824: {
			Slug:        "spade-suit",
			Character:   "♠️",
			UnicodeName: "spade suit",
			CodePoint:   "2660 FE0F",
			Group:       "activities",
			SubGroup:    "game",
		},

		9827: {
			Slug:        "club-suit",
			Character:   "♣️",
			UnicodeName: "club suit",
			CodePoint:   "2663 FE0F",
			Group:       "activities",
			SubGroup:    "game",
		},

		9829: {
			Slug:        "heart-suit",
			Character:   "♥️",
			UnicodeName: "heart suit",
			CodePoint:   "2665 FE0F",
			Group:       "activities",
			SubGroup:    "game",
		},

		9830: {
			Slug:        "diamond-suit",
			Character:   "♦️",
			UnicodeName: "diamond suit",
			CodePoint:   "2666 FE0F",
			Group:       "activities",
			SubGroup:    "game",
		},

		9832: {
			Slug:        "e0-6-hot-springs",
			Character:   "♨️",
			UnicodeName: "E0.6 hot springs",
			CodePoint:   "2668 FE0F",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		9851: {
			Slug:        "recycling-symbol",
			Character:   "♻️",
			UnicodeName: "recycling symbol",
			CodePoint:   "267B FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		9854: {
			Slug:        "infinity",
			Character:   "♾️",
			UnicodeName: "infinity",
			CodePoint:   "267E FE0F",
			Group:       "symbols",
			SubGroup:    "math",
		},

		9855: {
			Slug:        "wheelchair-symbol",
			Character:   "♿",
			UnicodeName: "wheelchair symbol",
			CodePoint:   "267F",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		9874: {
			Slug:        "hammer-and-pick",
			Character:   "⚒️",
			UnicodeName: "hammer and pick",
			CodePoint:   "2692 FE0F",
			Group:       "objects",
			SubGroup:    "tool",
		},

		9875: {
			Slug:        "e0-6-anchor",
			Character:   "⚓",
			UnicodeName: "E0.6 anchor",
			CodePoint:   "2693",
			Group:       "travel-places",
			SubGroup:    "transport-water",
		},

		9876: {
			Slug:        "crossed-swords",
			Character:   "⚔️",
			UnicodeName: "crossed swords",
			CodePoint:   "2694 FE0F",
			Group:       "objects",
			SubGroup:    "tool",
		},

		9877: {
			Slug:        "medical-symbol",
			Character:   "⚕️",
			UnicodeName: "medical symbol",
			CodePoint:   "2695 FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		9878: {
			Slug:        "balance-scale",
			Character:   "⚖️",
			UnicodeName: "balance scale",
			CodePoint:   "2696 FE0F",
			Group:       "objects",
			SubGroup:    "tool",
		},

		9879: {
			Slug:        "alembic",
			Character:   "⚗️",
			UnicodeName: "alembic",
			CodePoint:   "2697 FE0F",
			Group:       "objects",
			SubGroup:    "science",
		},

		9881: {
			Slug:        "gear",
			Character:   "⚙️",
			UnicodeName: "gear",
			CodePoint:   "2699 FE0F",
			Group:       "objects",
			SubGroup:    "tool",
		},

		9883: {
			Slug:        "atom-symbol",
			Character:   "⚛️",
			UnicodeName: "atom symbol",
			CodePoint:   "269B FE0F",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		9884: {
			Slug:        "fleur-de-lis",
			Character:   "⚜️",
			UnicodeName: "fleur-de-lis",
			CodePoint:   "269C FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		9888: {
			Slug:        "warning",
			Character:   "⚠️",
			UnicodeName: "warning",
			CodePoint:   "26A0 FE0F",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		9889: {
			Slug:        "high-voltage",
			Character:   "⚡",
			UnicodeName: "high voltage",
			CodePoint:   "26A1",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9895: {
			Slug:        "transgender-symbol",
			Character:   "⚧️",
			UnicodeName: "transgender symbol",
			CodePoint:   "26A7 FE0F",
			Group:       "symbols",
			SubGroup:    "gender",
		},

		9898: {
			Slug:        "white-circle",
			Character:   "⚪",
			UnicodeName: "white circle",
			CodePoint:   "26AA",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		9899: {
			Slug:        "black-circle",
			Character:   "⚫",
			UnicodeName: "black circle",
			CodePoint:   "26AB",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		9904: {
			Slug:        "coffin",
			Character:   "⚰️",
			UnicodeName: "coffin",
			CodePoint:   "26B0 FE0F",
			Group:       "objects",
			SubGroup:    "other-object",
		},

		9905: {
			Slug:        "funeral-urn",
			Character:   "⚱️",
			UnicodeName: "funeral urn",
			CodePoint:   "26B1 FE0F",
			Group:       "objects",
			SubGroup:    "other-object",
		},

		9917: {
			Slug:        "soccer-ball",
			Character:   "⚽",
			UnicodeName: "soccer ball",
			CodePoint:   "26BD",
			Group:       "activities",
			SubGroup:    "sport",
		},

		9918: {
			Slug:        "baseball",
			Character:   "⚾",
			UnicodeName: "baseball",
			CodePoint:   "26BE",
			Group:       "activities",
			SubGroup:    "sport",
		},

		9924: {
			Slug:        "snowman-without-snow",
			Character:   "⛄",
			UnicodeName: "snowman without snow",
			CodePoint:   "26C4",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9925: {
			Slug:        "sun-behind-cloud",
			Character:   "⛅",
			UnicodeName: "sun behind cloud",
			CodePoint:   "26C5",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9928: {
			Slug:        "cloud-with-lightning-and-rain",
			Character:   "⛈️",
			UnicodeName: "cloud with lightning and rain",
			CodePoint:   "26C8 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9934: {
			Slug:        "ophiuchus",
			Character:   "⛎",
			UnicodeName: "Ophiuchus",
			CodePoint:   "26CE",
			Group:       "symbols",
			SubGroup:    "zodiac",
		},

		9935: {
			Slug:        "pick",
			Character:   "⛏️",
			UnicodeName: "pick",
			CodePoint:   "26CF FE0F",
			Group:       "objects",
			SubGroup:    "tool",
		},

		9937: {
			Slug:        "rescue-worker-s-helmet",
			Character:   "⛑️",
			UnicodeName: "rescue worker’s helmet",
			CodePoint:   "26D1 FE0F",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		9939: {
			Slug:        "chains",
			Character:   "⛓️",
			UnicodeName: "chains",
			CodePoint:   "26D3 FE0F",
			Group:       "objects",
			SubGroup:    "tool",
		},

		9940: {
			Slug:        "no-entry",
			Character:   "⛔",
			UnicodeName: "no entry",
			CodePoint:   "26D4",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		9961: {
			Slug:        "e0-7-shinto-shrine",
			Character:   "⛩️",
			UnicodeName: "E0.7 shinto shrine",
			CodePoint:   "26E9 FE0F",
			Group:       "travel-places",
			SubGroup:    "place-religious",
		},

		9962: {
			Slug:        "e0-6-church",
			Character:   "⛪",
			UnicodeName: "E0.6 church",
			CodePoint:   "26EA",
			Group:       "travel-places",
			SubGroup:    "place-religious",
		},

		9968: {
			Slug:        "e0-7-mountain",
			Character:   "⛰️",
			UnicodeName: "E0.7 mountain",
			CodePoint:   "26F0 FE0F",
			Group:       "travel-places",
			SubGroup:    "place-geographic",
		},

		9969: {
			Slug:        "umbrella-on-ground",
			Character:   "⛱️",
			UnicodeName: "umbrella on ground",
			CodePoint:   "26F1 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		9970: {
			Slug:        "e0-6-fountain",
			Character:   "⛲",
			UnicodeName: "E0.6 fountain",
			CodePoint:   "26F2",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		9971: {
			Slug:        "flag-in-hole",
			Character:   "⛳",
			UnicodeName: "flag in hole",
			CodePoint:   "26F3",
			Group:       "activities",
			SubGroup:    "sport",
		},

		9972: {
			Slug:        "e0-7-ferry",
			Character:   "⛴️",
			UnicodeName: "E0.7 ferry",
			CodePoint:   "26F4 FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-water",
		},

		9973: {
			Slug:        "e0-6-sailboat",
			Character:   "⛵",
			UnicodeName: "E0.6 sailboat",
			CodePoint:   "26F5",
			Group:       "travel-places",
			SubGroup:    "transport-water",
		},

		9975: {
			Slug:        "e0-7-skier",
			Character:   "⛷️",
			UnicodeName: "E0.7 skier",
			CodePoint:   "26F7 FE0F",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		9976: {
			Slug:        "ice-skate",
			Character:   "⛸️",
			UnicodeName: "ice skate",
			CodePoint:   "26F8 FE0F",
			Group:       "activities",
			SubGroup:    "sport",
		},

		9977: {
			Slug:        "e4-0-woman-bouncing-ball-dark-skin-tone",
			Character:   "⛹🏿‍♀️",
			UnicodeName: "E4.0 woman bouncing ball: dark skin tone",
			CodePoint:   "26F9 1F3FF 200D 2640 FE0F",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		9978: {
			Slug:        "e0-6-tent",
			Character:   "⛺",
			UnicodeName: "E0.6 tent",
			CodePoint:   "26FA",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		9981: {
			Slug:        "e0-6-fuel-pump",
			Character:   "⛽",
			UnicodeName: "E0.6 fuel pump",
			CodePoint:   "26FD",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		9986: {
			Slug:        "scissors",
			Character:   "✂️",
			UnicodeName: "scissors",
			CodePoint:   "2702 FE0F",
			Group:       "objects",
			SubGroup:    "office",
		},

		9989: {
			Slug:        "check-mark-button",
			Character:   "✅",
			UnicodeName: "check mark button",
			CodePoint:   "2705",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		9992: {
			Slug:        "e0-6-airplane",
			Character:   "✈️",
			UnicodeName: "E0.6 airplane",
			CodePoint:   "2708 FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		9993: {
			Slug:        "envelope",
			Character:   "✉️",
			UnicodeName: "envelope",
			CodePoint:   "2709 FE0F",
			Group:       "objects",
			SubGroup:    "mail",
		},

		9994: {
			Slug:        "e0-6-raised-fist",
			Character:   "✊",
			UnicodeName: "E0.6 raised fist",
			CodePoint:   "270A",
			Group:       "people-body",
			SubGroup:    "hand-fingers-closed",
		},

		9995: {
			Slug:        "e0-6-raised-hand",
			Character:   "✋",
			UnicodeName: "E0.6 raised hand",
			CodePoint:   "270B",
			Group:       "people-body",
			SubGroup:    "hand-fingers-open",
		},

		9996: {
			Slug:        "e1-0-victory-hand-dark-skin-tone",
			Character:   "✌🏿",
			UnicodeName: "E1.0 victory hand: dark skin tone",
			CodePoint:   "270C 1F3FF",
			Group:       "people-body",
			SubGroup:    "hand-fingers-partial",
		},

		9997: {
			Slug:        "e1-0-writing-hand-dark-skin-tone",
			Character:   "✍🏿",
			UnicodeName: "E1.0 writing hand: dark skin tone",
			CodePoint:   "270D 1F3FF",
			Group:       "people-body",
			SubGroup:    "hand-prop",
		},

		9999: {
			Slug:        "pencil",
			Character:   "✏️",
			UnicodeName: "pencil",
			CodePoint:   "270F FE0F",
			Group:       "objects",
			SubGroup:    "writing",
		},

		10002: {
			Slug:        "black-nib",
			Character:   "✒️",
			UnicodeName: "black nib",
			CodePoint:   "2712 FE0F",
			Group:       "objects",
			SubGroup:    "writing",
		},

		10004: {
			Slug:        "check-mark",
			Character:   "✔️",
			UnicodeName: "check mark",
			CodePoint:   "2714 FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		10006: {
			Slug:        "multiply",
			Character:   "✖️",
			UnicodeName: "multiply",
			CodePoint:   "2716 FE0F",
			Group:       "symbols",
			SubGroup:    "math",
		},

		10013: {
			Slug:        "latin-cross",
			Character:   "✝️",
			UnicodeName: "latin cross",
			CodePoint:   "271D FE0F",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		10017: {
			Slug:        "star-of-david",
			Character:   "✡️",
			UnicodeName: "star of David",
			CodePoint:   "2721 FE0F",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		10024: {
			Slug:        "sparkles",
			Character:   "✨",
			UnicodeName: "sparkles",
			CodePoint:   "2728",
			Group:       "activities",
			SubGroup:    "event",
		},

		10035: {
			Slug:        "eight-spoked-asterisk",
			Character:   "✳️",
			UnicodeName: "eight-spoked asterisk",
			CodePoint:   "2733 FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		10036: {
			Slug:        "eight-pointed-star",
			Character:   "✴️",
			UnicodeName: "eight-pointed star",
			CodePoint:   "2734 FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		10052: {
			Slug:        "snowflake",
			Character:   "❄️",
			UnicodeName: "snowflake",
			CodePoint:   "2744 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		10055: {
			Slug:        "sparkle",
			Character:   "❇️",
			UnicodeName: "sparkle",
			CodePoint:   "2747 FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		10060: {
			Slug:        "cross-mark",
			Character:   "❌",
			UnicodeName: "cross mark",
			CodePoint:   "274C",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		10062: {
			Slug:        "cross-mark-button",
			Character:   "❎",
			UnicodeName: "cross mark button",
			CodePoint:   "274E",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		10067: {
			Slug:        "question-mark",
			Character:   "❓",
			UnicodeName: "question mark",
			CodePoint:   "2753",
			Group:       "symbols",
			SubGroup:    "punctuation",
		},

		10068: {
			Slug:        "white-question-mark",
			Character:   "❔",
			UnicodeName: "white question mark",
			CodePoint:   "2754",
			Group:       "symbols",
			SubGroup:    "punctuation",
		},

		10069: {
			Slug:        "white-exclamation-mark",
			Character:   "❕",
			UnicodeName: "white exclamation mark",
			CodePoint:   "2755",
			Group:       "symbols",
			SubGroup:    "punctuation",
		},

		10071: {
			Slug:        "exclamation-mark",
			Character:   "❗",
			UnicodeName: "exclamation mark",
			CodePoint:   "2757",
			Group:       "symbols",
			SubGroup:    "punctuation",
		},

		10083: {
			Slug:        "e1-0-heart-exclamation",
			Character:   "❣️",
			UnicodeName: "E1.0 heart exclamation",
			CodePoint:   "2763 FE0F",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		10084: {
			Slug:        "e0-6-red-heart",
			Character:   "❤️",
			UnicodeName: "E0.6 red heart",
			CodePoint:   "2764 FE0F",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		10133: {
			Slug:        "plus",
			Character:   "➕",
			UnicodeName: "plus",
			CodePoint:   "2795",
			Group:       "symbols",
			SubGroup:    "math",
		},

		10134: {
			Slug:        "minus",
			Character:   "➖",
			UnicodeName: "minus",
			CodePoint:   "2796",
			Group:       "symbols",
			SubGroup:    "math",
		},

		10135: {
			Slug:        "divide",
			Character:   "➗",
			UnicodeName: "divide",
			CodePoint:   "2797",
			Group:       "symbols",
			SubGroup:    "math",
		},

		10145: {
			Slug:        "right-arrow",
			Character:   "➡️",
			UnicodeName: "right arrow",
			CodePoint:   "27A1 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		10160: {
			Slug:        "curly-loop",
			Character:   "➰",
			UnicodeName: "curly loop",
			CodePoint:   "27B0",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		10175: {
			Slug:        "double-curly-loop",
			Character:   "➿",
			UnicodeName: "double curly loop",
			CodePoint:   "27BF",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		10548: {
			Slug:        "right-arrow-curving-up",
			Character:   "⤴️",
			UnicodeName: "right arrow curving up",
			CodePoint:   "2934 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		10549: {
			Slug:        "right-arrow-curving-down",
			Character:   "⤵️",
			UnicodeName: "right arrow curving down",
			CodePoint:   "2935 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		11013: {
			Slug:        "left-arrow",
			Character:   "⬅️",
			UnicodeName: "left arrow",
			CodePoint:   "2B05 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		11014: {
			Slug:        "up-arrow",
			Character:   "⬆️",
			UnicodeName: "up arrow",
			CodePoint:   "2B06 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		11015: {
			Slug:        "down-arrow",
			Character:   "⬇️",
			UnicodeName: "down arrow",
			CodePoint:   "2B07 FE0F",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		11035: {
			Slug:        "black-large-square",
			Character:   "⬛",
			UnicodeName: "black large square",
			CodePoint:   "2B1B",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		11036: {
			Slug:        "white-large-square",
			Character:   "⬜",
			UnicodeName: "white large square",
			CodePoint:   "2B1C",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		11088: {
			Slug:        "star",
			Character:   "⭐",
			UnicodeName: "star",
			CodePoint:   "2B50",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		11093: {
			Slug:        "hollow-red-circle",
			Character:   "⭕",
			UnicodeName: "hollow red circle",
			CodePoint:   "2B55",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		12336: {
			Slug:        "wavy-dash",
			Character:   "〰️",
			UnicodeName: "wavy dash",
			CodePoint:   "3030 FE0F",
			Group:       "symbols",
			SubGroup:    "punctuation",
		},

		12349: {
			Slug:        "part-alternation-mark",
			Character:   "〽️",
			UnicodeName: "part alternation mark",
			CodePoint:   "303D FE0F",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		12951: {
			Slug:        "japanese-congratulations-button",
			Character:   "㊗️",
			UnicodeName: "Japanese “congratulations” button",
			CodePoint:   "3297 FE0F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		12953: {
			Slug:        "japanese-secret-button",
			Character:   "㊙️",
			UnicodeName: "Japanese “secret” button",
			CodePoint:   "3299 FE0F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		126980: {
			Slug:        "mahjong-red-dragon",
			Character:   "🀄",
			UnicodeName: "mahjong red dragon",
			CodePoint:   "1F004",
			Group:       "activities",
			SubGroup:    "game",
		},

		127183: {
			Slug:        "joker",
			Character:   "🃏",
			UnicodeName: "joker",
			CodePoint:   "1F0CF",
			Group:       "activities",
			SubGroup:    "game",
		},

		127344: {
			Slug:        "a-button-blood-type",
			Character:   "🅰️",
			UnicodeName: "A button (blood type)",
			CodePoint:   "1F170 FE0F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127345: {
			Slug:        "b-button-blood-type",
			Character:   "🅱️",
			UnicodeName: "B button (blood type)",
			CodePoint:   "1F171 FE0F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127358: {
			Slug:        "o-button-blood-type",
			Character:   "🅾️",
			UnicodeName: "O button (blood type)",
			CodePoint:   "1F17E FE0F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127359: {
			Slug:        "p-button",
			Character:   "🅿️",
			UnicodeName: "P button",
			CodePoint:   "1F17F FE0F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127374: {
			Slug:        "ab-button-blood-type",
			Character:   "🆎",
			UnicodeName: "AB button (blood type)",
			CodePoint:   "1F18E",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127377: {
			Slug:        "cl-button",
			Character:   "🆑",
			UnicodeName: "CL button",
			CodePoint:   "1F191",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127378: {
			Slug:        "cool-button",
			Character:   "🆒",
			UnicodeName: "COOL button",
			CodePoint:   "1F192",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127379: {
			Slug:        "free-button",
			Character:   "🆓",
			UnicodeName: "FREE button",
			CodePoint:   "1F193",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127380: {
			Slug:        "id-button",
			Character:   "🆔",
			UnicodeName: "ID button",
			CodePoint:   "1F194",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127381: {
			Slug:        "new-button",
			Character:   "🆕",
			UnicodeName: "NEW button",
			CodePoint:   "1F195",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127382: {
			Slug:        "ng-button",
			Character:   "🆖",
			UnicodeName: "NG button",
			CodePoint:   "1F196",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127383: {
			Slug:        "ok-button",
			Character:   "🆗",
			UnicodeName: "OK button",
			CodePoint:   "1F197",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127384: {
			Slug:        "sos-button",
			Character:   "🆘",
			UnicodeName: "SOS button",
			CodePoint:   "1F198",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127385: {
			Slug:        "up-button",
			Character:   "🆙",
			UnicodeName: "UP! button",
			CodePoint:   "1F199",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127386: {
			Slug:        "vs-button",
			Character:   "🆚",
			UnicodeName: "VS button",
			CodePoint:   "1F19A",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127462: {
			Slug:        "flag-azerbaijan",
			Character:   "🇦🇿",
			UnicodeName: "flag: Azerbaijan",
			CodePoint:   "1F1E6 1F1FF",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127463: {
			Slug:        "flag-belize",
			Character:   "🇧🇿",
			UnicodeName: "flag: Belize",
			CodePoint:   "1F1E7 1F1FF",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127464: {
			Slug:        "flag-czechia",
			Character:   "🇨🇿",
			UnicodeName: "flag: Czechia",
			CodePoint:   "1F1E8 1F1FF",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127465: {
			Slug:        "flag-algeria",
			Character:   "🇩🇿",
			UnicodeName: "flag: Algeria",
			CodePoint:   "1F1E9 1F1FF",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127466: {
			Slug:        "flag-european-union",
			Character:   "🇪🇺",
			UnicodeName: "flag: European Union",
			CodePoint:   "1F1EA 1F1FA",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127467: {
			Slug:        "flag-france",
			Character:   "🇫🇷",
			UnicodeName: "flag: France",
			CodePoint:   "1F1EB 1F1F7",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127468: {
			Slug:        "flag-guyana",
			Character:   "🇬🇾",
			UnicodeName: "flag: Guyana",
			CodePoint:   "1F1EC 1F1FE",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127469: {
			Slug:        "flag-hungary",
			Character:   "🇭🇺",
			UnicodeName: "flag: Hungary",
			CodePoint:   "1F1ED 1F1FA",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127470: {
			Slug:        "flag-italy",
			Character:   "🇮🇹",
			UnicodeName: "flag: Italy",
			CodePoint:   "1F1EE 1F1F9",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127471: {
			Slug:        "flag-japan",
			Character:   "🇯🇵",
			UnicodeName: "flag: Japan",
			CodePoint:   "1F1EF 1F1F5",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127472: {
			Slug:        "flag-kazakhstan",
			Character:   "🇰🇿",
			UnicodeName: "flag: Kazakhstan",
			CodePoint:   "1F1F0 1F1FF",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127473: {
			Slug:        "flag-libya",
			Character:   "🇱🇾",
			UnicodeName: "flag: Libya",
			CodePoint:   "1F1F1 1F1FE",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127474: {
			Slug:        "flag-mozambique",
			Character:   "🇲🇿",
			UnicodeName: "flag: Mozambique",
			CodePoint:   "1F1F2 1F1FF",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127475: {
			Slug:        "flag-new-zealand",
			Character:   "🇳🇿",
			UnicodeName: "flag: New Zealand",
			CodePoint:   "1F1F3 1F1FF",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127476: {
			Slug:        "flag-oman",
			Character:   "🇴🇲",
			UnicodeName: "flag: Oman",
			CodePoint:   "1F1F4 1F1F2",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127477: {
			Slug:        "flag-paraguay",
			Character:   "🇵🇾",
			UnicodeName: "flag: Paraguay",
			CodePoint:   "1F1F5 1F1FE",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127478: {
			Slug:        "flag-qatar",
			Character:   "🇶🇦",
			UnicodeName: "flag: Qatar",
			CodePoint:   "1F1F6 1F1E6",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127479: {
			Slug:        "flag-rwanda",
			Character:   "🇷🇼",
			UnicodeName: "flag: Rwanda",
			CodePoint:   "1F1F7 1F1FC",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127480: {
			Slug:        "flag-eswatini",
			Character:   "🇸🇿",
			UnicodeName: "flag: Eswatini",
			CodePoint:   "1F1F8 1F1FF",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127481: {
			Slug:        "flag-tanzania",
			Character:   "🇹🇿",
			UnicodeName: "flag: Tanzania",
			CodePoint:   "1F1F9 1F1FF",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127482: {
			Slug:        "flag-uzbekistan",
			Character:   "🇺🇿",
			UnicodeName: "flag: Uzbekistan",
			CodePoint:   "1F1FA 1F1FF",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127483: {
			Slug:        "flag-vanuatu",
			Character:   "🇻🇺",
			UnicodeName: "flag: Vanuatu",
			CodePoint:   "1F1FB 1F1FA",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127484: {
			Slug:        "flag-samoa",
			Character:   "🇼🇸",
			UnicodeName: "flag: Samoa",
			CodePoint:   "1F1FC 1F1F8",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127485: {
			Slug:        "flag-kosovo",
			Character:   "🇽🇰",
			UnicodeName: "flag: Kosovo",
			CodePoint:   "1F1FD 1F1F0",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127486: {
			Slug:        "flag-mayotte",
			Character:   "🇾🇹",
			UnicodeName: "flag: Mayotte",
			CodePoint:   "1F1FE 1F1F9",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127487: {
			Slug:        "flag-zimbabwe",
			Character:   "🇿🇼",
			UnicodeName: "flag: Zimbabwe",
			CodePoint:   "1F1FF 1F1FC",
			Group:       "flags",
			SubGroup:    "country-flag",
		},

		127489: {
			Slug:        "japanese-here-button",
			Character:   "🈁",
			UnicodeName: "Japanese “here” button",
			CodePoint:   "1F201",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127490: {
			Slug:        "japanese-service-charge-button",
			Character:   "🈂️",
			UnicodeName: "Japanese “service charge” button",
			CodePoint:   "1F202 FE0F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127514: {
			Slug:        "japanese-free-of-charge-button",
			Character:   "🈚",
			UnicodeName: "Japanese “free of charge” button",
			CodePoint:   "1F21A",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127535: {
			Slug:        "japanese-reserved-button",
			Character:   "🈯",
			UnicodeName: "Japanese “reserved” button",
			CodePoint:   "1F22F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127538: {
			Slug:        "japanese-prohibited-button",
			Character:   "🈲",
			UnicodeName: "Japanese “prohibited” button",
			CodePoint:   "1F232",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127539: {
			Slug:        "japanese-vacancy-button",
			Character:   "🈳",
			UnicodeName: "Japanese “vacancy” button",
			CodePoint:   "1F233",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127540: {
			Slug:        "japanese-passing-grade-button",
			Character:   "🈴",
			UnicodeName: "Japanese “passing grade” button",
			CodePoint:   "1F234",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127541: {
			Slug:        "japanese-no-vacancy-button",
			Character:   "🈵",
			UnicodeName: "Japanese “no vacancy” button",
			CodePoint:   "1F235",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127542: {
			Slug:        "japanese-not-free-of-charge-button",
			Character:   "🈶",
			UnicodeName: "Japanese “not free of charge” button",
			CodePoint:   "1F236",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127543: {
			Slug:        "japanese-monthly-amount-button",
			Character:   "🈷️",
			UnicodeName: "Japanese “monthly amount” button",
			CodePoint:   "1F237 FE0F",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127544: {
			Slug:        "japanese-application-button",
			Character:   "🈸",
			UnicodeName: "Japanese “application” button",
			CodePoint:   "1F238",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127545: {
			Slug:        "japanese-discount-button",
			Character:   "🈹",
			UnicodeName: "Japanese “discount” button",
			CodePoint:   "1F239",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127546: {
			Slug:        "japanese-open-for-business-button",
			Character:   "🈺",
			UnicodeName: "Japanese “open for business” button",
			CodePoint:   "1F23A",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127568: {
			Slug:        "japanese-bargain-button",
			Character:   "🉐",
			UnicodeName: "Japanese “bargain” button",
			CodePoint:   "1F250",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127569: {
			Slug:        "japanese-acceptable-button",
			Character:   "🉑",
			UnicodeName: "Japanese “acceptable” button",
			CodePoint:   "1F251",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		127744: {
			Slug:        "cyclone",
			Character:   "🌀",
			UnicodeName: "cyclone",
			CodePoint:   "1F300",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127745: {
			Slug:        "e0-6-foggy",
			Character:   "🌁",
			UnicodeName: "E0.6 foggy",
			CodePoint:   "1F301",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127746: {
			Slug:        "closed-umbrella",
			Character:   "🌂",
			UnicodeName: "closed umbrella",
			CodePoint:   "1F302",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127747: {
			Slug:        "e0-6-night-with-stars",
			Character:   "🌃",
			UnicodeName: "E0.6 night with stars",
			CodePoint:   "1F303",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127748: {
			Slug:        "e0-6-sunrise-over-mountains",
			Character:   "🌄",
			UnicodeName: "E0.6 sunrise over mountains",
			CodePoint:   "1F304",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127749: {
			Slug:        "e0-6-sunrise",
			Character:   "🌅",
			UnicodeName: "E0.6 sunrise",
			CodePoint:   "1F305",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127750: {
			Slug:        "e0-6-cityscape-at-dusk",
			Character:   "🌆",
			UnicodeName: "E0.6 cityscape at dusk",
			CodePoint:   "1F306",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127751: {
			Slug:        "e0-6-sunset",
			Character:   "🌇",
			UnicodeName: "E0.6 sunset",
			CodePoint:   "1F307",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127752: {
			Slug:        "rainbow",
			Character:   "🌈",
			UnicodeName: "rainbow",
			CodePoint:   "1F308",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127753: {
			Slug:        "e0-6-bridge-at-night",
			Character:   "🌉",
			UnicodeName: "E0.6 bridge at night",
			CodePoint:   "1F309",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127754: {
			Slug:        "water-wave",
			Character:   "🌊",
			UnicodeName: "water wave",
			CodePoint:   "1F30A",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127755: {
			Slug:        "e0-6-volcano",
			Character:   "🌋",
			UnicodeName: "E0.6 volcano",
			CodePoint:   "1F30B",
			Group:       "travel-places",
			SubGroup:    "place-geographic",
		},

		127756: {
			Slug:        "milky-way",
			Character:   "🌌",
			UnicodeName: "milky way",
			CodePoint:   "1F30C",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127757: {
			Slug:        "e0-7-globe-showing-europe-africa",
			Character:   "🌍",
			UnicodeName: "E0.7 globe showing Europe-Africa",
			CodePoint:   "1F30D",
			Group:       "travel-places",
			SubGroup:    "place-map",
		},

		127758: {
			Slug:        "e0-7-globe-showing-americas",
			Character:   "🌎",
			UnicodeName: "E0.7 globe showing Americas",
			CodePoint:   "1F30E",
			Group:       "travel-places",
			SubGroup:    "place-map",
		},

		127759: {
			Slug:        "e0-6-globe-showing-asia-australia",
			Character:   "🌏",
			UnicodeName: "E0.6 globe showing Asia-Australia",
			CodePoint:   "1F30F",
			Group:       "travel-places",
			SubGroup:    "place-map",
		},

		127760: {
			Slug:        "e1-0-globe-with-meridians",
			Character:   "🌐",
			UnicodeName: "E1.0 globe with meridians",
			CodePoint:   "1F310",
			Group:       "travel-places",
			SubGroup:    "place-map",
		},

		127761: {
			Slug:        "new-moon",
			Character:   "🌑",
			UnicodeName: "new moon",
			CodePoint:   "1F311",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127762: {
			Slug:        "waxing-crescent-moon",
			Character:   "🌒",
			UnicodeName: "waxing crescent moon",
			CodePoint:   "1F312",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127763: {
			Slug:        "first-quarter-moon",
			Character:   "🌓",
			UnicodeName: "first quarter moon",
			CodePoint:   "1F313",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127764: {
			Slug:        "waxing-gibbous-moon",
			Character:   "🌔",
			UnicodeName: "waxing gibbous moon",
			CodePoint:   "1F314",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127765: {
			Slug:        "full-moon",
			Character:   "🌕",
			UnicodeName: "full moon",
			CodePoint:   "1F315",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127766: {
			Slug:        "waning-gibbous-moon",
			Character:   "🌖",
			UnicodeName: "waning gibbous moon",
			CodePoint:   "1F316",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127767: {
			Slug:        "last-quarter-moon",
			Character:   "🌗",
			UnicodeName: "last quarter moon",
			CodePoint:   "1F317",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127768: {
			Slug:        "waning-crescent-moon",
			Character:   "🌘",
			UnicodeName: "waning crescent moon",
			CodePoint:   "1F318",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127769: {
			Slug:        "crescent-moon",
			Character:   "🌙",
			UnicodeName: "crescent moon",
			CodePoint:   "1F319",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127770: {
			Slug:        "new-moon-face",
			Character:   "🌚",
			UnicodeName: "new moon face",
			CodePoint:   "1F31A",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127771: {
			Slug:        "first-quarter-moon-face",
			Character:   "🌛",
			UnicodeName: "first quarter moon face",
			CodePoint:   "1F31B",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127772: {
			Slug:        "last-quarter-moon-face",
			Character:   "🌜",
			UnicodeName: "last quarter moon face",
			CodePoint:   "1F31C",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127773: {
			Slug:        "full-moon-face",
			Character:   "🌝",
			UnicodeName: "full moon face",
			CodePoint:   "1F31D",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127774: {
			Slug:        "sun-with-face",
			Character:   "🌞",
			UnicodeName: "sun with face",
			CodePoint:   "1F31E",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127775: {
			Slug:        "glowing-star",
			Character:   "🌟",
			UnicodeName: "glowing star",
			CodePoint:   "1F31F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127776: {
			Slug:        "shooting-star",
			Character:   "🌠",
			UnicodeName: "shooting star",
			CodePoint:   "1F320",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127777: {
			Slug:        "thermometer",
			Character:   "🌡️",
			UnicodeName: "thermometer",
			CodePoint:   "1F321 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127780: {
			Slug:        "sun-behind-small-cloud",
			Character:   "🌤️",
			UnicodeName: "sun behind small cloud",
			CodePoint:   "1F324 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127781: {
			Slug:        "sun-behind-large-cloud",
			Character:   "🌥️",
			UnicodeName: "sun behind large cloud",
			CodePoint:   "1F325 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127782: {
			Slug:        "sun-behind-rain-cloud",
			Character:   "🌦️",
			UnicodeName: "sun behind rain cloud",
			CodePoint:   "1F326 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127783: {
			Slug:        "cloud-with-rain",
			Character:   "🌧️",
			UnicodeName: "cloud with rain",
			CodePoint:   "1F327 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127784: {
			Slug:        "cloud-with-snow",
			Character:   "🌨️",
			UnicodeName: "cloud with snow",
			CodePoint:   "1F328 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127785: {
			Slug:        "cloud-with-lightning",
			Character:   "🌩️",
			UnicodeName: "cloud with lightning",
			CodePoint:   "1F329 FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127786: {
			Slug:        "tornado",
			Character:   "🌪️",
			UnicodeName: "tornado",
			CodePoint:   "1F32A FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127787: {
			Slug:        "fog",
			Character:   "🌫️",
			UnicodeName: "fog",
			CodePoint:   "1F32B FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127788: {
			Slug:        "wind-face",
			Character:   "🌬️",
			UnicodeName: "wind face",
			CodePoint:   "1F32C FE0F",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		127789: {
			Slug:        "e1-0-hot-dog",
			Character:   "🌭",
			UnicodeName: "E1.0 hot dog",
			CodePoint:   "1F32D",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127790: {
			Slug:        "e1-0-taco",
			Character:   "🌮",
			UnicodeName: "E1.0 taco",
			CodePoint:   "1F32E",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127791: {
			Slug:        "e1-0-burrito",
			Character:   "🌯",
			UnicodeName: "E1.0 burrito",
			CodePoint:   "1F32F",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127792: {
			Slug:        "e0-6-chestnut",
			Character:   "🌰",
			UnicodeName: "E0.6 chestnut",
			CodePoint:   "1F330",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		127793: {
			Slug:        "e0-6-seedling",
			Character:   "🌱",
			UnicodeName: "E0.6 seedling",
			CodePoint:   "1F331",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127794: {
			Slug:        "e1-0-evergreen-tree",
			Character:   "🌲",
			UnicodeName: "E1.0 evergreen tree",
			CodePoint:   "1F332",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127795: {
			Slug:        "e1-0-deciduous-tree",
			Character:   "🌳",
			UnicodeName: "E1.0 deciduous tree",
			CodePoint:   "1F333",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127796: {
			Slug:        "e0-6-palm-tree",
			Character:   "🌴",
			UnicodeName: "E0.6 palm tree",
			CodePoint:   "1F334",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127797: {
			Slug:        "e0-6-cactus",
			Character:   "🌵",
			UnicodeName: "E0.6 cactus",
			CodePoint:   "1F335",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127798: {
			Slug:        "e0-7-hot-pepper",
			Character:   "🌶️",
			UnicodeName: "E0.7 hot pepper",
			CodePoint:   "1F336 FE0F",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		127799: {
			Slug:        "e0-6-tulip",
			Character:   "🌷",
			UnicodeName: "E0.6 tulip",
			CodePoint:   "1F337",
			Group:       "animals-nature",
			SubGroup:    "plant-flower",
		},

		127800: {
			Slug:        "e0-6-cherry-blossom",
			Character:   "🌸",
			UnicodeName: "E0.6 cherry blossom",
			CodePoint:   "1F338",
			Group:       "animals-nature",
			SubGroup:    "plant-flower",
		},

		127801: {
			Slug:        "e0-6-rose",
			Character:   "🌹",
			UnicodeName: "E0.6 rose",
			CodePoint:   "1F339",
			Group:       "animals-nature",
			SubGroup:    "plant-flower",
		},

		127802: {
			Slug:        "e0-6-hibiscus",
			Character:   "🌺",
			UnicodeName: "E0.6 hibiscus",
			CodePoint:   "1F33A",
			Group:       "animals-nature",
			SubGroup:    "plant-flower",
		},

		127803: {
			Slug:        "e0-6-sunflower",
			Character:   "🌻",
			UnicodeName: "E0.6 sunflower",
			CodePoint:   "1F33B",
			Group:       "animals-nature",
			SubGroup:    "plant-flower",
		},

		127804: {
			Slug:        "e0-6-blossom",
			Character:   "🌼",
			UnicodeName: "E0.6 blossom",
			CodePoint:   "1F33C",
			Group:       "animals-nature",
			SubGroup:    "plant-flower",
		},

		127805: {
			Slug:        "e0-6-ear-of-corn",
			Character:   "🌽",
			UnicodeName: "E0.6 ear of corn",
			CodePoint:   "1F33D",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		127806: {
			Slug:        "e0-6-sheaf-of-rice",
			Character:   "🌾",
			UnicodeName: "E0.6 sheaf of rice",
			CodePoint:   "1F33E",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127807: {
			Slug:        "e0-6-herb",
			Character:   "🌿",
			UnicodeName: "E0.6 herb",
			CodePoint:   "1F33F",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127808: {
			Slug:        "e0-6-four-leaf-clover",
			Character:   "🍀",
			UnicodeName: "E0.6 four leaf clover",
			CodePoint:   "1F340",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127809: {
			Slug:        "e0-6-maple-leaf",
			Character:   "🍁",
			UnicodeName: "E0.6 maple leaf",
			CodePoint:   "1F341",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127810: {
			Slug:        "e0-6-fallen-leaf",
			Character:   "🍂",
			UnicodeName: "E0.6 fallen leaf",
			CodePoint:   "1F342",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127811: {
			Slug:        "e0-6-leaf-fluttering-in-wind",
			Character:   "🍃",
			UnicodeName: "E0.6 leaf fluttering in wind",
			CodePoint:   "1F343",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		127812: {
			Slug:        "e0-6-mushroom",
			Character:   "🍄",
			UnicodeName: "E0.6 mushroom",
			CodePoint:   "1F344",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		127813: {
			Slug:        "e0-6-tomato",
			Character:   "🍅",
			UnicodeName: "E0.6 tomato",
			CodePoint:   "1F345",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127814: {
			Slug:        "e0-6-eggplant",
			Character:   "🍆",
			UnicodeName: "E0.6 eggplant",
			CodePoint:   "1F346",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		127815: {
			Slug:        "e0-6-grapes",
			Character:   "🍇",
			UnicodeName: "E0.6 grapes",
			CodePoint:   "1F347",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127816: {
			Slug:        "e0-6-melon",
			Character:   "🍈",
			UnicodeName: "E0.6 melon",
			CodePoint:   "1F348",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127817: {
			Slug:        "e0-6-watermelon",
			Character:   "🍉",
			UnicodeName: "E0.6 watermelon",
			CodePoint:   "1F349",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127818: {
			Slug:        "e0-6-tangerine",
			Character:   "🍊",
			UnicodeName: "E0.6 tangerine",
			CodePoint:   "1F34A",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127819: {
			Slug:        "e1-0-lemon",
			Character:   "🍋",
			UnicodeName: "E1.0 lemon",
			CodePoint:   "1F34B",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127820: {
			Slug:        "e0-6-banana",
			Character:   "🍌",
			UnicodeName: "E0.6 banana",
			CodePoint:   "1F34C",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127821: {
			Slug:        "e0-6-pineapple",
			Character:   "🍍",
			UnicodeName: "E0.6 pineapple",
			CodePoint:   "1F34D",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127822: {
			Slug:        "e0-6-red-apple",
			Character:   "🍎",
			UnicodeName: "E0.6 red apple",
			CodePoint:   "1F34E",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127823: {
			Slug:        "e0-6-green-apple",
			Character:   "🍏",
			UnicodeName: "E0.6 green apple",
			CodePoint:   "1F34F",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127824: {
			Slug:        "e1-0-pear",
			Character:   "🍐",
			UnicodeName: "E1.0 pear",
			CodePoint:   "1F350",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127825: {
			Slug:        "e0-6-peach",
			Character:   "🍑",
			UnicodeName: "E0.6 peach",
			CodePoint:   "1F351",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127826: {
			Slug:        "e0-6-cherries",
			Character:   "🍒",
			UnicodeName: "E0.6 cherries",
			CodePoint:   "1F352",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127827: {
			Slug:        "e0-6-strawberry",
			Character:   "🍓",
			UnicodeName: "E0.6 strawberry",
			CodePoint:   "1F353",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		127828: {
			Slug:        "e0-6-hamburger",
			Character:   "🍔",
			UnicodeName: "E0.6 hamburger",
			CodePoint:   "1F354",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127829: {
			Slug:        "e0-6-pizza",
			Character:   "🍕",
			UnicodeName: "E0.6 pizza",
			CodePoint:   "1F355",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127830: {
			Slug:        "e0-6-meat-on-bone",
			Character:   "🍖",
			UnicodeName: "E0.6 meat on bone",
			CodePoint:   "1F356",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127831: {
			Slug:        "e0-6-poultry-leg",
			Character:   "🍗",
			UnicodeName: "E0.6 poultry leg",
			CodePoint:   "1F357",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127832: {
			Slug:        "e0-6-rice-cracker",
			Character:   "🍘",
			UnicodeName: "E0.6 rice cracker",
			CodePoint:   "1F358",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127833: {
			Slug:        "e0-6-rice-ball",
			Character:   "🍙",
			UnicodeName: "E0.6 rice ball",
			CodePoint:   "1F359",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127834: {
			Slug:        "e0-6-cooked-rice",
			Character:   "🍚",
			UnicodeName: "E0.6 cooked rice",
			CodePoint:   "1F35A",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127835: {
			Slug:        "e0-6-curry-rice",
			Character:   "🍛",
			UnicodeName: "E0.6 curry rice",
			CodePoint:   "1F35B",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127836: {
			Slug:        "e0-6-steaming-bowl",
			Character:   "🍜",
			UnicodeName: "E0.6 steaming bowl",
			CodePoint:   "1F35C",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127837: {
			Slug:        "e0-6-spaghetti",
			Character:   "🍝",
			UnicodeName: "E0.6 spaghetti",
			CodePoint:   "1F35D",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127838: {
			Slug:        "e0-6-bread",
			Character:   "🍞",
			UnicodeName: "E0.6 bread",
			CodePoint:   "1F35E",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127839: {
			Slug:        "e0-6-french-fries",
			Character:   "🍟",
			UnicodeName: "E0.6 french fries",
			CodePoint:   "1F35F",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127840: {
			Slug:        "e0-6-roasted-sweet-potato",
			Character:   "🍠",
			UnicodeName: "E0.6 roasted sweet potato",
			CodePoint:   "1F360",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127841: {
			Slug:        "e0-6-dango",
			Character:   "🍡",
			UnicodeName: "E0.6 dango",
			CodePoint:   "1F361",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127842: {
			Slug:        "e0-6-oden",
			Character:   "🍢",
			UnicodeName: "E0.6 oden",
			CodePoint:   "1F362",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127843: {
			Slug:        "e0-6-sushi",
			Character:   "🍣",
			UnicodeName: "E0.6 sushi",
			CodePoint:   "1F363",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127844: {
			Slug:        "e0-6-fried-shrimp",
			Character:   "🍤",
			UnicodeName: "E0.6 fried shrimp",
			CodePoint:   "1F364",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127845: {
			Slug:        "e0-6-fish-cake-with-swirl",
			Character:   "🍥",
			UnicodeName: "E0.6 fish cake with swirl",
			CodePoint:   "1F365",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127846: {
			Slug:        "e0-6-soft-ice-cream",
			Character:   "🍦",
			UnicodeName: "E0.6 soft ice cream",
			CodePoint:   "1F366",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127847: {
			Slug:        "e0-6-shaved-ice",
			Character:   "🍧",
			UnicodeName: "E0.6 shaved ice",
			CodePoint:   "1F367",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127848: {
			Slug:        "e0-6-ice-cream",
			Character:   "🍨",
			UnicodeName: "E0.6 ice cream",
			CodePoint:   "1F368",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127849: {
			Slug:        "e0-6-doughnut",
			Character:   "🍩",
			UnicodeName: "E0.6 doughnut",
			CodePoint:   "1F369",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127850: {
			Slug:        "e0-6-cookie",
			Character:   "🍪",
			UnicodeName: "E0.6 cookie",
			CodePoint:   "1F36A",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127851: {
			Slug:        "e0-6-chocolate-bar",
			Character:   "🍫",
			UnicodeName: "E0.6 chocolate bar",
			CodePoint:   "1F36B",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127852: {
			Slug:        "e0-6-candy",
			Character:   "🍬",
			UnicodeName: "E0.6 candy",
			CodePoint:   "1F36C",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127853: {
			Slug:        "e0-6-lollipop",
			Character:   "🍭",
			UnicodeName: "E0.6 lollipop",
			CodePoint:   "1F36D",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127854: {
			Slug:        "e0-6-custard",
			Character:   "🍮",
			UnicodeName: "E0.6 custard",
			CodePoint:   "1F36E",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127855: {
			Slug:        "e0-6-honey-pot",
			Character:   "🍯",
			UnicodeName: "E0.6 honey pot",
			CodePoint:   "1F36F",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127856: {
			Slug:        "e0-6-shortcake",
			Character:   "🍰",
			UnicodeName: "E0.6 shortcake",
			CodePoint:   "1F370",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127857: {
			Slug:        "e0-6-bento-box",
			Character:   "🍱",
			UnicodeName: "E0.6 bento box",
			CodePoint:   "1F371",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		127858: {
			Slug:        "e0-6-pot-of-food",
			Character:   "🍲",
			UnicodeName: "E0.6 pot of food",
			CodePoint:   "1F372",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127859: {
			Slug:        "e0-6-cooking",
			Character:   "🍳",
			UnicodeName: "E0.6 cooking",
			CodePoint:   "1F373",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127860: {
			Slug:        "e0-6-fork-and-knife",
			Character:   "🍴",
			UnicodeName: "E0.6 fork and knife",
			CodePoint:   "1F374",
			Group:       "food-drink",
			SubGroup:    "dishware",
		},

		127861: {
			Slug:        "e0-6-teacup-without-handle",
			Character:   "🍵",
			UnicodeName: "E0.6 teacup without handle",
			CodePoint:   "1F375",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		127862: {
			Slug:        "e0-6-sake",
			Character:   "🍶",
			UnicodeName: "E0.6 sake",
			CodePoint:   "1F376",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		127863: {
			Slug:        "e0-6-wine-glass",
			Character:   "🍷",
			UnicodeName: "E0.6 wine glass",
			CodePoint:   "1F377",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		127864: {
			Slug:        "e0-6-cocktail-glass",
			Character:   "🍸",
			UnicodeName: "E0.6 cocktail glass",
			CodePoint:   "1F378",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		127865: {
			Slug:        "e0-6-tropical-drink",
			Character:   "🍹",
			UnicodeName: "E0.6 tropical drink",
			CodePoint:   "1F379",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		127866: {
			Slug:        "e0-6-beer-mug",
			Character:   "🍺",
			UnicodeName: "E0.6 beer mug",
			CodePoint:   "1F37A",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		127867: {
			Slug:        "e0-6-clinking-beer-mugs",
			Character:   "🍻",
			UnicodeName: "E0.6 clinking beer mugs",
			CodePoint:   "1F37B",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		127868: {
			Slug:        "e1-0-baby-bottle",
			Character:   "🍼",
			UnicodeName: "E1.0 baby bottle",
			CodePoint:   "1F37C",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		127869: {
			Slug:        "e0-7-fork-and-knife-with-plate",
			Character:   "🍽️",
			UnicodeName: "E0.7 fork and knife with plate",
			CodePoint:   "1F37D FE0F",
			Group:       "food-drink",
			SubGroup:    "dishware",
		},

		127870: {
			Slug:        "e1-0-bottle-with-popping-cork",
			Character:   "🍾",
			UnicodeName: "E1.0 bottle with popping cork",
			CodePoint:   "1F37E",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		127871: {
			Slug:        "e1-0-popcorn",
			Character:   "🍿",
			UnicodeName: "E1.0 popcorn",
			CodePoint:   "1F37F",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		127872: {
			Slug:        "ribbon",
			Character:   "🎀",
			UnicodeName: "ribbon",
			CodePoint:   "1F380",
			Group:       "activities",
			SubGroup:    "event",
		},

		127873: {
			Slug:        "wrapped-gift",
			Character:   "🎁",
			UnicodeName: "wrapped gift",
			CodePoint:   "1F381",
			Group:       "activities",
			SubGroup:    "event",
		},

		127874: {
			Slug:        "e0-6-birthday-cake",
			Character:   "🎂",
			UnicodeName: "E0.6 birthday cake",
			CodePoint:   "1F382",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		127875: {
			Slug:        "jack-o-lantern",
			Character:   "🎃",
			UnicodeName: "jack-o-lantern",
			CodePoint:   "1F383",
			Group:       "activities",
			SubGroup:    "event",
		},

		127876: {
			Slug:        "christmas-tree",
			Character:   "🎄",
			UnicodeName: "Christmas tree",
			CodePoint:   "1F384",
			Group:       "activities",
			SubGroup:    "event",
		},

		127877: {
			Slug:        "e0-6-santa-claus",
			Character:   "🎅",
			UnicodeName: "E0.6 Santa Claus",
			CodePoint:   "1F385",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		127878: {
			Slug:        "fireworks",
			Character:   "🎆",
			UnicodeName: "fireworks",
			CodePoint:   "1F386",
			Group:       "activities",
			SubGroup:    "event",
		},

		127879: {
			Slug:        "sparkler",
			Character:   "🎇",
			UnicodeName: "sparkler",
			CodePoint:   "1F387",
			Group:       "activities",
			SubGroup:    "event",
		},

		127880: {
			Slug:        "balloon",
			Character:   "🎈",
			UnicodeName: "balloon",
			CodePoint:   "1F388",
			Group:       "activities",
			SubGroup:    "event",
		},

		127881: {
			Slug:        "party-popper",
			Character:   "🎉",
			UnicodeName: "party popper",
			CodePoint:   "1F389",
			Group:       "activities",
			SubGroup:    "event",
		},

		127882: {
			Slug:        "confetti-ball",
			Character:   "🎊",
			UnicodeName: "confetti ball",
			CodePoint:   "1F38A",
			Group:       "activities",
			SubGroup:    "event",
		},

		127883: {
			Slug:        "tanabata-tree",
			Character:   "🎋",
			UnicodeName: "tanabata tree",
			CodePoint:   "1F38B",
			Group:       "activities",
			SubGroup:    "event",
		},

		127884: {
			Slug:        "crossed-flags",
			Character:   "🎌",
			UnicodeName: "crossed flags",
			CodePoint:   "1F38C",
			Group:       "flags",
			SubGroup:    "flag",
		},

		127885: {
			Slug:        "pine-decoration",
			Character:   "🎍",
			UnicodeName: "pine decoration",
			CodePoint:   "1F38D",
			Group:       "activities",
			SubGroup:    "event",
		},

		127886: {
			Slug:        "japanese-dolls",
			Character:   "🎎",
			UnicodeName: "Japanese dolls",
			CodePoint:   "1F38E",
			Group:       "activities",
			SubGroup:    "event",
		},

		127887: {
			Slug:        "carp-streamer",
			Character:   "🎏",
			UnicodeName: "carp streamer",
			CodePoint:   "1F38F",
			Group:       "activities",
			SubGroup:    "event",
		},

		127888: {
			Slug:        "wind-chime",
			Character:   "🎐",
			UnicodeName: "wind chime",
			CodePoint:   "1F390",
			Group:       "activities",
			SubGroup:    "event",
		},

		127889: {
			Slug:        "moon-viewing-ceremony",
			Character:   "🎑",
			UnicodeName: "moon viewing ceremony",
			CodePoint:   "1F391",
			Group:       "activities",
			SubGroup:    "event",
		},

		127890: {
			Slug:        "backpack",
			Character:   "🎒",
			UnicodeName: "backpack",
			CodePoint:   "1F392",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		127891: {
			Slug:        "graduation-cap",
			Character:   "🎓",
			UnicodeName: "graduation cap",
			CodePoint:   "1F393",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		127894: {
			Slug:        "military-medal",
			Character:   "🎖️",
			UnicodeName: "military medal",
			CodePoint:   "1F396 FE0F",
			Group:       "activities",
			SubGroup:    "award-medal",
		},

		127895: {
			Slug:        "reminder-ribbon",
			Character:   "🎗️",
			UnicodeName: "reminder ribbon",
			CodePoint:   "1F397 FE0F",
			Group:       "activities",
			SubGroup:    "event",
		},

		127897: {
			Slug:        "studio-microphone",
			Character:   "🎙️",
			UnicodeName: "studio microphone",
			CodePoint:   "1F399 FE0F",
			Group:       "objects",
			SubGroup:    "music",
		},

		127898: {
			Slug:        "level-slider",
			Character:   "🎚️",
			UnicodeName: "level slider",
			CodePoint:   "1F39A FE0F",
			Group:       "objects",
			SubGroup:    "music",
		},

		127899: {
			Slug:        "control-knobs",
			Character:   "🎛️",
			UnicodeName: "control knobs",
			CodePoint:   "1F39B FE0F",
			Group:       "objects",
			SubGroup:    "music",
		},

		127902: {
			Slug:        "film-frames",
			Character:   "🎞️",
			UnicodeName: "film frames",
			CodePoint:   "1F39E FE0F",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		127903: {
			Slug:        "admission-tickets",
			Character:   "🎟️",
			UnicodeName: "admission tickets",
			CodePoint:   "1F39F FE0F",
			Group:       "activities",
			SubGroup:    "event",
		},

		127904: {
			Slug:        "e0-6-carousel-horse",
			Character:   "🎠",
			UnicodeName: "E0.6 carousel horse",
			CodePoint:   "1F3A0",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127905: {
			Slug:        "e0-6-ferris-wheel",
			Character:   "🎡",
			UnicodeName: "E0.6 ferris wheel",
			CodePoint:   "1F3A1",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127906: {
			Slug:        "e0-6-roller-coaster",
			Character:   "🎢",
			UnicodeName: "E0.6 roller coaster",
			CodePoint:   "1F3A2",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127907: {
			Slug:        "fishing-pole",
			Character:   "🎣",
			UnicodeName: "fishing pole",
			CodePoint:   "1F3A3",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127908: {
			Slug:        "microphone",
			Character:   "🎤",
			UnicodeName: "microphone",
			CodePoint:   "1F3A4",
			Group:       "objects",
			SubGroup:    "music",
		},

		127909: {
			Slug:        "movie-camera",
			Character:   "🎥",
			UnicodeName: "movie camera",
			CodePoint:   "1F3A5",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		127910: {
			Slug:        "cinema",
			Character:   "🎦",
			UnicodeName: "cinema",
			CodePoint:   "1F3A6",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		127911: {
			Slug:        "headphone",
			Character:   "🎧",
			UnicodeName: "headphone",
			CodePoint:   "1F3A7",
			Group:       "objects",
			SubGroup:    "music",
		},

		127912: {
			Slug:        "artist-palette",
			Character:   "🎨",
			UnicodeName: "artist palette",
			CodePoint:   "1F3A8",
			Group:       "activities",
			SubGroup:    "arts-crafts",
		},

		127913: {
			Slug:        "top-hat",
			Character:   "🎩",
			UnicodeName: "top hat",
			CodePoint:   "1F3A9",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		127914: {
			Slug:        "e0-6-circus-tent",
			Character:   "🎪",
			UnicodeName: "E0.6 circus tent",
			CodePoint:   "1F3AA",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127915: {
			Slug:        "ticket",
			Character:   "🎫",
			UnicodeName: "ticket",
			CodePoint:   "1F3AB",
			Group:       "activities",
			SubGroup:    "event",
		},

		127916: {
			Slug:        "clapper-board",
			Character:   "🎬",
			UnicodeName: "clapper board",
			CodePoint:   "1F3AC",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		127917: {
			Slug:        "performing-arts",
			Character:   "🎭",
			UnicodeName: "performing arts",
			CodePoint:   "1F3AD",
			Group:       "activities",
			SubGroup:    "arts-crafts",
		},

		127918: {
			Slug:        "video-game",
			Character:   "🎮",
			UnicodeName: "video game",
			CodePoint:   "1F3AE",
			Group:       "activities",
			SubGroup:    "game",
		},

		127919: {
			Slug:        "direct-hit",
			Character:   "🎯",
			UnicodeName: "direct hit",
			CodePoint:   "1F3AF",
			Group:       "activities",
			SubGroup:    "game",
		},

		127920: {
			Slug:        "slot-machine",
			Character:   "🎰",
			UnicodeName: "slot machine",
			CodePoint:   "1F3B0",
			Group:       "activities",
			SubGroup:    "game",
		},

		127921: {
			Slug:        "pool-8-ball",
			Character:   "🎱",
			UnicodeName: "pool 8 ball",
			CodePoint:   "1F3B1",
			Group:       "activities",
			SubGroup:    "game",
		},

		127922: {
			Slug:        "game-die",
			Character:   "🎲",
			UnicodeName: "game die",
			CodePoint:   "1F3B2",
			Group:       "activities",
			SubGroup:    "game",
		},

		127923: {
			Slug:        "bowling",
			Character:   "🎳",
			UnicodeName: "bowling",
			CodePoint:   "1F3B3",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127924: {
			Slug:        "flower-playing-cards",
			Character:   "🎴",
			UnicodeName: "flower playing cards",
			CodePoint:   "1F3B4",
			Group:       "activities",
			SubGroup:    "game",
		},

		127925: {
			Slug:        "musical-note",
			Character:   "🎵",
			UnicodeName: "musical note",
			CodePoint:   "1F3B5",
			Group:       "objects",
			SubGroup:    "music",
		},

		127926: {
			Slug:        "musical-notes",
			Character:   "🎶",
			UnicodeName: "musical notes",
			CodePoint:   "1F3B6",
			Group:       "objects",
			SubGroup:    "music",
		},

		127927: {
			Slug:        "saxophone",
			Character:   "🎷",
			UnicodeName: "saxophone",
			CodePoint:   "1F3B7",
			Group:       "objects",
			SubGroup:    "musical-instrument",
		},

		127928: {
			Slug:        "guitar",
			Character:   "🎸",
			UnicodeName: "guitar",
			CodePoint:   "1F3B8",
			Group:       "objects",
			SubGroup:    "musical-instrument",
		},

		127929: {
			Slug:        "musical-keyboard",
			Character:   "🎹",
			UnicodeName: "musical keyboard",
			CodePoint:   "1F3B9",
			Group:       "objects",
			SubGroup:    "musical-instrument",
		},

		127930: {
			Slug:        "trumpet",
			Character:   "🎺",
			UnicodeName: "trumpet",
			CodePoint:   "1F3BA",
			Group:       "objects",
			SubGroup:    "musical-instrument",
		},

		127931: {
			Slug:        "violin",
			Character:   "🎻",
			UnicodeName: "violin",
			CodePoint:   "1F3BB",
			Group:       "objects",
			SubGroup:    "musical-instrument",
		},

		127932: {
			Slug:        "musical-score",
			Character:   "🎼",
			UnicodeName: "musical score",
			CodePoint:   "1F3BC",
			Group:       "objects",
			SubGroup:    "music",
		},

		127933: {
			Slug:        "running-shirt",
			Character:   "🎽",
			UnicodeName: "running shirt",
			CodePoint:   "1F3BD",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127934: {
			Slug:        "tennis",
			Character:   "🎾",
			UnicodeName: "tennis",
			CodePoint:   "1F3BE",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127935: {
			Slug:        "skis",
			Character:   "🎿",
			UnicodeName: "skis",
			CodePoint:   "1F3BF",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127936: {
			Slug:        "basketball",
			Character:   "🏀",
			UnicodeName: "basketball",
			CodePoint:   "1F3C0",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127937: {
			Slug:        "chequered-flag",
			Character:   "🏁",
			UnicodeName: "chequered flag",
			CodePoint:   "1F3C1",
			Group:       "flags",
			SubGroup:    "flag",
		},

		127938: {
			Slug:        "e0-6-snowboarder",
			Character:   "🏂",
			UnicodeName: "E0.6 snowboarder",
			CodePoint:   "1F3C2",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		127939: {
			Slug:        "e0-6-person-running",
			Character:   "🏃",
			UnicodeName: "E0.6 person running",
			CodePoint:   "1F3C3",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		127940: {
			Slug:        "e0-6-person-surfing",
			Character:   "🏄",
			UnicodeName: "E0.6 person surfing",
			CodePoint:   "1F3C4",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		127941: {
			Slug:        "sports-medal",
			Character:   "🏅",
			UnicodeName: "sports medal",
			CodePoint:   "1F3C5",
			Group:       "activities",
			SubGroup:    "award-medal",
		},

		127942: {
			Slug:        "trophy",
			Character:   "🏆",
			UnicodeName: "trophy",
			CodePoint:   "1F3C6",
			Group:       "activities",
			SubGroup:    "award-medal",
		},

		127943: {
			Slug:        "e1-0-horse-racing",
			Character:   "🏇",
			UnicodeName: "E1.0 horse racing",
			CodePoint:   "1F3C7",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		127944: {
			Slug:        "american-football",
			Character:   "🏈",
			UnicodeName: "american football",
			CodePoint:   "1F3C8",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127945: {
			Slug:        "rugby-football",
			Character:   "🏉",
			UnicodeName: "rugby football",
			CodePoint:   "1F3C9",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127946: {
			Slug:        "e0-6-person-swimming",
			Character:   "🏊",
			UnicodeName: "E0.6 person swimming",
			CodePoint:   "1F3CA",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		127947: {
			Slug:        "e4-0-woman-lifting-weights-dark-skin-tone",
			Character:   "🏋🏿‍♀️",
			UnicodeName: "E4.0 woman lifting weights: dark skin tone",
			CodePoint:   "1F3CB 1F3FF 200D 2640 FE0F",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		127948: {
			Slug:        "e4-0-woman-golfing-dark-skin-tone",
			Character:   "🏌🏿‍♀️",
			UnicodeName: "E4.0 woman golfing: dark skin tone",
			CodePoint:   "1F3CC 1F3FF 200D 2640 FE0F",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		127949: {
			Slug:        "e0-7-motorcycle",
			Character:   "🏍️",
			UnicodeName: "E0.7 motorcycle",
			CodePoint:   "1F3CD FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		127950: {
			Slug:        "e0-7-racing-car",
			Character:   "🏎️",
			UnicodeName: "E0.7 racing car",
			CodePoint:   "1F3CE FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		127951: {
			Slug:        "cricket-game",
			Character:   "🏏",
			UnicodeName: "cricket game",
			CodePoint:   "1F3CF",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127952: {
			Slug:        "volleyball",
			Character:   "🏐",
			UnicodeName: "volleyball",
			CodePoint:   "1F3D0",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127953: {
			Slug:        "field-hockey",
			Character:   "🏑",
			UnicodeName: "field hockey",
			CodePoint:   "1F3D1",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127954: {
			Slug:        "ice-hockey",
			Character:   "🏒",
			UnicodeName: "ice hockey",
			CodePoint:   "1F3D2",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127955: {
			Slug:        "ping-pong",
			Character:   "🏓",
			UnicodeName: "ping pong",
			CodePoint:   "1F3D3",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127956: {
			Slug:        "e0-7-snow-capped-mountain",
			Character:   "🏔️",
			UnicodeName: "E0.7 snow-capped mountain",
			CodePoint:   "1F3D4 FE0F",
			Group:       "travel-places",
			SubGroup:    "place-geographic",
		},

		127957: {
			Slug:        "e0-7-camping",
			Character:   "🏕️",
			UnicodeName: "E0.7 camping",
			CodePoint:   "1F3D5 FE0F",
			Group:       "travel-places",
			SubGroup:    "place-geographic",
		},

		127958: {
			Slug:        "e0-7-beach-with-umbrella",
			Character:   "🏖️",
			UnicodeName: "E0.7 beach with umbrella",
			CodePoint:   "1F3D6 FE0F",
			Group:       "travel-places",
			SubGroup:    "place-geographic",
		},

		127959: {
			Slug:        "e0-7-building-construction",
			Character:   "🏗️",
			UnicodeName: "E0.7 building construction",
			CodePoint:   "1F3D7 FE0F",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127960: {
			Slug:        "e0-7-houses",
			Character:   "🏘️",
			UnicodeName: "E0.7 houses",
			CodePoint:   "1F3D8 FE0F",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127961: {
			Slug:        "e0-7-cityscape",
			Character:   "🏙️",
			UnicodeName: "E0.7 cityscape",
			CodePoint:   "1F3D9 FE0F",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		127962: {
			Slug:        "e0-7-derelict-house",
			Character:   "🏚️",
			UnicodeName: "E0.7 derelict house",
			CodePoint:   "1F3DA FE0F",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127963: {
			Slug:        "e0-7-classical-building",
			Character:   "🏛️",
			UnicodeName: "E0.7 classical building",
			CodePoint:   "1F3DB FE0F",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127964: {
			Slug:        "e0-7-desert",
			Character:   "🏜️",
			UnicodeName: "E0.7 desert",
			CodePoint:   "1F3DC FE0F",
			Group:       "travel-places",
			SubGroup:    "place-geographic",
		},

		127965: {
			Slug:        "e0-7-desert-island",
			Character:   "🏝️",
			UnicodeName: "E0.7 desert island",
			CodePoint:   "1F3DD FE0F",
			Group:       "travel-places",
			SubGroup:    "place-geographic",
		},

		127966: {
			Slug:        "e0-7-national-park",
			Character:   "🏞️",
			UnicodeName: "E0.7 national park",
			CodePoint:   "1F3DE FE0F",
			Group:       "travel-places",
			SubGroup:    "place-geographic",
		},

		127967: {
			Slug:        "e0-7-stadium",
			Character:   "🏟️",
			UnicodeName: "E0.7 stadium",
			CodePoint:   "1F3DF FE0F",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127968: {
			Slug:        "e0-6-house",
			Character:   "🏠",
			UnicodeName: "E0.6 house",
			CodePoint:   "1F3E0",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127969: {
			Slug:        "e0-6-house-with-garden",
			Character:   "🏡",
			UnicodeName: "E0.6 house with garden",
			CodePoint:   "1F3E1",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127970: {
			Slug:        "e0-6-office-building",
			Character:   "🏢",
			UnicodeName: "E0.6 office building",
			CodePoint:   "1F3E2",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127971: {
			Slug:        "e0-6-japanese-post-office",
			Character:   "🏣",
			UnicodeName: "E0.6 Japanese post office",
			CodePoint:   "1F3E3",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127972: {
			Slug:        "e1-0-post-office",
			Character:   "🏤",
			UnicodeName: "E1.0 post office",
			CodePoint:   "1F3E4",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127973: {
			Slug:        "e0-6-hospital",
			Character:   "🏥",
			UnicodeName: "E0.6 hospital",
			CodePoint:   "1F3E5",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127974: {
			Slug:        "e0-6-bank",
			Character:   "🏦",
			UnicodeName: "E0.6 bank",
			CodePoint:   "1F3E6",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127975: {
			Slug:        "atm-sign",
			Character:   "🏧",
			UnicodeName: "ATM sign",
			CodePoint:   "1F3E7",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		127976: {
			Slug:        "e0-6-hotel",
			Character:   "🏨",
			UnicodeName: "E0.6 hotel",
			CodePoint:   "1F3E8",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127977: {
			Slug:        "e0-6-love-hotel",
			Character:   "🏩",
			UnicodeName: "E0.6 love hotel",
			CodePoint:   "1F3E9",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127978: {
			Slug:        "e0-6-convenience-store",
			Character:   "🏪",
			UnicodeName: "E0.6 convenience store",
			CodePoint:   "1F3EA",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127979: {
			Slug:        "e0-6-school",
			Character:   "🏫",
			UnicodeName: "E0.6 school",
			CodePoint:   "1F3EB",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127980: {
			Slug:        "e0-6-department-store",
			Character:   "🏬",
			UnicodeName: "E0.6 department store",
			CodePoint:   "1F3EC",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127981: {
			Slug:        "e0-6-factory",
			Character:   "🏭",
			UnicodeName: "E0.6 factory",
			CodePoint:   "1F3ED",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127982: {
			Slug:        "red-paper-lantern",
			Character:   "🏮",
			UnicodeName: "red paper lantern",
			CodePoint:   "1F3EE",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		127983: {
			Slug:        "e0-6-japanese-castle",
			Character:   "🏯",
			UnicodeName: "E0.6 Japanese castle",
			CodePoint:   "1F3EF",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127984: {
			Slug:        "e0-6-castle",
			Character:   "🏰",
			UnicodeName: "E0.6 castle",
			CodePoint:   "1F3F0",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		127987: {
			Slug:        "transgender-flag",
			Character:   "🏳️‍⚧️",
			UnicodeName: "transgender flag",
			CodePoint:   "1F3F3 FE0F 200D 26A7 FE0F",
			Group:       "flags",
			SubGroup:    "flag",
		},

		127988: {
			Slug:        "flag-england",
			Character:   "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
			UnicodeName: "flag: England",
			CodePoint:   "1F3F4 E0067 E0062 E0065 E006E E0067 E007F",
			Group:       "flags",
			SubGroup:    "subdivision-flag",
		},

		127989: {
			Slug:        "e0-7-rosette",
			Character:   "🏵️",
			UnicodeName: "E0.7 rosette",
			CodePoint:   "1F3F5 FE0F",
			Group:       "animals-nature",
			SubGroup:    "plant-flower",
		},

		127991: {
			Slug:        "label",
			Character:   "🏷️",
			UnicodeName: "label",
			CodePoint:   "1F3F7 FE0F",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		127992: {
			Slug:        "badminton",
			Character:   "🏸",
			UnicodeName: "badminton",
			CodePoint:   "1F3F8",
			Group:       "activities",
			SubGroup:    "sport",
		},

		127993: {
			Slug:        "bow-and-arrow",
			Character:   "🏹",
			UnicodeName: "bow and arrow",
			CodePoint:   "1F3F9",
			Group:       "objects",
			SubGroup:    "tool",
		},

		127994: {
			Slug:        "e1-0-amphora",
			Character:   "🏺",
			UnicodeName: "E1.0 amphora",
			CodePoint:   "1F3FA",
			Group:       "food-drink",
			SubGroup:    "dishware",
		},

		128000: {
			Slug:        "e1-0-rat",
			Character:   "🐀",
			UnicodeName: "E1.0 rat",
			CodePoint:   "1F400",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128001: {
			Slug:        "e1-0-mouse",
			Character:   "🐁",
			UnicodeName: "E1.0 mouse",
			CodePoint:   "1F401",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128002: {
			Slug:        "e1-0-ox",
			Character:   "🐂",
			UnicodeName: "E1.0 ox",
			CodePoint:   "1F402",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128003: {
			Slug:        "e1-0-water-buffalo",
			Character:   "🐃",
			UnicodeName: "E1.0 water buffalo",
			CodePoint:   "1F403",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128004: {
			Slug:        "e1-0-cow",
			Character:   "🐄",
			UnicodeName: "E1.0 cow",
			CodePoint:   "1F404",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128005: {
			Slug:        "e1-0-tiger",
			Character:   "🐅",
			UnicodeName: "E1.0 tiger",
			CodePoint:   "1F405",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128006: {
			Slug:        "e1-0-leopard",
			Character:   "🐆",
			UnicodeName: "E1.0 leopard",
			CodePoint:   "1F406",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128007: {
			Slug:        "e1-0-rabbit",
			Character:   "🐇",
			UnicodeName: "E1.0 rabbit",
			CodePoint:   "1F407",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128008: {
			Slug:        "e0-7-cat",
			Character:   "🐈",
			UnicodeName: "E0.7 cat",
			CodePoint:   "1F408",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128009: {
			Slug:        "e1-0-dragon",
			Character:   "🐉",
			UnicodeName: "E1.0 dragon",
			CodePoint:   "1F409",
			Group:       "animals-nature",
			SubGroup:    "animal-reptile",
		},

		128010: {
			Slug:        "e1-0-crocodile",
			Character:   "🐊",
			UnicodeName: "E1.0 crocodile",
			CodePoint:   "1F40A",
			Group:       "animals-nature",
			SubGroup:    "animal-reptile",
		},

		128011: {
			Slug:        "e1-0-whale",
			Character:   "🐋",
			UnicodeName: "E1.0 whale",
			CodePoint:   "1F40B",
			Group:       "animals-nature",
			SubGroup:    "animal-marine",
		},

		128012: {
			Slug:        "e0-6-snail",
			Character:   "🐌",
			UnicodeName: "E0.6 snail",
			CodePoint:   "1F40C",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		128013: {
			Slug:        "e0-6-snake",
			Character:   "🐍",
			UnicodeName: "E0.6 snake",
			CodePoint:   "1F40D",
			Group:       "animals-nature",
			SubGroup:    "animal-reptile",
		},

		128014: {
			Slug:        "e0-6-horse",
			Character:   "🐎",
			UnicodeName: "E0.6 horse",
			CodePoint:   "1F40E",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128015: {
			Slug:        "e1-0-ram",
			Character:   "🐏",
			UnicodeName: "E1.0 ram",
			CodePoint:   "1F40F",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128016: {
			Slug:        "e1-0-goat",
			Character:   "🐐",
			UnicodeName: "E1.0 goat",
			CodePoint:   "1F410",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128017: {
			Slug:        "e0-6-ewe",
			Character:   "🐑",
			UnicodeName: "E0.6 ewe",
			CodePoint:   "1F411",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128018: {
			Slug:        "e0-6-monkey",
			Character:   "🐒",
			UnicodeName: "E0.6 monkey",
			CodePoint:   "1F412",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128019: {
			Slug:        "e1-0-rooster",
			Character:   "🐓",
			UnicodeName: "E1.0 rooster",
			CodePoint:   "1F413",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		128020: {
			Slug:        "e0-6-chicken",
			Character:   "🐔",
			UnicodeName: "E0.6 chicken",
			CodePoint:   "1F414",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		128021: {
			Slug:        "e12-0-service-dog",
			Character:   "🐕‍🦺",
			UnicodeName: "E12.0 service dog",
			CodePoint:   "1F415 200D 1F9BA",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128022: {
			Slug:        "e1-0-pig",
			Character:   "🐖",
			UnicodeName: "E1.0 pig",
			CodePoint:   "1F416",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128023: {
			Slug:        "e0-6-boar",
			Character:   "🐗",
			UnicodeName: "E0.6 boar",
			CodePoint:   "1F417",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128024: {
			Slug:        "e0-6-elephant",
			Character:   "🐘",
			UnicodeName: "E0.6 elephant",
			CodePoint:   "1F418",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128025: {
			Slug:        "e0-6-octopus",
			Character:   "🐙",
			UnicodeName: "E0.6 octopus",
			CodePoint:   "1F419",
			Group:       "animals-nature",
			SubGroup:    "animal-marine",
		},

		128026: {
			Slug:        "e0-6-spiral-shell",
			Character:   "🐚",
			UnicodeName: "E0.6 spiral shell",
			CodePoint:   "1F41A",
			Group:       "animals-nature",
			SubGroup:    "animal-marine",
		},

		128027: {
			Slug:        "e0-6-bug",
			Character:   "🐛",
			UnicodeName: "E0.6 bug",
			CodePoint:   "1F41B",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		128028: {
			Slug:        "e0-6-ant",
			Character:   "🐜",
			UnicodeName: "E0.6 ant",
			CodePoint:   "1F41C",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		128029: {
			Slug:        "e0-6-honeybee",
			Character:   "🐝",
			UnicodeName: "E0.6 honeybee",
			CodePoint:   "1F41D",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		128030: {
			Slug:        "e0-6-lady-beetle",
			Character:   "🐞",
			UnicodeName: "E0.6 lady beetle",
			CodePoint:   "1F41E",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		128031: {
			Slug:        "e0-6-fish",
			Character:   "🐟",
			UnicodeName: "E0.6 fish",
			CodePoint:   "1F41F",
			Group:       "animals-nature",
			SubGroup:    "animal-marine",
		},

		128032: {
			Slug:        "e0-6-tropical-fish",
			Character:   "🐠",
			UnicodeName: "E0.6 tropical fish",
			CodePoint:   "1F420",
			Group:       "animals-nature",
			SubGroup:    "animal-marine",
		},

		128033: {
			Slug:        "e0-6-blowfish",
			Character:   "🐡",
			UnicodeName: "E0.6 blowfish",
			CodePoint:   "1F421",
			Group:       "animals-nature",
			SubGroup:    "animal-marine",
		},

		128034: {
			Slug:        "e0-6-turtle",
			Character:   "🐢",
			UnicodeName: "E0.6 turtle",
			CodePoint:   "1F422",
			Group:       "animals-nature",
			SubGroup:    "animal-reptile",
		},

		128035: {
			Slug:        "e0-6-hatching-chick",
			Character:   "🐣",
			UnicodeName: "E0.6 hatching chick",
			CodePoint:   "1F423",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		128036: {
			Slug:        "e0-6-baby-chick",
			Character:   "🐤",
			UnicodeName: "E0.6 baby chick",
			CodePoint:   "1F424",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		128037: {
			Slug:        "e0-6-front-facing-baby-chick",
			Character:   "🐥",
			UnicodeName: "E0.6 front-facing baby chick",
			CodePoint:   "1F425",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		128038: {
			Slug:        "e0-6-bird",
			Character:   "🐦",
			UnicodeName: "E0.6 bird",
			CodePoint:   "1F426",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		128039: {
			Slug:        "e0-6-penguin",
			Character:   "🐧",
			UnicodeName: "E0.6 penguin",
			CodePoint:   "1F427",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		128040: {
			Slug:        "e0-6-koala",
			Character:   "🐨",
			UnicodeName: "E0.6 koala",
			CodePoint:   "1F428",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128041: {
			Slug:        "e0-6-poodle",
			Character:   "🐩",
			UnicodeName: "E0.6 poodle",
			CodePoint:   "1F429",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128042: {
			Slug:        "e1-0-camel",
			Character:   "🐪",
			UnicodeName: "E1.0 camel",
			CodePoint:   "1F42A",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128043: {
			Slug:        "e0-6-two-hump-camel",
			Character:   "🐫",
			UnicodeName: "E0.6 two-hump camel",
			CodePoint:   "1F42B",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128044: {
			Slug:        "e0-6-dolphin",
			Character:   "🐬",
			UnicodeName: "E0.6 dolphin",
			CodePoint:   "1F42C",
			Group:       "animals-nature",
			SubGroup:    "animal-marine",
		},

		128045: {
			Slug:        "e0-6-mouse-face",
			Character:   "🐭",
			UnicodeName: "E0.6 mouse face",
			CodePoint:   "1F42D",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128046: {
			Slug:        "e0-6-cow-face",
			Character:   "🐮",
			UnicodeName: "E0.6 cow face",
			CodePoint:   "1F42E",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128047: {
			Slug:        "e0-6-tiger-face",
			Character:   "🐯",
			UnicodeName: "E0.6 tiger face",
			CodePoint:   "1F42F",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128048: {
			Slug:        "e0-6-rabbit-face",
			Character:   "🐰",
			UnicodeName: "E0.6 rabbit face",
			CodePoint:   "1F430",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128049: {
			Slug:        "e0-6-cat-face",
			Character:   "🐱",
			UnicodeName: "E0.6 cat face",
			CodePoint:   "1F431",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128050: {
			Slug:        "e0-6-dragon-face",
			Character:   "🐲",
			UnicodeName: "E0.6 dragon face",
			CodePoint:   "1F432",
			Group:       "animals-nature",
			SubGroup:    "animal-reptile",
		},

		128051: {
			Slug:        "e0-6-spouting-whale",
			Character:   "🐳",
			UnicodeName: "E0.6 spouting whale",
			CodePoint:   "1F433",
			Group:       "animals-nature",
			SubGroup:    "animal-marine",
		},

		128052: {
			Slug:        "e0-6-horse-face",
			Character:   "🐴",
			UnicodeName: "E0.6 horse face",
			CodePoint:   "1F434",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128053: {
			Slug:        "e0-6-monkey-face",
			Character:   "🐵",
			UnicodeName: "E0.6 monkey face",
			CodePoint:   "1F435",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128054: {
			Slug:        "e0-6-dog-face",
			Character:   "🐶",
			UnicodeName: "E0.6 dog face",
			CodePoint:   "1F436",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128055: {
			Slug:        "e0-6-pig-face",
			Character:   "🐷",
			UnicodeName: "E0.6 pig face",
			CodePoint:   "1F437",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128056: {
			Slug:        "e0-6-frog",
			Character:   "🐸",
			UnicodeName: "E0.6 frog",
			CodePoint:   "1F438",
			Group:       "animals-nature",
			SubGroup:    "animal-amphibian",
		},

		128057: {
			Slug:        "e0-6-hamster",
			Character:   "🐹",
			UnicodeName: "E0.6 hamster",
			CodePoint:   "1F439",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128058: {
			Slug:        "e0-6-wolf",
			Character:   "🐺",
			UnicodeName: "E0.6 wolf",
			CodePoint:   "1F43A",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128059: {
			Slug:        "e0-6-bear",
			Character:   "🐻",
			UnicodeName: "E0.6 bear",
			CodePoint:   "1F43B",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128060: {
			Slug:        "e0-6-panda",
			Character:   "🐼",
			UnicodeName: "E0.6 panda",
			CodePoint:   "1F43C",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128061: {
			Slug:        "e0-6-pig-nose",
			Character:   "🐽",
			UnicodeName: "E0.6 pig nose",
			CodePoint:   "1F43D",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128062: {
			Slug:        "e0-6-paw-prints",
			Character:   "🐾",
			UnicodeName: "E0.6 paw prints",
			CodePoint:   "1F43E",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128063: {
			Slug:        "e0-7-chipmunk",
			Character:   "🐿️",
			UnicodeName: "E0.7 chipmunk",
			CodePoint:   "1F43F FE0F",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		128064: {
			Slug:        "e0-6-eyes",
			Character:   "👀",
			UnicodeName: "E0.6 eyes",
			CodePoint:   "1F440",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		128065: {
			Slug:        "e0-7-eye",
			Character:   "👁️",
			UnicodeName: "E0.7 eye",
			CodePoint:   "1F441 FE0F",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		128066: {
			Slug:        "e0-6-ear",
			Character:   "👂",
			UnicodeName: "E0.6 ear",
			CodePoint:   "1F442",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		128067: {
			Slug:        "e0-6-nose",
			Character:   "👃",
			UnicodeName: "E0.6 nose",
			CodePoint:   "1F443",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		128068: {
			Slug:        "e0-6-mouth",
			Character:   "👄",
			UnicodeName: "E0.6 mouth",
			CodePoint:   "1F444",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		128069: {
			Slug:        "e0-6-tongue",
			Character:   "👅",
			UnicodeName: "E0.6 tongue",
			CodePoint:   "1F445",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		128070: {
			Slug:        "e0-6-backhand-index-pointing-up",
			Character:   "👆",
			UnicodeName: "E0.6 backhand index pointing up",
			CodePoint:   "1F446",
			Group:       "people-body",
			SubGroup:    "hand-single-finger",
		},

		128071: {
			Slug:        "e0-6-backhand-index-pointing-down",
			Character:   "👇",
			UnicodeName: "E0.6 backhand index pointing down",
			CodePoint:   "1F447",
			Group:       "people-body",
			SubGroup:    "hand-single-finger",
		},

		128072: {
			Slug:        "e0-6-backhand-index-pointing-left",
			Character:   "👈",
			UnicodeName: "E0.6 backhand index pointing left",
			CodePoint:   "1F448",
			Group:       "people-body",
			SubGroup:    "hand-single-finger",
		},

		128073: {
			Slug:        "e0-6-backhand-index-pointing-right",
			Character:   "👉",
			UnicodeName: "E0.6 backhand index pointing right",
			CodePoint:   "1F449",
			Group:       "people-body",
			SubGroup:    "hand-single-finger",
		},

		128074: {
			Slug:        "e0-6-oncoming-fist",
			Character:   "👊",
			UnicodeName: "E0.6 oncoming fist",
			CodePoint:   "1F44A",
			Group:       "people-body",
			SubGroup:    "hand-fingers-closed",
		},

		128075: {
			Slug:        "e0-6-waving-hand",
			Character:   "👋",
			UnicodeName: "E0.6 waving hand",
			CodePoint:   "1F44B",
			Group:       "people-body",
			SubGroup:    "hand-fingers-open",
		},

		128076: {
			Slug:        "e0-6-ok-hand",
			Character:   "👌",
			UnicodeName: "E0.6 OK hand",
			CodePoint:   "1F44C",
			Group:       "people-body",
			SubGroup:    "hand-fingers-partial",
		},

		128077: {
			Slug:        "e0-6-thumbs-up",
			Character:   "👍",
			UnicodeName: "E0.6 thumbs up",
			CodePoint:   "1F44D",
			Group:       "people-body",
			SubGroup:    "hand-fingers-closed",
		},

		128078: {
			Slug:        "e0-6-thumbs-down",
			Character:   "👎",
			UnicodeName: "E0.6 thumbs down",
			CodePoint:   "1F44E",
			Group:       "people-body",
			SubGroup:    "hand-fingers-closed",
		},

		128079: {
			Slug:        "e0-6-clapping-hands",
			Character:   "👏",
			UnicodeName: "E0.6 clapping hands",
			CodePoint:   "1F44F",
			Group:       "people-body",
			SubGroup:    "hands",
		},

		128080: {
			Slug:        "e0-6-open-hands",
			Character:   "👐",
			UnicodeName: "E0.6 open hands",
			CodePoint:   "1F450",
			Group:       "people-body",
			SubGroup:    "hands",
		},

		128081: {
			Slug:        "crown",
			Character:   "👑",
			UnicodeName: "crown",
			CodePoint:   "1F451",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128082: {
			Slug:        "woman-s-hat",
			Character:   "👒",
			UnicodeName: "woman’s hat",
			CodePoint:   "1F452",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128083: {
			Slug:        "glasses",
			Character:   "👓",
			UnicodeName: "glasses",
			CodePoint:   "1F453",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128084: {
			Slug:        "necktie",
			Character:   "👔",
			UnicodeName: "necktie",
			CodePoint:   "1F454",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128085: {
			Slug:        "t-shirt",
			Character:   "👕",
			UnicodeName: "t-shirt",
			CodePoint:   "1F455",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128086: {
			Slug:        "jeans",
			Character:   "👖",
			UnicodeName: "jeans",
			CodePoint:   "1F456",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128087: {
			Slug:        "dress",
			Character:   "👗",
			UnicodeName: "dress",
			CodePoint:   "1F457",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128088: {
			Slug:        "kimono",
			Character:   "👘",
			UnicodeName: "kimono",
			CodePoint:   "1F458",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128089: {
			Slug:        "bikini",
			Character:   "👙",
			UnicodeName: "bikini",
			CodePoint:   "1F459",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128090: {
			Slug:        "woman-s-clothes",
			Character:   "👚",
			UnicodeName: "woman’s clothes",
			CodePoint:   "1F45A",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128091: {
			Slug:        "purse",
			Character:   "👛",
			UnicodeName: "purse",
			CodePoint:   "1F45B",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128092: {
			Slug:        "handbag",
			Character:   "👜",
			UnicodeName: "handbag",
			CodePoint:   "1F45C",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128093: {
			Slug:        "clutch-bag",
			Character:   "👝",
			UnicodeName: "clutch bag",
			CodePoint:   "1F45D",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128094: {
			Slug:        "man-s-shoe",
			Character:   "👞",
			UnicodeName: "man’s shoe",
			CodePoint:   "1F45E",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128095: {
			Slug:        "running-shoe",
			Character:   "👟",
			UnicodeName: "running shoe",
			CodePoint:   "1F45F",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128096: {
			Slug:        "high-heeled-shoe",
			Character:   "👠",
			UnicodeName: "high-heeled shoe",
			CodePoint:   "1F460",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128097: {
			Slug:        "woman-s-sandal",
			Character:   "👡",
			UnicodeName: "woman’s sandal",
			CodePoint:   "1F461",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128098: {
			Slug:        "woman-s-boot",
			Character:   "👢",
			UnicodeName: "woman’s boot",
			CodePoint:   "1F462",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128099: {
			Slug:        "e0-6-footprints",
			Character:   "👣",
			UnicodeName: "E0.6 footprints",
			CodePoint:   "1F463",
			Group:       "people-body",
			SubGroup:    "person-symbol",
		},

		128100: {
			Slug:        "e0-6-bust-in-silhouette",
			Character:   "👤",
			UnicodeName: "E0.6 bust in silhouette",
			CodePoint:   "1F464",
			Group:       "people-body",
			SubGroup:    "person-symbol",
		},

		128101: {
			Slug:        "e1-0-busts-in-silhouette",
			Character:   "👥",
			UnicodeName: "E1.0 busts in silhouette",
			CodePoint:   "1F465",
			Group:       "people-body",
			SubGroup:    "person-symbol",
		},

		128102: {
			Slug:        "e0-6-boy",
			Character:   "👦",
			UnicodeName: "E0.6 boy",
			CodePoint:   "1F466",
			Group:       "people-body",
			SubGroup:    "person",
		},

		128103: {
			Slug:        "e0-6-girl",
			Character:   "👧",
			UnicodeName: "E0.6 girl",
			CodePoint:   "1F467",
			Group:       "people-body",
			SubGroup:    "person",
		},

		128104: {
			Slug:        "e4-0-family-man-boy",
			Character:   "👨‍👦",
			UnicodeName: "E4.0 family: man, boy",
			CodePoint:   "1F468 200D 1F466",
			Group:       "people-body",
			SubGroup:    "family",
		},

		128105: {
			Slug:        "e4-0-family-woman-boy",
			Character:   "👩‍👦",
			UnicodeName: "E4.0 family: woman, boy",
			CodePoint:   "1F469 200D 1F466",
			Group:       "people-body",
			SubGroup:    "family",
		},

		128106: {
			Slug:        "e0-6-family",
			Character:   "👪",
			UnicodeName: "E0.6 family",
			CodePoint:   "1F46A",
			Group:       "people-body",
			SubGroup:    "family",
		},

		128107: {
			Slug:        "e12-0-woman-and-man-holding-hands-dark-skin-tone",
			Character:   "👫🏿",
			UnicodeName: "E12.0 woman and man holding hands: dark skin tone",
			CodePoint:   "1F46B 1F3FF",
			Group:       "people-body",
			SubGroup:    "family",
		},

		128108: {
			Slug:        "e12-0-men-holding-hands-dark-skin-tone",
			Character:   "👬🏿",
			UnicodeName: "E12.0 men holding hands: dark skin tone",
			CodePoint:   "1F46C 1F3FF",
			Group:       "people-body",
			SubGroup:    "family",
		},

		128109: {
			Slug:        "e12-0-women-holding-hands-dark-skin-tone",
			Character:   "👭🏿",
			UnicodeName: "E12.0 women holding hands: dark skin tone",
			CodePoint:   "1F46D 1F3FF",
			Group:       "people-body",
			SubGroup:    "family",
		},

		128110: {
			Slug:        "e0-6-police-officer",
			Character:   "👮",
			UnicodeName: "E0.6 police officer",
			CodePoint:   "1F46E",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		128111: {
			Slug:        "e0-6-people-with-bunny-ears",
			Character:   "👯",
			UnicodeName: "E0.6 people with bunny ears",
			CodePoint:   "1F46F",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		128112: {
			Slug:        "e0-6-person-with-veil",
			Character:   "👰",
			UnicodeName: "E0.6 person with veil",
			CodePoint:   "1F470",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		128113: {
			Slug:        "e4-0-woman-blond-hair",
			Character:   "👱‍♀️",
			UnicodeName: "E4.0 woman: blond hair",
			CodePoint:   "1F471 200D 2640 FE0F",
			Group:       "people-body",
			SubGroup:    "person",
		},

		128114: {
			Slug:        "e0-6-person-with-skullcap",
			Character:   "👲",
			UnicodeName: "E0.6 person with skullcap",
			CodePoint:   "1F472",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		128115: {
			Slug:        "e0-6-person-wearing-turban",
			Character:   "👳",
			UnicodeName: "E0.6 person wearing turban",
			CodePoint:   "1F473",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		128116: {
			Slug:        "e0-6-old-man",
			Character:   "👴",
			UnicodeName: "E0.6 old man",
			CodePoint:   "1F474",
			Group:       "people-body",
			SubGroup:    "person",
		},

		128117: {
			Slug:        "e0-6-old-woman",
			Character:   "👵",
			UnicodeName: "E0.6 old woman",
			CodePoint:   "1F475",
			Group:       "people-body",
			SubGroup:    "person",
		},

		128118: {
			Slug:        "e0-6-baby",
			Character:   "👶",
			UnicodeName: "E0.6 baby",
			CodePoint:   "1F476",
			Group:       "people-body",
			SubGroup:    "person",
		},

		128119: {
			Slug:        "e0-6-construction-worker",
			Character:   "👷",
			UnicodeName: "E0.6 construction worker",
			CodePoint:   "1F477",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		128120: {
			Slug:        "e0-6-princess",
			Character:   "👸",
			UnicodeName: "E0.6 princess",
			CodePoint:   "1F478",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		128121: {
			Slug:        "e0-6-ogre",
			Character:   "👹",
			UnicodeName: "E0.6 ogre",
			CodePoint:   "1F479",
			Group:       "smileys-emotion",
			SubGroup:    "face-costume",
		},

		128122: {
			Slug:        "e0-6-goblin",
			Character:   "👺",
			UnicodeName: "E0.6 goblin",
			CodePoint:   "1F47A",
			Group:       "smileys-emotion",
			SubGroup:    "face-costume",
		},

		128123: {
			Slug:        "e0-6-ghost",
			Character:   "👻",
			UnicodeName: "E0.6 ghost",
			CodePoint:   "1F47B",
			Group:       "smileys-emotion",
			SubGroup:    "face-costume",
		},

		128124: {
			Slug:        "e0-6-baby-angel",
			Character:   "👼",
			UnicodeName: "E0.6 baby angel",
			CodePoint:   "1F47C",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		128125: {
			Slug:        "e0-6-alien",
			Character:   "👽",
			UnicodeName: "E0.6 alien",
			CodePoint:   "1F47D",
			Group:       "smileys-emotion",
			SubGroup:    "face-costume",
		},

		128126: {
			Slug:        "e0-6-alien-monster",
			Character:   "👾",
			UnicodeName: "E0.6 alien monster",
			CodePoint:   "1F47E",
			Group:       "smileys-emotion",
			SubGroup:    "face-costume",
		},

		128127: {
			Slug:        "e0-6-angry-face-with-horns",
			Character:   "👿",
			UnicodeName: "E0.6 angry face with horns",
			CodePoint:   "1F47F",
			Group:       "smileys-emotion",
			SubGroup:    "face-negative",
		},

		128128: {
			Slug:        "e0-6-skull",
			Character:   "💀",
			UnicodeName: "E0.6 skull",
			CodePoint:   "1F480",
			Group:       "smileys-emotion",
			SubGroup:    "face-negative",
		},

		128129: {
			Slug:        "e0-6-person-tipping-hand",
			Character:   "💁",
			UnicodeName: "E0.6 person tipping hand",
			CodePoint:   "1F481",
			Group:       "people-body",
			SubGroup:    "person-gesture",
		},

		128130: {
			Slug:        "e0-6-guard",
			Character:   "💂",
			UnicodeName: "E0.6 guard",
			CodePoint:   "1F482",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		128131: {
			Slug:        "e0-6-woman-dancing",
			Character:   "💃",
			UnicodeName: "E0.6 woman dancing",
			CodePoint:   "1F483",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		128132: {
			Slug:        "lipstick",
			Character:   "💄",
			UnicodeName: "lipstick",
			CodePoint:   "1F484",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128133: {
			Slug:        "e0-6-nail-polish",
			Character:   "💅",
			UnicodeName: "E0.6 nail polish",
			CodePoint:   "1F485",
			Group:       "people-body",
			SubGroup:    "hand-prop",
		},

		128134: {
			Slug:        "e0-6-person-getting-massage",
			Character:   "💆",
			UnicodeName: "E0.6 person getting massage",
			CodePoint:   "1F486",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		128135: {
			Slug:        "e0-6-person-getting-haircut",
			Character:   "💇",
			UnicodeName: "E0.6 person getting haircut",
			CodePoint:   "1F487",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		128136: {
			Slug:        "e0-6-barber-pole",
			Character:   "💈",
			UnicodeName: "E0.6 barber pole",
			CodePoint:   "1F488",
			Group:       "travel-places",
			SubGroup:    "place-other",
		},

		128137: {
			Slug:        "syringe",
			Character:   "💉",
			UnicodeName: "syringe",
			CodePoint:   "1F489",
			Group:       "objects",
			SubGroup:    "medical",
		},

		128138: {
			Slug:        "pill",
			Character:   "💊",
			UnicodeName: "pill",
			CodePoint:   "1F48A",
			Group:       "objects",
			SubGroup:    "medical",
		},

		128139: {
			Slug:        "e0-6-kiss-mark",
			Character:   "💋",
			UnicodeName: "E0.6 kiss mark",
			CodePoint:   "1F48B",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128140: {
			Slug:        "e0-6-love-letter",
			Character:   "💌",
			UnicodeName: "E0.6 love letter",
			CodePoint:   "1F48C",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128141: {
			Slug:        "ring",
			Character:   "💍",
			UnicodeName: "ring",
			CodePoint:   "1F48D",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128142: {
			Slug:        "gem-stone",
			Character:   "💎",
			UnicodeName: "gem stone",
			CodePoint:   "1F48E",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128143: {
			Slug:        "e0-6-kiss",
			Character:   "💏",
			UnicodeName: "E0.6 kiss",
			CodePoint:   "1F48F",
			Group:       "people-body",
			SubGroup:    "family",
		},

		128144: {
			Slug:        "e0-6-bouquet",
			Character:   "💐",
			UnicodeName: "E0.6 bouquet",
			CodePoint:   "1F490",
			Group:       "animals-nature",
			SubGroup:    "plant-flower",
		},

		128145: {
			Slug:        "e0-6-couple-with-heart",
			Character:   "💑",
			UnicodeName: "E0.6 couple with heart",
			CodePoint:   "1F491",
			Group:       "people-body",
			SubGroup:    "family",
		},

		128146: {
			Slug:        "e0-6-wedding",
			Character:   "💒",
			UnicodeName: "E0.6 wedding",
			CodePoint:   "1F492",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		128147: {
			Slug:        "e0-6-beating-heart",
			Character:   "💓",
			UnicodeName: "E0.6 beating heart",
			CodePoint:   "1F493",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128148: {
			Slug:        "e0-6-broken-heart",
			Character:   "💔",
			UnicodeName: "E0.6 broken heart",
			CodePoint:   "1F494",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128149: {
			Slug:        "e0-6-two-hearts",
			Character:   "💕",
			UnicodeName: "E0.6 two hearts",
			CodePoint:   "1F495",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128150: {
			Slug:        "e0-6-sparkling-heart",
			Character:   "💖",
			UnicodeName: "E0.6 sparkling heart",
			CodePoint:   "1F496",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128151: {
			Slug:        "e0-6-growing-heart",
			Character:   "💗",
			UnicodeName: "E0.6 growing heart",
			CodePoint:   "1F497",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128152: {
			Slug:        "e0-6-heart-with-arrow",
			Character:   "💘",
			UnicodeName: "E0.6 heart with arrow",
			CodePoint:   "1F498",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128153: {
			Slug:        "e0-6-blue-heart",
			Character:   "💙",
			UnicodeName: "E0.6 blue heart",
			CodePoint:   "1F499",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128154: {
			Slug:        "e0-6-green-heart",
			Character:   "💚",
			UnicodeName: "E0.6 green heart",
			CodePoint:   "1F49A",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128155: {
			Slug:        "e0-6-yellow-heart",
			Character:   "💛",
			UnicodeName: "E0.6 yellow heart",
			CodePoint:   "1F49B",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128156: {
			Slug:        "e0-6-purple-heart",
			Character:   "💜",
			UnicodeName: "E0.6 purple heart",
			CodePoint:   "1F49C",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128157: {
			Slug:        "e0-6-heart-with-ribbon",
			Character:   "💝",
			UnicodeName: "E0.6 heart with ribbon",
			CodePoint:   "1F49D",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128158: {
			Slug:        "e0-6-revolving-hearts",
			Character:   "💞",
			UnicodeName: "E0.6 revolving hearts",
			CodePoint:   "1F49E",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128159: {
			Slug:        "e0-6-heart-decoration",
			Character:   "💟",
			UnicodeName: "E0.6 heart decoration",
			CodePoint:   "1F49F",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128160: {
			Slug:        "diamond-with-a-dot",
			Character:   "💠",
			UnicodeName: "diamond with a dot",
			CodePoint:   "1F4A0",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128161: {
			Slug:        "light-bulb",
			Character:   "💡",
			UnicodeName: "light bulb",
			CodePoint:   "1F4A1",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128162: {
			Slug:        "e0-6-anger-symbol",
			Character:   "💢",
			UnicodeName: "E0.6 anger symbol",
			CodePoint:   "1F4A2",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128163: {
			Slug:        "e0-6-bomb",
			Character:   "💣",
			UnicodeName: "E0.6 bomb",
			CodePoint:   "1F4A3",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128164: {
			Slug:        "e0-6-zzz",
			Character:   "💤",
			UnicodeName: "E0.6 zzz",
			CodePoint:   "1F4A4",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128165: {
			Slug:        "e0-6-collision",
			Character:   "💥",
			UnicodeName: "E0.6 collision",
			CodePoint:   "1F4A5",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128166: {
			Slug:        "e0-6-sweat-droplets",
			Character:   "💦",
			UnicodeName: "E0.6 sweat droplets",
			CodePoint:   "1F4A6",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128167: {
			Slug:        "droplet",
			Character:   "💧",
			UnicodeName: "droplet",
			CodePoint:   "1F4A7",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		128168: {
			Slug:        "e0-6-dashing-away",
			Character:   "💨",
			UnicodeName: "E0.6 dashing away",
			CodePoint:   "1F4A8",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128169: {
			Slug:        "e0-6-pile-of-poo",
			Character:   "💩",
			UnicodeName: "E0.6 pile of poo",
			CodePoint:   "1F4A9",
			Group:       "smileys-emotion",
			SubGroup:    "face-costume",
		},

		128170: {
			Slug:        "e0-6-flexed-biceps",
			Character:   "💪",
			UnicodeName: "E0.6 flexed biceps",
			CodePoint:   "1F4AA",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		128171: {
			Slug:        "e0-6-dizzy",
			Character:   "💫",
			UnicodeName: "E0.6 dizzy",
			CodePoint:   "1F4AB",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128172: {
			Slug:        "e0-6-speech-balloon",
			Character:   "💬",
			UnicodeName: "E0.6 speech balloon",
			CodePoint:   "1F4AC",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128173: {
			Slug:        "e1-0-thought-balloon",
			Character:   "💭",
			UnicodeName: "E1.0 thought balloon",
			CodePoint:   "1F4AD",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128174: {
			Slug:        "e0-6-white-flower",
			Character:   "💮",
			UnicodeName: "E0.6 white flower",
			CodePoint:   "1F4AE",
			Group:       "animals-nature",
			SubGroup:    "plant-flower",
		},

		128175: {
			Slug:        "e0-6-hundred-points",
			Character:   "💯",
			UnicodeName: "E0.6 hundred points",
			CodePoint:   "1F4AF",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128176: {
			Slug:        "money-bag",
			Character:   "💰",
			UnicodeName: "money bag",
			CodePoint:   "1F4B0",
			Group:       "objects",
			SubGroup:    "money",
		},

		128177: {
			Slug:        "currency-exchange",
			Character:   "💱",
			UnicodeName: "currency exchange",
			CodePoint:   "1F4B1",
			Group:       "symbols",
			SubGroup:    "currency",
		},

		128178: {
			Slug:        "heavy-dollar-sign",
			Character:   "💲",
			UnicodeName: "heavy dollar sign",
			CodePoint:   "1F4B2",
			Group:       "symbols",
			SubGroup:    "currency",
		},

		128179: {
			Slug:        "credit-card",
			Character:   "💳",
			UnicodeName: "credit card",
			CodePoint:   "1F4B3",
			Group:       "objects",
			SubGroup:    "money",
		},

		128180: {
			Slug:        "yen-banknote",
			Character:   "💴",
			UnicodeName: "yen banknote",
			CodePoint:   "1F4B4",
			Group:       "objects",
			SubGroup:    "money",
		},

		128181: {
			Slug:        "dollar-banknote",
			Character:   "💵",
			UnicodeName: "dollar banknote",
			CodePoint:   "1F4B5",
			Group:       "objects",
			SubGroup:    "money",
		},

		128182: {
			Slug:        "euro-banknote",
			Character:   "💶",
			UnicodeName: "euro banknote",
			CodePoint:   "1F4B6",
			Group:       "objects",
			SubGroup:    "money",
		},

		128183: {
			Slug:        "pound-banknote",
			Character:   "💷",
			UnicodeName: "pound banknote",
			CodePoint:   "1F4B7",
			Group:       "objects",
			SubGroup:    "money",
		},

		128184: {
			Slug:        "money-with-wings",
			Character:   "💸",
			UnicodeName: "money with wings",
			CodePoint:   "1F4B8",
			Group:       "objects",
			SubGroup:    "money",
		},

		128185: {
			Slug:        "chart-increasing-with-yen",
			Character:   "💹",
			UnicodeName: "chart increasing with yen",
			CodePoint:   "1F4B9",
			Group:       "objects",
			SubGroup:    "money",
		},

		128186: {
			Slug:        "e0-6-seat",
			Character:   "💺",
			UnicodeName: "E0.6 seat",
			CodePoint:   "1F4BA",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128187: {
			Slug:        "laptop",
			Character:   "💻",
			UnicodeName: "laptop",
			CodePoint:   "1F4BB",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128188: {
			Slug:        "briefcase",
			Character:   "💼",
			UnicodeName: "briefcase",
			CodePoint:   "1F4BC",
			Group:       "objects",
			SubGroup:    "office",
		},

		128189: {
			Slug:        "computer-disk",
			Character:   "💽",
			UnicodeName: "computer disk",
			CodePoint:   "1F4BD",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128190: {
			Slug:        "floppy-disk",
			Character:   "💾",
			UnicodeName: "floppy disk",
			CodePoint:   "1F4BE",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128191: {
			Slug:        "optical-disk",
			Character:   "💿",
			UnicodeName: "optical disk",
			CodePoint:   "1F4BF",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128192: {
			Slug:        "dvd",
			Character:   "📀",
			UnicodeName: "dvd",
			CodePoint:   "1F4C0",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128193: {
			Slug:        "file-folder",
			Character:   "📁",
			UnicodeName: "file folder",
			CodePoint:   "1F4C1",
			Group:       "objects",
			SubGroup:    "office",
		},

		128194: {
			Slug:        "open-file-folder",
			Character:   "📂",
			UnicodeName: "open file folder",
			CodePoint:   "1F4C2",
			Group:       "objects",
			SubGroup:    "office",
		},

		128195: {
			Slug:        "page-with-curl",
			Character:   "📃",
			UnicodeName: "page with curl",
			CodePoint:   "1F4C3",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128196: {
			Slug:        "page-facing-up",
			Character:   "📄",
			UnicodeName: "page facing up",
			CodePoint:   "1F4C4",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128197: {
			Slug:        "calendar",
			Character:   "📅",
			UnicodeName: "calendar",
			CodePoint:   "1F4C5",
			Group:       "objects",
			SubGroup:    "office",
		},

		128198: {
			Slug:        "tear-off-calendar",
			Character:   "📆",
			UnicodeName: "tear-off calendar",
			CodePoint:   "1F4C6",
			Group:       "objects",
			SubGroup:    "office",
		},

		128199: {
			Slug:        "card-index",
			Character:   "📇",
			UnicodeName: "card index",
			CodePoint:   "1F4C7",
			Group:       "objects",
			SubGroup:    "office",
		},

		128200: {
			Slug:        "chart-increasing",
			Character:   "📈",
			UnicodeName: "chart increasing",
			CodePoint:   "1F4C8",
			Group:       "objects",
			SubGroup:    "office",
		},

		128201: {
			Slug:        "chart-decreasing",
			Character:   "📉",
			UnicodeName: "chart decreasing",
			CodePoint:   "1F4C9",
			Group:       "objects",
			SubGroup:    "office",
		},

		128202: {
			Slug:        "bar-chart",
			Character:   "📊",
			UnicodeName: "bar chart",
			CodePoint:   "1F4CA",
			Group:       "objects",
			SubGroup:    "office",
		},

		128203: {
			Slug:        "clipboard",
			Character:   "📋",
			UnicodeName: "clipboard",
			CodePoint:   "1F4CB",
			Group:       "objects",
			SubGroup:    "office",
		},

		128204: {
			Slug:        "pushpin",
			Character:   "📌",
			UnicodeName: "pushpin",
			CodePoint:   "1F4CC",
			Group:       "objects",
			SubGroup:    "office",
		},

		128205: {
			Slug:        "round-pushpin",
			Character:   "📍",
			UnicodeName: "round pushpin",
			CodePoint:   "1F4CD",
			Group:       "objects",
			SubGroup:    "office",
		},

		128206: {
			Slug:        "paperclip",
			Character:   "📎",
			UnicodeName: "paperclip",
			CodePoint:   "1F4CE",
			Group:       "objects",
			SubGroup:    "office",
		},

		128207: {
			Slug:        "straight-ruler",
			Character:   "📏",
			UnicodeName: "straight ruler",
			CodePoint:   "1F4CF",
			Group:       "objects",
			SubGroup:    "office",
		},

		128208: {
			Slug:        "triangular-ruler",
			Character:   "📐",
			UnicodeName: "triangular ruler",
			CodePoint:   "1F4D0",
			Group:       "objects",
			SubGroup:    "office",
		},

		128209: {
			Slug:        "bookmark-tabs",
			Character:   "📑",
			UnicodeName: "bookmark tabs",
			CodePoint:   "1F4D1",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128210: {
			Slug:        "ledger",
			Character:   "📒",
			UnicodeName: "ledger",
			CodePoint:   "1F4D2",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128211: {
			Slug:        "notebook",
			Character:   "📓",
			UnicodeName: "notebook",
			CodePoint:   "1F4D3",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128212: {
			Slug:        "notebook-with-decorative-cover",
			Character:   "📔",
			UnicodeName: "notebook with decorative cover",
			CodePoint:   "1F4D4",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128213: {
			Slug:        "closed-book",
			Character:   "📕",
			UnicodeName: "closed book",
			CodePoint:   "1F4D5",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128214: {
			Slug:        "open-book",
			Character:   "📖",
			UnicodeName: "open book",
			CodePoint:   "1F4D6",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128215: {
			Slug:        "green-book",
			Character:   "📗",
			UnicodeName: "green book",
			CodePoint:   "1F4D7",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128216: {
			Slug:        "blue-book",
			Character:   "📘",
			UnicodeName: "blue book",
			CodePoint:   "1F4D8",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128217: {
			Slug:        "orange-book",
			Character:   "📙",
			UnicodeName: "orange book",
			CodePoint:   "1F4D9",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128218: {
			Slug:        "books",
			Character:   "📚",
			UnicodeName: "books",
			CodePoint:   "1F4DA",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128219: {
			Slug:        "name-badge",
			Character:   "📛",
			UnicodeName: "name badge",
			CodePoint:   "1F4DB",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		128220: {
			Slug:        "scroll",
			Character:   "📜",
			UnicodeName: "scroll",
			CodePoint:   "1F4DC",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128221: {
			Slug:        "memo",
			Character:   "📝",
			UnicodeName: "memo",
			CodePoint:   "1F4DD",
			Group:       "objects",
			SubGroup:    "writing",
		},

		128222: {
			Slug:        "telephone-receiver",
			Character:   "📞",
			UnicodeName: "telephone receiver",
			CodePoint:   "1F4DE",
			Group:       "objects",
			SubGroup:    "phone",
		},

		128223: {
			Slug:        "pager",
			Character:   "📟",
			UnicodeName: "pager",
			CodePoint:   "1F4DF",
			Group:       "objects",
			SubGroup:    "phone",
		},

		128224: {
			Slug:        "fax-machine",
			Character:   "📠",
			UnicodeName: "fax machine",
			CodePoint:   "1F4E0",
			Group:       "objects",
			SubGroup:    "phone",
		},

		128225: {
			Slug:        "satellite-antenna",
			Character:   "📡",
			UnicodeName: "satellite antenna",
			CodePoint:   "1F4E1",
			Group:       "objects",
			SubGroup:    "science",
		},

		128226: {
			Slug:        "loudspeaker",
			Character:   "📢",
			UnicodeName: "loudspeaker",
			CodePoint:   "1F4E2",
			Group:       "objects",
			SubGroup:    "sound",
		},

		128227: {
			Slug:        "megaphone",
			Character:   "📣",
			UnicodeName: "megaphone",
			CodePoint:   "1F4E3",
			Group:       "objects",
			SubGroup:    "sound",
		},

		128228: {
			Slug:        "outbox-tray",
			Character:   "📤",
			UnicodeName: "outbox tray",
			CodePoint:   "1F4E4",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128229: {
			Slug:        "inbox-tray",
			Character:   "📥",
			UnicodeName: "inbox tray",
			CodePoint:   "1F4E5",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128230: {
			Slug:        "package",
			Character:   "📦",
			UnicodeName: "package",
			CodePoint:   "1F4E6",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128231: {
			Slug:        "e-mail",
			Character:   "📧",
			UnicodeName: "e-mail",
			CodePoint:   "1F4E7",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128232: {
			Slug:        "incoming-envelope",
			Character:   "📨",
			UnicodeName: "incoming envelope",
			CodePoint:   "1F4E8",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128233: {
			Slug:        "envelope-with-arrow",
			Character:   "📩",
			UnicodeName: "envelope with arrow",
			CodePoint:   "1F4E9",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128234: {
			Slug:        "closed-mailbox-with-lowered-flag",
			Character:   "📪",
			UnicodeName: "closed mailbox with lowered flag",
			CodePoint:   "1F4EA",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128235: {
			Slug:        "closed-mailbox-with-raised-flag",
			Character:   "📫",
			UnicodeName: "closed mailbox with raised flag",
			CodePoint:   "1F4EB",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128236: {
			Slug:        "open-mailbox-with-raised-flag",
			Character:   "📬",
			UnicodeName: "open mailbox with raised flag",
			CodePoint:   "1F4EC",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128237: {
			Slug:        "open-mailbox-with-lowered-flag",
			Character:   "📭",
			UnicodeName: "open mailbox with lowered flag",
			CodePoint:   "1F4ED",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128238: {
			Slug:        "postbox",
			Character:   "📮",
			UnicodeName: "postbox",
			CodePoint:   "1F4EE",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128239: {
			Slug:        "postal-horn",
			Character:   "📯",
			UnicodeName: "postal horn",
			CodePoint:   "1F4EF",
			Group:       "objects",
			SubGroup:    "sound",
		},

		128240: {
			Slug:        "newspaper",
			Character:   "📰",
			UnicodeName: "newspaper",
			CodePoint:   "1F4F0",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128241: {
			Slug:        "mobile-phone",
			Character:   "📱",
			UnicodeName: "mobile phone",
			CodePoint:   "1F4F1",
			Group:       "objects",
			SubGroup:    "phone",
		},

		128242: {
			Slug:        "mobile-phone-with-arrow",
			Character:   "📲",
			UnicodeName: "mobile phone with arrow",
			CodePoint:   "1F4F2",
			Group:       "objects",
			SubGroup:    "phone",
		},

		128243: {
			Slug:        "vibration-mode",
			Character:   "📳",
			UnicodeName: "vibration mode",
			CodePoint:   "1F4F3",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		128244: {
			Slug:        "mobile-phone-off",
			Character:   "📴",
			UnicodeName: "mobile phone off",
			CodePoint:   "1F4F4",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		128245: {
			Slug:        "no-mobile-phones",
			Character:   "📵",
			UnicodeName: "no mobile phones",
			CodePoint:   "1F4F5",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		128246: {
			Slug:        "antenna-bars",
			Character:   "📶",
			UnicodeName: "antenna bars",
			CodePoint:   "1F4F6",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		128247: {
			Slug:        "camera",
			Character:   "📷",
			UnicodeName: "camera",
			CodePoint:   "1F4F7",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128248: {
			Slug:        "camera-with-flash",
			Character:   "📸",
			UnicodeName: "camera with flash",
			CodePoint:   "1F4F8",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128249: {
			Slug:        "video-camera",
			Character:   "📹",
			UnicodeName: "video camera",
			CodePoint:   "1F4F9",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128250: {
			Slug:        "television",
			Character:   "📺",
			UnicodeName: "television",
			CodePoint:   "1F4FA",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128251: {
			Slug:        "radio",
			Character:   "📻",
			UnicodeName: "radio",
			CodePoint:   "1F4FB",
			Group:       "objects",
			SubGroup:    "music",
		},

		128252: {
			Slug:        "videocassette",
			Character:   "📼",
			UnicodeName: "videocassette",
			CodePoint:   "1F4FC",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128253: {
			Slug:        "film-projector",
			Character:   "📽️",
			UnicodeName: "film projector",
			CodePoint:   "1F4FD FE0F",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128255: {
			Slug:        "prayer-beads",
			Character:   "📿",
			UnicodeName: "prayer beads",
			CodePoint:   "1F4FF",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128256: {
			Slug:        "shuffle-tracks-button",
			Character:   "🔀",
			UnicodeName: "shuffle tracks button",
			CodePoint:   "1F500",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		128257: {
			Slug:        "repeat-button",
			Character:   "🔁",
			UnicodeName: "repeat button",
			CodePoint:   "1F501",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		128258: {
			Slug:        "repeat-single-button",
			Character:   "🔂",
			UnicodeName: "repeat single button",
			CodePoint:   "1F502",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		128259: {
			Slug:        "clockwise-vertical-arrows",
			Character:   "🔃",
			UnicodeName: "clockwise vertical arrows",
			CodePoint:   "1F503",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		128260: {
			Slug:        "counterclockwise-arrows-button",
			Character:   "🔄",
			UnicodeName: "counterclockwise arrows button",
			CodePoint:   "1F504",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		128261: {
			Slug:        "dim-button",
			Character:   "🔅",
			UnicodeName: "dim button",
			CodePoint:   "1F505",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		128262: {
			Slug:        "bright-button",
			Character:   "🔆",
			UnicodeName: "bright button",
			CodePoint:   "1F506",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		128263: {
			Slug:        "muted-speaker",
			Character:   "🔇",
			UnicodeName: "muted speaker",
			CodePoint:   "1F507",
			Group:       "objects",
			SubGroup:    "sound",
		},

		128264: {
			Slug:        "speaker-low-volume",
			Character:   "🔈",
			UnicodeName: "speaker low volume",
			CodePoint:   "1F508",
			Group:       "objects",
			SubGroup:    "sound",
		},

		128265: {
			Slug:        "speaker-medium-volume",
			Character:   "🔉",
			UnicodeName: "speaker medium volume",
			CodePoint:   "1F509",
			Group:       "objects",
			SubGroup:    "sound",
		},

		128266: {
			Slug:        "speaker-high-volume",
			Character:   "🔊",
			UnicodeName: "speaker high volume",
			CodePoint:   "1F50A",
			Group:       "objects",
			SubGroup:    "sound",
		},

		128267: {
			Slug:        "battery",
			Character:   "🔋",
			UnicodeName: "battery",
			CodePoint:   "1F50B",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128268: {
			Slug:        "electric-plug",
			Character:   "🔌",
			UnicodeName: "electric plug",
			CodePoint:   "1F50C",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128269: {
			Slug:        "magnifying-glass-tilted-left",
			Character:   "🔍",
			UnicodeName: "magnifying glass tilted left",
			CodePoint:   "1F50D",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128270: {
			Slug:        "magnifying-glass-tilted-right",
			Character:   "🔎",
			UnicodeName: "magnifying glass tilted right",
			CodePoint:   "1F50E",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128271: {
			Slug:        "locked-with-pen",
			Character:   "🔏",
			UnicodeName: "locked with pen",
			CodePoint:   "1F50F",
			Group:       "objects",
			SubGroup:    "lock",
		},

		128272: {
			Slug:        "locked-with-key",
			Character:   "🔐",
			UnicodeName: "locked with key",
			CodePoint:   "1F510",
			Group:       "objects",
			SubGroup:    "lock",
		},

		128273: {
			Slug:        "key",
			Character:   "🔑",
			UnicodeName: "key",
			CodePoint:   "1F511",
			Group:       "objects",
			SubGroup:    "lock",
		},

		128274: {
			Slug:        "locked",
			Character:   "🔒",
			UnicodeName: "locked",
			CodePoint:   "1F512",
			Group:       "objects",
			SubGroup:    "lock",
		},

		128275: {
			Slug:        "unlocked",
			Character:   "🔓",
			UnicodeName: "unlocked",
			CodePoint:   "1F513",
			Group:       "objects",
			SubGroup:    "lock",
		},

		128276: {
			Slug:        "bell",
			Character:   "🔔",
			UnicodeName: "bell",
			CodePoint:   "1F514",
			Group:       "objects",
			SubGroup:    "sound",
		},

		128277: {
			Slug:        "bell-with-slash",
			Character:   "🔕",
			UnicodeName: "bell with slash",
			CodePoint:   "1F515",
			Group:       "objects",
			SubGroup:    "sound",
		},

		128278: {
			Slug:        "bookmark",
			Character:   "🔖",
			UnicodeName: "bookmark",
			CodePoint:   "1F516",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128279: {
			Slug:        "link",
			Character:   "🔗",
			UnicodeName: "link",
			CodePoint:   "1F517",
			Group:       "objects",
			SubGroup:    "tool",
		},

		128280: {
			Slug:        "radio-button",
			Character:   "🔘",
			UnicodeName: "radio button",
			CodePoint:   "1F518",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128281: {
			Slug:        "back-arrow",
			Character:   "🔙",
			UnicodeName: "BACK arrow",
			CodePoint:   "1F519",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		128282: {
			Slug:        "end-arrow",
			Character:   "🔚",
			UnicodeName: "END arrow",
			CodePoint:   "1F51A",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		128283: {
			Slug:        "on-arrow",
			Character:   "🔛",
			UnicodeName: "ON! arrow",
			CodePoint:   "1F51B",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		128284: {
			Slug:        "soon-arrow",
			Character:   "🔜",
			UnicodeName: "SOON arrow",
			CodePoint:   "1F51C",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		128285: {
			Slug:        "top-arrow",
			Character:   "🔝",
			UnicodeName: "TOP arrow",
			CodePoint:   "1F51D",
			Group:       "symbols",
			SubGroup:    "arrow",
		},

		128286: {
			Slug:        "no-one-under-eighteen",
			Character:   "🔞",
			UnicodeName: "no one under eighteen",
			CodePoint:   "1F51E",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		128287: {
			Slug:        "keycap-10",
			Character:   "🔟",
			UnicodeName: "keycap: 10",
			CodePoint:   "1F51F",
			Group:       "symbols",
			SubGroup:    "keycap",
		},

		128288: {
			Slug:        "input-latin-uppercase",
			Character:   "🔠",
			UnicodeName: "input latin uppercase",
			CodePoint:   "1F520",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		128289: {
			Slug:        "input-latin-lowercase",
			Character:   "🔡",
			UnicodeName: "input latin lowercase",
			CodePoint:   "1F521",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		128290: {
			Slug:        "input-numbers",
			Character:   "🔢",
			UnicodeName: "input numbers",
			CodePoint:   "1F522",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		128291: {
			Slug:        "input-symbols",
			Character:   "🔣",
			UnicodeName: "input symbols",
			CodePoint:   "1F523",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		128292: {
			Slug:        "input-latin-letters",
			Character:   "🔤",
			UnicodeName: "input latin letters",
			CodePoint:   "1F524",
			Group:       "symbols",
			SubGroup:    "alphanum",
		},

		128293: {
			Slug:        "fire",
			Character:   "🔥",
			UnicodeName: "fire",
			CodePoint:   "1F525",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		128294: {
			Slug:        "flashlight",
			Character:   "🔦",
			UnicodeName: "flashlight",
			CodePoint:   "1F526",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128295: {
			Slug:        "wrench",
			Character:   "🔧",
			UnicodeName: "wrench",
			CodePoint:   "1F527",
			Group:       "objects",
			SubGroup:    "tool",
		},

		128296: {
			Slug:        "hammer",
			Character:   "🔨",
			UnicodeName: "hammer",
			CodePoint:   "1F528",
			Group:       "objects",
			SubGroup:    "tool",
		},

		128297: {
			Slug:        "nut-and-bolt",
			Character:   "🔩",
			UnicodeName: "nut and bolt",
			CodePoint:   "1F529",
			Group:       "objects",
			SubGroup:    "tool",
		},

		128298: {
			Slug:        "e0-6-kitchen-knife",
			Character:   "🔪",
			UnicodeName: "E0.6 kitchen knife",
			CodePoint:   "1F52A",
			Group:       "food-drink",
			SubGroup:    "dishware",
		},

		128299: {
			Slug:        "pistol",
			Character:   "🔫",
			UnicodeName: "pistol",
			CodePoint:   "1F52B",
			Group:       "objects",
			SubGroup:    "tool",
		},

		128300: {
			Slug:        "microscope",
			Character:   "🔬",
			UnicodeName: "microscope",
			CodePoint:   "1F52C",
			Group:       "objects",
			SubGroup:    "science",
		},

		128301: {
			Slug:        "telescope",
			Character:   "🔭",
			UnicodeName: "telescope",
			CodePoint:   "1F52D",
			Group:       "objects",
			SubGroup:    "science",
		},

		128302: {
			Slug:        "crystal-ball",
			Character:   "🔮",
			UnicodeName: "crystal ball",
			CodePoint:   "1F52E",
			Group:       "activities",
			SubGroup:    "game",
		},

		128303: {
			Slug:        "dotted-six-pointed-star",
			Character:   "🔯",
			UnicodeName: "dotted six-pointed star",
			CodePoint:   "1F52F",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		128304: {
			Slug:        "japanese-symbol-for-beginner",
			Character:   "🔰",
			UnicodeName: "Japanese symbol for beginner",
			CodePoint:   "1F530",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		128305: {
			Slug:        "trident-emblem",
			Character:   "🔱",
			UnicodeName: "trident emblem",
			CodePoint:   "1F531",
			Group:       "symbols",
			SubGroup:    "other-symbol",
		},

		128306: {
			Slug:        "black-square-button",
			Character:   "🔲",
			UnicodeName: "black square button",
			CodePoint:   "1F532",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128307: {
			Slug:        "white-square-button",
			Character:   "🔳",
			UnicodeName: "white square button",
			CodePoint:   "1F533",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128308: {
			Slug:        "red-circle",
			Character:   "🔴",
			UnicodeName: "red circle",
			CodePoint:   "1F534",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128309: {
			Slug:        "blue-circle",
			Character:   "🔵",
			UnicodeName: "blue circle",
			CodePoint:   "1F535",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128310: {
			Slug:        "large-orange-diamond",
			Character:   "🔶",
			UnicodeName: "large orange diamond",
			CodePoint:   "1F536",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128311: {
			Slug:        "large-blue-diamond",
			Character:   "🔷",
			UnicodeName: "large blue diamond",
			CodePoint:   "1F537",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128312: {
			Slug:        "small-orange-diamond",
			Character:   "🔸",
			UnicodeName: "small orange diamond",
			CodePoint:   "1F538",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128313: {
			Slug:        "small-blue-diamond",
			Character:   "🔹",
			UnicodeName: "small blue diamond",
			CodePoint:   "1F539",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128314: {
			Slug:        "red-triangle-pointed-up",
			Character:   "🔺",
			UnicodeName: "red triangle pointed up",
			CodePoint:   "1F53A",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128315: {
			Slug:        "red-triangle-pointed-down",
			Character:   "🔻",
			UnicodeName: "red triangle pointed down",
			CodePoint:   "1F53B",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128316: {
			Slug:        "upwards-button",
			Character:   "🔼",
			UnicodeName: "upwards button",
			CodePoint:   "1F53C",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		128317: {
			Slug:        "downwards-button",
			Character:   "🔽",
			UnicodeName: "downwards button",
			CodePoint:   "1F53D",
			Group:       "symbols",
			SubGroup:    "av-symbol",
		},

		128329: {
			Slug:        "om",
			Character:   "🕉️",
			UnicodeName: "om",
			CodePoint:   "1F549 FE0F",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		128330: {
			Slug:        "e0-7-dove",
			Character:   "🕊️",
			UnicodeName: "E0.7 dove",
			CodePoint:   "1F54A FE0F",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		128331: {
			Slug:        "e1-0-kaaba",
			Character:   "🕋",
			UnicodeName: "E1.0 kaaba",
			CodePoint:   "1F54B",
			Group:       "travel-places",
			SubGroup:    "place-religious",
		},

		128332: {
			Slug:        "e1-0-mosque",
			Character:   "🕌",
			UnicodeName: "E1.0 mosque",
			CodePoint:   "1F54C",
			Group:       "travel-places",
			SubGroup:    "place-religious",
		},

		128333: {
			Slug:        "e1-0-synagogue",
			Character:   "🕍",
			UnicodeName: "E1.0 synagogue",
			CodePoint:   "1F54D",
			Group:       "travel-places",
			SubGroup:    "place-religious",
		},

		128334: {
			Slug:        "menorah",
			Character:   "🕎",
			UnicodeName: "menorah",
			CodePoint:   "1F54E",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		128336: {
			Slug:        "one-o-clock",
			Character:   "🕐",
			UnicodeName: "one o’clock",
			CodePoint:   "1F550",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128337: {
			Slug:        "two-o-clock",
			Character:   "🕑",
			UnicodeName: "two o’clock",
			CodePoint:   "1F551",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128338: {
			Slug:        "three-o-clock",
			Character:   "🕒",
			UnicodeName: "three o’clock",
			CodePoint:   "1F552",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128339: {
			Slug:        "four-o-clock",
			Character:   "🕓",
			UnicodeName: "four o’clock",
			CodePoint:   "1F553",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128340: {
			Slug:        "five-o-clock",
			Character:   "🕔",
			UnicodeName: "five o’clock",
			CodePoint:   "1F554",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128341: {
			Slug:        "six-o-clock",
			Character:   "🕕",
			UnicodeName: "six o’clock",
			CodePoint:   "1F555",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128342: {
			Slug:        "seven-o-clock",
			Character:   "🕖",
			UnicodeName: "seven o’clock",
			CodePoint:   "1F556",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128343: {
			Slug:        "eight-o-clock",
			Character:   "🕗",
			UnicodeName: "eight o’clock",
			CodePoint:   "1F557",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128344: {
			Slug:        "nine-o-clock",
			Character:   "🕘",
			UnicodeName: "nine o’clock",
			CodePoint:   "1F558",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128345: {
			Slug:        "ten-o-clock",
			Character:   "🕙",
			UnicodeName: "ten o’clock",
			CodePoint:   "1F559",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128346: {
			Slug:        "eleven-o-clock",
			Character:   "🕚",
			UnicodeName: "eleven o’clock",
			CodePoint:   "1F55A",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128347: {
			Slug:        "twelve-o-clock",
			Character:   "🕛",
			UnicodeName: "twelve o’clock",
			CodePoint:   "1F55B",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128348: {
			Slug:        "one-thirty",
			Character:   "🕜",
			UnicodeName: "one-thirty",
			CodePoint:   "1F55C",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128349: {
			Slug:        "two-thirty",
			Character:   "🕝",
			UnicodeName: "two-thirty",
			CodePoint:   "1F55D",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128350: {
			Slug:        "three-thirty",
			Character:   "🕞",
			UnicodeName: "three-thirty",
			CodePoint:   "1F55E",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128351: {
			Slug:        "four-thirty",
			Character:   "🕟",
			UnicodeName: "four-thirty",
			CodePoint:   "1F55F",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128352: {
			Slug:        "five-thirty",
			Character:   "🕠",
			UnicodeName: "five-thirty",
			CodePoint:   "1F560",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128353: {
			Slug:        "six-thirty",
			Character:   "🕡",
			UnicodeName: "six-thirty",
			CodePoint:   "1F561",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128354: {
			Slug:        "seven-thirty",
			Character:   "🕢",
			UnicodeName: "seven-thirty",
			CodePoint:   "1F562",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128355: {
			Slug:        "eight-thirty",
			Character:   "🕣",
			UnicodeName: "eight-thirty",
			CodePoint:   "1F563",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128356: {
			Slug:        "nine-thirty",
			Character:   "🕤",
			UnicodeName: "nine-thirty",
			CodePoint:   "1F564",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128357: {
			Slug:        "ten-thirty",
			Character:   "🕥",
			UnicodeName: "ten-thirty",
			CodePoint:   "1F565",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128358: {
			Slug:        "eleven-thirty",
			Character:   "🕦",
			UnicodeName: "eleven-thirty",
			CodePoint:   "1F566",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128359: {
			Slug:        "twelve-thirty",
			Character:   "🕧",
			UnicodeName: "twelve-thirty",
			CodePoint:   "1F567",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128367: {
			Slug:        "candle",
			Character:   "🕯️",
			UnicodeName: "candle",
			CodePoint:   "1F56F FE0F",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		128368: {
			Slug:        "mantelpiece-clock",
			Character:   "🕰️",
			UnicodeName: "mantelpiece clock",
			CodePoint:   "1F570 FE0F",
			Group:       "travel-places",
			SubGroup:    "time",
		},

		128371: {
			Slug:        "e0-7-hole",
			Character:   "🕳️",
			UnicodeName: "E0.7 hole",
			CodePoint:   "1F573 FE0F",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128372: {
			Slug:        "e4-0-person-in-suit-levitating-dark-skin-tone",
			Character:   "🕴🏿",
			UnicodeName: "E4.0 person in suit levitating: dark skin tone",
			CodePoint:   "1F574 1F3FF",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		128373: {
			Slug:        "e4-0-woman-detective-dark-skin-tone",
			Character:   "🕵🏿‍♀️",
			UnicodeName: "E4.0 woman detective: dark skin tone",
			CodePoint:   "1F575 1F3FF 200D 2640 FE0F",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		128374: {
			Slug:        "sunglasses",
			Character:   "🕶️",
			UnicodeName: "sunglasses",
			CodePoint:   "1F576 FE0F",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128375: {
			Slug:        "e0-7-spider",
			Character:   "🕷️",
			UnicodeName: "E0.7 spider",
			CodePoint:   "1F577 FE0F",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		128376: {
			Slug:        "e0-7-spider-web",
			Character:   "🕸️",
			UnicodeName: "E0.7 spider web",
			CodePoint:   "1F578 FE0F",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		128377: {
			Slug:        "joystick",
			Character:   "🕹️",
			UnicodeName: "joystick",
			CodePoint:   "1F579 FE0F",
			Group:       "activities",
			SubGroup:    "game",
		},

		128378: {
			Slug:        "e3-0-man-dancing",
			Character:   "🕺",
			UnicodeName: "E3.0 man dancing",
			CodePoint:   "1F57A",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		128391: {
			Slug:        "linked-paperclips",
			Character:   "🖇️",
			UnicodeName: "linked paperclips",
			CodePoint:   "1F587 FE0F",
			Group:       "objects",
			SubGroup:    "office",
		},

		128394: {
			Slug:        "pen",
			Character:   "🖊️",
			UnicodeName: "pen",
			CodePoint:   "1F58A FE0F",
			Group:       "objects",
			SubGroup:    "writing",
		},

		128395: {
			Slug:        "fountain-pen",
			Character:   "🖋️",
			UnicodeName: "fountain pen",
			CodePoint:   "1F58B FE0F",
			Group:       "objects",
			SubGroup:    "writing",
		},

		128396: {
			Slug:        "paintbrush",
			Character:   "🖌️",
			UnicodeName: "paintbrush",
			CodePoint:   "1F58C FE0F",
			Group:       "objects",
			SubGroup:    "writing",
		},

		128397: {
			Slug:        "crayon",
			Character:   "🖍️",
			UnicodeName: "crayon",
			CodePoint:   "1F58D FE0F",
			Group:       "objects",
			SubGroup:    "writing",
		},

		128400: {
			Slug:        "e1-0-hand-with-fingers-splayed-dark-skin-tone",
			Character:   "🖐🏿",
			UnicodeName: "E1.0 hand with fingers splayed: dark skin tone",
			CodePoint:   "1F590 1F3FF",
			Group:       "people-body",
			SubGroup:    "hand-fingers-open",
		},

		128405: {
			Slug:        "e1-0-middle-finger",
			Character:   "🖕",
			UnicodeName: "E1.0 middle finger",
			CodePoint:   "1F595",
			Group:       "people-body",
			SubGroup:    "hand-single-finger",
		},

		128406: {
			Slug:        "e1-0-vulcan-salute",
			Character:   "🖖",
			UnicodeName: "E1.0 vulcan salute",
			CodePoint:   "1F596",
			Group:       "people-body",
			SubGroup:    "hand-fingers-open",
		},

		128420: {
			Slug:        "e3-0-black-heart",
			Character:   "🖤",
			UnicodeName: "E3.0 black heart",
			CodePoint:   "1F5A4",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128421: {
			Slug:        "desktop-computer",
			Character:   "🖥️",
			UnicodeName: "desktop computer",
			CodePoint:   "1F5A5 FE0F",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128424: {
			Slug:        "printer",
			Character:   "🖨️",
			UnicodeName: "printer",
			CodePoint:   "1F5A8 FE0F",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128433: {
			Slug:        "computer-mouse",
			Character:   "🖱️",
			UnicodeName: "computer mouse",
			CodePoint:   "1F5B1 FE0F",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128434: {
			Slug:        "trackball",
			Character:   "🖲️",
			UnicodeName: "trackball",
			CodePoint:   "1F5B2 FE0F",
			Group:       "objects",
			SubGroup:    "computer",
		},

		128444: {
			Slug:        "framed-picture",
			Character:   "🖼️",
			UnicodeName: "framed picture",
			CodePoint:   "1F5BC FE0F",
			Group:       "activities",
			SubGroup:    "arts-crafts",
		},

		128450: {
			Slug:        "card-index-dividers",
			Character:   "🗂️",
			UnicodeName: "card index dividers",
			CodePoint:   "1F5C2 FE0F",
			Group:       "objects",
			SubGroup:    "office",
		},

		128451: {
			Slug:        "card-file-box",
			Character:   "🗃️",
			UnicodeName: "card file box",
			CodePoint:   "1F5C3 FE0F",
			Group:       "objects",
			SubGroup:    "office",
		},

		128452: {
			Slug:        "file-cabinet",
			Character:   "🗄️",
			UnicodeName: "file cabinet",
			CodePoint:   "1F5C4 FE0F",
			Group:       "objects",
			SubGroup:    "office",
		},

		128465: {
			Slug:        "wastebasket",
			Character:   "🗑️",
			UnicodeName: "wastebasket",
			CodePoint:   "1F5D1 FE0F",
			Group:       "objects",
			SubGroup:    "office",
		},

		128466: {
			Slug:        "spiral-notepad",
			Character:   "🗒️",
			UnicodeName: "spiral notepad",
			CodePoint:   "1F5D2 FE0F",
			Group:       "objects",
			SubGroup:    "office",
		},

		128467: {
			Slug:        "spiral-calendar",
			Character:   "🗓️",
			UnicodeName: "spiral calendar",
			CodePoint:   "1F5D3 FE0F",
			Group:       "objects",
			SubGroup:    "office",
		},

		128476: {
			Slug:        "clamp",
			Character:   "🗜️",
			UnicodeName: "clamp",
			CodePoint:   "1F5DC FE0F",
			Group:       "objects",
			SubGroup:    "tool",
		},

		128477: {
			Slug:        "old-key",
			Character:   "🗝️",
			UnicodeName: "old key",
			CodePoint:   "1F5DD FE0F",
			Group:       "objects",
			SubGroup:    "lock",
		},

		128478: {
			Slug:        "rolled-up-newspaper",
			Character:   "🗞️",
			UnicodeName: "rolled-up newspaper",
			CodePoint:   "1F5DE FE0F",
			Group:       "objects",
			SubGroup:    "book-paper",
		},

		128481: {
			Slug:        "dagger",
			Character:   "🗡️",
			UnicodeName: "dagger",
			CodePoint:   "1F5E1 FE0F",
			Group:       "objects",
			SubGroup:    "tool",
		},

		128483: {
			Slug:        "e0-7-speaking-head",
			Character:   "🗣️",
			UnicodeName: "E0.7 speaking head",
			CodePoint:   "1F5E3 FE0F",
			Group:       "people-body",
			SubGroup:    "person-symbol",
		},

		128488: {
			Slug:        "e2-0-left-speech-bubble",
			Character:   "🗨️",
			UnicodeName: "E2.0 left speech bubble",
			CodePoint:   "1F5E8 FE0F",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128495: {
			Slug:        "e0-7-right-anger-bubble",
			Character:   "🗯️",
			UnicodeName: "E0.7 right anger bubble",
			CodePoint:   "1F5EF FE0F",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		128499: {
			Slug:        "ballot-box-with-ballot",
			Character:   "🗳️",
			UnicodeName: "ballot box with ballot",
			CodePoint:   "1F5F3 FE0F",
			Group:       "objects",
			SubGroup:    "mail",
		},

		128506: {
			Slug:        "e0-7-world-map",
			Character:   "🗺️",
			UnicodeName: "E0.7 world map",
			CodePoint:   "1F5FA FE0F",
			Group:       "travel-places",
			SubGroup:    "place-map",
		},

		128507: {
			Slug:        "e0-6-mount-fuji",
			Character:   "🗻",
			UnicodeName: "E0.6 mount fuji",
			CodePoint:   "1F5FB",
			Group:       "travel-places",
			SubGroup:    "place-geographic",
		},

		128508: {
			Slug:        "e0-6-tokyo-tower",
			Character:   "🗼",
			UnicodeName: "E0.6 Tokyo tower",
			CodePoint:   "1F5FC",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		128509: {
			Slug:        "e0-6-statue-of-liberty",
			Character:   "🗽",
			UnicodeName: "E0.6 Statue of Liberty",
			CodePoint:   "1F5FD",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		128510: {
			Slug:        "e0-6-map-of-japan",
			Character:   "🗾",
			UnicodeName: "E0.6 map of Japan",
			CodePoint:   "1F5FE",
			Group:       "travel-places",
			SubGroup:    "place-map",
		},

		128511: {
			Slug:        "moai",
			Character:   "🗿",
			UnicodeName: "moai",
			CodePoint:   "1F5FF",
			Group:       "objects",
			SubGroup:    "other-object",
		},

		128512: {
			Slug:        "e1-0-grinning-face",
			Character:   "😀",
			UnicodeName: "E1.0 grinning face",
			CodePoint:   "1F600",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128513: {
			Slug:        "e0-6-beaming-face-with-smiling-eyes",
			Character:   "😁",
			UnicodeName: "E0.6 beaming face with smiling eyes",
			CodePoint:   "1F601",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128514: {
			Slug:        "e0-6-face-with-tears-of-joy",
			Character:   "😂",
			UnicodeName: "E0.6 face with tears of joy",
			CodePoint:   "1F602",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128515: {
			Slug:        "e0-6-grinning-face-with-big-eyes",
			Character:   "😃",
			UnicodeName: "E0.6 grinning face with big eyes",
			CodePoint:   "1F603",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128516: {
			Slug:        "e0-6-grinning-face-with-smiling-eyes",
			Character:   "😄",
			UnicodeName: "E0.6 grinning face with smiling eyes",
			CodePoint:   "1F604",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128517: {
			Slug:        "e0-6-grinning-face-with-sweat",
			Character:   "😅",
			UnicodeName: "E0.6 grinning face with sweat",
			CodePoint:   "1F605",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128518: {
			Slug:        "e0-6-grinning-squinting-face",
			Character:   "😆",
			UnicodeName: "E0.6 grinning squinting face",
			CodePoint:   "1F606",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128519: {
			Slug:        "e1-0-smiling-face-with-halo",
			Character:   "😇",
			UnicodeName: "E1.0 smiling face with halo",
			CodePoint:   "1F607",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128520: {
			Slug:        "e1-0-smiling-face-with-horns",
			Character:   "😈",
			UnicodeName: "E1.0 smiling face with horns",
			CodePoint:   "1F608",
			Group:       "smileys-emotion",
			SubGroup:    "face-negative",
		},

		128521: {
			Slug:        "e0-6-winking-face",
			Character:   "😉",
			UnicodeName: "E0.6 winking face",
			CodePoint:   "1F609",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128522: {
			Slug:        "e0-6-smiling-face-with-smiling-eyes",
			Character:   "😊",
			UnicodeName: "E0.6 smiling face with smiling eyes",
			CodePoint:   "1F60A",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128523: {
			Slug:        "e0-6-face-savoring-food",
			Character:   "😋",
			UnicodeName: "E0.6 face savoring food",
			CodePoint:   "1F60B",
			Group:       "smileys-emotion",
			SubGroup:    "face-tongue",
		},

		128524: {
			Slug:        "e0-6-relieved-face",
			Character:   "😌",
			UnicodeName: "E0.6 relieved face",
			CodePoint:   "1F60C",
			Group:       "smileys-emotion",
			SubGroup:    "face-sleepy",
		},

		128525: {
			Slug:        "e0-6-smiling-face-with-heart-eyes",
			Character:   "😍",
			UnicodeName: "E0.6 smiling face with heart-eyes",
			CodePoint:   "1F60D",
			Group:       "smileys-emotion",
			SubGroup:    "face-affection",
		},

		128526: {
			Slug:        "e1-0-smiling-face-with-sunglasses",
			Character:   "😎",
			UnicodeName: "E1.0 smiling face with sunglasses",
			CodePoint:   "1F60E",
			Group:       "smileys-emotion",
			SubGroup:    "face-glasses",
		},

		128527: {
			Slug:        "e0-6-smirking-face",
			Character:   "😏",
			UnicodeName: "E0.6 smirking face",
			CodePoint:   "1F60F",
			Group:       "smileys-emotion",
			SubGroup:    "face-neutral-skeptical",
		},

		128528: {
			Slug:        "e0-7-neutral-face",
			Character:   "😐",
			UnicodeName: "E0.7 neutral face",
			CodePoint:   "1F610",
			Group:       "smileys-emotion",
			SubGroup:    "face-neutral-skeptical",
		},

		128529: {
			Slug:        "e1-0-expressionless-face",
			Character:   "😑",
			UnicodeName: "E1.0 expressionless face",
			CodePoint:   "1F611",
			Group:       "smileys-emotion",
			SubGroup:    "face-neutral-skeptical",
		},

		128530: {
			Slug:        "e0-6-unamused-face",
			Character:   "😒",
			UnicodeName: "E0.6 unamused face",
			CodePoint:   "1F612",
			Group:       "smileys-emotion",
			SubGroup:    "face-neutral-skeptical",
		},

		128531: {
			Slug:        "e0-6-downcast-face-with-sweat",
			Character:   "😓",
			UnicodeName: "E0.6 downcast face with sweat",
			CodePoint:   "1F613",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128532: {
			Slug:        "e0-6-pensive-face",
			Character:   "😔",
			UnicodeName: "E0.6 pensive face",
			CodePoint:   "1F614",
			Group:       "smileys-emotion",
			SubGroup:    "face-sleepy",
		},

		128533: {
			Slug:        "e1-0-confused-face",
			Character:   "😕",
			UnicodeName: "E1.0 confused face",
			CodePoint:   "1F615",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128534: {
			Slug:        "e0-6-confounded-face",
			Character:   "😖",
			UnicodeName: "E0.6 confounded face",
			CodePoint:   "1F616",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128535: {
			Slug:        "e1-0-kissing-face",
			Character:   "😗",
			UnicodeName: "E1.0 kissing face",
			CodePoint:   "1F617",
			Group:       "smileys-emotion",
			SubGroup:    "face-affection",
		},

		128536: {
			Slug:        "e0-6-face-blowing-a-kiss",
			Character:   "😘",
			UnicodeName: "E0.6 face blowing a kiss",
			CodePoint:   "1F618",
			Group:       "smileys-emotion",
			SubGroup:    "face-affection",
		},

		128537: {
			Slug:        "e1-0-kissing-face-with-smiling-eyes",
			Character:   "😙",
			UnicodeName: "E1.0 kissing face with smiling eyes",
			CodePoint:   "1F619",
			Group:       "smileys-emotion",
			SubGroup:    "face-affection",
		},

		128538: {
			Slug:        "e0-6-kissing-face-with-closed-eyes",
			Character:   "😚",
			UnicodeName: "E0.6 kissing face with closed eyes",
			CodePoint:   "1F61A",
			Group:       "smileys-emotion",
			SubGroup:    "face-affection",
		},

		128539: {
			Slug:        "e1-0-face-with-tongue",
			Character:   "😛",
			UnicodeName: "E1.0 face with tongue",
			CodePoint:   "1F61B",
			Group:       "smileys-emotion",
			SubGroup:    "face-tongue",
		},

		128540: {
			Slug:        "e0-6-winking-face-with-tongue",
			Character:   "😜",
			UnicodeName: "E0.6 winking face with tongue",
			CodePoint:   "1F61C",
			Group:       "smileys-emotion",
			SubGroup:    "face-tongue",
		},

		128541: {
			Slug:        "e0-6-squinting-face-with-tongue",
			Character:   "😝",
			UnicodeName: "E0.6 squinting face with tongue",
			CodePoint:   "1F61D",
			Group:       "smileys-emotion",
			SubGroup:    "face-tongue",
		},

		128542: {
			Slug:        "e0-6-disappointed-face",
			Character:   "😞",
			UnicodeName: "E0.6 disappointed face",
			CodePoint:   "1F61E",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128543: {
			Slug:        "e1-0-worried-face",
			Character:   "😟",
			UnicodeName: "E1.0 worried face",
			CodePoint:   "1F61F",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128544: {
			Slug:        "e0-6-angry-face",
			Character:   "😠",
			UnicodeName: "E0.6 angry face",
			CodePoint:   "1F620",
			Group:       "smileys-emotion",
			SubGroup:    "face-negative",
		},

		128545: {
			Slug:        "e0-6-pouting-face",
			Character:   "😡",
			UnicodeName: "E0.6 pouting face",
			CodePoint:   "1F621",
			Group:       "smileys-emotion",
			SubGroup:    "face-negative",
		},

		128546: {
			Slug:        "e0-6-crying-face",
			Character:   "😢",
			UnicodeName: "E0.6 crying face",
			CodePoint:   "1F622",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128547: {
			Slug:        "e0-6-persevering-face",
			Character:   "😣",
			UnicodeName: "E0.6 persevering face",
			CodePoint:   "1F623",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128548: {
			Slug:        "e0-6-face-with-steam-from-nose",
			Character:   "😤",
			UnicodeName: "E0.6 face with steam from nose",
			CodePoint:   "1F624",
			Group:       "smileys-emotion",
			SubGroup:    "face-negative",
		},

		128549: {
			Slug:        "e0-6-sad-but-relieved-face",
			Character:   "😥",
			UnicodeName: "E0.6 sad but relieved face",
			CodePoint:   "1F625",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128550: {
			Slug:        "e1-0-frowning-face-with-open-mouth",
			Character:   "😦",
			UnicodeName: "E1.0 frowning face with open mouth",
			CodePoint:   "1F626",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128551: {
			Slug:        "e1-0-anguished-face",
			Character:   "😧",
			UnicodeName: "E1.0 anguished face",
			CodePoint:   "1F627",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128552: {
			Slug:        "e0-6-fearful-face",
			Character:   "😨",
			UnicodeName: "E0.6 fearful face",
			CodePoint:   "1F628",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128553: {
			Slug:        "e0-6-weary-face",
			Character:   "😩",
			UnicodeName: "E0.6 weary face",
			CodePoint:   "1F629",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128554: {
			Slug:        "e0-6-sleepy-face",
			Character:   "😪",
			UnicodeName: "E0.6 sleepy face",
			CodePoint:   "1F62A",
			Group:       "smileys-emotion",
			SubGroup:    "face-sleepy",
		},

		128555: {
			Slug:        "e0-6-tired-face",
			Character:   "😫",
			UnicodeName: "E0.6 tired face",
			CodePoint:   "1F62B",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128556: {
			Slug:        "e1-0-grimacing-face",
			Character:   "😬",
			UnicodeName: "E1.0 grimacing face",
			CodePoint:   "1F62C",
			Group:       "smileys-emotion",
			SubGroup:    "face-neutral-skeptical",
		},

		128557: {
			Slug:        "e0-6-loudly-crying-face",
			Character:   "😭",
			UnicodeName: "E0.6 loudly crying face",
			CodePoint:   "1F62D",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128558: {
			Slug:        "e1-0-face-with-open-mouth",
			Character:   "😮",
			UnicodeName: "E1.0 face with open mouth",
			CodePoint:   "1F62E",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128559: {
			Slug:        "e1-0-hushed-face",
			Character:   "😯",
			UnicodeName: "E1.0 hushed face",
			CodePoint:   "1F62F",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128560: {
			Slug:        "e0-6-anxious-face-with-sweat",
			Character:   "😰",
			UnicodeName: "E0.6 anxious face with sweat",
			CodePoint:   "1F630",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128561: {
			Slug:        "e0-6-face-screaming-in-fear",
			Character:   "😱",
			UnicodeName: "E0.6 face screaming in fear",
			CodePoint:   "1F631",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128562: {
			Slug:        "e0-6-astonished-face",
			Character:   "😲",
			UnicodeName: "E0.6 astonished face",
			CodePoint:   "1F632",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128563: {
			Slug:        "e0-6-flushed-face",
			Character:   "😳",
			UnicodeName: "E0.6 flushed face",
			CodePoint:   "1F633",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128564: {
			Slug:        "e1-0-sleeping-face",
			Character:   "😴",
			UnicodeName: "E1.0 sleeping face",
			CodePoint:   "1F634",
			Group:       "smileys-emotion",
			SubGroup:    "face-sleepy",
		},

		128565: {
			Slug:        "e0-6-dizzy-face",
			Character:   "😵",
			UnicodeName: "E0.6 dizzy face",
			CodePoint:   "1F635",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		128566: {
			Slug:        "e1-0-face-without-mouth",
			Character:   "😶",
			UnicodeName: "E1.0 face without mouth",
			CodePoint:   "1F636",
			Group:       "smileys-emotion",
			SubGroup:    "face-neutral-skeptical",
		},

		128567: {
			Slug:        "e0-6-face-with-medical-mask",
			Character:   "😷",
			UnicodeName: "E0.6 face with medical mask",
			CodePoint:   "1F637",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		128568: {
			Slug:        "e0-6-grinning-cat-with-smiling-eyes",
			Character:   "😸",
			UnicodeName: "E0.6 grinning cat with smiling eyes",
			CodePoint:   "1F638",
			Group:       "smileys-emotion",
			SubGroup:    "cat-face",
		},

		128569: {
			Slug:        "e0-6-cat-with-tears-of-joy",
			Character:   "😹",
			UnicodeName: "E0.6 cat with tears of joy",
			CodePoint:   "1F639",
			Group:       "smileys-emotion",
			SubGroup:    "cat-face",
		},

		128570: {
			Slug:        "e0-6-grinning-cat",
			Character:   "😺",
			UnicodeName: "E0.6 grinning cat",
			CodePoint:   "1F63A",
			Group:       "smileys-emotion",
			SubGroup:    "cat-face",
		},

		128571: {
			Slug:        "e0-6-smiling-cat-with-heart-eyes",
			Character:   "😻",
			UnicodeName: "E0.6 smiling cat with heart-eyes",
			CodePoint:   "1F63B",
			Group:       "smileys-emotion",
			SubGroup:    "cat-face",
		},

		128572: {
			Slug:        "e0-6-cat-with-wry-smile",
			Character:   "😼",
			UnicodeName: "E0.6 cat with wry smile",
			CodePoint:   "1F63C",
			Group:       "smileys-emotion",
			SubGroup:    "cat-face",
		},

		128573: {
			Slug:        "e0-6-kissing-cat",
			Character:   "😽",
			UnicodeName: "E0.6 kissing cat",
			CodePoint:   "1F63D",
			Group:       "smileys-emotion",
			SubGroup:    "cat-face",
		},

		128574: {
			Slug:        "e0-6-pouting-cat",
			Character:   "😾",
			UnicodeName: "E0.6 pouting cat",
			CodePoint:   "1F63E",
			Group:       "smileys-emotion",
			SubGroup:    "cat-face",
		},

		128575: {
			Slug:        "e0-6-crying-cat",
			Character:   "😿",
			UnicodeName: "E0.6 crying cat",
			CodePoint:   "1F63F",
			Group:       "smileys-emotion",
			SubGroup:    "cat-face",
		},

		128576: {
			Slug:        "e0-6-weary-cat",
			Character:   "🙀",
			UnicodeName: "E0.6 weary cat",
			CodePoint:   "1F640",
			Group:       "smileys-emotion",
			SubGroup:    "cat-face",
		},

		128577: {
			Slug:        "e1-0-slightly-frowning-face",
			Character:   "🙁",
			UnicodeName: "E1.0 slightly frowning face",
			CodePoint:   "1F641",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		128578: {
			Slug:        "e1-0-slightly-smiling-face",
			Character:   "🙂",
			UnicodeName: "E1.0 slightly smiling face",
			CodePoint:   "1F642",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128579: {
			Slug:        "e1-0-upside-down-face",
			Character:   "🙃",
			UnicodeName: "E1.0 upside-down face",
			CodePoint:   "1F643",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		128580: {
			Slug:        "e1-0-face-with-rolling-eyes",
			Character:   "🙄",
			UnicodeName: "E1.0 face with rolling eyes",
			CodePoint:   "1F644",
			Group:       "smileys-emotion",
			SubGroup:    "face-neutral-skeptical",
		},

		128581: {
			Slug:        "e0-6-person-gesturing-no",
			Character:   "🙅",
			UnicodeName: "E0.6 person gesturing NO",
			CodePoint:   "1F645",
			Group:       "people-body",
			SubGroup:    "person-gesture",
		},

		128582: {
			Slug:        "e0-6-person-gesturing-ok",
			Character:   "🙆",
			UnicodeName: "E0.6 person gesturing OK",
			CodePoint:   "1F646",
			Group:       "people-body",
			SubGroup:    "person-gesture",
		},

		128583: {
			Slug:        "e0-6-person-bowing",
			Character:   "🙇",
			UnicodeName: "E0.6 person bowing",
			CodePoint:   "1F647",
			Group:       "people-body",
			SubGroup:    "person-gesture",
		},

		128584: {
			Slug:        "e0-6-see-no-evil-monkey",
			Character:   "🙈",
			UnicodeName: "E0.6 see-no-evil monkey",
			CodePoint:   "1F648",
			Group:       "smileys-emotion",
			SubGroup:    "monkey-face",
		},

		128585: {
			Slug:        "e0-6-hear-no-evil-monkey",
			Character:   "🙉",
			UnicodeName: "E0.6 hear-no-evil monkey",
			CodePoint:   "1F649",
			Group:       "smileys-emotion",
			SubGroup:    "monkey-face",
		},

		128586: {
			Slug:        "e0-6-speak-no-evil-monkey",
			Character:   "🙊",
			UnicodeName: "E0.6 speak-no-evil monkey",
			CodePoint:   "1F64A",
			Group:       "smileys-emotion",
			SubGroup:    "monkey-face",
		},

		128587: {
			Slug:        "e0-6-person-raising-hand",
			Character:   "🙋",
			UnicodeName: "E0.6 person raising hand",
			CodePoint:   "1F64B",
			Group:       "people-body",
			SubGroup:    "person-gesture",
		},

		128588: {
			Slug:        "e0-6-raising-hands",
			Character:   "🙌",
			UnicodeName: "E0.6 raising hands",
			CodePoint:   "1F64C",
			Group:       "people-body",
			SubGroup:    "hands",
		},

		128589: {
			Slug:        "e0-6-person-frowning",
			Character:   "🙍",
			UnicodeName: "E0.6 person frowning",
			CodePoint:   "1F64D",
			Group:       "people-body",
			SubGroup:    "person-gesture",
		},

		128590: {
			Slug:        "e0-6-person-pouting",
			Character:   "🙎",
			UnicodeName: "E0.6 person pouting",
			CodePoint:   "1F64E",
			Group:       "people-body",
			SubGroup:    "person-gesture",
		},

		128591: {
			Slug:        "e0-6-folded-hands",
			Character:   "🙏",
			UnicodeName: "E0.6 folded hands",
			CodePoint:   "1F64F",
			Group:       "people-body",
			SubGroup:    "hands",
		},

		128640: {
			Slug:        "e0-6-rocket",
			Character:   "🚀",
			UnicodeName: "E0.6 rocket",
			CodePoint:   "1F680",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128641: {
			Slug:        "e1-0-helicopter",
			Character:   "🚁",
			UnicodeName: "E1.0 helicopter",
			CodePoint:   "1F681",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128642: {
			Slug:        "e1-0-locomotive",
			Character:   "🚂",
			UnicodeName: "E1.0 locomotive",
			CodePoint:   "1F682",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128643: {
			Slug:        "e0-6-railway-car",
			Character:   "🚃",
			UnicodeName: "E0.6 railway car",
			CodePoint:   "1F683",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128644: {
			Slug:        "e0-6-high-speed-train",
			Character:   "🚄",
			UnicodeName: "E0.6 high-speed train",
			CodePoint:   "1F684",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128645: {
			Slug:        "e0-6-bullet-train",
			Character:   "🚅",
			UnicodeName: "E0.6 bullet train",
			CodePoint:   "1F685",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128646: {
			Slug:        "e1-0-train",
			Character:   "🚆",
			UnicodeName: "E1.0 train",
			CodePoint:   "1F686",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128647: {
			Slug:        "e0-6-metro",
			Character:   "🚇",
			UnicodeName: "E0.6 metro",
			CodePoint:   "1F687",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128648: {
			Slug:        "e1-0-light-rail",
			Character:   "🚈",
			UnicodeName: "E1.0 light rail",
			CodePoint:   "1F688",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128649: {
			Slug:        "e0-6-station",
			Character:   "🚉",
			UnicodeName: "E0.6 station",
			CodePoint:   "1F689",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128650: {
			Slug:        "e1-0-tram",
			Character:   "🚊",
			UnicodeName: "E1.0 tram",
			CodePoint:   "1F68A",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128651: {
			Slug:        "e1-0-tram-car",
			Character:   "🚋",
			UnicodeName: "E1.0 tram car",
			CodePoint:   "1F68B",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128652: {
			Slug:        "e0-6-bus",
			Character:   "🚌",
			UnicodeName: "E0.6 bus",
			CodePoint:   "1F68C",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128653: {
			Slug:        "e0-7-oncoming-bus",
			Character:   "🚍",
			UnicodeName: "E0.7 oncoming bus",
			CodePoint:   "1F68D",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128654: {
			Slug:        "e1-0-trolleybus",
			Character:   "🚎",
			UnicodeName: "E1.0 trolleybus",
			CodePoint:   "1F68E",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128655: {
			Slug:        "e0-6-bus-stop",
			Character:   "🚏",
			UnicodeName: "E0.6 bus stop",
			CodePoint:   "1F68F",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128656: {
			Slug:        "e1-0-minibus",
			Character:   "🚐",
			UnicodeName: "E1.0 minibus",
			CodePoint:   "1F690",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128657: {
			Slug:        "e0-6-ambulance",
			Character:   "🚑",
			UnicodeName: "E0.6 ambulance",
			CodePoint:   "1F691",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128658: {
			Slug:        "e0-6-fire-engine",
			Character:   "🚒",
			UnicodeName: "E0.6 fire engine",
			CodePoint:   "1F692",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128659: {
			Slug:        "e0-6-police-car",
			Character:   "🚓",
			UnicodeName: "E0.6 police car",
			CodePoint:   "1F693",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128660: {
			Slug:        "e0-7-oncoming-police-car",
			Character:   "🚔",
			UnicodeName: "E0.7 oncoming police car",
			CodePoint:   "1F694",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128661: {
			Slug:        "e0-6-taxi",
			Character:   "🚕",
			UnicodeName: "E0.6 taxi",
			CodePoint:   "1F695",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128662: {
			Slug:        "e1-0-oncoming-taxi",
			Character:   "🚖",
			UnicodeName: "E1.0 oncoming taxi",
			CodePoint:   "1F696",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128663: {
			Slug:        "e0-6-automobile",
			Character:   "🚗",
			UnicodeName: "E0.6 automobile",
			CodePoint:   "1F697",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128664: {
			Slug:        "e0-7-oncoming-automobile",
			Character:   "🚘",
			UnicodeName: "E0.7 oncoming automobile",
			CodePoint:   "1F698",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128665: {
			Slug:        "e0-6-sport-utility-vehicle",
			Character:   "🚙",
			UnicodeName: "E0.6 sport utility vehicle",
			CodePoint:   "1F699",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128666: {
			Slug:        "e0-6-delivery-truck",
			Character:   "🚚",
			UnicodeName: "E0.6 delivery truck",
			CodePoint:   "1F69A",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128667: {
			Slug:        "e1-0-articulated-lorry",
			Character:   "🚛",
			UnicodeName: "E1.0 articulated lorry",
			CodePoint:   "1F69B",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128668: {
			Slug:        "e1-0-tractor",
			Character:   "🚜",
			UnicodeName: "E1.0 tractor",
			CodePoint:   "1F69C",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128669: {
			Slug:        "e1-0-monorail",
			Character:   "🚝",
			UnicodeName: "E1.0 monorail",
			CodePoint:   "1F69D",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128670: {
			Slug:        "e1-0-mountain-railway",
			Character:   "🚞",
			UnicodeName: "E1.0 mountain railway",
			CodePoint:   "1F69E",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128671: {
			Slug:        "e1-0-suspension-railway",
			Character:   "🚟",
			UnicodeName: "E1.0 suspension railway",
			CodePoint:   "1F69F",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128672: {
			Slug:        "e1-0-mountain-cableway",
			Character:   "🚠",
			UnicodeName: "E1.0 mountain cableway",
			CodePoint:   "1F6A0",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128673: {
			Slug:        "e1-0-aerial-tramway",
			Character:   "🚡",
			UnicodeName: "E1.0 aerial tramway",
			CodePoint:   "1F6A1",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128674: {
			Slug:        "e0-6-ship",
			Character:   "🚢",
			UnicodeName: "E0.6 ship",
			CodePoint:   "1F6A2",
			Group:       "travel-places",
			SubGroup:    "transport-water",
		},

		128675: {
			Slug:        "e1-0-person-rowing-boat",
			Character:   "🚣",
			UnicodeName: "E1.0 person rowing boat",
			CodePoint:   "1F6A3",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		128676: {
			Slug:        "e0-6-speedboat",
			Character:   "🚤",
			UnicodeName: "E0.6 speedboat",
			CodePoint:   "1F6A4",
			Group:       "travel-places",
			SubGroup:    "transport-water",
		},

		128677: {
			Slug:        "e0-6-horizontal-traffic-light",
			Character:   "🚥",
			UnicodeName: "E0.6 horizontal traffic light",
			CodePoint:   "1F6A5",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128678: {
			Slug:        "e1-0-vertical-traffic-light",
			Character:   "🚦",
			UnicodeName: "E1.0 vertical traffic light",
			CodePoint:   "1F6A6",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128679: {
			Slug:        "e0-6-construction",
			Character:   "🚧",
			UnicodeName: "E0.6 construction",
			CodePoint:   "1F6A7",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128680: {
			Slug:        "e0-6-police-car-light",
			Character:   "🚨",
			UnicodeName: "E0.6 police car light",
			CodePoint:   "1F6A8",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128681: {
			Slug:        "triangular-flag",
			Character:   "🚩",
			UnicodeName: "triangular flag",
			CodePoint:   "1F6A9",
			Group:       "flags",
			SubGroup:    "flag",
		},

		128682: {
			Slug:        "door",
			Character:   "🚪",
			UnicodeName: "door",
			CodePoint:   "1F6AA",
			Group:       "objects",
			SubGroup:    "household",
		},

		128683: {
			Slug:        "prohibited",
			Character:   "🚫",
			UnicodeName: "prohibited",
			CodePoint:   "1F6AB",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		128684: {
			Slug:        "cigarette",
			Character:   "🚬",
			UnicodeName: "cigarette",
			CodePoint:   "1F6AC",
			Group:       "objects",
			SubGroup:    "other-object",
		},

		128685: {
			Slug:        "no-smoking",
			Character:   "🚭",
			UnicodeName: "no smoking",
			CodePoint:   "1F6AD",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		128686: {
			Slug:        "litter-in-bin-sign",
			Character:   "🚮",
			UnicodeName: "litter in bin sign",
			CodePoint:   "1F6AE",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128687: {
			Slug:        "no-littering",
			Character:   "🚯",
			UnicodeName: "no littering",
			CodePoint:   "1F6AF",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		128688: {
			Slug:        "potable-water",
			Character:   "🚰",
			UnicodeName: "potable water",
			CodePoint:   "1F6B0",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128689: {
			Slug:        "non-potable-water",
			Character:   "🚱",
			UnicodeName: "non-potable water",
			CodePoint:   "1F6B1",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		128690: {
			Slug:        "e0-6-bicycle",
			Character:   "🚲",
			UnicodeName: "E0.6 bicycle",
			CodePoint:   "1F6B2",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128691: {
			Slug:        "no-bicycles",
			Character:   "🚳",
			UnicodeName: "no bicycles",
			CodePoint:   "1F6B3",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		128692: {
			Slug:        "e1-0-person-biking",
			Character:   "🚴",
			UnicodeName: "E1.0 person biking",
			CodePoint:   "1F6B4",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		128693: {
			Slug:        "e1-0-person-mountain-biking",
			Character:   "🚵",
			UnicodeName: "E1.0 person mountain biking",
			CodePoint:   "1F6B5",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		128694: {
			Slug:        "e0-6-person-walking",
			Character:   "🚶",
			UnicodeName: "E0.6 person walking",
			CodePoint:   "1F6B6",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		128695: {
			Slug:        "no-pedestrians",
			Character:   "🚷",
			UnicodeName: "no pedestrians",
			CodePoint:   "1F6B7",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		128696: {
			Slug:        "children-crossing",
			Character:   "🚸",
			UnicodeName: "children crossing",
			CodePoint:   "1F6B8",
			Group:       "symbols",
			SubGroup:    "warning",
		},

		128697: {
			Slug:        "men-s-room",
			Character:   "🚹",
			UnicodeName: "men’s room",
			CodePoint:   "1F6B9",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128698: {
			Slug:        "women-s-room",
			Character:   "🚺",
			UnicodeName: "women’s room",
			CodePoint:   "1F6BA",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128699: {
			Slug:        "restroom",
			Character:   "🚻",
			UnicodeName: "restroom",
			CodePoint:   "1F6BB",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128700: {
			Slug:        "baby-symbol",
			Character:   "🚼",
			UnicodeName: "baby symbol",
			CodePoint:   "1F6BC",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128701: {
			Slug:        "toilet",
			Character:   "🚽",
			UnicodeName: "toilet",
			CodePoint:   "1F6BD",
			Group:       "objects",
			SubGroup:    "household",
		},

		128702: {
			Slug:        "water-closet",
			Character:   "🚾",
			UnicodeName: "water closet",
			CodePoint:   "1F6BE",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128703: {
			Slug:        "shower",
			Character:   "🚿",
			UnicodeName: "shower",
			CodePoint:   "1F6BF",
			Group:       "objects",
			SubGroup:    "household",
		},

		128704: {
			Slug:        "e0-6-person-taking-bath",
			Character:   "🛀",
			UnicodeName: "E0.6 person taking bath",
			CodePoint:   "1F6C0",
			Group:       "people-body",
			SubGroup:    "person-resting",
		},

		128705: {
			Slug:        "bathtub",
			Character:   "🛁",
			UnicodeName: "bathtub",
			CodePoint:   "1F6C1",
			Group:       "objects",
			SubGroup:    "household",
		},

		128706: {
			Slug:        "passport-control",
			Character:   "🛂",
			UnicodeName: "passport control",
			CodePoint:   "1F6C2",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128707: {
			Slug:        "customs",
			Character:   "🛃",
			UnicodeName: "customs",
			CodePoint:   "1F6C3",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128708: {
			Slug:        "baggage-claim",
			Character:   "🛄",
			UnicodeName: "baggage claim",
			CodePoint:   "1F6C4",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128709: {
			Slug:        "left-luggage",
			Character:   "🛅",
			UnicodeName: "left luggage",
			CodePoint:   "1F6C5",
			Group:       "symbols",
			SubGroup:    "transport-sign",
		},

		128715: {
			Slug:        "couch-and-lamp",
			Character:   "🛋️",
			UnicodeName: "couch and lamp",
			CodePoint:   "1F6CB FE0F",
			Group:       "objects",
			SubGroup:    "household",
		},

		128716: {
			Slug:        "e1-0-person-in-bed",
			Character:   "🛌",
			UnicodeName: "E1.0 person in bed",
			CodePoint:   "1F6CC",
			Group:       "people-body",
			SubGroup:    "person-resting",
		},

		128717: {
			Slug:        "shopping-bags",
			Character:   "🛍️",
			UnicodeName: "shopping bags",
			CodePoint:   "1F6CD FE0F",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		128718: {
			Slug:        "e0-7-bellhop-bell",
			Character:   "🛎️",
			UnicodeName: "E0.7 bellhop bell",
			CodePoint:   "1F6CE FE0F",
			Group:       "travel-places",
			SubGroup:    "hotel",
		},

		128719: {
			Slug:        "bed",
			Character:   "🛏️",
			UnicodeName: "bed",
			CodePoint:   "1F6CF FE0F",
			Group:       "objects",
			SubGroup:    "household",
		},

		128720: {
			Slug:        "place-of-worship",
			Character:   "🛐",
			UnicodeName: "place of worship",
			CodePoint:   "1F6D0",
			Group:       "symbols",
			SubGroup:    "religion",
		},

		128721: {
			Slug:        "e3-0-stop-sign",
			Character:   "🛑",
			UnicodeName: "E3.0 stop sign",
			CodePoint:   "1F6D1",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128722: {
			Slug:        "shopping-cart",
			Character:   "🛒",
			UnicodeName: "shopping cart",
			CodePoint:   "1F6D2",
			Group:       "objects",
			SubGroup:    "household",
		},

		128725: {
			Slug:        "e12-0-hindu-temple",
			Character:   "🛕",
			UnicodeName: "E12.0 hindu temple",
			CodePoint:   "1F6D5",
			Group:       "travel-places",
			SubGroup:    "place-religious",
		},

		128726: {
			Slug:        "e13-0-hut",
			Character:   "🛖",
			UnicodeName: "E13.0 hut",
			CodePoint:   "1F6D6",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		128727: {
			Slug:        "elevator",
			Character:   "🛗",
			UnicodeName: "elevator",
			CodePoint:   "1F6D7",
			Group:       "objects",
			SubGroup:    "household",
		},

		128736: {
			Slug:        "hammer-and-wrench",
			Character:   "🛠️",
			UnicodeName: "hammer and wrench",
			CodePoint:   "1F6E0 FE0F",
			Group:       "objects",
			SubGroup:    "tool",
		},

		128737: {
			Slug:        "shield",
			Character:   "🛡️",
			UnicodeName: "shield",
			CodePoint:   "1F6E1 FE0F",
			Group:       "objects",
			SubGroup:    "tool",
		},

		128738: {
			Slug:        "e0-7-oil-drum",
			Character:   "🛢️",
			UnicodeName: "E0.7 oil drum",
			CodePoint:   "1F6E2 FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128739: {
			Slug:        "e0-7-motorway",
			Character:   "🛣️",
			UnicodeName: "E0.7 motorway",
			CodePoint:   "1F6E3 FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128740: {
			Slug:        "e0-7-railway-track",
			Character:   "🛤️",
			UnicodeName: "E0.7 railway track",
			CodePoint:   "1F6E4 FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128741: {
			Slug:        "e0-7-motor-boat",
			Character:   "🛥️",
			UnicodeName: "E0.7 motor boat",
			CodePoint:   "1F6E5 FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-water",
		},

		128745: {
			Slug:        "e0-7-small-airplane",
			Character:   "🛩️",
			UnicodeName: "E0.7 small airplane",
			CodePoint:   "1F6E9 FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128747: {
			Slug:        "e1-0-airplane-departure",
			Character:   "🛫",
			UnicodeName: "E1.0 airplane departure",
			CodePoint:   "1F6EB",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128748: {
			Slug:        "e1-0-airplane-arrival",
			Character:   "🛬",
			UnicodeName: "E1.0 airplane arrival",
			CodePoint:   "1F6EC",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128752: {
			Slug:        "e0-7-satellite",
			Character:   "🛰️",
			UnicodeName: "E0.7 satellite",
			CodePoint:   "1F6F0 FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128755: {
			Slug:        "e0-7-passenger-ship",
			Character:   "🛳️",
			UnicodeName: "E0.7 passenger ship",
			CodePoint:   "1F6F3 FE0F",
			Group:       "travel-places",
			SubGroup:    "transport-water",
		},

		128756: {
			Slug:        "e3-0-kick-scooter",
			Character:   "🛴",
			UnicodeName: "E3.0 kick scooter",
			CodePoint:   "1F6F4",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128757: {
			Slug:        "e3-0-motor-scooter",
			Character:   "🛵",
			UnicodeName: "E3.0 motor scooter",
			CodePoint:   "1F6F5",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128758: {
			Slug:        "e3-0-canoe",
			Character:   "🛶",
			UnicodeName: "E3.0 canoe",
			CodePoint:   "1F6F6",
			Group:       "travel-places",
			SubGroup:    "transport-water",
		},

		128759: {
			Slug:        "sled",
			Character:   "🛷",
			UnicodeName: "sled",
			CodePoint:   "1F6F7",
			Group:       "activities",
			SubGroup:    "sport",
		},

		128760: {
			Slug:        "e5-0-flying-saucer",
			Character:   "🛸",
			UnicodeName: "E5.0 flying saucer",
			CodePoint:   "1F6F8",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		128761: {
			Slug:        "e11-0-skateboard",
			Character:   "🛹",
			UnicodeName: "E11.0 skateboard",
			CodePoint:   "1F6F9",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128762: {
			Slug:        "e12-0-auto-rickshaw",
			Character:   "🛺",
			UnicodeName: "E12.0 auto rickshaw",
			CodePoint:   "1F6FA",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128763: {
			Slug:        "e13-0-pickup-truck",
			Character:   "🛻",
			UnicodeName: "E13.0 pickup truck",
			CodePoint:   "1F6FB",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128764: {
			Slug:        "e13-0-roller-skate",
			Character:   "🛼",
			UnicodeName: "E13.0 roller skate",
			CodePoint:   "1F6FC",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		128992: {
			Slug:        "orange-circle",
			Character:   "🟠",
			UnicodeName: "orange circle",
			CodePoint:   "1F7E0",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128993: {
			Slug:        "yellow-circle",
			Character:   "🟡",
			UnicodeName: "yellow circle",
			CodePoint:   "1F7E1",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128994: {
			Slug:        "green-circle",
			Character:   "🟢",
			UnicodeName: "green circle",
			CodePoint:   "1F7E2",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128995: {
			Slug:        "purple-circle",
			Character:   "🟣",
			UnicodeName: "purple circle",
			CodePoint:   "1F7E3",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128996: {
			Slug:        "brown-circle",
			Character:   "🟤",
			UnicodeName: "brown circle",
			CodePoint:   "1F7E4",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128997: {
			Slug:        "red-square",
			Character:   "🟥",
			UnicodeName: "red square",
			CodePoint:   "1F7E5",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128998: {
			Slug:        "blue-square",
			Character:   "🟦",
			UnicodeName: "blue square",
			CodePoint:   "1F7E6",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		128999: {
			Slug:        "orange-square",
			Character:   "🟧",
			UnicodeName: "orange square",
			CodePoint:   "1F7E7",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		129000: {
			Slug:        "yellow-square",
			Character:   "🟨",
			UnicodeName: "yellow square",
			CodePoint:   "1F7E8",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		129001: {
			Slug:        "green-square",
			Character:   "🟩",
			UnicodeName: "green square",
			CodePoint:   "1F7E9",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		129002: {
			Slug:        "purple-square",
			Character:   "🟪",
			UnicodeName: "purple square",
			CodePoint:   "1F7EA",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		129003: {
			Slug:        "brown-square",
			Character:   "🟫",
			UnicodeName: "brown square",
			CodePoint:   "1F7EB",
			Group:       "symbols",
			SubGroup:    "geometric",
		},

		129292: {
			Slug:        "e13-0-pinched-fingers",
			Character:   "🤌",
			UnicodeName: "E13.0 pinched fingers",
			CodePoint:   "1F90C",
			Group:       "people-body",
			SubGroup:    "hand-fingers-partial",
		},

		129293: {
			Slug:        "e12-0-white-heart",
			Character:   "🤍",
			UnicodeName: "E12.0 white heart",
			CodePoint:   "1F90D",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		129294: {
			Slug:        "e12-0-brown-heart",
			Character:   "🤎",
			UnicodeName: "E12.0 brown heart",
			CodePoint:   "1F90E",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		129295: {
			Slug:        "e12-0-pinching-hand",
			Character:   "🤏",
			UnicodeName: "E12.0 pinching hand",
			CodePoint:   "1F90F",
			Group:       "people-body",
			SubGroup:    "hand-fingers-partial",
		},

		129296: {
			Slug:        "e1-0-zipper-mouth-face",
			Character:   "🤐",
			UnicodeName: "E1.0 zipper-mouth face",
			CodePoint:   "1F910",
			Group:       "smileys-emotion",
			SubGroup:    "face-neutral-skeptical",
		},

		129297: {
			Slug:        "e1-0-money-mouth-face",
			Character:   "🤑",
			UnicodeName: "E1.0 money-mouth face",
			CodePoint:   "1F911",
			Group:       "smileys-emotion",
			SubGroup:    "face-tongue",
		},

		129298: {
			Slug:        "e1-0-face-with-thermometer",
			Character:   "🤒",
			UnicodeName: "E1.0 face with thermometer",
			CodePoint:   "1F912",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		129299: {
			Slug:        "e1-0-nerd-face",
			Character:   "🤓",
			UnicodeName: "E1.0 nerd face",
			CodePoint:   "1F913",
			Group:       "smileys-emotion",
			SubGroup:    "face-glasses",
		},

		129300: {
			Slug:        "e1-0-thinking-face",
			Character:   "🤔",
			UnicodeName: "E1.0 thinking face",
			CodePoint:   "1F914",
			Group:       "smileys-emotion",
			SubGroup:    "face-hand",
		},

		129301: {
			Slug:        "e1-0-face-with-head-bandage",
			Character:   "🤕",
			UnicodeName: "E1.0 face with head-bandage",
			CodePoint:   "1F915",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		129302: {
			Slug:        "e1-0-robot",
			Character:   "🤖",
			UnicodeName: "E1.0 robot",
			CodePoint:   "1F916",
			Group:       "smileys-emotion",
			SubGroup:    "face-costume",
		},

		129303: {
			Slug:        "e1-0-hugging-face",
			Character:   "🤗",
			UnicodeName: "E1.0 hugging face",
			CodePoint:   "1F917",
			Group:       "smileys-emotion",
			SubGroup:    "face-hand",
		},

		129304: {
			Slug:        "e1-0-sign-of-the-horns",
			Character:   "🤘",
			UnicodeName: "E1.0 sign of the horns",
			CodePoint:   "1F918",
			Group:       "people-body",
			SubGroup:    "hand-fingers-partial",
		},

		129305: {
			Slug:        "e3-0-call-me-hand",
			Character:   "🤙",
			UnicodeName: "E3.0 call me hand",
			CodePoint:   "1F919",
			Group:       "people-body",
			SubGroup:    "hand-fingers-partial",
		},

		129306: {
			Slug:        "e3-0-raised-back-of-hand",
			Character:   "🤚",
			UnicodeName: "E3.0 raised back of hand",
			CodePoint:   "1F91A",
			Group:       "people-body",
			SubGroup:    "hand-fingers-open",
		},

		129307: {
			Slug:        "e3-0-left-facing-fist",
			Character:   "🤛",
			UnicodeName: "E3.0 left-facing fist",
			CodePoint:   "1F91B",
			Group:       "people-body",
			SubGroup:    "hand-fingers-closed",
		},

		129308: {
			Slug:        "e3-0-right-facing-fist",
			Character:   "🤜",
			UnicodeName: "E3.0 right-facing fist",
			CodePoint:   "1F91C",
			Group:       "people-body",
			SubGroup:    "hand-fingers-closed",
		},

		129309: {
			Slug:        "e3-0-handshake",
			Character:   "🤝",
			UnicodeName: "E3.0 handshake",
			CodePoint:   "1F91D",
			Group:       "people-body",
			SubGroup:    "hands",
		},

		129310: {
			Slug:        "e3-0-crossed-fingers",
			Character:   "🤞",
			UnicodeName: "E3.0 crossed fingers",
			CodePoint:   "1F91E",
			Group:       "people-body",
			SubGroup:    "hand-fingers-partial",
		},

		129311: {
			Slug:        "e5-0-love-you-gesture",
			Character:   "🤟",
			UnicodeName: "E5.0 love-you gesture",
			CodePoint:   "1F91F",
			Group:       "people-body",
			SubGroup:    "hand-fingers-partial",
		},

		129312: {
			Slug:        "e3-0-cowboy-hat-face",
			Character:   "🤠",
			UnicodeName: "E3.0 cowboy hat face",
			CodePoint:   "1F920",
			Group:       "smileys-emotion",
			SubGroup:    "face-hat",
		},

		129313: {
			Slug:        "e3-0-clown-face",
			Character:   "🤡",
			UnicodeName: "E3.0 clown face",
			CodePoint:   "1F921",
			Group:       "smileys-emotion",
			SubGroup:    "face-costume",
		},

		129314: {
			Slug:        "e3-0-nauseated-face",
			Character:   "🤢",
			UnicodeName: "E3.0 nauseated face",
			CodePoint:   "1F922",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		129315: {
			Slug:        "e3-0-rolling-on-the-floor-laughing",
			Character:   "🤣",
			UnicodeName: "E3.0 rolling on the floor laughing",
			CodePoint:   "1F923",
			Group:       "smileys-emotion",
			SubGroup:    "face-smiling",
		},

		129316: {
			Slug:        "e3-0-drooling-face",
			Character:   "🤤",
			UnicodeName: "E3.0 drooling face",
			CodePoint:   "1F924",
			Group:       "smileys-emotion",
			SubGroup:    "face-sleepy",
		},

		129317: {
			Slug:        "e3-0-lying-face",
			Character:   "🤥",
			UnicodeName: "E3.0 lying face",
			CodePoint:   "1F925",
			Group:       "smileys-emotion",
			SubGroup:    "face-neutral-skeptical",
		},

		129318: {
			Slug:        "e3-0-person-facepalming",
			Character:   "🤦",
			UnicodeName: "E3.0 person facepalming",
			CodePoint:   "1F926",
			Group:       "people-body",
			SubGroup:    "person-gesture",
		},

		129319: {
			Slug:        "e3-0-sneezing-face",
			Character:   "🤧",
			UnicodeName: "E3.0 sneezing face",
			CodePoint:   "1F927",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		129320: {
			Slug:        "e5-0-face-with-raised-eyebrow",
			Character:   "🤨",
			UnicodeName: "E5.0 face with raised eyebrow",
			CodePoint:   "1F928",
			Group:       "smileys-emotion",
			SubGroup:    "face-neutral-skeptical",
		},

		129321: {
			Slug:        "e5-0-star-struck",
			Character:   "🤩",
			UnicodeName: "E5.0 star-struck",
			CodePoint:   "1F929",
			Group:       "smileys-emotion",
			SubGroup:    "face-affection",
		},

		129322: {
			Slug:        "e5-0-zany-face",
			Character:   "🤪",
			UnicodeName: "E5.0 zany face",
			CodePoint:   "1F92A",
			Group:       "smileys-emotion",
			SubGroup:    "face-tongue",
		},

		129323: {
			Slug:        "e5-0-shushing-face",
			Character:   "🤫",
			UnicodeName: "E5.0 shushing face",
			CodePoint:   "1F92B",
			Group:       "smileys-emotion",
			SubGroup:    "face-hand",
		},

		129324: {
			Slug:        "e5-0-face-with-symbols-on-mouth",
			Character:   "🤬",
			UnicodeName: "E5.0 face with symbols on mouth",
			CodePoint:   "1F92C",
			Group:       "smileys-emotion",
			SubGroup:    "face-negative",
		},

		129325: {
			Slug:        "e5-0-face-with-hand-over-mouth",
			Character:   "🤭",
			UnicodeName: "E5.0 face with hand over mouth",
			CodePoint:   "1F92D",
			Group:       "smileys-emotion",
			SubGroup:    "face-hand",
		},

		129326: {
			Slug:        "e5-0-face-vomiting",
			Character:   "🤮",
			UnicodeName: "E5.0 face vomiting",
			CodePoint:   "1F92E",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		129327: {
			Slug:        "e5-0-exploding-head",
			Character:   "🤯",
			UnicodeName: "E5.0 exploding head",
			CodePoint:   "1F92F",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		129328: {
			Slug:        "e3-0-pregnant-woman",
			Character:   "🤰",
			UnicodeName: "E3.0 pregnant woman",
			CodePoint:   "1F930",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		129329: {
			Slug:        "e5-0-breast-feeding",
			Character:   "🤱",
			UnicodeName: "E5.0 breast-feeding",
			CodePoint:   "1F931",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		129330: {
			Slug:        "e5-0-palms-up-together",
			Character:   "🤲",
			UnicodeName: "E5.0 palms up together",
			CodePoint:   "1F932",
			Group:       "people-body",
			SubGroup:    "hands",
		},

		129331: {
			Slug:        "e3-0-selfie",
			Character:   "🤳",
			UnicodeName: "E3.0 selfie",
			CodePoint:   "1F933",
			Group:       "people-body",
			SubGroup:    "hand-prop",
		},

		129332: {
			Slug:        "e3-0-prince",
			Character:   "🤴",
			UnicodeName: "E3.0 prince",
			CodePoint:   "1F934",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		129333: {
			Slug:        "e3-0-person-in-tuxedo",
			Character:   "🤵",
			UnicodeName: "E3.0 person in tuxedo",
			CodePoint:   "1F935",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		129334: {
			Slug:        "e3-0-mrs-claus",
			Character:   "🤶",
			UnicodeName: "E3.0 Mrs. Claus",
			CodePoint:   "1F936",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		129335: {
			Slug:        "e3-0-person-shrugging",
			Character:   "🤷",
			UnicodeName: "E3.0 person shrugging",
			CodePoint:   "1F937",
			Group:       "people-body",
			SubGroup:    "person-gesture",
		},

		129336: {
			Slug:        "e3-0-person-cartwheeling",
			Character:   "🤸",
			UnicodeName: "E3.0 person cartwheeling",
			CodePoint:   "1F938",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		129337: {
			Slug:        "e3-0-person-juggling",
			Character:   "🤹",
			UnicodeName: "E3.0 person juggling",
			CodePoint:   "1F939",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		129338: {
			Slug:        "e3-0-person-fencing",
			Character:   "🤺",
			UnicodeName: "E3.0 person fencing",
			CodePoint:   "1F93A",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		129340: {
			Slug:        "e3-0-people-wrestling",
			Character:   "🤼",
			UnicodeName: "E3.0 people wrestling",
			CodePoint:   "1F93C",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		129341: {
			Slug:        "e3-0-person-playing-water-polo",
			Character:   "🤽",
			UnicodeName: "E3.0 person playing water polo",
			CodePoint:   "1F93D",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		129342: {
			Slug:        "e3-0-person-playing-handball",
			Character:   "🤾",
			UnicodeName: "E3.0 person playing handball",
			CodePoint:   "1F93E",
			Group:       "people-body",
			SubGroup:    "person-sport",
		},

		129343: {
			Slug:        "diving-mask",
			Character:   "🤿",
			UnicodeName: "diving mask",
			CodePoint:   "1F93F",
			Group:       "activities",
			SubGroup:    "sport",
		},

		129344: {
			Slug:        "e3-0-wilted-flower",
			Character:   "🥀",
			UnicodeName: "E3.0 wilted flower",
			CodePoint:   "1F940",
			Group:       "animals-nature",
			SubGroup:    "plant-flower",
		},

		129345: {
			Slug:        "drum",
			Character:   "🥁",
			UnicodeName: "drum",
			CodePoint:   "1F941",
			Group:       "objects",
			SubGroup:    "musical-instrument",
		},

		129346: {
			Slug:        "e3-0-clinking-glasses",
			Character:   "🥂",
			UnicodeName: "E3.0 clinking glasses",
			CodePoint:   "1F942",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		129347: {
			Slug:        "e3-0-tumbler-glass",
			Character:   "🥃",
			UnicodeName: "E3.0 tumbler glass",
			CodePoint:   "1F943",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		129348: {
			Slug:        "e3-0-spoon",
			Character:   "🥄",
			UnicodeName: "E3.0 spoon",
			CodePoint:   "1F944",
			Group:       "food-drink",
			SubGroup:    "dishware",
		},

		129349: {
			Slug:        "goal-net",
			Character:   "🥅",
			UnicodeName: "goal net",
			CodePoint:   "1F945",
			Group:       "activities",
			SubGroup:    "sport",
		},

		129351: {
			Slug:        "1st-place-medal",
			Character:   "🥇",
			UnicodeName: "1st place medal",
			CodePoint:   "1F947",
			Group:       "activities",
			SubGroup:    "award-medal",
		},

		129352: {
			Slug:        "2nd-place-medal",
			Character:   "🥈",
			UnicodeName: "2nd place medal",
			CodePoint:   "1F948",
			Group:       "activities",
			SubGroup:    "award-medal",
		},

		129353: {
			Slug:        "3rd-place-medal",
			Character:   "🥉",
			UnicodeName: "3rd place medal",
			CodePoint:   "1F949",
			Group:       "activities",
			SubGroup:    "award-medal",
		},

		129354: {
			Slug:        "boxing-glove",
			Character:   "🥊",
			UnicodeName: "boxing glove",
			CodePoint:   "1F94A",
			Group:       "activities",
			SubGroup:    "sport",
		},

		129355: {
			Slug:        "martial-arts-uniform",
			Character:   "🥋",
			UnicodeName: "martial arts uniform",
			CodePoint:   "1F94B",
			Group:       "activities",
			SubGroup:    "sport",
		},

		129356: {
			Slug:        "curling-stone",
			Character:   "🥌",
			UnicodeName: "curling stone",
			CodePoint:   "1F94C",
			Group:       "activities",
			SubGroup:    "sport",
		},

		129357: {
			Slug:        "lacrosse",
			Character:   "🥍",
			UnicodeName: "lacrosse",
			CodePoint:   "1F94D",
			Group:       "activities",
			SubGroup:    "sport",
		},

		129358: {
			Slug:        "softball",
			Character:   "🥎",
			UnicodeName: "softball",
			CodePoint:   "1F94E",
			Group:       "activities",
			SubGroup:    "sport",
		},

		129359: {
			Slug:        "flying-disc",
			Character:   "🥏",
			UnicodeName: "flying disc",
			CodePoint:   "1F94F",
			Group:       "activities",
			SubGroup:    "sport",
		},

		129360: {
			Slug:        "e3-0-croissant",
			Character:   "🥐",
			UnicodeName: "E3.0 croissant",
			CodePoint:   "1F950",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129361: {
			Slug:        "e3-0-avocado",
			Character:   "🥑",
			UnicodeName: "E3.0 avocado",
			CodePoint:   "1F951",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		129362: {
			Slug:        "e3-0-cucumber",
			Character:   "🥒",
			UnicodeName: "E3.0 cucumber",
			CodePoint:   "1F952",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		129363: {
			Slug:        "e3-0-bacon",
			Character:   "🥓",
			UnicodeName: "E3.0 bacon",
			CodePoint:   "1F953",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129364: {
			Slug:        "e3-0-potato",
			Character:   "🥔",
			UnicodeName: "E3.0 potato",
			CodePoint:   "1F954",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		129365: {
			Slug:        "e3-0-carrot",
			Character:   "🥕",
			UnicodeName: "E3.0 carrot",
			CodePoint:   "1F955",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		129366: {
			Slug:        "e3-0-baguette-bread",
			Character:   "🥖",
			UnicodeName: "E3.0 baguette bread",
			CodePoint:   "1F956",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129367: {
			Slug:        "e3-0-green-salad",
			Character:   "🥗",
			UnicodeName: "E3.0 green salad",
			CodePoint:   "1F957",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129368: {
			Slug:        "e3-0-shallow-pan-of-food",
			Character:   "🥘",
			UnicodeName: "E3.0 shallow pan of food",
			CodePoint:   "1F958",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129369: {
			Slug:        "e3-0-stuffed-flatbread",
			Character:   "🥙",
			UnicodeName: "E3.0 stuffed flatbread",
			CodePoint:   "1F959",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129370: {
			Slug:        "e3-0-egg",
			Character:   "🥚",
			UnicodeName: "E3.0 egg",
			CodePoint:   "1F95A",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129371: {
			Slug:        "e3-0-glass-of-milk",
			Character:   "🥛",
			UnicodeName: "E3.0 glass of milk",
			CodePoint:   "1F95B",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		129372: {
			Slug:        "e3-0-peanuts",
			Character:   "🥜",
			UnicodeName: "E3.0 peanuts",
			CodePoint:   "1F95C",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		129373: {
			Slug:        "e3-0-kiwi-fruit",
			Character:   "🥝",
			UnicodeName: "E3.0 kiwi fruit",
			CodePoint:   "1F95D",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		129374: {
			Slug:        "e3-0-pancakes",
			Character:   "🥞",
			UnicodeName: "E3.0 pancakes",
			CodePoint:   "1F95E",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129375: {
			Slug:        "e5-0-dumpling",
			Character:   "🥟",
			UnicodeName: "E5.0 dumpling",
			CodePoint:   "1F95F",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		129376: {
			Slug:        "e5-0-fortune-cookie",
			Character:   "🥠",
			UnicodeName: "E5.0 fortune cookie",
			CodePoint:   "1F960",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		129377: {
			Slug:        "e5-0-takeout-box",
			Character:   "🥡",
			UnicodeName: "E5.0 takeout box",
			CodePoint:   "1F961",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		129378: {
			Slug:        "e5-0-chopsticks",
			Character:   "🥢",
			UnicodeName: "E5.0 chopsticks",
			CodePoint:   "1F962",
			Group:       "food-drink",
			SubGroup:    "dishware",
		},

		129379: {
			Slug:        "e5-0-bowl-with-spoon",
			Character:   "🥣",
			UnicodeName: "E5.0 bowl with spoon",
			CodePoint:   "1F963",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129380: {
			Slug:        "e5-0-cup-with-straw",
			Character:   "🥤",
			UnicodeName: "E5.0 cup with straw",
			CodePoint:   "1F964",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		129381: {
			Slug:        "e5-0-coconut",
			Character:   "🥥",
			UnicodeName: "E5.0 coconut",
			CodePoint:   "1F965",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		129382: {
			Slug:        "e5-0-broccoli",
			Character:   "🥦",
			UnicodeName: "E5.0 broccoli",
			CodePoint:   "1F966",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		129383: {
			Slug:        "e5-0-pie",
			Character:   "🥧",
			UnicodeName: "E5.0 pie",
			CodePoint:   "1F967",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		129384: {
			Slug:        "e5-0-pretzel",
			Character:   "🥨",
			UnicodeName: "E5.0 pretzel",
			CodePoint:   "1F968",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129385: {
			Slug:        "e5-0-cut-of-meat",
			Character:   "🥩",
			UnicodeName: "E5.0 cut of meat",
			CodePoint:   "1F969",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129386: {
			Slug:        "e5-0-sandwich",
			Character:   "🥪",
			UnicodeName: "E5.0 sandwich",
			CodePoint:   "1F96A",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129387: {
			Slug:        "e5-0-canned-food",
			Character:   "🥫",
			UnicodeName: "E5.0 canned food",
			CodePoint:   "1F96B",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129388: {
			Slug:        "e11-0-leafy-green",
			Character:   "🥬",
			UnicodeName: "E11.0 leafy green",
			CodePoint:   "1F96C",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		129389: {
			Slug:        "e11-0-mango",
			Character:   "🥭",
			UnicodeName: "E11.0 mango",
			CodePoint:   "1F96D",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		129390: {
			Slug:        "e11-0-moon-cake",
			Character:   "🥮",
			UnicodeName: "E11.0 moon cake",
			CodePoint:   "1F96E",
			Group:       "food-drink",
			SubGroup:    "food-asian",
		},

		129391: {
			Slug:        "e11-0-bagel",
			Character:   "🥯",
			UnicodeName: "E11.0 bagel",
			CodePoint:   "1F96F",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129392: {
			Slug:        "e11-0-smiling-face-with-hearts",
			Character:   "🥰",
			UnicodeName: "E11.0 smiling face with hearts",
			CodePoint:   "1F970",
			Group:       "smileys-emotion",
			SubGroup:    "face-affection",
		},

		129393: {
			Slug:        "e12-0-yawning-face",
			Character:   "🥱",
			UnicodeName: "E12.0 yawning face",
			CodePoint:   "1F971",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		129394: {
			Slug:        "e13-0-smiling-face-with-tear",
			Character:   "🥲",
			UnicodeName: "E13.0 smiling face with tear",
			CodePoint:   "1F972",
			Group:       "smileys-emotion",
			SubGroup:    "face-affection",
		},

		129395: {
			Slug:        "e11-0-partying-face",
			Character:   "🥳",
			UnicodeName: "E11.0 partying face",
			CodePoint:   "1F973",
			Group:       "smileys-emotion",
			SubGroup:    "face-hat",
		},

		129396: {
			Slug:        "e11-0-woozy-face",
			Character:   "🥴",
			UnicodeName: "E11.0 woozy face",
			CodePoint:   "1F974",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		129397: {
			Slug:        "e11-0-hot-face",
			Character:   "🥵",
			UnicodeName: "E11.0 hot face",
			CodePoint:   "1F975",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		129398: {
			Slug:        "e11-0-cold-face",
			Character:   "🥶",
			UnicodeName: "E11.0 cold face",
			CodePoint:   "1F976",
			Group:       "smileys-emotion",
			SubGroup:    "face-unwell",
		},

		129399: {
			Slug:        "e13-0-ninja",
			Character:   "🥷",
			UnicodeName: "E13.0 ninja",
			CodePoint:   "1F977",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		129400: {
			Slug:        "e13-0-disguised-face",
			Character:   "🥸",
			UnicodeName: "E13.0 disguised face",
			CodePoint:   "1F978",
			Group:       "smileys-emotion",
			SubGroup:    "face-hat",
		},

		129402: {
			Slug:        "e11-0-pleading-face",
			Character:   "🥺",
			UnicodeName: "E11.0 pleading face",
			CodePoint:   "1F97A",
			Group:       "smileys-emotion",
			SubGroup:    "face-concerned",
		},

		129403: {
			Slug:        "sari",
			Character:   "🥻",
			UnicodeName: "sari",
			CodePoint:   "1F97B",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129404: {
			Slug:        "lab-coat",
			Character:   "🥼",
			UnicodeName: "lab coat",
			CodePoint:   "1F97C",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129405: {
			Slug:        "goggles",
			Character:   "🥽",
			UnicodeName: "goggles",
			CodePoint:   "1F97D",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129406: {
			Slug:        "hiking-boot",
			Character:   "🥾",
			UnicodeName: "hiking boot",
			CodePoint:   "1F97E",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129407: {
			Slug:        "flat-shoe",
			Character:   "🥿",
			UnicodeName: "flat shoe",
			CodePoint:   "1F97F",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129408: {
			Slug:        "e1-0-crab",
			Character:   "🦀",
			UnicodeName: "E1.0 crab",
			CodePoint:   "1F980",
			Group:       "food-drink",
			SubGroup:    "food-marine",
		},

		129409: {
			Slug:        "e1-0-lion",
			Character:   "🦁",
			UnicodeName: "E1.0 lion",
			CodePoint:   "1F981",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129410: {
			Slug:        "e1-0-scorpion",
			Character:   "🦂",
			UnicodeName: "E1.0 scorpion",
			CodePoint:   "1F982",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		129411: {
			Slug:        "e1-0-turkey",
			Character:   "🦃",
			UnicodeName: "E1.0 turkey",
			CodePoint:   "1F983",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		129412: {
			Slug:        "e1-0-unicorn",
			Character:   "🦄",
			UnicodeName: "E1.0 unicorn",
			CodePoint:   "1F984",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129413: {
			Slug:        "e3-0-eagle",
			Character:   "🦅",
			UnicodeName: "E3.0 eagle",
			CodePoint:   "1F985",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		129414: {
			Slug:        "e3-0-duck",
			Character:   "🦆",
			UnicodeName: "E3.0 duck",
			CodePoint:   "1F986",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		129415: {
			Slug:        "e3-0-bat",
			Character:   "🦇",
			UnicodeName: "E3.0 bat",
			CodePoint:   "1F987",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129416: {
			Slug:        "e3-0-shark",
			Character:   "🦈",
			UnicodeName: "E3.0 shark",
			CodePoint:   "1F988",
			Group:       "animals-nature",
			SubGroup:    "animal-marine",
		},

		129417: {
			Slug:        "e3-0-owl",
			Character:   "🦉",
			UnicodeName: "E3.0 owl",
			CodePoint:   "1F989",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		129418: {
			Slug:        "e3-0-fox",
			Character:   "🦊",
			UnicodeName: "E3.0 fox",
			CodePoint:   "1F98A",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129419: {
			Slug:        "e3-0-butterfly",
			Character:   "🦋",
			UnicodeName: "E3.0 butterfly",
			CodePoint:   "1F98B",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		129420: {
			Slug:        "e3-0-deer",
			Character:   "🦌",
			UnicodeName: "E3.0 deer",
			CodePoint:   "1F98C",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129421: {
			Slug:        "e3-0-gorilla",
			Character:   "🦍",
			UnicodeName: "E3.0 gorilla",
			CodePoint:   "1F98D",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129422: {
			Slug:        "e3-0-lizard",
			Character:   "🦎",
			UnicodeName: "E3.0 lizard",
			CodePoint:   "1F98E",
			Group:       "animals-nature",
			SubGroup:    "animal-reptile",
		},

		129423: {
			Slug:        "e3-0-rhinoceros",
			Character:   "🦏",
			UnicodeName: "E3.0 rhinoceros",
			CodePoint:   "1F98F",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129424: {
			Slug:        "e3-0-shrimp",
			Character:   "🦐",
			UnicodeName: "E3.0 shrimp",
			CodePoint:   "1F990",
			Group:       "food-drink",
			SubGroup:    "food-marine",
		},

		129425: {
			Slug:        "e3-0-squid",
			Character:   "🦑",
			UnicodeName: "E3.0 squid",
			CodePoint:   "1F991",
			Group:       "food-drink",
			SubGroup:    "food-marine",
		},

		129426: {
			Slug:        "e5-0-giraffe",
			Character:   "🦒",
			UnicodeName: "E5.0 giraffe",
			CodePoint:   "1F992",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129427: {
			Slug:        "e5-0-zebra",
			Character:   "🦓",
			UnicodeName: "E5.0 zebra",
			CodePoint:   "1F993",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129428: {
			Slug:        "e5-0-hedgehog",
			Character:   "🦔",
			UnicodeName: "E5.0 hedgehog",
			CodePoint:   "1F994",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129429: {
			Slug:        "e5-0-sauropod",
			Character:   "🦕",
			UnicodeName: "E5.0 sauropod",
			CodePoint:   "1F995",
			Group:       "animals-nature",
			SubGroup:    "animal-reptile",
		},

		129430: {
			Slug:        "e5-0-t-rex",
			Character:   "🦖",
			UnicodeName: "E5.0 T-Rex",
			CodePoint:   "1F996",
			Group:       "animals-nature",
			SubGroup:    "animal-reptile",
		},

		129431: {
			Slug:        "e5-0-cricket",
			Character:   "🦗",
			UnicodeName: "E5.0 cricket",
			CodePoint:   "1F997",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		129432: {
			Slug:        "e11-0-kangaroo",
			Character:   "🦘",
			UnicodeName: "E11.0 kangaroo",
			CodePoint:   "1F998",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129433: {
			Slug:        "e11-0-llama",
			Character:   "🦙",
			UnicodeName: "E11.0 llama",
			CodePoint:   "1F999",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129434: {
			Slug:        "e11-0-peacock",
			Character:   "🦚",
			UnicodeName: "E11.0 peacock",
			CodePoint:   "1F99A",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		129435: {
			Slug:        "e11-0-hippopotamus",
			Character:   "🦛",
			UnicodeName: "E11.0 hippopotamus",
			CodePoint:   "1F99B",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129436: {
			Slug:        "e11-0-parrot",
			Character:   "🦜",
			UnicodeName: "E11.0 parrot",
			CodePoint:   "1F99C",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		129437: {
			Slug:        "e11-0-raccoon",
			Character:   "🦝",
			UnicodeName: "E11.0 raccoon",
			CodePoint:   "1F99D",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129438: {
			Slug:        "e11-0-lobster",
			Character:   "🦞",
			UnicodeName: "E11.0 lobster",
			CodePoint:   "1F99E",
			Group:       "food-drink",
			SubGroup:    "food-marine",
		},

		129439: {
			Slug:        "e11-0-mosquito",
			Character:   "🦟",
			UnicodeName: "E11.0 mosquito",
			CodePoint:   "1F99F",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		129440: {
			Slug:        "e11-0-microbe",
			Character:   "🦠",
			UnicodeName: "E11.0 microbe",
			CodePoint:   "1F9A0",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		129441: {
			Slug:        "e11-0-badger",
			Character:   "🦡",
			UnicodeName: "E11.0 badger",
			CodePoint:   "1F9A1",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129442: {
			Slug:        "e11-0-swan",
			Character:   "🦢",
			UnicodeName: "E11.0 swan",
			CodePoint:   "1F9A2",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		129443: {
			Slug:        "e13-0-mammoth",
			Character:   "🦣",
			UnicodeName: "E13.0 mammoth",
			CodePoint:   "1F9A3",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129444: {
			Slug:        "e13-0-dodo",
			Character:   "🦤",
			UnicodeName: "E13.0 dodo",
			CodePoint:   "1F9A4",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		129445: {
			Slug:        "e12-0-sloth",
			Character:   "🦥",
			UnicodeName: "E12.0 sloth",
			CodePoint:   "1F9A5",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129446: {
			Slug:        "e12-0-otter",
			Character:   "🦦",
			UnicodeName: "E12.0 otter",
			CodePoint:   "1F9A6",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129447: {
			Slug:        "e12-0-orangutan",
			Character:   "🦧",
			UnicodeName: "E12.0 orangutan",
			CodePoint:   "1F9A7",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129448: {
			Slug:        "e12-0-skunk",
			Character:   "🦨",
			UnicodeName: "E12.0 skunk",
			CodePoint:   "1F9A8",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129449: {
			Slug:        "e12-0-flamingo",
			Character:   "🦩",
			UnicodeName: "E12.0 flamingo",
			CodePoint:   "1F9A9",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		129450: {
			Slug:        "e12-0-oyster",
			Character:   "🦪",
			UnicodeName: "E12.0 oyster",
			CodePoint:   "1F9AA",
			Group:       "food-drink",
			SubGroup:    "food-marine",
		},

		129451: {
			Slug:        "e13-0-beaver",
			Character:   "🦫",
			UnicodeName: "E13.0 beaver",
			CodePoint:   "1F9AB",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129452: {
			Slug:        "e13-0-bison",
			Character:   "🦬",
			UnicodeName: "E13.0 bison",
			CodePoint:   "1F9AC",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129453: {
			Slug:        "e13-0-seal",
			Character:   "🦭",
			UnicodeName: "E13.0 seal",
			CodePoint:   "1F9AD",
			Group:       "animals-nature",
			SubGroup:    "animal-marine",
		},

		129454: {
			Slug:        "e12-0-guide-dog",
			Character:   "🦮",
			UnicodeName: "E12.0 guide dog",
			CodePoint:   "1F9AE",
			Group:       "animals-nature",
			SubGroup:    "animal-mammal",
		},

		129455: {
			Slug:        "white-cane",
			Character:   "🦯",
			UnicodeName: "white cane",
			CodePoint:   "1F9AF",
			Group:       "objects",
			SubGroup:    "tool",
		},

		129460: {
			Slug:        "e11-0-bone",
			Character:   "🦴",
			UnicodeName: "E11.0 bone",
			CodePoint:   "1F9B4",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		129461: {
			Slug:        "e11-0-leg",
			Character:   "🦵",
			UnicodeName: "E11.0 leg",
			CodePoint:   "1F9B5",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		129462: {
			Slug:        "e11-0-foot",
			Character:   "🦶",
			UnicodeName: "E11.0 foot",
			CodePoint:   "1F9B6",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		129463: {
			Slug:        "e11-0-tooth",
			Character:   "🦷",
			UnicodeName: "E11.0 tooth",
			CodePoint:   "1F9B7",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		129464: {
			Slug:        "e11-0-superhero",
			Character:   "🦸",
			UnicodeName: "E11.0 superhero",
			CodePoint:   "1F9B8",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		129465: {
			Slug:        "e11-0-supervillain",
			Character:   "🦹",
			UnicodeName: "E11.0 supervillain",
			CodePoint:   "1F9B9",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		129466: {
			Slug:        "safety-vest",
			Character:   "🦺",
			UnicodeName: "safety vest",
			CodePoint:   "1F9BA",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129467: {
			Slug:        "e12-0-ear-with-hearing-aid",
			Character:   "🦻",
			UnicodeName: "E12.0 ear with hearing aid",
			CodePoint:   "1F9BB",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		129468: {
			Slug:        "e12-0-motorized-wheelchair",
			Character:   "🦼",
			UnicodeName: "E12.0 motorized wheelchair",
			CodePoint:   "1F9BC",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		129469: {
			Slug:        "e12-0-manual-wheelchair",
			Character:   "🦽",
			UnicodeName: "E12.0 manual wheelchair",
			CodePoint:   "1F9BD",
			Group:       "travel-places",
			SubGroup:    "transport-ground",
		},

		129470: {
			Slug:        "e12-0-mechanical-arm",
			Character:   "🦾",
			UnicodeName: "E12.0 mechanical arm",
			CodePoint:   "1F9BE",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		129471: {
			Slug:        "e12-0-mechanical-leg",
			Character:   "🦿",
			UnicodeName: "E12.0 mechanical leg",
			CodePoint:   "1F9BF",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		129472: {
			Slug:        "e1-0-cheese-wedge",
			Character:   "🧀",
			UnicodeName: "E1.0 cheese wedge",
			CodePoint:   "1F9C0",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129473: {
			Slug:        "e11-0-cupcake",
			Character:   "🧁",
			UnicodeName: "E11.0 cupcake",
			CodePoint:   "1F9C1",
			Group:       "food-drink",
			SubGroup:    "food-sweet",
		},

		129474: {
			Slug:        "e11-0-salt",
			Character:   "🧂",
			UnicodeName: "E11.0 salt",
			CodePoint:   "1F9C2",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129475: {
			Slug:        "e12-0-beverage-box",
			Character:   "🧃",
			UnicodeName: "E12.0 beverage box",
			CodePoint:   "1F9C3",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		129476: {
			Slug:        "e12-0-garlic",
			Character:   "🧄",
			UnicodeName: "E12.0 garlic",
			CodePoint:   "1F9C4",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		129477: {
			Slug:        "e12-0-onion",
			Character:   "🧅",
			UnicodeName: "E12.0 onion",
			CodePoint:   "1F9C5",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		129478: {
			Slug:        "e12-0-falafel",
			Character:   "🧆",
			UnicodeName: "E12.0 falafel",
			CodePoint:   "1F9C6",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129479: {
			Slug:        "e12-0-waffle",
			Character:   "🧇",
			UnicodeName: "E12.0 waffle",
			CodePoint:   "1F9C7",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129480: {
			Slug:        "e12-0-butter",
			Character:   "🧈",
			UnicodeName: "E12.0 butter",
			CodePoint:   "1F9C8",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129481: {
			Slug:        "e12-0-mate",
			Character:   "🧉",
			UnicodeName: "E12.0 mate",
			CodePoint:   "1F9C9",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		129482: {
			Slug:        "e12-0-ice",
			Character:   "🧊",
			UnicodeName: "E12.0 ice",
			CodePoint:   "1F9CA",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		129483: {
			Slug:        "e13-0-bubble-tea",
			Character:   "🧋",
			UnicodeName: "E13.0 bubble tea",
			CodePoint:   "1F9CB",
			Group:       "food-drink",
			SubGroup:    "drink",
		},

		129485: {
			Slug:        "e12-0-person-standing",
			Character:   "🧍",
			UnicodeName: "E12.0 person standing",
			CodePoint:   "1F9CD",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		129486: {
			Slug:        "e12-0-person-kneeling",
			Character:   "🧎",
			UnicodeName: "E12.0 person kneeling",
			CodePoint:   "1F9CE",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		129487: {
			Slug:        "e12-0-deaf-person",
			Character:   "🧏",
			UnicodeName: "E12.0 deaf person",
			CodePoint:   "1F9CF",
			Group:       "people-body",
			SubGroup:    "person-gesture",
		},

		129488: {
			Slug:        "e5-0-face-with-monocle",
			Character:   "🧐",
			UnicodeName: "E5.0 face with monocle",
			CodePoint:   "1F9D0",
			Group:       "smileys-emotion",
			SubGroup:    "face-glasses",
		},

		129489: {
			Slug:        "e12-0-people-holding-hands",
			Character:   "🧑‍🤝‍🧑",
			UnicodeName: "E12.0 people holding hands",
			CodePoint:   "1F9D1 200D 1F91D 200D 1F9D1",
			Group:       "people-body",
			SubGroup:    "family",
		},

		129490: {
			Slug:        "e5-0-child",
			Character:   "🧒",
			UnicodeName: "E5.0 child",
			CodePoint:   "1F9D2",
			Group:       "people-body",
			SubGroup:    "person",
		},

		129491: {
			Slug:        "e5-0-older-person",
			Character:   "🧓",
			UnicodeName: "E5.0 older person",
			CodePoint:   "1F9D3",
			Group:       "people-body",
			SubGroup:    "person",
		},

		129492: {
			Slug:        "e5-0-man-beard",
			Character:   "🧔",
			UnicodeName: "E5.0 man: beard",
			CodePoint:   "1F9D4",
			Group:       "people-body",
			SubGroup:    "person",
		},

		129493: {
			Slug:        "e5-0-woman-with-headscarf",
			Character:   "🧕",
			UnicodeName: "E5.0 woman with headscarf",
			CodePoint:   "1F9D5",
			Group:       "people-body",
			SubGroup:    "person-role",
		},

		129494: {
			Slug:        "e5-0-person-in-steamy-room",
			Character:   "🧖",
			UnicodeName: "E5.0 person in steamy room",
			CodePoint:   "1F9D6",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		129495: {
			Slug:        "e5-0-person-climbing",
			Character:   "🧗",
			UnicodeName: "E5.0 person climbing",
			CodePoint:   "1F9D7",
			Group:       "people-body",
			SubGroup:    "person-activity",
		},

		129496: {
			Slug:        "e5-0-person-in-lotus-position",
			Character:   "🧘",
			UnicodeName: "E5.0 person in lotus position",
			CodePoint:   "1F9D8",
			Group:       "people-body",
			SubGroup:    "person-resting",
		},

		129497: {
			Slug:        "e5-0-mage",
			Character:   "🧙",
			UnicodeName: "E5.0 mage",
			CodePoint:   "1F9D9",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		129498: {
			Slug:        "e5-0-fairy",
			Character:   "🧚",
			UnicodeName: "E5.0 fairy",
			CodePoint:   "1F9DA",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		129499: {
			Slug:        "e5-0-vampire",
			Character:   "🧛",
			UnicodeName: "E5.0 vampire",
			CodePoint:   "1F9DB",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		129500: {
			Slug:        "e5-0-merperson",
			Character:   "🧜",
			UnicodeName: "E5.0 merperson",
			CodePoint:   "1F9DC",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		129501: {
			Slug:        "e5-0-elf",
			Character:   "🧝",
			UnicodeName: "E5.0 elf",
			CodePoint:   "1F9DD",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		129502: {
			Slug:        "e5-0-genie",
			Character:   "🧞",
			UnicodeName: "E5.0 genie",
			CodePoint:   "1F9DE",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		129503: {
			Slug:        "e5-0-zombie",
			Character:   "🧟",
			UnicodeName: "E5.0 zombie",
			CodePoint:   "1F9DF",
			Group:       "people-body",
			SubGroup:    "person-fantasy",
		},

		129504: {
			Slug:        "e5-0-brain",
			Character:   "🧠",
			UnicodeName: "E5.0 brain",
			CodePoint:   "1F9E0",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		129505: {
			Slug:        "e5-0-orange-heart",
			Character:   "🧡",
			UnicodeName: "E5.0 orange heart",
			CodePoint:   "1F9E1",
			Group:       "smileys-emotion",
			SubGroup:    "emotion",
		},

		129506: {
			Slug:        "billed-cap",
			Character:   "🧢",
			UnicodeName: "billed cap",
			CodePoint:   "1F9E2",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129507: {
			Slug:        "scarf",
			Character:   "🧣",
			UnicodeName: "scarf",
			CodePoint:   "1F9E3",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129508: {
			Slug:        "gloves",
			Character:   "🧤",
			UnicodeName: "gloves",
			CodePoint:   "1F9E4",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129509: {
			Slug:        "coat",
			Character:   "🧥",
			UnicodeName: "coat",
			CodePoint:   "1F9E5",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129510: {
			Slug:        "socks",
			Character:   "🧦",
			UnicodeName: "socks",
			CodePoint:   "1F9E6",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129511: {
			Slug:        "red-envelope",
			Character:   "🧧",
			UnicodeName: "red envelope",
			CodePoint:   "1F9E7",
			Group:       "activities",
			SubGroup:    "event",
		},

		129512: {
			Slug:        "firecracker",
			Character:   "🧨",
			UnicodeName: "firecracker",
			CodePoint:   "1F9E8",
			Group:       "activities",
			SubGroup:    "event",
		},

		129513: {
			Slug:        "puzzle-piece",
			Character:   "🧩",
			UnicodeName: "puzzle piece",
			CodePoint:   "1F9E9",
			Group:       "activities",
			SubGroup:    "game",
		},

		129514: {
			Slug:        "test-tube",
			Character:   "🧪",
			UnicodeName: "test tube",
			CodePoint:   "1F9EA",
			Group:       "objects",
			SubGroup:    "science",
		},

		129515: {
			Slug:        "petri-dish",
			Character:   "🧫",
			UnicodeName: "petri dish",
			CodePoint:   "1F9EB",
			Group:       "objects",
			SubGroup:    "science",
		},

		129516: {
			Slug:        "dna",
			Character:   "🧬",
			UnicodeName: "dna",
			CodePoint:   "1F9EC",
			Group:       "objects",
			SubGroup:    "science",
		},

		129517: {
			Slug:        "e11-0-compass",
			Character:   "🧭",
			UnicodeName: "E11.0 compass",
			CodePoint:   "1F9ED",
			Group:       "travel-places",
			SubGroup:    "place-map",
		},

		129518: {
			Slug:        "abacus",
			Character:   "🧮",
			UnicodeName: "abacus",
			CodePoint:   "1F9EE",
			Group:       "objects",
			SubGroup:    "computer",
		},

		129519: {
			Slug:        "fire-extinguisher",
			Character:   "🧯",
			UnicodeName: "fire extinguisher",
			CodePoint:   "1F9EF",
			Group:       "objects",
			SubGroup:    "household",
		},

		129520: {
			Slug:        "toolbox",
			Character:   "🧰",
			UnicodeName: "toolbox",
			CodePoint:   "1F9F0",
			Group:       "objects",
			SubGroup:    "tool",
		},

		129521: {
			Slug:        "e11-0-brick",
			Character:   "🧱",
			UnicodeName: "E11.0 brick",
			CodePoint:   "1F9F1",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		129522: {
			Slug:        "magnet",
			Character:   "🧲",
			UnicodeName: "magnet",
			CodePoint:   "1F9F2",
			Group:       "objects",
			SubGroup:    "tool",
		},

		129523: {
			Slug:        "e11-0-luggage",
			Character:   "🧳",
			UnicodeName: "E11.0 luggage",
			CodePoint:   "1F9F3",
			Group:       "travel-places",
			SubGroup:    "hotel",
		},

		129524: {
			Slug:        "lotion-bottle",
			Character:   "🧴",
			UnicodeName: "lotion bottle",
			CodePoint:   "1F9F4",
			Group:       "objects",
			SubGroup:    "household",
		},

		129525: {
			Slug:        "thread",
			Character:   "🧵",
			UnicodeName: "thread",
			CodePoint:   "1F9F5",
			Group:       "activities",
			SubGroup:    "arts-crafts",
		},

		129526: {
			Slug:        "yarn",
			Character:   "🧶",
			UnicodeName: "yarn",
			CodePoint:   "1F9F6",
			Group:       "activities",
			SubGroup:    "arts-crafts",
		},

		129527: {
			Slug:        "safety-pin",
			Character:   "🧷",
			UnicodeName: "safety pin",
			CodePoint:   "1F9F7",
			Group:       "objects",
			SubGroup:    "household",
		},

		129528: {
			Slug:        "teddy-bear",
			Character:   "🧸",
			UnicodeName: "teddy bear",
			CodePoint:   "1F9F8",
			Group:       "activities",
			SubGroup:    "game",
		},

		129529: {
			Slug:        "broom",
			Character:   "🧹",
			UnicodeName: "broom",
			CodePoint:   "1F9F9",
			Group:       "objects",
			SubGroup:    "household",
		},

		129530: {
			Slug:        "basket",
			Character:   "🧺",
			UnicodeName: "basket",
			CodePoint:   "1F9FA",
			Group:       "objects",
			SubGroup:    "household",
		},

		129531: {
			Slug:        "roll-of-paper",
			Character:   "🧻",
			UnicodeName: "roll of paper",
			CodePoint:   "1F9FB",
			Group:       "objects",
			SubGroup:    "household",
		},

		129532: {
			Slug:        "soap",
			Character:   "🧼",
			UnicodeName: "soap",
			CodePoint:   "1F9FC",
			Group:       "objects",
			SubGroup:    "household",
		},

		129533: {
			Slug:        "sponge",
			Character:   "🧽",
			UnicodeName: "sponge",
			CodePoint:   "1F9FD",
			Group:       "objects",
			SubGroup:    "household",
		},

		129534: {
			Slug:        "receipt",
			Character:   "🧾",
			UnicodeName: "receipt",
			CodePoint:   "1F9FE",
			Group:       "objects",
			SubGroup:    "money",
		},

		129535: {
			Slug:        "nazar-amulet",
			Character:   "🧿",
			UnicodeName: "nazar amulet",
			CodePoint:   "1F9FF",
			Group:       "activities",
			SubGroup:    "game",
		},

		129648: {
			Slug:        "ballet-shoes",
			Character:   "🩰",
			UnicodeName: "ballet shoes",
			CodePoint:   "1FA70",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129649: {
			Slug:        "one-piece-swimsuit",
			Character:   "🩱",
			UnicodeName: "one-piece swimsuit",
			CodePoint:   "1FA71",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129650: {
			Slug:        "briefs",
			Character:   "🩲",
			UnicodeName: "briefs",
			CodePoint:   "1FA72",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129651: {
			Slug:        "shorts",
			Character:   "🩳",
			UnicodeName: "shorts",
			CodePoint:   "1FA73",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129652: {
			Slug:        "thong-sandal",
			Character:   "🩴",
			UnicodeName: "thong sandal",
			CodePoint:   "1FA74",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129656: {
			Slug:        "drop-of-blood",
			Character:   "🩸",
			UnicodeName: "drop of blood",
			CodePoint:   "1FA78",
			Group:       "objects",
			SubGroup:    "medical",
		},

		129657: {
			Slug:        "adhesive-bandage",
			Character:   "🩹",
			UnicodeName: "adhesive bandage",
			CodePoint:   "1FA79",
			Group:       "objects",
			SubGroup:    "medical",
		},

		129658: {
			Slug:        "stethoscope",
			Character:   "🩺",
			UnicodeName: "stethoscope",
			CodePoint:   "1FA7A",
			Group:       "objects",
			SubGroup:    "medical",
		},

		129664: {
			Slug:        "yo-yo",
			Character:   "🪀",
			UnicodeName: "yo-yo",
			CodePoint:   "1FA80",
			Group:       "activities",
			SubGroup:    "game",
		},

		129665: {
			Slug:        "kite",
			Character:   "🪁",
			UnicodeName: "kite",
			CodePoint:   "1FA81",
			Group:       "activities",
			SubGroup:    "game",
		},

		129666: {
			Slug:        "e12-0-parachute",
			Character:   "🪂",
			UnicodeName: "E12.0 parachute",
			CodePoint:   "1FA82",
			Group:       "travel-places",
			SubGroup:    "transport-air",
		},

		129667: {
			Slug:        "boomerang",
			Character:   "🪃",
			UnicodeName: "boomerang",
			CodePoint:   "1FA83",
			Group:       "objects",
			SubGroup:    "tool",
		},

		129668: {
			Slug:        "magic-wand",
			Character:   "🪄",
			UnicodeName: "magic wand",
			CodePoint:   "1FA84",
			Group:       "activities",
			SubGroup:    "game",
		},

		129669: {
			Slug:        "pinata",
			Character:   "🪅",
			UnicodeName: "piñata",
			CodePoint:   "1FA85",
			Group:       "activities",
			SubGroup:    "game",
		},

		129670: {
			Slug:        "nesting-dolls",
			Character:   "🪆",
			UnicodeName: "nesting dolls",
			CodePoint:   "1FA86",
			Group:       "activities",
			SubGroup:    "game",
		},

		129680: {
			Slug:        "ringed-planet",
			Character:   "🪐",
			UnicodeName: "ringed planet",
			CodePoint:   "1FA90",
			Group:       "travel-places",
			SubGroup:    "sky-weather",
		},

		129681: {
			Slug:        "chair",
			Character:   "🪑",
			UnicodeName: "chair",
			CodePoint:   "1FA91",
			Group:       "objects",
			SubGroup:    "household",
		},

		129682: {
			Slug:        "razor",
			Character:   "🪒",
			UnicodeName: "razor",
			CodePoint:   "1FA92",
			Group:       "objects",
			SubGroup:    "household",
		},

		129683: {
			Slug:        "axe",
			Character:   "🪓",
			UnicodeName: "axe",
			CodePoint:   "1FA93",
			Group:       "objects",
			SubGroup:    "tool",
		},

		129684: {
			Slug:        "diya-lamp",
			Character:   "🪔",
			UnicodeName: "diya lamp",
			CodePoint:   "1FA94",
			Group:       "objects",
			SubGroup:    "light-video",
		},

		129685: {
			Slug:        "banjo",
			Character:   "🪕",
			UnicodeName: "banjo",
			CodePoint:   "1FA95",
			Group:       "objects",
			SubGroup:    "musical-instrument",
		},

		129686: {
			Slug:        "military-helmet",
			Character:   "🪖",
			UnicodeName: "military helmet",
			CodePoint:   "1FA96",
			Group:       "objects",
			SubGroup:    "clothing",
		},

		129687: {
			Slug:        "accordion",
			Character:   "🪗",
			UnicodeName: "accordion",
			CodePoint:   "1FA97",
			Group:       "objects",
			SubGroup:    "musical-instrument",
		},

		129688: {
			Slug:        "long-drum",
			Character:   "🪘",
			UnicodeName: "long drum",
			CodePoint:   "1FA98",
			Group:       "objects",
			SubGroup:    "musical-instrument",
		},

		129689: {
			Slug:        "coin",
			Character:   "🪙",
			UnicodeName: "coin",
			CodePoint:   "1FA99",
			Group:       "objects",
			SubGroup:    "money",
		},

		129690: {
			Slug:        "carpentry-saw",
			Character:   "🪚",
			UnicodeName: "carpentry saw",
			CodePoint:   "1FA9A",
			Group:       "objects",
			SubGroup:    "tool",
		},

		129691: {
			Slug:        "screwdriver",
			Character:   "🪛",
			UnicodeName: "screwdriver",
			CodePoint:   "1FA9B",
			Group:       "objects",
			SubGroup:    "tool",
		},

		129692: {
			Slug:        "ladder",
			Character:   "🪜",
			UnicodeName: "ladder",
			CodePoint:   "1FA9C",
			Group:       "objects",
			SubGroup:    "tool",
		},

		129693: {
			Slug:        "hook",
			Character:   "🪝",
			UnicodeName: "hook",
			CodePoint:   "1FA9D",
			Group:       "objects",
			SubGroup:    "tool",
		},

		129694: {
			Slug:        "mirror",
			Character:   "🪞",
			UnicodeName: "mirror",
			CodePoint:   "1FA9E",
			Group:       "objects",
			SubGroup:    "household",
		},

		129695: {
			Slug:        "window",
			Character:   "🪟",
			UnicodeName: "window",
			CodePoint:   "1FA9F",
			Group:       "objects",
			SubGroup:    "household",
		},

		129696: {
			Slug:        "plunger",
			Character:   "🪠",
			UnicodeName: "plunger",
			CodePoint:   "1FAA0",
			Group:       "objects",
			SubGroup:    "household",
		},

		129697: {
			Slug:        "sewing-needle",
			Character:   "🪡",
			UnicodeName: "sewing needle",
			CodePoint:   "1FAA1",
			Group:       "activities",
			SubGroup:    "arts-crafts",
		},

		129698: {
			Slug:        "knot",
			Character:   "🪢",
			UnicodeName: "knot",
			CodePoint:   "1FAA2",
			Group:       "activities",
			SubGroup:    "arts-crafts",
		},

		129699: {
			Slug:        "bucket",
			Character:   "🪣",
			UnicodeName: "bucket",
			CodePoint:   "1FAA3",
			Group:       "objects",
			SubGroup:    "household",
		},

		129700: {
			Slug:        "mouse-trap",
			Character:   "🪤",
			UnicodeName: "mouse trap",
			CodePoint:   "1FAA4",
			Group:       "objects",
			SubGroup:    "household",
		},

		129701: {
			Slug:        "toothbrush",
			Character:   "🪥",
			UnicodeName: "toothbrush",
			CodePoint:   "1FAA5",
			Group:       "objects",
			SubGroup:    "household",
		},

		129702: {
			Slug:        "headstone",
			Character:   "🪦",
			UnicodeName: "headstone",
			CodePoint:   "1FAA6",
			Group:       "objects",
			SubGroup:    "other-object",
		},

		129703: {
			Slug:        "placard",
			Character:   "🪧",
			UnicodeName: "placard",
			CodePoint:   "1FAA7",
			Group:       "objects",
			SubGroup:    "other-object",
		},

		129704: {
			Slug:        "e13-0-rock",
			Character:   "🪨",
			UnicodeName: "E13.0 rock",
			CodePoint:   "1FAA8",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		129712: {
			Slug:        "e13-0-fly",
			Character:   "🪰",
			UnicodeName: "E13.0 fly",
			CodePoint:   "1FAB0",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		129713: {
			Slug:        "e13-0-worm",
			Character:   "🪱",
			UnicodeName: "E13.0 worm",
			CodePoint:   "1FAB1",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		129714: {
			Slug:        "e13-0-beetle",
			Character:   "🪲",
			UnicodeName: "E13.0 beetle",
			CodePoint:   "1FAB2",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		129715: {
			Slug:        "e13-0-cockroach",
			Character:   "🪳",
			UnicodeName: "E13.0 cockroach",
			CodePoint:   "1FAB3",
			Group:       "animals-nature",
			SubGroup:    "animal-bug",
		},

		129716: {
			Slug:        "e13-0-potted-plant",
			Character:   "🪴",
			UnicodeName: "E13.0 potted plant",
			CodePoint:   "1FAB4",
			Group:       "animals-nature",
			SubGroup:    "plant-other",
		},

		129717: {
			Slug:        "e13-0-wood",
			Character:   "🪵",
			UnicodeName: "E13.0 wood",
			CodePoint:   "1FAB5",
			Group:       "travel-places",
			SubGroup:    "place-building",
		},

		129718: {
			Slug:        "e13-0-feather",
			Character:   "🪶",
			UnicodeName: "E13.0 feather",
			CodePoint:   "1FAB6",
			Group:       "animals-nature",
			SubGroup:    "animal-bird",
		},

		129728: {
			Slug:        "e13-0-anatomical-heart",
			Character:   "🫀",
			UnicodeName: "E13.0 anatomical heart",
			CodePoint:   "1FAC0",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		129729: {
			Slug:        "e13-0-lungs",
			Character:   "🫁",
			UnicodeName: "E13.0 lungs",
			CodePoint:   "1FAC1",
			Group:       "people-body",
			SubGroup:    "body-parts",
		},

		129730: {
			Slug:        "e13-0-people-hugging",
			Character:   "🫂",
			UnicodeName: "E13.0 people hugging",
			CodePoint:   "1FAC2",
			Group:       "people-body",
			SubGroup:    "person-symbol",
		},

		129744: {
			Slug:        "e13-0-blueberries",
			Character:   "🫐",
			UnicodeName: "E13.0 blueberries",
			CodePoint:   "1FAD0",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		129745: {
			Slug:        "e13-0-bell-pepper",
			Character:   "🫑",
			UnicodeName: "E13.0 bell pepper",
			CodePoint:   "1FAD1",
			Group:       "food-drink",
			SubGroup:    "food-vegetable",
		},

		129746: {
			Slug:        "e13-0-olive",
			Character:   "🫒",
			UnicodeName: "E13.0 olive",
			CodePoint:   "1FAD2",
			Group:       "food-drink",
			SubGroup:    "food-fruit",
		},

		129747: {
			Slug:        "e13-0-flatbread",
			Character:   "🫓",
			UnicodeName: "E13.0 flatbread",
			CodePoint:   "1FAD3",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129748: {
			Slug:        "e13-0-tamale",
			Character:   "🫔",
			UnicodeName: "E13.0 tamale",
			CodePoint:   "1FAD4",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129749: {
			Slug:        "e13-0-fondue",
			Character:   "🫕",
			UnicodeName: "E13.0 fondue",
			CodePoint:   "1FAD5",
			Group:       "food-drink",
			SubGroup:    "food-prepared",
		},

		129750: {
			Slug:        "e13-0-teapot",
			Character:   "🫖",
			UnicodeName: "E13.0 teapot",
			CodePoint:   "1FAD6",
			Group:       "food-drink",
			SubGroup:    "drink",
		},
	}
)
