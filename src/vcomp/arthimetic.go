package vcomp

const Code_valu_bits byte = 8
const MAX_FREQ uint16 = 64

type SymbolStats struct {
	symbol byte
	freq   uint16
	high   uint16
	low    uint16
}

func Acompress(input []byte) {
	buildModel(input)
}

func buildModel(input []byte) {
	//	symDict := Count(input)

}

func scalRanges() {

}
