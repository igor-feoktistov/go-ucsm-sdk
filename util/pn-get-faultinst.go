package util

import (
	"github.com/igor-feoktistov/go-ucsm-sdk/api"
	"github.com/igor-feoktistov/go-ucsm-sdk/mo"
)

func FaultInstancesGet(c *api.Client, pnDn string) (faultInsts []mo.FaultInst, err error) {
	var out mo.FaultInsts
	req := api.ConfigResolveChildrenRequest {
		    Cookie: c.Cookie,
		    InDn: pnDn,
		    ClassId: "faultInst",
		    InHierarchical: "true",
	}
	if err = c.ConfigResolveChildren(req, &out); err == nil {
		faultInsts = out.Instances
	}
	return
}
