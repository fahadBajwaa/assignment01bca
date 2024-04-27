# Blockchain Implementation in Go

This repository contains a simple implementation of a blockchain in Go, created as part of an assignment.

## Author
- **Fahad Ashraf Bajwa**
- Student ID: 20i-2349

## Description
The project consists of two main files:
- `blockchain.go`: Defines the blockchain data structure and related functions.
- `main.go`: Demonstrates the usage of the blockchain implementation by creating a blockchain and adding blocks to it.

## Usage
To use this implementation, follow these steps:
1. Clone the repository.
2. Navigate to the directory containing the code.
3. Run the `main.go` file:  go run main.go
The output will display the details of each block in the blockchain and indicate whether the blockchain is valid or not.
Blockchain Structure

The blockchain consists of blocks, where each block contains the following information:

Transaction: Details of the transaction stored in the block.
Nonce: A random number used in the proof-of-work algorithm.
Previous Hash: Hash of the previous block in the chain.
Current Hash: Hash of the current block.
Functions

NewBlock
Creates a new block and adds it to the blockchain.

DisplayBlocks
Prints all the blocks in the blockchain.

CalculateHash
Calculates the hash of a block using the SHA-256 algorithm.

VerifyChain
Verifies the integrity of the blockchain by checking the hash links between blocks.



## Credits

This project was created by Fahad Bajwa.
