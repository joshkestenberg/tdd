package main

// Our hash function.
func Hash(x, y int) int {
	return (x + y) % 7
}

// Takes a list of elems and returns the merkle root using Hash
func MerkleRoot(elems []int) int {
	// TODO: implement
}

// Return a proof that the index'th element is in the merkle tree
func MerkleProof(elems []int, index int) (proof []int) {
	// TODO: implement
}

// Verify the proof
func MerkleVerify(root int, elem int, proof []int) bool {
	// TODO: implement
}
