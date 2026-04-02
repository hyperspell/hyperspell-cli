// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/hyperspell-cli/internal/mocktest"
)

func TestEvaluateGetQuery(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"evaluate", "get-query",
			"--query-id", "query_id",
		)
	})
}

func TestEvaluateScoreHighlight(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"evaluate", "score-highlight",
			"--highlight-id", "highlight_id",
			"--comment", "comment",
			"--score", "-1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"comment: comment\n" +
			"score: -1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"evaluate", "score-highlight",
			"--highlight-id", "highlight_id",
		)
	})
}

func TestEvaluateScoreQuery(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"evaluate", "score-query",
			"--query-id", "query_id",
			"--score", "-1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("score: -1")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"evaluate", "score-query",
			"--query-id", "query_id",
		)
	})
}
