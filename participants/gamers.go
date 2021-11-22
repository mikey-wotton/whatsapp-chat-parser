package participants

//go:generate stringer -type=Gamer -output=gen_gamer.go -trimprefix=Gamer
type Gamer int

const (
	GamerUnknown Gamer = iota //Un
	GamerMike
	GamerPaul
	GamerDayle
	GamerNoone
	GamerLockers
	GamerKirky
	GamerWezz
	GamerPiper
	GamerJames
	GamerDerren
	GamerNathan
	GamerQck
	GamerSantosh
	GamerTristan
	GamerAlex
	GamerLee
	GamerTapan
	GamerGemima
	GamerBayley
)

var Gamers = []Gamer{
	GamerMike,
	GamerPaul,
	GamerDayle,
	GamerNoone,
	GamerLockers,
	GamerKirky,
	GamerWezz,
	GamerPiper,
	GamerJames,
	GamerDerren,
	GamerNathan,
	GamerQck,
	GamerSantosh,
	GamerTristan,
	GamerAlex,
	GamerLee,
	GamerTapan,
	GamerGemima,
	GamerBayley,
}
