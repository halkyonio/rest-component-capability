package rest_component

import (
	beta1 "halkyon.io/api/v1beta1"
	"halkyon.io/operator-framework"
	"halkyon.io/operator-framework/plugins/capability"
)

var _ capability.PluginResource = &Resource{}

type Resource struct {
	capability.SimplePluginResourceStem
}

func (m Resource) GetDependentResourcesWith(owner beta1.HalkyonResource) []framework.DependentResource {
	c := NewComponent(owner)
	config := framework.NewDefaultSecretConfig()
	config.CheckedForReadiness = true
	return []framework.DependentResource{framework.NewSecret(c, config), c}
}

func NewPluginResource() capability.PluginResource {
	return &Resource{capability.NewSimplePluginResourceStem("api", capability.TypeInfo{Type: "rest-component", Versions: []string{"1"}})}
}
