// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authz

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testProviders map[string]*schema.Provider
var testProvider *schema.Provider

func init() {
	testProvider = Provider()
	testProviders = map[string]*schema.Provider{
		"authz": testProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("AUTHZ_HOST"); err == "" {
		t.Fatal("AUTHZ_HOST must be set for acceptance tests")
	}
	if err := os.Getenv("AUTHZ_CLIENT_SECRET"); err == "" {
		t.Fatal("AUTHZ_CLIENT_SECRET must be set for acceptance tests")
	}
}
