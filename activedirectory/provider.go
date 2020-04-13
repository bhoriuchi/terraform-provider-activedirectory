package activedirectory

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider terraform active directory provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Active Directory domain",
				DefaultFunc: schema.EnvDefaultFunc("MSAD_DOMAIN", nil),
			},
			"domain_controller": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Active Directory domain controller to connect to",
				DefaultFunc: schema.EnvDefaultFunc("MSAD_DOMAIN_CONTROLLER", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Active Directory username to use for authentication",
				DefaultFunc: schema.EnvDefaultFunc("MSAD_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Password for the Active Directory username provided",
				DefaultFunc: schema.EnvDefaultFunc("MSAD_PASSWORD", nil),
			},
		},
	}
}
