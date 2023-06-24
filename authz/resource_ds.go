package authz

import (
	"context"

	authz "github.com/eko/authz/backend/pkg/authz"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/grpc/metadata"
)

func dsResource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dsResourceRead,
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
				Required: false,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
			"attributes": computedTfMap,
		},
	}
}

func dsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	d.SetId(resp.Resource.Id)
	return diags
}
