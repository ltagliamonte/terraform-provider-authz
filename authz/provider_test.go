// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authz

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"authz": testAccProvider,
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
	if err := os.Getenv("AUTHZ_PASSWORD"); err == "" {
		t.Fatal("AUTHZ_PASSWORD must be set for acceptance tests")
	}
}
