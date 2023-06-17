// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authz

import (
	"context"

	authz "github.com/eko/authz/backend/pkg/authz"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	listOfStrings = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
)

func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePolicyCreate,
		ReadContext:   resourcePolicyRead,
		UpdateContext: resourcePolicyUpdate,
		DeleteContext: resourcePolicyDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type: schema.TypeString,
			},
			"resources":       listOfStrings,
			"actions":         listOfStrings,
			"attribute_rules": listOfStrings,
		},
	}
}

func resourcePolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(*AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pID := d.Get("id").(string)
	pResources := d.Get("resources").([]string)
	pActions := d.Get("actions").([]string)
	pAttribute := d.Get("attribute_rules").([]string)

	_, err := ac.client.PolicyCreate(context.Background(), &authz.PolicyCreateRequest{
		Id:             pID,
		Actions:        pActions,
		Resources:      pResources,
		AttributeRules: pAttribute,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	resourcePolicyRead(ctx, d, m)

	return diags
}

func resourcePolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(*AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pID := d.Get("id").(string)

	resp, err := ac.client.PolicyGet(context.Background(), &authz.PolicyGetRequest{
		Id: pID,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	policy := resp.Policy
	if err := d.Set("id", policy.Id); err != nil {
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

func resourcePolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(*AuthzClient)

	pID := d.Get("id").(string)
	pResources := d.Get("resources").([]string)
	pActions := d.Get("actions").([]string)
	pAttribute := d.Get("attribute_rules").([]string)

	ac.client.PolicyUpdate(context.Background(), &authz.PolicyUpdateRequest{
		Id:             pID,
		Actions:        pActions,
		Resources:      pResources,
		AttributeRules: pAttribute,
	})

	return resourcePolicyRead(ctx, d, m)
}

func resourcePolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(*AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pID := d.Get("id").(string)

	_, err := ac.client.PolicyDelete(context.Background(), &authz.PolicyDeleteRequest{
		Id: pID,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
