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

const (
	EndpointUrlKey        = "ENDPOINT_URL"
	EndpointOverrideParam = "halkyon.endpointKey"
	ComponentParam        = "component"
	PortParam             = "port"
	ContextParam          = "context"
)

type component struct {
	*framework.BaseDependentResource
	params map[string]string
}

func NewComponent(owner framework.SerializableResource) *component {
	config := framework.NewConfig(gvk)
	config.CheckedForReadiness = true
	config.Created = false
	config.Updated = true
	p := &component{BaseDependentResource: framework.NewConfiguredBaseDependentResource(owner, config)}
	c := ownerAsCapability(p)
	p.params = util.ParametersAsMap(c.Spec.Parameters)
	return p
}

func ownerAsCapability(res framework.DependentResource) *capability.Capability {
	return res.Owner().(*capability.Capability)
}

func (m *component) Name() string {
	return m.params["component"]
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
	key := EndpointUrlKey
	if override, ok := m.params[EndpointOverrideParam]; ok {
		key = override
	}
	result := map[string][]byte{
		key: []byte(fmt.Sprintf("http://%s:%s%s", m.params[ComponentParam], m.params[PortParam], m.params[ContextParam])),
	}
	return result
}

func (m *component) GetSecretName() string {
	return framework.DefaultSecretNameFor(m)
}
