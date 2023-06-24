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

func resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
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
			"kind": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"attributes": tfMap,
		},
	}
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	rName := d.Get("name").(string)
	rKind := d.Get("kind").(string)
	rValue := d.Get("value").(string)
	rAttributes := d.Get("attributes").(map[string]interface{})

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	resp, err := ac.client.ResourceCreate(outContext, &authz.ResourceCreateRequest{
		Id:         rName,
		Kind:       rKind,
		Value:      rValue,
		Attributes: getAttributesList(rAttributes),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Resource.Id)
	resourceRead(ctx, d, m)

	return diags
}

func resourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	rName := d.Get("name").(string)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	resp, err := ac.client.ResourceGet(outContext, &authz.ResourceGetRequest{
		Id: rName,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", resp.Resource.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("kind", resp.Resource.Kind); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("value", resp.Resource.Value); err != nil {
		return diag.FromErr(err)
	}
	attributes := flattenAttributesItems(resp.Resource.Attributes)
	if err := d.Set("attributes", attributes); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	rName := d.Get("name").(string)
	rKind := d.Get("kind").(string)
	rValue := d.Get("value").(string)
	rAttributes := d.Get("attributes").(map[string]interface{})

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	_, err := ac.client.ResourceUpdate(outContext, &authz.ResourceUpdateRequest{
		Id:         rName,
		Kind:       rKind,
		Value:      rValue,
		Attributes: getAttributesList(rAttributes),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceRead(ctx, d, m)
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(AuthzClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	rName := d.Get("name").(string)

	outContext := metadata.NewOutgoingContext(ctx, ac.md)
	_, err := ac.client.ResourceDelete(outContext, &authz.ResourceDeleteRequest{
		Id: rName,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
