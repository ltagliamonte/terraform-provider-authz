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
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AUTHZ_CLIENT_ID", nil),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AUTHZ_CLIENT_SECRET", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"authz_policy":    policy(),
			"authz_resource":  resource(),
			"authz_principal": principal(),
			"authz_role":      role(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"authz_policy":    dsPolicy(),
			"authz_principal": dsPrincipal(),
			"authz_resource":  dsResource(),
			"authz_role":      dsRole(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	cID := d.Get("client_id").(string)
	cSecret := d.Get("client_secret").(string)

	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (cID != "") && (cSecret != "") {
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
			ClientId:     cID,
			ClientSecret: cSecret,
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
		Detail:   "Empty client_id and/or client_secret",
	})

	return nil, diags
}
