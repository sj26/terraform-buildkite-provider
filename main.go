package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sj26/terraform-provider-buildkite/buildkite"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: buildkite.Provider,
	})
}
