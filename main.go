package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-archive/archive"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: archive.Provider})
}
