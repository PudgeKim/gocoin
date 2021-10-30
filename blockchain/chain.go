package blockchain

import (
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newest_hash"`
	Height     int    `json:"height"`
}

// for singleton
var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

func BlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			// 위에서 var로 선언했기 때문에 :=가 아닌 =
			b = &blockchain{"", 0}
			b.AddBlock("Genesis")
		})
	}
	return b
}
