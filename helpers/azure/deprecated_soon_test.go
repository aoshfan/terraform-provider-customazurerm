// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azure_test

import (
	"testing"

	"github.com/aoshfan/terraform-provider-customazurerm/helpers/azure"
)

func TestHelper_AzureResourceID(t *testing.T) {
	cases := []struct {
		ID     string
		Errors int
	}{
		{
			ID:     "",
			Errors: 1,
		},
		{
			ID:     "nonsense",
			Errors: 1,
		},
		{
			ID:     "/slash",
			Errors: 1,
		},
		{
			ID:     "/path/to/nothing",
			Errors: 1,
		},
		{
			ID:     "/subscriptions",
			Errors: 1,
		},
		{
			ID:     "/providers",
			Errors: 1,
		},
		{
			ID:     "/subscriptions/not-a-guid",
			Errors: 0,
		},
		{
			ID:     "/providers/test",
			Errors: 0,
		},
		{
			ID:     "/subscriptions/00000000-0000-0000-0000-00000000000/",
			Errors: 0,
		},
		{
			ID:     "/providers/provider.name/",
			Errors: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.ID, func(t *testing.T) {
			_, errors := azure.ValidateResourceID(tc.ID, "test")

			if len(errors) < tc.Errors {
				t.Fatalf("Expected ValidateResourceID to have %d not %d errors for %q", tc.Errors, len(errors), tc.ID)
			}
		})
	}
}

func TestAzureResourceIDOrEmpty(t *testing.T) {
	cases := []struct {
		ID     string
		Errors int
	}{
		{
			ID:     "",
			Errors: 0,
		},
		{
			ID:     "nonsense",
			Errors: 1,
		},
		// as this function just calls TestAzureResourceId lets not be as comprehensive
		{
			ID:     "/providers/provider.name/",
			Errors: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.ID, func(t *testing.T) {
			_, errors := azure.ValidateResourceIDOrEmpty(tc.ID, "test") // nolint: staticcheck

			if len(errors) < tc.Errors {
				t.Fatalf("Expected TestAzureResourceIdOrEmpty to have %d not %d errors for %q", tc.Errors, len(errors), tc.ID)
			}
		})
	}
}
