package rest_component

import (
	"fmt"
	v1beta12 "halkyon.io/api/component/v1beta1"
	"halkyon.io/api/v1beta1"
	framework "halkyon.io/operator-framework"
	"halkyon.io/rest-component-capability/pkg/plugin"
	"k8s.io/apimachinery/pkg/runtime"
)

var _ framework.DependentResource = &component{}
var gvk = v1beta12.SchemeGroupVersion.WithKind(v1beta12.Kind)

type component struct {
	*framework.BaseDependentResource
}

func NewComponent(owner v1beta1.HalkyonResource) *component {
	config := framework.NewConfig(gvk)
	config.CheckedForReadiness = true
	config.CreatedOrUpdated = false
	p := &component{framework.NewConfiguredBaseDependentResource(owner, config)}
	return p
}

func (m *component) Name() string {
	c := plugin.OwnerAsCapability(m)
	paramsMap := plugin.ParametersAsMap(c.Spec.Parameters)
	return paramsMap["component"]
}

func (m *component) NameFrom(underlying runtime.Object) string {
	return framework.DefaultNameFrom(m, underlying)
}

func (m *component) Fetch() (runtime.Object, error) {
	panic("should never be called")
}

func (m *component) Build(empty bool) (runtime.Object, error) {
	c := &v1beta12.Component{}
	if !empty {
		panic("should not call Build to build a Component")
	}
	return c, nil
}

func (m *component) Update(_ runtime.Object) (bool, error) {
	return false, nil
}

func (m *component) GetCondition(underlying runtime.Object, err error) *v1beta1.DependentCondition {
	return framework.DefaultGetConditionFor(m, err)
}

func (m *component) GetDataMap() map[string][]byte {
	c := plugin.OwnerAsCapability(m)
	paramsMap := plugin.ParametersAsMap(c.Spec.Parameters)
	/*
		parameters:
			            - name: context
			              value: /api/fruits
			            - name: port
			              value: "8080"
		name: "ENDPOINT_BACKEND"
			      value: "http://fruit-backend-sb:8080/api/fruits"
	*/
	key := "ENDPOINT_BACKEND"
	if override, ok := paramsMap["endpointKey"]; ok {
		key = override
	}
	result := map[string][]byte{
		key: []byte(fmt.Sprintf("http://%s:%s%s", paramsMap["component"], paramsMap["port"], paramsMap["context"])),
	}
	return result
}

func (m *component) GetSecretName() string {
	return plugin.DefaultSecretNameFor(m)
}
