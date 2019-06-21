package pagehash

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := Get("muzzammil.xyz/?pagehash-test")
	if err != nil {
		t.Fatalf("Expected nil, got %+v", err)
	}
	for _, hash := range resp.Hashes {
		fmt.Printf("%s:\t%s\n", hash.Algorithm, hash.Hash)
	}
}
