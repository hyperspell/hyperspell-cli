// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/hyperspell-cli/internal/mocktest"
)

func TestSessionsAdd(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"sessions", "add",
			"--history", "history",
			"--date", "'2019-12-27T18:11:19.117Z'",
			"--extract", "procedure",
			"--format", "vercel",
			"--metadata", "{foo: string}",
			"--session-id", "session_id",
			"--title", "title",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"history: history\n" +
			"date: '2019-12-27T18:11:19.117Z'\n" +
			"extract:\n" +
			"  - procedure\n" +
			"format: vercel\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"session_id: session_id\n" +
			"title: title\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"sessions", "add",
		)
	})
}
