package ankadb

import (
	"testing"
)

func Test_LoadConfig(t *testing.T) {
	cfg, err := LoadConfig("./test/test.yaml")
	if err != nil {
		t.Fatalf("Test_LoadConfig err %v", err)

		return
	}

	if cfg == nil {
		t.Fatalf("Test_LoadConfig cfg is nil")

		return
	}

	if cfg.AddrGRPC != "0.0.0.0:7788" {
		t.Fatalf("Test_LoadConfig invalid AddrGRPC %v", cfg.AddrGRPC)
	}

	if cfg.AddrHTTP != "0.0.0.0:7789" {
		t.Fatalf("Test_LoadConfig invalid AddrHTTP %v", cfg.AddrHTTP)
	}

	if cfg.PathDBRoot != "./test/dat" {
		t.Fatalf("Test_LoadConfig invalid PathDBRoot %v", cfg.PathDBRoot)
	}

	if cfg.ListDB == nil {
		t.Fatalf("Test_LoadConfig ListDB nil")
	}

	if len(cfg.ListDB) != 2 {
		t.Fatalf("Test_LoadConfig invalid ListDB length %v", len(cfg.ListDB))
	}

	if cfg.ListDB[0].Engine != "leveldb" || cfg.ListDB[0].Name != "test001" || cfg.ListDB[0].PathDB != "test001" {
		t.Fatalf("Test_LoadConfig invalid ListDB[0] %v", cfg.ListDB[0])
	}

	if cfg.ListDB[1].Engine != "leveldb" || cfg.ListDB[1].Name != "test002" || cfg.ListDB[1].PathDB != "test002" {
		t.Fatalf("Test_LoadConfig invalid ListDB[1] %v", cfg.ListDB[1])
	}

	t.Logf("Test_LoadConfig OK")
}
