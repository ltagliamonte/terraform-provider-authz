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

func principal() *schema.Resource {
	return &schema.Resource{
		CreateContext: principalCreate,
		ReadContext:   principalRead,
		UpdateContext: principalUpdate,
		DeleteContext: principalDelete,
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
			"roles":      tfSet,
			"attributes": tfMap,
		},
	}
}

func principalCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pName := d.Get("name").(string)
	pRoles := d.Get("roles").(*schema.Set)
	pAttributes := d.Get("attributes").(map[string]interface{})

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	resp, err := ac.client.PrincipalCreate(outContext, &authz.PrincipalCreateRequest{
		Id:         pName,
		Roles:      getSet(pRoles),
		Attributes: getAttributesList(pAttributes),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Principal.Id)
	principalRead(ctx, d, m)

	return diags
}

func principalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pName := d.Get("name").(string)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	resp, err := ac.client.PrincipalGet(outContext, &authz.PrincipalGetRequest{
		Id: pName,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	principal := resp.Principal
	if err := d.Set("name", principal.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("roles", principal.Roles); err != nil {
		return diag.FromErr(err)
	}
	attributes := flattenAttributesItems(principal.Attributes)
	if err := d.Set("attributes", attributes); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func principalUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	pName := d.Get("name").(string)
	pRoles := d.Get("roles").(*schema.Set)
	pAttributes := d.Get("attributes").(map[string]interface{})

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	_, err := ac.client.PrincipalUpdate(outContext, &authz.PrincipalUpdateRequest{
		Id:         pName,
		Roles:      getSet(pRoles),
		Attributes: getAttributesList(pAttributes),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return principalRead(ctx, d, m)
}

func principalDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pName := d.Get("name").(string)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	_, err := ac.client.PrincipalDelete(outContext, &authz.PrincipalDeleteRequest{
		Id: pName,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
