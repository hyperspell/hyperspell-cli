// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/hyperspell/hyperspell-cli/internal/mocktest"
)

func TestActionsAddReaction(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"actions", "add-reaction",
			"--channel", "channel",
			"--name", "name",
			"--provider", "slack",
			"--timestamp", "timestamp",
			"--connection", "connection",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"channel: channel\n" +
			"name: name\n" +
			"provider: slack\n" +
			"timestamp: timestamp\n" +
			"connection: connection\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"actions", "add-reaction",
		)
	})
}

func TestActionsSendMessage(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"actions", "send-message",
			"--provider", "slack",
			"--text", "text",
			"--channel", "channel",
			"--connection", "connection",
			"--parent", "parent",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"provider: slack\n" +
			"text: text\n" +
			"channel: channel\n" +
			"connection: connection\n" +
			"parent: parent\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--user-id", "string",
			"actions", "send-message",
		)
	})
}
