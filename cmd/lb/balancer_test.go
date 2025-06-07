package main

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	addr := "192.168.1.1:12345"
	hash1 := hash(addr)
	hash2 := hash(addr)

	if hash1 != hash2 {
		t.Errorf("Hash function should be consistent. Got %d and %d for same input", hash1, hash2)
	}

	addr2 := "192.168.1.2:54321"
	hash3 := hash(addr2)

	if hash1 == hash3 {
		t.Logf("Note: Different addresses produced same hash (collision): %s and %s", addr, addr2)
	}

	emptyHash := hash("")
	if emptyHash == 0 {
		t.Logf("Empty string hash is zero")
	}
}

func TestChooseServer(t *testing.T) {
	servers := []string{"server1:8080", "server2:8080", "server3:8080"}

	result := chooseServer("192.168.1.1:12345", []string{})
	if result != "" {
		t.Errorf("Expected empty string for empty server pool, got %s", result)
	}

	clientAddr := "192.168.1.1:12345"
	server1 := chooseServer(clientAddr, servers)
	server2 := chooseServer(clientAddr, servers)

	if server1 != server2 {
		t.Errorf("Same client should always get same server. Got %s and %s", server1, server2)
	}
