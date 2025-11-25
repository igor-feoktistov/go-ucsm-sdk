package util

import (
	"github.com/igor-feoktistov/go-ucsm-sdk/api"
	"github.com/igor-feoktistov/go-ucsm-sdk/mo"
)

func ChassisGet(c *api.Client, dn string) (chassis []mo.EquipmentChassis, err error) {
	var out mo.EquipmentChassisMos
	filter := api.FilterAnd{}
	if len(dn) > 0 {
		filter.Filters = append(filter.Filters, api.FilterEq{
			FilterProperty: api.FilterProperty {
				Class: "equipmentChassis",
				Property: "dn",
				Value: dn,
			},
		})
	}
	req := api.ConfigResolveClassRequest{
		Cookie: c.Cookie,
		ClassId: "equipmentChassis",
		InHierarchical: "false",
		InFilter: filter,
	}
	if err = c.ConfigResolveClass(req, &out); err == nil {
		chassis = out.EquipmentChassis
	}
	return
}
