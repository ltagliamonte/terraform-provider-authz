package authz

import (
	"context"

	authz "github.com/eko/authz/backend/pkg/authz"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/grpc/metadata"
)

func dsRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dsRoleRead,
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
			"policies": computedTfSet,
		},
	}
}

func dsRoleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	d.SetId(resp.Role.Id)
	return diags
}
