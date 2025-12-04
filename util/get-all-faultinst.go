package util

import (
	"github.com/igor-feoktistov/go-ucsm-sdk/api"
	"github.com/igor-feoktistov/go-ucsm-sdk/mo"
)

func FaultInstancesGetAll(c *api.Client) (faultInsts []mo.FaultInst, err error) {
	var out mo.FaultInsts
        req := api.ConfigResolveClassRequest {
                    Cookie: c.Cookie,
                    ClassId: "faultInst",
                    InHierarchical: "false",
        }
        if err = c.ConfigResolveClass(req, &out); err == nil {
                faultInsts = out.Instances
        }
	return
}
