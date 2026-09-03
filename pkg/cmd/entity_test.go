// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/hyperspell/hyperspell-cli/internal/mocktest"
)

func TestEntitiesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"entities", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "1",
			"--min-supporting-documents", "0",
			"--name-prefix", "name_prefix",
			"--search", "search",
			"--sort-by", "id",
			"--sort-dir", "asc",
			"--status", "provisional",
			"--type", "type",
		)
	})
}

func TestEntitiesGet(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"entities", "get",
			"--entity-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestEntitiesSearch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"entities", "search",
			"--query", "query",
			"--limit", "1",
			"--type", "type",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: query\n" +
			"limit: 1\n" +
			"type: type\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"entities", "search",
		)
	})
}
