// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/hyperspell/hyperspell-cli/internal/mocktest"
)

func TestLiveGetResource(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"live", "get-resource",
			"--source", "reddit",
			"--resource-id", "resource_id",
			"--connection-id", "connection_id",
			"--index=true",
		)
	})
}

func TestLiveListResources(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"live", "list-resources",
			"--max-items", "10",
			"--source", "reddit",
			"--connection-id", "connection_id",
			"--cursor", "cursor",
			"--size", "0",
		)
	})
}

func TestLiveListSources(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"live", "list-sources",
		)
	})
}

func TestLiveSearch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"live", "search",
			"--source", "reddit",
			"--query", "query",
			"--connection-id", "connection_id",
			"--index=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: query\n" +
			"connection_id: connection_id\n" +
			"index: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"live", "search",
			"--source", "reddit",
		)
	})
}
