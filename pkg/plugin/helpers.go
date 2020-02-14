package plugin

import (
	v1beta12 "halkyon.io/api/capability/v1beta1"
	"halkyon.io/api/v1beta1"
	framework "halkyon.io/operator-framework"
	"strings"
)

func OwnerAsCapability(res framework.DependentResource) *v1beta12.Capability {
	return res.Owner().(*v1beta12.Capability)
}

// Convert Array of parameters to a Map
func ParametersAsMap(parameters []v1beta1.NameValuePair) map[string]string {
	result := make(map[string]string)
	for _, parameter := range parameters {
		result[parameter.Name] = parameter.Value
	}
	return result
}

func DefaultSecretNameFor(secretOwner NeedsSecret) string {
	return strings.ToLower(secretOwner.Owner().GetName()) + "-config"
}

//getAppLabels returns an string map with the labels which wil be associated to the kubernetes/ocp resource which will be created and managed by this operator
func GetAppLabels(name string) map[string]string {
	return map[string]string{
		"app": name,
	}
}
