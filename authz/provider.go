// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authz

import (
	"context"
	"fmt"

	authz "github.com/eko/authz/backend/pkg/authz"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type AuthzClient struct {
	client authz.ApiClient
	md     metadata.MD
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AUTHZ_HOST", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AUTHZ_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AUTHZ_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"authz_policy":    policy(),
			"authz_resource":  resource(),
			"authz_principal": principal(),
			"authz_role":      role(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		conn, err := grpc.Dial(*host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failed to connect: %v",
			})
			return nil, diags
		}
		c := authz.NewApiClient(conn)
		resp, err := c.Authenticate(ctx, &authz.AuthenticateRequest{
			ClientId:     username,
			ClientSecret: password,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Error(),
			})

			return nil, diags
		}
		return AuthzClient{
			client: c,
			md: metadata.New(map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", resp.Token),
			}),
		}, diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create Authz client",
		Detail:   "Empty username and/or password",
	})

	return nil, diags
}
