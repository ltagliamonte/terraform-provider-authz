package authz

import (
	authz "github.com/eko/authz/backend/pkg/authz"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	tfSet = &schema.Schema{
		Type:     schema.TypeSet,
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

func getSet(set *schema.Set) []string {
	var strList []string
	for _, k := range set.List() {
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

func flattenAttributesItems(attributes []*authz.Attribute) map[string]interface{} {
	if attributes != nil {
		res := make(map[string]interface{}, len(attributes))
		for _, attribute := range attributes {
			res[attribute.Key] = attribute.Value
		}
		return res
	}

	return make(map[string]interface{}, 0)
}
