package rest_component

import (
	"fmt"
	"halkyon.io/api/capability/v1beta1"
	"halkyon.io/operator-framework"
	"halkyon.io/operator-framework/plugins/capability"
	"halkyon.io/operator-framework/util"
)

var _ capability.PluginResource = &Resource{}

type Resource struct {
	capability.SimplePluginResourceStem
}

func (m Resource) CheckValidity(owner framework.SerializableResource) []string {
	c := owner.(*v1beta1.Capability)
	parameters := util.ParametersAsMap(c.Spec.Parameters)
	missing := make([]string, 0, 3)
	missing = m.checkIfMissing(parameters, PortParam, missing)
	missing = m.checkIfMissing(parameters, ComponentParam, missing)
	missing = m.checkIfMissing(parameters, ContextParam, missing)
	if len(missing) > 0 {
		m.Logger.Info("validation", "missing", missing)
	}
	return missing
}

func (m Resource) checkIfMissing(parameters map[string]string, toCheck string, missing []string) []string {
	if _, ok := parameters[toCheck]; !ok {
		missing = append(missing, m.GetPrefixedValidationMessage(fmt.Sprintf("missing parameter %s", toCheck)))
	}
	return missing
}

func (m Resource) GetDependentResourcesWith(owner framework.SerializableResource) []framework.DependentResource {
	c := NewComponent(owner)
	config := framework.NewDefaultSecretConfig()
	config.CheckedForReadiness = true
	return []framework.DependentResource{framework.NewSecret(c, config), c}
}

func NewPluginResource() capability.PluginResource {
	return &Resource{capability.NewSimplePluginResourceStem("api", capability.TypeInfo{Type: "rest-component", Versions: []string{"1"}})}
}
