// Fahad Ashraf Bajwa
// 20i-2349
// Assignment 1

package main

import (
	"fmt"

	a1 "github.com/fahadBajwaa/assignment01bca"
)

func main() {
	blockchain := a1.NewBlockchain()

	blockchain.NewBlock("Genesis Block", 0, "")

	// Add more blocks
	blockchain.NewBlock("Bob to Alice", 1, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash)
	blockchain.NewBlock("Alice to Bob", 2, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash)

	blockchain.DisplayBlocks()

	if blockchain.VerifyChain() {
		fmt.Println("The blockchain is valid.")
	} else {
		fmt.Println("The blockchain is not valid.")
	}
}
