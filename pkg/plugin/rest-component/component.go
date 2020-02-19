package rest_component

import (
	"fmt"
	capability "halkyon.io/api/capability/v1beta1"
	v1beta12 "halkyon.io/api/component/v1beta1"
	"halkyon.io/api/v1beta1"
	framework "halkyon.io/operator-framework"
	"halkyon.io/operator-framework/util"
	"k8s.io/apimachinery/pkg/runtime"
)

var _ framework.DependentResource = &component{}
var gvk = v1beta12.SchemeGroupVersion.WithKind(v1beta12.Kind)

const EndpointUrlKey = "ENDPOINT_URL"

type component struct {
	*framework.BaseDependentResource
}

func NewComponent(owner v1beta1.HalkyonResource) *component {
	config := framework.NewConfig(gvk)
	config.CheckedForReadiness = true
	config.Created = false
	config.Updated = true
	p := &component{framework.NewConfiguredBaseDependentResource(owner, config)}
	return p
}

func ownerAsCapability(res framework.DependentResource) *capability.Capability {
	return res.Owner().(*capability.Capability)
}

func (m *component) Name() string {
	c := ownerAsCapability(m)
	paramsMap := util.ParametersAsMap(c.Spec.Parameters)
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

func (m *component) Update(toUpdate runtime.Object) (bool, runtime.Object, error) {
	return false, toUpdate, nil
}

func (m *component) GetCondition(underlying runtime.Object, err error) *v1beta1.DependentCondition {
	return framework.DefaultGetConditionFor(m, err)
}

func (m *component) GetDataMap() map[string][]byte {
	c := ownerAsCapability(m)
	paramsMap := util.ParametersAsMap(c.Spec.Parameters)
	key := EndpointUrlKey
	if override, ok := paramsMap["halkyon.endpointKey"]; ok {
		key = override
	}
	result := map[string][]byte{
		key: []byte(fmt.Sprintf("http://%s:%s%s", paramsMap["component"], paramsMap["port"], paramsMap["context"])),
	}
	return result
}

func (m *component) GetSecretName() string {
	return framework.DefaultSecretNameFor(m)
}
