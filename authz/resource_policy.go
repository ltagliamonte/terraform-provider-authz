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

var (
	listOfStrings = &schema.Schema{
		Type:     schema.TypeList,
		Computed: false,
		Required: true,
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
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"resources":       listOfStrings,
			"actions":         listOfStrings,
			"attribute_rules": listOfStrings,
		},
	}
}

func resourcePolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pName := d.Get("name").(string)
	pResources := d.Get("resources").([]interface{})
	pActions := d.Get("actions").([]interface{})
	pAttribute := d.Get("attribute_rules").([]interface{})

	rList := getStringListFromInterface(pResources)
	aList := getStringListFromInterface(pActions)
	atList := getStringListFromInterface(pAttribute)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	resp, err := ac.client.PolicyCreate(outContext, &authz.PolicyCreateRequest{
		Id:             pName,
		Actions:        aList,
		Resources:      rList,
		AttributeRules: atList,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Policy.Id)
	resourcePolicyRead(ctx, d, m)

	return diags
}

func getStringListFromInterface(list []interface{}) []string {
	var strList []string
	for _, k := range list {
		strList = append(strList, k.(string))
	}
	return strList
}

func resourcePolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourcePolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	pName := d.Get("name").(string)
	pResources := d.Get("resources").([]interface{})
	pActions := d.Get("actions").([]interface{})
	pAttribute := d.Get("attribute_rules").([]interface{})

	rList := getStringListFromInterface(pResources)
	aList := getStringListFromInterface(pActions)
	atList := getStringListFromInterface(pAttribute)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	ac.client.PolicyUpdate(outContext, &authz.PolicyUpdateRequest{
		Id:             pName,
		Actions:        aList,
		Resources:      rList,
		AttributeRules: atList,
	})

	return resourcePolicyRead(ctx, d, m)
}

func resourcePolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
