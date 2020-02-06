package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider is an Terraform Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"cloudfunc_packer": resourcePacker(),
		},
	}
}
