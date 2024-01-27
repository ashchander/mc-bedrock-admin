package main

import (
	"strings"
	"testing"
)

const SAMPLE_INPUT = `Jan 26 23:36:34 chander-server run_server.sh[820]: [2024-01-26 23:36:34:194 INFO] commandblockoutput = true, dodaylightcycle = true, doentitydrops = true`

func TestParseConfig(t *testing.T) {
	output, err := parseConfig(SAMPLE_INPUT)

	// TODO look for correct output
	if output != "{\"commandblockoutput\":true, \"dodaylightcycle\":true, \"doentitydrops\":true}" || err != nil {
		t.Fatalf("Failed to parse config: %s", output)
	}
}

func TestParseConfig_ReturnErrorForEmptyConfig(t *testing.T) {
	output, err := parseConfig("")

	if output != "" || err == nil {
		t.Fatal("Failed to return error for empty config")
	}
	if !strings.Contains(err.Error(), "Failed to parse config - No config provided") {
		t.Fatal("Returned wrong error")
	}
}

func TestParseConfig_ReturnErrorForInvalidConfig(t *testing.T) {
	output, err := parseConfig("INFO] ")

	if output != "" || err == nil {
		t.Fatal("Failed to return error for empty config")
	}
	if !strings.Contains(err.Error(), "Failed to parse config - Invalid config provided") {
		t.Fatal("Returned wrong error")
	}
}
