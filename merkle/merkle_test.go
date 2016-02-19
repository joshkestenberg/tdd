package main

import (
	"testing"
)

type example struct {
	elements []int
	root     []int
}

var examples = []example{
	example{[]int{3, 4}, 0},
	example{[]int{3, 4, 9, 6}, 1},
	example{[]int{5, 2, 3, 1, 6, 7, 1, 1}, 5},
}

func TestMerkleRoot(t *testing.T) {
	for _, e := range examples {
		elems := e.elements
		root := e.root
		root_ := MerkleRoot(elems)
		if root != root_ {
			t.Fatalf("Got %v; Expected %v", root_, root)
		}
	}
}

func TestMerkleProof(t *testing.T) {
	for _, e := range examples {
		elems := e.elements
		root := e.root

		for i, el := range elems {
			proof := MerkleProof(elems, i)
			legit := MerkleVerify(root, el, proof)
			if !legit {
				t.Fatal("Merkle proof failed to verify")
			}
		}
	}
}
