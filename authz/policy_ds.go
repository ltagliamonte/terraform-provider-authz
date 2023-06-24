package authz

import (
	"context"

	authz "github.com/eko/authz/backend/pkg/authz"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/grpc/metadata"
)

func dsPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dsPolicyRead,
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
			"resources":       computedTfSet,
			"actions":         computedTfSet,
			"attribute_rules": computedTfSet,
		},
	}
}

func dsPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	if err := d.Set("name", resp.Policy.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resources", resp.Policy.Resources); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("actions", resp.Policy.Actions); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("attribute_rules", resp.Policy.AttributeRules); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Policy.Id)
	return diags
}
