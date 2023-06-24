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

func policy() *schema.Resource {
	return &schema.Resource{
		CreateContext: policyCreate,
		ReadContext:   policyRead,
		UpdateContext: policyUpdate,
		DeleteContext: policyDelete,
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
			"resources":       tfSet,
			"actions":         tfSet,
			"attribute_rules": tfSet,
		},
	}
}

func policyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pName := d.Get("name").(string)
	pResources := d.Get("resources").(*schema.Set)
	pActions := d.Get("actions").(*schema.Set)
	pAttribute := d.Get("attribute_rules").(*schema.Set)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	resp, err := ac.client.PolicyCreate(outContext, &authz.PolicyCreateRequest{
		Id:             pName,
		Actions:        getSet(pActions),
		Resources:      getSet(pResources),
		AttributeRules: getSet(pAttribute),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Policy.Id)
	policyRead(ctx, d, m)

	return diags
}

func policyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pName := d.Get("name").(string)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	resp, err := ac.client.PolicyGet(outContext, &authz.PolicyGetRequest{
		Id: pName,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	policy := resp.Policy
	if err := d.Set("name", policy.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resources", policy.Resources); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("actions", policy.Actions); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("attribute_rules", policy.AttributeRules); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func policyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	pName := d.Get("name").(string)
	pResources := d.Get("resources").(*schema.Set)
	pActions := d.Get("actions").(*schema.Set)
	pAttribute := d.Get("attribute_rules").(*schema.Set)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	_, err := ac.client.PolicyUpdate(outContext, &authz.PolicyUpdateRequest{
		Id:             pName,
		Actions:        getSet(pActions),
		Resources:      getSet(pResources),
		AttributeRules: getSet(pAttribute),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return policyRead(ctx, d, m)
}

func policyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pName := d.Get("name").(string)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	_, err := ac.client.PolicyDelete(outContext, &authz.PolicyDeleteRequest{
		Id: pName,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
