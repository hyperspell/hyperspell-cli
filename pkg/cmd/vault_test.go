// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/hyperspell/hyperspell-cli/internal/mocktest"
)

func TestVaultsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--user-id", "string",
			"vaults", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--size", "0",
		)
	})
}
