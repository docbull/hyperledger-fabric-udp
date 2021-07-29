package rtfountain

import (
	fountain "gofountain"
)

// decoder decodes raptor symbols into
func decoder(encSymbols []fountain.LTBlock, codec fountain.Codec, dec fountain.Decoder) []byte {
	return dec.Decode()
}
