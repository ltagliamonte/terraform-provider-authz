// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authz

import (
	"context"

	authz "github.com/eko/authz/backend/pkg/authz"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/grpc/metadata"
)

func role() *schema.Resource {
	return &schema.Resource{
		CreateContext: roleCreate,
		ReadContext:   roleRead,
		UpdateContext: roleUpdate,
		DeleteContext: roleDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"policies": tfSet,
		},
	}
}

func roleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	rName := d.Get("name").(string)
	rPolicies := d.Get("policies").(*schema.Set)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	resp, err := ac.client.RoleCreate(outContext, &authz.RoleCreateRequest{
		Id:       rName,
		Policies: getSet(rPolicies),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Role.Id)
	roleRead(ctx, d, m)

	return diags
}

func roleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	rName := d.Get("name").(string)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	resp, err := ac.client.RoleGet(outContext, &authz.RoleGetRequest{
		Id: rName,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", resp.Role.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("policies", resp.Role.Policies); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func roleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	rName := d.Get("name").(string)
	rPolicies := d.Get("policies").(*schema.Set)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	_, err := ac.client.RoleUpdate(outContext, &authz.RoleUpdateRequest{
		Id:       rName,
		Policies: getSet(rPolicies),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return roleRead(ctx, d, m)
}

func roleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	rName := d.Get("name").(string)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	_, err := ac.client.RoleDelete(outContext, &authz.RoleDeleteRequest{
		Id: rName,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
