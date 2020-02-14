package main

import (
	"halkyon.io/rest-component-capability/pkg/plugin/rest-component"
	plugins "halkyon.io/operator-framework/plugins/capability"
)

func main() {
	plugins.StartPluginServerFor(rest_component.NewPluginResource())
}
