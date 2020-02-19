module halkyon.io/rest-component-capability

go 1.13

require (
	github.com/appscode/jsonpatch v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/prometheus/client_golang v1.4.1 // indirect
	go.uber.org/zap v1.13.0 // indirect
	halkyon.io/api v1.0.0-rc.4.0.20200217221003-af8973318c2d
	halkyon.io/operator-framework v1.0.0-beta.4.0.20200219153202-7e58159b95d1
	k8s.io/apiextensions-apiserver v0.17.3 // indirect
	k8s.io/apimachinery v0.17.3
	sigs.k8s.io/testing_frameworks v0.1.2 // indirect
)

replace (
	github.com/appscode/jsonpatch => github.com/appscode/jsonpatch v0.0.0-20190108182946-7c0e3b262f30 // indirect
	k8s.io/api => k8s.io/api v0.0.0-20181213150558-05914d821849
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20181127025237-2b1284ed4c93
	k8s.io/client-go => k8s.io/client-go v0.0.0-20181213151034-8d9ed539ba31
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.1.10
)
