// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authz

import (
	"context"
	"strconv"

	hc "github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePolicyRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coffee_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_teaser": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_price": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_image": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quantity": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	orderID := strconv.Itoa(d.Get("id").(int))

	order, err := c.GetOrder(orderID, &c.Token)
	if err != nil {
		return diag.FromErr(err)
	}

	orderItems := flattenOrderItemsData(&order.Items)
	if err := d.Set("items", orderItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(orderID)

	return diags
}

func flattenOrderItemsData(orderItems *[]hc.OrderItem) []interface{} {
	if orderItems != nil {
		ois := make([]interface{}, len(*orderItems), len(*orderItems))

		for i, orderItem := range *orderItems {
			oi := make(map[string]interface{})

			oi["coffee_id"] = orderItem.Coffee.ID
			oi["coffee_name"] = orderItem.Coffee.Name
			oi["coffee_teaser"] = orderItem.Coffee.Teaser
			oi["coffee_description"] = orderItem.Coffee.Description
			oi["coffee_price"] = orderItem.Coffee.Price
			oi["coffee_image"] = orderItem.Coffee.Image
			oi["quantity"] = orderItem.Quantity

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}
