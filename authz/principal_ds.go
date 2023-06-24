package authz

import (
	"context"

	authz "github.com/eko/authz/backend/pkg/authz"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/grpc/metadata"
)

func dsPrincipal() *schema.Resource {
	return &schema.Resource{
		ReadContext: dsPrincipalRead,
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
			"roles":      computedTfSet,
			"attributes": computedTfMap,
		},
	}
}

func dsPrincipalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	if err := d.Set("name", resp.Principal.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("roles", resp.Principal.Roles); err != nil {
		return diag.FromErr(err)
	}
	attributes := flattenAttributesItems(resp.Principal.Attributes)
	if err := d.Set("attributes", attributes); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.Principal.Id)
	return diags
}
