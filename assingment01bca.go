// Fahad Ashraf Bajwa
// 20i-2349
// Assignment 1

package assingment01bca

import (
	"crypto/sha256"
	"fmt"
)

// Block represents a block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

// Blockchain represents a blockchain.
type Blockchain struct {
	Blocks []*Block
}

// NewBlock creates a new block and adds it to the blockchain.
func (bc *Blockchain) NewBlock(transaction string, nonce int, previousHash string) {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CurrentHash = CalculateHash(block)
	bc.Blocks = append(bc.Blocks, block)
}

// DisplayBlocks prints all the blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.CurrentHash)
		fmt.Println()
	}
}

// CalculateHash calculates the hash of a block.
func CalculateHash(block *Block) string {
	data := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
