package authz

import (
	authz "github.com/eko/authz/backend/pkg/authz"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	tfList = &schema.Schema{
		Type:     schema.TypeList,
		Computed: false,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	tfMap = &schema.Schema{
		Type:     schema.TypeMap,
		Computed: false,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	}
)

func getStrList(list []interface{}) []string {
	var strList []string
	for _, k := range list {
		strList = append(strList, k.(string))
	}
	return strList
}

func getAttributesList(val map[string]interface{}) []*authz.Attribute {
	attributesList := make([]*authz.Attribute, 0)
	for k, v := range val {
		attributesList = append(attributesList, &authz.Attribute{
			Key:   k,
			Value: v.(string),
		})
	}
	return attributesList
}
