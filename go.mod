module github.com/alauda/helm-crds

go 1.12

require (
	github.com/alauda/component-base v0.0.0-20190628064654-a4dafcfd3446
	github.com/fatih/structs v1.1.0
	github.com/ghodss/yaml v1.0.0
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/thoas/go-funk v0.4.0
	// helm.sh/helm v3.0.0-alpha.1.0.20190613170622-c35dbb7aabf8+incompatible
	helm.sh/helm/v3 v3.5.4
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/klog v0.4.0
	rsc.io/letsencrypt v0.0.3 // indirect
)

replace (
	github.com/deislabs/oras => github.com/deislabs/oras v0.11.0
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	github.com/russross/blackfriday => github.com/russross/blackfriday v1.5.2

	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.4
	k8s.io/apiserver => k8s.io/apiserver v0.20.4
	k8s.io/client-go => k8s.io/client-go v0.20.4
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.6.0
)
