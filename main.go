package main

import (
	"github.com/claranet/terraform-provider-zabbix/zabbix"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	p := plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	}

	plugin.Serve(&p)
}
