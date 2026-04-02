// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/hyperspell-cli/internal/mocktest"
)

func TestIntegrationsSlackList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"integrations:slack", "list",
			"--channel", "string",
			"--exclude-archived=true",
			"--include-dms=true",
			"--include-group-dms=true",
			"--include-private=true",
		)
	})
}
